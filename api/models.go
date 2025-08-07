package main

type UserRecord struct {
	ID       uint32   // 4 bytes
	IsActive uint8    // 1 byte
	Username [27]byte // 27 bytes(char)
	Email    [31]byte // 31 bytes(char)
	Padding  [1]byte  // 1 bytes(completar o 64)
}

type UserPayload struct {
	ID       uint32 `json:"id"`
	IsActive uint8  `json:"is_active"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserResponse struct {
	ID       int
	IsActive int
	Username string
	Email    string
}
