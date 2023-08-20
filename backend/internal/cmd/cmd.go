package cmd

import (
	"flag"
	"media-down/backend/pkg/logs"
	"strings"
	"time"
)

func main() {

	url := flag.String("url", "", "mdeia page url")
	mediaType := flag.String("type", "mdeia", "媒体类型,如: video, music")
	waitTime := flag.Float64("waitTime", 0.5, "加载等待时间, 单位: 分钟")
	headless := flag.Bool("headless", true, "是否隐藏浏览器")
	flag.Parse()
	//url 低于12不合法,最短正确示例: http://a.cn/
	if len(*url) < 12 || !strings.HasPrefix(*url, "http") {
		logs.Error("url不合法!")
		flag.PrintDefaults()
		return
	}

	timeout := time.Second * time.Duration(*waitTime*60)
	// cmd.Chromedp(*url, *mediaType, timeout, *headless)
	Rod(*url, *mediaType, timeout, *headless)
}

type Cmd struct {
}

func NewCmd() *Cmd {
	return &Cmd{}
}

func (a *Cmd) MediaDown(url string) {
	mediaType := "media"
	timeout := time.Second * time.Duration(0.5*60)
	Rod(url, mediaType, timeout, true)
}
