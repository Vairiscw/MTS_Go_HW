package library

import "crypto/rand"

type idGenerator func(string) string

func IdGeneratorOne(key string) string {
	salt := rand.Text()
	return key + salt
}

func IdGeneratorTwo(key string) string {
	return "id_" + key
}
