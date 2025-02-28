package main

import (
	"fmt"
	"os"
)

type Operação struct {
	Nome string
	Func func(string) error
	Arg  string
}

type Arquivo struct {
	file *os.File

	ops chan Operação

	Caminho string
}

func ArquivoTemporário(nome string) (a Arquivo, err error) {

	a.Caminho = fmt.Sprintf("/tmp/%s", nome)

	f, err := os.OpenFile(a.Caminho, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

	if err != nil {
		return
	}

	a.file = f

	return
}

func (a *Arquivo) Observar() {

	a.ops = make(chan Operação)

	for operação := range a.ops {
		operação.Func(operação.Arg)
	}
}

func (a *Arquivo) Inserir(content string) {

	a.ops <- Operação{
		"Inserir",
		a.inserir,
		content,
	}
}

func (a *Arquivo) Limpar() {

	a.ops <- Operação{
		"Limpar",
		a.limpar,
		"",
	}
}

func (a *Arquivo) inserir(content string) (err error) {

	cont := []byte(content)
	cont = append(cont, '\n')

	_, err = a.file.Write([]byte(cont))

	return
}

func (a *Arquivo) limpar(_ string /* placeholder */) (err error) {

	err = a.file.Truncate(0)
	_, err = a.file.Seek(0, 0)

	return
}
