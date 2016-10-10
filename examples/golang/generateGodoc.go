package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	eventMap map[string]EventInfo

	RUN_ID string
)

type EventInfo struct {
	id  string
	url string
}

func init() {
	eventMap = make(map[string]EventInfo)
}

func main() {
	initEvent()

	// when start a component ,send a COMPONENT_START notify and register self in pipeline
	notifyEvent("COMPONENT_START", "application/json", strings.NewReader("component start..."))

	// on this func ,you can wait pipeline send the data you will use
	// waitForData()

	// after get all data you need, send a TASH_START notify to pipeline and start do what you need
	notifyEvent("TASK_START", "application/json", strings.NewReader("task start..."))

	result := generateGodoc()

	resultMap := make(map[string]interface{}, 0)
	resultMap["output"] = map[string]string{"doc": result}
	resultMap["status"] = "success"
	resultMap["result"] = true
	resultByte, _ := json.Marshal(resultMap)
	// after task done send a TASK_RESULT notify to pipeline
	notifyEvent("TASK_RESULT", "application/json", bytes.NewReader(resultByte))

	// on this func, you can do some after-task job,like recycle some resource
	// recycleResource()

	// after after-task job done, send a COMPONENT_STOP notify to pipeline ,so that pipeline will stop it as soon as possible
	// if you don't send a COMPONENT_STOP notify,pipelin will auto stop it after some time(default:60s)
	notifyEvent("COMPONENT_STOP", "application/json", strings.NewReader("component stoping..."))

	fmt.Println("done...")
	// wait the pipeline stop service,don't exit by yourself
	waitForKill()
}

func initEvent() {
	eventList := os.Getenv("EVENT_LIST")
	RUN_ID = os.Getenv("RUN_ID")

	for _, eventInfo := range strings.Split(eventList, ";") {
		if len(strings.Split(eventInfo, ",")) > 1 {
			eventKey := strings.Split(eventInfo, ",")[0]
			eventId := strings.Split(eventInfo, ",")[1]
			if os.Getenv(eventKey) != "" {
				event := new(EventInfo)
				event.id = eventId
				event.url = os.Getenv(eventKey)
				eventMap[eventKey] = *event
			}
		}
	}
}

func notifyEvent(eventKey, bodyType string, body io.Reader) {
	if event, ok := eventMap[eventKey]; ok {
		eventId := event.id
		notifyUrl := event.url
		if strings.Contains(notifyUrl, "?") {
			notifyUrl += "&runId=" + RUN_ID
		} else {
			notifyUrl += "?runId=" + RUN_ID
		}
		notifyUrl += "&event=" + eventKey + "&eventId=" + eventId

		client := &http.Client{}
		req, _ := http.NewRequest(http.MethodPut, notifyUrl, body)
		// resp, err := http.Post(notifyUrl, bodyType, body)
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("error when notify event", eventKey, body)
		}

		respBody, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println("notify resp is", string(respBody))
	}
}

func generateGodoc() string {
	cmd := exec.Command("godoc", "cmd/github.com/kubernetes/kubernetes/cmd/kubectl/app")
	var result bytes.Buffer
	cmd.Stdout = &result
	err := cmd.Run()
	if err != nil {
		return err.Error()
	}
	return result.String()
}

func waitForKill() {
	for {
		time.Sleep(1 * time.Second)
	}
}
