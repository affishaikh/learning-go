package main

import (
	"fmt"
	"strings"
)

func getCipherAlphabet(plainTextAlphabet, keywordAlphabet string) (cipherAlphabet string) {
	alphabets := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	indexOfCipherAlphabet := strings.Index(
		alphabets,
		strings.ToUpper(keywordAlphabet),
	) + strings.Index(
		alphabets,
		strings.ToUpper(plainTextAlphabet),
	)
	if indexOfCipherAlphabet >= 26 {
		indexOfCipherAlphabet = indexOfCipherAlphabet - 26
	}
	cipherAlphabet = string(alphabets[indexOfCipherAlphabet])
	return
}

func getPlainAlphabet(cipherTextAlphabet, keywordAlphabet string) (plainAlphabet string) {
	alphabets := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	indexOfPlainAlphabet := strings.Index(
		alphabets,
		strings.ToUpper(keywordAlphabet),
	) - strings.Index(
		alphabets,
		strings.ToUpper(cipherTextAlphabet),
	)
	if indexOfPlainAlphabet >= 26 {
		indexOfPlainAlphabet = indexOfPlainAlphabet - 26
	}
	plainAlphabet = string(alphabets[indexOfPlainAlphabet])
	return
}

func vignereCipherEncrypt(plainText, key string) (cipherText string) {
	for i := 0; i < len(plainText); i++ {
		cipherText += getCipherAlphabet(string(plainText[i]), string(key[i % 4]))
	}
	return
}

func vignereCipherDecrypt(cipherText, key string) (plainText string) {
	for i := 0; i < len(cipherText); i++ {
		plainText += getCipherAlphabet(string(cipherText[i]), string(key[i % 4]))
	}
	return
}

func main() {
	fmt.Println(vignereCipherEncrypt("thesunandthemaninthemoon", "king"))
	fmt.Println(vignereCipherDecrypt("DPRYEVNTNBUKWIAOXBUKWWBT", "king"))
}
