package handler

import (
	"net/http"
	"strconv"

	"github.com/Manifoldz/TodoListRESTAPI/internal/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id parameter")
		return
	}

	var input entities.ToDoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.ToDoItem.Create(listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
func (h *Handler) getAllItem(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id parameter")
		return
	}
	items, err := h.services.ToDoItem.GetAll(listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}
func (h *Handler) getItemById(c *gin.Context) {}
func (h *Handler) updateItem(c *gin.Context)  {}
func (h *Handler) deleteItem(c *gin.Context)  {}
