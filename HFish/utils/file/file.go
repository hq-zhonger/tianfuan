package file

import (
	"fmt"
	"os"
	"tianfuan/HFish/error"
)

func Output(result string, path string) {
	if path != "" {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			os.Mkdir("./scripts", os.ModePerm)
		}
		f_create, _ := os.Create(path)
		f_create.Close()
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		error.Check(err, "fail to open file")
		f.Write([]byte(result))
		f.Close()
	} else {
		fmt.Println(result)
	}
}
