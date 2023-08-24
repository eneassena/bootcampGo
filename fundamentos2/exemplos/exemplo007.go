package exemplos

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type produto struct {
	ID         int
	quantidade int
	preco      float64
}

func (p produto) geraLinhaCSV() string {
	return fmt.Sprintf("\n%d,%d,%.2f", p.ID, p.quantidade, p.preco)
}

func (p produto) gerarCabecalhoCSV() string {
	return "id,quantidade,preco"
}

func geraCSV(caminho string, produto []produto) error {
	// verifica se a lista de produto esta vasia
	if len(produto) == 0 {
		return errors.New("quantidade do produto invalida")
	}

	// abri um arquivo
	file, err := os.OpenFile(caminho, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo: %w", err)
	}

	defer file.Close()

	p := produto[0]

	fmt.Println(p)

	if _, err = file.WriteString(p.gerarCabecalhoCSV()); err != nil {
		return fmt.Errorf("Erro ao gerar cabecalho: %w", err)
	}

	for _, p = range produto {
		if _, err = file.WriteString(p.geraLinhaCSV()); err != nil {
			return fmt.Errorf("erro ao salvar linhas %w", err)
		}
	}
	return nil
}

func criaArquivo(filename string, content []byte, permission fs.FileMode) error {
	return os.WriteFile(filename, content, permission)
}
