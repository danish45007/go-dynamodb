package config

import "time"

type ServiceConfig interface {
	GetConfig() time.Duration
}

func GetConfig() time.Duration {

}
