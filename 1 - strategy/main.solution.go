package main

import "fmt"

type EmailNotification struct{}

func (e EmailNotification) Send(message string) {
	fmt.Println("send email::: ", message)
}

type NotificationStratery interface {
	Send(message string)
}

type NotificationService struct {
	strategy NotificationStratery
}

func NewNotificationService(strategy NotificationStratery) *NotificationService {
	return &NotificationService{
		strategy: strategy,
	}
}

func (s *NotificationService) Send(message string) {
	s.strategy.Send(message)
}

func main() {
	emailService := NewNotificationService(EmailNotification{})
	emailService.Send("Welcom you")
}
