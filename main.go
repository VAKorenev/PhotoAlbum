// PhotoAlbum project main.go
package main

import (
	"fmt"
	"ini-master"
	"io"
	"os"
)

type server struct {
	ip   string
	port string
}

func (s *server) load(f *ini.File) {
	ip, err := f.Section("server").GetKey("ip")
	if err != nil {
		s.ip = "127.0.0.1"
	}
	s.ip = ip.String()
	port, err := f.Section("server").GetKey("port")
	if err != nil {
		s.port = "8080"
	}
	s.port = port.String()
}

type data struct {
	folder string
}

func (d *data) load(f *ini.File) {
	folder, err := f.Section("data").GetKey("folder")
	if err != nil {
		d.folder = "."
	}
	d.folder = folder.String()
}

type html struct {
	title string
}

func (h *html) load(f *ini.File) {
	title, err := f.Section("html").GetKey("title")
	if err != nil {
		h.title = "Моя фотогаллерея"
	}
	h.title = title.String()
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
		panic(err)
	}
	s := new(server)
	s.load(c)
	d := new(data)
	d.load(c)
	h := new(html)
	h.load(c)
	readdir(d.folder)
	fmt.Println(s.ip, s.port, d.folder, h.title)
}
