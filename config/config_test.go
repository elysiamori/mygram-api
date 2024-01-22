package config

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestConfi(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Gagal memuat file .env: %v", err)
	}
}
