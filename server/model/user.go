package model

type User struct {
	// 用户信息数据结构
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
