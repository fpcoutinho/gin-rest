package models

type Aluno struct {
	Matricula int    `json:"matricula" bson:"matricula"`
	Nome      string `json:"nome" bson:"nome"`
	Idade     int    `json:"idade" bson:"idade"`
	Curso     string `json:"curso" bson:"curso"`
}
