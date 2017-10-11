package models

import "os"

// System represent a system...
type System struct {
	Name string
}

// Rom represents a rom file
type Rom struct {
	name     string
	hashList []Hash
}

// RomInfo represents a exported info of a rom file
type RomInfo struct {
	Name         string
	HashListInfo []Hash
}

// NewRom create a new rom with filepath, file and hash
func NewRom(filepath string, file os.FileInfo, hash []string) *Rom {
	hashList := []Hash{}
	for _, h := range hash {
		hashList = append(hashList, *NewHash(filepath, h))
	}
	return &Rom{name: file.Name(), hashList: hashList}
}

// Hash get specified hash for a rom file
func (r *Rom) Hash(hashType string) *Hash {
	for _, h := range r.hashList {
		if h.HashType == hashType {
			return &h
		}
	}
	return &Hash{}
}

// HashList get list of hash for a rom file
func (r *Rom) HashList() []Hash {
	return r.hashList
}

// Name get name of a rom file
func (r *Rom) Name() string {
	return r.name
}

// Info get all info for a rom
func (r *Rom) Info() *RomInfo {
	return &RomInfo{Name: r.name, HashListInfo: r.HashList()}
}
