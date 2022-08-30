// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.outputs.GetIpServiceRoutedTo;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetIpServiceResult {
    /**
     * @return can be terminated
     * 
     */
    private Boolean canBeTerminated;
    /**
     * @return country
     * 
     */
    private String country;
    /**
     * @return Custom description on your ip
     * 
     */
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return ip block
     * 
     */
    private String ip;
    /**
     * @return IP block organisation Id
     * 
     */
    private String organisationId;
    /**
     * @return Routage information
     * 
     */
    private List<GetIpServiceRoutedTo> routedTos;
    /**
     * @return Service where ip is routed to
     * 
     */
    private String serviceName;
    /**
     * @return Possible values for ip type (    &#34;cdn&#34;, &#34;cloud&#34;, &#34;dedicated&#34;, &#34;failover&#34;, &#34;hosted_ssl&#34;, &#34;housing&#34;, &#34;loadBalancing&#34;, &#34;mail&#34;, &#34;overthebox&#34;, &#34;pcc&#34;, &#34;pci&#34;, &#34;private&#34;, &#34;vpn&#34;, &#34;vps&#34;, &#34;vrack&#34;, &#34;xdsl&#34;)
     * 
     */
    private String type;

    private GetIpServiceResult() {}
    /**
     * @return can be terminated
     * 
     */
    public Boolean canBeTerminated() {
        return this.canBeTerminated;
    }
    /**
     * @return country
     * 
     */
    public String country() {
        return this.country;
    }
    /**
     * @return Custom description on your ip
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return ip block
     * 
     */
    public String ip() {
        return this.ip;
    }
    /**
     * @return IP block organisation Id
     * 
     */
    public String organisationId() {
        return this.organisationId;
    }
    /**
     * @return Routage information
     * 
     */
    public List<GetIpServiceRoutedTo> routedTos() {
        return this.routedTos;
    }
    /**
     * @return Service where ip is routed to
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return Possible values for ip type (    &#34;cdn&#34;, &#34;cloud&#34;, &#34;dedicated&#34;, &#34;failover&#34;, &#34;hosted_ssl&#34;, &#34;housing&#34;, &#34;loadBalancing&#34;, &#34;mail&#34;, &#34;overthebox&#34;, &#34;pcc&#34;, &#34;pci&#34;, &#34;private&#34;, &#34;vpn&#34;, &#34;vps&#34;, &#34;vrack&#34;, &#34;xdsl&#34;)
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpServiceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean canBeTerminated;
        private String country;
        private String description;
        private String id;
        private String ip;
        private String organisationId;
        private List<GetIpServiceRoutedTo> routedTos;
        private String serviceName;
        private String type;
        public Builder() {}
        public Builder(GetIpServiceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.canBeTerminated = defaults.canBeTerminated;
    	      this.country = defaults.country;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.ip = defaults.ip;
    	      this.organisationId = defaults.organisationId;
    	      this.routedTos = defaults.routedTos;
    	      this.serviceName = defaults.serviceName;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder canBeTerminated(Boolean canBeTerminated) {
            this.canBeTerminated = Objects.requireNonNull(canBeTerminated);
            return this;
        }
        @CustomType.Setter
        public Builder country(String country) {
            this.country = Objects.requireNonNull(country);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ip(String ip) {
            this.ip = Objects.requireNonNull(ip);
            return this;
        }
        @CustomType.Setter
        public Builder organisationId(String organisationId) {
            this.organisationId = Objects.requireNonNull(organisationId);
            return this;
        }
        @CustomType.Setter
        public Builder routedTos(List<GetIpServiceRoutedTo> routedTos) {
            this.routedTos = Objects.requireNonNull(routedTos);
            return this;
        }
        public Builder routedTos(GetIpServiceRoutedTo... routedTos) {
            return routedTos(List.of(routedTos));
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            this.serviceName = Objects.requireNonNull(serviceName);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetIpServiceResult build() {
            final var o = new GetIpServiceResult();
            o.canBeTerminated = canBeTerminated;
            o.country = country;
            o.description = description;
            o.id = id;
            o.ip = ip;
            o.organisationId = organisationId;
            o.routedTos = routedTos;
            o.serviceName = serviceName;
            o.type = type;
            return o;
        }
    }
}
