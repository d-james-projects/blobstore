package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/peterbourgon/diskv"
)

func strLess(a, b string) bool { return a < b }

func AdvancedTransformExample(key string) *diskv.PathKey {
	path := strings.Split(key, "/")
	last := len(path) - 1
	return &diskv.PathKey{
		Path:     path[:last],
		FileName: path[last] + ".bkv",
	}
}

// If you provide an AdvancedTransform, you must also provide its
// inverse:

func InverseTransformExample(pathKey *diskv.PathKey) (key string) {
	if 0 == len(pathKey.Path) {
		return strings.Join(pathKey.Path, "/")
	}
	bkv := pathKey.FileName[len(pathKey.FileName)-4:]
	if bkv != ".bkv" {
		panic("Invalid file found in storage folder!")
	}
	return strings.Join(pathKey.Path, "/") + pathKey.FileName[:len(pathKey.FileName)-4]
}

func main() {
	//	keys := []string{}
	//	Index := diskv.BTreeIndex{}

	d := diskv.New(diskv.Options{
		BasePath:          "my-data-dir",
		AdvancedTransform: AdvancedTransformExample,
		InverseTransform:  InverseTransformExample,
		CacheSizeMax:      1024 * 1024,
		Index:             &diskv.BTreeIndex{},
		IndexLess:         strLess,
	})
	// Write some text to the key "myindex/<time>".
	now := time.Now()
	secs := now.Unix()
	key := fmt.Sprintf("myindex/%d", secs)
	d.WriteString(key, "¡Hola!") // will be stored in "<basedir>/myindex/<time>.bkv"
	//fmt.Println(d.ReadString("alpha/beta/gamma"))

	time.Sleep(1500 * time.Millisecond)
	now = time.Now()
	secs = now.Unix()
	key = fmt.Sprintf("myindex/%d", secs)
	d.WriteString(key, "¡Hola!") // will be stored in "<basedir>/myindex/<time>.bkv"

	time.Sleep(1500 * time.Millisecond)
	now = time.Now()
	secs = now.Unix()
	key = fmt.Sprintf("myindex/%d", secs)
	d.WriteString(key, "¡Hola!") // will be stored in "<basedir>/myindex/<time>.bkv"
	/*
		for _, key := range d.Index.Keys("myindex/", 2) {
			keys = append(keys, key)
		}

		fmt.Println(keys)*/
}
