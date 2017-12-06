package models

type WeixinUser struct {
	ID int64

	Tokens  string
	HeadImg string
	Name    string
	Gender  int32

	Psn    string
	Stream string
	xbox   string

	VipDate int64
}
