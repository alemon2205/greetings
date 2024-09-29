# Saludos en GO

##Instalación: 
Ejecuta el sisguiente comando para instalar el paguete:
``` bash
go get -u github.com/alemon2205/greetings
```

 ## Uso
 Aqui tienes un ejemplo de ocmo utilizar el paquete en tu código:
 
 ``` go
 package main

import (
	"fmt"
	"log"

	"github.com/alemon2205/greetings"
)

func main() {

	log.SetPrefix("greetings:")
	//log.SetFlags(0)
	//message, err := greetings.Hello("ale")
	names := []string{"Ale", "Fer", "Bruno", "Sofi"}
	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	for _, saludo := range message {
		fmt.Println(saludo)
	}
}
```    



