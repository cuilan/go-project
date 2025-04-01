package utils

import (
	"time"
)

func Unit8ToTime(data []uint8) string {
	// 将[]uint8转换为字符串
	timeStr := string(data)

	// 解析字符串为时间
	t, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return ""
	}
	// 将UTC时间转换为北京时间（东八区，UTC+8）
	beijingTime := t.In(time.FixedZone("CST", 8*3600))
	return beijingTime.Format("20060102150405")
}

func CurrentTime() string {
	return time.Now().Format("20060102150405")
}

func FormatTimeToFloor10Minutes() string {
	now := time.Now().UTC()
	// 截断到最近的10分钟（向下取整）
	flooredTime := now.Truncate(10 * time.Minute)
	return flooredTime.Format("200601021504")
}
