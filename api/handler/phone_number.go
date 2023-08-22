package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/api/models"
	"github.com/user/pkg/helper"
)

// @Security ApiKeyAuth
// Create phone_number godoc
// @ID create_phone_number
// @Router /v1/phone_number [POST]
// @Summary Create PhoneNumber
// @Description Create PhoneNumber
// @Tags PhoneNumber
// @Accept json
// @Procedure json
// @Param PhoneNumber body models.CreatePhoneNumber true "CreatePhoneNumberRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreatePhoneNumber(c *gin.Context) {

	var createPhoneNumber models.CreatePhoneNumber
	err := c.ShouldBindJSON(&createPhoneNumber)
	if err != nil {
		h.handlerResponse(c, "error PhoneNumber should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.PhoneNumber().Create(c.Request.Context(), &createPhoneNumber)
	if err != nil {
		h.handlerResponse(c, "storage.PhoneNumber.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.PhoneNumber().GetByID(c.Request.Context(), &models.PhoneNumberPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.PhoneNumber.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create PhoneNumber resposne", http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// GetByID phone_number godoc
// @ID get_by_id_phone_number
// @Router /v1/phone_number/{id} [GET]
// @Summary Get By ID PhoneNumber
// @Description Get By ID PhoneNumber
// @Tags PhoneNumber
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdPhoneNumber(c *gin.Context) {
	var id string
	id = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.PhoneNumber().GetByID(c.Request.Context(), &models.PhoneNumberPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.PhoneNumber.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id PhoneNumber resposne", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// GetList phone_number godoc
// @ID get_list_phone_number
// @Router /v1/phone_number [GET]
// @Summary Get List PhoneNumber
// @Description Get List PhoneNumber
// @Tags PhoneNumber
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "name"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListPhoneNumber(c *gin.Context) {
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list PhoneNumber offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list PhoneNumber limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.PhoneNumber().GetList(c.Request.Context(), &models.PhoneNumberGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.PhoneNumber.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list PhoneNumber resposne", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Update phone_number godoc
// @ID update_phone_number
// @Router /v1/phone_number/{id} [PUT]
// @Summary Update PhoneNumber
// @Description Update PhoneNumber
// @Tags PhoneNumber
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param PhoneNumber body models.UpdatePhoneNumber true "UpdatePhoneNumberRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdatePhoneNumber(c *gin.Context) {

	var (
		id                string = c.Param("id")
		updatePhoneNumber models.UpdatePhoneNumber
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updatePhoneNumber)
	if err != nil {
		h.handlerResponse(c, "error PhoneNumber should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updatePhoneNumber.Id = id
	rowsAffected, err := h.strg.PhoneNumber().Update(c.Request.Context(), &updatePhoneNumber)
	if err != nil {
		h.handlerResponse(c, "storage.PhoneNumber.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.PhoneNumber.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.PhoneNumber().GetByID(c.Request.Context(), &models.PhoneNumberPrimaryKey{Id: updatePhoneNumber.Id})
	if err != nil {
		h.handlerResponse(c, "storage.PhoneNumber.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create PhoneNumber resposne", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete phone_number godoc
// @ID delete_phone_number
// @Router /v1/phone_number/{id} [DELETE]
// @Summary Delete PhoneNumber
// @Description Delete PhoneNumber
// @Tags PhoneNumber
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeletePhoneNumber(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.PhoneNumber().Delete(c.Request.Context(), &models.PhoneNumberPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.PhoneNumber.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create PhoneNumber resposne", http.StatusNoContent, nil)
}
