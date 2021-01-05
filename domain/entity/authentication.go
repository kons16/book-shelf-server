package entity

import "crypto/rand"

// ユーザーの認証周りの処理を行うエンティティ
type Auth struct {
}

func GenerateToken(string) string {
	table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_@"
	l := len(table)
	ret := make([]byte, 128)
	src := make([]byte, 128)
	rand.Read(src)
	for i := 0; i < 128; i++ {
		ret[i] = table[int(src[i])%l]
	}

	return string(ret)
}
