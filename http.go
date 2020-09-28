package goutils

//有关Http协议GET和POST请求的封装

import (
	"bytes"
	"net"
	"net/http"
)

//SocketHTTPRequest 通过远程WEB Socket套接字方式下发交换机操作指令,通过JSON传递参数和响应 (目前在使用的方式)
func SocketHTTPRequest(socketFile, httpURI string, msg []byte) (*http.Response, error) {
	addr := &net.UnixAddr{Name: socketFile, Net: "unix"}
	unixDial := func(_, _ string) (conn net.Conn, err error) {
		conn, err = net.DialUnix("unix", nil, addr)
		return
	}
	client := &http.Client{
		Transport: &http.Transport{
			Dial: unixDial,
		},
		Timeout: 0,
	}
	requestObj, _ := http.NewRequest("POST", "http://hostname"+httpURI, bytes.NewReader(msg))
	requestObj.Close = true
	resp, err := client.Do(requestObj)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
