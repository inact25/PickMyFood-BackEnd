package main

import (
	"github.com/inact25/PickMyFood-BackEnd/configs"
	"github.com/inact25/PickMyFood-BackEnd/masters/apis"
	"github.com/inact25/PickMyFood-BackEnd/utils"
)

func main() {
	conf := configs.NewAppConfig()
	db, err := configs.InitDB(conf)
	utils.ErrorCheck(err, "Print")
	myRoute := configs.CreateRouter()
	apis.Init(myRoute, db)
	configs.RunServer(myRoute)
}
