// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetPropertiesSearchProperty {
    private String accountId;
    private String assetId;
    private String contractId;
    private String edgeHostname;
    private String groupId;
    private String hostname;
    private String productionStatus;
    private String propertyId;
    private String propertyName;
    private Integer propertyVersion;
    private String stagingStatus;
    private String updatedByUser;
    private String updatedDate;

    private GetPropertiesSearchProperty() {}
    public String accountId() {
        return this.accountId;
    }
    public String assetId() {
        return this.assetId;
    }
    public String contractId() {
        return this.contractId;
    }
    public String edgeHostname() {
        return this.edgeHostname;
    }
    public String groupId() {
        return this.groupId;
    }
    public String hostname() {
        return this.hostname;
    }
    public String productionStatus() {
        return this.productionStatus;
    }
    public String propertyId() {
        return this.propertyId;
    }
    public String propertyName() {
        return this.propertyName;
    }
    public Integer propertyVersion() {
        return this.propertyVersion;
    }
    public String stagingStatus() {
        return this.stagingStatus;
    }
    public String updatedByUser() {
        return this.updatedByUser;
    }
    public String updatedDate() {
        return this.updatedDate;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPropertiesSearchProperty defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accountId;
        private String assetId;
        private String contractId;
        private String edgeHostname;
        private String groupId;
        private String hostname;
        private String productionStatus;
        private String propertyId;
        private String propertyName;
        private Integer propertyVersion;
        private String stagingStatus;
        private String updatedByUser;
        private String updatedDate;
        public Builder() {}
        public Builder(GetPropertiesSearchProperty defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accountId = defaults.accountId;
    	      this.assetId = defaults.assetId;
    	      this.contractId = defaults.contractId;
    	      this.edgeHostname = defaults.edgeHostname;
    	      this.groupId = defaults.groupId;
    	      this.hostname = defaults.hostname;
    	      this.productionStatus = defaults.productionStatus;
    	      this.propertyId = defaults.propertyId;
    	      this.propertyName = defaults.propertyName;
    	      this.propertyVersion = defaults.propertyVersion;
    	      this.stagingStatus = defaults.stagingStatus;
    	      this.updatedByUser = defaults.updatedByUser;
    	      this.updatedDate = defaults.updatedDate;
        }

        @CustomType.Setter
        public Builder accountId(String accountId) {
            this.accountId = Objects.requireNonNull(accountId);
            return this;
        }
        @CustomType.Setter
        public Builder assetId(String assetId) {
            this.assetId = Objects.requireNonNull(assetId);
            return this;
        }
        @CustomType.Setter
        public Builder contractId(String contractId) {
            this.contractId = Objects.requireNonNull(contractId);
            return this;
        }
        @CustomType.Setter
        public Builder edgeHostname(String edgeHostname) {
            this.edgeHostname = Objects.requireNonNull(edgeHostname);
            return this;
        }
        @CustomType.Setter
        public Builder groupId(String groupId) {
            this.groupId = Objects.requireNonNull(groupId);
            return this;
        }
        @CustomType.Setter
        public Builder hostname(String hostname) {
            this.hostname = Objects.requireNonNull(hostname);
            return this;
        }
        @CustomType.Setter
        public Builder productionStatus(String productionStatus) {
            this.productionStatus = Objects.requireNonNull(productionStatus);
            return this;
        }
        @CustomType.Setter
        public Builder propertyId(String propertyId) {
            this.propertyId = Objects.requireNonNull(propertyId);
            return this;
        }
        @CustomType.Setter
        public Builder propertyName(String propertyName) {
            this.propertyName = Objects.requireNonNull(propertyName);
            return this;
        }
        @CustomType.Setter
        public Builder propertyVersion(Integer propertyVersion) {
            this.propertyVersion = Objects.requireNonNull(propertyVersion);
            return this;
        }
        @CustomType.Setter
        public Builder stagingStatus(String stagingStatus) {
            this.stagingStatus = Objects.requireNonNull(stagingStatus);
            return this;
        }
        @CustomType.Setter
        public Builder updatedByUser(String updatedByUser) {
            this.updatedByUser = Objects.requireNonNull(updatedByUser);
            return this;
        }
        @CustomType.Setter
        public Builder updatedDate(String updatedDate) {
            this.updatedDate = Objects.requireNonNull(updatedDate);
            return this;
        }
        public GetPropertiesSearchProperty build() {
            final var o = new GetPropertiesSearchProperty();
            o.accountId = accountId;
            o.assetId = assetId;
            o.contractId = contractId;
            o.edgeHostname = edgeHostname;
            o.groupId = groupId;
            o.hostname = hostname;
            o.productionStatus = productionStatus;
            o.propertyId = propertyId;
            o.propertyName = propertyName;
            o.propertyVersion = propertyVersion;
            o.stagingStatus = stagingStatus;
            o.updatedByUser = updatedByUser;
            o.updatedDate = updatedDate;
            return o;
        }
    }
}
