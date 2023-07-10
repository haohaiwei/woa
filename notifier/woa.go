package notifier

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/haohaiwei/woa/model"
	"github.com/haohaiwei/woa/transformer"
)

// Send send markdown message to woa
func Send(notification model.Notification, defaultRobot string, cluster string) (err error) {

	markdown, robotURL, err := transformer.TransformToMarkdown(notification, cluster)

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
		log.Println("woa robot url not found ignore:")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		log.Println(data)
		return
	}

	defer resp.Body.Close()
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	return
}
