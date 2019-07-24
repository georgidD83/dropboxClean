package main

import (
	"fmt"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/users"
)

var (
	token = "-e2kZ4ViybAAAAAAAAAL2VaU2GHPfTciP60B56RJgvEVxzK2G9X5bR5Tg1reyvHp"
)

func main() {
	config := dropbox.Config{
		Token:    token,
		LogLevel: dropbox.LogInfo, // if needed, set the desired logging level. Default is off
	}
	dbx := users.New(config)
	// start making API calls

	fmt.Printf("%+v\n", dbx)
}
