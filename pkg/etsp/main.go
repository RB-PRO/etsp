package etsp

const URL string = "https://ws.etsp.ru"

type User struct {
	Login       string `json:"Login"`
	Password    string `json:"Password"`
	HashSession string `json:"-"`
}
