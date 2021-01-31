package Notification

type Service interface {
  // Config noti
  From(from string)
  To(to string)
  Via(channel string)

  // Send
  Delay(delay int)
  Send(users []string)
  SendNow(users []string)
}