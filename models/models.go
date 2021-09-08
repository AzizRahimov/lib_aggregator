package models

import (
	"encoding/xml"
	"time"
)

type Config struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Uri      string `json:"uri"`
	Port     string `json:"port"`
}


type RequestFromProc struct {
	XMLName    xml.Name          `xml:"request"`
	Command    string            `xml:"command"`
	Payment    []PaymentReq      `xml:"payment"`
	SenderInfo SenderInformation `xml:"sender_information"`
	Balance    *BalanceReq       `xml:"balance"`
	ServiceID  string            `xml:"matchingid"` // id service on FALCON
	Gateway    string            `xml:"gateway"`    // GATE->TAJPAY, OR EXPRESSPAY, BABILONM_GW_UNIQE
}



type SenderInformation struct {
	Agent struct {
		ID   int64  `xml:"id,attr"`
		Name string `xml:"name,attr"`
	} `xml:"agent"`
	Endpoint struct {
		ID      int64  `xml:"id,attr"`
		Name    string `xml:"name,attr"`
		Address string `xml:"address,attr"`
	} `xml:"endpoint"`
	Sender struct {
		Fio string `xml:"fio,attr"`
	} `xml:"sender"`
}

type PaymentReq struct {
	TerminalNum   int64      `xml:"terminalNum"`
	ClientType    *string    `xml:"clientType"`
	Trn           int64      `xml:"trn,attr"`
	Fsum          float64    `xml:"fsum,attr"`
	Tsum          float64    `xml:"tsum,attr"`
	Currency2gw   *string    `xml:"currency2gw,attr"`
	Currency      string     `xml:"currency"`
	Account       string     `xml:"account,attr"`
	Amount        float64    `xml:"amount,attr"`
	Param         string     `xml:"param"`
	NotifyFlag    *bool      `xml:"notify_flag"`
	Param2        string     `xml:"param2"`
	Time          *time.Time `xml:"-"`
	Time2         string     `xml:"time,attr"`
	Receipt       *string    `xml:"receipt"`
	ReferenceID   *string    `xml:"reference_id"`
	ExtStatus     *string    `xml:"extstatus"`
	Amount2Credit float64    `xml:"amount2credit"`
	Rate          float64    `xml:"rate"`
}

type BalanceReq struct {
	Account    string  `xml:"account,attr"`
	ClientType *string `xml:"clientType"`
}



type PaymentResponse struct {
	ReferenceID *string    `xml:"referenceid"`
	OsmpTrn     int64      `xml:"osmp_trn"`
	ReceiptID   *string    `xml:"receiptid"`
	ExtStatus   *string    `xml:"extstatus"`
	Notified    bool       `xml:"notified"`
	Date        *time.Time `xml:"date"`
}

type Response struct {
	XMLName      xml.Name          `xml:"response"`
	Status       int64             `xml:"status"`
	Message      string            `xml:"message"`
	PaymentResp  *PaymentResponse  `xml:"payment"`
	PrecheckResp *PrecheckResponse `xml:"precheck"`
	Balance      *string           `xml:"balance"`
	Overdraft    *string           `xml:"limit"`
	// Services     []ServiceResponse `xml:"services"`
}

type PrecheckResponse struct {
	Name      *string `xml:"name"`
	AccountID *string `xml:"account_id"`
	CardID    *string `xml:"card_id"`
	Pan       *string `xml:"pan"`
	Address   *string `xml:"address"`
	Previous  *string `xml:"previous"`
	Present   *string `xml:"present"`
	Date      *string `xml:"date"`
	Rest      *string `xml:"rest"`
	Item      *[]Item `xml:"item"`
}

type Item struct {
	Label string `xml:"label,attr"`
	Value string `xml:"value,attr"`
}
