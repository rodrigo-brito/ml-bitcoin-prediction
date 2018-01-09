package main

import (
	"sync"

	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Crawler started...")
	s := gocron.NewScheduler()
	s.Every(15).Minutes().Do(func() {
		logrus.Infof("Stating cron at %s", time.Now())
		wg := new(sync.WaitGroup)
		wg.Add(2)
		go func() {
			if err := RunCripto(); err != nil {
				logrus.Error(err)
			}
			wg.Done()
		}()

		go func() {
			if err := RunMarket(); err != nil {
				logrus.Error(err)
			}
			wg.Done()
		}()
		wg.Wait()
	})
	<-s.Start()
}
