package model

import "time"

type Test struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name      string    `gorm:"column:name;unique"`
	Age       int       `gorm:"default:10"`
	CreatedAt int64     `gorm:"autoUpdateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt time.Time `grom:"column:updated_at" json:"updated_at"`
}
