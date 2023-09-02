package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type IReversed interface {
	reverseText() (string, error)
}

type ReverseText struct {
	text string
}

func (r *ReverseText) reverseText() (string, error) {
	text := r.text
	if text == "" {
		return "", errors.New("La cadena esta vacia")
	}
	var result string = ""
	for index := len(text) - 1; index >= 0; index-- {
		result += string(text[index])
	}
	return result, nil
}

func displayNames(names ...string) []string {
	return names
}

type Person struct {
	name, age, lastName interface{}
}

func voidInterface(names ...interface{}) {
	for i, _ := range names {
		fmt.Println(names[i])
	}
}

func currentTime() time.Time {
	return time.Now()
}
func main() {
	reverse := ReverseText{text: "kevin"}
	if result, err := reverse.reverseText(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(displayNames("kevincin", "juan", "keysy"))
	kevin := Person{name: "kevin", lastName: "dueñas", age: 22}
	fmt.Println(kevin)

	listaAny := []interface{}{1, 2, 3.5, "kevin"}
	fmt.Println(listaAny)

	voidInterface("kevin", "asael", 22, 25.5)

	fmt.Println(currentTime())

	localLocation, err := time.LoadLocation("America/Mexico_City")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(localLocation)

	fmt.Println(byte('k'))

	fmt.Println(string(rune(65)))

	myRune := []rune("kevin")
	fmt.Println(myRune)

	second := 'k'
	fmt.Println(second)

	for i := 0; i < len(myRune); i++ {
		fmt.Println(convertRuneToCharacter(myRune[i]))
	}
	saltoLinea := '\n'
	fmt.Printf("My rune unicode character: %c%c",
		second, saltoLinea)
	runSlice := []rune{'k', 'e', 'v', 'i', 'n'}
	fmt.Println(runSlice)

	fmt.Println(countBytes("kevin"))

	msg := "one"
	n1 := len(msg)
	fmt.Println(n1)

	// File handling
	content, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(content)
	n, err := os.Stdout.Write(content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)

	fileName := "text.txt"
	_, er := os.Stat(fileName)

	if os.IsNotExist(er) {
		fmt.Printf("%v file does not exist\n", fileName)
	} else {
		fmt.Printf("%v file exist\n", fileName)
	}

	fileName2 := "text2.txt"
	if _, e := os.Stat(fileName2); errors.Is(e, os.ErrNotExist) {
		fmt.Printf("%v file does not exist\n", fileName2)
	} else {
		fmt.Printf("%v file exist\n", fileName2)
	}

}

func convertRuneToCharacter(r rune) string {
	return string(r)
}
func countBytes(s string) int {
	return len(s)
}
