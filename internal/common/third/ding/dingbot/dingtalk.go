package dingbot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// 获取时间戳和加密签名
func getKey(SecRet string) (string, string) {
	NowTime := time.Now().UnixNano() / 1e6
	timestamp := strconv.FormatInt(NowTime, 10)
	//拼接
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, SecRet)
	//创建一个新的使用哈希校验算法的hash.Hash接口
	hash := hmac.New(sha256.New, []byte(SecRet))
	//通过该哈希接口添加数据-添加数据
	hash.Write([]byte(stringToSign))
	//通过哈希接口计算结果
	signData := hash.Sum(nil)
	//url.QueryEscape字符串转义
	//将计算结果通过base64编码 转string格式
	sign := url.QueryEscape(base64.StdEncoding.EncodeToString(signData))
	return timestamp, sign
}

// SECf866f7a7b3968a1c898afc25d783f38f14783c5d3ef3f356c1e8bea0c59fe0f0
// https://oapi.dingtalk.com/robot/send?access_token=3d80d7a4c542e476ae9ab3af2f835e59503d4c960cdb021a797571c0da970e9d
// 发送
// messageContent发送内容
// access_token 机器人token
// SecRet
func Dingtalk(messageContent, access_token, SecRet string, atIds ...string) {
	timestamp, sign := getKey(SecRet)
	apiUrl := "https://oapi.dingtalk.com/robot/send?access_token=" + access_token + "&timestamp=" + timestamp + "&sign=" + sign
	//设置发送内容
	text := map[string]string{
		"content": messageContent,
	}

	postData := map[string]interface{}{
		"msgtype": "text",
		"text":    text,
	}

	//判断需要艾特
	if len(atIds) > 0 {
		postData["at"] = map[string][]string{
			"atUserIds": atIds,
		}
	}

	//类型转换
	data, _ := json.Marshal(postData)
	client := &http.Client{}
	//创建一个请求
	req, err := http.NewRequest("POST", apiUrl, bytes.NewReader(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	//添加标头
	req.Header.Add("Content-Type", "application/json")
	//执行调用
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	//读取返回内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

// 发送钉钉消息
func SendDingMsg(msg, accessToken, SecRet string) {
	Dingtalk(msg, accessToken, SecRet, "")
}
