package stream

import (
	"time"

	"github.com/robfig/cron/v3"
)

func (s *appImpl) Cron() {
	c := cron.New(cron.WithLocation(time.UTC))
	c.AddFunc("*/5 * * * *", s.Status)
	c.Start()
}
