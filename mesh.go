package main

import (
	"encoding/json"
	"net/http"
)

func PingCheck(w http.ResponseWriter, req *http.Request) {
	enableJsonCors(&w)

	ua := req.Header.Get("User-Agent")
	ips := req.Header.Get("X-Forwarded-For")
	ip := req.Header.Get("X-Client-IP")

	if ips == "" {
		ips = ip
	}

	msg, _ := json.Marshal(map[string]interface{}{"result": "ok"})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(msg)

	log.Info("Remote ping. UA: ", ua, " IPs: ", ips)
}
