package dingtalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"kiwi/internal/app/project/gitlab"
)

// Robot 钉钉机器人
type Robot struct {
	WebhookURL string
	Secret     string
	Keywords   []string
}

// NewRobot 创建钉钉机器人实例
func NewRobot(webhookURL, secret string, keywords []string) *Robot {
	return &Robot{
		WebhookURL: webhookURL,
		Secret:     secret,
		Keywords:   keywords,
	}
}

// Message 钉钉消息结构
type Message struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
}

// sign 生成签名
func (r *Robot) sign() (string, int64) {
	timestamp := time.Now().UnixMilli()
	strToSign := fmt.Sprintf("%d\n%s", timestamp, r.Secret)

	mac := hmac.New(sha256.New, []byte(r.Secret))
	mac.Write([]byte(strToSign))
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return sign, timestamp
}

// SendPushEvent 发送代码推送事件消息
func (r *Robot) SendPushEvent(event *gitlab.PushEvent) error {
	// 构建消息内容
	msg := Message{
		Msgtype: "markdown",
	}

	// 确保消息包含关键词
	keyword := r.Keyword()
	msg.Markdown.Title = fmt.Sprintf("%s: 【%s】新的代码提交", r.Secret, keyword)

	// 格式化消息内容
	msg.Markdown.Text = fmt.Sprintf("%s\n\n"+
		"**项目**: [%s](%s)\n\n"+
		"**提交者**: %s\n\n"+
		"**分支**: %s\n\n"+
		"**提交信息**:\n%s",
		keyword,
		event.Project.Name,
		event.Project.WebURL,
		event.UserName,
		event.GetBranchName(),
		event.FormatCommitMessages(),
	)

	slog.Info("准备发送钉钉消息",
		"project", event.Project.Name,
		"branch", event.GetBranchName(),
		"user", event.UserName)

	// 生成签名
	sign, timestamp := r.sign()

	// 构建请求URL
	webhookURL, _ := url.Parse(r.WebhookURL)
	q := webhookURL.Query()
	q.Set("timestamp", fmt.Sprintf("%d", timestamp))
	q.Set("sign", sign)
	webhookURL.RawQuery = q.Encode()

	// 发送请求
	payload, err := json.Marshal(msg)
	if err != nil {
		slog.Error("序列化消息失败", "error", err)
		return fmt.Errorf("序列化消息失败: %v", err)
	}

	resp, err := http.Post(webhookURL.String(), "application/json", bytes.NewBuffer(payload))
	if err != nil {
		slog.Error("发送消息失败", "error", err)
		return fmt.Errorf("发送消息失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("发送消息失败", "status_code", resp.StatusCode)
		return fmt.Errorf("发送消息失败，状态码: %d", resp.StatusCode)
	}

	slog.Info("成功发送钉钉消息",
		slog.String("project", event.Project.Name),
		slog.String("branch", event.GetBranchName()),
	)

	return nil
}

// Keyword 获取第一个关键词
func (r *Robot) Keyword() string {
	if len(r.Keywords) > 0 {
		return r.Keywords[0]
	}
	return "代码提交"
}
