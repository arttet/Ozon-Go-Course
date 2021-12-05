package main

import (
	"bytes"
	"strconv"
)

type contract struct {
	Posts []post `json:"posts"`
}

type post struct {
	ID    id     `json:"id"`
	Title string `json:"title"`
}

type id int64

func (i *id) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		b = bytes.Trim(b, `"`)
	}

	v, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}

	*i = id(v)
	return nil
}
