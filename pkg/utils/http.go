package utils

import (
	"bytes"
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

// HttpPostT 带超时的http post
func HttpPostT(ctx context.Context, addr string, data url.Values, timeout int) (ret []byte, err error) {
	if timeout < 0 {
		timeout = 3
	}

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*time.Duration(timeout)) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				_ = conn.SetDeadline(time.Now().Add(time.Second * time.Duration(timeout))) //设置发送接受数据超时
				return conn, nil
			},
		},
	}

	res, err := client.PostForm(addr, data)
	if err != nil {
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	return body, nil
}

// HttpPostJsonDataT 带超时的http post json string as data
func HttpPostJsonDataT(addr string, jsonDataBytes []byte, timeout int) (ret []byte, err error) {
	if timeout < 0 {
		timeout = 3
	}

	req, err := http.NewRequest("POST", addr, bytes.NewBuffer(jsonDataBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*time.Duration(timeout)) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				_ = conn.SetDeadline(time.Now().Add(time.Second * time.Duration(timeout))) //设置发送接受数据超时
				return conn, nil
			},
		},
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

//带超时的http get
func HttpGetT(addr string, timeout int) (ret []byte, err error) {
	if timeout < 0 {
		timeout = 3
	}

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*time.Duration(timeout)) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				_ = conn.SetDeadline(time.Now().Add(time.Second * time.Duration(timeout))) //设置发送接受数据超时
				return conn, nil
			},
		},
	}

	res, err := client.Get(addr)
	if err != nil {
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	return body, nil
}
