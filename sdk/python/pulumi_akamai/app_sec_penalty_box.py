# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecPenaltyBoxArgs', 'AppSecPenaltyBox']

@pulumi.input_type
class AppSecPenaltyBoxArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 penalty_box_action: pulumi.Input[str],
                 penalty_box_protection: pulumi.Input[bool],
                 security_policy_id: pulumi.Input[str],
                 version: pulumi.Input[int]):
        """
        The set of arguments for constructing a AppSecPenaltyBox resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] penalty_box_action: The action to take when penalty box protection is triggered: `alert` to record the trigger event, `deny` to block the request, `deny_custom_{custom_deny_id}` to execute a custom deny action, or `none` to take no action. Ignored if `penalty_box_protection` is set to `false`.
        :param pulumi.Input[bool] penalty_box_protection: A boolean value indicating whether to enable penalty box protection.
        :param pulumi.Input[str] security_policy_id: The ID of the security policy to use.
        :param pulumi.Input[int] version: The version number of the security configuration to use.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "penalty_box_action", penalty_box_action)
        pulumi.set(__self__, "penalty_box_protection", penalty_box_protection)
        pulumi.set(__self__, "security_policy_id", security_policy_id)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        The ID of the security configuration to use.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="penaltyBoxAction")
    def penalty_box_action(self) -> pulumi.Input[str]:
        """
        The action to take when penalty box protection is triggered: `alert` to record the trigger event, `deny` to block the request, `deny_custom_{custom_deny_id}` to execute a custom deny action, or `none` to take no action. Ignored if `penalty_box_protection` is set to `false`.
        """
        return pulumi.get(self, "penalty_box_action")

    @penalty_box_action.setter
    def penalty_box_action(self, value: pulumi.Input[str]):
        pulumi.set(self, "penalty_box_action", value)

    @property
    @pulumi.getter(name="penaltyBoxProtection")
    def penalty_box_protection(self) -> pulumi.Input[bool]:
        """
        A boolean value indicating whether to enable penalty box protection.
        """
        return pulumi.get(self, "penalty_box_protection")

    @penalty_box_protection.setter
    def penalty_box_protection(self, value: pulumi.Input[bool]):
        pulumi.set(self, "penalty_box_protection", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Input[str]:
        """
        The ID of the security policy to use.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_policy_id", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[int]:
        """
        The version number of the security configuration to use.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[int]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class _AppSecPenaltyBoxState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 penalty_box_action: Optional[pulumi.Input[str]] = None,
                 penalty_box_protection: Optional[pulumi.Input[bool]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering AppSecPenaltyBox resources.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] penalty_box_action: The action to take when penalty box protection is triggered: `alert` to record the trigger event, `deny` to block the request, `deny_custom_{custom_deny_id}` to execute a custom deny action, or `none` to take no action. Ignored if `penalty_box_protection` is set to `false`.
        :param pulumi.Input[bool] penalty_box_protection: A boolean value indicating whether to enable penalty box protection.
        :param pulumi.Input[str] security_policy_id: The ID of the security policy to use.
        :param pulumi.Input[int] version: The version number of the security configuration to use.
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if penalty_box_action is not None:
            pulumi.set(__self__, "penalty_box_action", penalty_box_action)
        if penalty_box_protection is not None:
            pulumi.set(__self__, "penalty_box_protection", penalty_box_protection)
        if security_policy_id is not None:
            pulumi.set(__self__, "security_policy_id", security_policy_id)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the security configuration to use.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="penaltyBoxAction")
    def penalty_box_action(self) -> Optional[pulumi.Input[str]]:
        """
        The action to take when penalty box protection is triggered: `alert` to record the trigger event, `deny` to block the request, `deny_custom_{custom_deny_id}` to execute a custom deny action, or `none` to take no action. Ignored if `penalty_box_protection` is set to `false`.
        """
        return pulumi.get(self, "penalty_box_action")

    @penalty_box_action.setter
    def penalty_box_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "penalty_box_action", value)

    @property
    @pulumi.getter(name="penaltyBoxProtection")
    def penalty_box_protection(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean value indicating whether to enable penalty box protection.
        """
        return pulumi.get(self, "penalty_box_protection")

    @penalty_box_protection.setter
    def penalty_box_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "penalty_box_protection", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the security policy to use.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_policy_id", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        The version number of the security configuration to use.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class AppSecPenaltyBox(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 penalty_box_action: Optional[pulumi.Input[str]] = None,
                 penalty_box_protection: Optional[pulumi.Input[bool]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Use the `AppSecPenaltyBox` resource to update the penalty box settings for a given security policy.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
        penalty_box = akamai.AppSecPenaltyBox("penaltyBox",
            config_id=configuration.config_id,
            version=configuration.latest_version,
            security_policy_id=var["security_policy_id"],
            penalty_box_protection=True,
            penalty_box_action=var["action"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] penalty_box_action: The action to take when penalty box protection is triggered: `alert` to record the trigger event, `deny` to block the request, `deny_custom_{custom_deny_id}` to execute a custom deny action, or `none` to take no action. Ignored if `penalty_box_protection` is set to `false`.
        :param pulumi.Input[bool] penalty_box_protection: A boolean value indicating whether to enable penalty box protection.
        :param pulumi.Input[str] security_policy_id: The ID of the security policy to use.
        :param pulumi.Input[int] version: The version number of the security configuration to use.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecPenaltyBoxArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use the `AppSecPenaltyBox` resource to update the penalty box settings for a given security policy.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
        penalty_box = akamai.AppSecPenaltyBox("penaltyBox",
            config_id=configuration.config_id,
            version=configuration.latest_version,
            security_policy_id=var["security_policy_id"],
            penalty_box_protection=True,
            penalty_box_action=var["action"])
        ```

        :param str resource_name: The name of the resource.
        :param AppSecPenaltyBoxArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecPenaltyBoxArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 penalty_box_action: Optional[pulumi.Input[str]] = None,
                 penalty_box_protection: Optional[pulumi.Input[bool]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
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
            __props__ = AppSecPenaltyBoxArgs.__new__(AppSecPenaltyBoxArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if penalty_box_action is None and not opts.urn:
                raise TypeError("Missing required property 'penalty_box_action'")
            __props__.__dict__["penalty_box_action"] = penalty_box_action
            if penalty_box_protection is None and not opts.urn:
                raise TypeError("Missing required property 'penalty_box_protection'")
            __props__.__dict__["penalty_box_protection"] = penalty_box_protection
            if security_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_policy_id'")
            __props__.__dict__["security_policy_id"] = security_policy_id
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__.__dict__["version"] = version
        super(AppSecPenaltyBox, __self__).__init__(
            'akamai:index/appSecPenaltyBox:AppSecPenaltyBox',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            penalty_box_action: Optional[pulumi.Input[str]] = None,
            penalty_box_protection: Optional[pulumi.Input[bool]] = None,
            security_policy_id: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'AppSecPenaltyBox':
        """
        Get an existing AppSecPenaltyBox resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] penalty_box_action: The action to take when penalty box protection is triggered: `alert` to record the trigger event, `deny` to block the request, `deny_custom_{custom_deny_id}` to execute a custom deny action, or `none` to take no action. Ignored if `penalty_box_protection` is set to `false`.
        :param pulumi.Input[bool] penalty_box_protection: A boolean value indicating whether to enable penalty box protection.
        :param pulumi.Input[str] security_policy_id: The ID of the security policy to use.
        :param pulumi.Input[int] version: The version number of the security configuration to use.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecPenaltyBoxState.__new__(_AppSecPenaltyBoxState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["penalty_box_action"] = penalty_box_action
        __props__.__dict__["penalty_box_protection"] = penalty_box_protection
        __props__.__dict__["security_policy_id"] = security_policy_id
        __props__.__dict__["version"] = version
        return AppSecPenaltyBox(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        The ID of the security configuration to use.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="penaltyBoxAction")
    def penalty_box_action(self) -> pulumi.Output[str]:
        """
        The action to take when penalty box protection is triggered: `alert` to record the trigger event, `deny` to block the request, `deny_custom_{custom_deny_id}` to execute a custom deny action, or `none` to take no action. Ignored if `penalty_box_protection` is set to `false`.
        """
        return pulumi.get(self, "penalty_box_action")

    @property
    @pulumi.getter(name="penaltyBoxProtection")
    def penalty_box_protection(self) -> pulumi.Output[bool]:
        """
        A boolean value indicating whether to enable penalty box protection.
        """
        return pulumi.get(self, "penalty_box_protection")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Output[str]:
        """
        The ID of the security policy to use.
        """
        return pulumi.get(self, "security_policy_id")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        The version number of the security configuration to use.
        """
        return pulumi.get(self, "version")

