package helper

import (
	"crypto/rand"
	"math"
	"math/big"
	"strconv"
	"time"
)

func GeneratePoId(numberOfDigits int) (string, error) {
	prefix := "POID"
	uniqueNumber, err := GenerateRandomSecureToken(numberOfDigits)
	if err != nil {
		return "", err
	}
	uniqueYear := getLastTwoDigitsOfYear()
	day := getDayOfMonth()
	timeValue := getCurrentTime()

	generate := prefix + strconv.Itoa(uniqueNumber) + uniqueYear + day + timeValue
	return generate, nil
}

func GenerateHeaderBoq(numberOfDigits int) (string, error) {
	prefix := ""
	uniqueNumber, err := GenerateRandomSecureToken(numberOfDigits)
	if err != nil {
		return "", err
	}
	uniqueYear := getLastTwoDigitsOfYear()
	day := getDayOfMonth()
	timeValue := getCurrentTime()

	generate := prefix + strconv.Itoa(uniqueNumber) + uniqueYear + day + timeValue
	return generate, nil
}

func GenerateRandomSecureToken(numberOfDigits int) (int, error) {
	maxLimit := int(math.Pow10(numberOfDigits)) - 1
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(maxLimit)))
	if err != nil {
		return 0, err
	}
	randomNumberInt := int(randomNumber.Int64())
	return randomNumberInt, nil
}

func getLastTwoDigitsOfYear() string {
	currentTime := time.Now()
	lastTwoDigits := currentTime.Format("06")
	return lastTwoDigits
}

func getDayOfMonth() string {
	currentTime := time.Now()
	dayOfMonth := strconv.Itoa(currentTime.Day())
	return dayOfMonth
}

func getCurrentTime() string {
	currentTime := time.Now()
	timeValue := currentTime.Format("1504")
	return timeValue
}
