package models

type KaveRequestModel struct {
	Message  string `json:"message"`
	Receptor string `json:"receptor"`
}

type MeliRequestModel struct {
	Message  string `json:"message"`
	Receptor string `json:"receptor"`
}

type MediaRequestModel struct {
}
type SMSIRRequestModel struct {
	Messages      string `json:"messages"`
	MobileNumbers string `json:"mobile_numbers"`
	LineNumber    string `json:"line_nuumber"`
}
