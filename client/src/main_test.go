package main

import (
	"form3-client/config"
	"form3-client/model"
	"form3-client/operation"
	"form3-client/util"
	"log"
	"testing"
)

const givenAccountNumber = "400300"
const givenAccountId = "eb0bd6f5-c3f5-44b2-b677-acd23cdde516"

var givenVersion = util.CreateNumberPointer(0)

func TestInit(t *testing.T) {
	const baseUrl = "http://interview-accountapi:8080"
	// const baseUrl = "http://localhost:8080"
	config.SetBaseUrl(baseUrl)
}

func TestCreate_expect201Created(t *testing.T) {

	var accountData model.AccountData
	accountData.ID = givenAccountId
	accountData.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde616"
	accountData.Type = "accounts"
	accountData.Version = givenVersion

	var accountAttr model.AccountAttributes
	accountAttr.AccountClassification = util.CreateStringPointer("Personal")
	accountAttr.AccountMatchingOptOut = util.CreateBooleanPointer(false)
	accountAttr.AccountNumber = givenAccountNumber
	accountAttr.AlternativeNames = []string{"Jana", "Rathan"}
	accountAttr.BankID = "400300"
	accountAttr.BankIDCode = "GBDSC"
	accountAttr.BaseCurrency = "SGD"
	accountAttr.Bic = "NWBKGB22"
	accountAttr.Country = util.CreateStringPointer("SG")
	accountAttr.Name = []string{"Jana", "Param"}

	accountData.Attributes = &accountAttr

	resp := operation.Create(accountData)
	// log.Println(resp)

	const createdStatus = "201 Created"
	var actualStatus = resp.Status

	if actualStatus != createdStatus {
		t.Errorf("Operation status actual:%q expected:%q", actualStatus, createdStatus)
	}
}

func TestCreate_expect409Conflict(t *testing.T) {

	var accountData model.AccountData
	accountData.ID = givenAccountId
	accountData.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde616"
	accountData.Type = "accounts"
	accountData.Version = givenVersion

	var accountAttr model.AccountAttributes
	accountAttr.AccountClassification = util.CreateStringPointer("Personal")
	accountAttr.AccountMatchingOptOut = util.CreateBooleanPointer(false)
	accountAttr.AccountNumber = givenAccountNumber
	accountAttr.AlternativeNames = []string{"Jana", "Rathan"}
	accountAttr.BankID = "400300"
	accountAttr.BankIDCode = "GBDSC"
	accountAttr.BaseCurrency = "SGD"
	accountAttr.Bic = "NWBKGB22"
	accountAttr.Country = util.CreateStringPointer("SG")
	accountAttr.Name = []string{"Jana", "Param"}

	accountData.Attributes = &accountAttr

	resp := operation.Create(accountData)
	log.Println(resp)

	const conflictStatus = "409 Conflict"
	var actualStatus = resp.Status

	if actualStatus != conflictStatus {
		t.Errorf("Operation status actual:%q expected:%q", actualStatus, conflictStatus)
	}
}

func TestCreate_expect400BadRequest_missingField(t *testing.T) {

	var accountData model.AccountData
	accountData.ID = givenAccountId
	// accountData.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde616"
	accountData.Type = "accounts"
	accountData.Version = givenVersion

	var accountAttr model.AccountAttributes
	accountAttr.AccountClassification = util.CreateStringPointer("Personal")
	accountAttr.AccountMatchingOptOut = util.CreateBooleanPointer(false)
	accountAttr.AccountNumber = givenAccountNumber
	accountAttr.AlternativeNames = []string{"Jana", "Rathan"}
	accountAttr.BankID = "400300"
	accountAttr.BankIDCode = "GBDSC"
	accountAttr.BaseCurrency = "SGD"
	accountAttr.Bic = "NWBKGB22"
	accountAttr.Country = util.CreateStringPointer("SG")
	accountAttr.Name = []string{"Jana", "Param"}

	accountData.Attributes = &accountAttr

	resp := operation.Create(accountData)
	log.Println(resp)

	const badRequestStatus = "400 Bad Request"
	var actualStatus = resp.Status

	if actualStatus != badRequestStatus {
		t.Errorf("Operation status actual:%q expected:%q", actualStatus, badRequestStatus)
	}
}

func TestCreate_expect400BadRequest_InvalidInput(t *testing.T) {

	var accountData model.AccountData
	accountData.ID = givenAccountId
	accountData.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde616"
	accountData.Type = "accounts"
	accountData.Version = givenVersion

	var accountAttr model.AccountAttributes
	accountAttr.AccountClassification = util.CreateStringPointer("Personal")
	accountAttr.AccountMatchingOptOut = util.CreateBooleanPointer(false)
	accountAttr.AccountNumber = givenAccountNumber
	accountAttr.AlternativeNames = []string{"Jana", "Rathan"}
	accountAttr.BankID = "400300"
	accountAttr.BankIDCode = "GBDSC PQRS" // <--- incorrect format
	accountAttr.BaseCurrency = "SGD"
	accountAttr.Bic = "NWBKGB22"
	accountAttr.Country = util.CreateStringPointer("SG")
	accountAttr.Name = []string{"Jana", "Param"}

	accountData.Attributes = &accountAttr

	resp := operation.Create(accountData)
	log.Println(resp)

	const badRequestStatus = "400 Bad Request"
	var actualStatus = resp.Status

	if actualStatus != badRequestStatus {
		t.Errorf("Operation status actual:%q expected:%q", actualStatus, badRequestStatus)
	}
}

func TestFetch(t *testing.T) {
	apiResponse := operation.Fetch()
	accountDataList := apiResponse.AccountDataList
	if len(accountDataList) != 1 {
		t.Error("Expected one account")
	}
}

func TestFetchMapped(t *testing.T) {
	accounts := operation.FetchMapped()
	// log.Println("AccountsAfterCreation:", accounts)

	if len(accounts) != 1 {
		t.Error("Expected one account")
	}

	account := accounts[0]
	var fetchedAccountNumber = account.AccountNumber

	if fetchedAccountNumber != givenAccountNumber {
		t.Errorf("AccountNumber expected:%q actual:%q", givenAccountNumber, fetchedAccountNumber)
	}
}

func TestDelete(t *testing.T) {

	operation.Delete(givenAccountId, *givenVersion)

	accounts := operation.FetchMapped()
	log.Println("AccountsAfterDeletion:", accounts)

	actualCount := len(accounts)
	const expectedCount = 0

	if actualCount != expectedCount {
		t.Errorf("No of accounts actual:%q expected:%q", actualCount, expectedCount)
	}

}