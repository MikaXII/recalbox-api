package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func CRC32ToString(fileDir string) string {

	file, err := os.Open(fileDir)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	hashCRC := crc32.NewIEEE()
	if _, err := io.Copy(hashCRC, file); err != nil {
		return ""
	}
	hashInBytes := hashCRC.Sum(nil)
	crcString := hex.EncodeToString(hashInBytes)

	return crcString
}

func MD5ToString(fileDir string) string {

	file, err := os.Open(fileDir)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	hashMD5 := md5.New()
	if _, err := io.Copy(hashMD5, file); err != nil {
		return ""
	}

	hashInBytes := hashMD5.Sum(nil)
	md5String := hex.EncodeToString(hashInBytes)

	return md5String
}

func SHA1ToString(fileDir string) string {

	file, err := os.Open(fileDir)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	hashSHA1 := sha1.New()
	if _, err := io.Copy(hashSHA1, file); err != nil {
		return ""
	}
	hashInBytes := hashSHA1.Sum(nil)
	sha1String := hex.EncodeToString(hashInBytes)
	return sha1String
}
