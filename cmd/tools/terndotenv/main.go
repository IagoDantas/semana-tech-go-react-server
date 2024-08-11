package main

import (
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	// Tentar carregar o arquivo .env
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	
	// Executar o comando tern
	cmd := exec.Command(
		"tern", 
		"migrate", 
		"--migrations", 
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	if err := cmd.Run(); err != nil {
		panic(err)
	}

}