package main

import (
	"fmt"

	"github.com/d-james-projects/blobstore"
)

func main() {

	/*
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
	*/
	/*
		for _, key := range d.Index.Keys("myindex/", 2) {
			keys = append(keys, key)
		}

		fmt.Println(keys)*/
	val := []byte("testing")
	b := blobstore.Create("my-base-dir/index", "my-index")
	key := b.Store(val)
	fmt.Println("stored key ", key)
	read, _ := b.Read(key)
	fmt.Println("read back ", read)
}
