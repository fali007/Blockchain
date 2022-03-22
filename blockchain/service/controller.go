package service

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"blockchain/types"
)

func IndexController(w http.ResponseWriter, r *http.Request){
	document,err:=ioutil.ReadFile("info.json")
	if err!=nil{
		fmt.Println("Error opening info file", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(document)
}

func GenesisController(w http.ResponseWriter, r *http.Request){
	document,err:=ioutil.ReadFile("genesis.json")
	if err!=nil{
		fmt.Println("Error opening info file", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(document)
}

func TransactionController(w http.ResponseWriter, r *http.Request){
	from:=r.URL.Query()["from"][0]
	to:=r.URL.Query()["to"][0]
	value:=r.URL.Query()["value"][0]
	data:=r.URL.Query()["data"][0]

	status:=Transaction(from,to,data,value)
	if status==true{
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(GetJsonEncoding(types.TxResponse{http.StatusOK}))
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(GetJsonEncoding(types.TxResponse{http.StatusBadRequest}))
	}
}

func CloseTransaction(w http.ResponseWriter, r *http.Request){
	status:=CloseChannel()
	if status==true{
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(GetJsonEncoding(types.TxResponse{http.StatusOK}))
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(GetJsonEncoding(types.TxResponse{http.StatusBadRequest}))
	}
}

func GetBalances(w http.ResponseWriter, r *http.Request){
	state:=LoadState()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(GetJsonEncoding(state.Balances))
}

func Verify(w http.ResponseWriter, r *http.Request){
	resp:=ValidateStateWithResponse()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}