package conf

import (
	"encoding/json"
	"io"
	"os"

	"gopkg.in/go-playground/validator.v9"
)

type Conf struct {
	Workers int `json:"workers" validate:"required,gt=0"`
}

func Validate(cfg Conf) error {
	v := validator.New()
	if err := v.Struct(cfg); err != nil {
		return err
	}
	return nil
}

func Load(fName string) (*Conf, error) {
	f, err := os.Open(fName)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var cfg Conf
	if err = json.Unmarshal(body, &cfg); err != nil {
		return nil, err
	}

	if err = Validate(cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
