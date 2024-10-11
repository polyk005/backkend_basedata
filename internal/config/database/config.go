package database

//Config - КОнфигурация для подключения к БД
type Config struct {
	User     string
	Password string
	Host     string
	Port     int
	DbName   string
}
