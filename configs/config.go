package configs

import (
	"github.com/inact25/PickMyFood-BackEnd/utils/environtment"
)

type dbConf struct {
	DbUser               string
	DbPass               string
	DbHost               string
	DbPort               string
	DbSchema             string
	AllowNativePasswords bool
}

type Conf struct {
	Db dbConf
}

func NewAppConfig() *Conf {
	return &Conf{dbConf{
		DbUser:               environtment.Get("DbUser", "uckl0iyn6dk4rych"),
		DbPass:               environtment.Get("DbPass", "ZdBnoKA3gCOVPtTGW6md"),
		DbHost:               environtment.Get("DbHost", "bfdggwsfnu3oq2r32g2n-mysql.services.clever-cloud.com"),
		DbPort:               environtment.Get("DbPort", "3306"),
		DbSchema:             environtment.Get("DbSchema", "bfdggwsfnu3oq2r32g2n"),
		AllowNativePasswords: true,
	}}
}
