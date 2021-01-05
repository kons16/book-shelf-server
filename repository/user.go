package repository

import "github.com/kons16/book-shelf-server/domain/entity"

// User はユーザの永続化を担当するリポジトリ
type User interface {
	Create(user *entity.User) error
}
