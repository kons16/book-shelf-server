package usecase

import (
	"fmt"
	"github.com/kons16/book-shelf-server/repository"
	"golang.org/x/crypto/bcrypt"
	"strconv"
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
func (uc *UserUseCase) Create(email string, password string) (string, error) {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	passwordHashStr := string(passwordHash)

	newUser, err := uc.userRepo.CreateUser(email, passwordHashStr)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	id := strconv.Itoa(newUser.ID)
	_, err = uc.userRepo.CreateToken(id)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	return "", nil
}
