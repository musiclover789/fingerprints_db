package fingerprint

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var screenList = []string{
	"height: 720, width: 1280, availHeight: 680, availWidth: 1240, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.5",
	"height: 1080, width: 1920, availHeight: 1040, availWidth: 1880, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1",
	"height: 1366, width: 768, availHeight: 1326, availWidth: 724, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1",
	"height: 900, width: 1600, availHeight: 860, availWidth: 1560, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.5",
	"height: 800, width: 1280, availHeight: 760, availWidth: 1240, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.25",
	"height: 1440, width: 2560, availHeight: 1400, availWidth: 2520, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 2",
	"height: 1366, width: 768, availHeight: 1326, availWidth: 724, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1",
	"height: 1080, width: 1920, availHeight: 1040, availWidth: 1880, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 2",
	"height: 720, width: 1280, availHeight: 680, availWidth: 1240, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.5",
	"height: 1440, width: 2560, availHeight: 1400, availWidth: 2520, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 2",
	"height: 768, width: 1366, availHeight: 728, availWidth: 1326, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.25",
	"height: 900, width: 1600, availHeight: 860, availWidth: 1560, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.5",
	"height: 1080, width: 1920, availHeight: 1040, availWidth: 1880, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 2.25",
	"height: 800, width: 1280, availHeight: 760, availWidth: 1240, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.25",
	"height: 768, width: 1366, availHeight: 728, availWidth: 1326, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.5",
	"height: 1440, width: 2560, availHeight: 1400, availWidth: 2520, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 2",
	"height: 1080, width: 1920, availHeight: 1040, availWidth: 1880, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.5",
	"height: 768, width: 1366, availHeight: 728, availWidth: 1326, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.25",
	"height: 900, width: 1600, availHeight: 860, availWidth: 1560, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.5",
	"height: 800, width: 1280, availHeight: 760, availWidth: 1240, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.25",
	"height: 1440, width: 2560, availHeight: 1400, availWidth: 2520, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 2",
	"height: 768, width: 1366, availHeight: 728, availWidth: 1326, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.25",
	"height: 1080, width: 1920, availHeight: 1040, availWidth: 1880, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.5",
	"height: 768, width: 1366, availHeight: 728, availWidth: 1326, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.25",
	"height: 900, width: 1600, availHeight: 860, availWidth: 1560, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.5",
	"height: 720, width: 1280, availHeight: 680, availWidth: 1240, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.5",
	"height: 1080, width: 1920, availHeight: 1040, availWidth: 1880, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 2.25",
	"height: 768, width: 1366, availHeight: 728, availWidth: 1326, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.5",
	"height: 800, width: 1280, availHeight: 760, availWidth: 1240, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 1.25",
	"height: 1440, width: 2560, availHeight: 1400, availWidth: 2520, availLeft: 0, availTop: 0, internal: true, primary: true, top: 0, left: 0, scaleFactor: 2",
}

// ScreenProperties represents the screen properties
type ScreenProperties struct {
	Height      int
	Width       int
	AvailHeight int
	AvailWidth  int
	AvailLeft   int
	AvailTop    int
	Internal    bool
	Primary     bool
	Top         int
	Left        int
	ScaleFactor float64
}

// GenerateScreenProperties generates simulated screen properties
func GenerateScreenProperties() (string, string) {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Define ranges for screen properties based on real data statistics
	heightRange := []int{600, 800, 900, 1024, 1050, 1080, 1200, 1440, 1600, 1920, 2048, 2160, 2560, 2880, 3440, 3840}
	widthRange := []int{800, 1024, 1280, 1366, 1440, 1600, 1680, 1920, 2048, 2560, 2880, 3440, 3840, 4096, 5120, 7680}
	availHeightRange := []int{600, 800, 900, 1024, 1050, 1080, 1200, 1440, 1600, 1920, 2048, 2160, 2560, 2880, 3440, 3840}
	availWidthRange := []int{800, 1024, 1280, 1366, 1440, 1600, 1680, 1920, 2048, 2560, 2880, 3440, 3840, 4096, 5120, 7680}
	availLeftRange := []int{0, 50, 100, 200, 300, 400, 500}
	availTopRange := []int{0, 50, 100, 200, 300, 400, 500}
	scaleFactorRange := []float64{1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0}

	// Randomly select values from the defined ranges
	screenHeight := selectRandomValue(heightRange)
	screenWidth := selectRandomValue(widthRange)
	availHeight := selectRandomValue(availHeightRange)
	availWidth := selectRandomValue(availWidthRange)
	availLeft := selectRandomValue(availLeftRange)
	availTop := selectRandomValue(availTopRange)
	scaleFactor := selectRandomFloat(scaleFactorRange)

	//--luna_screen=height:803,width:360,availHeight:803,availWidth:360,availLeft:0,availTop:0,colorDepth:24,pixelDepth:24"
	// Format the properties as a string
	properties := fmt.Sprintf("height: %d, width: %d, availHeight: %d, availWidth: %d, availLeft: %d, availTop: %d,colorDepth:24,pixelDepth:24",
		screenHeight, screenWidth, availHeight, availWidth, availLeft, availTop)

	return properties, strconv.FormatFloat(scaleFactor, 'f', -1, 64)
}

// selectRandomValue randomly selects a value from a given range
func selectRandomValue(values []int) int {
	index := rand.Intn(len(values))
	return values[index]
}

// selectRandomFloat randomly selects a value from a given range of floats
func selectRandomFloat(values []float64) float64 {
	index := rand.Intn(len(values))
	return values[index]
}

// getScreen 从字符串数组中随机获取一个值，并去掉两端的空格
func GetScreen() (string, string) {
	//if len(screenList) == 0 {
	//	return "" // 数组为空，返回空字符串
	//}
	//rand.Seed(time.Now().UnixNano())            // 使用当前时间的纳秒部分作为随机种子
	//index := rand.Intn(len(screenList))         // 生成随机索引
	//return strings.TrimSpace(screenList[index]) // 去除两端空格后返回随机值
	return GenerateScreenProperties()
}
