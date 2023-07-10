package transformer

import (
	"bytes"
	"fmt"

	"github.com/haohaiwei/woa/model"
)

// TransformToMarkdown transform alertmanager notification to dingtalk markdow message
func TransformToMarkdown(notification model.Notification, cluster string) (markdown *model.WoaMarkdown, robotURL string, err error) {

	status := notification.Status
	alertname := notification.GroupLabels["alertname"]

	annotations := notification.CommonAnnotations
	robotURL = annotations["woaRobot"]

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("#### 告警集群: %s\n", cluster))
	buffer.WriteString(fmt.Sprintf("##### 告警项: %s\n", alertname))
	buffer.WriteString(fmt.Sprintf("##### 当前状态: %s\n", status))

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
