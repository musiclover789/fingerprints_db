package fingerprint

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

/*
**
设计目的：计算无限靠谱的指纹库

proxy 根据proxy来计算、时区

查询IP 可以去官网免费注册 获取；每个月查5w次是免费的。
https://ipinfo.io/account/home


查询IP 所属
*/

// getTimezone 函数发送 HTTP 请求，并解析返回结果中的 timezone 字段
func GetTimezone(ip string) (string, string, error) {
	// 创建 HTTP 客户端
	client := &http.Client{
		Timeout: 5 * time.Second, // 设置超时时间为 5 秒钟
	}
	url := "https://ipinfo.io/" + ip + "/?token=a31d3a598a4c1c"
	if len(ip) == 0 {
		url = "https://ipinfo.io?token=a31d3a598a4c1c"
	}
	// 发送 HTTP GET 请求
	resp, err := client.Get(url)
	if err != nil {
		return "", "", err // 发生错误，返回错误信息
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err // 读取响应体失败，返回错误信息
	}

	// 解析 JSON 响应，提取 timezone 字段
	timezone := gjson.Get(string(body), "timezone").String()
	if timezone == "" {
		return "", "", errors.New("timezone not found") // 没有找到 timezone 字段，返回错误信息
	}
	ip = gjson.Get(string(body), "ip").String()

	return timezone, ip, nil // 返回解析的 timezone 字符串
}
