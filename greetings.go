package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return name, errors.New("Nombre vacio")
	}
	message := fmt.Sprintf(RandonSaludo(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {

	saludos := make(map[string]string)
	for _, name := range names {
		message, error := Hello(name)
		if error != nil {
			return nil, error
		}
		saludos[name] = message
	}
	return saludos, nil

}

func RandonSaludo() string {

	saludos := []string{
		"Hola, Bienvenida %v ! ",
		"Hola, %v ! ",
		"Bienvenida %v ! ",
		"Te estabamos esperando %v ! "}

	return saludos[rand.Intn(len(saludos))]
}
