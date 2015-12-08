package main

import "net"
import "time"
import "net/http"
import "fmt"

import "nowprovision/webhookproxy"

func main() {

	fmt.Println("Starting webhookproxy server on port 8080")

	config := webhookproxy.Config{}
	config.TryLaterStatusCode = 503
	config.BackQueueSize = 100
	config.MaxWaitSeconds = 10 * time.Second
	config.UseLongPoll = true
	config.MaxPayloadSize = 5000000
	config.Secret = "whatever"
	config.Hostname = "" // leave blank for not multi-tenant

	_, localNetwork, _ := net.ParseCIDR("127.0.0.1/24")
	config.WebhookWhiteList = []*net.IPNet{localNetwork}
	config.PollReplyWhiteList = []*net.IPNet{localNetwork}

	webhookproxy.RegisterHandlers(&config, http.DefaultServeMux)
	http.ListenAndServe(":8080", nil)
}
