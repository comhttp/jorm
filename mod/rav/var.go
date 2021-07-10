package rav

import (
	"github.com/comhttp/jorm/pkg/utl"
)

type Media struct {
	WebSite     []string   `json:"web"`
	Explorer    []string   `json:"explorer"`
	Chat        []string   `json:"chat"`
	Twitter     string     `json:"tw"`
	Facebook    string     `json:"facebook"`
	Telegram    string     `json:"telegram"`
	Reddit      string     `json:"reddit"`
	Github      []string   `json:"github"`
	BitcoinTalk string     `json:"bitcointalk"`
	WhitePaper  string     `json:"whitepaper"`
	isLogo      bool       `json:"islogo" form:"islogo"`
	Logo        utl.Images `json:"logo" form:"logo"`
}

//User struct declaration
type User struct {
	ID       string
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Gender   string `json:"Gender"`
	Password string `json:"Password"`
}
