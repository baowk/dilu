package sms

import (
	"math/rand/v2"
	"strings"
)

type SmsSend interface {
	Send(phone string, code string, tmpId string)
}

var SMSSend SmsSend

func Setup(smsSend SmsSend) {
	SMSSend = smsSend
}

func GenerateSmsCode(length int) string {
	numberic := [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(numberic[rand.IntN(len(numberic))])
	}
	return sb.String()
}
