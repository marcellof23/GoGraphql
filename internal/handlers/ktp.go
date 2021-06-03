package handlers

import (
	"context"
	"time"

	"github.com/marcellof23/GoGraphql/configs/database"
	"github.com/marcellof23/GoGraphql/graph/model"
	"github.com/marcellof23/GoGraphql/internal/models"
)

// CreateUserHandler create a new user to db
func CreateUserHandler(ctx context.Context, input *model.NewUser) (*models.User, error) {
	user := models.User{
		NIK:           input.Nik,
		Nama:          input.Nama,
		Jenis_kelamin: input.JenisKelamin,
		Tanggal_lahir: input.TanggalLahir,
		Alamat:        input.Alamat,
		Agama:         input.Agama,
		Created_at:    time.Now(),
	}

	res := database.DB.Create(&user) // pass reference of data to Create

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
	// newuser return reference object of new created user
}

// DeleteUserHandler delete a user by id
func DeleteUserHandler(ctx context.Context, id int64) (bool, error) {
	res := database.DB.Delete(&models.User{}, id)

	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
	// return true if user deleted
}

// UpdateUserHandler update a user by id
func UpdateUserHandler(ctx context.Context, id int64, input *model.NewUser) (*models.User, error) {

	newuser := models.User{
		NIK:           input.Nik,
		Nama:          input.Nama,
		Jenis_kelamin: input.JenisKelamin,
		Tanggal_lahir: input.TanggalLahir,
		Alamat:        input.Alamat,
		Agama:         input.Agama,
		Updated_at:    time.Now(),
	}
	res := database.DB.Model(&models.User{}).Where("id = ?", id).Updates(newuser)

	if res.Error != nil {
		return nil, res.Error
	}

	return &newuser, nil
	// newuser return reference object of new updated user
}

// GetAllUserHandler retrieve all list of user
func GetAllUserHandler(ctx context.Context) ([]*models.User, error) {
	var user []*models.User
	res := database.DB.Find(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}
