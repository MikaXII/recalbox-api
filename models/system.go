package models

import "os"

type System struct {
	Name string
}

type Rom struct {
	name     string
	hashList []Hash
}

type RomInfo struct {
	Name         string
	HashListInfo []Hash
}

func NewRom(filepath string, file os.FileInfo, hash []string) *Rom {
	hashList := []Hash{}
	for _, h := range hash {
		hashList = append(hashList, NewHash(filepath, h))
	}
	return &Rom{name: file.Name(), hashList: hashList}
}

func (r *Rom) Hash(hashType string) *Hash {
	for _, h := range r.hashList {
		if h.HashType == hashType {
			return &h
		}
	}
	return &Hash{}
}

func (r *Rom) HashList() []Hash {
	return r.hashList
}

func (r *Rom) Name() string {
	return r.name
}

func (r *Rom) Info() *RomInfo {
	return &RomInfo{Name: r.name, HashListInfo: r.HashList()}
}
