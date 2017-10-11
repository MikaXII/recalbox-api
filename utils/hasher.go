package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
)

// CRC32ToString get CRC32 string of a file
func CRC32ToString(fileDir string) string {

	file, err := os.Open(fileDir)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	hashCRC := crc32.NewIEEE()
	if _, err := io.Copy(hashCRC, file); err != nil {
		return ""
	}
	hashInBytes := hashCRC.Sum(nil)
	crcString := hex.EncodeToString(hashInBytes)

	return crcString
}

// MD5ToString get MD5 string of a file
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

// SHA1ToString get SHA1 string of a file
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

// CRC32 get CRC32 hash of a file
func CRC32(fileDir string) ([]byte, string) {

	file, err := os.Open(fileDir)
	if err != nil {
		return nil, ""
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	hashCRC := crc32.NewIEEE()
	if _, err := io.Copy(hashCRC, file); err != nil {
		return nil, ""
	}
	hashInBytes := hashCRC.Sum(nil)
	crcString := hex.EncodeToString(hashInBytes)

	return hashInBytes, crcString
}

// MD5 get MD5 hash of a file
func MD5(fileDir string) ([]byte, string) {
	file, err := os.Open(fileDir)
	if err != nil {
		fmt.Println(err)
		return nil, ""
	}
	hashMD5 := md5.New()
	if _, err := io.Copy(hashMD5, file); err != nil {
		return nil, ""
	}

	hashInBytes := hashMD5.Sum(nil)
	md5String := hex.EncodeToString(hashInBytes)
	return hashInBytes, md5String
}

// SHA1 get SHA1 hash of a file
func SHA1(fileDir string) ([]byte, string) {

	file, err := os.Open(fileDir)
	if err != nil {
		fmt.Println(err)
		return nil, ""
	}

	hashSHA1 := sha1.New()
	if _, err := io.Copy(hashSHA1, file); err != nil {
		return nil, ""
	}
	hashInBytes := hashSHA1.Sum(nil)
	sha1String := hex.EncodeToString(hashInBytes)
	return hashInBytes, sha1String
}
