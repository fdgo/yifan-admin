package cron

import (
	"github.com/robfig/cron"
)

func (s *cronServiceImpl) Start() {
	release := NewRecycleService(s.db, s.cache)
	s.AddJob("* * * * *", release)
	s.cron.Start()
}

func (s *cronServiceImpl) AddJob(spec string, job cron.Job) {
	s.cron.AddJob(spec, job)
}

func (s *cronServiceImpl) Stop() {
	s.cron.Stop()
}
