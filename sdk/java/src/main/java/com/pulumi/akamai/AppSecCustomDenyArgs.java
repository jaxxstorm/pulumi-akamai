// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class AppSecCustomDenyArgs extends com.pulumi.resources.ResourceArgs {

    public static final AppSecCustomDenyArgs Empty = new AppSecCustomDenyArgs();

    /**
     * . Unique identifier of the security configuration associated with the custom deny.
     * 
     */
    @Import(name="configId", required=true)
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the custom deny.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }

    /**
     * . Path to a JSON file containing properties and property values for the custom deny.
     * 
     */
    @Import(name="customDeny", required=true)
    private Output<String> customDeny;

    /**
     * @return . Path to a JSON file containing properties and property values for the custom deny.
     * 
     */
    public Output<String> customDeny() {
        return this.customDeny;
    }

    private AppSecCustomDenyArgs() {}

    private AppSecCustomDenyArgs(AppSecCustomDenyArgs $) {
        this.configId = $.configId;
        this.customDeny = $.customDeny;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecCustomDenyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecCustomDenyArgs $;

        public Builder() {
            $ = new AppSecCustomDenyArgs();
        }

        public Builder(AppSecCustomDenyArgs defaults) {
            $ = new AppSecCustomDenyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the custom deny.
         * 
         * @return builder
         * 
         */
        public Builder configId(Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the custom deny.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param customDeny . Path to a JSON file containing properties and property values for the custom deny.
         * 
         * @return builder
         * 
         */
        public Builder customDeny(Output<String> customDeny) {
            $.customDeny = customDeny;
            return this;
        }

        /**
         * @param customDeny . Path to a JSON file containing properties and property values for the custom deny.
         * 
         * @return builder
         * 
         */
        public Builder customDeny(String customDeny) {
            return customDeny(Output.of(customDeny));
        }

        public AppSecCustomDenyArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            $.customDeny = Objects.requireNonNull($.customDeny, "expected parameter 'customDeny' to be non-null");
            return $;
        }
    }

}
