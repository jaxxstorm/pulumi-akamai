// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAppSecContractsGroupsResult {
    private @Nullable String contractid;
    private String defaultContractid;
    private Integer defaultGroupid;
    private @Nullable Integer groupid;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String json;
    private String outputText;

    private GetAppSecContractsGroupsResult() {}
    public Optional<String> contractid() {
        return Optional.ofNullable(this.contractid);
    }
    public String defaultContractid() {
        return this.defaultContractid;
    }
    public Integer defaultGroupid() {
        return this.defaultGroupid;
    }
    public Optional<Integer> groupid() {
        return Optional.ofNullable(this.groupid);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String json() {
        return this.json;
    }
    public String outputText() {
        return this.outputText;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAppSecContractsGroupsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String contractid;
        private String defaultContractid;
        private Integer defaultGroupid;
        private @Nullable Integer groupid;
        private String id;
        private String json;
        private String outputText;
        public Builder() {}
        public Builder(GetAppSecContractsGroupsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.contractid = defaults.contractid;
    	      this.defaultContractid = defaults.defaultContractid;
    	      this.defaultGroupid = defaults.defaultGroupid;
    	      this.groupid = defaults.groupid;
    	      this.id = defaults.id;
    	      this.json = defaults.json;
    	      this.outputText = defaults.outputText;
        }

        @CustomType.Setter
        public Builder contractid(@Nullable String contractid) {
            this.contractid = contractid;
            return this;
        }
        @CustomType.Setter
        public Builder defaultContractid(String defaultContractid) {
            this.defaultContractid = Objects.requireNonNull(defaultContractid);
            return this;
        }
        @CustomType.Setter
        public Builder defaultGroupid(Integer defaultGroupid) {
            this.defaultGroupid = Objects.requireNonNull(defaultGroupid);
            return this;
        }
        @CustomType.Setter
        public Builder groupid(@Nullable Integer groupid) {
            this.groupid = groupid;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder json(String json) {
            this.json = Objects.requireNonNull(json);
            return this;
        }
        @CustomType.Setter
        public Builder outputText(String outputText) {
            this.outputText = Objects.requireNonNull(outputText);
            return this;
        }
        public GetAppSecContractsGroupsResult build() {
            final var o = new GetAppSecContractsGroupsResult();
            o.contractid = contractid;
            o.defaultContractid = defaultContractid;
            o.defaultGroupid = defaultGroupid;
            o.groupid = groupid;
            o.id = id;
            o.json = json;
            o.outputText = outputText;
            return o;
        }
    }
}
