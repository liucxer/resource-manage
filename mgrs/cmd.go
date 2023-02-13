package mgrs

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/go-cmd/cmd"
	"github.com/sirupsen/logrus"
)

type Mgr struct {
}

func NewMgr() *Mgr {
	return &Mgr{}
}

var GlobalMgr = NewMgr()

var LineSeparator string

func init() {
	switch runtime.GOOS {
	case "windows":
		LineSeparator = "\r\n"
	case "android", "darwin", "freebsd", "linux", "netbsd", "openbsd", "solaris":
		LineSeparator = "\n"
	}
}

type CmdStatus struct {
	Cmd         string   `json:"cmd"`
	PID         int      `json:"pid"`
	Complete    bool     `json:"complete"`                   // false if stopped or signaled
	Exit        int      `json:"exit"`                       // exit code of process
	Error       error    `json:"error" swaggertype:"string"` // Go error
	StartTs     int64    `json:"startTs"`                    // Unix ts (nanoseconds), zero if Cmd not started
	StopTs      int64    `json:"stopTs"`                     // Unix ts (nanoseconds), zero if Cmd not started or running
	Runtime     float64  `json:"runtime"`                    // seconds, zero if Cmd not started
	StdoutLines []string `json:"stdout_lines"`
	StderrLines []string `json:"stderr_lines"`
	Stderr      string   `json:"stderr"`
	Stdout      string   `json:"stdout"`
}

var (
	DefaultTimeOut = int64(30)
)

type Option func(opts *Options)

type Options struct {
	Timeout int64
}

func WithTimeoutOption(timeout int64) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

func loadOptions(options ...Option) *Options {
	opts := new(Options)
	for _, option := range options {
		option(opts)
	}
	return opts
}

func (mgr *Mgr) SyncExecute(cmdStr string, options ...Option) (*CmdStatus, error) {
	var (
		res     CmdStatus
		err     error
		timeout int64
	)
	// 开始时间
	startTime := time.Now()

	opts := loadOptions(options...)
	if opts.Timeout != 0 {
		timeout = opts.Timeout
	} else {
		timeout = DefaultTimeOut
	}

	command := cmd.NewCmd("sh", []string{"-c", cmdStr}...)
	go func() {
		<-time.After(time.Duration(timeout) * time.Second)
		err = command.Stop()
		if err != nil {
			logrus.Errorf("command.Stop err:%v. cmdStr:%s", err, cmdStr)
		}
		err = errors.New(fmt.Sprintf("Mgr SyncExecute timeout:%d. cmdStr:%s", timeout, cmdStr))
	}()

	cmdStatus := <-command.Start()
	if cmdStatus.Error != nil {
		// 执行时间
		latencyTime := time.Now().Sub(startTime)
		logrus.Errorf("SyncExecute command.Start. cmdStr:%s, err:%v. cmdStatus:%+v, cost:%0.2fs",
			cmdStr, cmdStatus.Error, cmdStatus, float64(latencyTime)/float64(time.Second))
		if err != nil { // timeout
			return nil, err
		}
		return nil, cmdStatus.Error
	}

	// 执行时间
	latencyTime := time.Now().Sub(startTime)
	logrus.Debugf("SyncExecute success. cost:%0.2fs, cmdStr:%s, cmdStatus:%+v",
		float64(latencyTime)/float64(time.Second), cmdStr, cmdStatus)

	res.Cmd = cmdStr
	res.PID = cmdStatus.PID
	res.Complete = cmdStatus.Complete
	res.Exit = cmdStatus.Exit
	res.Error = cmdStatus.Error
	res.StartTs = cmdStatus.StartTs
	res.StopTs = cmdStatus.StopTs
	res.Runtime = cmdStatus.Runtime
	res.StdoutLines = cmdStatus.Stdout
	res.StderrLines = cmdStatus.Stderr

	stdout := strings.Join(RemoveEmptyElm(cmdStatus.Stdout), LineSeparator)
	stderr := strings.Join(RemoveEmptyElm(cmdStatus.Stderr), LineSeparator)
	res.Stdout = stdout
	res.Stderr = stderr

	return &res, cmdStatus.Error
}

// RemoveEmptyElm 数组去除空白元素
func RemoveEmptyElm(arr []string) []string {
	var ret []string

	if len(arr) == 0 {
		return ret
	}

	for _, val := range arr {
		if len(val) != 0 {
			ret = append(ret, val)
		}
	}
	return ret
}

func (mgr *Mgr) SyncExecuteToJson(response interface{}, cmdStr string, options ...Option) (*CmdStatus, error) {
	var (
		res *CmdStatus
		err error
	)
	res, err = mgr.SyncExecute(cmdStr, options...)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(res.Stdout), &response)
	if err != nil {
		logrus.Errorf("SyncExecuteToJson json.Unmarshal err:%v. cmdStatus:%+v", err, res)
		return nil, err
	}
	logrus.Debugf("SyncExecuteToJson success. cmdStr:%s, cmdStatus:%+v", cmdStr, res)

	return res, err
}
