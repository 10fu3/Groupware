package Entity

import "fmt"

type DbSettings struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func (s *DbSettings) GetDbStrings() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", s.User, s.Password, s.Host, s.Port, s.DbName)
}
