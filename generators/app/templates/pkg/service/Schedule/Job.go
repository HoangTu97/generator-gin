package Schedule

type Job interface {
  Trigger(trigger Trigger) Job
  GetTrigger() Trigger
  Run()
}
