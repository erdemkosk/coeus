package models

type ServerConfiguration struct {
	Server struct {
		Port int
	}
	Redis struct {
		Host     string
		Port     int
		Password string
	}
}
