package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/masahiro331/gexto"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("required [ext4] argument")
	}

	fs, _ := gexto.NewFileSystem(os.Args[1])
	files, _ := fs.List()
	for _, filename := range files {
		targetFile := "os-release"
		if len(os.Args) == 3 {
			targetFile = os.Args[2]
		}
		if strings.Contains(filename, targetFile) {
			g, _ := fs.Open(filename)
			b, _ := ioutil.ReadAll(g)
			fmt.Println(string(b))
		}
	}
}
