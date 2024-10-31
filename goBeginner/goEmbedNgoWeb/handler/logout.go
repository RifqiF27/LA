package handler

import (
	"fmt"
	"os"
)

func Logout() {

	sessionFile := "session.json"

	if _, err := os.Stat(sessionFile); os.IsNotExist(err) {
		fmt.Println("Anda sudah logout atau belum login.")
		return
	}

	err := os.Remove(sessionFile)
	if err != nil {
		fmt.Println("Gagal menghapus sesi:", err)
		return
	}

	fmt.Println("Logout berhasil, sesi telah dihapus.")
}
