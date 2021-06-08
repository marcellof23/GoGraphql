package helpers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
)

type CursorStruct struct {
	ID int64
}

// encodeCursor encode from cursorStruct into string with base64
func EncodeCursor(curs CursorStruct) string {
	data := []byte(fmt.Sprintf("%d", curs.ID))

	// sEnc store result of encoded string data
	sEnc := base64.StdEncoding.EncodeToString((data))

	return sEnc
}

// decodeCursor decode from  string into cursorStruct  with base64
func DecodeCursor(cursor string) (*CursorStruct, error) {

	// sDec store result of decoded string cursor
	sDec, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return &CursorStruct{}, err
	}

	// vals stringify sDec
	vals := string(sDec)

	// convert vals from string into integer
	id, err := strconv.Atoi(vals)
	if err != nil {
		return &CursorStruct{}, errors.New("invalid_cursor")
	}

	// return in cursorStruct from decoded cursor
	return &CursorStruct{
		ID: int64(id),
	}, nil
}
