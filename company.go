package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/rs/xid"
)

type Company struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Sector    string `json:"sector"`
	Category  string `json:"category"`
	IsStartup bool   `json:"is_startup"`
	CEO       string `json:"ceo"`
	Revenue   string `json:"revenue"`
}

func readCompaniesFromFile() []Company {
	companyJson, err := ioutil.ReadFile("./companies.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	var companies []Company

	if err := json.Unmarshal(companyJson, &companies); err != nil {
		log.Fatal(err.Error())
	}
	return companies
}

func getCompanies() ([]Company, error) {
	companies := readCompaniesFromFile()

	var newCompanies []Company

	for _, company := range companies {
		newCompanies = append(newCompanies, Company{
			ID:        company.ID,
			Name:      company.Name,
			Sector:    company.Sector,
			Category:  company.Category,
			IsStartup: company.IsStartup,
			CEO:       company.CEO,
			Revenue:   company.Revenue,
		})
	}
	return newCompanies, nil
}

func addNew(name, sector, category, ceo, revenue string, is_startup bool) Company {

	var newCompany Company

	fmt.Println("Fields", name)
	newCompany.ID = xid.New().String()
	newCompany.Name = name
	newCompany.Sector = sector
	newCompany.Category = category
	newCompany.IsStartup = is_startup
	newCompany.CEO = ceo
	newCompany.Revenue = revenue

	companies := readCompaniesFromFile()

	companies = append(companies, newCompany)

	companyByte, err := json.Marshal(companies)
	if err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
	}

	err = ioutil.WriteFile("companies.json", companyByte, 0644)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	return newCompany
}
