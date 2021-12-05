package product_service

import (
	"encoding/json"
	"errors"
)

type Product struct {
	ID         int64    `db:"id"`
	Name       string   `db:"name"`
	CategoryID int64    `db:"category_id"`
	Attributes ProductAttributes`db:"info"`
}

type ProductAttribute struct {
	ID         int64    `json:"id"`
	Value       string   `json:"value"`
}
type ProductAttributes []ProductAttribute

func (pa *ProductAttributes) Scan(src interface{}) (err error) {
	var ProductAttributes []ProductAttribute
	if src == nil {
		return nil
	}
	switch src.(type) {
	case string:
		err = json.Unmarshal([]byte(src.(string)), &ProductAttributes)
	case []byte:
		err = json.Unmarshal(src.([]byte), &ProductAttributes)
	default:
		return errors.New("Incompatible type")
	}

	if err != nil {
		return err
	}

	*pa = ProductAttributes
	return nil
}