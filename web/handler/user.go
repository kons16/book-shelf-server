package handler

import (
	"github.com/kons16/book-shelf-server/usecase"
	"github.com/labstack/echo"
	"net/http"
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
	token, err := uh.userUC.Create()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, &userRes{
		token: token,
	})
}

// user関連のレスポンスに使用する構造体
type userRes struct {
	token string `json:"token"`
}
