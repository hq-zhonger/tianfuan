package Regedit

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
	"os"
)

type Regedit struct {
}

func (r *Regedit) CheckRegedit() {
	key, err := registry.OpenKey(registry.CURRENT_USER, "Environment", registry.QUERY_VALUE)
	defer key.Close()

	if err != nil {
		log.Fatal(err)
	}

	value, _, err := key.GetStringValue("tianfuan")
	if err != nil {
		r.AddRegedit()
	}

	dir, _ := os.Getwd()
	if value != dir {
		r.AddRegedit()
	}
}

func (r *Regedit) AddRegedit() {
	key, err := registry.OpenKey(registry.CURRENT_USER, "Environment", registry.ALL_ACCESS)
	defer key.Close()
	if err != nil {
		fmt.Println(4)
		log.Fatal(err)
	}

	dir, _ := os.Getwd()
	err = key.SetStringValue("tianfuan", dir)
	if err != nil {
		fmt.Println(5)
		log.Fatal(err)
	}

	r.CheckRegedit()
}

func (r *Regedit) DeleteRegedit() {
	key, err := registry.OpenKey(registry.CURRENT_USER, "Environment", registry.ALL_ACCESS)
	defer key.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = key.DeleteValue("tianfuan")
	if err != nil {
		return
	}
}
