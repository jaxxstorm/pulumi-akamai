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
public final class GetAppSecRatePolicyActionsResult {
    private Integer configId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String outputText;
    private @Nullable Integer ratePolicyId;
    private String securityPolicyId;

    private GetAppSecRatePolicyActionsResult() {}
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
    public String outputText() {
        return this.outputText;
    }
    public Optional<Integer> ratePolicyId() {
        return Optional.ofNullable(this.ratePolicyId);
    }
    public String securityPolicyId() {
        return this.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAppSecRatePolicyActionsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer configId;
        private String id;
        private String outputText;
        private @Nullable Integer ratePolicyId;
        private String securityPolicyId;
        public Builder() {}
        public Builder(GetAppSecRatePolicyActionsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.configId = defaults.configId;
    	      this.id = defaults.id;
    	      this.outputText = defaults.outputText;
    	      this.ratePolicyId = defaults.ratePolicyId;
    	      this.securityPolicyId = defaults.securityPolicyId;
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
        public Builder outputText(String outputText) {
            this.outputText = Objects.requireNonNull(outputText);
            return this;
        }
        @CustomType.Setter
        public Builder ratePolicyId(@Nullable Integer ratePolicyId) {
            this.ratePolicyId = ratePolicyId;
            return this;
        }
        @CustomType.Setter
        public Builder securityPolicyId(String securityPolicyId) {
            this.securityPolicyId = Objects.requireNonNull(securityPolicyId);
            return this;
        }
        public GetAppSecRatePolicyActionsResult build() {
            final var o = new GetAppSecRatePolicyActionsResult();
            o.configId = configId;
            o.id = id;
            o.outputText = outputText;
            o.ratePolicyId = ratePolicyId;
            o.securityPolicyId = securityPolicyId;
            return o;
        }
    }
}
