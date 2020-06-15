package utils

import "time"

//TimeToTimestamp 将日期时间格式转换为时间戳
func TimeToTimestamp(layout, timestr string) (int64, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return -1, err
	}
	times, err := time.ParseInLocation(layout, timestr, loc)
	if err != nil {
		return -1, err
	}
	trantimestamp := times.Unix()
	return trantimestamp, nil
}

//TimeFormat 转换时间格式
func TimeFormat(oldLayout, newLayout, oldTimeStr string) (string, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return oldTimeStr, err
	}
	times, err := time.ParseInLocation(oldLayout, oldTimeStr, loc)
	if err != nil {
		return oldTimeStr, err
	}
	return times.Format(newLayout), nil
}
