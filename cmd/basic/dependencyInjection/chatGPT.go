package main

import "fmt"

// 定义抽象接口
type MessageSender interface {
	SendMessage(string)
}

// 定义具体实现
type EmailSender struct{}

func (es *EmailSender) SendMessage(msg string) {
	fmt.Println("Sending email:", msg)
}

// 定义具体实现
type SmsSender struct{}

func (ss *SmsSender) SendMessage(msg string) {
	fmt.Println("Sending SMS:", msg)
}

// 使用依赖注入
type UserService struct {
	sender MessageSender
}

func (us *UserService) SetMessageSender(sender MessageSender) {
	us.sender = sender
}

func (us *UserService) SendWelcomeMessage(name string) {
	us.sender.SendMessage("Welcome, " + name)
}

func main() {
	// 使用EmailSender作为依赖对象
	emailSender := &EmailSender{}
	userService := &UserService{sender: emailSender}
	userService.SendWelcomeMessage("John")

	// 使用SmsSender作为依赖对象
	smsSender := &SmsSender{}
	userService.SetMessageSender(smsSender)
	userService.SendWelcomeMessage("Jane")
}
