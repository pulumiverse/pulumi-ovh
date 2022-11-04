// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Me.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.Me.outputs.GetMeCurrency;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetMeResult {
    private String address;
    private String area;
    private String birthCity;
    private String birthDay;
    private String city;
    private String companyNationalIdentificationNumber;
    private String corporationType;
    private String country;
    private List<GetMeCurrency> currencies;
    private String customerCode;
    private String email;
    private String fax;
    private String firstname;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String italianSdi;
    private String language;
    private String legalform;
    private String name;
    private String nationalIdentificationNumber;
    private String nichandle;
    private String organisation;
    private String ovhCompany;
    private String ovhSubsidiary;
    private String phone;
    private String phoneCountry;
    private String sex;
    private String spareEmail;
    private String state;
    private String vat;
    private String zip;

    private GetMeResult() {}
    public String address() {
        return this.address;
    }
    public String area() {
        return this.area;
    }
    public String birthCity() {
        return this.birthCity;
    }
    public String birthDay() {
        return this.birthDay;
    }
    public String city() {
        return this.city;
    }
    public String companyNationalIdentificationNumber() {
        return this.companyNationalIdentificationNumber;
    }
    public String corporationType() {
        return this.corporationType;
    }
    public String country() {
        return this.country;
    }
    public List<GetMeCurrency> currencies() {
        return this.currencies;
    }
    public String customerCode() {
        return this.customerCode;
    }
    public String email() {
        return this.email;
    }
    public String fax() {
        return this.fax;
    }
    public String firstname() {
        return this.firstname;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String italianSdi() {
        return this.italianSdi;
    }
    public String language() {
        return this.language;
    }
    public String legalform() {
        return this.legalform;
    }
    public String name() {
        return this.name;
    }
    public String nationalIdentificationNumber() {
        return this.nationalIdentificationNumber;
    }
    public String nichandle() {
        return this.nichandle;
    }
    public String organisation() {
        return this.organisation;
    }
    public String ovhCompany() {
        return this.ovhCompany;
    }
    public String ovhSubsidiary() {
        return this.ovhSubsidiary;
    }
    public String phone() {
        return this.phone;
    }
    public String phoneCountry() {
        return this.phoneCountry;
    }
    public String sex() {
        return this.sex;
    }
    public String spareEmail() {
        return this.spareEmail;
    }
    public String state() {
        return this.state;
    }
    public String vat() {
        return this.vat;
    }
    public String zip() {
        return this.zip;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String address;
        private String area;
        private String birthCity;
        private String birthDay;
        private String city;
        private String companyNationalIdentificationNumber;
        private String corporationType;
        private String country;
        private List<GetMeCurrency> currencies;
        private String customerCode;
        private String email;
        private String fax;
        private String firstname;
        private String id;
        private String italianSdi;
        private String language;
        private String legalform;
        private String name;
        private String nationalIdentificationNumber;
        private String nichandle;
        private String organisation;
        private String ovhCompany;
        private String ovhSubsidiary;
        private String phone;
        private String phoneCountry;
        private String sex;
        private String spareEmail;
        private String state;
        private String vat;
        private String zip;
        public Builder() {}
        public Builder(GetMeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.address = defaults.address;
    	      this.area = defaults.area;
    	      this.birthCity = defaults.birthCity;
    	      this.birthDay = defaults.birthDay;
    	      this.city = defaults.city;
    	      this.companyNationalIdentificationNumber = defaults.companyNationalIdentificationNumber;
    	      this.corporationType = defaults.corporationType;
    	      this.country = defaults.country;
    	      this.currencies = defaults.currencies;
    	      this.customerCode = defaults.customerCode;
    	      this.email = defaults.email;
    	      this.fax = defaults.fax;
    	      this.firstname = defaults.firstname;
    	      this.id = defaults.id;
    	      this.italianSdi = defaults.italianSdi;
    	      this.language = defaults.language;
    	      this.legalform = defaults.legalform;
    	      this.name = defaults.name;
    	      this.nationalIdentificationNumber = defaults.nationalIdentificationNumber;
    	      this.nichandle = defaults.nichandle;
    	      this.organisation = defaults.organisation;
    	      this.ovhCompany = defaults.ovhCompany;
    	      this.ovhSubsidiary = defaults.ovhSubsidiary;
    	      this.phone = defaults.phone;
    	      this.phoneCountry = defaults.phoneCountry;
    	      this.sex = defaults.sex;
    	      this.spareEmail = defaults.spareEmail;
    	      this.state = defaults.state;
    	      this.vat = defaults.vat;
    	      this.zip = defaults.zip;
        }

        @CustomType.Setter
        public Builder address(String address) {
            this.address = Objects.requireNonNull(address);
            return this;
        }
        @CustomType.Setter
        public Builder area(String area) {
            this.area = Objects.requireNonNull(area);
            return this;
        }
        @CustomType.Setter
        public Builder birthCity(String birthCity) {
            this.birthCity = Objects.requireNonNull(birthCity);
            return this;
        }
        @CustomType.Setter
        public Builder birthDay(String birthDay) {
            this.birthDay = Objects.requireNonNull(birthDay);
            return this;
        }
        @CustomType.Setter
        public Builder city(String city) {
            this.city = Objects.requireNonNull(city);
            return this;
        }
        @CustomType.Setter
        public Builder companyNationalIdentificationNumber(String companyNationalIdentificationNumber) {
            this.companyNationalIdentificationNumber = Objects.requireNonNull(companyNationalIdentificationNumber);
            return this;
        }
        @CustomType.Setter
        public Builder corporationType(String corporationType) {
            this.corporationType = Objects.requireNonNull(corporationType);
            return this;
        }
        @CustomType.Setter
        public Builder country(String country) {
            this.country = Objects.requireNonNull(country);
            return this;
        }
        @CustomType.Setter
        public Builder currencies(List<GetMeCurrency> currencies) {
            this.currencies = Objects.requireNonNull(currencies);
            return this;
        }
        public Builder currencies(GetMeCurrency... currencies) {
            return currencies(List.of(currencies));
        }
        @CustomType.Setter
        public Builder customerCode(String customerCode) {
            this.customerCode = Objects.requireNonNull(customerCode);
            return this;
        }
        @CustomType.Setter
        public Builder email(String email) {
            this.email = Objects.requireNonNull(email);
            return this;
        }
        @CustomType.Setter
        public Builder fax(String fax) {
            this.fax = Objects.requireNonNull(fax);
            return this;
        }
        @CustomType.Setter
        public Builder firstname(String firstname) {
            this.firstname = Objects.requireNonNull(firstname);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder italianSdi(String italianSdi) {
            this.italianSdi = Objects.requireNonNull(italianSdi);
            return this;
        }
        @CustomType.Setter
        public Builder language(String language) {
            this.language = Objects.requireNonNull(language);
            return this;
        }
        @CustomType.Setter
        public Builder legalform(String legalform) {
            this.legalform = Objects.requireNonNull(legalform);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder nationalIdentificationNumber(String nationalIdentificationNumber) {
            this.nationalIdentificationNumber = Objects.requireNonNull(nationalIdentificationNumber);
            return this;
        }
        @CustomType.Setter
        public Builder nichandle(String nichandle) {
            this.nichandle = Objects.requireNonNull(nichandle);
            return this;
        }
        @CustomType.Setter
        public Builder organisation(String organisation) {
            this.organisation = Objects.requireNonNull(organisation);
            return this;
        }
        @CustomType.Setter
        public Builder ovhCompany(String ovhCompany) {
            this.ovhCompany = Objects.requireNonNull(ovhCompany);
            return this;
        }
        @CustomType.Setter
        public Builder ovhSubsidiary(String ovhSubsidiary) {
            this.ovhSubsidiary = Objects.requireNonNull(ovhSubsidiary);
            return this;
        }
        @CustomType.Setter
        public Builder phone(String phone) {
            this.phone = Objects.requireNonNull(phone);
            return this;
        }
        @CustomType.Setter
        public Builder phoneCountry(String phoneCountry) {
            this.phoneCountry = Objects.requireNonNull(phoneCountry);
            return this;
        }
        @CustomType.Setter
        public Builder sex(String sex) {
            this.sex = Objects.requireNonNull(sex);
            return this;
        }
        @CustomType.Setter
        public Builder spareEmail(String spareEmail) {
            this.spareEmail = Objects.requireNonNull(spareEmail);
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            this.state = Objects.requireNonNull(state);
            return this;
        }
        @CustomType.Setter
        public Builder vat(String vat) {
            this.vat = Objects.requireNonNull(vat);
            return this;
        }
        @CustomType.Setter
        public Builder zip(String zip) {
            this.zip = Objects.requireNonNull(zip);
            return this;
        }
        public GetMeResult build() {
            final var o = new GetMeResult();
            o.address = address;
            o.area = area;
            o.birthCity = birthCity;
            o.birthDay = birthDay;
            o.city = city;
            o.companyNationalIdentificationNumber = companyNationalIdentificationNumber;
            o.corporationType = corporationType;
            o.country = country;
            o.currencies = currencies;
            o.customerCode = customerCode;
            o.email = email;
            o.fax = fax;
            o.firstname = firstname;
            o.id = id;
            o.italianSdi = italianSdi;
            o.language = language;
            o.legalform = legalform;
            o.name = name;
            o.nationalIdentificationNumber = nationalIdentificationNumber;
            o.nichandle = nichandle;
            o.organisation = organisation;
            o.ovhCompany = ovhCompany;
            o.ovhSubsidiary = ovhSubsidiary;
            o.phone = phone;
            o.phoneCountry = phoneCountry;
            o.sex = sex;
            o.spareEmail = spareEmail;
            o.state = state;
            o.vat = vat;
            o.zip = zip;
            return o;
        }
    }
}