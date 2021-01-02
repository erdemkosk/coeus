package models

type LoginInput struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type ConfigInput struct {
	Key   string      `json:"key"`
	Type  string      `json:"type"`
	Value interface{} ``
}

type AuthToken struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
