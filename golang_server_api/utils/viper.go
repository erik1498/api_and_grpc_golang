package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	_viper "github.com/spf13/viper"
)

type Viper interface {
	Config
	myBoost() Config
}

type viper struct {
	viper *_viper.Viper
	*sync.Mutex
}

func NewViper() (Config, error) {
	v := _viper.New()

	v.SetConfigName(".env")
	v.SetConfigType("json")
	v.AddConfigPath(".")
	fmt.Println("trying read .env")
	var errNotFound *_viper.ConfigFileAlreadyExistsError
	if err := v.ReadInConfig(); err != nil {
		if ok := errors.As(err, &errNotFound); !ok {
			return nil, errNotFound
		}
		fmt.Println(err)
		return nil, errNotFound
	}

	return &viper{
		v,
		&sync.Mutex{},
	}, nil
}

func GetGCPSettings() (string, error) {
	v := _viper.New()
	v.SetConfigName(".env")
	v.SetConfigType("json")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		return "", errors.New("error reading .env file")
	}

	settings := v.GetStringMapString(EnvGcpRootKey)

	jsonString, err := json.Marshal(settings)
	if err != nil {
		return "", errors.New("error marshalling settings to JSON")
	}

	return string(jsonString), nil
}

func (v *viper) GetInt(key string) int64 {
	return v.viper.GetInt64(EnvRootKey + key)
}

func (v *viper) GetString(key string) string {
	return v.viper.GetString(EnvRootKey + key)
}

func (v *viper) GetFloat64(key string) float64 {
	return v.viper.GetFloat64(EnvRootKey + key)
}

func (v *viper) GetBool(key string) bool {
	return v.viper.GetBool(EnvRootKey + key)
}
