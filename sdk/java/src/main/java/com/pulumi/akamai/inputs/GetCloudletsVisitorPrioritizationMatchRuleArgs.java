// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.akamai.inputs.GetCloudletsVisitorPrioritizationMatchRuleMatchRuleArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCloudletsVisitorPrioritizationMatchRuleArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudletsVisitorPrioritizationMatchRuleArgs Empty = new GetCloudletsVisitorPrioritizationMatchRuleArgs();

    /**
     * - (Optional) A list of Cloudlet-specific match rules for a policy.
     * 
     */
    @Import(name="matchRules")
    private @Nullable Output<List<GetCloudletsVisitorPrioritizationMatchRuleMatchRuleArgs>> matchRules;

    /**
     * @return - (Optional) A list of Cloudlet-specific match rules for a policy.
     * 
     */
    public Optional<Output<List<GetCloudletsVisitorPrioritizationMatchRuleMatchRuleArgs>>> matchRules() {
        return Optional.ofNullable(this.matchRules);
    }

    private GetCloudletsVisitorPrioritizationMatchRuleArgs() {}

    private GetCloudletsVisitorPrioritizationMatchRuleArgs(GetCloudletsVisitorPrioritizationMatchRuleArgs $) {
        this.matchRules = $.matchRules;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudletsVisitorPrioritizationMatchRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudletsVisitorPrioritizationMatchRuleArgs $;

        public Builder() {
            $ = new GetCloudletsVisitorPrioritizationMatchRuleArgs();
        }

        public Builder(GetCloudletsVisitorPrioritizationMatchRuleArgs defaults) {
            $ = new GetCloudletsVisitorPrioritizationMatchRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param matchRules - (Optional) A list of Cloudlet-specific match rules for a policy.
         * 
         * @return builder
         * 
         */
        public Builder matchRules(@Nullable Output<List<GetCloudletsVisitorPrioritizationMatchRuleMatchRuleArgs>> matchRules) {
            $.matchRules = matchRules;
            return this;
        }

        /**
         * @param matchRules - (Optional) A list of Cloudlet-specific match rules for a policy.
         * 
         * @return builder
         * 
         */
        public Builder matchRules(List<GetCloudletsVisitorPrioritizationMatchRuleMatchRuleArgs> matchRules) {
            return matchRules(Output.of(matchRules));
        }

        /**
         * @param matchRules - (Optional) A list of Cloudlet-specific match rules for a policy.
         * 
         * @return builder
         * 
         */
        public Builder matchRules(GetCloudletsVisitorPrioritizationMatchRuleMatchRuleArgs... matchRules) {
            return matchRules(List.of(matchRules));
        }

        public GetCloudletsVisitorPrioritizationMatchRuleArgs build() {
            return $;
        }
    }

}
