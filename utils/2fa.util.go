package utils

import (
	"github.com/mdp/qrterminal"
	"github.com/skip2/go-qrcode"
	"github.com/xlzd/gotp"
	"log"
	"time"
)

// GenerateSecret is a function that generates a secret key for OTP.
func GenerateSecret() string {
	return gotp.RandomSecret(16)
}

// GenerateTOTPWithSecret is a function that generates a TOTP URI with the secret key.
func GenerateTOTPWithSecret(secret string) string {
	uri := gotp.NewDefaultTOTP(secret).ProvisioningUri("swe@yopmail.com", "SweApp")
	log.Printf("[2FA] TOTP URI: %s", uri)

	err := qrcode.WriteFile(uri, qrcode.Medium, 256, "qr.png")
	if err != nil {
		log.Printf("[2FA] Error generating QR code: %s", err.Error())
		return ""
	}

	qrterminal.GenerateWithConfig(uri, qrterminal.Config{
		Level:     qrterminal.L,
		Writer:    log.Writer(),
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
	})

	return uri
}

// VerifyOTP is a function that verifies the OTP code with the secret key.
func VerifyOTP(secret string, otp string) bool {
	totp := gotp.NewDefaultTOTP(secret)
	if totp.Verify(otp, time.Now().Unix()) {
		log.Printf("[2FA] OTP code [%s] is valid. Access granted.", otp)
		return true
	}

	return false
}
