package Schedule

import (
  "log"
  "strconv"
)

type JobHttp interface {
  Trigger(trigger Trigger) Job
  GetTrigger() Trigger

  Url(url string) JobHttp
  Body(body string) JobHttp
  Method(method string) JobHttp

  CanRun() bool

  Run()
  Retry() bool
}

type jobHttp struct {
  action string
  name string
  trigger Trigger

  url string
  body string
  method string
  limitRetries int
  numRetries int
}

func NewJobHttp(action string, name string) JobHttp {
  return &jobHttp{action: action, name: name, numRetries: 0, limitRetries: 3}
}

func (j *jobHttp) Trigger(trigger Trigger) Job {
  j.trigger = trigger
  return j
}

func (j *jobHttp) GetTrigger() Trigger {
  return j.trigger
}

func (j *jobHttp) Url(url string) JobHttp {
  j.url = url
  return j
}

func (j *jobHttp) Body(body string) JobHttp {
  j.body = body
  return j
}

func (j *jobHttp) Method(method string) JobHttp {
  j.method = method
  return j
}

func (j *jobHttp) CanRun() bool {
  return j.numRetries < j.limitRetries
}

func (j *jobHttp) Run() {
  if !j.CanRun() {
    log.Println("Run -> failed limit retries -> name:" + j.name+ " numRetries:" + strconv.Itoa(j.numRetries))
    return
  }
  log.Println("Run finish job : name:" + j.name+ " action:"+j.action)
}

func (j *jobHttp) Retry() bool {
  j.numRetries++
  j.Run()
  return true
}
