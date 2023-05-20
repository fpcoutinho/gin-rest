package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/fpcoutinho/gin-rest/configs"
	"github.com/fpcoutinho/gin-rest/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func All(c *gin.Context) {
	alunoCollection := configs.DB.Collection("alunos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var alunos []models.Aluno
	defer cancel()

	results, err := alunoCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer results.Close(ctx)

	if err = results.All(ctx, &alunos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"alunos": alunos})
}

func CriaAluno(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var aluno models.Aluno
	defer cancel()
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	configs.DB.Collection("alunos").InsertOne(ctx, aluno)
}

func FindAluno(c *gin.Context) {
	alunoCollection := configs.DB.Collection("alunos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var aluno models.Aluno
	defer cancel()
	m := c.Param("matricula")
	matricula, err := strconv.Atoi(m)
	if err != nil {
		panic(err)
	}

	if err := alunoCollection.FindOne(ctx, bson.M{"matricula": matricula}).Decode(&aluno); err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"aluno": aluno})
}

func DeleteAluno(c *gin.Context) {
	alunoCollection := configs.DB.Collection("alunos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var aluno models.Aluno
	defer cancel()

	m := c.Param("matricula")
	matricula, err := strconv.Atoi(m)
	if err != nil {
		panic(err)
	}

	if err := alunoCollection.FindOneAndDelete(ctx, bson.M{"matricula": matricula}).Decode(&aluno); err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aluno '" + aluno.Nome + "' de matrícula '" + m + "' deletado com sucesso."})
}

func UpdateAluno(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var aluno models.Aluno
	defer cancel()

	m := c.Param("matricula")
	matricula, err := strconv.Atoi(m)
	if err != nil {
		panic(err)
	}

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := aluno
	configs.DB.Collection("alunos").FindOneAndUpdate(ctx, bson.M{"matricula": matricula}, bson.M{"$set": bson.M{"nome": aluno.Nome, "curso": aluno.Curso, "idade": aluno.Idade}}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&result)

	if result.Matricula == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"aluno": result})
}

func FindAlunosPorCurso(c *gin.Context) {
	alunoCollection := configs.DB.Collection("alunos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var alunos []models.Aluno
	defer cancel()

	curso := c.Param("curso")

	results, err := alunoCollection.Find(ctx, bson.M{"curso": curso})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer results.Close(ctx)

	if err = results.All(ctx, &alunos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"alunos de " + curso: alunos})
}
