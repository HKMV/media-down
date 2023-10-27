package cmd

import (
	"context"
	"media-down/backend/pkg/logs"
	"media-down/backend/pkg/media"
	"os"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func Rod(url, mediaType string, waitTime time.Duration, headless bool) {
	ctx := context.Background()
	l := launcher.New().Context(ctx).Headless(headless).MustLaunch()
	page := rod.New().ControlURL(l).MustConnect().MustPage("")
	rodDownMedia(url, mediaType, waitTime, page)
	ctx.Deadline()
}

func rodDownMedia(url, mediaType string, timeout time.Duration, page *rod.Page) {
	mediaChan := make(chan media.Media)

	dir, _ := os.Getwd()
	go func() {
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
				continue
			}
			logs.Info("down success path: %s", outpath)
		}
	}()

	router := page.HijackRequests()
	router.MustAdd("*", func(ctx *rod.Hijack) {
		// 你可以使用很多其他 enum 类型，比如 NetworkResourceTypeScript 用于 javascript
		// 这个例子里我们使用 NetworkResourceTypeImage 来阻止图片
		if ctx.Request.Type() == proto.NetworkResourceTypeMedia {
			m := media.Media{
				Url:  ctx.Request.URL().String(),
				Name: page.MustInfo().Title,
			}
			//真正的发送请求等待响应,获取真正的响应信息
			ctx.MustLoadResponse()
			m.Type = ctx.Response.Headers().Get("Content-Type")

			mediaChan <- m
			return
		}
		ctx.ContinueRequest(&proto.FetchContinueRequest{})
	})
	// 因为我们只劫持特定页面，即便不使用 "*" 也不会太多性能影响
	go router.Run()

	page.Timeout(timeout).MustNavigate(url).MustWaitStable()
}
