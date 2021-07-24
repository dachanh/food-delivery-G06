package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type Image struct {
	Url    string `json:"url" gorm:"column:url;"`
	Width  int    `json:"width" gorm:"column:width;"`
	Height int    `json:"height" gorm:"column:height;"`
}

func (Image) TableName() string { return "images" }

func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		log.Println(ok)
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
