// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.akamai.inputs.GetCloudletsPhasedReleaseMatchRuleMatchRule;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCloudletsPhasedReleaseMatchRulePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudletsPhasedReleaseMatchRulePlainArgs Empty = new GetCloudletsPhasedReleaseMatchRulePlainArgs();

    /**
     * - (Optional) A list of Cloudlet-specific match rules for a policy.
     * 
     */
    @Import(name="matchRules")
    private @Nullable List<GetCloudletsPhasedReleaseMatchRuleMatchRule> matchRules;

    /**
     * @return - (Optional) A list of Cloudlet-specific match rules for a policy.
     * 
     */
    public Optional<List<GetCloudletsPhasedReleaseMatchRuleMatchRule>> matchRules() {
        return Optional.ofNullable(this.matchRules);
    }

    private GetCloudletsPhasedReleaseMatchRulePlainArgs() {}

    private GetCloudletsPhasedReleaseMatchRulePlainArgs(GetCloudletsPhasedReleaseMatchRulePlainArgs $) {
        this.matchRules = $.matchRules;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudletsPhasedReleaseMatchRulePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudletsPhasedReleaseMatchRulePlainArgs $;

        public Builder() {
            $ = new GetCloudletsPhasedReleaseMatchRulePlainArgs();
        }

        public Builder(GetCloudletsPhasedReleaseMatchRulePlainArgs defaults) {
            $ = new GetCloudletsPhasedReleaseMatchRulePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param matchRules - (Optional) A list of Cloudlet-specific match rules for a policy.
         * 
         * @return builder
         * 
         */
        public Builder matchRules(@Nullable List<GetCloudletsPhasedReleaseMatchRuleMatchRule> matchRules) {
            $.matchRules = matchRules;
            return this;
        }

        /**
         * @param matchRules - (Optional) A list of Cloudlet-specific match rules for a policy.
         * 
         * @return builder
         * 
         */
        public Builder matchRules(GetCloudletsPhasedReleaseMatchRuleMatchRule... matchRules) {
            return matchRules(List.of(matchRules));
        }

        public GetCloudletsPhasedReleaseMatchRulePlainArgs build() {
            return $;
        }
    }

}
