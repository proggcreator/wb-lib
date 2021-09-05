package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.wildberries.ru/finance/go-infrastructure/rfc7807"
	"github.com/gin-gonic/gin"
	restful "github.com/proggcreator/wb-lib"
)

const (
	id = "5"
)

func (h *Handler) employee_add(c *gin.Context) {

	//create new employee
	exampl := restful.RetEmployee()
	//
	id, err := h.services.EmplWork.CreateEmpl(exampl)
	if err != nil {
		rfc7807.NewProblem().SetTitle("BadRequest").
			SetDetail("error create employee").SetStatusCode(http.StatusBadRequest).
			Write(c.Writer)
		return
	}
	fmt.Fprintf(c.Writer, "%s Created", id)
}

func (h *Handler) employee_remove(c *gin.Context) {

	//remove employee
	err := h.services.EmplWork.DeleteEmpl(id)
	if err != nil {
		rfc7807.NewProblem().SetTitle("BadRequest").
			SetDetail("error delete employee").SetStatusCode(http.StatusBadRequest).
			Write(c.Writer)

		return
	}

}
func (h *Handler) employee_upd(c *gin.Context) {

	//update employee
	newEmp := restful.RetEmployee()
	//
	err := h.services.EmplWork.UpdateEmpl(newEmp)
	if err != nil {
		rfc7807.NewProblem().SetTitle("BadRequest").
			SetDetail("error update employee").SetStatusCode(http.StatusBadRequest).
			Write(c.Writer)
		return
	}
}
func (h *Handler) get_all(c *gin.Context) {

	list, err := h.services.EmplWork.GetAllEmpl()
	if err != nil {

		rfc7807.NewProblem().SetTitle("BadRequest").
			SetDetail("error create employee").SetStatusCode(http.StatusBadRequest).
			Write(c.Writer)
		return
	}
	fmt.Fprint(c.Writer, "Все пользователи:  ")
	fmt.Fprint(c.Writer, list)
}
func (h *Handler) employee_get(c *gin.Context) {

	//parse id param
	id := c.Param("id")
	list, err := h.services.EmplWork.GetByIdEmpl(id)
	if err != nil {

		rfc7807.NewProblem().SetTitle("BadRequest").
			SetDetail("error get by id").SetStatusCode(http.StatusBadRequest).
			Write(c.Writer)

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
