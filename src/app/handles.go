package app

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/LATOKEN/relayer-smart-contract.git/src/common"
	"github.com/LATOKEN/relayer-smart-contract.git/src/models"
)

const numPerPage = 100

// Endpoints ...
func (a *App) Endpoints(w http.ResponseWriter, r *http.Request) {
	endpoints := struct {
		Endpoints []string `json:"endpoints"`
	}{
		Endpoints: []string{
			"/status",
			"/status/{destination_chain}/{sender}/{receipt}/{amount}/{tx_hash}",
		},
	}

	common.ResponJSON(w, http.StatusOK, endpoints)
}

// StatusHandler ...
func (a *App) SwapStatusHandler(w http.ResponseWriter, r *http.Request) {
	msg := models.SwapStatus{
		Chain:   mux.Vars(r)["destination_chain"],
		Sender:  mux.Vars(r)["sender"],
		Receipt: mux.Vars(r)["receipt"],
		Amount:  mux.Vars(r)["amount"],
		TxHash:  mux.Vars(r)["tx_hash"],
	}
	if msg.Chain == "" || msg.Sender == "" || msg.Receipt == "" || msg.Amount == "" || msg.TxHash == "" {
		a.logger.Errorf("Empty request(/status/{destination_chain}/{sender}/{receipt}/{amount}/{tx_hash})")
		common.ResponJSON(w, http.StatusInternalServerError, createNewError("empty request", ""))
		return
	}

	status, err := a.relayer.GetSwapStatus(&msg)
	if err != nil {
		common.ResponJSON(w, http.StatusInternalServerError, createNewError("get swap from database", err.Error()))
		return
	}

	common.ResponJSON(w, http.StatusOK, status)
}

// StatusHandler ...
func (a *App) StatusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := a.relayer.StatusOfWorkers()
	if err != nil {
		common.ResponError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.ResponJSON(w, http.StatusOK, status)
}
