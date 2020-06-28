package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"math/rand"
	"strings"
	"time"
)

func fetch(url string) (retVal string, err error) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	err = chromedp.Run(ctx,
		chromedp.Emulate(device.IPhoneXR),
		chromedp.Sleep(time.Duration(3*1000*1000*1000+rand.Intn(5*1000*1000*1000-3*1000*1000*1000+1))),
		chromedp.Navigate(url),
		chromedp.WaitVisible(`main`, chromedp.ByID),
		chromedp.OuterHTML("html", &retVal),
	)
	if strings.Contains(retVal, "验证按钮") {
		err = fmt.Errorf("需要验证")
		return
	}

	if strings.Contains(retVal, "禁止访问") {
		err = fmt.Errorf("禁止访问")
		return
	}
	return
}
