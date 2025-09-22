// Package gemini — gemini internal logic
package gemini

import (
	. "github.com/voltgate/voltgate/v6/internal/constant"
	"github.com/voltgate/voltgate/v6/internal/interfaces"
	"github.com/voltgate/voltgate/v6/internal/translator/translator"
)

func init() {
	translator.Register(
		Gemini,
		Codex,
		ConvertGeminiRequestToCodex,
		interfaces.TranslateResponse{
			Stream:     ConvertCodexResponseToGemini,
			NonStream:  ConvertCodexResponseToGeminiNonStream,
			TokenCount: GeminiTokenCount,
		},
	)
}
