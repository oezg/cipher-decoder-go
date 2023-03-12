package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	bobSecretKey   = 7
	bobMessage     = "Will you marry me?"
	happyReaction  = "Great!"
	sadReaction    = "What a pity!"
	positiveAnswer = "Yeah, okay!"
	negativeAnswer = "Let's be friends."
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var numbers = readNumbers(*reader)
	var publicBase, publicMode = numbers[0], numbers[1]
	fmt.Println("OK")
	var alicePublicKey = readNumbers(*reader)[0]
	var bobPublicKey = modularExponentiation(publicBase, bobSecretKey, publicMode)
	fmt.Printf("B is %d\n", bobPublicKey)
	var sharedSecret = modularExponentiation(alicePublicKey, bobSecretKey, publicMode)
	fmt.Println(encryptCaesarCipher(bobMessage, sharedSecret))
	showReaction(readLine(*reader), sharedSecret)
}

func modularExponentiation(base, exponent, mode int) int {
	var result = 1
	for i := 0; i < exponent; i++ {
		result = result * base % mode
	}
	return result
}

func encryptCaesarCipher(message string, secret int) string {
	var out []rune
	var substitute rune
	var rangeLength = 'Z' - 'A' + 1
	for _, ch := range message {
		if 'A' <= ch && ch <= 'Z' {
			substitute = (ch-'A'+rune(secret))%rangeLength + 'A'
		} else if 'a' <= ch && ch <= 'z' {
			substitute = (ch-'a'+rune(secret))%rangeLength + 'a'
		} else {
			substitute = ch
		}
		out = append(out, substitute)
	}
	return string(out)
}

func decryptCaesarCipher(cipher string, secret int) string {
	return encryptCaesarCipher(cipher, 26-secret%26)
}

func showReaction(answer string, secret int) {
	switch decryptCaesarCipher(answer, secret) {
	case positiveAnswer:
		fmt.Println(encryptCaesarCipher(happyReaction, secret))
	case negativeAnswer:
		fmt.Println(encryptCaesarCipher(sadReaction, secret))
	}
}

func readNumbers(reader bufio.Reader) (numbers []int) {
	for _, word := range strings.Split(readLine(reader), " ") {
		if num, err := strconv.Atoi(word); err == nil {
			numbers = append(numbers, num)
		}
	}
	return
}

func readLine(reader bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(input)
}
