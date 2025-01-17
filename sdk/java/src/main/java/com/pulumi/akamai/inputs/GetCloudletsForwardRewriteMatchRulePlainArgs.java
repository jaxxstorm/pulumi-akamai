// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.akamai.inputs.GetCloudletsForwardRewriteMatchRuleMatchRule;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCloudletsForwardRewriteMatchRulePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudletsForwardRewriteMatchRulePlainArgs Empty = new GetCloudletsForwardRewriteMatchRulePlainArgs();

    /**
     * - (Optional) A list of Cloudlet-specific match rules for a policy.
     * 
     */
    @Import(name="matchRules")
    private @Nullable List<GetCloudletsForwardRewriteMatchRuleMatchRule> matchRules;

    /**
     * @return - (Optional) A list of Cloudlet-specific match rules for a policy.
     * 
     */
    public Optional<List<GetCloudletsForwardRewriteMatchRuleMatchRule>> matchRules() {
        return Optional.ofNullable(this.matchRules);
    }

    private GetCloudletsForwardRewriteMatchRulePlainArgs() {}

    private GetCloudletsForwardRewriteMatchRulePlainArgs(GetCloudletsForwardRewriteMatchRulePlainArgs $) {
        this.matchRules = $.matchRules;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudletsForwardRewriteMatchRulePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudletsForwardRewriteMatchRulePlainArgs $;

        public Builder() {
            $ = new GetCloudletsForwardRewriteMatchRulePlainArgs();
        }

        public Builder(GetCloudletsForwardRewriteMatchRulePlainArgs defaults) {
            $ = new GetCloudletsForwardRewriteMatchRulePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param matchRules - (Optional) A list of Cloudlet-specific match rules for a policy.
         * 
         * @return builder
         * 
         */
        public Builder matchRules(@Nullable List<GetCloudletsForwardRewriteMatchRuleMatchRule> matchRules) {
            $.matchRules = matchRules;
            return this;
        }

        /**
         * @param matchRules - (Optional) A list of Cloudlet-specific match rules for a policy.
         * 
         * @return builder
         * 
         */
        public Builder matchRules(GetCloudletsForwardRewriteMatchRuleMatchRule... matchRules) {
            return matchRules(List.of(matchRules));
        }

        public GetCloudletsForwardRewriteMatchRulePlainArgs build() {
            return $;
        }
    }

}
