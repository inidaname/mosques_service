package handler

import (
	"net/http"

	"github.com/inidaname/mosque/mosques-service/internal/service"
	"github.com/inidaname/mosque/mosques-service/internal/util"
	"github.com/inidaname/mosque_location/protos"
)

type MosqueHttpHandler struct {
	mosqueService service.MosqueService
}

func NewHttpMosqueService(mosqueService service.MosqueService) *MosqueHttpHandler {
	handler := &MosqueHttpHandler{
		mosqueService: mosqueService,
	}

	return handler
}

func (h *MosqueHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /mosque", h.CreateMosque)
	router.HandleFunc("GET /mosque", h.ListMosque)
	router.HandleFunc("PUT /mosque", h.UpdateMosque)
}

func (h *MosqueHttpHandler) CreateMosque(w http.ResponseWriter, r *http.Request) {
	var req protos.CreateMosqueRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	resp, err := h.mosqueService.CreateMosque(r.Context(), &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	util.WriteJSON(w, http.StatusCreated, resp)
}

func (h *MosqueHttpHandler) ListMosque(w http.ResponseWriter, r *http.Request) {
	var req protos.ListMosquesRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.mosqueService.ListMosque(r.Context(), &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, resp)

}

func (h *MosqueHttpHandler) UpdateMosque(w http.ResponseWriter, r *http.Request) {
	var req protos.UpdateMosqueRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.mosqueService.UpdateMpsque(r.Context(), &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, resp)

}
