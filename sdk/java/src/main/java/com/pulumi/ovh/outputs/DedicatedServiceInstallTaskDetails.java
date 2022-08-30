// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DedicatedServiceInstallTaskDetails {
    /**
     * @return Template change log details.
     * 
     * @deprecated
     * field is not used anymore
     * 
     */
    @Deprecated /* field is not used anymore */
    private @Nullable String changeLog;
    /**
     * @return Set up the server using the provided hostname instead of the default hostname.
     * 
     */
    private @Nullable String customHostname;
    /**
     * @return Disk group id.
     * 
     */
    private @Nullable Integer diskGroupId;
    /**
     * @return set to true to install RTM.
     * 
     */
    private @Nullable Boolean installRtm;
    /**
     * @return set to true to install sql server (Windows template only).
     * 
     */
    private @Nullable Boolean installSqlServer;
    /**
     * @return language.
     * 
     */
    private @Nullable String language;
    /**
     * @return set to true to disable RAID.
     * 
     */
    private @Nullable Boolean noRaid;
    /**
     * @return Indicate the URL where your postinstall customisation script is located.
     * 
     */
    private @Nullable String postInstallationScriptLink;
    /**
     * @return Indicate the string returned by your postinstall customisation script on successful execution. Advice: your script should return a unique validation string in case of succes. A good example is &#39;loh1Xee7eo OK OK OK UGh8Ang1Gu&#39;.
     * 
     */
    private @Nullable String postInstallationScriptReturn;
    /**
     * @return set to true to make a hardware raid reset.
     * 
     */
    private @Nullable Boolean resetHwRaid;
    /**
     * @return soft raid devices.
     * 
     */
    private @Nullable Integer softRaidDevices;
    /**
     * @return Name of the ssh key that should be installed. Password login will be disabled.
     * 
     */
    private @Nullable String sshKeyName;
    /**
     * @return Use the distribution&#39;s native kernel instead of the recommended OVH Kernel.
     * 
     */
    private @Nullable Boolean useDistribKernel;
    /**
     * @return set to true to use SPLA.
     * 
     */
    private @Nullable Boolean useSpla;

    private DedicatedServiceInstallTaskDetails() {}
    /**
     * @return Template change log details.
     * 
     * @deprecated
     * field is not used anymore
     * 
     */
    @Deprecated /* field is not used anymore */
    public Optional<String> changeLog() {
        return Optional.ofNullable(this.changeLog);
    }
    /**
     * @return Set up the server using the provided hostname instead of the default hostname.
     * 
     */
    public Optional<String> customHostname() {
        return Optional.ofNullable(this.customHostname);
    }
    /**
     * @return Disk group id.
     * 
     */
    public Optional<Integer> diskGroupId() {
        return Optional.ofNullable(this.diskGroupId);
    }
    /**
     * @return set to true to install RTM.
     * 
     */
    public Optional<Boolean> installRtm() {
        return Optional.ofNullable(this.installRtm);
    }
    /**
     * @return set to true to install sql server (Windows template only).
     * 
     */
    public Optional<Boolean> installSqlServer() {
        return Optional.ofNullable(this.installSqlServer);
    }
    /**
     * @return language.
     * 
     */
    public Optional<String> language() {
        return Optional.ofNullable(this.language);
    }
    /**
     * @return set to true to disable RAID.
     * 
     */
    public Optional<Boolean> noRaid() {
        return Optional.ofNullable(this.noRaid);
    }
    /**
     * @return Indicate the URL where your postinstall customisation script is located.
     * 
     */
    public Optional<String> postInstallationScriptLink() {
        return Optional.ofNullable(this.postInstallationScriptLink);
    }
    /**
     * @return Indicate the string returned by your postinstall customisation script on successful execution. Advice: your script should return a unique validation string in case of succes. A good example is &#39;loh1Xee7eo OK OK OK UGh8Ang1Gu&#39;.
     * 
     */
    public Optional<String> postInstallationScriptReturn() {
        return Optional.ofNullable(this.postInstallationScriptReturn);
    }
    /**
     * @return set to true to make a hardware raid reset.
     * 
     */
    public Optional<Boolean> resetHwRaid() {
        return Optional.ofNullable(this.resetHwRaid);
    }
    /**
     * @return soft raid devices.
     * 
     */
    public Optional<Integer> softRaidDevices() {
        return Optional.ofNullable(this.softRaidDevices);
    }
    /**
     * @return Name of the ssh key that should be installed. Password login will be disabled.
     * 
     */
    public Optional<String> sshKeyName() {
        return Optional.ofNullable(this.sshKeyName);
    }
    /**
     * @return Use the distribution&#39;s native kernel instead of the recommended OVH Kernel.
     * 
     */
    public Optional<Boolean> useDistribKernel() {
        return Optional.ofNullable(this.useDistribKernel);
    }
    /**
     * @return set to true to use SPLA.
     * 
     */
    public Optional<Boolean> useSpla() {
        return Optional.ofNullable(this.useSpla);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DedicatedServiceInstallTaskDetails defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String changeLog;
        private @Nullable String customHostname;
        private @Nullable Integer diskGroupId;
        private @Nullable Boolean installRtm;
        private @Nullable Boolean installSqlServer;
        private @Nullable String language;
        private @Nullable Boolean noRaid;
        private @Nullable String postInstallationScriptLink;
        private @Nullable String postInstallationScriptReturn;
        private @Nullable Boolean resetHwRaid;
        private @Nullable Integer softRaidDevices;
        private @Nullable String sshKeyName;
        private @Nullable Boolean useDistribKernel;
        private @Nullable Boolean useSpla;
        public Builder() {}
        public Builder(DedicatedServiceInstallTaskDetails defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.changeLog = defaults.changeLog;
    	      this.customHostname = defaults.customHostname;
    	      this.diskGroupId = defaults.diskGroupId;
    	      this.installRtm = defaults.installRtm;
    	      this.installSqlServer = defaults.installSqlServer;
    	      this.language = defaults.language;
    	      this.noRaid = defaults.noRaid;
    	      this.postInstallationScriptLink = defaults.postInstallationScriptLink;
    	      this.postInstallationScriptReturn = defaults.postInstallationScriptReturn;
    	      this.resetHwRaid = defaults.resetHwRaid;
    	      this.softRaidDevices = defaults.softRaidDevices;
    	      this.sshKeyName = defaults.sshKeyName;
    	      this.useDistribKernel = defaults.useDistribKernel;
    	      this.useSpla = defaults.useSpla;
        }

        @CustomType.Setter
        public Builder changeLog(@Nullable String changeLog) {
            this.changeLog = changeLog;
            return this;
        }
        @CustomType.Setter
        public Builder customHostname(@Nullable String customHostname) {
            this.customHostname = customHostname;
            return this;
        }
        @CustomType.Setter
        public Builder diskGroupId(@Nullable Integer diskGroupId) {
            this.diskGroupId = diskGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder installRtm(@Nullable Boolean installRtm) {
            this.installRtm = installRtm;
            return this;
        }
        @CustomType.Setter
        public Builder installSqlServer(@Nullable Boolean installSqlServer) {
            this.installSqlServer = installSqlServer;
            return this;
        }
        @CustomType.Setter
        public Builder language(@Nullable String language) {
            this.language = language;
            return this;
        }
        @CustomType.Setter
        public Builder noRaid(@Nullable Boolean noRaid) {
            this.noRaid = noRaid;
            return this;
        }
        @CustomType.Setter
        public Builder postInstallationScriptLink(@Nullable String postInstallationScriptLink) {
            this.postInstallationScriptLink = postInstallationScriptLink;
            return this;
        }
        @CustomType.Setter
        public Builder postInstallationScriptReturn(@Nullable String postInstallationScriptReturn) {
            this.postInstallationScriptReturn = postInstallationScriptReturn;
            return this;
        }
        @CustomType.Setter
        public Builder resetHwRaid(@Nullable Boolean resetHwRaid) {
            this.resetHwRaid = resetHwRaid;
            return this;
        }
        @CustomType.Setter
        public Builder softRaidDevices(@Nullable Integer softRaidDevices) {
            this.softRaidDevices = softRaidDevices;
            return this;
        }
        @CustomType.Setter
        public Builder sshKeyName(@Nullable String sshKeyName) {
            this.sshKeyName = sshKeyName;
            return this;
        }
        @CustomType.Setter
        public Builder useDistribKernel(@Nullable Boolean useDistribKernel) {
            this.useDistribKernel = useDistribKernel;
            return this;
        }
        @CustomType.Setter
        public Builder useSpla(@Nullable Boolean useSpla) {
            this.useSpla = useSpla;
            return this;
        }
        public DedicatedServiceInstallTaskDetails build() {
            final var o = new DedicatedServiceInstallTaskDetails();
            o.changeLog = changeLog;
            o.customHostname = customHostname;
            o.diskGroupId = diskGroupId;
            o.installRtm = installRtm;
            o.installSqlServer = installSqlServer;
            o.language = language;
            o.noRaid = noRaid;
            o.postInstallationScriptLink = postInstallationScriptLink;
            o.postInstallationScriptReturn = postInstallationScriptReturn;
            o.resetHwRaid = resetHwRaid;
            o.softRaidDevices = softRaidDevices;
            o.sshKeyName = sshKeyName;
            o.useDistribKernel = useDistribKernel;
            o.useSpla = useSpla;
            return o;
        }
    }
}
