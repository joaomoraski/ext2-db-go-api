package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

const EXEC_PATH = "../ext2-db-engine/ext2shell"
const DB_NAME = "user_record"

func insertRecord(record UserRecord) (string, error) {
	recordBytes, err := convert_to_bytes(record)
	if err != nil {
		return "", fmt.Errorf("erro ao converter para binario: %w", err)
	}

	cmd := exec.Command(EXEC_PATH, "db", "insert", DB_NAME)
	cmd.Dir = "../ext2-db-engine/"

	// joga na stdin os bytes
	cmd.Stdin = recordBytes

	output, err := cmd.CombinedOutput()
	if err != nil || strings.Contains(string(output), "Erro de Banco de Dados") {
		return string(output), errors.New("erro de Banco de Dados")
	}
	return string(output), nil
}

func getRecords(limit int, filters string) ([]UserResponse, string, error) {
	var cmd *exec.Cmd
	if filters == "" {
		cmd = exec.Command(EXEC_PATH, "db", "select", DB_NAME, strconv.Itoa(limit))
	} else {
		cmd = exec.Command(EXEC_PATH, "db", "select-where", DB_NAME, filters, strconv.Itoa(limit))
	}
	cmd.Dir = "../ext2-db-engine/"

	output, err := cmd.CombinedOutput()

	users := convert_to_entity(strings.Split(string(output), "\n"))
	if err != nil {
		return nil, string(output), fmt.Errorf("erro ao converter para binario: %w", err)
	}
	return users, string(output), nil
}
