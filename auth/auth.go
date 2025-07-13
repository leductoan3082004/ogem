package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
)

type VirtualKey struct {
	ID          string            `json:"id"`
	Key         string            `json:"key"`
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Models      []string          `json:"models,omitempty"`
	MaxTokens   *int64            `json:"max_tokens,omitempty"`
	MaxRequests *int64            `json:"max_requests,omitempty"`
	Budget      *float64          `json:"budget,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	ExpiresAt   *time.Time        `json:"expires_at,omitempty"`
	IsActive    bool              `json:"is_active"`
	UsageStats  *UsageStats       `json:"usage_stats,omitempty"`
	Permissions []string          `json:"permissions,omitempty"`
}

type UsageStats struct {
	TotalTokens   int64      `json:"total_tokens"`
	TotalRequests int64      `json:"total_requests"`
	TotalCost     float64    `json:"total_cost"`
	LastUsed      *time.Time `json:"last_used,omitempty"`
}

// AuthManager interface for authentication management
type AuthManager interface {
	Manager
}

// AuthContext holds authentication information
type AuthContext struct {
	UserID   string                 `json:"user_id"`
	TenantID string                 `json:"tenant_id,omitempty"`
	KeyID    string                 `json:"key_id,omitempty"`
	Claims   map[string]interface{} `json:"claims,omitempty"`
}

// GetAuthContextFromRequest extracts auth context from HTTP request
func GetAuthContextFromRequest(r *http.Request) *AuthContext {
	// Check for auth context in request context
	if authCtx := r.Context().Value("auth_context"); authCtx != nil {
		if ctx, ok := authCtx.(*AuthContext); ok {
			return ctx
		}
	}
	// Fallback: create basic auth context from headers
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		return nil
	}
	return &AuthContext{
		UserID:   userID,
		TenantID: r.Header.Get("X-Tenant-ID"),
		KeyID:    r.Header.Get("X-Key-ID"),
		Claims:   make(map[string]interface{}),
	}
}

// VirtualKeyManager interface for managing virtual keys
type VirtualKeyManager interface {
	CreateKey(ctx context.Context, req *KeyRequest) (*VirtualKey, error)
	ListKeys(ctx context.Context) ([]*VirtualKey, error)
	DeleteKey(ctx context.Context, keyID string) error
	GetKey(ctx context.Context, keyID string) (*VirtualKey, error)
	UpdateKey(ctx context.Context, keyID string, updates map[string]interface{}) (*VirtualKey, error)
}

type KeyRequest struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Models      []string          `json:"models,omitempty"`
	MaxTokens   *int64            `json:"max_tokens,omitempty"`
	MaxRequests *int64            `json:"max_requests,omitempty"`
	Budget      *float64          `json:"budget,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	ExpiresAt   *time.Time        `json:"expires_at,omitempty"`
}

type KeyResponse struct {
	ID          string            `json:"id"`
	Key         string            `json:"key"`
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Models      []string          `json:"models,omitempty"`
	MaxTokens   *int64            `json:"max_tokens,omitempty"`
	MaxRequests *int64            `json:"max_requests,omitempty"`
	Budget      *float64          `json:"budget,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	ExpiresAt   *time.Time        `json:"expires_at,omitempty"`
	IsActive    bool              `json:"is_active"`
	UsageStats  *UsageStats       `json:"usage_stats,omitempty"`
}

type Manager interface {
	VirtualKeyManager
	GetKeyByValue(ctx context.Context, keyValue string) (*VirtualKey, error)
	ValidateKey(ctx context.Context, keyValue string, modelName string) (*VirtualKey, error)
	UpdateUsage(ctx context.Context, keyValue string, tokens int64, cost float64) error
}

func GenerateAPIKey() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}
	return "ogem-" + hex.EncodeToString(bytes), nil
}

func GenerateKeyID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}
	return "key_" + hex.EncodeToString(bytes), nil
}

func VirtualKeyToKeyResponse(key *VirtualKey) *KeyResponse {
	return &KeyResponse{
		ID:          key.ID,
		Key:         key.Key,
		Name:        key.Name,
		Description: key.Description,
		Models:      key.Models,
		MaxTokens:   key.MaxTokens,
		MaxRequests: key.MaxRequests,
		Budget:      key.Budget,
		Metadata:    key.Metadata,
		CreatedAt:   key.CreatedAt,
		ExpiresAt:   key.ExpiresAt,
		IsActive:    key.IsActive,
		UsageStats:  key.UsageStats,
	}
}

// VirtualKeysToKeyResponses converts a slice of VirtualKey to slice of KeyResponse
func VirtualKeysToKeyResponses(keys []*VirtualKey) []*KeyResponse {
	responses := make([]*KeyResponse, len(keys))
	for i, key := range keys {
		responses[i] = VirtualKeyToKeyResponse(key)
	}
	return responses
}
