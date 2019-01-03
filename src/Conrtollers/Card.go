package conrtollers

import (
	"fmt"
	"net/http"
	"time"

	gov "github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"

	models "Go/rest-api/src/Models"
)

// GetCard - get all cards from collection
func (h *Conrtollers) GetCard(c echo.Context) error {
	cardz := []models.Card{}
	h.DB.DB("CardsDB").C("Cards").Find(bson.M{}).All(&cardz)

	return c.JSON(http.StatusOK, cardz)
}

// CreateCard - create a new card from payload
func (h *Conrtollers) CreateCard(c echo.Context) error {
	card := &models.Card{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Active:    true,
	}

	c.Bind(card)

	_, err := gov.ValidateStruct(card)
	if err != nil {
		fmt.Println("Validate error: ", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	println("card: ", card.Type, card.Code)

	err = h.DB.DB("CardsDB").C("Cards").Insert(card)
	if err != nil {
		fmt.Println("ERR: ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusCreated, "OK")
}

// DeleteCard - make a card as deleted
func (h *Conrtollers) DeleteCard(c echo.Context) error {
	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		return c.String(http.StatusBadRequest, "You passed incorrect id")
	}

	fmt.Println("id: ", id)

	err := h.DB.DB("CardsDB").C("Cards").UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"active": false}})
	if err != nil {
		fmt.Println("ERR: ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusOK, "Ok")
}

// UpdateCard - update a card data from payload
func (h *Conrtollers) UpdateCard(c echo.Context) error {
	id := c.Param("id")
	card := &models.Card{}
	c.Bind(card)

	_, err := gov.ValidateStruct(card)
	if err != nil {
		fmt.Println("Validate error: ", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println("card: ", card)

	if !bson.IsObjectIdHex(id) {
		return c.String(http.StatusBadRequest, "You passed incorrect id")
	}

	err = h.DB.DB("CardsDB").C("Cards").UpdateId(bson.ObjectIdHex(id), card)

	if err != nil {
		fmt.Println("ERR: ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusCreated, "OK")
}
