package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

func getCred() {
	key, cred, _ := anaconda.AuthorizationURL("")
	fmt.Printf("access: %v\n", key)
	fmt.Printf("Key: ")
	buf := bufio.NewReader(os.Stdin)
	line, _, _ := buf.ReadLine()
	cred, _, _ = anaconda.GetCredentials(cred, string(line))
	fmt.Printf("cred: %s\n", cred)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(a, e) {
			return true
		}
	}
	return false
}

func count(target, s string) int {
	cnt := 0
	for _, a := range []rune(s) {
		if strings.ContainsRune(target, a) {
			cnt++
		}
	}
	return cnt
}
