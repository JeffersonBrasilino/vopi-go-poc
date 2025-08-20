package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vopi-go-poc/internal/chat"
	"github.com/vopi-go-poc/internal/core"
)

func main() {
	dbport,_:=strconv.Atoi(os.Getenv("DB_PORT"))
	dbConn := core.NewPostgresConnection(
		os.Getenv("DB_HOST"),
		dbport,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)

	r := gin.Default()
	chat.NewChatModule(dbConn).WithHttp(r)
	r.Run(":3000")
}