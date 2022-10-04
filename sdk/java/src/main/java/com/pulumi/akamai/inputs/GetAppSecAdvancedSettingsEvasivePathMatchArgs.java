// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAppSecAdvancedSettingsEvasivePathMatchArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecAdvancedSettingsEvasivePathMatchArgs Empty = new GetAppSecAdvancedSettingsEvasivePathMatchArgs();

    /**
     * The configuration ID.
     * 
     */
    @Import(name="configId", required=true)
    private Output<Integer> configId;

    /**
     * @return The configuration ID.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }

    /**
     * The ID of the security policy to use.
     * 
     */
    @Import(name="securityPolicyId")
    private @Nullable Output<String> securityPolicyId;

    /**
     * @return The ID of the security policy to use.
     * 
     */
    public Optional<Output<String>> securityPolicyId() {
        return Optional.ofNullable(this.securityPolicyId);
    }

    private GetAppSecAdvancedSettingsEvasivePathMatchArgs() {}

    private GetAppSecAdvancedSettingsEvasivePathMatchArgs(GetAppSecAdvancedSettingsEvasivePathMatchArgs $) {
        this.configId = $.configId;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecAdvancedSettingsEvasivePathMatchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecAdvancedSettingsEvasivePathMatchArgs $;

        public Builder() {
            $ = new GetAppSecAdvancedSettingsEvasivePathMatchArgs();
        }

        public Builder(GetAppSecAdvancedSettingsEvasivePathMatchArgs defaults) {
            $ = new GetAppSecAdvancedSettingsEvasivePathMatchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId The configuration ID.
         * 
         * @return builder
         * 
         */
        public Builder configId(Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId The configuration ID.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param securityPolicyId The ID of the security policy to use.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(@Nullable Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId The ID of the security policy to use.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        public GetAppSecAdvancedSettingsEvasivePathMatchArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            return $;
        }
    }

}
