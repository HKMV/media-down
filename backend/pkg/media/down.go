package media

import (
	"media-down/backend/pkg/logs"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	client *resty.Client
)

func init() {
	client = resty.New().
		SetRetryCount(1).
		SetTimeout(10 * time.Second)
	//SetHeader(pgt.UserAgentHeaderName, pgt.UserAgentHeaderValue).
	//SetHeader(pgt.OriginHeaderName, pgt.GeekBang).
	//SetLogger(logger.DiscardLogger{})
}

func DownFile(outpath, url string) (err error) {
	resp, err := client.R().
		//SetContext(ctx).
		SetOutput(outpath).
		Get(url)
	if err != nil {
		return err
	}
	logs.Info("download resp: %v", resp)
	return nil
}
