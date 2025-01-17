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


public final class AppSecAttackGroupState extends com.pulumi.resources.ResourceArgs {

    public static final AppSecAttackGroupState Empty = new AppSecAttackGroupState();

    /**
     * . Unique name of the attack group being modified.
     * 
     */
    @Import(name="attackGroup")
    private @Nullable Output<String> attackGroup;

    /**
     * @return . Unique name of the attack group being modified.
     * 
     */
    public Optional<Output<String>> attackGroup() {
        return Optional.ofNullable(this.attackGroup);
    }

    /**
     * . Action taken any time the attack group is triggered. Allowed values are:
     * - **alert**. Record information about the request.
     * - **deny**. Block the request,
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     * 
     */
    @Import(name="attackGroupAction")
    private @Nullable Output<String> attackGroupAction;

    /**
     * @return . Action taken any time the attack group is triggered. Allowed values are:
     * - **alert**. Record information about the request.
     * - **deny**. Block the request,
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     * 
     */
    public Optional<Output<String>> attackGroupAction() {
        return Optional.ofNullable(this.attackGroupAction);
    }

    /**
     * . Path to a JSON file containing the conditions and exceptions to be assigned to the attack group.
     * 
     */
    @Import(name="conditionException")
    private @Nullable Output<String> conditionException;

    /**
     * @return . Path to a JSON file containing the conditions and exceptions to be assigned to the attack group.
     * 
     */
    public Optional<Output<String>> conditionException() {
        return Optional.ofNullable(this.conditionException);
    }

    /**
     * . Unique identifier of the security configuration associated with the attack group being modified.
     * 
     */
    @Import(name="configId")
    private @Nullable Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the attack group being modified.
     * 
     */
    public Optional<Output<Integer>> configId() {
        return Optional.ofNullable(this.configId);
    }

    /**
     * . Unique identifier of the security policy associated with the attack group being modified.
     * 
     */
    @Import(name="securityPolicyId")
    private @Nullable Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the attack group being modified.
     * 
     */
    public Optional<Output<String>> securityPolicyId() {
        return Optional.ofNullable(this.securityPolicyId);
    }

    private AppSecAttackGroupState() {}

    private AppSecAttackGroupState(AppSecAttackGroupState $) {
        this.attackGroup = $.attackGroup;
        this.attackGroupAction = $.attackGroupAction;
        this.conditionException = $.conditionException;
        this.configId = $.configId;
        this.securityPolicyId = $.securityPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecAttackGroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecAttackGroupState $;

        public Builder() {
            $ = new AppSecAttackGroupState();
        }

        public Builder(AppSecAttackGroupState defaults) {
            $ = new AppSecAttackGroupState(Objects.requireNonNull(defaults));
        }

        /**
         * @param attackGroup . Unique name of the attack group being modified.
         * 
         * @return builder
         * 
         */
        public Builder attackGroup(@Nullable Output<String> attackGroup) {
            $.attackGroup = attackGroup;
            return this;
        }

        /**
         * @param attackGroup . Unique name of the attack group being modified.
         * 
         * @return builder
         * 
         */
        public Builder attackGroup(String attackGroup) {
            return attackGroup(Output.of(attackGroup));
        }

        /**
         * @param attackGroupAction . Action taken any time the attack group is triggered. Allowed values are:
         * - **alert**. Record information about the request.
         * - **deny**. Block the request,
         * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
         * - **none**. Take no action.
         * 
         * @return builder
         * 
         */
        public Builder attackGroupAction(@Nullable Output<String> attackGroupAction) {
            $.attackGroupAction = attackGroupAction;
            return this;
        }

        /**
         * @param attackGroupAction . Action taken any time the attack group is triggered. Allowed values are:
         * - **alert**. Record information about the request.
         * - **deny**. Block the request,
         * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
         * - **none**. Take no action.
         * 
         * @return builder
         * 
         */
        public Builder attackGroupAction(String attackGroupAction) {
            return attackGroupAction(Output.of(attackGroupAction));
        }

        /**
         * @param conditionException . Path to a JSON file containing the conditions and exceptions to be assigned to the attack group.
         * 
         * @return builder
         * 
         */
        public Builder conditionException(@Nullable Output<String> conditionException) {
            $.conditionException = conditionException;
            return this;
        }

        /**
         * @param conditionException . Path to a JSON file containing the conditions and exceptions to be assigned to the attack group.
         * 
         * @return builder
         * 
         */
        public Builder conditionException(String conditionException) {
            return conditionException(Output.of(conditionException));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the attack group being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(@Nullable Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the attack group being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the attack group being modified.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(@Nullable Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the attack group being modified.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        public AppSecAttackGroupState build() {
            return $;
        }
    }

}
