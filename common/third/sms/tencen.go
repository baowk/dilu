package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/baowk/dilu-core/common/utils/regexp_util"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	terrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
)

func Send(phone string, code string) error {
	// 实例化一个认证对象，入参需要传入腾讯云账户密钥对secretId,secretKey.
	credential := common.NewCredential(
		"",
		"",
	)

	cpf := profile.NewClientProfile()

	cpf.HttpProfile.ReqMethod = "POST"

	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"

	client, _ := sms.NewClient(credential, "ap-beijing", cpf)

	/* 实例化一个请求对象，根据调用的接口和实际情况*/
	request := sms.NewSendSmsRequest()

	// 应用 ID 可前往 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage) 查看
	request.SmsSdkAppId = common.StringPtr("1400816005")

	// 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名
	request.SignName = common.StringPtr("辰链信息科技")

	/* 模板 ID: 必须填写已审核通过的模板 ID */
	request.TemplateId = common.StringPtr("1779176")

	/* 模板参数: 模板参数的个数需要与 TemplateId 对应模板的变量个数保持一致，若无模板参数，则设置为空*/
	request.TemplateParamSet = common.StringPtrs([]string{code})

	/* 下发手机号码，采用 E.164 标准，+[国家或地区码][手机号]
	 * 示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号，最多不要超过200个手机号*/
	if regexp_util.CheckMobile(phone) {
		phone = "+86" + phone
	} else {
		return errors.New("")
	}
	request.PhoneNumberSet = common.StringPtrs([]string{phone})

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.SendSms(request)
	// 处理异常
	if _, ok := err.(*terrors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return err
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。

	b, _ := json.Marshal(response.Response)
	// 打印返回的json字符串
	fmt.Printf("%s", b)

	return err

}

// GenerateSmsCode 生成验证码;length代表验证码的长度
func GenerateSmsCode(length int) string {
	numberic := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().Unix())
	var sb strings.Builder
	for i := 0; i < length; i++ {
		fmt.Fprintf(&sb, "%d", numberic[rand.Intn(len(numberic))])
	}
	return sb.String()
}
