package cmd

import (
	"context"
	"media-down/backend/pkg/common"
	"media-down/backend/pkg/logs"
	"media-down/backend/pkg/media"
	"os"
	"strings"
	"time"

	ccdp "media-down/backend/pkg/chromedp"

	"github.com/chromedp/chromedp"
)

func Chromedp(url, mediaType string, waitTime time.Duration, headless bool) {
	// var (
	// 	wg sync.WaitGroup
	// )
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", headless),
		// chromedp.DisableGPU,
		chromedp.Flag("ignore-certificate-errors", true), //忽略错误
		chromedp.Flag("disable-web-security", true),      //禁用网络安全标志
		chromedp.NoFirstRun,                              //设置网站不是首次运行
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"), //设置UserAgent
	)
	if common.IsWindows() {
		edgePath := "C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe"
		if common.PathIsExis(edgePath) {
			opts = append(opts,
				chromedp.NoDefaultBrowserCheck, //不检查默认浏览器
				chromedp.ExecPath(edgePath))
		}
	}
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	// create context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	downMedia(url, mediaType, waitTime, ctx)
	logs.Info("download finish!")

	// 这将阻塞，直到 chromedp 侦听器关闭通道
	ctx.Deadline()
}

func downMedia(url, mediaType string, timeout time.Duration, ctx context.Context) {
	//设置一个通道，以便我们稍后在监控下载时阻止进程
	// fired := make(chan struct{})
	logs.Info("timeout time: %v s", timeout.Seconds())

	mediaChan := make(chan media.Media)

	dir, _ := os.Getwd()
	go func() {
		i := 0
		for {
			m := <-mediaChan
			mtype := ""
			s := strings.Split(m.Type, "/")
			if len(s) == 2 {
				mtype = s[1]
			} else {
				//直接从url获取文件名
				s2 := strings.Split(m.Url, "/")
				m.Name = m.Name + "-" + s2[len(s2)-1]
			}
			if len(mtype) > 0 {
				m.Name = m.Name + "." + mtype
			}
			m.Name = FileNameFix(m.Name)
			logs.Info("start down [%s] url: %s", m.Name, m.Url)

			outpath := strings.Join([]string{dir, mediaType, m.Name}, "/")
			err := media.DownFile(outpath, m.Url)
			if err != nil {
				logs.Error("down error: %s", err.Error())
			} else {
				logs.Info("down success path: %s", outpath)
			}
			if i == 3 {
				// close(fired)
				ctx.Done()
				return
			}
			i++
		}
	}()

	time.Sleep(time.Second)
	//新建标签页获取媒体URL地址
	newCtx, _ := chromedp.NewContext(ctx)
	if err := chromedp.Run(newCtx, ccdp.GenGetMediaTask(url, timeout, mediaChan)); err != nil {
		logs.Error("Get media error: %s", err.Error())
	}
}
