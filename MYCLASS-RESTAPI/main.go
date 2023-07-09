package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type data struct {
	ID int `json:id`
	Name string `json:"name"`
	NPM string `json:"NPM"`
}

func setupRouter() *gin.Engine{
	r := gin.Default();

	v1 := r.Group("v1")
	v1.GET("/2IA07", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
		"data": []data{
			data{
				ID: 1,
				Name: "ABYAN MUNTAQO ILYAS",
				NPM: "50421018",
			},
			data{
				ID: 2,
				Name: "ADRIAN NOVA",
				NPM: "50421055",
			},
			data{
				ID: 3,
				Name: "AHMAD MARSHAL HAKIM",
				NPM: "50421073",
			},
			data{
				ID: 4,
				Name: "AHMAD RAIHAN WAHYULLAH",
				NPM: "50421078",
			},
			data{
				ID: 5,
				Name: "ALGAVRA ZHAPANDYA BINANDITA",
				NPM: "50421117",
			},
			data{
				ID: 6,
				Name: "ARINDA ABRAAR",
				NPM: "50421207",
			},
			data{
				ID: 7,
				Name: "ARYA DUTA FAUSTINE",
				NPM: "50421219",
			},
			data{
				ID: 8,
				Name: "BENEDICTUS ADOLF JOSEVAN SARUM",
				NPM: "50421272",
			},
			data{
				ID: 9,
				Name: "DAVID MATTHEW LAMSIHAR",
				NPM: "50421339",
			},
			data{
				ID: 10,
				Name: "DONI MANGAMPA",
				NPM: "50421405",
			},
			data{
				ID: 11,
				Name: "DYFAN AL-FATH DJIDAN ALDANI",
				NPM: "50421412",
			},
			data{
				ID: 12,
				Name: "FARAH NAWAL AZALIA",
				NPM: "50421468",
			},
			data{
				ID: 13,
				Name: "FARHAN LUTFI",
				NPM: "50421475",
			},
			data{
				ID: 14,
				Name: "FITRIA AMANDA PRATIWI",
				NPM: "50421530",
			},
			data{
				ID: 15,
				Name: "FITRIYANI",
				NPM: "50421533",
			},
			data{
				ID: 16,
				Name: "FLAMELIA SAKHAWATININGRUM",
				NPM: "50421536",
			},
			data{
				ID: 17,
				Name: "HENDRA KUSUMA",
				NPM: "50421601",
			},
			data{
				ID: 18,
				Name: "JONATHAN RIYADI LIMBONG",
				NPM: "51421606",
			},
			data{
				ID: 19,
				Name: "KEVIN FEBRIAN BIMANTARA",
				NPM: "50421723",
			},
			data{
				ID: 20,
				Name: "M RAYDZAN D M",
				NPM: "51421609",
			},
			data{
				ID: 21,
				Name: "MATTHEW MICHAEL GIDEON",
				NPM: "50421790",
			},
			data{
				ID: 22,
				Name: "	MIRANTI PRAMUJI",
				NPM: "50421817",
			},
			data{
				ID: 23,
				Name: "MUHAMAD FAJRI",
				NPM: "50421862",
			},
			data{
				ID: 24,
				Name: "MUHAMAD ZIDAN",
				NPM: "50421887",
			},
			data{
				ID: 25,
				Name: "MUHAMMAD AZRIEL APRIJAL",
				NPM: "50421928",
			},
			data{
				ID: 26,
				Name: "MUHAMMAD FAUZI ILHAMSYAH",
				NPM: "50421964",
			},
			data{
				ID: 27,
				Name: "MUHAMMAD HILDAN FIRDAWAN",
				NPM: "50421988",
			},
			data{
				ID: 28,
				Name: "MUHAMMAD SAHAL NURDIN",
				NPM: "51421075",
			},
			data{
				ID: 29,
				Name: "NI NYOMAN AYU NIRMALA LUKITA",
				NPM: "51421136",
			},
			data{
				ID: 30,
				Name: "NUR ISMA ALIANTI",
				NPM: "51421136",
			},
			data{
				ID: 31,
				Name: "RAHUL JAIKISHIN GIANI",
				NPM: "51421224",
			},
			data{
				ID: 32,
				Name: "RANGGA SAPUTRA",
				NPM: "51421256",
			},
			data{
				ID: 33,
				Name: "RAYHAN FAIZ",
				NPM: "51421265",
			},
			data{
				ID: 34,
				Name: "RIZKI EKA FACHRUDIN",
				NPM: "51421340",
			},
			data{
				ID: 35,
				Name: "RIZKY PUTRA PERDANA",
				NPM: "51421348",
			},
			data{
				ID: 36,
				Name: "RUTHNISSI ASMORO",
				NPM: "51421368",
			},
			data{
				ID: 37,
				Name: "RYVERA BR. PASARIBU",
				NPM: "51421373",
			},
			data{
				ID: 38,
				Name: "SHAQILA SAFA RINALDI",
				NPM: "51421415",
			},
			data{
				ID: 39,
				Name: "TIRTA MUKTI WICAKSONO",
				NPM: "51421480",
			},
			data{
				ID: 40,
				Name: "VITO BAHARI",
				NPM: "51421508",
			},


			},
		})
	})

	

return r
}


func main()  {
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}