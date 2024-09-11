# fingerprints_db
基于luna框架的浏览器指纹部分代码示例|浏览器指纹设置示例｜fingerprints
这是一个简单的示例代码,介绍如何使用luna设置浏览器指纹；


原项目地址:https://github.com/musiclover789/luna

## 前提条件



- 已经有授权文件

- 有代理IP，可以是https\http\sockts5类型的。如果您的IP 是域名方式的，您需要先获取您代理IP的具体IP

  否则webrtc将出现IP不匹配的情况。

  

## 限制

目前显卡信息、示例中并不是特别多，您可以自行增加。







## 如何使用

1、将本项目，拷贝到luna框架的源代码里面

或者如果您是引用包的方式，就直接拷贝这个代码包到您自己的项目里面即可。



2、找到示例代码的这个部分main.go 如下，根据注释内容，修改您成您自己的信息



```
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

func main() {
	luna_utils.KillProcess()
	/***
   目前示例中提供的指纹设置的关键代码是这一行，
   Fingerprint: fingerprint.GetFingerprintNum("118.27.40.147", i),
   这个函数，您需要传入2个参数；
   第一个参数是您的代理IP的具体IP
   第二个是1-99的数字，也就意味着您可以产生1-99的套指纹，如果您需要更多的指纹，您可以自行修改代码逻辑。
   这一版本暂时只提供这个示例。将根据用户数量，考虑，是否会增加手机指纹如何设置的代码，和扩大这个随机指纹数量。
   
   
   
	*/
	/********************************/

	for i := 2; i < 3; i++ {
		time.Sleep(time.Second)
		go func(i int) {
		//这个路径您可以换成您自己的luna浏览器路径
			chromiumPath := "C:\\workspace\\chrome\\chrome\\src\\out\\Default\\chrome.exe"
			_, browserObj := devtools.NewBrowser(chromiumPath, &devtools.BrowserOptions{
			  //这个路径您也需要换成您自己设置的缓存路径
				CachePath:   
				luna_utils.CreateCacheDirInSubDir("C:\\workspace\\luna\\luna_new\\luna_new_case\\cache\\"),
				Fingerprint: fingerprint.GetFingerprintNum("111.1.40.111", i),
				Headless:    false,
				WindowSize: &devtools.WindowSize{
					Width:  1496,
					Height: 967,
				},
				ProxyStr: "http://uname:password@111.1.40.111:3128",
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

```






