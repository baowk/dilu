package utils

import (
	"strings"

	"github.com/mozillazg/go-pinyin"
)

func GetPinyin(text string) string {
	return strings.Join(pinyin.LazyConvert(text, nil), "-")
}
