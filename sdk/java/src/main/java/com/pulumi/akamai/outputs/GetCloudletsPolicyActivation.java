// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.akamai.outputs.GetCloudletsPolicyActivationPolicyInfo;
import com.pulumi.akamai.outputs.GetCloudletsPolicyActivationPropertyInfo;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetCloudletsPolicyActivation {
    private String apiVersion;
    private String network;
    private List<GetCloudletsPolicyActivationPolicyInfo> policyInfos;
    private List<GetCloudletsPolicyActivationPropertyInfo> propertyInfos;

    private GetCloudletsPolicyActivation() {}
    public String apiVersion() {
        return this.apiVersion;
    }
    public String network() {
        return this.network;
    }
    public List<GetCloudletsPolicyActivationPolicyInfo> policyInfos() {
        return this.policyInfos;
    }
    public List<GetCloudletsPolicyActivationPropertyInfo> propertyInfos() {
        return this.propertyInfos;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCloudletsPolicyActivation defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String apiVersion;
        private String network;
        private List<GetCloudletsPolicyActivationPolicyInfo> policyInfos;
        private List<GetCloudletsPolicyActivationPropertyInfo> propertyInfos;
        public Builder() {}
        public Builder(GetCloudletsPolicyActivation defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apiVersion = defaults.apiVersion;
    	      this.network = defaults.network;
    	      this.policyInfos = defaults.policyInfos;
    	      this.propertyInfos = defaults.propertyInfos;
        }

        @CustomType.Setter
        public Builder apiVersion(String apiVersion) {
            this.apiVersion = Objects.requireNonNull(apiVersion);
            return this;
        }
        @CustomType.Setter
        public Builder network(String network) {
            this.network = Objects.requireNonNull(network);
            return this;
        }
        @CustomType.Setter
        public Builder policyInfos(List<GetCloudletsPolicyActivationPolicyInfo> policyInfos) {
            this.policyInfos = Objects.requireNonNull(policyInfos);
            return this;
        }
        public Builder policyInfos(GetCloudletsPolicyActivationPolicyInfo... policyInfos) {
            return policyInfos(List.of(policyInfos));
        }
        @CustomType.Setter
        public Builder propertyInfos(List<GetCloudletsPolicyActivationPropertyInfo> propertyInfos) {
            this.propertyInfos = Objects.requireNonNull(propertyInfos);
            return this;
        }
        public Builder propertyInfos(GetCloudletsPolicyActivationPropertyInfo... propertyInfos) {
            return propertyInfos(List.of(propertyInfos));
        }
        public GetCloudletsPolicyActivation build() {
            final var o = new GetCloudletsPolicyActivation();
            o.apiVersion = apiVersion;
            o.network = network;
            o.policyInfos = policyInfos;
            o.propertyInfos = propertyInfos;
            return o;
        }
    }
}
