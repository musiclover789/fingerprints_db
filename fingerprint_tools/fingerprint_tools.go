package fingerprint_tools

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(min, max int) int {
	// 设置随机种子为当前时间的纳秒数
	rand.Seed(time.Now().UnixNano())

	// 生成并返回范围内的随机数
	return rand.Intn(max-min+1) + min
}
