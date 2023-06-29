package main

/*
   #include <windows.h>
   BOOL SelfDelete();
   //----------------------------------------------------------------
   int DelSelf()
   {
       SelfDelete();
       exit(0);
       return 0;
   }
   //---------------------------------------------------------------
   BOOL SelfDelete()
   {
    TCHAR szModule [MAX_PATH],
     szComspec[MAX_PATH],
     szParams [MAX_PATH];

    // get file path names:
    if((GetModuleFileName(0,szModule,MAX_PATH)!=0) &&
     (GetShortPathName(szModule,szModule,MAX_PATH)!=0) &&
     (GetEnvironmentVariable("COMSPEC",szComspec,MAX_PATH)!=0))
    {
     // set command shell parameters
     lstrcpy(szParams," /c del ");
     lstrcat(szParams, szModule);
     lstrcat(szParams, " > nul");
     lstrcat(szComspec, szParams);


     // set struct members
     STARTUPINFO  si={0};
     PROCESS_INFORMATION pi={0};
     si.cb = sizeof(si);
     si.dwFlags = STARTF_USESHOWWINDOW;
     si.wShowWindow = SW_HIDE;

     // increase resource allocation to program
     SetPriorityClass(GetCurrentProcess(),
      REALTIME_PRIORITY_CLASS);
     SetThreadPriority(GetCurrentThread(),
      THREAD_PRIORITY_TIME_CRITICAL);

     // invoke command shell
     if(CreateProcess(0, szComspec, 0, 0, 0,CREATE_SUSPENDED|
      DETACHED_PROCESS, 0, 0, &si, &pi))
     {
      // suppress command shell process until program exits
      SetPriorityClass(pi.hProcess,IDLE_PRIORITY_CLASS);
      SetThreadPriority(pi.hThread,THREAD_PRIORITY_IDLE);

      // resume shell process with new low priority
      ResumeThread(pi.hThread);

      // everything seemed to work
      return TRUE;
     }
     else // if error, normalize allocation
     {
      SetPriorityClass(GetCurrentProcess(),
       NORMAL_PRIORITY_CLASS);
      SetThreadPriority(GetCurrentThread(),
       THREAD_PRIORITY_NORMAL);
     }
    }
    return FALSE;
   }
*/
import "C"

import (
	"bytes"
	"fmt"
	"github.com/flopp/go-findfont"
	"github.com/kardianos/service"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

type program struct{}

func (p program) Start(s service.Service) error {
	//TODO implement me
	panic("implement me")
}

func (p program) Stop(s service.Service) error {
	//TODO implement me
	s.Stop()
	panic("implement me")
}

func init() {
	if !IsAdmin() {
		err := RunMeElevated()
		if err != nil {
			os.Exit(0)
		}
		os.Exit(0)
	}
	// 设置中文环境变量
	paths := findfont.List()
	for _, path := range paths {
		// msyh.ttc
		if strings.Contains(path, "msyh.ttc") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

func main() {
	//a := app.New()
	//a.Settings().SetTheme(theme.DarkTheme())
	//w := a.NewWindow("天府安卸载程序")
	//abs, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//form := widget.NewForm(
	//	widget.NewFormItem("名称", widget.NewLabel("天府安")),
	//	widget.NewFormItem("版本", widget.NewLabel("0.0.1")),
	//	widget.NewFormItem("路径", widget.NewLabel(abs)),
	//)
	//
	//form.SubmitText = "卸载"
	//form.CancelText = "取消"
	//
	//form.OnCancel = func() {
	//	a.Quit()
	//}
	//
	//form.OnSubmit = func() {
	//	AppIsRunning()
	//	DeleteDesktopLnk()
	//	DeleteRegedit()
	//	DeleteServer()
	//	DeleteDir()
	//	DeleteSelf()
	//}
	//
	//w.SetContent(container.NewCenter(form))
	//w.CenterOnScreen()
	//w.ShowAndRun()

	AppIsRunning()
	DeleteDesktopLnk()
	DeleteRegedit()
	DeleteServer()
	DeleteDir()
	DeleteSelf()
}

// AppIsRunning 检查程序是否运行
func AppIsRunning() {
	// 检查某个程序是否在运行
	out, err := exec.Command("tasklist").Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	if bytes.Contains(out, []byte("tianfuan.exe")) {
		// 结束程序运行
		cmd := exec.Command("taskkill", "/IM", "tianfuan.exe", "/F")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}
	}
}

// DeleteDesktopLnk 删除快捷方式
func DeleteDesktopLnk() {
	// 获取当前用户的home目录路径
	HomeDir, _ := os.UserHomeDir()
	// 获取桌面路径
	DesktopPath := filepath.Join(HomeDir, "Desktop")

	DesktopPath += "\\天府安.lnk"
	DesktopPath = strings.ReplaceAll(DesktopPath, "\\", "\\\\")
	os.Remove(DesktopPath)
}

// DeleteRegedit 删除注册表
func DeleteRegedit() {
	err := registry.DeleteKey(registry.LOCAL_MACHINE, "SOFTWARE\\\\Microsoft\\\\Windows\\\\CurrentVersion\\\\Uninstall\\tianfuan")
	if err != nil {
		return
	}
}

// DeleteServer 删除服务
func DeleteServer() {
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

	s.Stop()
	s.Uninstall()
}

// DeleteDir 删除目录
func DeleteDir() {
	abs, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.RemoveAll(abs)
}

// DeleteSelf 删除自身
func DeleteSelf() {
	abs, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Chdir(abs)
	C.SelfDelete()
}

// IsAdmin 判断当前权限是否是管理员
func IsAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	return true
}

// RunMeElevated 以管理员身份运行
func RunMeElevated() error {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 0 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		return err
	}
	return nil
}
