package model

type PwdListRes struct {
	page  int
	count int
	list  []PwdRes
}

type PwdRes struct {
	Id          uint
	Company     string
	Domain      string
	AccountInfo string
	Account     string
}
