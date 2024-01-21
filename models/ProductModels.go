package models

import (
	"time"
)

type ProdectModel  struct {

	Id int
	CatagoryId int
	ProdectType string
	Name string
	Quantity int
	Available bool
	Price int
	CreatedAt time.Time
	UptadedAt time.Time

}