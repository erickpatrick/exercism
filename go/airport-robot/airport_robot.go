package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type Italian struct{}

type Portuguese struct{}

func SayHello(visitorName string, robotGreeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", robotGreeter.LanguageName(), robotGreeter.Greet(visitorName))
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}
