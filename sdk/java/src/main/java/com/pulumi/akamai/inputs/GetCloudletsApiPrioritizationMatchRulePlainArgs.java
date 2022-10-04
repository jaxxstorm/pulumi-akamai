// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.akamai.inputs.GetCloudletsApiPrioritizationMatchRuleMatchRule;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCloudletsApiPrioritizationMatchRulePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudletsApiPrioritizationMatchRulePlainArgs Empty = new GetCloudletsApiPrioritizationMatchRulePlainArgs();

    /**
     * - (Optional) A list of Cloudlet-specific match rules for a policy.
     * 
     */
    @Import(name="matchRules")
    private @Nullable List<GetCloudletsApiPrioritizationMatchRuleMatchRule> matchRules;

    /**
     * @return - (Optional) A list of Cloudlet-specific match rules for a policy.
     * 
     */
    public Optional<List<GetCloudletsApiPrioritizationMatchRuleMatchRule>> matchRules() {
        return Optional.ofNullable(this.matchRules);
    }

    private GetCloudletsApiPrioritizationMatchRulePlainArgs() {}

    private GetCloudletsApiPrioritizationMatchRulePlainArgs(GetCloudletsApiPrioritizationMatchRulePlainArgs $) {
        this.matchRules = $.matchRules;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudletsApiPrioritizationMatchRulePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudletsApiPrioritizationMatchRulePlainArgs $;

        public Builder() {
            $ = new GetCloudletsApiPrioritizationMatchRulePlainArgs();
        }

        public Builder(GetCloudletsApiPrioritizationMatchRulePlainArgs defaults) {
            $ = new GetCloudletsApiPrioritizationMatchRulePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param matchRules - (Optional) A list of Cloudlet-specific match rules for a policy.
         * 
         * @return builder
         * 
         */
        public Builder matchRules(@Nullable List<GetCloudletsApiPrioritizationMatchRuleMatchRule> matchRules) {
            $.matchRules = matchRules;
            return this;
        }

        /**
         * @param matchRules - (Optional) A list of Cloudlet-specific match rules for a policy.
         * 
         * @return builder
         * 
         */
        public Builder matchRules(GetCloudletsApiPrioritizationMatchRuleMatchRule... matchRules) {
            return matchRules(List.of(matchRules));
        }

        public GetCloudletsApiPrioritizationMatchRulePlainArgs build() {
            return $;
        }
    }

}
