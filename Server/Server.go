package Server

import (
	"fmt"
	"github.com/kardianos/service"
	"golang.org/x/sys/windows/svc/mgr"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Server struct {
	program
}

type program struct {
}

func (p *program) CheckServer() {
	manager, err := mgr.Connect()
	if err != nil {
		fmt.Println("权限不足 服务开启失败", err)
		return
	}

	defer manager.Disconnect()

	// Get the openService name from the user
	// Open the openService
	openService, err := manager.OpenService("tianfuan")
	if err != nil {
		// 服务打开失败 或 不存在
		fmt.Println("服务不存在", err)
		p.Install()
		return
	}

	defer openService.Close()

	// Check the openService status
	status, err := openService.Query()
	if err != nil {
		// 服务状态查询失败
		fmt.Println("服务状态查询失败 权限不足", err)
		return
	}

	// Print the openService status
	fmt.Println("服务状态", status.State)
}

func (p *program) run() {
	// 具体的服务实现
	select {
	case <-time.NewTimer(time.Minute * 1).C:
		p.IsProcessRunning()
	}
}

func (p *program) IsProcessRunning() {
	cmd := exec.Command("tasklist")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Check if the process name is in the list of running processes
	if strings.Contains(string(output), "tianfuan") {
		return
	} else {
		cmd = exec.Command("cmd", "/c", "%tianfuan%\\tianfuan.exe")
		err := cmd.Run()
		if err != nil {
			return
		}
	}
}

func (p *program) Install() {
	srvConfig := &service.Config{
		Name:        "tianfuan",
		DisplayName: "tianfuan保护程序",
		Description: "tianfuan安全",
	}

	prg := &program{}

	s, err := service.New(prg, srvConfig)
	if err != nil {
		fmt.Println(err)
	}

	err = s.Install()
	if err != nil {
		log.Fatal(err)
	}

	go s.Start()
	go s.Run()
}

func (p *program) Uninstall() {
	srvConfig := &service.Config{
		Name:        "tianfuan",
		DisplayName: "tianfuan保护程序",
		Description: "tianfuan安全",
	}

	prg := &program{}

	s, err := service.New(prg, srvConfig)
	if err != nil {
		fmt.Println(err)
	}

	err = s.Uninstall()
	if err != nil {
		log.Fatal(err)
	}
}

func (p *program) Start(s service.Service) error {
	fmt.Println("服务运行...")
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	return nil
}
