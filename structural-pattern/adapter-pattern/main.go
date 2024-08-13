package main

import (
	"log"

	"github.com/sirupsen/logrus"
)

func main() {
	var logger Logger
	// Old logger
	logger = &OldLogger{}
	logger.LogInfo("Hello World")
	logger.LogError("Hello World")

	// New logger
	logger = NewLogrusAdapter()
	logger.LogInfo("logrus: Hello World")
	logger.LogError("logrus: Hello World")
}

// target
type Logger interface {
	LogInfo(message string)
	LogError(message string)
}

// OldLogger is the old logger
type OldLogger struct{}

func (l *OldLogger) LogInfo(message string) {
	log.Println(message)
}

func (l *OldLogger) LogError(message string) {
	log.Println(message)
}

// NewLogger is the new logger
// adaptee
type LogrusLogger struct {
	logrus *logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{
		logrus: logrus.New(),
	}
}

func (l *LogrusLogger) LogInfo(message string) {
	l.logrus.Info(message)
}

func (l *LogrusLogger) LogError(message string) {
	l.logrus.Error(message)
}

// Adapter
type LogrusAdapter struct {
	logrusLogger *LogrusLogger
}

func NewLogrusAdapter() *LogrusAdapter {
	return &LogrusAdapter{
		logrusLogger: NewLogrusLogger(),
	}
}

func (l *LogrusAdapter) LogInfo(message string) {
	l.logrusLogger.LogInfo(message)
}

func (l *LogrusAdapter) LogError(message string) {
	l.logrusLogger.LogError(message)
}

// type paypal struct {
// }

// func NewPaypal() *paypal {
// 	return &paypal{}
// }

// func (p *paypal) MakePayment(amount int) {
// 	fmt.Println("Paypal payment is done for amount: ", amount)
// }

// type amazon struct {
// }

// func NewAmazon() *amazon {
// 	return &amazon{}
// }

// func (a *amazon) PayAmazon(amount int) {
// 	fmt.Println("Amazon payment is done for amount: ", amount)
// }

// type paymentGateway interface {
// 	processPayment(amount int)
// }

// type paypalAdapter struct {
// 	paypalGateway *paypal
// }

// func (p *paypalAdapter) processPayment(amount int) {
// 	p.paypalGateway.MakePayment(amount)
// }

// type amazonAdapter struct {
// 	amazonGateway *amazon
// }

// func (a *amazonAdapter) processPayment(amount int) {
// 	a.amazonGateway.PayAmazon(amount)
// }

// func main() {
// 	paypal := NewPaypal()
// 	paypalAdapter := &paypalAdapter{
// 		paypalGateway: paypal,
// 	}

// 	amazon := NewAmazon()
// 	amazonAdapter := &amazonAdapter{
// 		amazonGateway: amazon,
// 	}

// 	paypalAdapter.processPayment(100)

// 	amazonAdapter.processPayment(200)
// }
