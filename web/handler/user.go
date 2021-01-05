package handler

import (
	"github.com/kons16/book-shelf-server/usecase"
	"github.com/labstack/echo"
)

// UserHandler は /users 以下のエンドポイントを管理する構造体
type UserHandler struct {
	userUC *usecase.UserUseCase
}

// NewUserHandler は UserHandler のポインタを生成する関数
func NewUserHandler(userUC *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUC: userUC}
}

// Create は POST /user に対応するハンドラー
func (uh *UserHandler) Create(c echo.Context) error {
	nil, _ := uh.userUC.Create()
	return nil
}
