package controllers

import (
	"net/http"

	db "github.com/arifikhsan/gin_rest/databases"

	"github.com/arifikhsan/gin_rest/app/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetAllPerson(c *gin.Context) {
	var persons []models.Person
	var _persons []models.TransformedPerson
	db.DB.Find(&persons)
	if len(persons) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Record not found!"})
		return
	}

	for _, item := range persons {
		_persons = append(_persons, models.TransformedPerson{ID: item.ID, Name: item.Name, Age: item.Age, UpdatedAt: item.UpdatedAt})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": _persons})
}

func CreatePerson(c *gin.Context) {
	var person models.Person
	var err error
	c.BindJSON(&person)
	if err = db.DB.Save(&person).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "messages": "Failed insert record"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"messages": "Created successfully",
		"data": models.TransformedPerson{
			ID:        person.ID,
			Name:      person.Name,
			Age:       person.Age,
			UpdatedAt: person.UpdatedAt,
		},
	})
}

func GetOnePerson(c *gin.Context) {
	var person models.Person
	personID := c.Param("id")
	db.DB.First(&person, personID)
	if person.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no person found!"})
		return
	}

	_persons := models.TransformedPerson{ID: person.ID, Name: person.Name, Age: person.Age, UpdatedAt: person.UpdatedAt}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _persons})
}

func UpdatePerson(c *gin.Context) {
	var person models.Person
	var person_input models.Person
	c.BindJSON(&person_input)

	personID := c.Param("id")
	db.DB.First(&person, personID)

	if person.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no person found!"})
		return
	}

	db.DB.Model(&person).Update(models.Person{Name: person_input.Name, Age: person_input.Age})
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Person updated successfully!",
		"data": models.TransformedPerson{
			ID:   person.ID,
			Name: person.Name,
			Age:  person.Age,
		},
	})
}

func DeletePerson(c *gin.Context) {
	var person models.Person
	personID := c.Param("id")
	db.DB.First(&person, personID)

	if person.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no person found!"})
		return
	}

	db.DB.Delete(&person)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "This person deleted successfully!"})
}
