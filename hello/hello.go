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
	"os"
	"strings"
)

func main() {
	var name := "AB12cdEF34gh"
	fmt.Printf("Your password is %s. Have fun! :)", name)
}

func input() {
	// fmt.Println(stringutil.Reverse("!selpmaxe oG ,olleH"))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	city, _ := reader.ReadString('\n')
	var name = strings.TrimSuffix(city, "\r\n")
	fmt.Printf("Hi, %s. Welcome! :)", name)

}
