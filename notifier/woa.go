package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/haohaiwei/woa/model"
	"github.com/haohaiwei/woa/transformer"
)

// Send send markdown message to woa
func Send(notification model.Notification, defaultRobot string) (err error) {

	markdown, robotURL, err := transformer.TransformToMarkdown(notification)

	if err != nil {
		return
	}

	data, err := json.Marshal(markdown)
	if err != nil {
		return
	}

	var woaRobotURL string

	if robotURL != "" {
		woaRobotURL = robotURL
	} else {
		woaRobotURL = defaultRobot
	}

	if len(woaRobotURL) == 0 {
		return nil
	}

	req, err := http.NewRequest(
		"POST",
		woaRobotURL,
		bytes.NewBuffer(data))

	if err != nil {
		fmt.Println("woa robot url not found ignore:")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	return
}
