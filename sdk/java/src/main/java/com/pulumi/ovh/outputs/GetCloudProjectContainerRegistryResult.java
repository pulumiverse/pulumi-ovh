// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetCloudProjectContainerRegistryResult {
    /**
     * @return Registry creation date
     * 
     */
    private String createdAt;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Registry name
     * 
     */
    private String name;
    /**
     * @return Project ID of your registry
     * 
     */
    private String projectId;
    /**
     * @return Region of the registry
     * 
     */
    private String region;
    private String registryId;
    private String serviceName;
    /**
     * @return Current size of the registry (bytes)
     * 
     */
    private Integer size;
    /**
     * @return Registry status
     * 
     */
    private String status;
    /**
     * @return Registry last update date
     * 
     */
    private String updatedAt;
    /**
     * @return Access url of the registry
     * 
     */
    private String url;
    /**
     * @return Version of your registry
     * 
     */
    private String version;

    private GetCloudProjectContainerRegistryResult() {}
    /**
     * @return Registry creation date
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Registry name
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Project ID of your registry
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return Region of the registry
     * 
     */
    public String region() {
        return this.region;
    }
    public String registryId() {
        return this.registryId;
    }
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return Current size of the registry (bytes)
     * 
     */
    public Integer size() {
        return this.size;
    }
    /**
     * @return Registry status
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return Registry last update date
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }
    /**
     * @return Access url of the registry
     * 
     */
    public String url() {
        return this.url;
    }
    /**
     * @return Version of your registry
     * 
     */
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCloudProjectContainerRegistryResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String id;
        private String name;
        private String projectId;
        private String region;
        private String registryId;
        private String serviceName;
        private Integer size;
        private String status;
        private String updatedAt;
        private String url;
        private String version;
        public Builder() {}
        public Builder(GetCloudProjectContainerRegistryResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.registryId = defaults.registryId;
    	      this.serviceName = defaults.serviceName;
    	      this.size = defaults.size;
    	      this.status = defaults.status;
    	      this.updatedAt = defaults.updatedAt;
    	      this.url = defaults.url;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
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
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder registryId(String registryId) {
            this.registryId = Objects.requireNonNull(registryId);
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            this.serviceName = Objects.requireNonNull(serviceName);
            return this;
        }
        @CustomType.Setter
        public Builder size(Integer size) {
            this.size = Objects.requireNonNull(size);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }
        @CustomType.Setter
        public Builder version(String version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }
        public GetCloudProjectContainerRegistryResult build() {
            final var o = new GetCloudProjectContainerRegistryResult();
            o.createdAt = createdAt;
            o.id = id;
            o.name = name;
            o.projectId = projectId;
            o.region = region;
            o.registryId = registryId;
            o.serviceName = serviceName;
            o.size = size;
            o.status = status;
            o.updatedAt = updatedAt;
            o.url = url;
            o.version = version;
            return o;
        }
    }
}
