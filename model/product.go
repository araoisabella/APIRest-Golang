package model

//model para estrutura de produto levando em consideração
// as colunas da tabela no banco de dados

type Product struct {
	ID    int     `json:"id_product"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
