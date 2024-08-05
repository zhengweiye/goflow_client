package goflow_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io"
	"net"
	"net/http"
	"time"
)

// 创建客户端
var httpClient = &http.Client{
	//请求超时时间
	Timeout: time.Second * 90,
	// 创建连接池
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 连接超时
			KeepAlive: 30 * time.Second, // 长连接超时时间
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接
		IdleConnTimeout:       90 * time.Second, // 空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, // tls握手超时时间
		ExpectContinueTimeout: 90 * time.Second, // 100-continue状态码超时时间
	},
}

type Result[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type PageData[T any] struct {
	List          []T   `json:"list"`          // 响应内容
	TotalElements int64 `json:"totalElements"` // 总记录数
	TotalPage     int64 `json:"totalPage"`     // 总页数
}

func httpPost[T any](client *Client, methodPath string, param any) (data Result[T], err error) {
	var body *bytes.Reader
	if param == nil {
		param = make(map[string]any)
	}
	paramBytes, err := json.Marshal(param)
	if err != nil {
		return
	}
	body = bytes.NewReader(paramBytes)

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/%s", client.host, "goflow", methodPath), body)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("appId", client.appId)
	request.Header.Set("appSecret", client.appSecret)
	request.Header.Set("env", client.env)
	resp, err := httpClient.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	//TODO 状态码判断
	if resp.StatusCode != 200 {
		err = fmt.Errorf("http请求返回的状态为%d", resp.StatusCode)
		return
	}

	//TODO 读取io
	resBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if len(resBytes) == 0 {
		err = fmt.Errorf("返回字节数组为空")
		return
	}

	//TODO 反序列化
	fmt.Printf("[工作流] 响应数据: %+v\n", string(resBytes))
	var m = make(map[string]any)
	err = json.Unmarshal(resBytes, &m)
	if err != nil {
		return
	}
	err = mapstructure.Decode(m, &data)
	if err != nil {
		return
	}
	return
}
