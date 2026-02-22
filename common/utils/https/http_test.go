package https

import (
	"fmt"
	"testing"
)

func TestAddHeader(t *testing.T) {
	res, err := New().AddHeader("Content-Type", "application/json").Get("https://www.baidu.com")
	if err != nil {
		t.Errorf("Req err %v", err)
	}
	fmt.Println(string(res))
}
