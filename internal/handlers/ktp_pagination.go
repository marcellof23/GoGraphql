package handlers

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/marcellof23/GoGraphql/configs/database"
	"github.com/marcellof23/GoGraphql/graph/model"
	"github.com/marcellof23/GoGraphql/internal/models"
)

type cursorStruct struct {
	ID int64
}

// encodeCursor encode from cursorStruct into string with base64
func encodeCursor(curs cursorStruct) string {
	data := []byte(fmt.Sprintf("%d", curs.ID))

	sEnc := base64.StdEncoding.EncodeToString((data))

	return sEnc
}

// decodeCursor decode from  string into cursorStruct  with base64
func decodeCursor(cursor string) (*cursorStruct, error) {
	sDec, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return &cursorStruct{}, err
	}
	vals := string(sDec)
	id, err := strconv.Atoi(vals)
	if err != nil {
		return &cursorStruct{}, errors.New("invalid_cursor")
	}

	return &cursorStruct{
		ID: int64(id),
	}, nil
}

// GetPaginationHandler return models of PaginationResultUser with the input arguments from Pagination models
func GetPaginationHandler(ctx context.Context, input model.Pagination) (*model.PaginationResultUser, error) {
	var users []*models.User
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
		decoded, _ := decodeCursor(*input.After)
		first = decoded.ID
	}

	// retrieve all total of id from table user
	totalId, _, _ := sq.Select("COUNT(id)").From("user").ToSql()
	rsTotal := database.DB.Raw(totalId).Scan(&PageCount)
	if rsTotal.Error != nil {
		return nil, rsTotal.Error
	}

	// fetch all user data
	fetchUserData := sq.Select("*").From("user")

	// query by name or nik
	PageQueryUserFirstId := fetchUserData.Where("(id >= ? AND nama LIKE ?) OR (id >= ? AND nik LIKE ?)", first, fmt.Sprint("%", input.Query, "%"), first, fmt.Sprint("%", input.Query, "%"))

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
	PageQueryLimit := PageQueryOrder.Limit(uint64(input.First + input.Offset))

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
		cur := encodeCursor(
			cursorStruct{ID: user.ID},
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
		isNextPage = userEdges[int64(len(userEdges))-1].Node.ID < PageCount
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
