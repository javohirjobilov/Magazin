package helper

import (
	"Magazin/models"
	
)
 

func MaxIdCatagory(PostArray []models.CatagoryModel) int {
	var maxID =  0
	 for i := 0; i <len(PostArray); i++ {
		if maxID < PostArray[i].ID {
			maxID = PostArray[i].ID
		}
	 }
	 return maxID+1

}
	 
func MaxIdProduct(PostArray []models.ProdectModel) int {
	var maxID =  0
	 for i := 0; i <len(PostArray); i++ {
		if maxID < PostArray[i].Id {
			maxID = PostArray[i].Id
		}
	 }
	 return maxID+1

}