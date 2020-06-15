package goutils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

//GetIPRange 得到ip范围 比如192.168.1.1-254
func GetIPRange(ipRange, splitStr string) (ipStart, ipEnd string) {
	ipArr := strings.Split(ipRange, splitStr)
	if len(ipArr) != 2 {
		return "", ""
	}
	ipStart = ipArr[0]
	ipEnd = ipArr[1]
	sections := strings.Split(ipStart, ".")
	if len(sections) != 4 {
		return "", ""
	}
	sections[3] = ipEnd
	ipEnd = fmt.Sprintf("%s.%s.%s.%s", sections[0], sections[1], sections[2], ipEnd)
	return
}

//IP2long 字符串IP转化为long IP
func IP2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

//Long2IP long IP转化为字符串
func Long2IP(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}

//InIPRange 判断某个ip是否在某个ip范围内 比如192.168.1.5 在 192.168.1.5-192.168.1.254段内
func InIPRange(ip string, ipRange string, isEqual bool) (inRange bool) {
	trial := net.ParseIP(ip)
	if trial.To4() == nil {
		return false
	}

	ip1 := net.ParseIP(strings.Split(ipRange, "-")[0])
	if ip1.To4() == nil {
		return false
	}
	ip2 := net.ParseIP(strings.Split(ipRange, "-")[1])
	if ip2.To4() == nil {
		return false
	}

	if isEqual {
		if bytes.Compare(trial, ip1) >= 0 && bytes.Compare(trial, ip2) <= 0 {
			return true
		}
	} else {
		if bytes.Compare(trial, ip1) > 0 && bytes.Compare(trial, ip2) < 0 {
			return true
		}
	}

	return false
}

//IsInSameSection 判断两个ip是否在同一个网段
func IsInSameSection(IP1, IP2, mask string) bool {
	if IP1 == "" || IP2 == "" || mask == "" {
		return false
	}

	maskLong := IP2long(mask)
	if IP2long(IP1)&maskLong == IP2long(IP2)&maskLong {
		return true
	}
	return false
}
