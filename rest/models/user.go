package models

type User struct {
	Id       int    `yaml:"id"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}
