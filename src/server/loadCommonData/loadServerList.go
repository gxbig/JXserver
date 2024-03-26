package loadCommonData

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"server/conf"
	"server/msg"
	"server/tool"
	"server/util"
	"time"
)

// 初始获取服务器地址
func init() {
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop() // 确保计时器停止
		util.ServerList = HttpGetServerList()
		// 使用无限循环来处理定时任务
		for range ticker.C {
			if util.ServerList != nil {
				return
			}
			util.ServerList = HttpGetServerList()
		}
	}()
}

// HttpGetServerList 获取所有服务器
func HttpGetServerList() []msg.Server {
	// 目标URL
	targetURL, err := url.Parse(conf.Server.HttpClientAddr + "/severGetServerList")
	if err != nil {
		tool.Error("http获取全部服务器信息失败:" + err.Error())
		return nil
	}
	// 创建请求
	request := &http.Request{
		Method: "GET",
		URL:    targetURL,
		Header: http.Header{
			"User-Agent": []string{"MyClient/0.1"},     // 设置User-Agent请求头
			"Accept":     []string{"application/json"}, // 设置Accept请求头
		},
	}

	// 发送请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		tool.Error("http获取全部服务器-获取请求:" + err.Error())
		return nil
	}

	// 读取响应体
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			tool.Error("http获取全部服务器信息请求异常:" + err.Error())
		}
	}(response.Body)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		tool.Error("http获取全部服务器-读取请求:" + err.Error())
		return nil
	}

	result := &requestResults{}
	err = json.Unmarshal(body, result)
	var serverList = make([]msg.Server, 10)
	if result.Code == "200" {
		//取内部数据
		err = json.Unmarshal(*result.Data, &serverList)
		tool.Debug("http获取全部服务器-成功:", serverList)
		return serverList
	} else {
		tool.Debug("获取用户异常:%s:", result.Message)

	}
	return nil
}
