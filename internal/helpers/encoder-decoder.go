package helpers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
)

type cursorStruct struct {
	ID int64
}

// encodeCursor encode from cursorStruct into string with base64
func encodeCursor(curs cursorStruct) string {
	data := []byte(fmt.Sprintf("%d", curs.ID))

	// sEnc store result of encoded string data
	sEnc := base64.StdEncoding.EncodeToString((data))

	return sEnc
}

// decodeCursor decode from  string into cursorStruct  with base64
func decodeCursor(cursor string) (*cursorStruct, error) {

	// sDec store result of decoded string cursor
	sDec, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return &cursorStruct{}, err
	}

	// vals stringify sDec
	vals := string(sDec)

	// convert vals from string into integer
	id, err := strconv.Atoi(vals)
	if err != nil {
		return &cursorStruct{}, errors.New("invalid_cursor")
	}

	// return in cursorStruct from decoded cursor
	return &cursorStruct{
		ID: int64(id),
	}, nil
}
