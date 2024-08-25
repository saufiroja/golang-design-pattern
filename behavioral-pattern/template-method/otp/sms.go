package otp

import "fmt"

type Sms struct {
	Otp
}

func (s *Sms) GenRandomOTP(otpLength int) string {
	randomOtp := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOtp)
	return randomOtp
}

func (s *Sms) SaveOTP(otp string) {
	fmt.Printf("SMS: saving otp: %s to the file\n", otp)
}

func (s *Sms) GetMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) SendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}
