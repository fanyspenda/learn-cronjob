package main

import (
	"time"

	cron "github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func main() {

	logging := logrus.New()
	jakartaTime, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		logging.Errorf("Error setting time - %s", err.Error())
	}

	timeCurrent := time.Now().In(jakartaTime).Format("2006-01-02 15:04:05")

	logging.Infof("Jakarta Time: %s", timeCurrent)

	customParser := cron.WithParser(cron.NewParser(
		cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow,
	))

	scheduler := cron.New(cron.WithLocation(jakartaTime), customParser)

	_, err = scheduler.AddFunc("*/5 * * * * *", func() {
		timeCurrent := time.Now().In(jakartaTime).Format("2006-01-02 15:04:05")
		logging.Infof("sending message every 5 seconds. Current Time: %s", timeCurrent)
	})

	if err != nil {
		logging.Errorf("Error creating task for cronjob - %s", err.Error())
	}

	go scheduler.Start()

	sig := make(chan bool, 1)
	<-sig
}
