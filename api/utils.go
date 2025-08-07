package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func convert_to_bytes(user UserRecord) (io.Reader, error) {
	buf := new(bytes.Buffer)
	// joga os bytes de user para o buffer
	// feito campo a campo para ter certeza que vai estar de acordo no C
	err := binary.Write(buf, binary.LittleEndian, user.ID)
	if err != nil {
		fmt.Println("Erro ao converter para binario o campo: 'id': %w\n ", err)
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, user.IsActive)
	if err != nil {
		fmt.Println("Erro ao converter para binario o campo 'isActive': %w\n", err)
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, user.Username)
	if err != nil {
		fmt.Println("Erro ao converter para binario o campo 'username': %w\n ", err)
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, user.Email)
	if err != nil {
		fmt.Println("Erro ao converter para binario o campo 'email': %w\n", err)
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, user.Padding)
	if err != nil {
		return nil, fmt.Errorf("falha ao escrever Padding: %w", err)
	}

	return buf, nil
}

func convert_to_entity(output []string) []UserResponse {
	var users []UserResponse

	for i := range output {
		record := UserResponse{}
		field_with_name := strings.Split(output[i], ";")
		for _, elem := range field_with_name {
			fields := strings.Split(elem, ":")
			if len(fields) == 2 {
				key := fields[0]
				value := fields[1]
				switch key {
				case "id":
					record.ID, _ = strconv.Atoi(value)
				case "is_active":
					fmt.Println(value)
					record.IsActive, _ = strconv.Atoi(value)
				case "username":
					record.Username = value
				case "email":
					record.Email = value
				}
			}
		}
		if record.ID > 0 {
			users = append(users, record)
		}
	}
	return users
}
