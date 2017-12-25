package configuration

import "os"

var ConfMap map[string]string

func InitializeConfiguration() {

	ConfMap = make(map[string]string)

	ConfMap["DBNAME"] = os.Getenv("DBNAME")
	ConfMap["DBUSR"] = os.Getenv("DBUSR")
	ConfMap["DBPASSWORD"] = os.Getenv("DBPASSWORD")
	ConfMap["DBHOST"] = os.Getenv("DBHOST")
	ConfMap["DBPORT"] = os.Getenv("DBPORT")

}
