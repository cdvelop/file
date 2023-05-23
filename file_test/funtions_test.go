package file_test

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() string {
	// Crear una nueva fuente de números aleatorios
	source := rand.NewSource(time.Now().UnixNano())

	// Crear un nuevo generador de números aleatorios a partir de la fuente
	generator := rand.New(source)

	// Generar un número aleatorio de 3 cifras
	number := generator.Intn(900) + 100

	return fmt.Sprint(number)
}
