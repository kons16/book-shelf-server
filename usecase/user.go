package usecase

import (
	"github.com/kons16/book-shelf-server/repository"
)

// UserUseCase はユーザに関係するアプリケーションロジックを担当する構造体
type UserUseCase struct {
	userRepo repository.User
}

// NewUserUseCase はUserUseCaseのポインタを生成する関数。引数は *MySQL.UserRepository が渡される。
func NewUserUseCase(userRepo repository.User) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

// Create は新規でユーザーを作成するユースケース
func (u *UserUseCase) Create() (string, error) {
	return "", nil
}
