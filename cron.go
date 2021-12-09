package main

import (
	"fmt"
	"github.com/robfig/cron"
	"sync"
	"time"
)

func StartCron() {
	c := cron.New()
	c.AddFunc("*/5 * * * * *", DoPing)
	c.Start()
}

// DoPing -> send GET requests to every server from the list
func DoPing() {
	hosts := GetSettingValues("check_ping")
	if len(hosts) == 0 {
		log.Info("No server for ping_check available")
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(hosts))

	for _, host := range hosts {

		go func(host string) {
			defer wg.Done()

			url := fmt.Sprintf("http://%s:1982/api/v1/ping", host)
			t1 := time.Now()
			_, err := GetRequest(url)
			t2 := time.Now()
			if err != nil {
				log.Error("Cannot ping ", host)
				AddMetric("check_ping_ms", host, float64(-1))
				return
			}
			diff := t2.Sub(t1)

			AddMetric("check_ping_ms", host, float64(diff)/float64(time.Millisecond))
			log.Info("Ping ", host, " is okay and took ", diff)
		}(host)

	}
}
