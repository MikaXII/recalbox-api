package models

import (
	"gitlab.com/MikaXII/recalbox-api/utils"
)

type Hash struct {
	StringValue string
	HashType    string
	byteValue   []byte
}

func NewHash(filePath string, hashType string) *Hash {
	bytes := []byte{}
	value := ""
	switch hashType {
	case "MD5":
		bytes, value = utils.MD5(filePath)
		break
	case "SHA1":
		bytes, value = utils.SHA1(filePath)
		break
	case "CRC":
		bytes, value = utils.CRC32(filePath)
		break
	}
	return &Hash{byteValue: bytes, StringValue: value, HashType: hashType}
}

func SupportedHash() []string {
	return []string{"MD5", "SHA1", "CRC"}
}
