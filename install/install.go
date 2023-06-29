package main

/*
#include "stdio.h"
#include "windows.h"

int CreateSubKey(){
HKEY hKey;
LPCTSTR subKey = TEXT("SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Uninstall\\tianfuan");
LONG createStatus = RegCreateKeyEx(HKEY_LOCAL_MACHINE, subKey, 0, NULL, REG_OPTION_NON_VOLATILE, KEY_ALL_ACCESS, NULL, &hKey, NULL);
if (createStatus == ERROR_SUCCESS) {
    // Key created successfully
    RegCloseKey(hKey);
} else {
    // Failed to create key
}
}
*/
import "C"
import (
	"archive/zip"
	"bytes"
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

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

//go:embed Icon.png
var icon []byte

//go:embed app.zip
var AppZip []byte

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	a.SetIcon(fyne.NewStaticResource("icon", icon))
	w := a.NewWindow("天府安安装程序")

	FilePath := widget.NewEntry()
	bar := widget.NewProgressBar()
	infinite := widget.NewProgressBarInfinite()
	bar.Hide()
	infinite.Hide()

	FilePath.ActionItem = widget.NewButtonWithIcon("", theme.FileIcon(), func() {
		OpenDir := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				return
			}

			FilePath.SetText(uri.Path())
		}, w)
		OpenDir.Show()
	})

	FilePath.SetText("C:\\Program Files\\tianfuan")
	FilePath.Disable()

	InstallButton := widget.NewButton("Install", func() {
		bar.Show()
		infinite.Show()
		// Create a temporary file to write the embedded zip archive to
		tmpfile, err := os.CreateTemp("", "example")
		if err != nil {
			log.Fatal(err)
		}

		defer os.Remove(tmpfile.Name())

		// Write the embedded zip archive to the temporary file
		if _, err := tmpfile.Write(AppZip); err != nil {
			log.Fatal(err)
		}

		// Open the temporary file as a zip archive
		r, err := zip.OpenReader(tmpfile.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()

		// Iterate through each file in the zip file
		var totalSize int64
		for _, f := range r.File {
			totalSize += int64(f.UncompressedSize64)
		}
		// Open each file and write its contents to a new file
		var writtenSize int64
		var decodeName string
		for _, f := range r.File {
			if f.Flags == 0 {
				//如果标致位是0  则是默认的本地编码   默认为gbk
				i := bytes.NewReader([]byte(f.Name))
				decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
				content, _ := ioutil.ReadAll(decoder)
				decodeName = string(content)
			} else {
				//如果标志为是 1 << 11也就是 2048  则是utf-8编码
				decodeName = f.Name
			}

			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}

			defer rc.Close()

			// Create the new file
			path := filepath.Join(FilePath.Text, decodeName)
			if f.FileInfo().IsDir() {
				os.MkdirAll(path, f.Mode())
			} else {
				os.MkdirAll(filepath.Dir(path), f.Mode())
				outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
				if err != nil {
					log.Fatal(err)
				}
				defer outFile.Close()

				// Write the file contents to the new file
				written, err := io.Copy(outFile, rc)
				if err != nil {
					log.Fatal(err)
				}
				writtenSize += written

				// Calculate and print the percentage of the zip file that has been extracted
				percentage := float64(writtenSize) / float64(totalSize) * 100
				bar.SetValue(percentage)
			}
		}
		infinite.Hide()
		CreateRegedit(FilePath.Text)
		CreateDesktopInk(fmt.Sprintf("%s\\tianfuan.exe", FilePath.Text))
		os.Exit(0)
	})

	form := widget.NewForm(
		widget.NewFormItem("安装路径", FilePath),
		widget.NewFormItem("", bar),
		widget.NewFormItem("", infinite),
	)

	w.SetContent(container.NewVBox(form, InstallButton))
	w.Resize(fyne.NewSize(1080, 720))
	w.CenterOnScreen()
	w.ShowAndRun()
}

func CreateRegedit(src string) {
	C.CreateSubKey() // 创建项
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\\\Microsoft\\\\Windows\\\\CurrentVersion\\\\Uninstall\\tianfuan", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	key.SetStringValue("DisplayIcon", fmt.Sprintf("%s\\tianfuan.exe", src))
	key.SetStringValue("DisplayName", "天府安")
	key.SetStringValue("DisplayVersion", "0.0.1")
	key.SetStringValue("Publisher", "天府安")
	key.SetStringValue("QuietUninstallString", fmt.Sprintf("%s\\uninstall.exe", src))
	key.SetStringValue("UninstallString", fmt.Sprintf("%s\\uninstall", src))
}

func CreateDesktopInk(src string) {
	runtime.LockOSThread()
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		log.Fatal(err)
	}
	defer oleShellObject.Release()
	wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		log.Fatal(err)
	}
	defer wshell.Release()

	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", fmt.Sprintf("%s\\天府安.lnk", filepath.Join(os.Getenv("USERPROFILE"), "Desktop")))
	if err != nil {
		log.Fatal(err)
	}
	idispatch := cs.ToIDispatch()
	oleutil.PutProperty(idispatch, "TargetPath", src)
	oleutil.CallMethod(idispatch, "Save")
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
