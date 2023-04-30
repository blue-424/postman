package Cnf

type Configuration struct {
	Port   string `json:"port"`
	Served string `json:"served"`
}

type Person struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Address     string `json:"address"`
	phoneNumber string `json:"phone_number"`
}
