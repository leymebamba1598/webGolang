package models

type User struct {
	Id       int    `xml:"id"`
	UserName string `xml:"username"`
	Password string `xml:"password"`
}
