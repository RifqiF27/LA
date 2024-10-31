package service

import (
	"main/model"
	"main/reposiitory"
)

type StudentService struct {
	Repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) *StudentService {
	return &StudentService{Repo: repo}
}

func (ss *StudentService) AddStudent(userId int, name, phoneNumber, address string) error {

	student := model.Student{
		UserID:      userId,
		Name:        name,
		PhoneNumber: phoneNumber,
		Address:     address,
	}

	err := ss.Repo.AddStudent(&student)
	if err != nil {
		return err
	}

	return nil
}

func (ss *StudentService) GetAllStudents() ([]model.Student, error) {
	var students []model.Student
	err := ss.Repo.GetAllStudents(&students)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (ss *StudentService) UpdateStudent(ID uint16, userId int, name, phoneNumber, address string) error {
	student := &model.Student{
		ID:          ID,
		UserID:      userId,
		Name:        name,
		PhoneNumber: phoneNumber,
		Address:     address,
	}

	// fmt.Println(student,"<<<<")
	return ss.Repo.UpdateStudent(student)
}

func (ss *StudentService) DeleteStudent(id, userID int) error {
	return ss.Repo.DeleteStudent(id, userID)
}
