package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	restful "github.com/proggcreator/wb-Restful"
)

const (
	id      = "5"
	timeout = 100 * time.Millisecond
)

func (h *Handler) employee_add(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	//create new employee
	exampleEmp := restful.RetEmployee()
	//
	id, err := h.services.EmplWork.CreateEmpl(exampleEmp, ctx)
	if err != nil {
		NewJsonError(c, JsonError{
			Status: http.StatusBadRequest,
			Title:  "BadRequest",
			Detail: "error create employee"})
		return
	}
	fmt.Fprintf(c.Writer, "%s Created", id)
}

func (h *Handler) employee_remove(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	//remove employee
	err := h.services.EmplWork.DeleteEmpl(id, ctx)
	if err != nil {
		NewJsonError(c, JsonError{
			Status: http.StatusBadRequest,
			Title:  "BadRequest",
			Detail: "error delete employee"})
		return
	}

}
func (h *Handler) employee_upd(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	//update employee
	newEmp := restful.RetEmployee()
	//
	err := h.services.EmplWork.UpdateEmpl(newEmp, ctx)
	if err != nil {
		NewJsonError(c, JsonError{
			Status: http.StatusBadRequest,
			Title:  "BadRequest",
			Detail: "error update employee"})
		return
	}
}
func (h *Handler) get_all(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	list, err := h.services.EmplWork.GetAllEmpl(ctx)
	if err != nil {
		NewJsonError(c, JsonError{
			Status: http.StatusBadRequest,
			Title:  "BadRequest",
			Detail: "error create employee"})
		return
	}
	fmt.Fprint(c.Writer, "Все пользователи:  ")
	fmt.Fprint(c.Writer, list)
}
func (h *Handler) employee_get(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	//parse id param
	id := c.Param("id")
	list, err := h.services.EmplWork.GetByIdEmpl(id, ctx)
	if err != nil {
		NewJsonError(c, JsonError{
			Status: http.StatusBadRequest,
			Title:  "BadRequest",
			Detail: "error get by id"})
		return
	}
	fmt.Fprint(c.Writer, list)

}

type infoApp struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (h *Handler) employee_tech(c *gin.Context) {

	mystr := infoApp{
		Name:    "employees",
		Version: "1.0.0",
	}

	out, _ := json.MarshalIndent(mystr, "", "  ")

	fmt.Fprintln(c.Writer, string(out))
}
