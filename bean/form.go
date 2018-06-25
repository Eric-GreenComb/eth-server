package bean

import ()

// FormParams FormParams
type FormParams struct {
	Params   string `form:"params" json:"params"`     // params
	Key      string `form:"key" json:"key"`           // key
	Value    string `form:"value" json:"value"`       // value
	TTL      string `form:"ttl" json:"ttl"`           // ttl
	Address  string `form:"address" json:"address"`   //
	From     string `form:"from" json:"from"`         //
	To       string `form:"to" json:"to"`             //
	Amount   string `form:"amount" json:"amount"`     //
	Decimals string `form:"decimals" json:"decimals"` //
	Pwd      string `form:"pwd" json:"pwd"`           //
}
