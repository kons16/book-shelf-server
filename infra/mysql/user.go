package MySQL

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kons16/book-shelf-server/domain/entity"
)

// UserRepository は repository.UserRepository を満たす構造体
type UserRepository struct {
	dbMap *gorm.DB
}

// NewUserRepository は UserRepository のポインタを生成する
func NewUserRepository(dbMap *gorm.DB) *UserRepository {
	return &UserRepository{dbMap: dbMap}
}

// CreateUser はユーザーを新規登録するSQLを記述
func (ur *UserRepository) CreateUser(email string, passwordHash string) (*entity.User, error) {
	user := &entity.User{Email: email, PasswordHash: passwordHash}
	r := ur.dbMap.Create(user)
	if r.Error != nil {
		fmt.Println(r.Error)
		return nil, r.Error
	}

	return &entity.User{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

// CreateToken は新規登録・ログインしたとき、token を新しく発行し保存する
func (ur *UserRepository) CreateToken(id int) (string, error) {
	auth := entity.GenerateToken()
	userSession := &entity.UserSession{UserId: id, Token: auth.Token}

	r := ur.dbMap.Create(userSession)
	if r.Error != nil {
		fmt.Println(r.Error)
		return "", r.Error
	}

	return auth.Token, nil
}

type userDTO struct {
	ID    int
	Email string
}
