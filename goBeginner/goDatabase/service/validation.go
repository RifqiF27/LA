package service

import "errors"

// ValidasiDataDriver memeriksa apakah semua data driver valid sebelum disimpan
func ValidasiDataDriver(name, phoneNumber, address, vehicle string) error {
	if name == "" || phoneNumber == "" || address == "" || vehicle == "" {
		return errors.New("data driver tidak lengkap")
	}
	return nil
}

// ValidasiDataCustomer memeriksa apakah semua data customer valid sebelum disimpan
func ValidasiDataCustomer(name, phoneNumber, address string) error {
	if name == "" || phoneNumber == "" || address == "" {
		return errors.New("data customer tidak lengkap")
	}
	return nil
}
