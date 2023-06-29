package main

import (
	"github.com/flopp/go-findfont"
	"log"
	"os"
	"path/filepath"
	"strings"
	"tianfuan/Gui"
)

func init() {
	// 设置中文环境变量
	paths := findfont.List()
	for _, path := range paths {
		// msyh.ttc
		if strings.Contains(path, "msyh.ttc") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}

	// 设置软件 运行目录
	abs, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	err := os.Chdir(abs)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var MyApp Gui.App
	MyApp.ShowGui()
}
