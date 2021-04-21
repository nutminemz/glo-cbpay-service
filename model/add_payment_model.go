package model

import "encoding/xml"

type AddPaymentRequest struct {
}

type AddPaymentResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    ResultBodyPayment `json:"data"`
}

type ResponsePayment struct {
	XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SoapBody *SOAPBodyPaymentResponse
}

type SOAPBodyPaymentResponse struct {
	XMLName xml.Name `xml:"Body"`
	Resp    *ResponsePaymentBody
}

type ResponsePaymentBody struct {
	XMLName  xml.Name `xml:"BillPaymentAddReturn"`
	Response *BodyPayment
}

type BodyPayment struct {
	XMLName xml.Name `xml:"return"`
	Code    string   `xml:"code"`
	Info    string   `xml:"info"`
	Result  *ResultBodyPayment
}

type ResultBodyPayment struct {
	Name     string   `xml:"print1" json:"name"`
	Cid      string   `xml:"print2" json:"cid"`
	Info     string   `xml:"print3" json:"info"`
	PostDesc string   `xml:"print4" json:"post_desc"`
	XMLName  xml.Name `xml:"result" json:"-"`
}

///// inquiry coses
type ResponseInquiry struct {
	XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SoapBody *SOAPBodyInquiryResponse
}

type SOAPBodyInquiryResponse struct {
	XMLName xml.Name `xml:"Body"`
	Resp    *ResponseInquiryBody
}

type ResponseInquiryBody struct {
	XMLName  xml.Name `xml:"inquiryCosesResponse"`
	Response *BodyInquiry
}

type BodyInquiry struct {
	XMLName xml.Name `xml:"return"`
	Code    string   `xml:"code"`
	Info    string   `xml:"info"`
	Result  *ResultBodyInquiry
}

type ResultBodyInquiry struct {
	XMLName xml.Name `xml:"cosesPaymentResult"`
	Info    string   `xml:"print1"`
}
