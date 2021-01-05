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

// Create は ユーザーを新規登録するSQLを記述
func (ur *UserRepository) Create(user *entity.User) error {
	return nil
}
