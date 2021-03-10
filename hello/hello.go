/*
Copyright 2014 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// pw := "AB12cdEF34gh"
	lowerCaseLetters := "abcdefghijklmnopqrstuvwxyz"
	// upperCaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	upperCaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// rand.Seed(seed int64)
	rand.Seed(time.Now().UnixNano())

	char1 := upperCaseLetters[rand.Intn(len(upperCaseLetters))]
	char2 := upperCaseLetters[rand.Intn(len(upperCaseLetters))]
	char3 := rand.Intn(9)
	char4 := rand.Intn(9)
	char5 := lowerCaseLetters[rand.Intn(len(lowerCaseLetters))]
	char6 := lowerCaseLetters[rand.Intn(len(lowerCaseLetters))]
	char7 := upperCaseLetters[rand.Intn(len(upperCaseLetters))]
	char8 := upperCaseLetters[rand.Intn(len(upperCaseLetters))]
	char9 := rand.Intn(9)
	char10 := rand.Intn(9)
	char11 := lowerCaseLetters[rand.Intn(len(lowerCaseLetters))]
	char12 := lowerCaseLetters[rand.Intn(len(lowerCaseLetters))]

	// fmt.Printf("Your password's first character is an upper-case letter: %c (from %d). Have fun! :)\n\n", firstChar, firstChar)
	fmt.Printf("Your password is %c%c%d%d%c%c%c%c%d%d%c%c. Have fun! :)\n\n", char1, char2, char3, char4, char5, char6, char7, char8, char9, char10, char11, char12)

	/*
		for k, v := range s {
			fmt.Printf("k：%d,v：%c == %d\n", k, v, v)
		}
	//*/
}

func input() {
	// fmt.Println(stringutil.Reverse("!selpmaxe oG ,olleH"))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	city, _ := reader.ReadString('\n')
	var name = strings.TrimSuffix(city, "\r\n")
	fmt.Printf("Hi, %s. Welcome! :)", name)

}
