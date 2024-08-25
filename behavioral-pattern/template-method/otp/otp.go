package otp

type IOtp interface {
	GenRandomOTP(int) string
	SaveOTP(string)
	GetMessage(string) string
	SendNotification(string) error
}

type Otp struct {
	IOtp
}

func (o *Otp) GenAndSendOTP(otpLength int) error {
	otp := o.GenRandomOTP(otpLength)
	o.SaveOTP(otp)
	message := o.GetMessage(otp)
	return o.SendNotification(message)
}
