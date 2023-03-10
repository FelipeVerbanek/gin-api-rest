package main

import (
	"bytes"
	"encoding/json"
	"gin-api-rest/controllers"
	database "gin-api-rest/db"
	"gin-api-rest/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func InputDB() database.InputConnectDB {
	return database.InputConnectDB{
		Host:     "localhost",
		User:     "postgres",
		Password: "postgres",
		Dbname:   "api_gin",
		Port:     "5432",
	}
}

func SetupRoutesTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.Default()
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Teste", CPF: "99999999999", RG: "999999999"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeleteAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestListAlunosHandler(t *testing.T) {
	database.Connection(InputDB())
	CriaAlunoMock()
	defer DeleteAlunoMock()

	r := SetupRoutesTest()
	r.GET("/api/alunos", controllers.ListarAlunos)

	req, _ := http.NewRequest("GET", "/api/alunos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestBuscaAlunoCPFHandler(t *testing.T) {
	database.Connection(InputDB())
	CriaAlunoMock()
	defer DeleteAlunoMock()
	r := SetupRoutesTest()

	r.GET("/api/alunos/cpf/:cpf", controllers.BuscaALunoCPF)
	req, _ := http.NewRequest("GET", "/api/alunos/cpf/99999999999", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.Connection(InputDB())
	CriaAlunoMock()
	defer DeleteAlunoMock()
	r := SetupRoutesTest()
	r.GET("/api/alunos/:id", controllers.BuscaAlunoID)

	path := "/api/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var alunoMock models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Teste", alunoMock.Nome)
	assert.Equal(t, "99999999999", alunoMock.CPF)
	assert.Equal(t, "999999999", alunoMock.RG)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.Connection(InputDB())
	CriaAlunoMock()
	defer DeleteAlunoMock()

	r := SetupRoutesTest()
	r.DELETE("/api/alunos/:id", controllers.DeleteAluno)

	path := "/api/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var aluno models.Aluno
	database.DB.First(&aluno, ID)
	assert.Equal(t, int(aluno.ID), 0)
}

func TestEditaAlunoHandler(t *testing.T) {
	database.Connection(InputDB())
	CriaAlunoMock()

	r := SetupRoutesTest()
	r.PATCH("/api/alunos/:id", controllers.EditarAluno)

	aluno := models.Aluno{Nome: "Testando Editar", CPF: "12345678912", RG: "123123945"}
	path := "/api/alunos/" + strconv.Itoa(ID)
	resultJson, _ := json.Marshal(aluno)

	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(resultJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunoMockAtualizado)

	assert.Equal(t, "Testando Editar", alunoMockAtualizado.Nome)
	assert.Equal(t, "12345678912", alunoMockAtualizado.CPF)
	assert.Equal(t, "123123945", alunoMockAtualizado.RG)
}
