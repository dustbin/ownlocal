package main

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

//holds a single row from the CSV contains meta data for JSON marshalling
type Business struct {
	ID         int    `json:"id"`
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Address2   string `json:"address2"`
	City       string `json:"city"`
	State      string `json:"state"`
	Zip        string `json:"zip"`
	Country    string `json:"country"`
	Phone      string `json:"phone"`
	Website    string `json:"website"`
	Created_at string `json:"created_at"`
}

//holds a group of rows and contains meta data for JSON marshalling
type BusinessDB struct {
	Businesses []*Business `json:"businesses"`
	Page       int         `json:"page"`
	Size       int         `json:"size"`
}

//creates a BusinessDB object from an io.Reader to a CSV, such as a File created by an os.Open call
func NewBusinessDB(reader io.Reader) (*BusinessDB, error) {
	businessDB := BusinessDB{}
	csvdb, err := NewCSVDB(reader)
	if err != nil {
		return &businessDB, err
	}
	businessDB.Page = 0
	businessDB.Size = csvdb.GetSize()
	businessDB.Businesses = make([]*Business, csvdb.GetSize())
	for i := 0; i < csvdb.GetSize(); i++ {
		business, err := NewBusiness(csvdb.GetRow(i))
		if err != nil {
			return &businessDB, err
		}
		businessDB.Businesses[i] = business
	}
	return &businessDB, nil
}

//creates a Business object from a CSV row, does basic data checking on the id
func NewBusiness(row []string) (*Business, error) {
	business := Business{}
	if len(row) != 12 {
		return &business, errors.New(fmt.Sprintf("invalid length row %d", len(row)))
	}
	i, err := strconv.Atoi(row[0])
	if err != nil {
		return &business, err
	}
	if i < 0 {
		return &business, errors.New(fmt.Sprintf("invalid id %d", i))
	}
	business.ID = i
	business.UUID = row[1]
	business.Name = row[2]
	business.Address = row[3]
	business.Address2 = row[4]
	business.City = row[5]
	business.State = row[6]
	business.Zip = row[7]
	business.Country = row[8]
	business.Phone = row[9]
	business.Website = row[10]
	business.Created_at = row[11]
	return &business, nil
}

//gets a page of arbitrary size from a BusinessDB object and returns as a BusinessDB, does basic error checking
func (bdb *BusinessDB) GetPage(page, size int) (*BusinessDB, error) {
	businessDB := BusinessDB{}
	if page < 0 {
		return &businessDB, errors.New(fmt.Sprintf("invalid page number %d", page))
	}
	if size < 0 {
		return &businessDB, errors.New(fmt.Sprintf("invalid size value %d", size))
	}
	start := page * size
	end := start + size
	if bdb.Size <= start {
		return &businessDB, errors.New(fmt.Sprintf("invalid page %d at size %d", page, size))
	}
	if bdb.Size < end {
		end = bdb.Size
	}
	businessDB.Page = page
	businessDB.Size = size
	businessDB.Businesses = bdb.Businesses[start:end]
	return &businessDB, nil
}

//gets a single Business object from a BusinessDB, does basic error checking
func (bdb *BusinessDB) GetBusiness(id int) (*Business, error) {
	if len(bdb.Businesses) <= id || id < 0 {
		return &Business{}, errors.New(fmt.Sprintf("id %d out of range", id))
	}
	return bdb.Businesses[id], nil
}
