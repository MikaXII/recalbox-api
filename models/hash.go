package models

import (
	"gitlab.com/MikaXII/recalbox-api/utils"
)

// Hash represent a hash for a rom file
type Hash struct {
	StringValue string
	HashType    string
	byteValue   []byte
}

// NewHash Get a hash for a specific file
func NewHash(filePath string, hashType string) *Hash {
	bytes := []byte{}
	value := ""
	switch hashType {
	case "MD5":
		bytes, value = utils.MD5(filePath)

	case "SHA1":
		bytes, value = utils.SHA1(filePath)

	case "CRC":
		bytes, value = utils.CRC32(filePath)

	}
	return &Hash{byteValue: bytes, StringValue: value, HashType: hashType}
}

// SupportedHash list of hash currently supported
func SupportedHash() []string {
	return []string{"MD5", "SHA1", "CRC"}
}
