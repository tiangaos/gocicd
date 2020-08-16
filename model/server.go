package model

type Environment int32

const (
	DEV Environment = iota
	FAT
	UAT
	TEST
	PROD
)

type Server struct {
	Id   string
	Name string
	IP   string
	Env  Environment
}
