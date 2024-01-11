package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/darlio88/go-api-test/api"
	"github.com/darlio88/go-api-test/internal/tools"
	"github.com/gorilla/schema"

	log "github.com/sirupsen/logrus"
)

func GetCoinsBalance(w http.ResponseWriter, r *http.Request) {

	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()

	var err error
	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	fmt.Println(tokenDetails)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*&tokenDetails.Coins),
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
