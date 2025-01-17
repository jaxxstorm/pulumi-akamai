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


public final class AppSecApiRequestConstraintsState extends com.pulumi.resources.ResourceArgs {

    public static final AppSecApiRequestConstraintsState Empty = new AppSecApiRequestConstraintsState();

    /**
     * . Action to assign to the API request constraint. Allowed values are:
     * - **alert**, Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     * 
     */
    @Import(name="action")
    private @Nullable Output<String> action;

    /**
     * @return . Action to assign to the API request constraint. Allowed values are:
     * - **alert**, Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     * 
     */
    public Optional<Output<String>> action() {
        return Optional.ofNullable(this.action);
    }

    /**
     * . ID of the API endpoint the constraint will be assigned to.
     * 
     */
    @Import(name="apiEndpointId")
    private @Nullable Output<Integer> apiEndpointId;

    /**
     * @return . ID of the API endpoint the constraint will be assigned to.
     * 
     */
    public Optional<Output<Integer>> apiEndpointId() {
        return Optional.ofNullable(this.apiEndpointId);
    }

    /**
     * . Unique identifier of the security configuration associated with the API request constraint settings being modified.
     * 
     */
    @Import(name="configId")
    private @Nullable Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the API request constraint settings being modified.
     * 
     */
    public Optional<Output<Integer>> configId() {
        return Optional.ofNullable(this.configId);
    }

    /**
     * . Unique identifier of the security policy associated with the API request constraint settings being modified.
     * 
     */
    @Import(name="securityPolicyId")
    private @Nullable Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the API request constraint settings being modified.
     * 
     */
    public Optional<Output<String>> securityPolicyId() {
        return Optional.ofNullable(this.securityPolicyId);
    }

    private AppSecApiRequestConstraintsState() {}

    private AppSecApiRequestConstraintsState(AppSecApiRequestConstraintsState $) {
        this.action = $.action;
        this.apiEndpointId = $.apiEndpointId;
        this.configId = $.configId;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecApiRequestConstraintsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecApiRequestConstraintsState $;

        public Builder() {
            $ = new AppSecApiRequestConstraintsState();
        }

        public Builder(AppSecApiRequestConstraintsState defaults) {
            $ = new AppSecApiRequestConstraintsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param action . Action to assign to the API request constraint. Allowed values are:
         * - **alert**, Record the event.
         * - **deny**. Block the request.
         * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
         * - **none**. Take no action.
         * 
         * @return builder
         * 
         */
        public Builder action(@Nullable Output<String> action) {
            $.action = action;
            return this;
        }

        /**
         * @param action . Action to assign to the API request constraint. Allowed values are:
         * - **alert**, Record the event.
         * - **deny**. Block the request.
         * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
         * - **none**. Take no action.
         * 
         * @return builder
         * 
         */
        public Builder action(String action) {
            return action(Output.of(action));
        }

        /**
         * @param apiEndpointId . ID of the API endpoint the constraint will be assigned to.
         * 
         * @return builder
         * 
         */
        public Builder apiEndpointId(@Nullable Output<Integer> apiEndpointId) {
            $.apiEndpointId = apiEndpointId;
            return this;
        }

        /**
         * @param apiEndpointId . ID of the API endpoint the constraint will be assigned to.
         * 
         * @return builder
         * 
         */
        public Builder apiEndpointId(Integer apiEndpointId) {
            return apiEndpointId(Output.of(apiEndpointId));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the API request constraint settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(@Nullable Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the API request constraint settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the API request constraint settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(@Nullable Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the API request constraint settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        public AppSecApiRequestConstraintsState build() {
            return $;
        }
    }

}
