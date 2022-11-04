// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about the current OVHcloud account.
func GetMe(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetMeResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetMeResult
	err := ctx.Invoke("ovh:Me/getMe:getMe", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getMe.
type GetMeResult struct {
	Address                             string          `pulumi:"address"`
	Area                                string          `pulumi:"area"`
	BirthCity                           string          `pulumi:"birthCity"`
	BirthDay                            string          `pulumi:"birthDay"`
	City                                string          `pulumi:"city"`
	CompanyNationalIdentificationNumber string          `pulumi:"companyNationalIdentificationNumber"`
	CorporationType                     string          `pulumi:"corporationType"`
	Country                             string          `pulumi:"country"`
	Currencies                          []GetMeCurrency `pulumi:"currencies"`
	CustomerCode                        string          `pulumi:"customerCode"`
	Email                               string          `pulumi:"email"`
	Fax                                 string          `pulumi:"fax"`
	Firstname                           string          `pulumi:"firstname"`
	// The provider-assigned unique ID for this managed resource.
	Id                           string `pulumi:"id"`
	ItalianSdi                   string `pulumi:"italianSdi"`
	Language                     string `pulumi:"language"`
	Legalform                    string `pulumi:"legalform"`
	Name                         string `pulumi:"name"`
	NationalIdentificationNumber string `pulumi:"nationalIdentificationNumber"`
	Nichandle                    string `pulumi:"nichandle"`
	Organisation                 string `pulumi:"organisation"`
	OvhCompany                   string `pulumi:"ovhCompany"`
	OvhSubsidiary                string `pulumi:"ovhSubsidiary"`
	Phone                        string `pulumi:"phone"`
	PhoneCountry                 string `pulumi:"phoneCountry"`
	Sex                          string `pulumi:"sex"`
	SpareEmail                   string `pulumi:"spareEmail"`
	State                        string `pulumi:"state"`
	Vat                          string `pulumi:"vat"`
	Zip                          string `pulumi:"zip"`
}
