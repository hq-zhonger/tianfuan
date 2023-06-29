package Antivirus

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	. "strconv"
	"strings"
	"syscall"
)

type Rule struct {
	Flag     bool
	FilePath chan string     // 文件路径
	Result   chan [][]string // 结果
	Process  chan string     // 当前进程
	Thread   int             // 线程
	Suffix   []string        // 后缀名
	Rule     []Antivirus     `json:"finger"`
}

type Antivirus struct {
	MD5    string `json:"md5"`
	SHA256 string `json:"sha256"`
	Hex    string `json:"hex"`
	Type   string `json:"type"`
	Family string `json:"family"`
}

// LoadRules 加载规则
func (r *Rule) LoadRules() {
	file, err := os.Open("Antivirus\\rules.json")
	defer file.Close()

	if err != nil {
		fmt.Println("规则加载错误")
		log.Fatal(err)
	}

	// 加载后缀 默认检查选项
	r.Suffix = []string{
		// Shell
		".exe",
		//".dll",
		//".com",
		//".ocx",
		//".vxd",
		//".sys",
		// WebShell
		//".vbs",
		//"js",
		//".php",
		//"asp",
		//"jsp",
		//"html",
		//"hta",
		//"htm",
		//// 移动
		//"apk",
		//"jar",
		//"java",
		//// 宏
		//"doc",
		//"docx",
		//"ppt",
		//"xls",
		//"xlsx",
		//"pdf",
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&r)

	if err != nil {
		fmt.Println("规则解析错误")
		log.Fatal(err)
	}
}

// SingleFileAntivirusScan 单文件反病毒扫描
func (r *Rule) SingleFileAntivirusScan() {
	for path := range r.FilePath {
		// fmt.Println("执行 -》 1")
		data, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			r.Result <- [][]string{{path, "检测错误", "检测错误"}}
			if r.Flag {
				return
			} else {
				continue
			}
		}

		MD5 := fmt.Sprintf("%x", md5.Sum(data))
		SHA256 := fmt.Sprintf("%x", sha256.Sum256(data))
		HEX := hex.EncodeToString(data)

		// fmt.Println("执行 -> 3")
		for _, rule := range r.Rule {
			strings.ToLower(strings.TrimSpace(rule.Hex))
			if MD5 == rule.MD5 || SHA256 == rule.SHA256 || strings.Contains(HEX, strings.ToLower(strings.TrimSpace(rule.Hex))) == true {
				// 文件路径 病毒类型 病毒家族
				r.Result <- [][]string{{path, rule.Type, rule.Family}}
				if r.Flag {
					return
				} else {
					continue
				}
			} else {
				r.Result <- [][]string{{path, "无检出", "无检出"}}
				if r.Flag {
					return
				} else {
					continue
				}
			}
		}
	}
}

// GetLogicalDrives  获取系统盘符
func (r *Rule) GetLogicalDrives() []string {
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

// FullAntivirusScan 全盘扫描
func (r *Rule) FullAntivirusScan() {
	for _, v := range r.GetLogicalDrives() {
		fmt.Println(v)
	}

	for i := 1; i <= r.Thread; i++ {
		go r.SingleFileAntivirusScan()
	}

	for _, v := range r.GetLogicalDrives() {
		err := filepath.Walk(v, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
			} else {
				if !info.IsDir() {
					if info.Size() <= 104857600 && info.Size() != 0 {
						for _, ext := range r.Suffix {
							if filepath.Ext(path) == ext {
								r.Process <- path
								r.FilePath <- path
							}
						}
					}
				}
			}
			return nil
		})

		if err != nil {
			fmt.Println(err)
		}
	}
	r.Flag = true
}
