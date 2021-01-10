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

// CreateUser はユーザーを新規登録するSQLを記述
func (ur *UserRepository) CreateUser(email string, passwordHash string) (*entity.User, error) {
	//r := ur.dbMap.Create()
	//if r.Error != nil {
	//	return "", r.Error
	//}

	// auth := entity.GenerateToken()
	// token := auth.Token

	return nil, nil
}

// CreateToken は新規登録・ログインしたとき、token を新しく発行し保存する
func (ur *UserRepository) CreateToken(id string) (string, error) {
	return "", nil
}
