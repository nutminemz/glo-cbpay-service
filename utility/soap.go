package utility

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"

	"api.inno/glo-profile-service/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type requestPayment struct {
	CusAcct string
	Dtm     string
	Tel     string
	TxRef   string
	Cid     string
	Ref1    string
}

type requestInquiry struct {
	CusAcct  string
	Dtm      string
	TxRef    string
	TranDate string
	TranTime string
}

func CallSOAPClientSteps(acc string, txRef string, cid string, tel string) (string, string, model.ResultBodyPayment, error) {
	moDes := model.ResultBodyPayment{}
	reqInq := populateRequestInquiry(acc, txRef, cid)
	httpReqInq, err := generateSOAPRequestInquiry(reqInq)
	if err != nil {
		fmt.Println("Some problem occurred in inquiry cosses request generation")
		return "", "", moDes, err
	}
	responseInq, err := soapCallInquiry(httpReqInq)
	if err != nil {
		fmt.Println("Problem occurred in making a SOAP inquiry cosses call")
		return "", "", moDes, err
	}
	code := responseInq.SoapBody.Resp.Response.Code
	log.Println(code)
	if code == "00000" {
		log.Println(responseInq.SoapBody.Resp.Response.Result.Info)
		ref1 := responseInq.SoapBody.Resp.Response.Result.Info

		reqPay := populateRequestPayment(acc, txRef, cid, tel, ref1)
		httpReq, err := generateSOAPRequestPayment(reqPay)
		if err != nil {
			fmt.Println("Some problem occurred in billaddpayment request generation")
			return "", "", moDes, err
		}

		response, err := soapCallPayment(httpReq)
		if err != nil {
			fmt.Println("Problem occurred in making a SOAP billaddpayment call")
			return "", "", moDes, err
		}
		moDes.Cid = response.SoapBody.Resp.Response.Result.Cid
		moDes.Info = response.SoapBody.Resp.Response.Result.Info
		moDes.Name = response.SoapBody.Resp.Response.Result.Name
		moDes.PostDesc = response.SoapBody.Resp.Response.Result.PostDesc
		log.Print(response.SoapBody.Resp.Response.Result.Info)
		paymentResCode := response.SoapBody.Resp.Response.Code

		return paymentResCode, response.SoapBody.Resp.Response.Info, moDes, nil
	}
	inqErrorTxt := responseInq.SoapBody.Resp.Response.Info
	return code, inqErrorTxt, moDes, err
}

func populateRequestPayment(acc string, txRef string, cid string, tel string, ref1 string) *requestPayment {

	currentTime := time.Now()
	req := requestPayment{}
	req.CusAcct = acc
	req.Dtm = currentTime.Format("02/01/2006 15:04:05")
	req.Tel = tel
	req.TxRef = txRef
	req.Cid = cid
	req.Ref1 = ref1
	return &req
}

func populateRequestInquiry(acc string, txRef string, cid string) *requestInquiry {

	currentTime := time.Now()
	req := requestInquiry{}
	req.CusAcct = acc
	req.Dtm = currentTime.Format("02/01/2006 15:04:05")
	req.TranDate = currentTime.Format("02/01/2006")
	req.TranTime = currentTime.Format("15:04:05")
	req.TxRef = txRef
	return &req
}

func generateSOAPRequestPayment(req *requestPayment) (*http.Request, error) {
	// Using the var getTemplate to construct request
	template, err := template.New("InputRequestPayment").Parse(getTemplatePaymentAddPayment)
	if err != nil {
		fmt.Println("Error while marshling object. %s ", err.Error())
		return nil, err
	}

	doc := &bytes.Buffer{}
	// Replacing the doc from template with actual req values
	err = template.Execute(doc, req)
	if err != nil {
		fmt.Println("template.Execute error. %s ", err.Error())
		return nil, err
	}

	buffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(buffer)
	err = encoder.Encode(doc.String())
	if err != nil {
		fmt.Println("encoder.Encode error. %s ", err.Error())
		return nil, err
	}

	r, err := http.NewRequest(http.MethodPost, "http://kcsjbsap90.kcs:49101/billpresentment-ws/BillPresentment", bytes.NewBuffer([]byte(doc.String())))
	if err != nil {
		fmt.Println("Error making a request. %s ", err.Error())
		return nil, err
	}

	return r, nil
}

func generateSOAPRequestInquiry(req *requestInquiry) (*http.Request, error) {
	// Using the var getTemplate to construct request
	template, err := template.New("InputRequestInquiry").Parse(getTemplatePaymentInquiry)
	if err != nil {
		fmt.Println("Error while marshling object. %s ", err.Error())
		return nil, err
	}

	doc := &bytes.Buffer{}
	// Replacing the doc from template with actual req values
	err = template.Execute(doc, req)
	if err != nil {
		fmt.Println("template.Execute error. %s ", err.Error())
		return nil, err
	}

	buffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(buffer)
	err = encoder.Encode(doc.String())
	if err != nil {
		fmt.Println("encoder.Encode error. %s ", err.Error())
		return nil, err
	}

	r, err := http.NewRequest(http.MethodPost, "http://kcsjbsap90.kcs:49101/payment-ws/PaymentWSService", bytes.NewBuffer([]byte(doc.String())))
	if err != nil {
		fmt.Println("Error making a request. %s ", err.Error())
		return nil, err
	}

	return r, nil
}

func soapCallPayment(req *http.Request) (*model.ResponsePayment, error) {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := &model.ResponsePayment{}
	err = xml.Unmarshal(body, &r)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func soapCallInquiry(req *http.Request) (*model.ResponseInquiry, error) {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := &model.ResponseInquiry{}
	err = xml.Unmarshal(body, &r)

	if err != nil {
		return nil, err
	}

	return r, nil
}

var getTemplatePaymentAddPayment = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ws="http://ws.regbill.cbpay.kcs.com/">
  <soapenv:Header/>
  <soapenv:Body>
    <ws:BillPaymentAdd>
      <!--Optional:-->
      <BillPaymentAddRequest>
        <bankCD>6</bankCD>
        <brCD>0</brCD>
        <chID>8</chID>
        <code>0200.500000</code>
        <currency>THB</currency>
        <dtm>{{.Dtm}}</dtm>
        <language>th</language>
        <password/>
        <termID>mKTB</termID>
        <txRef>mKTB.{{.TxRef}}</txRef>
        <userID>0</userID>
        <param>
          <dtc>0</dtc>
          <payeeCode>91865</payeeCode>
          <payment>
            <comAmt>35200.0</comAmt>
            <comFeeAmt>10</comFeeAmt>
            <comFeeAmtParam>
              <total>10</total>
            </comFeeAmtParam>
            <commitFlg>Y</commitFlg>
            <cusAcct>{{.CusAcct}}</cusAcct>
            <cusAmt>35200.0</cusAmt>
            <effFlg>Y</effFlg>
            <pmtDtm>{{.Dtm}}</pmtDtm>
            <pmtRef>.{{.TxRef}}</pmtRef>
          </payment>
          <ref1>{{.Ref1}}</ref1>
          <ref2>5</ref2>
          <ref3>{{.Cid}}</ref3>
          <revFlg>N</revFlg>
          <tCmt/>
        </param>
      </BillPaymentAddRequest>
    </ws:BillPaymentAdd>
  </soapenv:Body>
</soapenv:Envelope>`

var getTemplatePaymentInquiry = `
<x:Envelope xmlns:x="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ws="http://ws.payment.cbpay.kcs.com/">
    <x:Header/>
    <x:Body>
        <ws:inquiryCoses>
            <request>
                <bankCD>6</bankCD>
                <brCD>0</brCD>
                <chID>8</chID>
                <code>0200.500000</code>
                <dtm>{{.Dtm}}</dtm>
                <termID>mKTB</termID>
                <txRef>mKTB.{{.TxRef}}</txRef>
                <userID>1200534</userID>
                <param>
                    <amt1>0</amt1>
                    <amt2>0</amt2>
                    <amt3>0</amt3>
                    <amt4>0</amt4>
                    <channel>8</channel>
                    <expiryDate>{{.TranDate}}</expiryDate>
                    <fromAcctId></fromAcctId>
                    <fromBankCd></fromBankCd>
                    <msgLength>0</msgLength>
                    <ref1>{{.CusAcct}}</ref1>
                    <rqUID>0000</rqUID>
                    <servProviderId>91865</servProviderId>
                    <systemId>0000</systemId>
                    <tranCode>00000</tranCode>
                    <tranDate>{{.TranDate}}</tranDate>
                    <tranTime>{{.TranTime}}</tranTime>
                </param>
            </request>
        </ws:inquiryCoses>
    </x:Body>
</x:Envelope>`
