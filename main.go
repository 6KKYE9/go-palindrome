package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// 判断是不是回文。忽略大小写、空白和标点，只比字母数字。
func isPalindrome(s string) bool {
	var clean []rune
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			clean = append(clean, r)
		}
	}
	for i := 0; i < len(clean)/2; i++ {
		if clean[i] != clean[len(clean)-1-i] {
			return false
		}
	}
	return len(clean) > 0
}

func main() {
	args := os.Args[1:]
	var texts []string
	for _, a := range args {
		if a == "-h" || a == "--help" {
			fmt.Println("go-palindrome 判断回文")
			fmt.Println("用法: go-palindrome <文本> ...  或  管道 | go-palindrome")
			return
		}
		texts = append(texts, a)
	}
	if len(texts) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			texts = []string{scanner.Text()}
		}
	}
	for _, t := range texts {
		if isPalindrome(t) {
			fmt.Printf("%s => 是回文\n", t)
		} else {
			fmt.Printf("%s => 不是\n", t)
		}
	}
}
