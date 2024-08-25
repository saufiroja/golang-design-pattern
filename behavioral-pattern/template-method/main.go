package main

import "fundamental/golang-design-pattern/behavioral-pattern/template-method/otp"

func main() {
	smsOtp := &otp.Sms{}
	o := otp.Otp{
		IOtp: smsOtp,
	}

	_ = o.GenAndSendOTP(4)

	emailOtp := &otp.Email{}

	o = otp.Otp{
		IOtp: emailOtp,
	}

	_ = o.GenAndSendOTP(4)
}
