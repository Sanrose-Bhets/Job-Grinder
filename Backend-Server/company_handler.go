package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sanrose-bhets/Job-Grinder/internal/database"
)

// a http handler that the go standard library expects
func (apiCfg *apiConfig) companyCreateHandler(w http.ResponseWriter, r *http.Request) {

	type parameter struct {
		CompanyName       string `json:"name"`
		CompanyLocation   string `json:"location"`
		CompanyWebsite    string `json:"site"`
		CompanyEmail      string `json:"mail"`
		HiringManagerName string `json:"hManager"`
	}

	decode := json.NewDecoder(r.Body)

	param := parameter{}
	err := decode.Decode(&param)
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error parsing json body: %v", err))
		return
	}

	comp, err := apiCfg.DB.CreateCompany(r.Context(), database.CreateCompanyParams{
		Companyname:     param.CompanyName,
		Companylocation: param.CompanyLocation,
		Companywebsite:  param.CompanyWebsite,
		Companyemail:    param.CompanyEmail,
		Hiringmanagername: sql.NullString{
			String: param.HiringManagerName,
			Valid:  param.HiringManagerName != "",
		},
	})
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Unable to create user: %v", err))
		return
	}

	respondwithJson(w, 200, databaseComptoComp(comp))

}
