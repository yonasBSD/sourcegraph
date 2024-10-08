package openai

// openAIChatCompletionsRequestParameters request object for openAI chat endpoint.
// https://platform.openai.com/docs/api-reference/chat/create
type openAIChatCompletionsRequestParameters struct {
	Model            string             `json:"model"`                       // request.Model
	Messages         []message          `json:"messages"`                    // request.Messages
	Temperature      float32            `json:"temperature,omitempty"`       // request.Temperature
	TopP             float32            `json:"top_p,omitempty"`             // request.TopP
	N                int                `json:"n,omitempty"`                 // always 1
	Stream           bool               `json:"stream,omitempty"`            // request.Stream
	Stop             []string           `json:"stop,omitempty"`              // request.StopSequences
	MaxTokens        int                `json:"max_tokens,omitempty"`        // request.MaxTokensToSample
	PresencePenalty  float32            `json:"presence_penalty,omitempty"`  // unused
	FrequencyPenalty float32            `json:"frequency_penalty,omitempty"` // unused
	LogitBias        map[string]float32 `json:"logit_bias,omitempty"`        // unused
	User             string             `json:"user,omitempty"`              // unused
}

// openAICompletionsRequestParameters payload for openAI completions endpoint.
// https://platform.openai.com/docs/api-reference/completions/create
type openAICompletionsRequestParameters struct {
	Model            string             `json:"model"`                       // request.Model
	Prompt           string             `json:"prompt"`                      // request.Messages[0] - formatted prompt expected to be the only message
	Temperature      float32            `json:"temperature,omitempty"`       // request.Temperature
	TopP             float32            `json:"top_p,omitempty"`             // request.TopP
	N                int                `json:"n,omitempty"`                 // always 1
	Stream           bool               `json:"stream,omitempty"`            // request.Stream
	Stop             []string           `json:"stop,omitempty"`              // request.StopSequences
	MaxTokens        int                `json:"max_tokens,omitempty"`        // request.MaxTokensToSample
	PresencePenalty  float32            `json:"presence_penalty,omitempty"`  // unused
	FrequencyPenalty float32            `json:"frequency_penalty,omitempty"` // unused
	LogitBias        map[string]float32 `json:"logit_bias,omitempty"`        // unused
	Suffix           string             `json:"suffix,omitempty"`            // unused
	User             string             `json:"user,omitempty"`              // unused
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openaiUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type openaiChoiceDelta struct {
	Content string `json:"content"`
}

type openaiMessage struct {
	Content string `json:"content"`
}

type openaiChoice struct {
	Delta        openaiChoiceDelta `json:"delta"`
	Message      openaiMessage     `json:"message"`
	Role         string            `json:"role"`
	FinishReason string            `json:"finish_reason"`
}

type openaiResponse struct {
	// Usage is only available for non-streaming requests.
	Usage   openaiUsage    `json:"usage"`
	Model   string         `json:"model"`
	Choices []openaiChoice `json:"choices"`
}
