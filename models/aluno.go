package models

import "gopkg.in/validator.v2"

type Aluno struct {
	Matricula string `json:"matricula" bson:"matricula" validate:"len=11,regexp=^20([0-1][0-9]|[2][0-3])\\d{7}$"`
	Nome      string `json:"nome" bson:"nome" validate:"min=2,max=50"`
	Idade     int    `json:"idade" bson:"idade" validate:"min=15,max=100"`
	Curso     string `json:"curso" bson:"curso" validate:"nonzero"`
}

func ValidaDadosDeAluno(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}
	return nil
}
