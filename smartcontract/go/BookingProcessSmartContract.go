package main
//this main calls the function and package it 

import (
	"encoding/json"
	"fmt"
	"log"

	// "strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//not required
const ChannelName = "mychannel"

// "receiptTypeAtOrigin": "CY",
// "deliveryTypeAtDestination": "CY",
// "cargoMovementTypeAtOrigin": "FCL",
// "cargoMovementTypeAtDestination": "FCL",
// "serviceContractReference": "string",
// "vesselName": "King of the Seas",
// "carrierServiceName": "Great Lion Service",
// "carrierServiceCode": "FE1",
// "universalServiceReference": "SR12345A",
// "carrierExportVoyageNumber": "2103S",
// "universalExportVoyageReference": "2103N",
// "declaredValue": 1231.1,
// "declaredValueCurrency": "DKK",
// "paymentTermCode": "PRE",
// "isPartialLoadAllowed": true,
// "isExportDeclarationRequired": true,
// "exportDeclarationReference": "ABC123123",
// "isImportLicenseRequired": true,
// "importLicenseReference": "ABC123123",
// "isCustomsFilingSubmissionByShipper": true,
// "contractQuotationReference": "DKK",
// "expectedDepartureDate": "2021-05-17",
// "expectedArrivalAtPlaceOfDeliveryStartDate": "2021-05-17",
// "expectedArrivalAtPlaceOfDeliveryEndDate": "2021-05-19",
// "transportDocumentTypeCode": "SWB",
// "transportDocumentReference": "string",
// "bookingChannelReference": "ABC12313",
// "incoTerms": "FCA",
// "communicationChannelCode": "AO",
// "isEquipmentSubstitutionAllowed": true,
// "vesselIMONumber": "9321483",
// "preCarriageModeOfTransportCode": "VESSEL",
// "invoicePayableAt": {
//   "locationName": "Eiffel Tower",
//   "UNLocationCode": "FRPAR"
// },
// "placeOfBLIssue": {
//   "locationName": "DCSA Headquarters",
//   "UNLocationCode": "NLAMS"
// },
// "commodities": [
//   {
// 	"commodityType": "Mobile phones",
// 	"HSCode": "string",
// 	"cargoGrossWeight": 12000,
// 	"cargoGrossWeightUnit": "KGM",
// 	"cargoGrossVolume": 120,
// 	"cargoGrossVolumeUnit": "MTQ",
// 	"numberOfPackages": 18,
// 	"exportLicenseIssueDate": "2021-05-14",
// 	"exportLicenseExpiryDate": "2021-05-21",
// 	"commodityRequestedEquipmentLink": "001"
//   }
// ],
// "valueAddedServices": [
//   {
// 	"valueAddedServiceCode": "SCON"
//   }
// ],
// "references": [
//   {
// 	"type": "FF",
// 	"value": "string"
//   }
// ],
// "requestedEquipments": [
//   {
// 	"ISOEquipmentCode": "22RT",
// 	"tareWeight": 4800,
// 	"tareWeightUnit": "KGM",
// 	"units": 3,
// 	"equipmentReferences": [
// 	  "PSSU2109481"
// 	],
// 	"isShipperOwned": true,
// 	"commodityRequestedEquipmentLink": "001"
//   }
// ],
// "documentParties": [
//   {
// 	"party": {
// 	  "partyName": "Asseco Denmark",
// 	  "taxReference1": "CVR-25645774",
// 	  "taxReference2": "CVR-25645774",
// 	  "publicKey": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkFzaW",
// 	  "address": {
// 		"name": "Henrik",
// 		"street": "Kronprinsessegade",
// 		"streetNumber": "54",
// 		"floor": "5. sal",
// 		"postCode": "1306",
// 		"city": "København",
// 		"stateRegion": "N/A",
// 		"country": "Denmark"
// 	  },
// 	  "partyContactDetails": [
// 		{
// 		  "name": "Henrik",
// 		  "phone": "+45 33364660",
// 		  "email": "info@dcsa.org",
// 		  "url": "https://www.dcsa.org"
// 		}
// 	  ],
// 	  "identifyingCodes": [
// 		{
// 		  "DCSAResponsibleAgencyCode": "SMDG",
// 		  "partyCode": "MSK",
// 		  "codeListName": "LCL"
// 		}
// 	  ]
// 	},
// 	"partyFunction": "DDS",
// 	"displayedAddress": [
// 	  "Kronprincessegade 54"
// 	],
// 	"isToBeNotified": true
//   }
// ],
// "shipmentLocations": [
//   {
// 	"location": {
// 	  "locationName": "CMP Container Terminal Copenhagen",
// 	  "UNLocationCode": "DKCPH",
// 	  "facilityCode": "CMPDK",
// 	  "facilityCodeListProvider": "SMDG"
// 	},
// 	"shipmentLocationTypeCode": "PRE",
// 	"eventDateTime": "2021-11-03T10:23:00+01:00"
//   }
// ]

// BookingContract represents the chaincode implementation

type BookingSmartContract struct {
	contractapi.Contract
}

// Booking represents the structure of a booking
//add more json obj as per the requirement 
// type BookingRequestAsset struct {
// 	BookingID       string `json:"bookingId"`
// 	BillingAddress  string `json:"billingAddress"`
// 	ShippingAddress string `json:"shippingAddress"`
// 	TypeOfGoods     string `json:"typesOfGoods"`
// 	Quantity        string `json:"quantity"`
// 	ServiceType     string `json:"serviceType"`
// 	Remarks         string `json:"remarks"` 
// 	receiptTypeAtOrigin string `json:"receiptTypeAtOrigin"`
// 	cargoMovementTypeAtOrigin string `json:"cargoMovementTypeAtOrigin"`
// 	deliveryTypeAtDestination string `json:"deliveryTypeAtDestination"`
// 	serviceContractReference string `json:"serviceContractReference"`
// 	vesselName  string `json:"vesselName "`
// 	carrierServiceName string `json:"carrierServiceName"`
// 	carrierServiceCode string `json:"carrierServiceCode"`
// 	universalServiceReference string `json:"universalServiceReference"`
// 	carrierExportVoyageNumber string `json:"carrierExportVoyageNumber"`
// 	universalExportVoyageReference string `json:"universalExportVoyageReference"`
// 	declaredValueCurrency string `json:"declaredValueCurrency"`
// 	declaredValue string `json:"declaredValue"`
// 	paymentTermCode string `json:"paymentTermCode"`
// 	isPartialLoadAllowed bool `json:"isPartialLoadAllowed"`
// 	isExportDeclarationRequired bool `json:"isExportDeclarationRequired"`
// 	exportDeclarationReference string `json:"exportDeclarationReference"`
// 	isImportLicenseRequired bool `json:"isImportLicenseRequired"`
// 	importLicenseReference string `json:"importLicenseReference"`
// 	isCustomsFilingSubmissionByShipper bool `json:"isCustomsFilingSubmissionByShipper"`
// 	contractQuotationReference string `json:"contractQuotationReference"`
// 	expectedDepartureDate string `json:"expectedDepartureDate"`
// 	expectedArrivalAtPlaceOfDeliveryStartDate string `json:"expectedArrivalAtPlaceOfDeliveryStartDate"`
// 	expectedArrivalAtPlaceOfDeliveryEndDate string `json:"expectedArrivalAtPlaceOfDeliveryEndDate"`
// 	transportDocumentTypeCode string `json:"transportDocumentTypeCode"`
// 	transportDocumentReference string `json:"transportDocumentReference"`
// 	bookingChannelReference string `json:"bookingChannelReference"`
// 	incoTerms string `json:"incoTerms"`
// 	communicationChannelCode string `json:"communicationChannelCode"`
// 	isEquipmentSubstitutionAllowed bool `json:"isEquipmentSubstitutionAllowed"`
// 	vesselIMONumber string `json:"vesselIMONumber"`
// 	preCarriageModeOfTransportCode string `json:"preCarriageModeOfTransportCode"`
// 	// cargoMovementTypeAtDestination string `json:"cargoMovementTypeAtDestination"`
// 	// invoicePayableAt string `json:"invoicePayableAt"`
// 	// placeOfBLIssue string `json:"placeOfBLIssue"`
// 	// commodities  string `json:"commodities"`
// 	CommodityType               string `json:"commodityType"`
// 	HSCode                      string `json:"HSCode"`
// 	CargoGrossWeight            int    `json:"cargoGrossWeight"`
// 	CargoGrossWeightUnit        string `json:"cargoGrossWeightUnit"`
// 	CargoGrossVolume            int    `json:"cargoGrossVolume"`
// 	CargoGrossVolumeUnit        string `json:"cargoGrossVolumeUnit"`
// 	NumberOfPackages            int    `json:"numberOfPackages"`
// 	ExportLicenseIssueDate      string `json:"exportLicenseIssueDate"`
// 	ExportLicenseExpiryDate     string `json:"exportLicenseExpiryDate"`
// 	CommodityRequestedEquipment string `json:"commodityRequestedEquipmentLink"`


// }
type Commodity struct {
	CommodityType               string `json:"commodityType"`
	HSCode                      string `json:"HSCode"`
	CargoGrossWeight            int    `json:"cargoGrossWeight"`
	CargoGrossWeightUnit        string `json:"cargoGrossWeightUnit"`
	CargoGrossVolume            int    `json:"cargoGrossVolume"`
	CargoGrossVolumeUnit        string `json:"cargoGrossVolumeUnit"`
	NumberOfPackages            int    `json:"numberOfPackages"`
	ExportLicenseIssueDate      string `json:"exportLicenseIssueDate"`
	ExportLicenseExpiryDate     string `json:"exportLicenseExpiryDate"`
	CommodityRequestedEquipment string `json:"commodityRequestedEquipmentLink"`
}

type Location struct {
	LocationName     string `json:"locationName"`
	UNLocationCode   string `json:"UNLocationCode"`
}

type PartyContactDetails struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	URL   string `json:"url"`
}

type Address struct {
	Name         string               `json:"name"`
	Street       string               `json:"street"`
	StreetNumber string               `json:"streetNumber"`
	Floor        string               `json:"floor"`
	PostCode     string               `json:"postCode"`
	City         string               `json:"city"`
	StateRegion  string               `json:"stateRegion"`
	Country      string               `json:"country"`
}

type IdentifyingCodes struct {
	DCSAResponsibleAgencyCode string `json:"DCSAResponsibleAgencyCode"`
	PartyCode                 string `json:"partyCode"`
	CodeListName              string `json:"codeListName"`
}

type Party struct {
	PartyName             string             `json:"partyName"`
	TaxReference1         string             `json:"taxReference1"`
	TaxReference2         string             `json:"taxReference2"`
	PublicKey             string             `json:"publicKey"`
	Address               Address            `json:"address"`
	PartyContactDetails   []PartyContactDetails `json:"partyContactDetails"`
	IdentifyingCodes      []IdentifyingCodes `json:"identifyingCodes"`
}

type DocumentParty struct {
	Party           Party  `json:"party"`
	PartyFunction   string `json:"partyFunction"`
	DisplayedAddress []string `json:"displayedAddress"`
	IsToBeNotified  bool   `json:"isToBeNotified"`
}

type ShipmentLocation struct {
	Location               Location `json:"location"`
	ShipmentLocationTypeCode string   `json:"shipmentLocationTypeCode"`
	EventDateTime         string   `json:"eventDateTime"`
}

type ValueAddedService struct {
	ValueAddedServiceCode string `json:"valueAddedServiceCode"`
}

type Reference struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type RequestedEquipment struct {
	ISOEquipmentCode              string   `json:"ISOEquipmentCode"`
	TareWeight                    int      `json:"tareWeight"`
	TareWeightUnit                string   `json:"tareWeightUnit"`
	Units                         int      `json:"units"`
	EquipmentReferences           []string `json:"equipmentReferences"`
	IsShipperOwned                bool     `json:"isShipperOwned"`
	CommodityRequestedEquipment   string   `json:"commodityRequestedEquipmentLink"`
}

type BookingRequestAsset struct {
	ReceiptTypeAtOrigin                      string             `json:"receiptTypeAtOrigin"`
	DeliveryTypeAtDestination                string             `json:"deliveryTypeAtDestination"`
	CargoMovementTypeAtOrigin                string             `json:"cargoMovementTypeAtOrigin"`
	CargoMovementTypeAtDestination           string             `json:"cargoMovementTypeAtDestination"`
	ServiceContractReference                 string             `json:"serviceContractReference"`
	VesselName                               string             `json:"vesselName"`
	CarrierServiceName                       string             `json:"carrierServiceName"`
	CarrierServiceCode                       string             `json:"carrierServiceCode"`
	UniversalServiceReference                string             `json:"universalServiceReference"`
	CarrierExportVoyageNumber                string             `json:"carrierExportVoyageNumber"`
	UniversalExportVoyageReference           string             `json:"universalExportVoyageReference"`
	DeclaredValue                            float64            `json:"declaredValue"`
	DeclaredValueCurrency                    string             `json:"declaredValueCurrency"`
	PaymentTermCode                          string             `json:"paymentTermCode"`
	IsPartialLoadAllowed                     bool               `json:"isPartialLoadAllowed"`
	IsExportDeclarationRequired              bool               `json:"isExportDeclarationRequired"`
	ExportDeclarationReference               string             `json:"exportDeclarationReference"`
	IsImportLicenseRequired                  bool               `json:"isImportLicenseRequired"`
	ImportLicenseReference                   string             `json:"importLicenseReference"`
	IsCustomsFilingSubmissionByShipper       bool               `json:"isCustomsFilingSubmissionByShipper"`
	ContractQuotationReference               string             `json:"contractQuotationReference"`
	ExpectedDepartureDate                    string             `json:"expectedDepartureDate"`
	ExpectedArrivalAtPlaceOfDeliveryStartDate string             `json:"expectedArrivalAtPlaceOfDeliveryStartDate"`
	ExpectedArrivalAtPlaceOfDeliveryEndDate   string             `json:"expectedArrivalAtPlaceOfDeliveryEndDate"`
	TransportDocumentTypeCode                string             `json:"transportDocumentTypeCode"`
	TransportDocumentReference               string             `json:"transportDocumentReference"`
	BookingChannelReference                  string             `json:"bookingChannelReference"`
	IncoTerms                                string             `json:"incoTerms"`
	CommunicationChannelCode                 string             `json:"communicationChannelCode"`
	IsEquipmentSubstitutionAllowed           bool               `json:"isEquipmentSubstitutionAllowed"`
	VesselIMONumber                          string             `json:"vesselIMONumber"`
	PreCarriageModeOfTransportCode            string             `json:"preCarriageModeOfTransportCode"`
	InvoicePayableAt                         Location           `json:"invoicePayableAt"`
	PlaceOfBLIssue                           Location           `json:"placeOfBLIssue"`
	Commodities                              []Commodity        `json:"commodities"`
	ValueAddedServices                       []ValueAddedService `json:"valueAddedServices"`
	References                               []Reference        `json:"references"`
	RequestedEquipments                      []RequestedEquipment `json:"requestedEquipments"`
	DocumentParties                          []DocumentParty    `json:"documentParties"`
	ShipmentLocations                        []ShipmentLocation `json:"shipmentLocations"`
}




//checks whether the booking id exists or not 
func (t *BookingSmartContract) BookingRequestAssetExists(ctx contractapi.TransactionContextInterface, bookingid string) (bool, error) {
	assetBytes, err := ctx.GetStub().GetState(bookingid)
	if err != nil {
		log.Println("Failed to read booking request asset from world state for bookingid" + bookingid)
		return false, fmt.Errorf("failed to read asset %s from world state. %v", bookingid, err)
	}

	return assetBytes != nil, nil
}

//event


// CreateAsset initializes a new asset in the ledger
func (t *BookingSmartContract) CreateBookingRequest(ctx contractapi.TransactionContextInterface, bookingid string) error {

		//calls the function to check the booking id exits in the ledger or not 
	exists, err := t.BookingRequestAssetExists(ctx, bookingid)
	if err != nil {
		log.Println("CreateBookingRequest:Failed to get booking request asset with bookingid" + bookingid)
		return fmt.Errorf("failed to get booking request: %v", err)
	}
	if exists {
		log.Println(" test k CreateBookingRequest:BookingRequest asset already exists with bookingid" + bookingid)
		return fmt.Errorf("BookingRequest already exists: %s", bookingid)
	}
	
	// obj build which is pushed to the ledger based on the structure 
	// asset := BookingRequestAsset{

		

	// }
	asset := BookingRequestAsset{
		ReceiptTypeAtOrigin:                      "CY",
		DeliveryTypeAtDestination:                "CY",
		CargoMovementTypeAtOrigin:                "FCL",
		CargoMovementTypeAtDestination:           "FCL",
		ServiceContractReference:                 "string",
		VesselName:                               "King of the Seas",
		CarrierServiceName:                       "Great Lion Service",
		CarrierServiceCode:                       "FE1",
		UniversalServiceReference:                "SR12345A",
		CarrierExportVoyageNumber:                "2103S",
		UniversalExportVoyageReference:           "2103N",
		DeclaredValue:                            1231.1,
		DeclaredValueCurrency:                    "DKK",
		PaymentTermCode:                          "PRE",
		IsPartialLoadAllowed:                     true,
		IsExportDeclarationRequired:              true,
		ExportDeclarationReference:               "ABC123123",
		IsImportLicenseRequired:                  true,
		ImportLicenseReference:                   "ABC123123",
		IsCustomsFilingSubmissionByShipper:       true,
		ContractQuotationReference:               "DKK",
		ExpectedDepartureDate:                    "2021-05-17",
		ExpectedArrivalAtPlaceOfDeliveryStartDate: "2021-05-17",
		ExpectedArrivalAtPlaceOfDeliveryEndDate:   "2021-05-19",
		TransportDocumentTypeCode:                "SWB",
		TransportDocumentReference:               "string",
		BookingChannelReference:                  "ABC12313",
		IncoTerms:                                "FCA",
		CommunicationChannelCode:                 "AO",
		IsEquipmentSubstitutionAllowed:           true,
		VesselIMONumber:                          "9321483",
		PreCarriageModeOfTransportCode:            "VESSEL",
		InvoicePayableAt: Location{
			LocationName:   "Eiffel Tower",
			UNLocationCode: "FRPAR",
		},
		PlaceOfBLIssue: Location{
			LocationName:   "DCSA Headquarters",
			UNLocationCode: "NLAMS",
		},
		Commodities: []Commodity{
			{
				CommodityType:               "Mobile phones",
				HSCode:                      "string",
				CargoGrossWeight:            12000,
				CargoGrossWeightUnit:        "KGM",
				CargoGrossVolume:            120,
				CargoGrossVolumeUnit:        "MTQ",
				NumberOfPackages:            18,
				ExportLicenseIssueDate:      "2021-05-14",
				ExportLicenseExpiryDate:     "2021-05-21",
				CommodityRequestedEquipment: "001",
			},
		},
		ValueAddedServices: []ValueAddedService{
			{
				ValueAddedServiceCode: "SCON",
			},
		},
		References: []Reference{
			{
				Type:  "FF",
				Value: "string",
			},
		},
		RequestedEquipments: []RequestedEquipment{
			{
				ISOEquipmentCode:              "22RT",
				TareWeight:                    4800,
				TareWeightUnit:                "KGM",
				Units:                         3,
				EquipmentReferences:           []string{"PSSU2109481"},
				IsShipperOwned:                true,
				CommodityRequestedEquipment:   "001",
			},
		},
		DocumentParties: []DocumentParty{
			{
				Party: Party{
					PartyName:     "Asseco Denmark",
					TaxReference1: "CVR-25645774",
					TaxReference2: "CVR-25645774",
					PublicKey:     "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkFzaW",
					Address: Address{
						Name:         "Henrik",
						Street:       "Kronprinsessegade",
						StreetNumber: "54",
						Floor:        "5. sal",
						PostCode:     "1306",
						City:         "København",
						StateRegion:  "N/A",
						Country:      "Denmark",
					},
					PartyContactDetails: []PartyContactDetails{
						{
							Name:  "Henrik",
							Phone: "+45 33364660",
							Email: "info@dcsa.org",
							URL:   "https://www.dcsa.org",
						},
					},
					IdentifyingCodes: []IdentifyingCodes{
						{
							DCSAResponsibleAgencyCode: "SMDG",
							PartyCode:                 "MSK",
							CodeListName:              "LCL",
						},
					},
				},
				PartyFunction:   "DDS",
				DisplayedAddress: []string{"Kronprincessegade 54"},
				IsToBeNotified:  true,
			},
		},
		ShipmentLocations: []ShipmentLocation{
			{
				Location: Location{
					LocationName:             "CMP Container Terminal Copenhagen",
					UNLocationCode:           "DKCPH",
					FacilityCode:             "CMPDK",
					FacilityCodeListProvider: "SMDG",
				},
				ShipmentLocationTypeCode: "PRE",
				EventDateTime:            "2021-11-03T10:23:00+01:00",
			},
		},
	}
	
//marshal the obj
	assetBytes, err := json.Marshal(asset)
	if err != nil {
		log.Println("CreateBookingRequest:Unable to marshal Booking Request field for the bookingid " + bookingid)

		return err
	}

	// test, err := t.shipmentEvent(ctx, bookingid)
	
	//end of updating secondary details
	ctx.GetStub().SetEvent("Create BookingRequest", assetBytes)

	return ctx.GetStub().PutState(bookingid, assetBytes)
	// if err != nil {
	// 	log.Println("CreateBookingRequest:Error while creation of booking request info asset for bookingid" + bookingid)
	// 	return err
	// }
	// ctx.GetStub().SetEvent("Create BookingRequest", assetBytes)
	// return nil
}

// ReadAsset retrieves an asset from the ledger
func (s *BookingSmartContract) ReadBookingRequest(ctx contractapi.TransactionContextInterface, bookingid string) (*BookingRequestAsset, error) {
	assetBytes, err := ctx.GetStub().GetState(bookingid)
	if err != nil {
		// log.Println("ReadBookingRequest:Failed to read booking request asset with bookingid" + bookingid)
		return nil, fmt.Errorf("failed to get asset %s: %v", bookingid, err)
	}
	if assetBytes == nil {
		// log.Println("ReadBookingRequest:Booking requestasset does not exists for bookingid" + bookingid)
		return nil, fmt.Errorf("asset %s does not exist", bookingid)
	}

	var asset BookingRequestAsset
	//unmarshal the obj to get the data 

	err = json.Unmarshal(assetBytes, &asset)
	if err != nil {
		log.Println("ReadBookingRequest:Unable to unmarshal booking request asset for bookingid" + bookingid)
		return nil, err
	}

	return &asset, nil
}

// DeleteAsset removes an asset key-value pair from the ledger
func (t *BookingSmartContract) DeleteBookingRequest(ctx contractapi.TransactionContextInterface, bookingid string) error {
	asset, err := t.ReadBookingRequest(ctx, bookingid)

	if err != nil {
		log.Printf("DeleteBookingRequest: unable to read Booking Request details for bookingid %s,%v", bookingid, err)

		return err
	}
	log.Printf("BookingRequest DeleteAsset:Deleting asset %v", asset)
	ctx.GetStub().DelState(bookingid)

	return nil
}



func main() {
	chaincode, err := contractapi.NewChaincode(new(BookingSmartContract))
	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
	// err := shim.Start(new(SmartContract))
	// if err != nil {
	//     fmt.Printf("Error starting chaincode: %s", err)
	// }
}

