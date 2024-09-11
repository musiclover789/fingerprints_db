package main

import (
	"fmt"
	"github.com/musiclover789/luna/base_devtools/page"
	"github.com/musiclover789/luna/devtools"
	fingerprint "github.com/musiclover789/luna/fingerprints_db/fingerprints"
	"github.com/musiclover789/luna/luna_utils"
	"github.com/musiclover789/luna/protocol"
	"sync"
	"time"
)

/*
官方文档
https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-navigate
*/
func main() {
	luna_utils.KillProcess()
	/***
	所有测试示例 均依托已下几个前提
	1、基于代理IP状态下测试 如http://uname:password@111.1.40.111:3128
	*/
	/********************************/

	for i := 2; i < 3; i++ {
		time.Sleep(time.Second)
		go func(i int) {
			chromiumPath := "C:\\workspace\\chrome\\chrome\\src\\out\\Default\\chrome.exe"
			_, browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
				CachePath:   luna_utils.CreateCacheDirInSubDir("C:\\workspace\\luna\\luna_new\\luna_new_case\\cache\\"),
				Fingerprint: fingerprint.GetFingerprintNum("118.27.40.147", i),
				Headless:    false,
				WindowSize: &devtools.WindowSize{
					Width:  1496,
					Height: 967,
				},
				ProxyStr: "http://razor:razor9200@118.27.40.147:3128",
			})
			//===================
			var wg sync.WaitGroup
			wg.Add(1)
			wg.Add(1)
			browserObj.OpenPageAndListen("https://abrahamjuliot.github.io/creepjs/", func(devToolsConn *protocol.Session) {
				devToolsConn.ShowLog(false)
				page.PageEnable(devToolsConn)
				devToolsConn.SubscribeOneTimeEvent("Page.loadEventFired", func(param interface{}) {
					fmt.Println("load ok")
					wg.Done()
				})
			})
			browserObj.OpenPageAndListen("https://www.browserscan.net/zh/", func(devToolsConn *protocol.Session) {
				devToolsConn.ShowLog(false)
				page.PageEnable(devToolsConn)
				devToolsConn.SubscribeOneTimeEvent("Page.loadEventFired", func(param interface{}) {
					fmt.Println("load ok")
					wg.Done()
				})
			})
			wg.Wait()
		}(i)
	}
	time.Sleep(time.Hour)

}
