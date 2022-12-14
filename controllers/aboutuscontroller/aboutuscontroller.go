package aboutuscontroller

import (
	"fmt"
	"net/http"

	"github.com/R-Media-Solutions/rmediasolutions-website-backend/helper"
	"github.com/R-Media-Solutions/rmediasolutions-website-backend/models"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var aboutus []models.AboutUs

	if err := models.DB.Find(&aboutus).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		fmt.Println(err)
		return
	}

	ResponseJson(w, http.StatusOK, aboutus)
}

func Show(w http.ResponseWriter, r *http.Request) {

}

func Create(w http.ResponseWriter, r *http.Request) {

}

func Update(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
