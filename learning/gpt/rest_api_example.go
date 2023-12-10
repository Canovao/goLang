package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Estrutura de dados para representar um objeto simples
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// "Banco de dados" simulado em memória
var people = []Person{
	{1, "Alice"},
	{2, "Bob"},
	{3, "Charlie"},
}

// Manipulador para a rota "/people"
func getPeopleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

// Manipulador para a rota "/people/{id}"
func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extrai o parâmetro da URL
	id := extractIDFromURL(r)
	if id == -1 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Procura a pessoa no "banco de dados"
	var person Person
	for _, p := range people {
		if p.ID == id {
			person = p
			break
		}
	}

	// Retorna a pessoa encontrada ou um erro se não encontrada
	if person.ID != 0 {
		json.NewEncoder(w).Encode(person)
	} else {
		http.Error(w, "Person not found", http.StatusNotFound)
	}
}

// Função utilitária para extrair o parâmetro da URL
func extractIDFromURL(r *http.Request) int {
	id := -1
	fmt.Sscanf(r.URL.Path, "/people/%d", &id)
	return id
}

func main() {
	// Configuração dos manipuladores
	http.HandleFunc("/people", getPeopleHandler)
	http.HandleFunc("/people/", getPersonHandler)

	// Inicia o servidor na porta 8080
	port := 8080
	log.Printf("Server listening on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
