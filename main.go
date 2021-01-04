package main

import (
	"fmt"
	"os"

	"github.com/kons16/book-shelf-server/infra/MySQL"
)

func main() {
	_, err := MySQL.NewMySQLDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
