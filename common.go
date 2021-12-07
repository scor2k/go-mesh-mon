package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func enableJsonCors(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	enableJsonCors(&w)

	msg, _ := json.Marshal(map[string]interface{}{"result": "ok", "msg": fmt.Sprintf("go-mesh-mon v.%v", appVersion)})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(msg)
}
