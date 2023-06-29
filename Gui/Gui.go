package Gui

import (
	"embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fyne.io/systray"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/fyne-io/terminal"
	_ "image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net"
	"os"
	"os/exec"
	"tianfuan/Antivirus"
	"tianfuan/CloudInspection"
	"tianfuan/DirectoryMonitoring"
	"tianfuan/HFish/utils/setting"
	"tianfuan/Regedit"
	"tianfuan/Server"
	"tianfuan/System"
	"time"
)

// 静态资源打包 区域
//
//go:embed static
var static embed.FS

type App struct {
	Regedit.Regedit
	Server.Server
	System.System
	DirectoryMonitoring.Watch
	CloudInspection.CloudInspection
	Antivirus.Rule
	a                         fyne.App
	w                         fyne.Window
	Gif                       *AnimatedGif
	FileDetectionButton       *widget.Button
	HorseHuntingButton        *widget.Button
	CloudInspectionButton     *widget.Button
	ComputerCleaningButton    *widget.Button
	HoneypotSystemButton      *widget.Button
	DirectoryMonitoringButton *widget.Button
	CleanButton               *widget.Button
}

// ReturnStartAnimationContainer 返回启动动画容器 开场白
func (MyApp *App) ReturnStartAnimationContainer() *fyne.Container {
	// 开场白
	file, _ := static.ReadFile("static/draw.gif")
	bgm, _ := static.Open("static/bgm.mp3")

	go func() {
		streamer, format, err := mp3.Decode(bgm)
		if err != nil {
			return
		}

		defer streamer.Close()
		done := make(chan bool)
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- true
		})))
		<-done
	}()

	MyApp.Gif, _ = NewAnimatedGifFromResource(fyne.NewStaticResource("draw", file))
	MyApp.Gif.Start()
	return container.NewMax(MyApp.Gif)
}

// ReturnWidgetContainer 返回控件容器
func (MyApp *App) ReturnWidgetContainer() *fyne.Container {
	MyApp.FileDetectionButton = widget.NewButton("文件检测", func() {
		go func() {
			w := MyApp.a.NewWindow("文件检测")
			text := canvas.NewText("检测目标 :", color.RGBA{255, 0, 255, 0})
			entry := widget.NewEntry()
			entry.TextStyle = fyne.TextStyle{Bold: true, Symbol: true}
			data := [][]string{
				{"文件路径", "病毒类型", "病毒家族"},
			}

			table := widget.NewTable(
				func() (int, int) {
					return len(data), len(data[0])
				},
				func() fyne.CanvasObject {
					label := widget.NewLabel("")
					label.Alignment = 1
					return label
				},
				func(i widget.TableCellID, o fyne.CanvasObject) {
					o.(*widget.Label).SetText(data[i.Row][i.Col])
				})

			table.SetColumnWidth(0, 450)
			table.SetColumnWidth(1, 300)
			table.SetColumnWidth(2, 300)

			entry.SetPlaceHolder("c:\\windows\\win.ini")
			entry.ActionItem = widget.NewButtonWithIcon("", theme.FileIcon(), func() {
				OpenFile := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
					if err != nil {
						return
					}

					if closer == nil {
						return
					}

					entry.SetText(closer.URI().Path())
				}, w)

				OpenFile.Resize(fyne.NewSize(1080, 720))
				OpenFile.Show()
			})

			StartButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
				if entry.Text == "" {
					return
				}

				MyApp.Rule.Result = make(chan [][]string, 10)
				MyApp.Rule.FilePath = make(chan string, 10)

				go func() {
					for r := range MyApp.Rule.Result {
						data = append(data, r...)
						table.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {
							template.(*widget.Label).SetText(data[id.Row][id.Col])
						}
					}
				}()

				MyApp.Rule.FilePath <- entry.Text
				go func() {
					MyApp.Rule.LoadRules()
					MyApp.Rule.Flag = true
					MyApp.Rule.SingleFileAntivirusScan()
					time.Sleep(time.Second * 5)
					close(MyApp.Rule.FilePath)
					close(MyApp.Rule.Result)
				}()
			})

			// 目标检测左边盒子内容
			ContainerDetectionLeftBox := container.New(layout.NewBorderLayout(nil, nil, nil, StartButton), StartButton, entry)
			// 检测目标最终容器
			ContainerDetectionTarget := container.New(layout.NewFormLayout(), text, ContainerDetectionLeftBox)

			lineEntry := widget.NewMultiLineEntry()
			lineEntry.Disable()
			lineEntry.SetMinRowsVisible(30)

			file, _ := static.ReadFile("static/bg.png")

			w.SetContent(container.NewMax(canvas.NewImageFromResource(fyne.NewStaticResource("", file)), container.NewVBox(ContainerDetectionTarget, container.NewAppTabs(container.NewTabItem("检测结果", table), container.NewTabItem("状态栏", lineEntry)))))

			w.Resize(fyne.NewSize(1080, 720))
			w.Show()
		}()
	})

	MyApp.HorseHuntingButton = widget.NewButton("全盘查杀", func() {
		go func() {
			w := MyApp.a.NewWindow("全盘查杀")
			text := canvas.NewText("检测目标 :", color.RGBA{255, 0, 255, 0})
			barInfinite := widget.NewProgressBarInfinite()
			barInfinite.Stop()

			entry := widget.NewEntry()
			entry.Disable()
			entry.TextStyle = fyne.TextStyle{Bold: true, Symbol: true}
			data := [][]string{
				{"文件路径", "病毒类型", "病毒家族"},
			}

			table := widget.NewTable(
				func() (int, int) {
					return len(data), len(data[0])
				},
				func() fyne.CanvasObject {
					label := widget.NewLabel("")
					label.Alignment = 1
					return label
				},
				func(i widget.TableCellID, o fyne.CanvasObject) {
					o.(*widget.Label).SetText(data[i.Row][i.Col])
				})

			table.SetColumnWidth(0, 450)
			table.SetColumnWidth(1, 300)
			table.SetColumnWidth(2, 300)

			entry.SetPlaceHolder("c:\\windows\\win.ini")
			entry.ActionItem = widget.NewButtonWithIcon("", theme.FileIcon(), func() {
				OpenFile := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
					if err != nil {
						return
					}

					if closer == nil {
						return
					}

					entry.SetText(closer.URI().Path())
				}, w)

				OpenFile.Resize(fyne.NewSize(1080, 720))
				OpenFile.Show()
			})

			StartButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
				entry.SetText("正在查杀中...")
				barInfinite.Start()

				MyApp.Rule.Thread = 50
				MyApp.Rule.FilePath = make(chan string, 10)
				MyApp.Rule.Result = make(chan [][]string, 10)
				MyApp.Rule.Process = make(chan string, 10)

				go func() {
					for r := range MyApp.Rule.Process {
						entry.SetText(r)
					}
				}()

				go func() {
					for r := range MyApp.Rule.Result {
						if r[0][1] == "无检出" || r[0][1] == "检测错误" {
							continue
						}

						data = append(data, r...)
						table.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {
							template.(*widget.Label).SetText(data[id.Row][id.Col])
						}
					}
				}()

				go func() {
					MyApp.Rule.LoadRules()
					MyApp.Rule.Flag = false
					MyApp.Rule.FullAntivirusScan()
					time.Sleep(time.Second * 30)
					entry.SetText("任务结束")
					barInfinite.Stop()
					close(MyApp.Rule.FilePath)
					close(MyApp.Rule.Process)
					close(MyApp.Rule.Result)
				}()
			})

			// 目标检测左边盒子内容
			ContainerDetectionLeftBox := container.New(layout.NewBorderLayout(nil, nil, nil, StartButton), StartButton, entry)
			// 检测目标最终容器
			ContainerDetectionTarget := container.New(layout.NewFormLayout(), text, ContainerDetectionLeftBox)

			lineEntry := widget.NewMultiLineEntry()
			lineEntry.Disable()
			lineEntry.SetMinRowsVisible(30)

			file, _ := static.ReadFile("static/bg.png")

			w.SetContent(container.NewMax(canvas.NewImageFromResource(fyne.NewStaticResource("", file)), container.NewVBox(ContainerDetectionTarget, barInfinite, container.NewAppTabs(container.NewTabItem("检测结果", table), container.NewTabItem("状态栏", lineEntry)))))

			w.Resize(fyne.NewSize(1080, 720))
			w.Show()
		}()
	})

	MyApp.CloudInspectionButton = widget.NewButton("云查杀", func() {
		go func() {
			w := MyApp.a.NewWindow("云查杀")
			text := canvas.NewText("检测目标 :", color.RGBA{255, 0, 255, 0})
			entry := widget.NewEntry()
			entry.TextStyle = fyne.TextStyle{Bold: true, Symbol: true}
			infinite := widget.NewProgressBarInfinite()
			infinite.Stop()
			data := [][]string{
				{"平台", "威胁等级", "引擎数量", "实际数量", "检出个数", "病毒类型", "病毒家族"},
			}

			table := widget.NewTable(
				func() (int, int) {
					return len(data), len(data[0])
				},
				func() fyne.CanvasObject {
					label := widget.NewLabel("")
					label.Alignment = 1
					return label
				},
				func(i widget.TableCellID, o fyne.CanvasObject) {
					o.(*widget.Label).SetText(data[i.Row][i.Col])
				})

			table.SetColumnWidth(0, 80)
			table.SetColumnWidth(1, 95)
			table.SetColumnWidth(2, 80)
			table.SetColumnWidth(3, 80)
			table.SetColumnWidth(4, 80)
			table.SetColumnWidth(5, 300)
			table.SetColumnWidth(6, 300)

			entry.SetPlaceHolder("c:\\windows\\win.ini")
			entry.ActionItem = widget.NewButtonWithIcon("", theme.FileIcon(), func() {
				OpenFile := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
					if err != nil {
						return
					}

					if closer == nil {
						return
					}

					entry.SetText(closer.URI().Path())
				}, w)

				OpenFile.Resize(fyne.NewSize(1080, 720))
				OpenFile.Show()
			})

			StartButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
				if entry.Text == "" {
					return
				}

				MyApp.CloudInspection.FilePath = entry.Text

				go func() {
					for r := range MyApp.CloudInspection.Result {
						data = append(data, r...)
						table.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {
							template.(*widget.Label).SetText(data[id.Row][id.Col])
						}
					}
				}()

				go func() {
					infinite.Start()
					MyApp.CloudInspection.Run()
					infinite.Stop()
				}()
			})

			// 目标检测左边盒子内容
			ContainerDetectionLeftBox := container.New(layout.NewBorderLayout(nil, nil, nil, StartButton), StartButton, entry)
			// 检测目标最终容器
			ContainerDetectionTarget := container.New(layout.NewFormLayout(), text, ContainerDetectionLeftBox)

			lineEntry := widget.NewMultiLineEntry()
			lineEntry.Disable()
			lineEntry.SetMinRowsVisible(30)

			file, _ := static.ReadFile("static/bg.png")

			w.SetContent(container.NewMax(canvas.NewImageFromResource(fyne.NewStaticResource("", file)), container.NewVBox(ContainerDetectionTarget, infinite, container.NewAppTabs(container.NewTabItem("检测结果", table), container.NewTabItem("状态栏", lineEntry)))))

			w.Resize(fyne.NewSize(1080, 720))
			w.Show()
		}()
	})

	MyApp.ComputerCleaningButton = widget.NewButton("电脑清理", func() {
		go func() {
			w := MyApp.a.NewWindow("电脑清理")
			t := terminal.New()

			go func() {
				_ = t.RunLocalShell()
				MyApp.a.Quit()
			}()

			CleanButton := widget.NewButton("开始清理", func() {
				go func() {
					_, err := os.Stat("clean.bat")
					if err != nil {
						bat, err := static.ReadFile("static/clean.bat")
						if err != nil {
							fmt.Println(err)
							return
						}
						file, err := os.OpenFile("clean.bat", os.O_WRONLY|os.O_CREATE, 0777)
						if err != nil {
							fmt.Println(err)
							return
						}
						_, err = file.Write(bat)
						if err != nil {
							fmt.Println(err)
							return
						}
						file.Close()
					}

					event := fyne.KeyEvent{
						Name: fyne.KeyEnter,
					}

					t.Write([]byte("cmd /c \"color a && clean.bat\"\n"))
					t.TypedKey(&event)
				}()
			})

			file, _ := static.ReadFile("static/bg.png")
			w.SetContent(container.NewMax(canvas.NewImageFromResource(fyne.NewStaticResource("", file)), container.NewVSplit(container.NewVBox(CleanButton), container.NewMax(t))))
			w.Resize(fyne.NewSize(1080, 720))
			w.Show()
		}()
	})

	MyApp.HoneypotSystemButton = widget.NewButton("蜜罐系统", func() {
		go func() {
			w := MyApp.a.NewWindow("蜜罐系统")
			label := widget.NewLabel("关闭")

			f := func() {
				_, err := net.Dial("tcp", "127.0.0.1:9001")
				if err != nil {
					label.SetText("关闭")
				} else {
					label.SetText("开启")
				}
			}

			f()

			StartButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
				if label.Text == "开启" {
					cmd := exec.Command("cmd", "/c", "start http://127.0.0.1:9001/login")
					err := cmd.Run()
					if err != nil {
						return
					}
					return
				}

				go func() {
					go setting.Run()
					cmd := exec.Command("cmd", "/c", "start http://127.0.0.1:9001/login")
					err := cmd.Run()
					if err != nil {
						return
					}
					f()
				}()
			})

			form := widget.NewForm(
				widget.NewFormItem("状态", label),
				widget.NewFormItem("开启", StartButton),
			)

			file, _ := static.ReadFile("static/bg.png")

			w.SetContent(container.NewMax(canvas.NewImageFromResource(fyne.NewStaticResource("", file)), container.NewCenter(form)))
			w.Resize(fyne.NewSize(1080, 720))
			w.Show()
		}()
	})

	MyApp.DirectoryMonitoringButton = widget.NewButton("目录监控", func() {
		go func() {
			w := MyApp.a.NewWindow("目录监控")
			data := [][]string{
				{"记录时间", "文件路径", "触发行为"},
			}

			entry := widget.NewEntry()
			entry.TextStyle = fyne.TextStyle{Bold: true, Symbol: true}
			entry.Disable()

			table := widget.NewTable(
				func() (int, int) {
					return len(data), len(data[0])
				},
				func() fyne.CanvasObject {
					label := widget.NewLabel("")
					label.Alignment = 1
					return label
				},
				func(i widget.TableCellID, o fyne.CanvasObject) {
					o.(*widget.Label).SetText(data[i.Row][i.Col])
				})

			table.SetColumnWidth(0, 200)
			table.SetColumnWidth(1, 700)
			table.SetColumnWidth(2, 150)

			StartButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
				MyApp.Watch.Result = make(chan [][]string, 10)
				MyApp.Watch.Process = make(chan string, 10)

				go func() {
					for p := range MyApp.Watch.Process {
						entry.SetText(p)
					}
				}()

				go func() {
					for r := range MyApp.Watch.Result {
						data = append(data, r...)
						table.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {
							template.(*widget.Label).SetText(data[id.Row][id.Col])
						}
					}
				}()

				go func() {
					// MyApp.Watch.WatchAllDir() // 全盘监控
					MyApp.Watch.WatchDir()
					MyApp.Watch.WatchDangerDir() // 敏感目录监控
					entry.SetText("正在监控中...")
				}()
			})

			// 目标检测左边盒子内容

			ContainerDetectionLeftBox := container.New(layout.NewBorderLayout(nil, nil, nil, container.NewMax(entry, StartButton)), container.NewMax(entry, StartButton), widget.NewLabel(""))
			// 检测目标最终容器
			ContainerDetectionTarget := container.New(layout.NewFormLayout(), widget.NewLabel(""), ContainerDetectionLeftBox)

			lineEntry := widget.NewMultiLineEntry()
			lineEntry.Disable()
			lineEntry.SetMinRowsVisible(30)

			file, _ := static.ReadFile("static/bg.png")

			w.SetContent(container.NewMax(canvas.NewImageFromResource(fyne.NewStaticResource("", file)), container.NewVBox(ContainerDetectionTarget, container.NewAppTabs(container.NewTabItem("检测结果", table), container.NewTabItem("状态栏", lineEntry)))))
			w.Resize(fyne.NewSize(1080, 720))
			w.Show()
		}()
	})

	return container.NewMax(container.NewGridWithColumns(3, MyApp.FileDetectionButton, MyApp.HorseHuntingButton, MyApp.CloudInspectionButton, MyApp.ComputerCleaningButton, MyApp.HoneypotSystemButton, MyApp.DirectoryMonitoringButton))
}

// ReturnMyAppContainer 返回最终容器
func (MyApp *App) ReturnMyAppContainer() *fyne.Container {
	file, _ := static.ReadFile("static/bg.png")
	bg := canvas.NewImageFromResource(fyne.NewStaticResource("bg", file))
	return container.NewMax(bg, MyApp.ReturnWidgetContainer())
}

func (MyApp *App) ShowGui() {
	// 鉴权 是否是管理员权限
	MyApp.System.CheckAdministrator()
	// 鉴权 是否是超级管理员权限
	MyApp.System.CheckSystem()
	// 检查注册表
	MyApp.Regedit.CheckRegedit()
	// 检查服务
	MyApp.Server.CheckServer()

	// 环境正常 进入程序
	MyApp.a = app.New()
	MyApp.a.Settings().SetTheme(CustomTheme{"Dark"})
	ico, _ := static.ReadFile("static/ico.png")
	MyApp.a.SetIcon(fyne.NewStaticResource("ico", ico))
	MyApp.w = MyApp.a.NewWindow("天府安")

	// 启动动画
	MyApp.w.SetContent(MyApp.ReturnStartAnimationContainer())
	go createTrayIcon()

	go func() {
		select {
		case <-time.After(time.Second * 3):
			MyApp.w.SetContent(MyApp.ReturnMyAppContainer())
			MyApp.Gif.Stop()
			MyApp.Gif = nil
			return
		}
	}()

	MyApp.w.Resize(fyne.NewSize(1080, 720))
	MyApp.w.CenterOnScreen()
	MyApp.w.ShowAndRun()
}

// Define the function to create the system tray icon
func createTrayIcon() {
	// Create a new systray
	systray.Run(onReady, onExit)
}

// Define the onReady function to be called when the systray is ready
func onReady() {
	// Add an icon to the systray
	file, _ := static.ReadFile("static/ico.ico")
	systray.SetIcon(file)
	// Add a tooltip to the systray icon
	systray.SetTooltip("tianfuan")
}

// Define the onExit function to be called when the systray is exited
func onExit() {
	// Do any cleanup here
	systray.Quit()
}
