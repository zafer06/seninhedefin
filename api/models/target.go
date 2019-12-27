package models

import "time"

// Member struct
type Member struct {
	ID          int    `orm:"auto;column(id)"`
	Name        string `orm:"size(150)"`
	Surname     string `orm:"size(150)"`
	Email       string `orm:"size(150)"`
	Token       string `orm:"size(250)"`
	ExpiredDate time.Time
}

// Target struct
type Target struct {
	ID     int    `orm:"auto;column(id)"`
	Rank   int    `orm:"type(int)"`
	Target string `orm:"size(250)"`
}
