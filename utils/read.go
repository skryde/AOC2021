package utils

import (
	"os"
	"strings"
)

func ReadFile(fileName string) ([]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// Se eliminan los blank characters que puedan haber al inicio o al final del string.
	// Particularmente me interesa eliminar la ultima linea en blanco que tiene el archivo de entrada.
	noBlankChars := strings.TrimSpace(string(data))
	return strings.Split(noBlankChars, "\n"), nil
}
