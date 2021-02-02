package Mailer

import (
  MailMessage "<%= appName %>/pkg/service/Mail/Message"

  "log"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ses"
  "github.com/aws/aws-sdk-go/aws/awserr"
)

type sesMailer struct {
  username string
  sess *session.Session
}

func NewSes(username, region string) *sesMailer {
  sess, _ := session.NewSession(&aws.Config{
    Region: aws.String(region)},
  )
  return &sesMailer{sess: sess}
}

func (m *sesMailer) Send(message MailMessage.Message) {
  // Create an SES session.
  svc := ses.New(m.sess)
  // Assemble the email.
  input := &ses.SendEmailInput{
    Destination: &ses.Destination{
      CcAddresses: []*string{
      },
      ToAddresses: []*string{
        aws.String(message.GetTo()[0]),
      },
    },
    Message: &ses.Message{
      Body: &ses.Body{
        // Html: &ses.Content{
        // 		Charset: aws.String(CharSet),
        // 		Data:    aws.String(HtmlBody),
        // },
        Text: &ses.Content{
            Charset: aws.String("UTF-8"),
            Data:    aws.String(message.GetBody()),
        },
      },
      Subject: &ses.Content{
        Charset: aws.String("UTF-8"),
        Data:    aws.String(message.GetSubject()),
      },
    },
    Source: aws.String(m.username),
    // Uncomment to use a configuration set
    //ConfigurationSetName: aws.String(ConfigurationSet),
  }
  // Attempt to send the email.
  result, err := svc.SendEmail(input)
  // Display error messages if they occur.
  if err != nil {
    if aerr, ok := err.(awserr.Error); ok {
      switch aerr.Code() {
      case ses.ErrCodeMessageRejected:
        log.Println(ses.ErrCodeMessageRejected, aerr.Error())
      case ses.ErrCodeMailFromDomainNotVerifiedException:
        log.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
      case ses.ErrCodeConfigurationSetDoesNotExistException:
        log.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
      default:
        log.Println(aerr.Error())
      }
    } else {
      // Print the error, cast err to awserr.Error to get the Code and
      // Message from an error.
      log.Println(err.Error())
    }
    return
  }

  log.Println("Email Sent to address: " + message.GetTo()[0])
  log.Println(result)
}
