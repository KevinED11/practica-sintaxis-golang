package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
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
	fmt.Printf("The content is %v", string(content))
	fmt.Println("contentido leido con exito")

	fileName := "text.txt"

	_, er := os.Stat(fileName)
	if os.IsNotExist(er) {
		fmt.Printf("%v file does not exist\n", fileName)
	} else {
		fmt.Printf("%v file exist\n", fileName)
	}

	if _, er := os.Stat(fileName); os.IsNotExist(er) {
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

	showNames("kevin", "juan", "keysy", "asael", "kevin")

	newMap := make(map[string]interface{})
	newMap["name"] = "kevin"
	newMap["age"] = 22
	newMap["lastName"] = "dueñas"
	fmt.Println(newMap)
	fmt.Println(newMap["name"])

	for key, value := range newMap {
		fmt.Println(key, value)
	}

	counter := 1
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	kevinc := struct {
		name string
		age  int
	}{"kevin", 22}

	fmt.Printf("%+v\n", kevinc)

	newArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(newArray)
	fmt.Println(newSlice)

	fmt.Println(time.Now().Weekday())

	switch day := time.Now().Weekday() - 1; day {
	case 0:
		fmt.Println("Lunes")
	case 1:
		fmt.Println("Martes")
	case 2:
		fmt.Println("Miercoles")
	case 3:
		fmt.Println("Jueves")
	case 4:
		fmt.Println("Viernes")
	case 5:
		fmt.Println("Sabado")
	case 6:
		fmt.Println("Domingo")
	default:
		fmt.Println("Unknown day")
	}

	kevinAge := kevinc.age

	switch {
	case kevinAge < 18:
		fmt.Println("Eres menor de edad")
	case kevinAge >= 18:
		fmt.Println("Eres mayor de edad")
	default:
		fmt.Println("Edad desconocida")

	}

	fmt.Println("For loop:")
	num := 5
	for i := 1; i <= num; i++ {
		defer fmt.Println(i)
	}

	for i, letter := range "hola mundo" {
		fmt.Println(i, string(letter))
	}

	/*
		for {
			fmt.Println("hola")
		}
	*/

	f, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Println(f.Name())

	// Se imprime al final
	defer fmt.Println("primera")
	// Se imprime despues de "primera"
	defer fmt.Println("segunda")
	fmt.Println("Hola mundo")
	// Se imprime despues de "Hola mundo"
	defer fmt.Println("tercera")

	// buffer
	var buf bytes.Buffer
	buf.Write([]byte("Hola me llamo kevin asael"))
	buf.WriteByte(32)
	buf.WriteByte(32)
	buf.WriteString(" espinoza duéñas")

	fmt.Println(buf)
	fmt.Println(buf.String())

	goVer := runtime.Version()
	osVer := runtime.GOOS
	arch := runtime.GOARCH

	fmt.Println(goVer)
	fmt.Println(osVer)
	fmt.Println(arch)
}
func showNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func convertRuneToCharacter(r rune) string {
	return string(r)
}
func countBytes(s string) int {
	return len(s)
}
