package main

import (
	"fmt"

	"github.com/d-james-projects/blobstore"
)

func main() {

	val := []byte("testing")
	b := blobstore.Create("my-base-dir", "test-index")
	key := b.Store(val)
	fmt.Println("stored key ", key)
	read, _ := b.Read(key)
	fmt.Println("read back ", read)

	all := b.GetAllIndex()
	fmt.Println(all)

	latest := b.GetLatestIndex()
	fmt.Println(latest)
	l, _ := b.Read(key)
	fmt.Println("read back latest ", l)

	c := blobstore.Create("my-base-dir", "a-test-index")
	k := c.Store(val)
	fmt.Println("stored key ", k)

}
