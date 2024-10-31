package handler

import (
	"database/sql"
	"fmt"
	"main/reposiitory"
	"main/service"
	"main/utils"
)

func session() (int, bool) {
	session, err := utils.ReadSession()
	if err != nil {
		utils.SendJSONResponse(401, "Unauthorized", nil)
		return 0, false
	}
	role, ok := session["Role"].(string)
	userID, okID := session["ID"].(float64)
	if !ok || role != "admin" || !okID {
		utils.SendJSONResponse(403, "Forbidden", nil)
		return 0, false
	}

	return int(userID), role == "admin"
}

func AddStudent(db *sql.DB) {

	// session, err := utils.ReadSession()
	// if err != nil {
	// 	utils.SendJSONResponse(400, "Gagal membaca session", nil)
	// 	return
	// }
	// userID, ok := session["ID"].(float64)
	// if !ok {
	// 	utils.SendJSONResponse(400, "ID tidak ditemukan di session", nil)
	// 	return
	// }
	userID, valid := session()
	if !valid {
		return
	}

	var name, phoneNumber, address string
	fmt.Print("Masukkan nama siswa: ")
	fmt.Scan(&name)

	fmt.Print("Masukkan nomor telepon: ")
	fmt.Scan(&phoneNumber)

	fmt.Print("Masukkan alamat: ")
	fmt.Scan(&address)

	repo := repository.NewStudentRepo(db)
	studentService := service.NewStudentService(repo)

	// fmt.Println(reflect.TypeOf(int(userID)), "<<<<<")
	err := studentService.AddStudent(int(userID), name, phoneNumber, address)
	if err != nil {
		utils.SendJSONResponse(400, err.Error(), nil)
		return
	}

	utils.SendJSONResponse(201, "Siswa berhasil ditambahkan", nil)
}

func GetAllStudents(db *sql.DB) {
	_, valid := session()
	if !valid {
		return
	}

	repo := repository.NewStudentRepo(db)
	studentService := service.NewStudentService(repo)

	students, err := studentService.GetAllStudents()
	if err != nil {
		utils.SendJSONResponse(500, "Gagal mengambil data siswa", nil)
		return
	}

	utils.SendJSONResponse(200, "Berhasil mengambil data siswa", students)

}

func UpdateStudent(db *sql.DB) {
	userID, valid := session()
	if !valid {
		return
	}

	var ID uint16
	var name, phoneNumber, address string

	fmt.Print("Masukkan ID siswa yang akan diperbarui: ")
	fmt.Scan(&ID)

	fmt.Print("Masukkan nama siswa baru: ")
	fmt.Scan(&name)

	fmt.Print("Masukkan nomor telepon baru: ")
	fmt.Scan(&phoneNumber)

	fmt.Print("Masukkan alamat baru: ")
	fmt.Scan(&address)

	repo := repository.NewStudentRepo(db)
	studentService := service.NewStudentService(repo)

	err := studentService.UpdateStudent(ID, userID, name, phoneNumber, address)
	if err != nil {
		utils.SendJSONResponse(400, err.Error(), nil)
		return
	}

	utils.SendJSONResponse(200, "Siswa berhasil diperbarui", nil)
}

func DeleteStudent(db *sql.DB) {
	userID, valid := session()
	if !valid {
		return
	}

	var studentID int
	fmt.Print("Masukkan ID siswa yang akan dihapus: ")
	fmt.Scan(&studentID)

	repo := repository.NewStudentRepo(db)
	studentService := service.NewStudentService(repo)

	err := studentService.DeleteStudent(studentID, userID)
	if err != nil {
		utils.SendJSONResponse(400, err.Error(), nil)
		return
	}

	utils.SendJSONResponse(200, "delete success", nil)
}
