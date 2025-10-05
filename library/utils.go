package library

import "crypto/rand"

type IdGenerator func(string) string

func IdGeneratorAddRandomText(key string) string {
	salt := rand.Text()
	return key + salt
}

func IdGeneratorAddPrefix(key string) string {
	return "id_" + key
}
