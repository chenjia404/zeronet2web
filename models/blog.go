package models

type Blog struct {
	Id          int    `gorm:"primary_key"`
	Address     string `json:"address" gorm:"index:unique"`
	Modified    int64  `json:"modified"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Links       string `json:"links"`
}
