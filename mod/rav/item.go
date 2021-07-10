package rav

import (
	"time"
)

type Item struct {
	Name        string          `json:"name" form:"name"`
	Slug        string          `json:"slug" form:"slug"`
	Rank        int             `json:"rank" form:"rank"`
	Platform    string          `json:"platform" form:"platform"`
	Description string          `json:"description" form:"description"`
	Published   bool            `json:"published" form:"published"`
	Selected    bool            `json:"selected" form:"selected"`
	Favorite    bool            `json:"fav" form:"favorite"`
	Checked     map[string]bool `json:"checked"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
