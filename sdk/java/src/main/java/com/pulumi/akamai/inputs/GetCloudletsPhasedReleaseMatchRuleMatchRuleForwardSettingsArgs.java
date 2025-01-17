// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsArgs Empty = new GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsArgs();

    /**
     * - (Required) The ID of the new origin requests are forwarded to. This type of origin is known as a Conditional Origin. See Property requirements for Cloudlets that forward requests to learn more.
     * 
     */
    @Import(name="originId", required=true)
    private Output<String> originId;

    /**
     * @return - (Required) The ID of the new origin requests are forwarded to. This type of origin is known as a Conditional Origin. See Property requirements for Cloudlets that forward requests to learn more.
     * 
     */
    public Output<String> originId() {
        return this.originId;
    }

    /**
     * - (Required)
     * 
     */
    @Import(name="percent", required=true)
    private Output<Integer> percent;

    /**
     * @return - (Required)
     * 
     */
    public Output<Integer> percent() {
        return this.percent;
    }

    private GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsArgs() {}

    private GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsArgs(GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsArgs $) {
        this.originId = $.originId;
        this.percent = $.percent;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsArgs $;

        public Builder() {
            $ = new GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsArgs();
        }

        public Builder(GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsArgs defaults) {
            $ = new GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param originId - (Required) The ID of the new origin requests are forwarded to. This type of origin is known as a Conditional Origin. See Property requirements for Cloudlets that forward requests to learn more.
         * 
         * @return builder
         * 
         */
        public Builder originId(Output<String> originId) {
            $.originId = originId;
            return this;
        }

        /**
         * @param originId - (Required) The ID of the new origin requests are forwarded to. This type of origin is known as a Conditional Origin. See Property requirements for Cloudlets that forward requests to learn more.
         * 
         * @return builder
         * 
         */
        public Builder originId(String originId) {
            return originId(Output.of(originId));
        }

        /**
         * @param percent - (Required)
         * 
         * @return builder
         * 
         */
        public Builder percent(Output<Integer> percent) {
            $.percent = percent;
            return this;
        }

        /**
         * @param percent - (Required)
         * 
         * @return builder
         * 
         */
        public Builder percent(Integer percent) {
            return percent(Output.of(percent));
        }

        public GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsArgs build() {
            $.originId = Objects.requireNonNull($.originId, "expected parameter 'originId' to be non-null");
            $.percent = Objects.requireNonNull($.percent, "expected parameter 'percent' to be non-null");
            return $;
        }
    }

}
