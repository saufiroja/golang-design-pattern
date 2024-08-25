package otp

import "fmt"

type Email struct {
	Otp
}

func (e *Email) GenRandomOTP(otpLength int) string {
	randomOtp := "1234"
	fmt.Printf("Email: generating random otp %s\n", randomOtp)
	return randomOtp
}

func (e *Email) SaveOTP(otp string) {
	fmt.Printf("Email: saving otp: %s to the file\n", otp)
}

func (e *Email) GetMessage(otp string) string {
	return "Email OTP for login is " + otp
}

func (e *Email) SendNotification(message string) error {
	fmt.Printf("Email: sending email: %s\n", message)
	return nil
}
