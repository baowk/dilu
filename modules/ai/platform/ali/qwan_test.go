package ali

import (
	"testing"
)

const token = "sk-"

func TestQwen(t *testing.T) {
	msg := []Message{
		{
			Role:    "user",
			Content: "你好",
		},
	}
	params := Parameters{
		// RepetitionPenalty: 1.0,
		// ResultFormat:      "text",
		// EnableSearch:      true,
		// MaxTokens:         100,
		// TopP:              0.9,
		// Seed:              0,
		// TopK:              1,
		// IncrementalOutput: true,
	}
	resp, err := Qwen(token, QwenMax, msg, params)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(resp)
	}
}
