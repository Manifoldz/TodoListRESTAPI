package handler

import (
	"net/http"

	"github.com/Manifoldz/TodoListRESTAPI/internal/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	var input entities.ToDoList

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.ToDoList.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) getAllList(c *gin.Context)  {}
func (h *Handler) getListById(c *gin.Context) {}
func (h *Handler) updateList(c *gin.Context)  {}
func (h *Handler) deleteList(c *gin.Context)  {}
