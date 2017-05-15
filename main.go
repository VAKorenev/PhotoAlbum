// PhotoAlbum project main.go
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func readdir(dir string) {
	dh, _ := os.Open(dir)
	defer dh.Close()
	for {
		f, err := dh.Readdir(10)
		if err == io.EOF {
			break
		}
		for _, fi := range f {
			fmt.Printf("%s/%s\n", dir, fi.Name())
			if fi.IsDir() {
				readdir(dir + "/" + fi.Name())
			}
		}
	}
}

func main() {

	readdir(".")

}
