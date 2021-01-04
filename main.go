package main

import (
	"fmt"
	"github.com/kons16/book-shelf-server/web"
	"os"

	"github.com/kons16/book-shelf-server/infra/MySQL"
)

func main() {
	_, err := MySQL.NewMySQLDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	s := web.NewServer()

	if err := s.Start(":" + "8000"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
