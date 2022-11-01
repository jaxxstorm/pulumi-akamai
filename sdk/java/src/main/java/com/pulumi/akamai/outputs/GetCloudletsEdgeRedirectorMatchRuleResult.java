// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.akamai.outputs.GetCloudletsEdgeRedirectorMatchRuleMatchRule;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetCloudletsEdgeRedirectorMatchRuleResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String json;
    private @Nullable List<GetCloudletsEdgeRedirectorMatchRuleMatchRule> matchRules;

    private GetCloudletsEdgeRedirectorMatchRuleResult() {}
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
    public List<GetCloudletsEdgeRedirectorMatchRuleMatchRule> matchRules() {
        return this.matchRules == null ? List.of() : this.matchRules;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCloudletsEdgeRedirectorMatchRuleResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String json;
        private @Nullable List<GetCloudletsEdgeRedirectorMatchRuleMatchRule> matchRules;
        public Builder() {}
        public Builder(GetCloudletsEdgeRedirectorMatchRuleResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.json = defaults.json;
    	      this.matchRules = defaults.matchRules;
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
        public Builder matchRules(@Nullable List<GetCloudletsEdgeRedirectorMatchRuleMatchRule> matchRules) {
            this.matchRules = matchRules;
            return this;
        }
        public Builder matchRules(GetCloudletsEdgeRedirectorMatchRuleMatchRule... matchRules) {
            return matchRules(List.of(matchRules));
        }
        public GetCloudletsEdgeRedirectorMatchRuleResult build() {
            final var o = new GetCloudletsEdgeRedirectorMatchRuleResult();
            o.id = id;
            o.json = json;
            o.matchRules = matchRules;
            return o;
        }
    }
}