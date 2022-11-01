// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.akamai.outputs.GetCloudletsApiPrioritizationMatchRuleMatchRule;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetCloudletsApiPrioritizationMatchRuleResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String json;
    private @Nullable List<GetCloudletsApiPrioritizationMatchRuleMatchRule> matchRules;

    private GetCloudletsApiPrioritizationMatchRuleResult() {}
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
    public List<GetCloudletsApiPrioritizationMatchRuleMatchRule> matchRules() {
        return this.matchRules == null ? List.of() : this.matchRules;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCloudletsApiPrioritizationMatchRuleResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String json;
        private @Nullable List<GetCloudletsApiPrioritizationMatchRuleMatchRule> matchRules;
        public Builder() {}
        public Builder(GetCloudletsApiPrioritizationMatchRuleResult defaults) {
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
        public Builder matchRules(@Nullable List<GetCloudletsApiPrioritizationMatchRuleMatchRule> matchRules) {
            this.matchRules = matchRules;
            return this;
        }
        public Builder matchRules(GetCloudletsApiPrioritizationMatchRuleMatchRule... matchRules) {
            return matchRules(List.of(matchRules));
        }
        public GetCloudletsApiPrioritizationMatchRuleResult build() {
            final var o = new GetCloudletsApiPrioritizationMatchRuleResult();
            o.id = id;
            o.json = json;
            o.matchRules = matchRules;
            return o;
        }
    }
}