package main

import (
	"bufio"
	"fmt"
	"os"
)

func encrypeted_position(letter byte) int {
	alphabet := "abcdefghijklmnopqrstuvwxyz "
	for i := 0; i < len(alphabet); i++ {
		if letter == alphabet[i] {
			return i
		}
	}
	return -1
}

func encrypeted_text(encrypted_num []int, lenght int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz "
	text := ""

	for j := 0; j < lenght; j++ {
		text += string(alphabet[encrypted_num[j]])
	}
	return text
}

func encrypt(plainText string, key int) string {
	var C [100]int

	for i := 0; i < len(plainText); i++ {
		C[i] = (encrypeted_position(plainText[i]) + key) % 27
	}
	return encrypeted_text(C[:], len(plainText))
}

func decrypt(cipherText string, key int) string {
	var C [100]int

	for i := 0; i < len(cipherText); i++ {
		if (encrypeted_position(cipherText[i]) - key) >= 0 {
			C[i] = (encrypeted_position(cipherText[i]) - key) % 27
		} else {
			C[i] = 27 - (key-encrypeted_position(cipherText[i]))%27
		}

	}
	return encrypeted_text(C[:], len(cipherText))
}

func main() {
	var alphabet = "a b c d e f g h i  j  k  l  m  n  o  p  q  r  s  t  u  v  w  x  y  z  _"
	var pos = "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27"
	fmt.Println("Used alphabet: " + alphabet)
	fmt.Println("Position:      " + pos)

	fmt.Println("Enter a string to encrypt or decrypt: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	plainText := scanner.Text()
	fmt.Println(plainText)

	fmt.Println("Enter a key value for encryption: ")
	var key int
	fmt.Scanln(&key)

	fmt.Println("Please select an option: \n 1. Encrypt \n 2. Decrypt")
	var option int
	fmt.Scanln(&option)

	if option == 1 {
		fmt.Println("Encrypted text: " + encrypt(plainText, key))
	} else if option == 2 {
		fmt.Println("Decrypted text: " + decrypt(plainText, key))
	} else {
		fmt.Println("Invalid option selected")
	}
}
