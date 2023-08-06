package protocol

type SetLimitHostArgs struct {
	LimitHosts string `json:"limit_hosts"`
	Enable     bool   `json:"enable"`
}

type SetLimitHostReply struct {
}

type GetLimitHostArgs struct {
}

type GetLimitHostReply struct {
	Enable     bool   `json:"enable"`
	LimitHosts string `json:"limit_hosts"`
}
