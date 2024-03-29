package main

import (
	"fmt"
	"strings"

	"github.com/peterbourgon/diskv"
)

func strLess(a, b string) bool { return a < b }

func strGreater(a, b string) bool { return a > b }

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
		//		return strings.Join(pathKey.Path, "/")
		return ""
	}
	bkv := pathKey.FileName[len(pathKey.FileName)-4:]
	if bkv != ".bkv" {
		panic("Invalid file found in storage folder!")
	}
	return strings.Join(pathKey.Path, "/") + "/" + pathKey.FileName[:len(pathKey.FileName)-4]
}

func main() {
	flatTransform := func(s string) []string { return []string{} }

	d := diskv.New(diskv.Options{
		BasePath: "my-data-dir/myindex",
		/*AdvancedTransform: AdvancedTransformExample,
		InverseTransform:  InverseTransformExample,*/
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
		Index:        &diskv.BTreeIndex{},
		IndexLess:    strGreater,
	})

	for k, got := range d.Index.Keys("", 1) {
		fmt.Println(k, got)
	}
	/*
		for _, key := range d.Index.Keys("myindex/", 2) {
			keys = append(keys, key)
		}

		fmt.Println(keys)*/

	rkey := "myindex/1572022944"
	if readVal, err := d.Read(rkey); err != nil {
		fmt.Printf("\nread: %s", err)
	} else {
		fmt.Printf("\n%s", readVal)
	}

	rkey = "myindex/1572022941"
	if readVal, err := d.Read(rkey); err != nil {
		fmt.Printf("\nread: %s", err)
	} else {
		fmt.Printf("\n%s", readVal)
	}

	rkey = "myindex/rubbish"
	if readVal, err := d.Read(rkey); err != nil {
		fmt.Printf("\nread: %s", err)
	} else {
		fmt.Printf("\n%s", readVal)
	}

}
