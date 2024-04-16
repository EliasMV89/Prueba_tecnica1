package main

import (
	"fmt"
	"os"
)

/*import (
	"fmt"
	"strings"
)

type Player struct {
	Name       string
	Identifier int
}

func main() {
	//Lista jugadores

	players := []string{"John Smith", "David Johnson", "Michael Garcia", "Sarah Williams", "Daniel Martinez", "Emily Brown", "James Rodriguez", "Sophia Lee", "Lucas Oliveira", "Olivia Taylor", "Mateo Hernandez", "Emma Wilson", "Gabriel Silva"}

	newPlayers := []string{"New player 1", "New player 2", "New player 3", "New player 4", "New player 5", "New player 6", "New player 7"}

	//borra las siete primeras posiciones
	players = players[7:]

	//agregar
	for _, name := range newPlayers {
		players = append(players, name)
	}

	//Mostrar lista completa
	fmt.Println("Lista completa de jugadiores: ")
	for i, player := range players {
		fmt.Println("%d. %s\n", i+1001, sanitizeName(player))
	}

}


func sanitizeName(name string) string {

	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == ' ' {
			return r
		}
		return -1
	}, name)
}*/

// FizzBuzz

/*func main() {

	n:=100

	for i:=1; i<=n; i++{

		if i%3==0 && i%5==0 {
			fmt.Println("fizzbuzz")
		}else if i%3==0 {
			fmt.Print("fizz")
		}else if i%5==0 {
			fmt.Println("buzz")
		}else {
			fmt.Println(i)
		}
		if i < n {
			fmt.Println(", ")
		}
	}
}*/

//menu-pacckaging

var (
	adminUsername      = "admin"
	adminPassword      = "root"
	supervisorUsername = "seeker223"
	superisorPassword  = "seekr"
	adminActions       = []string{"Crear laborer", "Eliminar laborer", "Generar archivos"}
	laborers           []string
)

func main() {
	for {
		fmt.Println("Bienvenido a Labora")
		fmt.Println("1. Iniciar sesion como administrador")
		fmt.Println("2. Iniciar  sesion como supervisor")
		fmt.Println("3. Salir")

		var choice int
		fmt.Println("Ingrese su opcion")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			adminLogin()
		case 2:
			supervisorLogin()
		case 3:
			fmt.Println("¡Hasta luego!")
			os.Exit(0)
		default:
			fmt.Println("Opcion ninvalida. Por favor, intente de nuevo")
		}
	}
}

func adminLogin() {
	fmt.Println("Ingrese las credenciales  de Administrador")
	fmt.Println("Usuario: ")
	var username string
	fmt.Scanln(&username)
	fmt.Println("Contraseña: ")
	var password string
	fmt.Scanln(&password)

	if username == adminUsername && password == adminPassword {
		fmt.Println("Inicio de sesion exitoso como Administrador")
		adminMenu()
	} else {
		fmt.Println("Credenciales incorrectas. Por favor, intente de nuevo")
	}
}

func supervisorLogin() {
	fmt.Println("Ingrese las credenciales de Supervisor")
	fmt.Println("Usuario: ")
	var username string
	fmt.Scanln(&username)
	fmt.Println("Contraseña: ")
	var password string
	fmt.Scanln(&password)

	if username == supervisorUsername && password == superisorPassword {
		fmt.Println("Inicio de sesion exitoso como Supervisor")
		//supervisorMenu()
	} else {
		fmt.Println("Credenciales incorrectas. Por favor, intente de nuevo")
	}
}

func adminMenu() {
	for {
		fmt.Println("Menu de Administrador")
		for i, action := range adminActions {
			fmt.Printf("%d. %s\n", i+1, action)
		}
		fmt.Println("0. Cerrar sesion")

		var choice int
		fmt.Println("Seleccione su opcion: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Accion: Crear laborer")
			createLaborer()
		case 2:
			fmt.Println("Accion: Eliminar laborer")
			deleteLaborer()
		case 3:
			fmt.Println("Accion: Generar archivo de texto personalizado")
			generateTextFile()
		case 0:
			fmt.Println("Cerrando sesion de Administrador")
			return
		default:
			fmt.Println("Opcion invalida. Por favor, intente de nuevo")
		}
	}
}

func createLaborer() {
	laborers = append(laborers, fmt.Sprintf("laborer %d", len(laborers)))
	fmt.Println("Laborer creado exitosamente")
}

func deleteLaborer() {
	if len(laborers) == 0 {
		fmt.Println(("No hay laborers para eliminar"))
		return
	}

	laborers = laborers[:len(laborers)-1]
	fmt.Println("laborer eliminado exitosamente")
}

func generateTextFile() {
	//fileName := fmt.Sprintf("mensajes_%s.txt", time.Now().Format("20060102_152032"))

	//Preguntar al usuario el nombre del archivo
	fmt.Print("Ingrese el nombre del archivo: ")
	var fileName string
	fmt.Scanln(&fileName)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error al crear archivo: %v\n", err)
		return
	}
	defer file.Close()
	for _, laborer := range laborers {
		_, err := file.WriteString(fmt.Sprintf("%s\n", laborer))
		if err != nil {
			fmt.Printf("error al escribir el archivo %v\n", err)
			return
		}
	}
	fmt.Printf("Archivo %s creado exitosamente. \n", fileName)

	//Contar la cantidad de letras en el nombre del archivo
	letterCount := len(fileName)
	fmt.Println("Cantidad de letras en el nombre del archivo: ", letterCount)
}
