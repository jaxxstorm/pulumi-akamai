// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAppSecSecurityPolicyResult {
    private Integer configId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String outputText;
    private String securityPolicyId;
    private List<String> securityPolicyIdLists;
    private @Nullable String securityPolicyName;

    private GetAppSecSecurityPolicyResult() {}
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
    public String securityPolicyId() {
        return this.securityPolicyId;
    }
    public List<String> securityPolicyIdLists() {
        return this.securityPolicyIdLists;
    }
    public Optional<String> securityPolicyName() {
        return Optional.ofNullable(this.securityPolicyName);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAppSecSecurityPolicyResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer configId;
        private String id;
        private String outputText;
        private String securityPolicyId;
        private List<String> securityPolicyIdLists;
        private @Nullable String securityPolicyName;
        public Builder() {}
        public Builder(GetAppSecSecurityPolicyResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.configId = defaults.configId;
    	      this.id = defaults.id;
    	      this.outputText = defaults.outputText;
    	      this.securityPolicyId = defaults.securityPolicyId;
    	      this.securityPolicyIdLists = defaults.securityPolicyIdLists;
    	      this.securityPolicyName = defaults.securityPolicyName;
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
        public Builder securityPolicyId(String securityPolicyId) {
            this.securityPolicyId = Objects.requireNonNull(securityPolicyId);
            return this;
        }
        @CustomType.Setter
        public Builder securityPolicyIdLists(List<String> securityPolicyIdLists) {
            this.securityPolicyIdLists = Objects.requireNonNull(securityPolicyIdLists);
            return this;
        }
        public Builder securityPolicyIdLists(String... securityPolicyIdLists) {
            return securityPolicyIdLists(List.of(securityPolicyIdLists));
        }
        @CustomType.Setter
        public Builder securityPolicyName(@Nullable String securityPolicyName) {
            this.securityPolicyName = securityPolicyName;
            return this;
        }
        public GetAppSecSecurityPolicyResult build() {
            final var o = new GetAppSecSecurityPolicyResult();
            o.configId = configId;
            o.id = id;
            o.outputText = outputText;
            o.securityPolicyId = securityPolicyId;
            o.securityPolicyIdLists = securityPolicyIdLists;
            o.securityPolicyName = securityPolicyName;
            return o;
        }
    }
}
