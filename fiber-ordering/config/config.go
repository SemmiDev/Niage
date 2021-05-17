package config

var Server 	*server
var MySQL 	*mySQL

func init() {
	MySQL = setupMySQL()
	Server = setupServer()
}