# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecRuleArgs', 'AppSecRule']

@pulumi.input_type
class AppSecRuleArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 rule_id: pulumi.Input[int],
                 security_policy_id: pulumi.Input[str],
                 condition_exception: Optional[pulumi.Input[str]] = None,
                 rule_action: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AppSecRule resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
        :param pulumi.Input[int] rule_id: . Unique identifier of the rule being modified.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
        :param pulumi.Input[str] condition_exception: . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule. You can view a sample JSON file in the [Modify the conditions and exceptions of a rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleconditionexception) section of the Application Security API documentation.
        :param pulumi.Input[str] rule_action: Allowed values are:
               - **alert**. Record the event.
               - **deny**. Block the request.
               - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
               - **none**. Take no action. or `none` to take no action.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "rule_id", rule_id)
        pulumi.set(__self__, "security_policy_id", security_policy_id)
        if condition_exception is not None:
            pulumi.set(__self__, "condition_exception", condition_exception)
        if rule_action is not None:
            pulumi.set(__self__, "rule_action", rule_action)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Input[int]:
        """
        . Unique identifier of the rule being modified.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Input[str]:
        """
        . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_policy_id", value)

    @property
    @pulumi.getter(name="conditionException")
    def condition_exception(self) -> Optional[pulumi.Input[str]]:
        """
        . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule. You can view a sample JSON file in the [Modify the conditions and exceptions of a rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleconditionexception) section of the Application Security API documentation.
        """
        return pulumi.get(self, "condition_exception")

    @condition_exception.setter
    def condition_exception(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "condition_exception", value)

    @property
    @pulumi.getter(name="ruleAction")
    def rule_action(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values are:
        - **alert**. Record the event.
        - **deny**. Block the request.
        - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
        - **none**. Take no action. or `none` to take no action.
        """
        return pulumi.get(self, "rule_action")

    @rule_action.setter
    def rule_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_action", value)


@pulumi.input_type
class _AppSecRuleState:
    def __init__(__self__, *,
                 condition_exception: Optional[pulumi.Input[str]] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 rule_action: Optional[pulumi.Input[str]] = None,
                 rule_id: Optional[pulumi.Input[int]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AppSecRule resources.
        :param pulumi.Input[str] condition_exception: . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule. You can view a sample JSON file in the [Modify the conditions and exceptions of a rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleconditionexception) section of the Application Security API documentation.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
        :param pulumi.Input[str] rule_action: Allowed values are:
               - **alert**. Record the event.
               - **deny**. Block the request.
               - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
               - **none**. Take no action. or `none` to take no action.
        :param pulumi.Input[int] rule_id: . Unique identifier of the rule being modified.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
        """
        if condition_exception is not None:
            pulumi.set(__self__, "condition_exception", condition_exception)
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if rule_action is not None:
            pulumi.set(__self__, "rule_action", rule_action)
        if rule_id is not None:
            pulumi.set(__self__, "rule_id", rule_id)
        if security_policy_id is not None:
            pulumi.set(__self__, "security_policy_id", security_policy_id)

    @property
    @pulumi.getter(name="conditionException")
    def condition_exception(self) -> Optional[pulumi.Input[str]]:
        """
        . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule. You can view a sample JSON file in the [Modify the conditions and exceptions of a rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleconditionexception) section of the Application Security API documentation.
        """
        return pulumi.get(self, "condition_exception")

    @condition_exception.setter
    def condition_exception(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "condition_exception", value)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="ruleAction")
    def rule_action(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values are:
        - **alert**. Record the event.
        - **deny**. Block the request.
        - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
        - **none**. Take no action. or `none` to take no action.
        """
        return pulumi.get(self, "rule_action")

    @rule_action.setter
    def rule_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_action", value)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[pulumi.Input[int]]:
        """
        . Unique identifier of the rule being modified.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_policy_id", value)


class AppSecRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition_exception: Optional[pulumi.Input[str]] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 rule_action: Optional[pulumi.Input[str]] = None,
                 rule_id: Optional[pulumi.Input[int]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        **Scopes**: Rule

        Modifies a Kona Rule Set rule's action, conditions, and exceptions.

        **Related API Endpoints**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rules/{ruleId}](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleaction) *and* [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rules/{ruleId}/condition-exception](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleconditionexception)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        rule = akamai.AppSecRule("rule",
            config_id=configuration.config_id,
            security_policy_id="gms1_134637",
            rule_id=60029316,
            rule_action="deny",
            condition_exception=(lambda path: open(path).read())(f"{path['module']}/condition_exception.json"))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] condition_exception: . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule. You can view a sample JSON file in the [Modify the conditions and exceptions of a rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleconditionexception) section of the Application Security API documentation.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
        :param pulumi.Input[str] rule_action: Allowed values are:
               - **alert**. Record the event.
               - **deny**. Block the request.
               - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
               - **none**. Take no action. or `none` to take no action.
        :param pulumi.Input[int] rule_id: . Unique identifier of the rule being modified.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        **Scopes**: Rule

        Modifies a Kona Rule Set rule's action, conditions, and exceptions.

        **Related API Endpoints**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rules/{ruleId}](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleaction) *and* [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rules/{ruleId}/condition-exception](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleconditionexception)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        rule = akamai.AppSecRule("rule",
            config_id=configuration.config_id,
            security_policy_id="gms1_134637",
            rule_id=60029316,
            rule_action="deny",
            condition_exception=(lambda path: open(path).read())(f"{path['module']}/condition_exception.json"))
        ```

        :param str resource_name: The name of the resource.
        :param AppSecRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition_exception: Optional[pulumi.Input[str]] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 rule_action: Optional[pulumi.Input[str]] = None,
                 rule_id: Optional[pulumi.Input[int]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppSecRuleArgs.__new__(AppSecRuleArgs)

            __props__.__dict__["condition_exception"] = condition_exception
            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            __props__.__dict__["rule_action"] = rule_action
            if rule_id is None and not opts.urn:
                raise TypeError("Missing required property 'rule_id'")
            __props__.__dict__["rule_id"] = rule_id
            if security_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_policy_id'")
            __props__.__dict__["security_policy_id"] = security_policy_id
        super(AppSecRule, __self__).__init__(
            'akamai:index/appSecRule:AppSecRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition_exception: Optional[pulumi.Input[str]] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            rule_action: Optional[pulumi.Input[str]] = None,
            rule_id: Optional[pulumi.Input[int]] = None,
            security_policy_id: Optional[pulumi.Input[str]] = None) -> 'AppSecRule':
        """
        Get an existing AppSecRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] condition_exception: . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule. You can view a sample JSON file in the [Modify the conditions and exceptions of a rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleconditionexception) section of the Application Security API documentation.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
        :param pulumi.Input[str] rule_action: Allowed values are:
               - **alert**. Record the event.
               - **deny**. Block the request.
               - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
               - **none**. Take no action. or `none` to take no action.
        :param pulumi.Input[int] rule_id: . Unique identifier of the rule being modified.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecRuleState.__new__(_AppSecRuleState)

        __props__.__dict__["condition_exception"] = condition_exception
        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["rule_action"] = rule_action
        __props__.__dict__["rule_id"] = rule_id
        __props__.__dict__["security_policy_id"] = security_policy_id
        return AppSecRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="conditionException")
    def condition_exception(self) -> pulumi.Output[Optional[str]]:
        """
        . Path to a JSON file containing a description of the conditions and exceptions to be associated with a rule. You can view a sample JSON file in the [Modify the conditions and exceptions of a rule](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putruleconditionexception) section of the Application Security API documentation.
        """
        return pulumi.get(self, "condition_exception")

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        . Unique identifier of the security configuration associated with the Kona Rule Set rule being modified.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="ruleAction")
    def rule_action(self) -> pulumi.Output[str]:
        """
        Allowed values are:
        - **alert**. Record the event.
        - **deny**. Block the request.
        - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
        - **none**. Take no action. or `none` to take no action.
        """
        return pulumi.get(self, "rule_action")

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Output[int]:
        """
        . Unique identifier of the rule being modified.
        """
        return pulumi.get(self, "rule_id")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Output[str]:
        """
        . Unique identifier of the security policy associated with the Kona Rule Set rule being modified.
        """
        return pulumi.get(self, "security_policy_id")

