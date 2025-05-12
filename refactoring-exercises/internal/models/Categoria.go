package models

// Modelo para a tabela "categorias"
type Categoria struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	Nome      string `json:"nome"`
	Codigo    string `gorm:"unique" json:"codigo"`
	Descricao string `json:"descricao"`
}
