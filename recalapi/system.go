package recalapi

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

// System represent a system...
type System struct {
	Name string
}

// Rom represents a rom file
type Media struct {
	Name     string
	Filepath string
	HashList []Hash
}

type Hash struct {
	HashType string
	Value    string
}

// NewRom create a new rom with filepath, file and hash
func NewMedia(filepath string, file os.FileInfo) *Media {
	rom := &Media{Name: file.Name(), Filepath: filepath}
	rom.SetHash()
	return rom
}

func (r *Media) SetHash() {
	r.HashList = append(r.HashList, Hash{HashType: "MD5", Value: r.MD5()})
	r.HashList = append(r.HashList, Hash{HashType: "SHA1", Value: r.SHA1()})
	r.HashList = append(r.HashList, Hash{HashType: "CRC", Value: r.CRC()})
}

func (r *Media) MD5() string {
	file, err := os.Open(r.Filepath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	hashMD5 := md5.New()
	if _, err := io.Copy(hashMD5, file); err != nil {
		return ""
	}
	return hex.EncodeToString(hashMD5.Sum(nil))
}

func (r *Media) SHA1() string {
	file, err := os.Open(r.Filepath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	hashSHA1 := sha1.New()
	if _, err := io.Copy(hashSHA1, file); err != nil {
		return ""
	}
	return hex.EncodeToString(hashSHA1.Sum(nil))
}

func (r *Media) CRC() string {
	file, err := os.Open(r.Filepath)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	hashCRC := crc32.NewIEEE()

	if _, err := io.Copy(hashCRC, file); err != nil {
		return ""
	}

	return hex.EncodeToString(hashCRC.Sum(nil))
}
