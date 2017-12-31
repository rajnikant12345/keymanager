package controller


type Helper interface {
	ValidateKeySize(interface{}) error
	ValidateDeletable(interface{}) error
	ValidateExportable(interface{}) error
	ValidateOwner(interface{}) error
	CreateKey(interface{}) error
}
