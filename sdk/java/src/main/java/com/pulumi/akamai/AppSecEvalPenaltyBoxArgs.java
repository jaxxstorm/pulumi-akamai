// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class AppSecEvalPenaltyBoxArgs extends com.pulumi.resources.ResourceArgs {

    public static final AppSecEvalPenaltyBoxArgs Empty = new AppSecEvalPenaltyBoxArgs();

    /**
     * . Unique identifier of the security configuration associated with the evaluation penalty box settings being modified.
     * 
     */
    @Import(name="configId", required=true)
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the evaluation penalty box settings being modified.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }

    /**
     * . Action taken any time evaluation penalty box protection is triggered. Allowed values are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     * 
     */
    @Import(name="penaltyBoxAction", required=true)
    private Output<String> penaltyBoxAction;

    /**
     * @return . Action taken any time evaluation penalty box protection is triggered. Allowed values are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     * 
     */
    public Output<String> penaltyBoxAction() {
        return this.penaltyBoxAction;
    }

    /**
     * . Set to **true** to enable evaluation penalty box protection; set to **false** to disable evaluation penalty box protection.
     * 
     */
    @Import(name="penaltyBoxProtection", required=true)
    private Output<Boolean> penaltyBoxProtection;

    /**
     * @return . Set to **true** to enable evaluation penalty box protection; set to **false** to disable evaluation penalty box protection.
     * 
     */
    public Output<Boolean> penaltyBoxProtection() {
        return this.penaltyBoxProtection;
    }

    /**
     * . Unique identifier of the security policy associated with the evaluation penalty box settings being modified.
     * 
     */
    @Import(name="securityPolicyId", required=true)
    private Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the evaluation penalty box settings being modified.
     * 
     */
    public Output<String> securityPolicyId() {
        return this.securityPolicyId;
    }

    private AppSecEvalPenaltyBoxArgs() {}

    private AppSecEvalPenaltyBoxArgs(AppSecEvalPenaltyBoxArgs $) {
        this.configId = $.configId;
        this.penaltyBoxAction = $.penaltyBoxAction;
        this.penaltyBoxProtection = $.penaltyBoxProtection;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecEvalPenaltyBoxArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecEvalPenaltyBoxArgs $;

        public Builder() {
            $ = new AppSecEvalPenaltyBoxArgs();
        }

        public Builder(AppSecEvalPenaltyBoxArgs defaults) {
            $ = new AppSecEvalPenaltyBoxArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the evaluation penalty box settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the evaluation penalty box settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param penaltyBoxAction . Action taken any time evaluation penalty box protection is triggered. Allowed values are:
         * - **alert**. Record the event.
         * - **deny**. Block the request.
         * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
         * - **none**. Take no action.
         * 
         * @return builder
         * 
         */
        public Builder penaltyBoxAction(Output<String> penaltyBoxAction) {
            $.penaltyBoxAction = penaltyBoxAction;
            return this;
        }

        /**
         * @param penaltyBoxAction . Action taken any time evaluation penalty box protection is triggered. Allowed values are:
         * - **alert**. Record the event.
         * - **deny**. Block the request.
         * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
         * - **none**. Take no action.
         * 
         * @return builder
         * 
         */
        public Builder penaltyBoxAction(String penaltyBoxAction) {
            return penaltyBoxAction(Output.of(penaltyBoxAction));
        }

        /**
         * @param penaltyBoxProtection . Set to **true** to enable evaluation penalty box protection; set to **false** to disable evaluation penalty box protection.
         * 
         * @return builder
         * 
         */
        public Builder penaltyBoxProtection(Output<Boolean> penaltyBoxProtection) {
            $.penaltyBoxProtection = penaltyBoxProtection;
            return this;
        }

        /**
         * @param penaltyBoxProtection . Set to **true** to enable evaluation penalty box protection; set to **false** to disable evaluation penalty box protection.
         * 
         * @return builder
         * 
         */
        public Builder penaltyBoxProtection(Boolean penaltyBoxProtection) {
            return penaltyBoxProtection(Output.of(penaltyBoxProtection));
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the evaluation penalty box settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the evaluation penalty box settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        public AppSecEvalPenaltyBoxArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            $.penaltyBoxAction = Objects.requireNonNull($.penaltyBoxAction, "expected parameter 'penaltyBoxAction' to be non-null");
            $.penaltyBoxProtection = Objects.requireNonNull($.penaltyBoxProtection, "expected parameter 'penaltyBoxProtection' to be non-null");
            $.securityPolicyId = Objects.requireNonNull($.securityPolicyId, "expected parameter 'securityPolicyId' to be non-null");
            return $;
        }
    }

}