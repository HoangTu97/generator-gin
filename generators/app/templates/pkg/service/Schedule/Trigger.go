package Schedule

import (
  "time"
)

type Trigger interface {
  Cron(cron string) Trigger
  At(firetime int64) Trigger
  Next(firetime int64) Trigger
  // everyMinute()
  // everyTwoMinutes()
  // everyThreeMinutes()
  // everyFourMinutes()
  // everyFiveMinutes()
  // everyTenMinutes()
  // everyFifteenMinutes()
  // everyThirtyMinutes()
  // hourly()
  // hourlyAt(firetime int)
  // everyTwoHours()
  // everyThreeHours()
  // everyFourHours()
  // everySixHours()
  // daily()
  // dailyAt(timeFormat string)
  // twiceDaily(t1 int, t2 int)
  // weekly()
  // weeklyOn(day int, timeFormat string)
  // monthly()
  // monthlyOn(day int, timeFormat string)
  // twiceMonthly(d1 int, d2 int, timeFormat string)
  // lastDayOfMonth(timeFormat string)
  // quarterly()
  // yearly()
  // yearlyOn(day int, month int, timeFormat string)
  // timezone(zoneId string)
  IsInvalid() bool
  IsCron() bool
}

type trigger struct {
  name string
  firetime int64
  cron string
}

func NewTrigger(name string) Trigger {
  return &trigger{name: name}
}

func (t *trigger) Cron(cron string) Trigger {
  t.cron = cron
  return t
}

func (t *trigger) At(firetime int64) Trigger {
  t.firetime = firetime
  return t
}

func (t *trigger) Next(firetime int64) Trigger {
  t.firetime = time.Now().Unix() + firetime
  return t
}

func (t *trigger) IsInvalid() bool {
  return t.cron == "" && t.firetime == 0
}

func (t *trigger) IsCron() bool {
  return t.cron != ""
}
