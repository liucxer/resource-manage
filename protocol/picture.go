package protocol

type PictureCreateArgs struct {
}

type PictureCreateReply struct {
	Url string `json:"url"`
}
