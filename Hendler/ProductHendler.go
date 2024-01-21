package Hendler

import (
	"Magazin/helper"
	"Magazin/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func PostHendler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAllProduct(w, r)

	case "POST":
		CreateProduct(w, r)
	case "PUT":
		UpdateProduct(w, r)
	case "DELETE":
		DeleteProduct(w, r)
	}

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var NewProdect models.ProdectModel
	json.NewDecoder(r.Body).Decode(&NewProdect)

	var NewData []models.ProdectModel
	ProdectByte, _ := os.ReadFile("db/Prodect.json")
	json.Unmarshal(ProdectByte, &NewData)

	NewProdect.Id = helper.MaxIdProduct(NewData)
	NewProdect.Available = true

	NewData = append(NewData, NewProdect)

	res, _ := json.Marshal(NewData)
	os.WriteFile("db/user.json", res, 0)
	fmt.Println("User Created ", NewProdect.Id)
	fmt.Fprintln(w, "User Created ")
	json.NewEncoder(w).Encode(NewProdect)

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var UpdateProdect models.ProdectModel
	json.NewDecoder(r.Body).Decode(&UpdateProdect)

	var NewData []models.ProdectModel
	ProdectByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(ProdectByte, &NewData)

	var ProdectFound bool
	for i := 0; i < len(NewData); i++ {
		if UpdateProdect.Id == NewData[i].Id {

			if UpdateProdect.ProdectType != "" {
				NewData[i].ProdectType = UpdateProdect.ProdectType

			}
			if UpdateProdect.Name != "" {
				NewData[i].Name = UpdateProdect.Name
			}
			if UpdateProdect.Quantity != 0 {
				NewData[i].Quantity = UpdateProdect.Quantity
			}
			if UpdateProdect.Price != 0 {
				NewData[i].Price = UpdateProdect.Price
			}

			NewData[i].UptadedAt = time.Now()
			NewData[i].Available = true

			ProdectFound = true
			break
		}

	}
	if !ProdectFound {
		fmt.Fprintln(w, "Prodect can not found with ID: ", UpdateProdect.Id)
		fmt.Println("Prodect can not found with ID: ", UpdateProdect.Id)
		w.WriteHeader(http.StatusNotFound)
		return

	}

	res, _ := json.Marshal(NewData)
	os.WriteFile("db/Prodect.json", res, 0)
	json.NewEncoder(w).Encode(UpdateProdect)
	fmt.Println("Prodect Created ", UpdateProdect.Id)
	fmt.Fprintln(w, "Prodect Created ")
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var DeleteProdect models.GetModel
	json.NewDecoder(r.Body).Decode(&DeleteProdect)

	var ProdectData []models.ProdectModel
	ProdectByte, _ := os.ReadFile("db/Prodect.json")
	json.Unmarshal(ProdectByte, &ProdectData)

	var ProdectFound bool

	for i := 0; i < len(ProdectData); i++ {
		if ProdectData[i].Id == DeleteProdect.ID {
			ProdectData = append(ProdectData[:i], ProdectData[i+1:]...)
			ProdectFound = true
		}

	}

	if ProdectFound {
		fmt.Println("Prodect deleted with ID: ", DeleteProdect.ID)
		fmt.Fprintln(w, "Prodect deleted with ID: ", DeleteProdect.ID)
		w.WriteHeader(http.StatusOK)

	} else {
		fmt.Println("Prodect can not found with ID ", DeleteProdect.ID)
		fmt.Fprintln(w, "Prodect can nor found with ID: ", DeleteProdect.ID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(ProdectData)
	os.WriteFile("db.Prodect.json", res, 0)

}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	var NewProdect models.GetModel
	json.NewDecoder(r.Body).Decode(&NewProdect)

	var CatagoryData []models.CatagoryModel
	CatagoryByte, _ := os.ReadFile("db/Catagory.json")
	json.Unmarshal(CatagoryByte, &CatagoryData)

	var ProdectData []models.ProdectModel
	ProdectByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(ProdectByte, &ProdectData)

	if NewProdect.ID >= 1 {

		for p := 0; p < len(ProdectData); p++ {
			if NewProdect.ID == ProdectData[p].Id {
				fmt.Fprintln(w, "--------------------------------------")
				fmt.Fprintln(w, "Product's ID", ProdectData[p].Id)
				fmt.Fprintln(w, "Product's ProdectType", ProdectData[p].ProdectType)
				fmt.Fprintln(w, "Product's Name", ProdectData[p].Name)
				fmt.Fprintln(w, "Product's Quantity", ProdectData[p].Quantity)
				fmt.Fprintln(w, "Product's Price", ProdectData[p].Price)
				fmt.Fprintln(w, "Product's available", ProdectData[p].Available)
				fmt.Fprintln(w, "Product's CreatedAt", ProdectData[p].CreatedAt)
				fmt.Fprintln(w, "Product's UptadedAt", ProdectData[p].UptadedAt)

			}

		}

	} else {
		for i := 0; i < len(ProdectData); i++ {
			fmt.Fprintln(w, "--------------------------------------")
			fmt.Fprintln(w, "Product's ID ", ProdectData[i].Id)
			fmt.Fprintln(w, "Product's ProdectType  ", ProdectData[i].ProdectType)
			fmt.Fprintln(w, "Product's Name  ", ProdectData[i].Name)
			fmt.Fprintln(w, "Product's Quantity ", ProdectData[i].Quantity)
			fmt.Fprintln(w, "Product's Price ", ProdectData[i].Price)
			fmt.Fprintln(w, "Product's available", ProdectData[i].Available)
			fmt.Fprintln(w, "Product's CreatedAt ", ProdectData[i].CreatedAt)
			fmt.Fprintln(w, "Product's UptadedAt", ProdectData[i].UptadedAt)

		}

	}
}
