package DirectoryMonitoring

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
	. "strconv"
	"syscall"
	"time"
)

type Watch struct {
	Process chan string
	Result  chan [][]string
	watch   *fsnotify.Watcher
}

// GetLogicalDrives  获取系统盘符
func (w *Watch) GetLogicalDrives() []string {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
	n, _, _ := GetLogicalDrives.Call()
	s := FormatInt(int64(n), 2)
	var DrivesAll = []string{"A:\\\\", "B:\\\\", "C:\\\\", "D:\\\\", "E:\\\\", "F:\\\\", "G:\\\\", "H:\\\\", "I:\\\\", "J:\\\\", "K:\\\\", "L:\\\\", "M:\\\\", "N:\\\\", "O:\\\\", "P:\\\\", "Q:\\\\", "R:\\\\", "S:\\\\", "T:\\\\", "U:\\\\", "V:\\\\", "W:\\\\", "X:\\\\", "Y:\\\\", "Z:\\\\"}
	temp := DrivesAll[0:len(s)]
	var d []string
	for i, v := range s {
		if v == 49 {
			l := len(s) - i - 1
			d = append(d, temp[l])
		}
	}
	var drives []string
	for i, v := range d {
		drives = append(drives[i:], append([]string{v}, drives[:i]...)...)
	}
	return drives
}

// WatchAllDir 全盘监控
func (w *Watch) WatchAllDir() {
	for _, driver := range w.GetLogicalDrives() {
		err := filepath.Walk(driver, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
			} else {
				if info.IsDir() {
					path, err := filepath.Abs(path)
					if err != nil {
						fmt.Println(err)
					} else {
						err = w.watch.Add(path)
						if err != nil {
							fmt.Println(err)
						}
						w.Process <- path
					}
				}
			}
			return nil
		})

		if err != nil {
			fmt.Println(err)
		}
	}
}

// WatchDangerDir 敏感目录监控
func (w *Watch) WatchDangerDir() {
	for _, driver := range w.GetLogicalDrives() {
		path := driver
		err := w.watch.Add(path)
		if err != nil {
			fmt.Println(err)
		}
		w.Process <- path
	}

	DangerPath := []string{
		"C:\\Windows",
		"C:\\Windows\\System32",
		"C:\\Windows\\SysWOW64",
		"C:\\Program Files",
		"C:\\Program Files (x86)",
	}

	for _, path := range DangerPath {
		err := w.watch.Add(path)
		if err != nil {
			fmt.Println(err)
		}
		w.Process <- path
	}
}

// WatchDir 监控目录
func (w *Watch) WatchDir() {
	w.watch, _ = fsnotify.NewWatcher()

	go func() {
		for {
			select {
			case ev := <-w.watch.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						w.Result <- [][]string{{time.Now().Format(time.DateTime), ev.Name, "创建文件"}}
						//这里获取新创建文件的信息，如果是目录，则加入监控中
						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							w.watch.Add(ev.Name)
							w.Result <- [][]string{{time.Now().Format(time.DateTime), ev.Name, "添加监控"}}
						}
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						w.Result <- [][]string{{time.Now().Format(time.DateTime), ev.Name, "写入文件"}}
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						w.Result <- [][]string{{time.Now().Format(time.DateTime), ev.Name, "删除文件"}}
						//如果删除文件是目录，则移除监控
						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							w.watch.Remove(ev.Name)
							w.Result <- [][]string{{time.Now().Format(time.DateTime), ev.Name, "删除监控"}}
						}
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						w.Result <- [][]string{{time.Now().Format(time.DateTime), ev.Name, "重命名文件"}}
						//如果重命名文件是目录，则移除监控
						//注意这里无法使用os.Stat来判断是否是目录了
						//因为重命名后，go已经无法找到原文件来获取信息了
						//所以这里就简单粗爆的直接remove好了
						w.watch.Remove(ev.Name)
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						w.Result <- [][]string{{time.Now().Format(time.DateTime), ev.Name, "修改权限"}}
					}
				}

			case err := <-w.watch.Errors:
				{
					fmt.Println("error : ", err)
					return
				}
			}
		}
	}()
}
