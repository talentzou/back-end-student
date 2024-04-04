package notice

import "time"

type SysNotice struct {
	Id        uint      `json:"-" gorm:"primarykey"`
	Title     string    `json:"title" gorm:"size:512"`
	Author    string    `json:"author" gorm:"size:256"`
	Timestamp time.Time `json:"timestamp" gorm:"type:date"`
}
