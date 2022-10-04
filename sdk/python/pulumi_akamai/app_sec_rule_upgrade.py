# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecRuleUpgradeArgs', 'AppSecRuleUpgrade']

@pulumi.input_type
class AppSecRuleUpgradeArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 security_policy_id: pulumi.Input[str],
                 upgrade_mode: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AppSecRuleUpgrade resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the ruleset being upgraded.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the ruleset being upgraded.
               - `upgrade_mode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
               - **ASE_AUTO**. Akamai automatically updates your rulesets.
               - **ASE_MANUAL**. Manually updates your rulesets.
        :param pulumi.Input[str] upgrade_mode: Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "security_policy_id", security_policy_id)
        if upgrade_mode is not None:
            pulumi.set(__self__, "upgrade_mode", upgrade_mode)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        . Unique identifier of the security configuration associated with the ruleset being upgraded.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Input[str]:
        """
        . Unique identifier of the security policy associated with the ruleset being upgraded.
        - `upgrade_mode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
        - **ASE_AUTO**. Akamai automatically updates your rulesets.
        - **ASE_MANUAL**. Manually updates your rulesets.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_policy_id", value)

    @property
    @pulumi.getter(name="upgradeMode")
    def upgrade_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
        """
        return pulumi.get(self, "upgrade_mode")

    @upgrade_mode.setter
    def upgrade_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upgrade_mode", value)


@pulumi.input_type
class _AppSecRuleUpgradeState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 current_ruleset: Optional[pulumi.Input[str]] = None,
                 eval_status: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 upgrade_mode: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AppSecRuleUpgrade resources.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the ruleset being upgraded.
        :param pulumi.Input[str] current_ruleset: Versioning information for the current KRS rule set
        :param pulumi.Input[str] eval_status: Whether an evaluation is currently in progress
        :param pulumi.Input[str] mode: Upgrade mode (KRS, AAG, ASE_MANUAL or ASE_AUTO)
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the ruleset being upgraded.
               - `upgrade_mode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
               - **ASE_AUTO**. Akamai automatically updates your rulesets.
               - **ASE_MANUAL**. Manually updates your rulesets.
        :param pulumi.Input[str] upgrade_mode: Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if current_ruleset is not None:
            pulumi.set(__self__, "current_ruleset", current_ruleset)
        if eval_status is not None:
            pulumi.set(__self__, "eval_status", eval_status)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if security_policy_id is not None:
            pulumi.set(__self__, "security_policy_id", security_policy_id)
        if upgrade_mode is not None:
            pulumi.set(__self__, "upgrade_mode", upgrade_mode)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        . Unique identifier of the security configuration associated with the ruleset being upgraded.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="currentRuleset")
    def current_ruleset(self) -> Optional[pulumi.Input[str]]:
        """
        Versioning information for the current KRS rule set
        """
        return pulumi.get(self, "current_ruleset")

    @current_ruleset.setter
    def current_ruleset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "current_ruleset", value)

    @property
    @pulumi.getter(name="evalStatus")
    def eval_status(self) -> Optional[pulumi.Input[str]]:
        """
        Whether an evaluation is currently in progress
        """
        return pulumi.get(self, "eval_status")

    @eval_status.setter
    def eval_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eval_status", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Upgrade mode (KRS, AAG, ASE_MANUAL or ASE_AUTO)
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        . Unique identifier of the security policy associated with the ruleset being upgraded.
        - `upgrade_mode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
        - **ASE_AUTO**. Akamai automatically updates your rulesets.
        - **ASE_MANUAL**. Manually updates your rulesets.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_policy_id", value)

    @property
    @pulumi.getter(name="upgradeMode")
    def upgrade_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
        """
        return pulumi.get(self, "upgrade_mode")

    @upgrade_mode.setter
    def upgrade_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upgrade_mode", value)


class AppSecRuleUpgrade(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 upgrade_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        rule_upgrade = akamai.AppSecRuleUpgrade("ruleUpgrade",
            config_id=configuration.config_id,
            security_policy_id="gms1_134637")
        pulumi.export("ruleUpgradeCurrentRuleset", rule_upgrade.current_ruleset)
        pulumi.export("ruleUpgradeMode", rule_upgrade.mode)
        pulumi.export("ruleUpgradeEvalStatus", rule_upgrade.eval_status)
        ```
        ## Output Options

        The following options can be used to determine the information returned and how that returned information is formatted:

        - `current_ruleset`. Versioning information for your current KRS rule set.
        - `mode`. Specifies the current upgrade mode type. Valid values are:
          - **KRS**. Rulesets must be manually upgraded.
          
          - **AAG**. Rulesets are automatically upgraded by Akamai.
          
          - **ASE_MANUAL**. Adaptive Security Engine rulesets must be manually upgraded.
          
          - **ASE_AUTO**. Adaptive Security Engine rulesets are automatically updated by Akamai.

        - `eval_status`. Returns **enabled** if an evaluation is currently in progress; otherwise returns **disabled**.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the ruleset being upgraded.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the ruleset being upgraded.
               - `upgrade_mode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
               - **ASE_AUTO**. Akamai automatically updates your rulesets.
               - **ASE_MANUAL**. Manually updates your rulesets.
        :param pulumi.Input[str] upgrade_mode: Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecRuleUpgradeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        rule_upgrade = akamai.AppSecRuleUpgrade("ruleUpgrade",
            config_id=configuration.config_id,
            security_policy_id="gms1_134637")
        pulumi.export("ruleUpgradeCurrentRuleset", rule_upgrade.current_ruleset)
        pulumi.export("ruleUpgradeMode", rule_upgrade.mode)
        pulumi.export("ruleUpgradeEvalStatus", rule_upgrade.eval_status)
        ```
        ## Output Options

        The following options can be used to determine the information returned and how that returned information is formatted:

        - `current_ruleset`. Versioning information for your current KRS rule set.
        - `mode`. Specifies the current upgrade mode type. Valid values are:
          - **KRS**. Rulesets must be manually upgraded.
          
          - **AAG**. Rulesets are automatically upgraded by Akamai.
          
          - **ASE_MANUAL**. Adaptive Security Engine rulesets must be manually upgraded.
          
          - **ASE_AUTO**. Adaptive Security Engine rulesets are automatically updated by Akamai.

        - `eval_status`. Returns **enabled** if an evaluation is currently in progress; otherwise returns **disabled**.

        :param str resource_name: The name of the resource.
        :param AppSecRuleUpgradeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecRuleUpgradeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 upgrade_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppSecRuleUpgradeArgs.__new__(AppSecRuleUpgradeArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if security_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_policy_id'")
            __props__.__dict__["security_policy_id"] = security_policy_id
            __props__.__dict__["upgrade_mode"] = upgrade_mode
            __props__.__dict__["current_ruleset"] = None
            __props__.__dict__["eval_status"] = None
            __props__.__dict__["mode"] = None
        super(AppSecRuleUpgrade, __self__).__init__(
            'akamai:index/appSecRuleUpgrade:AppSecRuleUpgrade',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            current_ruleset: Optional[pulumi.Input[str]] = None,
            eval_status: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            security_policy_id: Optional[pulumi.Input[str]] = None,
            upgrade_mode: Optional[pulumi.Input[str]] = None) -> 'AppSecRuleUpgrade':
        """
        Get an existing AppSecRuleUpgrade resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the ruleset being upgraded.
        :param pulumi.Input[str] current_ruleset: Versioning information for the current KRS rule set
        :param pulumi.Input[str] eval_status: Whether an evaluation is currently in progress
        :param pulumi.Input[str] mode: Upgrade mode (KRS, AAG, ASE_MANUAL or ASE_AUTO)
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the ruleset being upgraded.
               - `upgrade_mode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
               - **ASE_AUTO**. Akamai automatically updates your rulesets.
               - **ASE_MANUAL**. Manually updates your rulesets.
        :param pulumi.Input[str] upgrade_mode: Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecRuleUpgradeState.__new__(_AppSecRuleUpgradeState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["current_ruleset"] = current_ruleset
        __props__.__dict__["eval_status"] = eval_status
        __props__.__dict__["mode"] = mode
        __props__.__dict__["security_policy_id"] = security_policy_id
        __props__.__dict__["upgrade_mode"] = upgrade_mode
        return AppSecRuleUpgrade(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        . Unique identifier of the security configuration associated with the ruleset being upgraded.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="currentRuleset")
    def current_ruleset(self) -> pulumi.Output[str]:
        """
        Versioning information for the current KRS rule set
        """
        return pulumi.get(self, "current_ruleset")

    @property
    @pulumi.getter(name="evalStatus")
    def eval_status(self) -> pulumi.Output[str]:
        """
        Whether an evaluation is currently in progress
        """
        return pulumi.get(self, "eval_status")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        """
        Upgrade mode (KRS, AAG, ASE_MANUAL or ASE_AUTO)
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Output[str]:
        """
        . Unique identifier of the security policy associated with the ruleset being upgraded.
        - `upgrade_mode`. (Optional). Modifies the upgrade type for organizations running the ASE beta. Allowed values are:
        - **ASE_AUTO**. Akamai automatically updates your rulesets.
        - **ASE_MANUAL**. Manually updates your rulesets.
        """
        return pulumi.get(self, "security_policy_id")

    @property
    @pulumi.getter(name="upgradeMode")
    def upgrade_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Modifies the upgrade type for organizations running the ASE beta (ASE_AUTO or ASE_MANUAL)
        """
        return pulumi.get(self, "upgrade_mode")

