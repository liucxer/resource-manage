package protocol

type VideoCreateArgs struct {
}

type VideoCreateReply struct {
	Url string `json:"url"`
}
