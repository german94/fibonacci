package main

import (
	"fibonacci/fibcalc"
	"fmt"
	"math/big"
	"net/http"
)

var (
	missingInput = "Parameter n (integer >= 0) must be provided"
	wrongInput   = "There was an error manipulating the number. Please check that it is made only with 0-9 digits"
)

func calcHandler(w http.ResponseWriter, r *http.Request) {
	if !validateCalcInput(w, r) {
		return
	}
	n, ok := getNParam(r)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, wrongInput)
		return
	}
	fmt.Fprint(w, fibcalc.NewFibonacciCalculator().GetTerm(n).String())
}

func validateCalcInput(w http.ResponseWriter, r *http.Request) bool {
	_, ok := r.URL.Query()["n"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, missingInput)
		return false
	}
	return true
}

func getNParam(r *http.Request) (*big.Int, bool) {
	nQueryParam := r.URL.Query().Get("n")
	n, okConvert := big.NewInt(0).SetString(nQueryParam, 10)
	if !okConvert {
		return nil, false
	}
	return n, true
}
