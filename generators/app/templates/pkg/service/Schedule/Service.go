package Schedule

import (
  "log"
)

type Service interface {
  // Trigger
  NewTrigger(name string) Trigger

  // Actions
  // Command(cmd string) Job
  // Job(jobName string) Job
  // Exec(cmd string) Job
  NewJob(name string) Job

  // Run
  Schedule(job Job)
  // CreateJob(group string, job Job)
  // UpdateJob(group string, name string, job Job)
  // DeleteJob(group string, name string)
}

type service struct {}

func NewService() Service {
  return &service{}
}

func (s *service) NewTrigger(name string) Trigger {
  return NewTrigger(name)
}

func (s *service) NewJob(name string) Job {
  return NewJobHttp("http", name)
}

func (s *service) Schedule(job Job) {
  if (job.GetTrigger().IsInvalid()) {
    log.Println("Invalid trigger")
    return
  }
  if (job.GetTrigger().IsCron()) {
    log.Println("Schedule cron job")
    job.Run()
    return
  }
  log.Println("Schedule firetime job")
  job.Run()
}
