// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.outputs.GetCloudProjectDatabaseEndpoint;
import com.pulumi.ovh.outputs.GetCloudProjectDatabaseNode;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetCloudProjectDatabaseResult {
    /**
     * @return Time on which backups start every day.
     * 
     */
    private String backupTime;
    private String clusterId;
    /**
     * @return Date of the creation of the cluster.
     * 
     */
    private String createdAt;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String description;
    /**
     * @return List of all endpoints objects of the service.
     * 
     */
    private List<GetCloudProjectDatabaseEndpoint> endpoints;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String engine;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String flavor;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Time on which maintenances can start every day.
     * 
     */
    private String maintenanceTime;
    /**
     * @return Type of network of the cluster.
     * 
     */
    private String networkType;
    /**
     * @return See Argument Reference above.
     * 
     */
    private List<GetCloudProjectDatabaseNode> nodes;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String plan;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String serviceName;
    /**
     * @return Current status of the cluster.
     * 
     */
    private String status;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String version;

    private GetCloudProjectDatabaseResult() {}
    /**
     * @return Time on which backups start every day.
     * 
     */
    public String backupTime() {
        return this.backupTime;
    }
    public String clusterId() {
        return this.clusterId;
    }
    /**
     * @return Date of the creation of the cluster.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return List of all endpoints objects of the service.
     * 
     */
    public List<GetCloudProjectDatabaseEndpoint> endpoints() {
        return this.endpoints;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String engine() {
        return this.engine;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String flavor() {
        return this.flavor;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Time on which maintenances can start every day.
     * 
     */
    public String maintenanceTime() {
        return this.maintenanceTime;
    }
    /**
     * @return Type of network of the cluster.
     * 
     */
    public String networkType() {
        return this.networkType;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public List<GetCloudProjectDatabaseNode> nodes() {
        return this.nodes;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String plan() {
        return this.plan;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return Current status of the cluster.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCloudProjectDatabaseResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String backupTime;
        private String clusterId;
        private String createdAt;
        private String description;
        private List<GetCloudProjectDatabaseEndpoint> endpoints;
        private String engine;
        private String flavor;
        private String id;
        private String maintenanceTime;
        private String networkType;
        private List<GetCloudProjectDatabaseNode> nodes;
        private String plan;
        private String serviceName;
        private String status;
        private String version;
        public Builder() {}
        public Builder(GetCloudProjectDatabaseResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backupTime = defaults.backupTime;
    	      this.clusterId = defaults.clusterId;
    	      this.createdAt = defaults.createdAt;
    	      this.description = defaults.description;
    	      this.endpoints = defaults.endpoints;
    	      this.engine = defaults.engine;
    	      this.flavor = defaults.flavor;
    	      this.id = defaults.id;
    	      this.maintenanceTime = defaults.maintenanceTime;
    	      this.networkType = defaults.networkType;
    	      this.nodes = defaults.nodes;
    	      this.plan = defaults.plan;
    	      this.serviceName = defaults.serviceName;
    	      this.status = defaults.status;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder backupTime(String backupTime) {
            this.backupTime = Objects.requireNonNull(backupTime);
            return this;
        }
        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            this.clusterId = Objects.requireNonNull(clusterId);
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder endpoints(List<GetCloudProjectDatabaseEndpoint> endpoints) {
            this.endpoints = Objects.requireNonNull(endpoints);
            return this;
        }
        public Builder endpoints(GetCloudProjectDatabaseEndpoint... endpoints) {
            return endpoints(List.of(endpoints));
        }
        @CustomType.Setter
        public Builder engine(String engine) {
            this.engine = Objects.requireNonNull(engine);
            return this;
        }
        @CustomType.Setter
        public Builder flavor(String flavor) {
            this.flavor = Objects.requireNonNull(flavor);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder maintenanceTime(String maintenanceTime) {
            this.maintenanceTime = Objects.requireNonNull(maintenanceTime);
            return this;
        }
        @CustomType.Setter
        public Builder networkType(String networkType) {
            this.networkType = Objects.requireNonNull(networkType);
            return this;
        }
        @CustomType.Setter
        public Builder nodes(List<GetCloudProjectDatabaseNode> nodes) {
            this.nodes = Objects.requireNonNull(nodes);
            return this;
        }
        public Builder nodes(GetCloudProjectDatabaseNode... nodes) {
            return nodes(List.of(nodes));
        }
        @CustomType.Setter
        public Builder plan(String plan) {
            this.plan = Objects.requireNonNull(plan);
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            this.serviceName = Objects.requireNonNull(serviceName);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder version(String version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }
        public GetCloudProjectDatabaseResult build() {
            final var o = new GetCloudProjectDatabaseResult();
            o.backupTime = backupTime;
            o.clusterId = clusterId;
            o.createdAt = createdAt;
            o.description = description;
            o.endpoints = endpoints;
            o.engine = engine;
            o.flavor = flavor;
            o.id = id;
            o.maintenanceTime = maintenanceTime;
            o.networkType = networkType;
            o.nodes = nodes;
            o.plan = plan;
            o.serviceName = serviceName;
            o.status = status;
            o.version = version;
            return o;
        }
    }
}
