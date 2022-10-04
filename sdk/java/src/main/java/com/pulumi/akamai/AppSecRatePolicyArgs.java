// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class AppSecRatePolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final AppSecRatePolicyArgs Empty = new AppSecRatePolicyArgs();

    /**
     * . Unique identifier of the security configuration associated with the rate policy being modified.
     * 
     */
    @Import(name="configId", required=true)
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the rate policy being modified.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }

    /**
     * . Path to a JSON file containing a rate policy definition.
     * 
     */
    @Import(name="ratePolicy", required=true)
    private Output<String> ratePolicy;

    /**
     * @return . Path to a JSON file containing a rate policy definition.
     * 
     */
    public Output<String> ratePolicy() {
        return this.ratePolicy;
    }

    private AppSecRatePolicyArgs() {}

    private AppSecRatePolicyArgs(AppSecRatePolicyArgs $) {
        this.configId = $.configId;
        this.ratePolicy = $.ratePolicy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecRatePolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecRatePolicyArgs $;

        public Builder() {
            $ = new AppSecRatePolicyArgs();
        }

        public Builder(AppSecRatePolicyArgs defaults) {
            $ = new AppSecRatePolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the rate policy being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the rate policy being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param ratePolicy . Path to a JSON file containing a rate policy definition.
         * 
         * @return builder
         * 
         */
        public Builder ratePolicy(Output<String> ratePolicy) {
            $.ratePolicy = ratePolicy;
            return this;
        }

        /**
         * @param ratePolicy . Path to a JSON file containing a rate policy definition.
         * 
         * @return builder
         * 
         */
        public Builder ratePolicy(String ratePolicy) {
            return ratePolicy(Output.of(ratePolicy));
        }

        public AppSecRatePolicyArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            $.ratePolicy = Objects.requireNonNull($.ratePolicy, "expected parameter 'ratePolicy' to be non-null");
            return $;
        }
    }

}
