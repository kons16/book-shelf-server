package MySQL

import (
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

// Create はユーザーを新規登録するSQLを記述
func (ur *UserRepository) Create(user *entity.User) (string, error) {
	r := ur.dbMap.Create(&user)
	if r.Error != nil {
		return "", r.Error
	}

	// auth := entity.GenerateToken()
	// token := auth.Token

	return "", nil
}
