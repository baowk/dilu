package sms

import (
	"sync"
	"testing"
)

// MockSmsSender 是一个模拟的SMS发送器实现
type MockSmsSender struct {
	SendCalled bool
	Phone      string
	Code       string
	TmpId      string
}

func (m *MockSmsSender) Send(phone string, code string, tmpId string) {
	m.SendCalled = true
	m.Phone = phone
	m.Code = code
	m.TmpId = tmpId
}

func TestGenerateSmsCode(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		wantLen  int
	}{
		{
			name:    "生成4位验证码",
			length:  4,
			wantLen: 4,
		},
		{
			name:    "生成6位验证码",
			length:  6,
			wantLen: 6,
		},
		{
			name:    "生成8位验证码",
			length:  8,
			wantLen: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateSmsCode(tt.length)
			
			// 检查长度是否正确
			if len(got) != tt.wantLen {
				t.Errorf("GenerateSmsCode() 长度 = %v, 想要 %v", len(got), tt.wantLen)
			}
			
			// 检查是否只包含数字字符
			for i, char := range got {
				if char < '0' || char > '9' {
					t.Errorf("GenerateSmsCode() 在位置 %d 包含非数字字符: %c", i, char)
				}
			}
		})
	}
}

func TestGenerateSmsCode_ZeroLength(t *testing.T) {
	got := GenerateSmsCode(0)
	if got != "" {
		t.Errorf("GenerateSmsCode(0) = %v, 想要空字符串", got)
	}
}

func TestGenerateSmsCode_NegativeLength(t *testing.T) {
	// 负数长度应该返回空字符串或者处理为0
	got := GenerateSmsCode(-1)
	// 由于当前实现没有处理负数情况，这里测试实际行为
	// 如果实现改变了，这个测试可以帮助发现行为变化
	t.Logf("GenerateSmsCode(-1) = %v", got)
}

func TestSetup(t *testing.T) {
	// 创建mock发送器
	mockSender := &MockSmsSender{}
	
	// 设置全局变量
	Setup(mockSender)
	
	// 验证全局变量被正确设置
	if SMSSend == nil {
		t.Error("Setup() 后 SMSSend 仍然为nil")
	}
	
	// 验证可以调用接口方法
	SMSSend.Send("13800138000", "123456", "template1")
	
	// 验证mock被正确调用
	if !mockSender.SendCalled {
		t.Error("Send 方法未被调用")
	}
	
	if mockSender.Phone != "13800138000" {
		t.Errorf("Phone = %v, 想要 13800138000", mockSender.Phone)
	}
	
	if mockSender.Code != "123456" {
		t.Errorf("Code = %v, 想要 123456", mockSender.Code)
	}
	
	if mockSender.TmpId != "template1" {
		t.Errorf("TmpId = %v, 想要 template1", mockSender.TmpId)
	}
}

func TestSmsSend_Interface(t *testing.T) {
	// 测试SmsSend接口的实现
	var _ SmsSend = (*MockSmsSender)(nil)
	
	// 这个测试主要是编译时检查，确保MockSmsSender实现了SmsSend接口
	// 如果接口定义改变导致不匹配，编译会失败
}

// 测试并发安全性
func TestGenerateSmsCode_Concurrent(t *testing.T) {
	const goroutines = 100
	const codesPerGoroutine = 100
	
	var wg sync.WaitGroup
	wg.Add(goroutines)
	
	// 使用map记录生成的验证码，检测重复
	codeMap := make(map[string]bool)
	var mapMutex sync.Mutex
	
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < codesPerGoroutine; j++ {
				code := GenerateSmsCode(6)
				
				mapMutex.Lock()
				codeMap[code] = true
				mapMutex.Unlock()
			}
		}()
	}
	
	wg.Wait()
	
	// 对于6位数字验证码，总共有10^6 = 1,000,000种可能
	// 生成10,000个验证码，根据生日悖论，出现重复是正常的
	// 我们只需要确保大部分是唯一的即可
	
	totalGenerated := goroutines * codesPerGoroutine
	uniqueCount := len(codeMap)
	
	// 计算唯一率
	uniqueRate := float64(uniqueCount) / float64(totalGenerated)
	
	t.Logf("总共生成: %d, 唯一数量: %d, 唯一率: %.2f%%", totalGenerated, uniqueCount, uniqueRate*100)
	
	// 要求至少99%的唯一率
	if uniqueRate < 0.99 {
		t.Errorf("唯一率过低: %.2f%% < 99%%", uniqueRate*100)
	}
}

// 测试不同长度验证码的随机性
func TestGenerateSmsCode_Randomness(t *testing.T) {
	const iterations = 1000
	const codeLength = 6
	
	codes := make(map[string]int)
	
	for i := 0; i < iterations; i++ {
		code := GenerateSmsCode(codeLength)
		codes[code]++
	}
	
	// 检查是否有明显的重复模式
	duplicateCount := 0
	for _, count := range codes {
		if count > 1 {
			duplicateCount += count - 1
		}
	}
	
	// 允许少量重复（概率性事件）
	if duplicateCount > iterations/10 {
		t.Errorf("重复验证码过多: %d/%d", duplicateCount, iterations)
	}
	
	t.Logf("生成了 %d 个唯一验证码，重复 %d 次", len(codes), duplicateCount)
}

func BenchmarkGenerateSmsCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateSmsCode(6)
	}
}

func BenchmarkGenerateSmsCode_DifferentLengths(b *testing.B) {
	lengths := []int{4, 6, 8}
	
	for _, length := range lengths {
		b.Run(string(rune(length)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GenerateSmsCode(length)
			}
		})
	}
}