package Sender

import (
  SmsMessage "<%= appName %>/pkg/service/Sms/Message"

  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
)

type vonageSmsRequestBody struct {
  From      string `json:"from"`
  Text      string `json:"text"`
  To        string `json:"to"`
  APIKey    string `json:"api_key"`
  APISecret string `json:"api_secret"`
}

type vonageSender struct {
  apiKey string
  apiSecret string
}

func NewVonage(apiKey, apiSecret string) *vonageSender {
  return &vonageSender{apiKey: apiKey, apiSecret: apiSecret}
}

func (s *vonageSender) Send(message SmsMessage.Message) {
  body := SMSRequestBody{
    APIKey:    s.apiKey,
    APISecret: s.apiSecret,
    To:        message.GetTo(),
    From:      message.GetFrom(),
    Text:      message.GetText(),
  }

  smsBody, err := json.Marshal(body)
  if err != nil {
    panic(err)
  }

  resp, err := http.Post("https://rest.nexmo.com/sms/json", "application/json", bytes.NewBuffer(smsBody))
  if err != nil {
    panic(err)
  }

  defer resp.Body.Close()

  respBody, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }

  fmt.Println(string(respBody))
}
