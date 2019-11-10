package blobstore

import (
	//	"bytes"
	"fmt"
	//	"io"
	"log"
	"strings"
	"time"

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

func strLess(a, b string) bool { return a > b }

func AdvancedTransformExample(key string) *diskv.PathKey {
	path := strings.Split(key, "/")
	last := len(path) - 1
	return &diskv.PathKey{
		Path:     path[:last],
		FileName: path[last] + defaultBlobExt,
	}
}

// If you provide an AdvancedTransform, you must also provide its
// inverse:

func InverseTransformExample(pathKey *diskv.PathKey) (key string) {
	if 0 == len(pathKey.Path) {
		return strings.Join(pathKey.Path, "/")
	}
	bkv := pathKey.FileName[len(pathKey.FileName)-len(defaultBlobExt):]
	if bkv != ".blob" {
		panic("Invalid file found in storage folder!")
	}
	return strings.Join(pathKey.Path, "/") + pathKey.FileName[:len(pathKey.FileName)-len(defaultBlobExt)]
}

func Create(Base string, Index string) Store {
	log.Println("start create")
	log.Println(Base, Index)
	r := Store{}

	d := diskv.New(diskv.Options{
		BasePath:          Base,
		AdvancedTransform: AdvancedTransformExample,
		InverseTransform:  InverseTransformExample,
		CacheSizeMax:      1024 * 1024,
		Index:             &diskv.BTreeIndex{},
		IndexLess:         strLess,
	})

	r.BaseDir = Base
	r.IndexName = Index
	r.d = d

	return r
}

func (b *Store) Smash() {
	defer b.d.EraseAll()
}

func (b *Store) Store(blob []byte) int64 {
	log.Println("start store")

	now := time.Now()
	secs := now.UnixNano()
	key := fmt.Sprintf("%d", secs)
	if e := b.d.Write(key, blob); e != nil {
		log.Println("error ", e)
		return 0
	} else {
		return secs
	}
}

func (b *Store) Read(key int64) ([]byte, error) {
	k := fmt.Sprintf("%d", key)
	r, e := b.d.Read(k)
	if e != nil {
		log.Println("error ", e)
	}
	//	r := []byte("testing")
	//	return r
	//r := bytes.NewBufferString("testing")
	//return r.Bytes()
	return r, e
}

//func Del()

//func GetLastIndex()

//func GetAllIndex()
