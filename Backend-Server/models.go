package main

import (
	"database/sql"

	"github.com/sanrose-bhets/Job-Grinder/internal/database"
)

type Account struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"mail"`
	Role   string `json:"role"`
	Apikey string `json:"api"`
}
type Company struct {
	Companyid         int32          `json:"id"`
	Companyname       string         `json:"name"`
	Companylocation   string         `json:"location"`
	Companywebsite    string         `json:"site"`
	Companyemail      string         `json:"mail"`
	Hiringmanagername sql.NullString `json:"hiringmanager"`
}

func databaseAcctoAcc(dbUser database.Account) Account {
	return Account{
		ID:     dbUser.ID,
		Name:   dbUser.Name,
		Email:  dbUser.Email,
		Role:   dbUser.Role,
		Apikey: dbUser.Apikey,
	}
}

func databaseComptoComp(dbUser database.Company) Company {
	return Company{
		Companyid:         dbUser.Companyid,
		Companyname:       dbUser.Companyname,
		Companylocation:   dbUser.Companylocation,
		Companywebsite:    dbUser.Companywebsite,
		Companyemail:      dbUser.Companyemail,
		Hiringmanagername: dbUser.Hiringmanagername,
	}
}
