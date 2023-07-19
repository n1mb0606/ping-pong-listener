package main

import (
	"net/http"
	"os"
	"time"

	"github.com/n1mb0606/ping-pong-listener/influxdbclient"
)

func main() {
	org := os.Getenv("PING_PONG_ORG")
	bucket := os.Getenv("PING_PONG_BUCKET")
	token := os.Getenv("PING_PONG_TOKEN")
	url := os.Getenv("PING_PONG_URL")

	influxClient := influxdbclient.GetNewInfluxClient(org, bucket, url, token)

	for true {
		resp, _ := http.Get(os.Getenv("PING_PONG_TARGET"))

		tags := map[string]string{
			"type": "status",
		}
		fields := map[string]interface{}{
			"status": resp.StatusCode,
		}

		influxClient.Write("wlog", tags, fields)
		time.Sleep(time.Second * 1)
	}
}
