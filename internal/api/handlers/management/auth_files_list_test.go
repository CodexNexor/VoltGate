// Package management — management internal logic
package management

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/voltgate/voltgate/v6/internal/registry"
	coreauth "github.com/voltgate/voltgate/v6/sdk/voltgate/auth"
)

func TestBuildAuthFileEntry_IncludesRuntimeDetailsForDashboard(t *testing.T) {
	authDir := t.TempDir()
	authPath := filepath.Join(authDir, "antigravity-user@example.com.json")
	if err := os.WriteFile(authPath, []byte(`{"type":"antigravity"}`), 0o600); err != nil {
		t.Fatalf("failed to write auth file: %v", err)
	}

	nextRetry := time.Now().Add(3 * time.Minute).UTC()
	expiresAt := time.Now().Add(2 * time.Hour).UTC()
	lastRefresh := time.Now().Add(-15 * time.Minute).UTC()

	record := &coreauth.Auth{
		ID:             "antigravity-user@example.com.json",
		FileName:       "antigravity-user@example.com.json",
		Provider:       "antigravity",
		Label:          "user@example.com",
		Prefix:         "team-a",
		ProxyURL:       "http://127.0.0.1:8080",
		Status:         coreauth.StatusActive,
		StatusMessage:  "ready",
		LastRefreshedAt: lastRefresh,
		NextRetryAfter: nextRetry,
		Quota: coreauth.QuotaState{
			Exceeded:      true,
			Reason:        "quota",
			NextRecoverAt: nextRetry,
			BackoffLevel:  2,
		},
		Attributes: map[string]string{
			"path": authPath,
		},
		Metadata: map[string]any{
			"email":      "user@example.com",
			"project_id": "project-123",
			"expired":    expiresAt.Format(time.RFC3339),
		},
		ModelStates: map[string]*coreauth.ModelState{
			"gemini-3-pro": {
				Status:         coreauth.StatusError,
				StatusMessage:  "quota",
				Unavailable:    true,
				NextRetryAfter: nextRetry,
				Quota: coreauth.QuotaState{
					Exceeded:      true,
					Reason:        "quota",
					NextRecoverAt: nextRetry,
					BackoffLevel:  1,
				},
				UpdatedAt: time.Now().UTC(),
			},
		},
	}

	registry.GetGlobalRegistry().RegisterClient(record.ID, record.Provider, []*registry.ModelInfo{
		{ID: "gemini-3-pro"},
		{ID: "gemini-3.1-flash"},
	})
	t.Cleanup(func() {
		registry.GetGlobalRegistry().UnregisterClient(record.ID)
	})

	handler := &Handler{}
	entry := handler.buildAuthFileEntry(record)
	if entry == nil {
		t.Fatal("expected auth entry, got nil")
	}

	if got, _ := entry["project_id"].(string); got != "project-123" {
		t.Fatalf("project_id = %q, want %q", got, "project-123")
	}
	if got, _ := entry["prefix"].(string); got != "team-a" {
		t.Fatalf("prefix = %q, want %q", got, "team-a")
	}
	if got, _ := entry["proxy_url"].(string); got != "http://127.0.0.1:8080" {
		t.Fatalf("proxy_url = %q, want %q", got, "http://127.0.0.1:8080")
	}
	if got, _ := entry["models_count"].(int); got != 2 {
		t.Fatalf("models_count = %d, want 2", got)
	}

	modelsPreview, ok := entry["models_preview"].([]string)
	if !ok {
		t.Fatalf("models_preview type = %T, want []string", entry["models_preview"])
	}
	if len(modelsPreview) != 2 {
		t.Fatalf("models_preview length = %d, want 2", len(modelsPreview))
	}
	if modelsPreview[0] != "gemini-3-pro" || modelsPreview[1] != "gemini-3.1-flash" {
		t.Fatalf("models_preview = %#v, want sorted model ids", modelsPreview)
	}

	quota, ok := entry["quota"].(coreauth.QuotaState)
	if !ok {
		t.Fatalf("quota type = %T, want coreauth.QuotaState", entry["quota"])
	}
	if !quota.Exceeded || quota.Reason != "quota" {
		t.Fatalf("quota = %#v, want exceeded quota state", quota)
	}

	modelStates, ok := entry["model_states"].(map[string]*coreauth.ModelState)
	if !ok {
		t.Fatalf("model_states type = %T, want map[string]*coreauth.ModelState", entry["model_states"])
	}
	state, ok := modelStates["gemini-3-pro"]
	if !ok || state == nil {
		t.Fatalf("expected gemini-3-pro model state, got %#v", modelStates)
	}
	if !state.Unavailable {
		t.Fatalf("state.Unavailable = false, want true")
	}

	expiresValue, ok := entry["expires_at"].(time.Time)
	if !ok {
		t.Fatalf("expires_at type = %T, want time.Time", entry["expires_at"])
	}
	if expiresValue.IsZero() {
		t.Fatal("expires_at should not be zero")
	}
}
