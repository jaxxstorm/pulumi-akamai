// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AppSecEvalRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final AppSecEvalRuleArgs Empty = new AppSecEvalRuleArgs();

    /**
     * . Path to a JSON file containing the conditions and exceptions to be applied to the evaluation rule.
     * 
     */
    @Import(name="conditionException")
    private @Nullable Output<String> conditionException;

    /**
     * @return . Path to a JSON file containing the conditions and exceptions to be applied to the evaluation rule.
     * 
     */
    public Optional<Output<String>> conditionException() {
        return Optional.ofNullable(this.conditionException);
    }

    /**
     * . Unique identifier of the security configuration in evaluation mode.
     * 
     */
    @Import(name="configId", required=true)
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration in evaluation mode.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }

    /**
     * . Action to be taken any time the evaluation rule is triggered, Allowed actions are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     * 
     */
    @Import(name="ruleAction", required=true)
    private Output<String> ruleAction;

    /**
     * @return . Action to be taken any time the evaluation rule is triggered, Allowed actions are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     * 
     */
    public Output<String> ruleAction() {
        return this.ruleAction;
    }

    /**
     * . Unique identifier of the evaluation rule being modified.
     * 
     */
    @Import(name="ruleId", required=true)
    private Output<Integer> ruleId;

    /**
     * @return . Unique identifier of the evaluation rule being modified.
     * 
     */
    public Output<Integer> ruleId() {
        return this.ruleId;
    }

    /**
     * . Unique identifier of the security policy associated with the evaluation process.
     * 
     */
    @Import(name="securityPolicyId", required=true)
    private Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the evaluation process.
     * 
     */
    public Output<String> securityPolicyId() {
        return this.securityPolicyId;
    }

    private AppSecEvalRuleArgs() {}

    private AppSecEvalRuleArgs(AppSecEvalRuleArgs $) {
        this.conditionException = $.conditionException;
        this.configId = $.configId;
        this.ruleAction = $.ruleAction;
        this.ruleId = $.ruleId;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecEvalRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecEvalRuleArgs $;

        public Builder() {
            $ = new AppSecEvalRuleArgs();
        }

        public Builder(AppSecEvalRuleArgs defaults) {
            $ = new AppSecEvalRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param conditionException . Path to a JSON file containing the conditions and exceptions to be applied to the evaluation rule.
         * 
         * @return builder
         * 
         */
        public Builder conditionException(@Nullable Output<String> conditionException) {
            $.conditionException = conditionException;
            return this;
        }

        /**
         * @param conditionException . Path to a JSON file containing the conditions and exceptions to be applied to the evaluation rule.
         * 
         * @return builder
         * 
         */
        public Builder conditionException(String conditionException) {
            return conditionException(Output.of(conditionException));
        }

        /**
         * @param configId . Unique identifier of the security configuration in evaluation mode.
         * 
         * @return builder
         * 
         */
        public Builder configId(Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration in evaluation mode.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param ruleAction . Action to be taken any time the evaluation rule is triggered, Allowed actions are:
         * - **alert**. Record the event.
         * - **deny**. Block the request.
         * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
         * - **none**. Take no action.
         * 
         * @return builder
         * 
         */
        public Builder ruleAction(Output<String> ruleAction) {
            $.ruleAction = ruleAction;
            return this;
        }

        /**
         * @param ruleAction . Action to be taken any time the evaluation rule is triggered, Allowed actions are:
         * - **alert**. Record the event.
         * - **deny**. Block the request.
         * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
         * - **none**. Take no action.
         * 
         * @return builder
         * 
         */
        public Builder ruleAction(String ruleAction) {
            return ruleAction(Output.of(ruleAction));
        }

        /**
         * @param ruleId . Unique identifier of the evaluation rule being modified.
         * 
         * @return builder
         * 
         */
        public Builder ruleId(Output<Integer> ruleId) {
            $.ruleId = ruleId;
            return this;
        }

        /**
         * @param ruleId . Unique identifier of the evaluation rule being modified.
         * 
         * @return builder
         * 
         */
        public Builder ruleId(Integer ruleId) {
            return ruleId(Output.of(ruleId));
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the evaluation process.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the evaluation process.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        public AppSecEvalRuleArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            $.ruleAction = Objects.requireNonNull($.ruleAction, "expected parameter 'ruleAction' to be non-null");
            $.ruleId = Objects.requireNonNull($.ruleId, "expected parameter 'ruleId' to be non-null");
            $.securityPolicyId = Objects.requireNonNull($.securityPolicyId, "expected parameter 'securityPolicyId' to be non-null");
            return $;
        }
    }

}