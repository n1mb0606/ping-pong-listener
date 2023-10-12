package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/n1mb0606/ping-pong-listener/influxdbclient"
)

var org, bucket, token, url string

func listenRoutine() {
	influxClient := influxdbclient.GetNewInfluxClient(org, bucket, url, token)

	st := time.Now()
	client := http.Client {
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(os.Getenv("PING_PONG_TARGET"))

	var tags map[string]string
	var fields map[string]interface{}
	if err != nil {
		fmt.Println(err)
		tags = map[string]string{
			"type": "stat",
		}
		fields = map[string]interface{}{
			"up": false,
			"status": 0,
			"ping": -1.0,
		}
	} else {
		tags = map[string]string{
			"type": "stat",
		}
		fields = map[string]interface{}{
			"up": true,
			"status": resp.StatusCode,
			"ping": time.Since(st).Seconds(),
		}
	}
	influxClient.Write("wlog", tags, fields)
}

func main() {
	org = os.Getenv("PING_PONG_ORG")
	bucket = os.Getenv("PING_PONG_BUCKET")
	token = os.Getenv("PING_PONG_TOKEN")
	url = os.Getenv("PING_PONG_URL")

	for true {
		go listenRoutine()
		time.Sleep(1 * time.Second)
	}
}
