package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload UserPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Corpo da requisiçao invalido", http.StatusBadRequest)
		return
	}

	if payload.Username == "" {
		http.Error(w, "O campo 'Username' não pode estar vazio", http.StatusBadRequest)
		return
	}
	if len(payload.Username) > 27 {
		http.Error(w, "O campo 'Username' esta muito longo", http.StatusBadRequest)
		return
	}
	if payload.Email == "" {
		http.Error(w, "O campo 'Email' não pode estar vazio", http.StatusBadRequest)
		return
	}
	if len(payload.Email) > 31 {
		http.Error(w, "O campo 'Email' esta muito longo", http.StatusBadRequest)
		return
	}

	var record UserRecord
	record.ID = payload.ID
	record.IsActive = payload.IsActive
	// copiar as strings para o array de bytes
	copy(record.Username[:], []byte(payload.Username))
	copy(record.Email[:], []byte(payload.Email))

	output, err := insertRecord(record)
	if err != nil {
		fmt.Printf("Erro ao inserir usuário: %s\n", output)
		http.Error(w, "Erro ao inserir usuário", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payload)
}

func SelectUserHandler(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	var filters string
	if r.URL.Query().Get("filters") != "" {
		filters = r.URL.Query().Get("filters")
	}

	users, output, err := getRecords(limit, filters)
	if err != nil {
		fmt.Printf("Erro ao listar usuários: %s\n", output)
		http.Error(w, "Erro ao listar usuários", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		fmt.Printf("Erro ao codificar resposta Json %v\n", err)
	}
}
