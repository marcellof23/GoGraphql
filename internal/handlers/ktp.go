package handlers

import (
	"context"
	"fmt"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/marcellof23/GoGraphql/configs/database"
	"github.com/marcellof23/GoGraphql/graph/model"
	"github.com/marcellof23/GoGraphql/internal/helpers"
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

	// res pass reference of data to Create user
	res := database.DB.Create(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
	// user return reference object of new created user
}

// DeleteUserHandler delete a user by id
func DeleteUserHandler(ctx context.Context, id int64) (bool, error) {

	deletedUser, err := GetUserByIDHandler(ctx, id)

	if err != nil {
		return false, err
	}

	res := database.DB.Delete(deletedUser, "id = ?", id)

	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
	// return true if user deleted else false
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

	// res find the user with id from the input in the database
	res := database.DB.Model(&models.User{}).Where("id = ?", id).Updates(newuser)

	if res.Error != nil {
		return nil, res.Error
	}

	return &newuser, nil
	// newuser return reference object of new updated user
}

// GetUserByIDHandler retrieve user by id
func GetUserByIDHandler(ctx context.Context, id int64) (*models.User, error) {
	idUser := &models.User{}
	rs := database.DB.First(idUser, "id = ?", id)

	if rs.Error != nil {
		return nil, rs.Error
	}

	return idUser, nil
}

// GetAllUserHandler retrieve all list of user
func GetAllUserHandler(ctx context.Context) ([]*models.User, error) {
	var user []*models.User
	res := database.DB.Find(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
	// user return all user in the database
}

// GetPaginationHandler return models of PaginationResultUser with the input arguments from Pagination models
func GetPaginationHandler(ctx context.Context, input model.Pagination) (*model.PaginationResultUser, error) {
	var users, userstotal []*models.User
	var userEdges []*model.PaginationEdge
	var first, PageCount, firstRowUser int64
	var res model.PaginationResultUser

	// check if there is no After in input
	if input.After == nil {
		firstID, _, firsterr := sq.Select("id").From("user").Limit(1).ToSql()
		database.DB.Raw(firstID).Scan(&firstRowUser)
		if firsterr != nil {
			return nil, firsterr
		}
		first = firstRowUser
	} else {
		decoded, _ := helpers.DecodeCursor(*input.After)
		first = decoded.ID
	}

	// fetch all user data and
	// query by name or nik
	PageQueryUserFirstId := sq.Select("*").From("user").Where("(id >= ? AND nama LIKE ?) OR (id >= ? AND nik LIKE ?)", first, fmt.Sprint("%", input.Query, "%"), first, fmt.Sprint("%", input.Query, "%"))

	// Generate sql query and store into PageCount
	rsTotal, rtTotalargs, rtTotalerr := PageQueryUserFirstId.ToSql()
	if rtTotalerr != nil {
		return nil, rtTotalerr
	}
	rstot := database.DB.Raw(rsTotal, rtTotalargs...).Scan(&userstotal)
	if rstot.Error != nil {
		return nil, rstot.Error
	}
	PageCount = int64(len(userstotal))

	// order by sort, desceding if - in front of attribute, else ascending
	PageQueryOrder := PageQueryUserFirstId
	for _, val := range input.Sort {
		desc := strings.HasPrefix(val, "-")
		if desc {
			PageQueryOrder = PageQueryOrder.OrderBy(strings.Replace(val, "-", "", 1) + " desc")
		} else {
			PageQueryOrder = PageQueryOrder.OrderBy(val + " asc")
		}
	}

	// PageQueryLimit of the page with addition offset
	PageQueryLimit := PageQueryOrder
	if input.First != 0 {
		PageQueryLimit = PageQueryOrder.Limit(uint64(input.First + input.Offset))
	} else {
		PageQueryLimit = PageQueryOrder.Limit(10)
	}

	//Convert the sqlbuilder into string, arguments and error
	Pagequery, Pageargs, Pageerr := PageQueryLimit.ToSql()

	if Pageerr != nil {
		return nil, Pageerr
	}
	rs := database.DB.Raw(Pagequery, Pageargs...).Scan(&users)
	if rs.Error != nil {
		return nil, rs.Error
	}

	// loop over all user to insert into userEdges
	for i, user := range users {
		if i <= int(input.Offset-1) {
			continue
		}
		cur := helpers.EncodeCursor(
			helpers.CursorStruct{ID: user.ID},
		)
		userEdges = append(userEdges, &model.PaginationEdge{
			Node:   user,
			Cursor: cur,
		})
	}

	// isNextPage is a boolean checking, true if has next page, else false
	var isNextPage bool

	// error checking if userEdges empty or not
	if len(userEdges) > 0 {
		isNextPage = input.First+input.Offset < PageCount
		res = model.PaginationResultUser{
			TotalCount: PageCount,
			Edges:      userEdges,
			PageInfo: &model.PaginationInfo{
				EndCursor:   userEdges[int64(len(userEdges))-1].Cursor,
				HasNextPage: isNextPage,
			},
		}

	} else {
		isNextPage = false
	}
	return &res, nil

}
