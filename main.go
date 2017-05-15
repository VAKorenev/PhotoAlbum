// PhotoAlbum project main.go
package main

import (
	"PhotoAlbum/ini-master"
	"fmt"
	"io"
	"os"
)

type server struct {
	ip   string
	port string
}

type data struct {
	folder string
}

type html struct {
	title string
}

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
	c, err := ini.Load("PhotoAlbum.conf")
	if err != nil {
		fmt.Println("Загрузили конфиг:", &c)
	}
	//	path := c.Section("data").GetKey("folder").String()
	section, err := c.GetSection("server")
	fmt.Println(section.GetKey("ip"))
	//	readdir("./Новая папка")
}
