package delivery

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/NuriCareers/Sreekanth-Cheriyanath-coding-challenge/blockinfo/domain"
	"github.com/gorilla/mux"
)

type BCInfoController struct {
	Svc domain.Service
}

func NewController(r *mux.Router, s domain.Service) {
	controller := &BCInfoController{
		Svc: s,
	}
	Route(r, controller)
}

// GetBlockInfo handles the fucntion for the /block endpoint
// reads the payload body and fetches the required info
func (con *BCInfoController) GetBlockInfo(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		con.HandlePOSTError(400, errors.New("empty request body"), w)
		return
	}
	req, e := ioutil.ReadAll(r.Body)
	var resp *domain.Block
	if len(req) == 0 {
		con.HandlePOSTError(400, errors.New("empty request body"), w)
		return
	}
	if e != nil {
		con.HandlePOSTError(400, e, w)
		return
	}
	var request domain.Request
	json.Unmarshal(req, &request)

	resp, e = con.Svc.GetBlockInfo(&request)
	if e != nil {
		con.HandlePOSTError(http.StatusBadRequest, e, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)
}

// GetTranInfo handles the fucntion for the /transaction endpoint
// reads the payload body and fetches the required info
func (con *BCInfoController) GetTranInfo(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		con.HandlePOSTError(400, errors.New("empty request body"), w)
		return
	}
	req, e := ioutil.ReadAll(r.Body)
	var resp *domain.Transaction
	if len(req) == 0 {
		con.HandlePOSTError(400, errors.New("empty request body"), w)
		return
	}
	if e != nil {
		con.HandlePOSTError(400, e, w)
		return
	}
	var request domain.Request
	json.Unmarshal(req, &request)

	resp, e = con.Svc.GetTranInfo(&request)
	if e != nil {
		con.HandlePOSTError(http.StatusBadRequest, e, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)
}

// HandlePOSTError Handling errors response here
func (con *BCInfoController) HandlePOSTError(status int, e error, w http.ResponseWriter) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(e.Error())
}
