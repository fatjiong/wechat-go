package wechat

//MixMessage 存放所有微信发送过来的消息和事件
type MixMessage struct {
	//CommonToken

	//基本消息
	MsgID        int64   `xml:"MsgId"`
	Content      string  `xml:"Content"`
	PicURL       string  `xml:"PicUrl"`
	MediaID      string  `xml:"MediaId"`
	Format       string  `xml:"Format"`
	ThumbMediaID string  `xml:"ThumbMediaId"`
	LocationX    float64 `xml:"Location_X"`
	LocationY    float64 `xml:"Location_Y"`
	Scale        float64 `xml:"Scale"`
	Label        string  `xml:"Label"`
	Title        string  `xml:"Title"`
	Description  string  `xml:"Description"`
	URL          string  `xml:"Url"`

	//事件相关
	Event     string `xml:"Event"`
	EventKey  string `xml:"EventKey"`
	Ticket    string `xml:"Ticket"`
	Latitude  string `xml:"Latitude"`
	Longitude string `xml:"Longitude"`
	Precision string `xml:"Precision"`

	MenuID string `xml:"MenuId"`

	//扫码事件
	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"`
		ScanResult string `xml:"ScanResult"`
	} `xml:"ScanCodeInfo"`

	//发图事件
	SendPicsInfo struct {
		Count   int32      `xml:"Count"`
		//PicList []EventPic `xml:"PicList>item"`
	} `xml:"SendPicsInfo"`

	//发送地理位置事件
	SendLocationInfo struct {
		LocationX float64 `xml:"Location_X"`
		LocationY float64 `xml:"Location_Y"`
		Scale     float64 `xml:"Scale"`
		Label     string  `xml:"Label"`
		Poiname   string  `xml:"Poiname"`
	}
}
