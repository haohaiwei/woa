package transformer

import (
	"bytes"
	"fmt"

	"github.com/haohaiwei/woa/model"
)

// TransformToMarkdown transform alertmanager notification to dingtalk markdow message
func TransformToMarkdown(notification model.Notification) (markdown *model.WoaMarkdown, robotURL string, err error) {

	groupKey := notification.GroupKey
	status := notification.Status

	annotations := notification.CommonAnnotations
	robotURL = annotations["woaRobot"]

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("### 通知组%s(当前状态:%s) \n", groupKey, status))

	buffer.WriteString(fmt.Sprintf("#### 告警项:\n"))

	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("##### %s\n > %s\n", annotations["summary"], annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n> 开始时间：%s\n", alert.StartsAt.Format("15:04:05")))
	}

	markdown = &model.WoaMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Text: buffer.String(),
		},
	}

	return
}
