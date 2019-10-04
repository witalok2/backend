package model

type Atividade struct {
	Id        int    `json:"id" db:"id"`
	Titulo    string `json:"titulo" db:"atividade_titulo"`
	SubTitulo string `json:"subtitulo" db:"atividade_subtitulo"`
	Descricao string `json:"descricao" db:"atividade_descricao"`
	Situacao  string `json:"situacao" db:"atividade_situacao"` // Pendente, Em Andamento, Finalizadas
}

type Atividades struct {
	Atividades []Atividade `json:"atividades"`
}
