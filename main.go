package main

import "fmt"

func main() {
	//uma função pode retornar mais de um valor
	//a construção := cria a variável caso ela não exista, já definindo seu tipo
	ok, err := say("Hello World")
	if err != nil {
		panic(err.Error())
	}
	switch ok {
	case true:
		fmt.Println("Deu certo")
	default:
		fmt.Println("deu errado")
	}

}

//as funções devem declarar o tipo de cada variável que recebe ou que retorna
func say(what string) (bool, error) {
	if what == "" {
		return false, fmt.Errorf("Empty string")
	}
	fmt.Println(what)
	return true, nil
}
