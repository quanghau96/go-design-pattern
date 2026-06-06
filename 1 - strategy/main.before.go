// package main

// import "fmt"

// type NotificationService struct{}

// func (s NotificationService) Send(channel string, message string) {
// 	if channel == "email" {
// 		fmt.Println("Send email:", message)
// 	} else if channel == "telegram" {
// 		fmt.Println("Send telegram:", message)
// 	} else if channel == "discord" {
// 		fmt.Println("Send discord:", message)
// 	} else {
// 		fmt.Println("Unsupported channel")
// 	}
// }

// func main() {
// 	service := NotificationService{}

// 	service.Send("email", "Welcome you")
// 	service.Send("telegram", "New order created")
// }

// problem
// - NotificationService knows too many logic
// - when sending a slack, we have to modify the code
// - want to test a single channel, it's too difficult
// - Violate the rule Open/Closed Principle, Code should be open for expansion, but limit editing of old code.
