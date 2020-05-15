package orm

type WeChat struct {
	UserID         int64
	OpenID         string
	UnionID        string
	IDCard         string
	IDCardFrontPic string
	IDCardBackPic  string
}

func (wc *WeChat) TableName() string {
	return "we_chat"
}
