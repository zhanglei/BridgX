package model

type Org struct {
	Base
	OrgName string `json:"org_name"`
}

func (u *Org) TableName() string {
	return "org"
}
