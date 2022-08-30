// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.outputs.GetCloudProjectRegionService;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetCloudProjectRegionResult {
    /**
     * @return the code of the geographic continent the region is running.
     * E.g.: EU for Europe, US for America...
     * 
     */
    private String continentCode;
    /**
     * @return The location code of the datacenter.
     * E.g.: &#34;GRA&#34;, meaning Gravelines, for region &#34;GRA1&#34;
     * 
     */
    private String datacenterLocation;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return the name of the public cloud service
     * 
     */
    private String name;
    private String serviceName;
    /**
     * @return The list of public cloud services running within the region
     * 
     */
    private List<GetCloudProjectRegionService> services;

    private GetCloudProjectRegionResult() {}
    /**
     * @return the code of the geographic continent the region is running.
     * E.g.: EU for Europe, US for America...
     * 
     */
    public String continentCode() {
        return this.continentCode;
    }
    /**
     * @return The location code of the datacenter.
     * E.g.: &#34;GRA&#34;, meaning Gravelines, for region &#34;GRA1&#34;
     * 
     */
    public String datacenterLocation() {
        return this.datacenterLocation;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return the name of the public cloud service
     * 
     */
    public String name() {
        return this.name;
    }
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return The list of public cloud services running within the region
     * 
     */
    public List<GetCloudProjectRegionService> services() {
        return this.services;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCloudProjectRegionResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String continentCode;
        private String datacenterLocation;
        private String id;
        private String name;
        private String serviceName;
        private List<GetCloudProjectRegionService> services;
        public Builder() {}
        public Builder(GetCloudProjectRegionResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.continentCode = defaults.continentCode;
    	      this.datacenterLocation = defaults.datacenterLocation;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.serviceName = defaults.serviceName;
    	      this.services = defaults.services;
        }

        @CustomType.Setter
        public Builder continentCode(String continentCode) {
            this.continentCode = Objects.requireNonNull(continentCode);
            return this;
        }
        @CustomType.Setter
        public Builder datacenterLocation(String datacenterLocation) {
            this.datacenterLocation = Objects.requireNonNull(datacenterLocation);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            this.serviceName = Objects.requireNonNull(serviceName);
            return this;
        }
        @CustomType.Setter
        public Builder services(List<GetCloudProjectRegionService> services) {
            this.services = Objects.requireNonNull(services);
            return this;
        }
        public Builder services(GetCloudProjectRegionService... services) {
            return services(List.of(services));
        }
        public GetCloudProjectRegionResult build() {
            final var o = new GetCloudProjectRegionResult();
            o.continentCode = continentCode;
            o.datacenterLocation = datacenterLocation;
            o.id = id;
            o.name = name;
            o.serviceName = serviceName;
            o.services = services;
            return o;
        }
    }
}
