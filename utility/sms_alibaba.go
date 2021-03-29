package utility

import (
	"encoding/json"
	"fmt"

	aliyunsmsclient "github.com/KenmyZhang/aliyun-communicate"
)

const (
	gatewayUrl = "http://dysmsapi.aliyuncs.com/"
)

var (
	accessKeyId     string
	accessKeySecret string
	signName        string

	smsClient = aliyunsmsclient.New(gatewayUrl)
)

func SetupSMS(sign, Key, Secret string) {
	accessKeyId = Key
	accessKeySecret = Secret
	signName = sign
}

func SendSMS(phoneNumber, templateCode, templateParam string) error {
	result, err := smsClient.Execute(accessKeyId, accessKeySecret, phoneNumber, signName, templateCode, templateParam)
	if err != nil {
		return err
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return err
	}

	if result.IsSuccessful() {
		fmt.Println("A SMS is sent successfully:", string(resultJson))
		return nil
	} else {
		return fmt.Errorf("Failed to send a SMS:%s", string(resultJson))
	}
}
