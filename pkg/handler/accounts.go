package handler

import (
	"echo_ifElse"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) registration(c echo.Context) error {
	var jsonBody echo_ifElse.AccountRequest
	var jsonResponse echo_ifElse.AccountResponse
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	id, err := h.services.Registration(jsonBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	jsonResponse.Email = jsonBody.Email
	jsonResponse.FirstName = jsonBody.FirstName
	jsonResponse.LastName = jsonBody.LastName
	jsonResponse.Id = id
	jsonResponse.Role = "USER"

	return c.JSON(http.StatusOK, jsonResponse)
}

func (h *Handler) getAcc(c echo.Context) error {
	var jsonResponse echo_ifElse.AccountResponse
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	jsonResponse, err = h.services.GetAcc(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, jsonResponse)
}

func (h *Handler) addAcc(c echo.Context) error {
	var jsonBody echo_ifElse.AccountRequest
	var jsonResponse echo_ifElse.AccountResponse
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	id, err := h.services.AddAcc(jsonBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	jsonResponse.Email = jsonBody.Email
	jsonResponse.FirstName = jsonBody.FirstName
	jsonResponse.LastName = jsonBody.LastName
	jsonResponse.Id = id
	jsonResponse.Role = jsonBody.Role

	return c.JSON(http.StatusOK, jsonResponse)
}

func (h *Handler) updateAcc(c echo.Context) error {
	var jsonRequest echo_ifElse.UpdateAccountResponse
	var jsonResponse echo_ifElse.AccountResponse
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	err = json.NewDecoder(c.Request().Body).Decode(&jsonRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	jsonResponse, err = h.services.UpdateAcc(id, jsonRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, jsonResponse)
}

func (h *Handler) deleteAcc(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	err = h.services.DeleteAcc(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, "ok")
}
