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


public final class AppSecRuleUpgradeArgs extends com.pulumi.resources.ResourceArgs {

    public static final AppSecRuleUpgradeArgs Empty = new AppSecRuleUpgradeArgs();

    /**
     * . Unique identifier of the security configuration associated with the ruleset being upgraded.
     * 
     */
    @Import(name="configId", required=true)
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the ruleset being upgraded.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }

    /**
     * . Unique identifier of the security policy associated with the ruleset being upgraded.
     * - `upgrade_mode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
     * - **ASE_AUTO**. Akamai automatically updates your rulesets.
     * - **ASE_MANUAL**. Manually updates your rulesets.
     * 
     */
    @Import(name="securityPolicyId", required=true)
    private Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the ruleset being upgraded.
     * - `upgrade_mode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
     * - **ASE_AUTO**. Akamai automatically updates your rulesets.
     * - **ASE_MANUAL**. Manually updates your rulesets.
     * 
     */
    public Output<String> securityPolicyId() {
        return this.securityPolicyId;
    }

    /**
     * Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
     * 
     */
    @Import(name="upgradeMode")
    private @Nullable Output<String> upgradeMode;

    /**
     * @return Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
     * 
     */
    public Optional<Output<String>> upgradeMode() {
        return Optional.ofNullable(this.upgradeMode);
    }

    private AppSecRuleUpgradeArgs() {}

    private AppSecRuleUpgradeArgs(AppSecRuleUpgradeArgs $) {
        this.configId = $.configId;
        this.securityPolicyId = $.securityPolicyId;
        this.upgradeMode = $.upgradeMode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecRuleUpgradeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecRuleUpgradeArgs $;

        public Builder() {
            $ = new AppSecRuleUpgradeArgs();
        }

        public Builder(AppSecRuleUpgradeArgs defaults) {
            $ = new AppSecRuleUpgradeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the ruleset being upgraded.
         * 
         * @return builder
         * 
         */
        public Builder configId(Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the ruleset being upgraded.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the ruleset being upgraded.
         * - `upgrade_mode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
         * - **ASE_AUTO**. Akamai automatically updates your rulesets.
         * - **ASE_MANUAL**. Manually updates your rulesets.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId . Unique identifier of the security policy associated with the ruleset being upgraded.
         * - `upgrade_mode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
         * - **ASE_AUTO**. Akamai automatically updates your rulesets.
         * - **ASE_MANUAL**. Manually updates your rulesets.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        /**
         * @param upgradeMode Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
         * 
         * @return builder
         * 
         */
        public Builder upgradeMode(@Nullable Output<String> upgradeMode) {
            $.upgradeMode = upgradeMode;
            return this;
        }

        /**
         * @param upgradeMode Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
         * 
         * @return builder
         * 
         */
        public Builder upgradeMode(String upgradeMode) {
            return upgradeMode(Output.of(upgradeMode));
        }

        public AppSecRuleUpgradeArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            $.securityPolicyId = Objects.requireNonNull($.securityPolicyId, "expected parameter 'securityPolicyId' to be non-null");
            return $;
        }
    }

}