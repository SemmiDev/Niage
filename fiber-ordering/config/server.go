package config

import "os"

type server struct {
	Port string
}

func setupServer() *server {
	v := server{
		Port: os.Getenv("PORT"),
	}

	if v.Port == "" {
		v.Port = "9090"
	}

	return &v
}