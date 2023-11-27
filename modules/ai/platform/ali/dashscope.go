package ali

import (
	"dilu/modules/ai/service/dto"
	"encoding/json"
	"fmt"

	"github.com/baowk/dilu-core/common/utils/https"
	"github.com/baowk/dilu-core/core"
	"go.uber.org/zap"
)

const (
	API_URL = "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"
	// QwenTurbo 问问机器人
	QwenTurbo       DashscopeModel = "qwen-turbo"
	QwenPlus        DashscopeModel = "qwen-plus"
	QwenMax         DashscopeModel = "qwen-max"
	Qwen14bChat     DashscopeModel = "qwen-14b-chat"
	Qwen7bChat      DashscopeModel = "qwen-7b-chat"
	Llama27bChatV2  DashscopeModel = "llama2-7b-chat-v2"
	Llama213bChatV2 DashscopeModel = "llama2-13b-chat-v2"
)

type DashscopeModel string

type DashscopeReq struct {
	Model string `json:"model"`
	Input struct {
		Messages []dto.Message `json:"messages"`
	} `json:"input"`
	Parameters Parameters `json:"parameters,omitempty"`
}

// type Stop[T int | string] [][]T
type Parameters struct {
	ResultFormat      string     `json:"result_format,omitempty"`      // N "text"表示旧版本的text "message"表示兼容openai的message
	Seed              uint64     `json:"seed,omitempty"`               // N 生成时，随机数的种子，用于控制模型生成的随机性。如果使用相同的种子，每次运行生成的结果都将相同；当需要复现模型的生成结果时，可以使用相同的种子。seed参数支持无符号64位整数类型。默认值 1234
	MaxTokens         int        `json:"max_tokens,omitempty"`         // N 用于限制模型生成token的数量，max_tokens设置的是生成上限，并不表示一定会生成这么多的token数量。最大值和默认值均为1500
	TopP              float32    `json:"top_p,omitempty"`              // N 生成时，核采样方法的概率阈值。例如，取值为0.8时，仅保留累计概率之和大于等于0.8的概率分布中的token，作为随机采样的候选集。取值范围为（0,1.0)，取值越大，生成的随机性越高；取值越低，生成的随机性越低。默认值为0.8。注意，取值不要大于等于1
	TopK              int        `json:"top_k,omitempty"`              // N 生成时，采样候选集的大小。例如，取值为50时，仅将单次生成中得分最高的50个token组成随机采样的候选集。取值越大，生成的随机性越高；取值越小，生成的确定性越高。注意：如果top_k参数为空或者top_k的值大于100，表示不启用top_k策略，此时仅有top_p策略生效，默认是空。
	RepetitionPenalty float32    `json:"repetition_penalty,omitempty"` // N 用于控制模型生成时的重复度。提高repetition_penalty时可以降低模型生成的重复度。1.0表示不做惩罚。默认为1.1。
	Temperature       float32    `json:"temperature,omitempty"`        // N 用于控制随机性和多样性的程度。具体来说，temperature值控制了生成文本时对每个候选词的概率分布进行平滑的程度。较高的temperature值会降低概率分布的峰值，使得更多的低概率词被选择，生成结果更加多样化；而较低的temperature值则会增强概率分布的峰值，使得高概率词更容易被选择，生成结果更加确定。取值范围： (0, 2),系统默认值1.0
	Stop              [][]string `json:"stop,omitempty"`               // N 用于控制生成时遇到某些内容则停止。如果指定了字符串或者token_ids，模型将要生成指定字符串或者token_ids时会停止生成，生成结果不包含指定的内容。例如指定stop为"你好"，表示将要生成"你好"时停止；指定stop为[37763, 367]，表示将要生成"Observation"时停止。同时，stop参数支持以list方式传入字符串数组或者token_ids数组，以期支持使用多个stop的场景。注意，list模式下不支持字符串和token_ids混用，list模式下每个元素类型要相同。
	EnableSearch      bool       `json:"EnableSearch,omitempty"`       // N 生成时，是否参考搜索的结果。注意：打开搜索并不意味着一定会使用搜索结果；如果打开搜索，模型会将搜索结果作为prompt，进而“自行判断”是否生成结合搜索结果的文本，默认为false
	IncrementalOutput bool       `json:"incremental_output,omitempty"` // N 用于控制流式输出模式，默认False，即后面内容会包含已经输出的内容；设置为True，将开启增量输出模式，后面输出不会包含已经输出的内容，您需要自行拼接整体输出，参考流式输出示例代码。默认False：I I like i like apple True:I like apple 该参数只能与stream输出模式配合使用。
}

type Output struct {
	Text         string    `json:"text"`
	Choices      []Choices `json:"choices"`
	FinishReason string    `json:"finish_reason"`
}

type Choices struct {
	FinishReason string      `json:"finish_reason"`
	Message      dto.Message `json:"message"`
}

type Usage struct {
	OutputTokens int `json:"output_tokens"`
	InputTokens  int `json:"input_tokens"`
}

type DashscopeRes struct {
	Output     Output `json:"output"`
	Usage      Usage  `json:"usage"`
	RequestId  string `json:"request_id"`
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func Dashscope(token string, model DashscopeModel, messages []dto.Message, params Parameters) (resp DashscopeRes, err error) {
	req := DashscopeReq{
		Model: string(model),
		Input: struct {
			Messages []dto.Message `json:"messages"`
		}{
			Messages: messages,
		},
		Parameters: params,
	}
	// if params[0].MaxTokens <= 0 || params[0].MaxTokens > 1500 {
	// 	params[0].MaxTokens = 1500
	// }
	// if params[0].TopP <= 0 || params[0].TopP > 1.0 {
	// 	params[0].TopP = 0.8
	// }
	// if params[0].Temperature <= 0 || params[0].Temperature > 2.0 {
	// 	params[0].TopK = 1.0
	// }
	// req.Parameters = params[0]

	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	core.Log.Info(string(data))
	res, err := https.New().AddHeader("Content-Type", "application/json").AddHeader("Authorization", fmt.Sprintf("Bearer %s", token)).Post(API_URL, data)

	if err != nil {
		core.Log.Error("chat err ", zap.Error(err))
		return
	} else {
		core.Log.Info(string(res))
		err = json.Unmarshal(res, &resp)
		return
	}
}
