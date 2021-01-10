package main

import (
	"fmt"
	"github.com/kons16/book-shelf-server/usecase"
	"github.com/kons16/book-shelf-server/web"
	"os"

	"github.com/kons16/book-shelf-server/infra/MySQL"
)

func main() {
	dbMap, err := MySQL.NewMySQLDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	userRepo := MySQL.NewUserRepository(dbMap)

	// ここで Repository の構造体と interface を紐付けさせる
	userUC := usecase.NewUserUseCase(userRepo)

	s := web.NewServer(userUC)

	if err := s.Start(":" + "8000"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
