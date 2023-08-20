package chromedp

import (
	"context"
	"fmt"
	"media-down/backend/pkg/logs"
	"media-down/backend/pkg/media"
	"sync"
	"time"

	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

var (
	ShowWindowContext, _ = chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.NoDefaultBrowserCheck,                   //不检查默认浏览器
			chromedp.Flag("headless", false),                 //开启图像界面,重点是开启这个
			chromedp.Flag("ignore-certificate-errors", true), //忽略错误
			chromedp.Flag("disable-web-security", true),      //禁用网络安全标志
			chromedp.NoFirstRun,                              //设置网站不是首次运行
			chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"), //设置UserAgent
		)...,
	)
)

func GenGetMediaTask(url string, timeout time.Duration, mediaChan chan media.Media) chromedp.Tasks {
	m := media.Media{}
	return chromedp.Tasks{
		runtime.Enable(),
		network.Enable(),
		// 开启响应拦截
		fetch.Enable().WithPatterns([]*fetch.RequestPattern{
			{
				URLPattern:   "*",
				ResourceType: network.ResourceTypeMedia,
				RequestStage: fetch.RequestStageResponse,
			},
		}),
		chromedp.ActionFunc(
			func(ctx context.Context) error {
				var (
					// once  sync.Once
					wg sync.WaitGroup
					// fired = make(chan struct{})
				)
				lctx, cancel := context.WithCancel(ctx)
				defer cancel()

				chromedp.ListenTarget(lctx, func(ev interface{}) {
					switch e := ev.(type) {
					case *fetch.EventRequestPaused:
						wg.Add(1)
						go func() { // convert javascript
							defer wg.Done()
							//fmt.Println(e.Request.URL)

							// mediaUrlChan <- fmt.Sprintf("%s:=%s", "media", e.Request.URL)
							m.Url = e.Request.URL
							he := e.ResponseHeaders
							for _, v := range he {
								if v.Name == "Content-Type" {
									m.Type = v.Value
								}
							}
							logs.Info("mediaType:%#v", m.Type)
							// logs.Info("type", e.Request.MixedContentType.String())
							chromedp.Run(ctx, chromedp.Title(&m.Name))
							mediaChan <- m
							// once.Do(func() { close(fired) })
						}()
					}
				})

				var err error
				go func(e *error) {
					_, _, errorText, err := page.Navigate(url).Do(ctx)
					if err != nil {
						*e = err
					}
					if errorText != "" {
						*e = fmt.Errorf("page load error %s", errorText)
					}
				}(&err)

				select {
				case <-time.After(timeout):
					logs.Warn("load timeout")
					if err := page.StopLoading().Do(ctx); err != nil {
						return err
					}
					// case <-fired:
				}
				wg.Wait()
				return err
			},
		),
		// chromedp.WaitVisible("body", chromedp.ByQuery),
		page.Close(),
	}
}
