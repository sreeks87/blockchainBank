package delivery

import "github.com/gorilla/mux"

// Basic route handling will be setup here
func Route(r *mux.Router, controller *BCInfoController) {
	r.HandleFunc("/block", controller.GetBlockInfo).Methods("POST")
	r.HandleFunc("/transaction", controller.GetTranInfo).Methods("POST")
}
