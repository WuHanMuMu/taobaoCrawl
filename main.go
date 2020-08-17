package main

import (
	"context"
	"crawler/taobao"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/chromedp/chromedp"
	"github.com/flopp/go-findfont"
	"os"
	"strconv"
	"strings"
)


func startChrome(){
	// 启动浏览器配置
	opt := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opt...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	res := "abc"

	fmt.Print("===>", res)
	chromedp.Run(ctx, taobao.MyHeaders(&res))
	fmt.Print("===>",res, "\r\n")
}

func init(){
	// 获取中文字体路径
	fontsPath := findfont.List()
	for _,path:= range fontsPath{
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		fmt.Println(path)
		if strings.Contains(path, "华文新魏.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
		if strings.Contains(path, "simhei.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

func info() fyne.CanvasObject{
	searchInput:= widget.NewEntry()
	searchInput.SetPlaceHolder("请输入搜索关键词")
	form:= widget.NewForm()
	form.Append("搜索词",searchInput)
	//form.SubmitText = "提交"
	form.Append("",widget.NewButton("提交", func() {
		fmt.Println("提交时间" + searchInput.Text)
		searchInput.SetText("")
	}))
	return widget.NewScrollContainer(widget.NewGroup("查询",form))
}

func main() {
	defer os.Unsetenv("FYNE_FONT")
	fmt.Println("hello,world")
	count:=0
	a:= app.New()
	window:= a.NewWindow("测试一波gui")
	window.Resize(fyne.NewSize(800,400))


	window.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayoutWithColumns(2),
		info(),
		widget.NewButton("heloo", func() {
			count ++
			window.SetTitle(strconv.Itoa(count))
			//startChrome()
		}),
	))
	window.ShowAndRun()
}
