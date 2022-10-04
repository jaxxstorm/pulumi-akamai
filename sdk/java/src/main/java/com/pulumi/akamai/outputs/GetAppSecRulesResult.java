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
public final class GetAppSecRulesResult {
    private String conditionException;
    private Integer configId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String json;
    private String outputText;
    private String ruleAction;
    private @Nullable Integer ruleId;
    private String securityPolicyId;

    private GetAppSecRulesResult() {}
    public String conditionException() {
        return this.conditionException;
    }
    public Integer configId() {
        return this.configId;
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
    public String ruleAction() {
        return this.ruleAction;
    }
    public Optional<Integer> ruleId() {
        return Optional.ofNullable(this.ruleId);
    }
    public String securityPolicyId() {
        return this.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAppSecRulesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String conditionException;
        private Integer configId;
        private String id;
        private String json;
        private String outputText;
        private String ruleAction;
        private @Nullable Integer ruleId;
        private String securityPolicyId;
        public Builder() {}
        public Builder(GetAppSecRulesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.conditionException = defaults.conditionException;
    	      this.configId = defaults.configId;
    	      this.id = defaults.id;
    	      this.json = defaults.json;
    	      this.outputText = defaults.outputText;
    	      this.ruleAction = defaults.ruleAction;
    	      this.ruleId = defaults.ruleId;
    	      this.securityPolicyId = defaults.securityPolicyId;
        }

        @CustomType.Setter
        public Builder conditionException(String conditionException) {
            this.conditionException = Objects.requireNonNull(conditionException);
            return this;
        }
        @CustomType.Setter
        public Builder configId(Integer configId) {
            this.configId = Objects.requireNonNull(configId);
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
        @CustomType.Setter
        public Builder ruleAction(String ruleAction) {
            this.ruleAction = Objects.requireNonNull(ruleAction);
            return this;
        }
        @CustomType.Setter
        public Builder ruleId(@Nullable Integer ruleId) {
            this.ruleId = ruleId;
            return this;
        }
        @CustomType.Setter
        public Builder securityPolicyId(String securityPolicyId) {
            this.securityPolicyId = Objects.requireNonNull(securityPolicyId);
            return this;
        }
        public GetAppSecRulesResult build() {
            final var o = new GetAppSecRulesResult();
            o.conditionException = conditionException;
            o.configId = configId;
            o.id = id;
            o.json = json;
            o.outputText = outputText;
            o.ruleAction = ruleAction;
            o.ruleId = ruleId;
            o.securityPolicyId = securityPolicyId;
            return o;
        }
    }
}
