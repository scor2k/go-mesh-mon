package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// JsonResponse -> map for parse json
type JsonResponse map[string]interface{}

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

// GetSettingValues -> ask Settings for values for the specific key
func GetSettingValues(key string) []string {
	var response []string
	var result []Settings
	DB.Model(&Settings{}).Where("Key = ?", key).Find(&result)

	for _, item := range result {
		response = append(response, item.Value)
	}

	return response
}

// GetRequest return unstructured json
func GetRequest(url string) (response JsonResponse, err error) {
	ua := fmt.Sprintf("go-mesh-mon/%s", appVersion)
	timeout, _ := strconv.Atoi(CheckPingTimeout)
	httpClient := http.Client{Timeout: time.Second * time.Duration(timeout)}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Error("getRequest error: ", err)
		return nil, err
	}

	req.Header.Set("User-Agent", ua)

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Error("getRequest getError: ", getErr)
		return nil, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Error("getRequest readError: ", readErr)
		return nil, readErr
	}

	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Error("getRequest parseError: ", jsonErr)
		return nil, jsonErr
	}

	return response, nil
}

func AddMetric(name string, instance string, value float64) {
	record := Metrics{Name: name, Instance: instance, Value: value, Timestamp: time.Now().UnixNano() / int64(time.Millisecond)}
	res := DB.Create(&record)
	if res.Error != nil {
		log.Error("Cannot add new DB record")
	}
}
