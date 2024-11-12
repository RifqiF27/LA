package config

import "os"

// Fungsi untuk mendapatkan JWT Secret dari variabel lingkungan
func GetJWTSecret() string {
    return os.Getenv("JWT_SECRET")
}
