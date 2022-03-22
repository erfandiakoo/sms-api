package models

type ConfigModel struct {
	MelipayamakUsername string `json:"mp_username"`
	MelipayamakPasswd   string `json:"mp_passwd"`
	MelipayamakNum      string `json:"mp_number"`
	KaveNegarApi        string `json:"kavenegar_api"`
	KaveNegarSender     string `json:"kavenegar_sender"`
}
type KaveRequestModel struct {
	Message  string `json:"message"`
	Receptor string `json:"receptor"`
}

type MeliRequestModel struct {
	Message  string `json:"message"`
	Receptor string `json:"receptor"`
}
type SMSRequestModel struct {
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
