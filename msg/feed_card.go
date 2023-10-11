package msg

type FeedCardLink struct {
	Title      string `json:"title"`
	MessageUrl string `json:"messageURL"`
	PicUrl     string `json:"picURL"`
}
type FeedCard struct {
	msgType
	FeedCard struct {
		Links []FeedCardLink `json:"links"`
	} `json:"feedCard"`
}

func (f *FeedCard) SetFeedCardLinks(links []FeedCardLink) {
	f.FeedCard.Links = links
}
func (f *FeedCard) SetMsgType() {
	f.MsgType = "feedCard"
}

func NewFeedCardLink() *FeedCard {
	return &FeedCard{}
}
