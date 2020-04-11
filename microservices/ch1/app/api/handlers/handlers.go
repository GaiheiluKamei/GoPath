package handlers

import (
	"encoding/json"
	"github.com/GaiheiluKamei/books/microservices/ch1/app/api/repository"
	"github.com/gorilla/mux"
	"net/http"
)

type Handlers struct {
	Repo *repository.Repository
}

func (h *Handlers) All(w http.ResponseWriter, r *http.Request) {
	tzcs, err := h.Repo.FindAll()
	if err != nil {
		error500(w, err)
		return
	}
	jr, err := json.Marshal(tzcs)
	if err != nil {
		error500(w, err)
		return
	}
	ok200(w, string(jr))
}

func (h *Handlers) GetByTZ(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tz, ok := params["timeZone"]
	if !ok {
		error400(w, "timeZone is required.")
		return
	}

	tzc, err := h.Repo.FindByTimeZone(tz)
	if err != nil {
		error404(w, "timeZone not found.")
		return
	}

	jr, err := json.Marshal(tzc)
	if err != nil {
		error500(w, err)
		return
	}
	ok200(w, string(jr))
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tz, ok := params["timeZone"]
	if !ok {
		error400(w, "timeZone not found.")
		return
	}

	err := h.Repo.Delete(repository.TZConversion{TimeZone:tz})
	if err != nil {
		error500(w, err)
		return
	}
	ok200(w, "Element successfully deleted.")
}

func (h *Handlers) Insert(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var tzc repository.TZConversion
	err := json.NewDecoder(r.Body).Decode(&tzc)
	if err != nil {
		error400(w, "Invalid json.")
		return
	}

	err = h.Repo.Insert(tzc)
	if err != nil {
		error500(w, err)
		return
	}
	ok200(w, "Element successfully inserted.")
}

func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tz, ok := params["timeZone"]
	if !ok {
		error400(w, "timeZone is required.")
		return
	}

	var tzc repository.TZConversion
	err := json.NewDecoder(r.Body).Decode(&tzc)
	if err != nil {
		error400(w, "Invalid json.")
		return
	}

	err = h.Repo.Update(tz, tzc)
	if err != nil {
		error500(w, err)
		return
	}
	ok200(w, "Element successfully updated.")
}
