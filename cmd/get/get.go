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
	//	flatTransform := func(s string) []string { return []string{} }
	getIndex := "myindex"
	getKey := "1572018745"

	d := diskv.New(diskv.Options{
		BasePath:          "my-data-dir",
		AdvancedTransform: AdvancedTransformExample,
		InverseTransform:  InverseTransformExample,
		CacheSizeMax:      1024 * 1024,
		//		Index:        &diskv.BTreeIndex{},
		//		IndexLess:    strGreater,
	})

	if readVal, err := d.Read(getIndex + "/" + getKey); err != nil {
		fmt.Printf("\nread: %s", err)
	} else {
		fmt.Printf("\n%s", readVal)
	}

	getIndex = "rubbish"

	if readVal, err := d.Read(getIndex + "/" + getKey); err != nil {
		fmt.Printf("\nread: %s", err)
	} else {
		fmt.Printf("\n%s", readVal)
	}

}
