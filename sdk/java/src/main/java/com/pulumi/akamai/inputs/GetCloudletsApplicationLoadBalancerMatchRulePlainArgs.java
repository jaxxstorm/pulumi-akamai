// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.akamai.inputs.GetCloudletsApplicationLoadBalancerMatchRuleMatchRule;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCloudletsApplicationLoadBalancerMatchRulePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudletsApplicationLoadBalancerMatchRulePlainArgs Empty = new GetCloudletsApplicationLoadBalancerMatchRulePlainArgs();

    /**
     * - (Optional) A list of Cloudlet-specific match rules for a policy.
     * 
     */
    @Import(name="matchRules")
    private @Nullable List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRule> matchRules;

    /**
     * @return - (Optional) A list of Cloudlet-specific match rules for a policy.
     * 
     */
    public Optional<List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRule>> matchRules() {
        return Optional.ofNullable(this.matchRules);
    }

    private GetCloudletsApplicationLoadBalancerMatchRulePlainArgs() {}

    private GetCloudletsApplicationLoadBalancerMatchRulePlainArgs(GetCloudletsApplicationLoadBalancerMatchRulePlainArgs $) {
        this.matchRules = $.matchRules;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudletsApplicationLoadBalancerMatchRulePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudletsApplicationLoadBalancerMatchRulePlainArgs $;

        public Builder() {
            $ = new GetCloudletsApplicationLoadBalancerMatchRulePlainArgs();
        }

        public Builder(GetCloudletsApplicationLoadBalancerMatchRulePlainArgs defaults) {
            $ = new GetCloudletsApplicationLoadBalancerMatchRulePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param matchRules - (Optional) A list of Cloudlet-specific match rules for a policy.
         * 
         * @return builder
         * 
         */
        public Builder matchRules(@Nullable List<GetCloudletsApplicationLoadBalancerMatchRuleMatchRule> matchRules) {
            $.matchRules = matchRules;
            return this;
        }

        /**
         * @param matchRules - (Optional) A list of Cloudlet-specific match rules for a policy.
         * 
         * @return builder
         * 
         */
        public Builder matchRules(GetCloudletsApplicationLoadBalancerMatchRuleMatchRule... matchRules) {
            return matchRules(List.of(matchRules));
        }

        public GetCloudletsApplicationLoadBalancerMatchRulePlainArgs build() {
            return $;
        }
    }

}
