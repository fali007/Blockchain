package types

type IndexPage struct{
	Name           string        `json:"name"`
	Endpoints      []Endpoint    `json:"endpoints"`
}

type Endpoint struct{
	Name           string        `json:"name"`
	Url            string        `json:"url"`
	Comment        string        `json:"comment"`
}

type TxResponse struct{
	Status 		   int  		 `json:"status"`
}

type ValidateResponse struct{
	Status 		   string		 `json:"status"`
	Error 		   string		 `json:"error"`
}