package Services

import (
	"rental_easy.in/m/pkg/database"
	rental "rental_easy.in/m/pkg/rentalmgmt"
)

type ServerSideImplementation struct {
	rental.UnimplementedRental_Easy_FunctionalitiesServer
	Db database.DataBase
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
