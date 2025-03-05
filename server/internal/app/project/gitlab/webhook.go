package gitlab

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

// PushEvent GitLab push事件的数据结构
type PushEvent struct {
	ObjectKind string `json:"object_kind"`
	Before     string `json:"before"`
	After      string `json:"after"`
	Ref        string `json:"ref"`
	UserName   string `json:"user_name"`
	Project    struct {
		Name            string `json:"name"`
		Description     string `json:"description"`
		WebURL          string `json:"web_url"`
		VisibilityLevel int    `json:"visibility_level"`
		Namespace       string `json:"namespace"`
	} `json:"project"`
	Commits []struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
}

// WebhookHandler 处理GitLab webhook请求
type WebhookHandler struct {
	SecretToken string
}

// NewWebhookHandler 创建webhook处理器
func NewWebhookHandler(secretToken string) *WebhookHandler {
	return &WebhookHandler{SecretToken: secretToken}
}

// ValidateSignature 验证请求签名
func (h *WebhookHandler) ValidateSignature(signature string, body []byte) bool {
	if h.SecretToken == "" {
		return true
	}

	expectedSignature := computeHMAC(body, []byte(h.SecretToken))
	return signature == expectedSignature
}

// computeHMAC 计算HMAC签名
func computeHMAC(message, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return "sha256=" + hex.EncodeToString(mac.Sum(nil))
}

// HandlePushEvent 处理push事件
func (h *WebhookHandler) HandlePushEvent(w http.ResponseWriter, r *http.Request) (*PushEvent, error) {
	if r.Method != http.MethodPost {
		slog.Error("不支持的HTTP方法", "method", r.Method)
		return nil, fmt.Errorf("不支持的HTTP方法: %s", r.Method)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("读取请求体失败", "error", err)
		return nil, fmt.Errorf("读取请求体失败: %v", err)
	}
	defer r.Body.Close()

	// 验证签名
	signature := r.Header.Get("X-Gitlab-Token")
	// 只与本地 secret_token 做比较
	if h.SecretToken != "" && signature != h.SecretToken {
		slog.Error("签名验证失败", "signature", signature)
		return nil, fmt.Errorf("签名验证失败")
	}

	// 解析事件数据
	var event PushEvent
	if err := json.Unmarshal(body, &event); err != nil {
		slog.Error("解析JSON失败", "error", err)
		return nil, fmt.Errorf("解析JSON失败: %v", err)
	}

	slog.Info("成功处理Push事件",
		slog.String("project", event.Project.Name),
		slog.String("branch", event.GetBranchName()),
		slog.String("user", event.UserName),
		slog.Int("commits_count", event.TotalCommitsCount),
	)

	return &event, nil
}

// GetBranchName 从ref中提取分支名
func (e *PushEvent) GetBranchName() string {
	return strings.TrimPrefix(e.Ref, "refs/heads/")
}

// FormatCommitMessages 格式化提交信息
func (e *PushEvent) FormatCommitMessages() string {
	var messages []string
	for _, commit := range e.Commits {
		messages = append(messages, fmt.Sprintf("- %s\n  %s",
			commit.Message,
			commit.URL))
	}
	return strings.Join(messages, "\n")
}
