// Package geminiCLI — gemini-cli internal logic
package geminiCLI

import (
	. "github.com/voltgate/voltgate/v6/internal/constant"
	"github.com/voltgate/voltgate/v6/internal/interfaces"
	"github.com/voltgate/voltgate/v6/internal/translator/translator"
)

func init() {
	translator.Register(
		GeminiCLI,
		Claude,
		ConvertGeminiCLIRequestToClaude,
		interfaces.TranslateResponse{
			Stream:     ConvertClaudeResponseToGeminiCLI,
			NonStream:  ConvertClaudeResponseToGeminiCLINonStream,
			TokenCount: GeminiCLITokenCount,
		},
	)
}
