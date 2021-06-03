package handlers

import (
	"context"
	"time"

	"github.com/marcellof23/GoGraphql/configs/database"
	"github.com/marcellof23/GoGraphql/graph/model"
	"github.com/marcellof23/GoGraphql/internal/models"
)

func stringToDate(value string) time.Time {
	var layoutFormat = "2006-01-02 15:04:05"
	var date, _ = time.Parse(layoutFormat, value)

	return date
}

func CreateUserHandler(ctx context.Context, input *model.NewUser) (*models.User, error) {
	user := models.User{
		NIK:          input.Nik,
		Nama:         input.Nama,
		JenisKelamin: input.JenisKelamin,
		TanggalLahir: input.TanggalLahir,
		Alamat:       input.Alamat,
		Agama:        input.Agama,
		CreatedAt:    time.Now(),
	}

	res := database.DB.Create(&user) // pass pointer of data to Create

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func DeleteUserHandler(ctx context.Context, id int64) (bool, error) {
	res := database.DB.Delete(&models.User{}, id)

	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}

func UpdateUserHandler(ctx context.Context, id int64, input *model.NewUser) (*models.User, error) {
	// Assumption where update all fields get updated
	newuser := models.User{
		NIK:          input.Nik,
		Nama:         input.Nama,
		JenisKelamin: input.JenisKelamin,
		TanggalLahir: input.TanggalLahir,
		Alamat:       input.Alamat,
		Agama:        input.Agama,
		UpdatedAt:    time.Now(),
	}
	res := database.DB.Model(&models.User{}).Where("id = ?", id).Updates(newuser)

	if res.Error != nil {
		return nil, res.Error
	}

	return &newuser, nil
}

func GetAllUserHandler(ctx context.Context) ([]*models.User, error) {
	var user []*models.User
	res := database.DB.Find(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}
