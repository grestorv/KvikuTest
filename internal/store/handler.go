package store

import (
	"Server/internal/handlers"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/getAll", h.GetList)
	router.POST("/save/:id", h.SaveValue)
	router.GET("/get/:id", h.GetValue)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(200)
	_, err := w.Write([]byte(fmt.Sprintf("Get list of values")))
	if err != nil {
		panic(err)
	}
}

func (h *handler) SaveValue(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(201)
	_, err := w.Write([]byte(fmt.Sprintf("Save value")))
	if err != nil {
		panic(err)
	}
}
func (h *handler) GetValue(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(200)
	_, err := w.Write([]byte(fmt.Sprintf("Get value")))
	if err != nil {
		panic(err)
	}
}
