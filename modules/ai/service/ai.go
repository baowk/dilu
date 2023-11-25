package service

import (
	"dilu/common/codes"
	"dilu/common/config"
	"dilu/modules/ai/platform/ali"
	"dilu/modules/ai/service/dto"

	"github.com/baowk/dilu-core/core/base"
)

type AiService struct {
	*base.BaseService
}

var SerAi = AiService{
	//base.NewService(consts.DB_CRM),
}

func (e *AiService) Chat(req dto.AiMsg, reqId string, msg *string) error {
	if len(req.Messages) == 0 {
		return codes.ErrInvalidParameter(reqId, "message is null")
	}
	if req.Platform == "ali" {
		model := ali.QwenMax
		if req.ModelName == "qwen-turbo" {
			model = ali.QwenTurbo
		} else if req.ModelName == "qwen-plus" {
			model = ali.QwenPlus
		}
		token := config.Ext.Ai.Ali.SK
		if token == "" {
			return codes.ErrInvalidParameter(reqId, "token is null")
		}
		params := ali.Parameters{
			ResultFormat: "message",
			// RepetitionPenalty: 1.0,
			// EnableSearch:      true,
			// MaxTokens:         100,
			// TopP:              0.9,
			// Seed:              0,
			// TopK:              1,
			// IncrementalOutput: true,
		}
		resp, err := ali.Qwen(token, model, req.Messages, params)
		if err != nil {
			return err
		} else {
			for _, choice := range resp.Output.Choices {
				*msg += choice.Message.Content + "\n"
			}
		}
	}
	return nil
}
