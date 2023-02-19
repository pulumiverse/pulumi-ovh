// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Hosting.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ovh.Hosting.inputs.PrivateDatabaseOrderArgs;
import com.pulumi.ovh.Hosting.inputs.PrivateDatabasePlanArgs;
import com.pulumi.ovh.Hosting.inputs.PrivateDatabasePlanOptionArgs;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PrivateDatabaseState extends com.pulumi.resources.ResourceArgs {

    public static final PrivateDatabaseState Empty = new PrivateDatabaseState();

    /**
     * Number of CPU on your private database
     * 
     */
    @Import(name="cpu")
    private @Nullable Output<Integer> cpu;

    /**
     * @return Number of CPU on your private database
     * 
     */
    public Optional<Output<Integer>> cpu() {
        return Optional.ofNullable(this.cpu);
    }

    /**
     * Datacenter where this private database is located
     * 
     */
    @Import(name="datacenter")
    private @Nullable Output<String> datacenter;

    /**
     * @return Datacenter where this private database is located
     * 
     */
    public Optional<Output<String>> datacenter() {
        return Optional.ofNullable(this.datacenter);
    }

    /**
     * Name displayed in customer panel for your private database
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Name displayed in customer panel for your private database
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Private database hostname
     * 
     */
    @Import(name="hostname")
    private @Nullable Output<String> hostname;

    /**
     * @return Private database hostname
     * 
     */
    public Optional<Output<String>> hostname() {
        return Optional.ofNullable(this.hostname);
    }

    /**
     * Private database FTP hostname
     * 
     */
    @Import(name="hostnameFtp")
    private @Nullable Output<String> hostnameFtp;

    /**
     * @return Private database FTP hostname
     * 
     */
    public Optional<Output<String>> hostnameFtp() {
        return Optional.ofNullable(this.hostnameFtp);
    }

    /**
     * Infrastructure where service was stored
     * 
     */
    @Import(name="infrastructure")
    private @Nullable Output<String> infrastructure;

    /**
     * @return Infrastructure where service was stored
     * 
     */
    public Optional<Output<String>> infrastructure() {
        return Optional.ofNullable(this.infrastructure);
    }

    /**
     * Type of the private database offer
     * 
     */
    @Import(name="offer")
    private @Nullable Output<String> offer;

    /**
     * @return Type of the private database offer
     * 
     */
    public Optional<Output<String>> offer() {
        return Optional.ofNullable(this.offer);
    }

    /**
     * Details about your Order
     * 
     */
    @Import(name="orders")
    private @Nullable Output<List<PrivateDatabaseOrderArgs>> orders;

    /**
     * @return Details about your Order
     * 
     */
    public Optional<Output<List<PrivateDatabaseOrderArgs>>> orders() {
        return Optional.ofNullable(this.orders);
    }

    /**
     * OVHcloud Subsidiary
     * 
     */
    @Import(name="ovhSubsidiary")
    private @Nullable Output<String> ovhSubsidiary;

    /**
     * @return OVHcloud Subsidiary
     * 
     */
    public Optional<Output<String>> ovhSubsidiary() {
        return Optional.ofNullable(this.ovhSubsidiary);
    }

    /**
     * OVHcloud payment mode (One of &#34;default-payment-mean&#34;, &#34;fidelity&#34;, &#34;ovh-account&#34;)
     * 
     */
    @Import(name="paymentMean")
    private @Nullable Output<String> paymentMean;

    /**
     * @return OVHcloud payment mode (One of &#34;default-payment-mean&#34;, &#34;fidelity&#34;, &#34;ovh-account&#34;)
     * 
     */
    public Optional<Output<String>> paymentMean() {
        return Optional.ofNullable(this.paymentMean);
    }

    /**
     * Product Plan to order
     * 
     */
    @Import(name="plan")
    private @Nullable Output<PrivateDatabasePlanArgs> plan;

    /**
     * @return Product Plan to order
     * 
     */
    public Optional<Output<PrivateDatabasePlanArgs>> plan() {
        return Optional.ofNullable(this.plan);
    }

    /**
     * Product Plan to order
     * 
     */
    @Import(name="planOptions")
    private @Nullable Output<List<PrivateDatabasePlanOptionArgs>> planOptions;

    /**
     * @return Product Plan to order
     * 
     */
    public Optional<Output<List<PrivateDatabasePlanOptionArgs>>> planOptions() {
        return Optional.ofNullable(this.planOptions);
    }

    /**
     * Private database service port
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return Private database service port
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * Private database FTP port
     * 
     */
    @Import(name="portFtp")
    private @Nullable Output<Integer> portFtp;

    /**
     * @return Private database FTP port
     * 
     */
    public Optional<Output<Integer>> portFtp() {
        return Optional.ofNullable(this.portFtp);
    }

    /**
     * Space allowed (in MB) on your private database
     * 
     */
    @Import(name="quotaSize")
    private @Nullable Output<Integer> quotaSize;

    /**
     * @return Space allowed (in MB) on your private database
     * 
     */
    public Optional<Output<Integer>> quotaSize() {
        return Optional.ofNullable(this.quotaSize);
    }

    /**
     * Sapce used (in MB) on your private database
     * 
     */
    @Import(name="quotaUsed")
    private @Nullable Output<Integer> quotaUsed;

    /**
     * @return Sapce used (in MB) on your private database
     * 
     */
    public Optional<Output<Integer>> quotaUsed() {
        return Optional.ofNullable(this.quotaUsed);
    }

    /**
     * Amount of ram (in MB) on your private database
     * 
     */
    @Import(name="ram")
    private @Nullable Output<Integer> ram;

    /**
     * @return Amount of ram (in MB) on your private database
     * 
     */
    public Optional<Output<Integer>> ram() {
        return Optional.ofNullable(this.ram);
    }

    /**
     * Private database server name
     * 
     */
    @Import(name="server")
    private @Nullable Output<String> server;

    /**
     * @return Private database server name
     * 
     */
    public Optional<Output<String>> server() {
        return Optional.ofNullable(this.server);
    }

    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * Private database state
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return Private database state
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    /**
     * Private database type
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Private database type
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * Private database available versions
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return Private database available versions
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    /**
     * Private database version label
     * 
     */
    @Import(name="versionLabel")
    private @Nullable Output<String> versionLabel;

    /**
     * @return Private database version label
     * 
     */
    public Optional<Output<String>> versionLabel() {
        return Optional.ofNullable(this.versionLabel);
    }

    /**
     * Private database version number
     * 
     */
    @Import(name="versionNumber")
    private @Nullable Output<Double> versionNumber;

    /**
     * @return Private database version number
     * 
     */
    public Optional<Output<Double>> versionNumber() {
        return Optional.ofNullable(this.versionNumber);
    }

    private PrivateDatabaseState() {}

    private PrivateDatabaseState(PrivateDatabaseState $) {
        this.cpu = $.cpu;
        this.datacenter = $.datacenter;
        this.displayName = $.displayName;
        this.hostname = $.hostname;
        this.hostnameFtp = $.hostnameFtp;
        this.infrastructure = $.infrastructure;
        this.offer = $.offer;
        this.orders = $.orders;
        this.ovhSubsidiary = $.ovhSubsidiary;
        this.paymentMean = $.paymentMean;
        this.plan = $.plan;
        this.planOptions = $.planOptions;
        this.port = $.port;
        this.portFtp = $.portFtp;
        this.quotaSize = $.quotaSize;
        this.quotaUsed = $.quotaUsed;
        this.ram = $.ram;
        this.server = $.server;
        this.serviceName = $.serviceName;
        this.state = $.state;
        this.type = $.type;
        this.version = $.version;
        this.versionLabel = $.versionLabel;
        this.versionNumber = $.versionNumber;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PrivateDatabaseState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PrivateDatabaseState $;

        public Builder() {
            $ = new PrivateDatabaseState();
        }

        public Builder(PrivateDatabaseState defaults) {
            $ = new PrivateDatabaseState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cpu Number of CPU on your private database
         * 
         * @return builder
         * 
         */
        public Builder cpu(@Nullable Output<Integer> cpu) {
            $.cpu = cpu;
            return this;
        }

        /**
         * @param cpu Number of CPU on your private database
         * 
         * @return builder
         * 
         */
        public Builder cpu(Integer cpu) {
            return cpu(Output.of(cpu));
        }

        /**
         * @param datacenter Datacenter where this private database is located
         * 
         * @return builder
         * 
         */
        public Builder datacenter(@Nullable Output<String> datacenter) {
            $.datacenter = datacenter;
            return this;
        }

        /**
         * @param datacenter Datacenter where this private database is located
         * 
         * @return builder
         * 
         */
        public Builder datacenter(String datacenter) {
            return datacenter(Output.of(datacenter));
        }

        /**
         * @param displayName Name displayed in customer panel for your private database
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Name displayed in customer panel for your private database
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param hostname Private database hostname
         * 
         * @return builder
         * 
         */
        public Builder hostname(@Nullable Output<String> hostname) {
            $.hostname = hostname;
            return this;
        }

        /**
         * @param hostname Private database hostname
         * 
         * @return builder
         * 
         */
        public Builder hostname(String hostname) {
            return hostname(Output.of(hostname));
        }

        /**
         * @param hostnameFtp Private database FTP hostname
         * 
         * @return builder
         * 
         */
        public Builder hostnameFtp(@Nullable Output<String> hostnameFtp) {
            $.hostnameFtp = hostnameFtp;
            return this;
        }

        /**
         * @param hostnameFtp Private database FTP hostname
         * 
         * @return builder
         * 
         */
        public Builder hostnameFtp(String hostnameFtp) {
            return hostnameFtp(Output.of(hostnameFtp));
        }

        /**
         * @param infrastructure Infrastructure where service was stored
         * 
         * @return builder
         * 
         */
        public Builder infrastructure(@Nullable Output<String> infrastructure) {
            $.infrastructure = infrastructure;
            return this;
        }

        /**
         * @param infrastructure Infrastructure where service was stored
         * 
         * @return builder
         * 
         */
        public Builder infrastructure(String infrastructure) {
            return infrastructure(Output.of(infrastructure));
        }

        /**
         * @param offer Type of the private database offer
         * 
         * @return builder
         * 
         */
        public Builder offer(@Nullable Output<String> offer) {
            $.offer = offer;
            return this;
        }

        /**
         * @param offer Type of the private database offer
         * 
         * @return builder
         * 
         */
        public Builder offer(String offer) {
            return offer(Output.of(offer));
        }

        /**
         * @param orders Details about your Order
         * 
         * @return builder
         * 
         */
        public Builder orders(@Nullable Output<List<PrivateDatabaseOrderArgs>> orders) {
            $.orders = orders;
            return this;
        }

        /**
         * @param orders Details about your Order
         * 
         * @return builder
         * 
         */
        public Builder orders(List<PrivateDatabaseOrderArgs> orders) {
            return orders(Output.of(orders));
        }

        /**
         * @param orders Details about your Order
         * 
         * @return builder
         * 
         */
        public Builder orders(PrivateDatabaseOrderArgs... orders) {
            return orders(List.of(orders));
        }

        /**
         * @param ovhSubsidiary OVHcloud Subsidiary
         * 
         * @return builder
         * 
         */
        public Builder ovhSubsidiary(@Nullable Output<String> ovhSubsidiary) {
            $.ovhSubsidiary = ovhSubsidiary;
            return this;
        }

        /**
         * @param ovhSubsidiary OVHcloud Subsidiary
         * 
         * @return builder
         * 
         */
        public Builder ovhSubsidiary(String ovhSubsidiary) {
            return ovhSubsidiary(Output.of(ovhSubsidiary));
        }

        /**
         * @param paymentMean OVHcloud payment mode (One of &#34;default-payment-mean&#34;, &#34;fidelity&#34;, &#34;ovh-account&#34;)
         * 
         * @return builder
         * 
         */
        public Builder paymentMean(@Nullable Output<String> paymentMean) {
            $.paymentMean = paymentMean;
            return this;
        }

        /**
         * @param paymentMean OVHcloud payment mode (One of &#34;default-payment-mean&#34;, &#34;fidelity&#34;, &#34;ovh-account&#34;)
         * 
         * @return builder
         * 
         */
        public Builder paymentMean(String paymentMean) {
            return paymentMean(Output.of(paymentMean));
        }

        /**
         * @param plan Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder plan(@Nullable Output<PrivateDatabasePlanArgs> plan) {
            $.plan = plan;
            return this;
        }

        /**
         * @param plan Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder plan(PrivateDatabasePlanArgs plan) {
            return plan(Output.of(plan));
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(@Nullable Output<List<PrivateDatabasePlanOptionArgs>> planOptions) {
            $.planOptions = planOptions;
            return this;
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(List<PrivateDatabasePlanOptionArgs> planOptions) {
            return planOptions(Output.of(planOptions));
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(PrivateDatabasePlanOptionArgs... planOptions) {
            return planOptions(List.of(planOptions));
        }

        /**
         * @param port Private database service port
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Private database service port
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param portFtp Private database FTP port
         * 
         * @return builder
         * 
         */
        public Builder portFtp(@Nullable Output<Integer> portFtp) {
            $.portFtp = portFtp;
            return this;
        }

        /**
         * @param portFtp Private database FTP port
         * 
         * @return builder
         * 
         */
        public Builder portFtp(Integer portFtp) {
            return portFtp(Output.of(portFtp));
        }

        /**
         * @param quotaSize Space allowed (in MB) on your private database
         * 
         * @return builder
         * 
         */
        public Builder quotaSize(@Nullable Output<Integer> quotaSize) {
            $.quotaSize = quotaSize;
            return this;
        }

        /**
         * @param quotaSize Space allowed (in MB) on your private database
         * 
         * @return builder
         * 
         */
        public Builder quotaSize(Integer quotaSize) {
            return quotaSize(Output.of(quotaSize));
        }

        /**
         * @param quotaUsed Sapce used (in MB) on your private database
         * 
         * @return builder
         * 
         */
        public Builder quotaUsed(@Nullable Output<Integer> quotaUsed) {
            $.quotaUsed = quotaUsed;
            return this;
        }

        /**
         * @param quotaUsed Sapce used (in MB) on your private database
         * 
         * @return builder
         * 
         */
        public Builder quotaUsed(Integer quotaUsed) {
            return quotaUsed(Output.of(quotaUsed));
        }

        /**
         * @param ram Amount of ram (in MB) on your private database
         * 
         * @return builder
         * 
         */
        public Builder ram(@Nullable Output<Integer> ram) {
            $.ram = ram;
            return this;
        }

        /**
         * @param ram Amount of ram (in MB) on your private database
         * 
         * @return builder
         * 
         */
        public Builder ram(Integer ram) {
            return ram(Output.of(ram));
        }

        /**
         * @param server Private database server name
         * 
         * @return builder
         * 
         */
        public Builder server(@Nullable Output<String> server) {
            $.server = server;
            return this;
        }

        /**
         * @param server Private database server name
         * 
         * @return builder
         * 
         */
        public Builder server(String server) {
            return server(Output.of(server));
        }

        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param state Private database state
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state Private database state
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        /**
         * @param type Private database type
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Private database type
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param version Private database available versions
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Private database available versions
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        /**
         * @param versionLabel Private database version label
         * 
         * @return builder
         * 
         */
        public Builder versionLabel(@Nullable Output<String> versionLabel) {
            $.versionLabel = versionLabel;
            return this;
        }

        /**
         * @param versionLabel Private database version label
         * 
         * @return builder
         * 
         */
        public Builder versionLabel(String versionLabel) {
            return versionLabel(Output.of(versionLabel));
        }

        /**
         * @param versionNumber Private database version number
         * 
         * @return builder
         * 
         */
        public Builder versionNumber(@Nullable Output<Double> versionNumber) {
            $.versionNumber = versionNumber;
            return this;
        }

        /**
         * @param versionNumber Private database version number
         * 
         * @return builder
         * 
         */
        public Builder versionNumber(Double versionNumber) {
            return versionNumber(Output.of(versionNumber));
        }

        public PrivateDatabaseState build() {
            return $;
        }
    }

}