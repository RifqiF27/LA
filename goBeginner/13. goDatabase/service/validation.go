package service

import "errors"

func ValidasiDataDriver(name, phoneNumber, address, vehicle string) error {
	if name == "" || phoneNumber == "" || address == "" || vehicle == "" {
		return errors.New("data driver tidak lengkap")
	}
	return nil
}

func ValidasiDataCustomer(name, phoneNumber, address string) error {
	if name == "" || phoneNumber == "" || address == "" {
		return errors.New("data customer tidak lengkap")
	}
	return nil
}
