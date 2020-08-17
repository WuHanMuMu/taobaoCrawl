package taobao

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"time"
)

// 获取商品列表
func List(searhWords []string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			expr:= cdp.TimeSinceEpoch(time.Now().Add(10 * time.Hour))
			_,err:= network.SetCookie("ass","xxxx").WithExpires(&expr).Do(ctx)
			return err
		}),
		chromedp.Navigate("https://www.taobao.com"),
	}
}

// 获取请求头
func MyHeaders(res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://www.whatsmyua.info/?a`),
		chromedp.WaitVisible(`#custom-ua-string`),
		chromedp.InnerHTML(`#custom-ua-string`, res),
		chromedp.Sleep(10 * time.Second),
	}
}
