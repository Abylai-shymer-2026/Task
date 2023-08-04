package main

import "fmt"

type Battery struct {
	text string
}

func (s Battery) String() string {
	a := 0
	for _, i := range s.text {
		if i == '1' {
			a++
		}
	}
	sc := ""
	ss := ""
	for i := 0; i < 10-a; i++ {
		ss += " "
	}
	for i := 0; i < a; i++ {
		sc += "X"
	}
	return fmt.Sprintf("[%s%s]", ss, sc)

}
func main() {
	var c string
	fmt.Scan(&c)

	batteryForTest := Battery{text: c}
	fmt.Print(batteryForTest)
}
