package account_usecase

import (
	rand "math/rand/v2"
)

var handleLetters = []rune("0123456789")

func genHandle() string {
	b := make([]rune, 12) // 16 (username max length) - len("user")
	for i := range b {
		b[i] = handleLetters[rand.IntN(len(handleLetters))]
	}

	return "user" + string(b)
}

var otpLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func genOtpCode() string {
	b := make([]rune, 18) // secure length: https://www.reddit.com/media?url=https%3A%2F%2Fi.redd.it%2F5g3ayy7pwxl51.jpg
	for i := range b {
		b[i] = otpLetters[rand.IntN(len(handleLetters))]
	}

	return string(b)
}
