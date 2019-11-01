package blobstore

import (
	"bytes"
	"io"
	"log"

	"github.com/peterbourgon/diskv"
)

const (
	defaultBaseDir = "data-dir"
	defaultBlobExt = ".blob"
)

type Store struct {
	BaseDir   string
	IndexName string
	d         *diskv.Diskv
}

func Create(Base string, Index string) Store {
	log.Println("start create")
	r := Store{}
	return r
}

func (b *Store) Store(i io.Reader) {
	log.Println("start store")
}

func (b *Store) Read(key int64) []byte {
	//	r := []byte("testing")
	//	return r
	r := bytes.NewBufferString("testing")
	return r.Bytes()
}

//func Del()

//func GetLastIndex()

//func GetAllIndex()
