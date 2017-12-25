package configuration



var ConfMap map[string]string

func InitializeConfiguration() {
	ConfMap["keydatabase"] = "Keymanager"
	ConfMap["user"] = "root"
	ConfMap["password"] = "root"
	ConfMap["host"] = "0.0.0.0"
	ConfMap["port"] = "3456"

}
