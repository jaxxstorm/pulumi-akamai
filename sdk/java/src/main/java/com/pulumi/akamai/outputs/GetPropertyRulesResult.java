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
public final class GetPropertyRulesResult {
    private String contractId;
    private String errors;
    private String groupId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String propertyId;
    private @Nullable String ruleFormat;
    private String rules;
    private Integer version;

    private GetPropertyRulesResult() {}
    public String contractId() {
        return this.contractId;
    }
    public String errors() {
        return this.errors;
    }
    public String groupId() {
        return this.groupId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String propertyId() {
        return this.propertyId;
    }
    public Optional<String> ruleFormat() {
        return Optional.ofNullable(this.ruleFormat);
    }
    public String rules() {
        return this.rules;
    }
    public Integer version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPropertyRulesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String contractId;
        private String errors;
        private String groupId;
        private String id;
        private String propertyId;
        private @Nullable String ruleFormat;
        private String rules;
        private Integer version;
        public Builder() {}
        public Builder(GetPropertyRulesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.contractId = defaults.contractId;
    	      this.errors = defaults.errors;
    	      this.groupId = defaults.groupId;
    	      this.id = defaults.id;
    	      this.propertyId = defaults.propertyId;
    	      this.ruleFormat = defaults.ruleFormat;
    	      this.rules = defaults.rules;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder contractId(String contractId) {
            this.contractId = Objects.requireNonNull(contractId);
            return this;
        }
        @CustomType.Setter
        public Builder errors(String errors) {
            this.errors = Objects.requireNonNull(errors);
            return this;
        }
        @CustomType.Setter
        public Builder groupId(String groupId) {
            this.groupId = Objects.requireNonNull(groupId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder propertyId(String propertyId) {
            this.propertyId = Objects.requireNonNull(propertyId);
            return this;
        }
        @CustomType.Setter
        public Builder ruleFormat(@Nullable String ruleFormat) {
            this.ruleFormat = ruleFormat;
            return this;
        }
        @CustomType.Setter
        public Builder rules(String rules) {
            this.rules = Objects.requireNonNull(rules);
            return this;
        }
        @CustomType.Setter
        public Builder version(Integer version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }
        public GetPropertyRulesResult build() {
            final var o = new GetPropertyRulesResult();
            o.contractId = contractId;
            o.errors = errors;
            o.groupId = groupId;
            o.id = id;
            o.propertyId = propertyId;
            o.ruleFormat = ruleFormat;
            o.rules = rules;
            o.version = version;
            return o;
        }
    }
}
