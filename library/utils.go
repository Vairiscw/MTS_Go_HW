package library

import "crypto/rand"

type IDGenerator func(string) string

func IDGeneratorAddRandomText(key string) string {
	salt := rand.Text()
	return key + salt
}

func IDGeneratorAddPrefix(key string) string {
	return "id_" + key
}
