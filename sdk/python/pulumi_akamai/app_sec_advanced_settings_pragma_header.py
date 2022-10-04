# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecAdvancedSettingsPragmaHeaderArgs', 'AppSecAdvancedSettingsPragmaHeader']

@pulumi.input_type
class AppSecAdvancedSettingsPragmaHeaderArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 pragma_header: pulumi.Input[str],
                 security_policy_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AppSecAdvancedSettingsPragmaHeader resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the pragma header settings being modified.
        :param pulumi.Input[str] pragma_header: . Path to a JSON file containing information about the conditions to exclude from the default remove action. By default, the Pragma header debugging information is stripped from an operation's response except in cases where you set `excludeCondition`.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the pragma header settings being modified. If not included, pragma header settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "pragma_header", pragma_header)
        if security_policy_id is not None:
            pulumi.set(__self__, "security_policy_id", security_policy_id)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        . Unique identifier of the security configuration associated with the pragma header settings being modified.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="pragmaHeader")
    def pragma_header(self) -> pulumi.Input[str]:
        """
        . Path to a JSON file containing information about the conditions to exclude from the default remove action. By default, the Pragma header debugging information is stripped from an operation's response except in cases where you set `excludeCondition`.
        """
        return pulumi.get(self, "pragma_header")

    @pragma_header.setter
    def pragma_header(self, value: pulumi.Input[str]):
        pulumi.set(self, "pragma_header", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        . Unique identifier of the security policy associated with the pragma header settings being modified. If not included, pragma header settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_policy_id", value)


@pulumi.input_type
class _AppSecAdvancedSettingsPragmaHeaderState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 pragma_header: Optional[pulumi.Input[str]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AppSecAdvancedSettingsPragmaHeader resources.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the pragma header settings being modified.
        :param pulumi.Input[str] pragma_header: . Path to a JSON file containing information about the conditions to exclude from the default remove action. By default, the Pragma header debugging information is stripped from an operation's response except in cases where you set `excludeCondition`.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the pragma header settings being modified. If not included, pragma header settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if pragma_header is not None:
            pulumi.set(__self__, "pragma_header", pragma_header)
        if security_policy_id is not None:
            pulumi.set(__self__, "security_policy_id", security_policy_id)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        . Unique identifier of the security configuration associated with the pragma header settings being modified.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="pragmaHeader")
    def pragma_header(self) -> Optional[pulumi.Input[str]]:
        """
        . Path to a JSON file containing information about the conditions to exclude from the default remove action. By default, the Pragma header debugging information is stripped from an operation's response except in cases where you set `excludeCondition`.
        """
        return pulumi.get(self, "pragma_header")

    @pragma_header.setter
    def pragma_header(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pragma_header", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        . Unique identifier of the security policy associated with the pragma header settings being modified. If not included, pragma header settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_policy_id", value)


class AppSecAdvancedSettingsPragmaHeader(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 pragma_header: Optional[pulumi.Input[str]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        **Scopes**: Security configuration; security policy

        Specifies the headers you can exclude from inspection when you are working with a Pragma debug header, a header that provides information about such things as: the edge routers used in a transaction; the Akamai IP addresses involved; whether a request was cached or not; etc. By default, pragma headers are removed from all responses.

        This operation can be applied at the security configuration level (in which case it applies to all the security policies in the configuration), or can be customized for an individual security policy.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/advanced-settings/pragma-header](https://techdocs.akamai.com/application-security/reference/put-policies-pragma-header)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        pragma_header = akamai.AppSecAdvancedSettingsPragmaHeader("pragmaHeader",
            config_id=configuration.config_id,
            security_policy_id="gms1_134637",
            pragma_header=(lambda path: open(path).read())(f"{path['module']}/pragma_header.json"))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the pragma header settings being modified.
        :param pulumi.Input[str] pragma_header: . Path to a JSON file containing information about the conditions to exclude from the default remove action. By default, the Pragma header debugging information is stripped from an operation's response except in cases where you set `excludeCondition`.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the pragma header settings being modified. If not included, pragma header settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecAdvancedSettingsPragmaHeaderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        **Scopes**: Security configuration; security policy

        Specifies the headers you can exclude from inspection when you are working with a Pragma debug header, a header that provides information about such things as: the edge routers used in a transaction; the Akamai IP addresses involved; whether a request was cached or not; etc. By default, pragma headers are removed from all responses.

        This operation can be applied at the security configuration level (in which case it applies to all the security policies in the configuration), or can be customized for an individual security policy.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/advanced-settings/pragma-header](https://techdocs.akamai.com/application-security/reference/put-policies-pragma-header)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        pragma_header = akamai.AppSecAdvancedSettingsPragmaHeader("pragmaHeader",
            config_id=configuration.config_id,
            security_policy_id="gms1_134637",
            pragma_header=(lambda path: open(path).read())(f"{path['module']}/pragma_header.json"))
        ```

        :param str resource_name: The name of the resource.
        :param AppSecAdvancedSettingsPragmaHeaderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecAdvancedSettingsPragmaHeaderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 pragma_header: Optional[pulumi.Input[str]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppSecAdvancedSettingsPragmaHeaderArgs.__new__(AppSecAdvancedSettingsPragmaHeaderArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if pragma_header is None and not opts.urn:
                raise TypeError("Missing required property 'pragma_header'")
            __props__.__dict__["pragma_header"] = pragma_header
            __props__.__dict__["security_policy_id"] = security_policy_id
        super(AppSecAdvancedSettingsPragmaHeader, __self__).__init__(
            'akamai:index/appSecAdvancedSettingsPragmaHeader:AppSecAdvancedSettingsPragmaHeader',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            pragma_header: Optional[pulumi.Input[str]] = None,
            security_policy_id: Optional[pulumi.Input[str]] = None) -> 'AppSecAdvancedSettingsPragmaHeader':
        """
        Get an existing AppSecAdvancedSettingsPragmaHeader resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the pragma header settings being modified.
        :param pulumi.Input[str] pragma_header: . Path to a JSON file containing information about the conditions to exclude from the default remove action. By default, the Pragma header debugging information is stripped from an operation's response except in cases where you set `excludeCondition`.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the pragma header settings being modified. If not included, pragma header settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecAdvancedSettingsPragmaHeaderState.__new__(_AppSecAdvancedSettingsPragmaHeaderState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["pragma_header"] = pragma_header
        __props__.__dict__["security_policy_id"] = security_policy_id
        return AppSecAdvancedSettingsPragmaHeader(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        . Unique identifier of the security configuration associated with the pragma header settings being modified.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="pragmaHeader")
    def pragma_header(self) -> pulumi.Output[str]:
        """
        . Path to a JSON file containing information about the conditions to exclude from the default remove action. By default, the Pragma header debugging information is stripped from an operation's response except in cases where you set `excludeCondition`.
        """
        return pulumi.get(self, "pragma_header")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Output[Optional[str]]:
        """
        . Unique identifier of the security policy associated with the pragma header settings being modified. If not included, pragma header settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
        """
        return pulumi.get(self, "security_policy_id")

