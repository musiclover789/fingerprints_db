package fingerprint

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*
**
设计目的：计算无限靠谱的指纹库

因为涉及到网络、所以这个函数未必是固定的

proxy 根据proxy来计算、时区、语言、地区 已经国际API

platform=win64 这个目前只针对windows设计；

根据当前版本计算出相近的确实存在的 useragent 来计算userAgentData

列出一些列相近的显卡库；这个暂时

列出一些比较正常的screen

内存、硬盘、也考虑显卡来进行大体调整或者不调整

cavans webgl 则相对随机;

声卡、字体、插件 指纹权重几乎太低，在这里不予考虑。

查询IP 可以去官网免费注册 获取；每个月查5w次是免费的。
https://ipinfo.io/account/home

https://ipinfo.io/104.234.230.91/?token=a31d3a598a4c1c
查询IP 所属
*/

func GetFingerprintNum(proxyIP string, num int) []string {
	// 调用 getTimezone 函数获取 timezone
	timezone, _, err := GetTimezone(proxyIP)
	offset, err := GetTimeZoneOffset(timezone)
	offset = offset * 3600 * 1000
	if err != nil {
		fmt.Println("Error:", err) // 输出错误信息
		return nil
	}

	fmt.Println("Timezone:", timezone) // 输出解析的 timezone 字符串
	userAgent, fullVersion, majorVersion := GenerateUserAgentByNum(num)
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/" + majorVersion + ".0.0.0 Safari/537.36"
	userAgentData := GenerateUserAgentData(majorVersion, fullVersion)
	//---
	var arr = []string{}
	arr = append(arr, "--luna_user_agent="+userAgent)
	arr = append(arr, "--luna_header_set=true")
	arr = append(arr, "--luna_header_1=User-Agent-lunareplace-"+userAgent)
	arr = append(arr, "--luna_header_2=Accept-Language-lunareplace-"+MapTimezoneToLanguage(timezone))
	//meta-data
	arr = append(arr, `--luna_header_3=sec-ch-ua-full-version-lunareplace-"`+fullVersion+`"`)
	arr = append(arr, `--luna_header_5=Sec-Ch-Ua-lunareplace-"Chromium";v="`+majorVersion+`", "Google Chrome";v="`+majorVersion+`", "Not-A.Brand";v="99`)
	arr = append(arr, `--luna_header_8=Sec-Ch-Ua-Full-Version-List-lunareplace-"Chromium";v="`+fullVersion+`", "Google Chrome";v="`+fullVersion+`", "Not-A.Brand";v="99.0.0.0"`)

	arr = append(arr, `--luna_header_9=luna-lunaadd-"`+fullVersion+`"`)
	arr = append(arr, userAgentData)

	arr = append(arr, "--luna_platform=Win32")
	//time zone
	arr = append(arr, "--luna_timezone="+timezone)
	arr = append(arr, "--luna_timezone_offset="+strconv.Itoa(offset))

	arr = append(arr, "--luna_languages="+MapTimezoneToLanguage(timezone))
	arr = append(arr, "--luna_language="+MapTimezoneToLanguage(timezone))
	arr = append(arr, "--luna_deviceMemory=8")
	arr = append(arr, "--luna_hardwareConcurrency=16")
	arr = append(arr, "--luna_cavans_random_int_number="+strconv.Itoa(num))    //+
	arr = append(arr, "--luna_audio_random_int_number="+strconv.Itoa(num))     //+
	arr = append(arr, "--luna_client_rects_int_number="+strconv.Itoa(num*num)) //+

	//webrtc
	//建议不设置IP6公网IP，而是将自己的网络禁用IP6 原因是即便设置了，他依然可以通过你设置的IP6找到你的地区、可能会造成其他指纹并不一致
	//arr = append(arr, "--luna_webrtc_public_ip6="+"2409:8a5e:aa9f:8a4:1869:e970:2645:a62b")
	arr = append(arr, "--luna_webrtc_public_ip="+proxyIP)
	arr = append(arr, "--luna_webrtc_local_ip6_ip="+GenerateRandomIPv6())

	//fonts //这里不做设置,意义不大
	//arr = append(arr, `--luna_font_white_list=Tahoma,Arial,Helvetica,arial,Arial Black,Arial Narrow,Bahnschrift,Bahnschrift Light,Bahnschrift SemiBold,Calibri,Calibri Light,Cambria,Cambria Math,Candara,Candara Light,Comic Sans MS,Consolas,Constantia,Corbel,Corbel Light,Courier,Courier New,Ebrima,Gadugi,Leelawadee UI,Segoe UI Emoji,Segoe UI Historic,Franklin Gothic Heavy,Franklin Gothic Medium,Gabriola,Georgia,Ink Free,Javanese Text,Lucida Console,Lucida Sans Unicode,MS Gothic,MS PGothic,MS UI Gothic,MS Sans Serif,Microsoft Sans Serif,MS Serif,Times,Times New Roman,MV Boli,Malgun Gothic,Marlett,Microsoft Himalaya,Microsoft JhengHei,Microsoft JhengHei Regular,Microsoft JhengHei Light,Microsoft JhengHei UI,Microsoft JhengHei UI Regular,Microsoft JhengHei UI Light,Microsoft New Tai Lue,Microsoft PhagsPa,Microsoft Tai Le,Microsoft YaHei Light,Microsoft YaHei UI,Microsoft YaHei UI Light,Microsoft Yi Baiti,SimSun-ExtB,MingLiU-ExtB,MingLiU_HKSCS-ExtB,Mongolian Baiti,Myanmar Text,Nirmala UI,PMingLiU-ExtB,Palatino Linotype,Segoe MDL2 Assets,Segoe Print,Segoe Script,Segoe UI Black,Segoe UI Light,Segoe UI Semibold,Segoe UI Symbol,Sitka Banner,Sitka Display,Sitka Heading,Sitka Small,Sitka Subheading,Sitka Text,Sylfaen,Symbol,Trebuchet MS,Verdana,Webdings,Wingdings,Yu Gothic,Yu Gothic Regular,Yu Gothic Light,Yu Gothic Medium,Yu Gothic UI,Yu Gothic UI Regular,Yu Gothic UI Light,Yu Gothic UI Semibold,Franklin Gothic,PMingLiU,Impact,Microsoft YaHei,SimSun,Gulim,MingLiU,MingLiU_HKSCS,Gabriola Regular,Impact Regular,Javanese Text Regular,Lucida Console Regular,Lucida Sans Unicode Regular,Microsoft Himalaya Regular,Microsoft Sans Serif Regular,Microsoft Yi Baiti Regular,MingLiU_HKSCS-ExtB Regular,MingLiu-ExtB Regular,MS Gothic Regular,MS PGothic Regular,MS UI Gothic Regular,MV Boli Regular,NSimSun Regular,PMingLiU-ExtB Regular,Segoe MDL2 Assets Regular,Segoe UI Emoji Regular,Segoe UI Historic Regular,Segoe UI Symbol Regular,SimSun Regular,SimSun-ExtB Regular,Sylfaen Regular,Symbol Regular,Webdings Regular,Wingdings Regular,NSimSun,system-uiLeelawadee,Old English Text MT,Imprint MT Shadow,Californian FB,Gill Sans MT Condensed,Wingdings 2,Juice ITC,SimHei,Engravers MT,Rockwell Condensed,Matura MT Script Capitals,Lucida Sans,Playbill,Castellar,Tw Cen MT Condensed,Lucida Sans Typewriter,Monotype Corsiva,Harrington,High Tower Text,Baskerville Old Face,Jokerman,Mistral,Wingdings 3,Goudy Stout,Cooper Black,Berlin Sans FB,Blackadder ITC,Wide Latin,Papyrus Condensed,Elephant,Papyrus,DejaVu Sans Mono,Stencil,Rockwell,Footlight MT Light,Goudy Old Style,Algerian,Edwardian Script ITC,Broadway,Brush Script MT,Poor Richard,Bell MT,MS Reference Specialty,FangSong,Agency FB,Calisto MT,Lucida Calligraphy,Tw Cen MT,Bernard MT Condensed,Informal Roman,Parchment,PMingLiU,Copperplate Gothic,STFangsong,Showcard Gothic,Century Gothic,Felix Titling,DengXian Light,Perpetua,Lucida Bright,Colonna MT,Ravie,HoloLens MDL2 Assets,Maiandra GD,Chiller,Vivaldi,Perpetua Titling MT,Niagara Solid,HoloLens MDL2 Assets Regular,STKaiti`)

	//也不做设置，
	//screen, devicePixelRatio := GetScreen()
	//arr = append(arr, "--luna_screen="+screen)
	//arr = append(arr, "--luna_devicePixelRatio="+devicePixelRatio)

	//webgl-显卡
	arr = append(arr, "--ignore-gpu-blocklist")
	arr = append(arr, "--enable-unsafe-webgpu")
	arr = append(arr, "--enable-webgpu-developer-features")
	for _, gl := range GetGLByNum(num) {
		arr = append(arr, gl)
	}

	for _, item := range arr {
		fmt.Println(item)
		fmt.Println()
	}
	fmt.Println("====================================")

	return arr
}

func RandomInRangeString(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	randomInt := min + rand.Intn(max-min+1)
	return strconv.Itoa(randomInt)
}
