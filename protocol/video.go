package protocol

type VideoCreateArgs struct {
}

type VideoCreateReply struct {
	Url string `json:"url"`
}

type MultipartVideoCreateArgs struct {
}

type Item struct {
	Video   string `json:"video"`
	Picture string `json:"picture"`
}

type MultipartVideoCreateReply struct {
	Items []Item `json:"items"`
}
