package mgrs

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/liucxer/resource-manage/config"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func FileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

type Result struct {
	Video   string
	Picture string
}

func (mgr *Mgr) Upload(path string) ([]Result, error) {
	var (
		res []Result
	)
	// 读取磁盘目录
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return []Result{}, err
	}

	for _, entry := range dirEntries {
		if entry.IsDir() {
			reply, err := mgr.Upload(path + "/" + entry.Name())
			if err != nil {
				return []Result{}, err
			}
			res = append(res, reply...)
			continue
		}
		// 计算md5sum
		filePath := path + "/" + entry.Name()
		md5Str, err := FileMD5(filePath)
		if err != nil {
			return []Result{}, err
		}

		// 生成封面
		coverPath := config.G_GlobalConfig.ResourcePath + "/picture/" + md5Str + ".jpg"
		cmd := "ffmpeg -i " + filePath + " -vframes 1 " + coverPath
		_, err = mgr.SyncExecute(cmd)
		if err != nil {
			return []Result{}, err
		}

		// 文件迁移
		bts, err := ioutil.ReadFile(filePath)
		if err != nil {
			return []Result{}, err
		}

		fileItems := strings.Split(filePath, ".")
		if err != nil {
			return []Result{}, err
		}
		fileExt := fileItems[len(fileItems)-1]
		videoPath := config.G_GlobalConfig.ResourcePath + "/video/" + md5Str + "." + fileExt
		err = ioutil.WriteFile(videoPath, bts, os.ModePerm)
		if err != nil {
			return []Result{}, err
		}

		res = append(res, Result{
			Video:   "/video/" + md5Str + "." + fileExt,
			Picture: "/picture/" + md5Str + ".jpg",
		})
	}

	return res, err
}
