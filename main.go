package main

import (
	"github.com/Hypes-Astro/task-5-vix-btpns-Muhamad-Alif-Anwar/database"
	"github.com/Hypes-Astro/task-5-vix-btpns-Muhamad-Alif-Anwar/models"
	"github.com/Hypes-Astro/task-5-vix-btpns-Muhamad-Alif-Anwar/router"
)

func main() {
	db := database.SetupDB()
	db.AutoMigrate(&models.User{})

	r := router.SetupRoutes(db)
	r.Run()
}
