package models

import (
	"main/utils"
)

func NewUser(userFile, sessionFile string) *UserServices {
	service := &UserServices{
		UsersFile:    userFile,
		SessionsFile: sessionFile,
	}
	service.LoadData()
	return service
}

func (s *UserServices) LoadData() {
	utils.ReadFile(s.UsersFile, &s.Users)
	utils.ReadFile(s.SessionsFile, &s.Session)
}

func (s *UserServices) saveSession() {
	utils.WriteFile(s.SessionsFile, s.Session)
}

func (s *UserServices) Register(name, phoneNumber, username, password string) error {
	for _, user := range s.Users {
		if user.Username == username {
			return utils.ErrUsernameExists
		}
	}
	newUser := User{Name: name, PhoneNumber: phoneNumber, Username: username, Password: password}
	s.Users = append(s.Users, newUser)
	utils.WriteFile(s.UsersFile, s.Users)
	return nil
}

func (s *UserServices) Login(username, password string) error {
	for _, user := range s.Users {
		if user.Username == username && user.Password == password {
			s.Session = &user
			s.saveSession()
			return nil
		}
	}
	return utils.ErrInvalidLogin
}

func (s *UserServices) Logout() {
	s.Session = nil
	s.saveSession()
}

func (s *UserServices) IsLoggedIn() bool {
	return s.Session != nil
}

func (s *UserServices) GetLoggedInUser() *User {
	if s.IsLoggedIn() {
		return s.Session
	}
	return nil
}

func (s *UserServices) EditAccount(newName, newPhoneNumber, password string) error {
	if s.Session == nil {
		return utils.ErrNotLoggedIn
	}

	s.Session.Name = newName
	s.Session.PhoneNumber = newPhoneNumber
	s.Session.Password = password

	for i, user := range s.Users {
		if user.Username == s.Session.Username {
			s.Users[i] = *s.Session
			break
		}
	}

	utils.WriteFile(s.UsersFile, s.Users)
	s.saveSession()
	return nil
}
