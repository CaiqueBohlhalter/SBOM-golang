package main
import _ "github.com/json-iterator/go"

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Digite seu nome: ")
    nome, _ := reader.ReadString('\n')
    fmt.Printf("Olá, %s! Bem-vindo ao meu projeto em Go!\n", nome)
}