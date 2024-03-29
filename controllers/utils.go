package controllers

import (
	"crypto/rand"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"ulfr/config"
)

func Include(path string) []string {
	files, _ := filepath.Glob("views/templates/*.html")
	libFiles, _ := filepath.Glob("views/templates/lib/*")
	pathFiles, _ := filepath.Glob("views/" + path + "/*.html")

	files = append(files, append(pathFiles, libFiles...)...)

	return files
}

func GenerateRandomString(length int) (string, error) {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	maxIndex := big.NewInt(int64(len(charSet)))

	randomString := make([]byte, length)

	for i := range randomString {
		randomIndex, err := rand.Int(rand.Reader, maxIndex)
		if err != nil {
			return "", err
		}
		randomString[i] = charSet[randomIndex.Int64()]
	}

	return string(randomString), nil
}

func WriteToFile(path string, data string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(data)
}

func SelfControl(host string) bool {
	if host == config.HostName || host == config.HostName+":"+config.Port {
		return true
	}
	return false
}

func validIP4(ipAddress string) bool {
	ipAddress = strings.Trim(ipAddress, " ")

	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ipAddress) {
		return true
	}
	return false
}
