// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Me.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIdentityUserResult {
    /**
     * @return Creation date of this user.
     * 
     */
    private String creation;
    /**
     * @return User description.
     * 
     */
    private String description;
    /**
     * @return User&#39;s email.
     * 
     */
    private String email;
    /**
     * @return User&#39;s group.
     * 
     */
    private String group;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Last update of this user.
     * 
     */
    private String lastUpdate;
    /**
     * @return User&#39;s login suffix.
     * 
     */
    private String login;
    /**
     * @return When the user changed his password for the last time.
     * 
     */
    private String passwordLastUpdate;
    /**
     * @return Current user&#39;s status.
     * 
     */
    private String status;
    private String user;

    private GetIdentityUserResult() {}
    /**
     * @return Creation date of this user.
     * 
     */
    public String creation() {
        return this.creation;
    }
    /**
     * @return User description.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return User&#39;s email.
     * 
     */
    public String email() {
        return this.email;
    }
    /**
     * @return User&#39;s group.
     * 
     */
    public String group() {
        return this.group;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Last update of this user.
     * 
     */
    public String lastUpdate() {
        return this.lastUpdate;
    }
    /**
     * @return User&#39;s login suffix.
     * 
     */
    public String login() {
        return this.login;
    }
    /**
     * @return When the user changed his password for the last time.
     * 
     */
    public String passwordLastUpdate() {
        return this.passwordLastUpdate;
    }
    /**
     * @return Current user&#39;s status.
     * 
     */
    public String status() {
        return this.status;
    }
    public String user() {
        return this.user;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIdentityUserResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String creation;
        private String description;
        private String email;
        private String group;
        private String id;
        private String lastUpdate;
        private String login;
        private String passwordLastUpdate;
        private String status;
        private String user;
        public Builder() {}
        public Builder(GetIdentityUserResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.creation = defaults.creation;
    	      this.description = defaults.description;
    	      this.email = defaults.email;
    	      this.group = defaults.group;
    	      this.id = defaults.id;
    	      this.lastUpdate = defaults.lastUpdate;
    	      this.login = defaults.login;
    	      this.passwordLastUpdate = defaults.passwordLastUpdate;
    	      this.status = defaults.status;
    	      this.user = defaults.user;
        }

        @CustomType.Setter
        public Builder creation(String creation) {
            this.creation = Objects.requireNonNull(creation);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder email(String email) {
            this.email = Objects.requireNonNull(email);
            return this;
        }
        @CustomType.Setter
        public Builder group(String group) {
            this.group = Objects.requireNonNull(group);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder lastUpdate(String lastUpdate) {
            this.lastUpdate = Objects.requireNonNull(lastUpdate);
            return this;
        }
        @CustomType.Setter
        public Builder login(String login) {
            this.login = Objects.requireNonNull(login);
            return this;
        }
        @CustomType.Setter
        public Builder passwordLastUpdate(String passwordLastUpdate) {
            this.passwordLastUpdate = Objects.requireNonNull(passwordLastUpdate);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder user(String user) {
            this.user = Objects.requireNonNull(user);
            return this;
        }
        public GetIdentityUserResult build() {
            final var o = new GetIdentityUserResult();
            o.creation = creation;
            o.description = description;
            o.email = email;
            o.group = group;
            o.id = id;
            o.lastUpdate = lastUpdate;
            o.login = login;
            o.passwordLastUpdate = passwordLastUpdate;
            o.status = status;
            o.user = user;
            return o;
        }
    }
}