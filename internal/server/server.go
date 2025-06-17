package server

import (
	"check_republic/internal/database"
)

type Server struct {
	DB database.DataStore
}
