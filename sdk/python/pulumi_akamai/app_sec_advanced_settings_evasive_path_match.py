# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecAdvancedSettingsEvasivePathMatchArgs', 'AppSecAdvancedSettingsEvasivePathMatch']

@pulumi.input_type
class AppSecAdvancedSettingsEvasivePathMatchArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 enable_path_match: pulumi.Input[bool],
                 security_policy_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AppSecAdvancedSettingsEvasivePathMatch resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[bool] enable_path_match: Whether to enable path match.
        :param pulumi.Input[str] security_policy_id: The ID of a specific security policy to which the evasive path match setting should be applied. If not supplied, the indicated setting will be applied to all policies within the configuration.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "enable_path_match", enable_path_match)
        if security_policy_id is not None:
            pulumi.set(__self__, "security_policy_id", security_policy_id)

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
    @pulumi.getter(name="enablePathMatch")
    def enable_path_match(self) -> pulumi.Input[bool]:
        """
        Whether to enable path match.
        """
        return pulumi.get(self, "enable_path_match")

    @enable_path_match.setter
    def enable_path_match(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enable_path_match", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of a specific security policy to which the evasive path match setting should be applied. If not supplied, the indicated setting will be applied to all policies within the configuration.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_policy_id", value)


@pulumi.input_type
class _AppSecAdvancedSettingsEvasivePathMatchState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 enable_path_match: Optional[pulumi.Input[bool]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AppSecAdvancedSettingsEvasivePathMatch resources.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[bool] enable_path_match: Whether to enable path match.
        :param pulumi.Input[str] security_policy_id: The ID of a specific security policy to which the evasive path match setting should be applied. If not supplied, the indicated setting will be applied to all policies within the configuration.
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if enable_path_match is not None:
            pulumi.set(__self__, "enable_path_match", enable_path_match)
        if security_policy_id is not None:
            pulumi.set(__self__, "security_policy_id", security_policy_id)

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
    @pulumi.getter(name="enablePathMatch")
    def enable_path_match(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable path match.
        """
        return pulumi.get(self, "enable_path_match")

    @enable_path_match.setter
    def enable_path_match(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_path_match", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of a specific security policy to which the evasive path match setting should be applied. If not supplied, the indicated setting will be applied to all policies within the configuration.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_policy_id", value)


class AppSecAdvancedSettingsEvasivePathMatch(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 enable_path_match: Optional[pulumi.Input[bool]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        **Scopes**: Security configuration; security policy

        The `resource_akamai_appsec_advanced_settings_evasive_path_match` resource allows you to enable, disable, or update the evasive path match setting for a configuration.
        This setting determines whether fuzzy matching is used to make URL matching more inclusive.
        This operation applies at the configuration level, and therefore applies to all policies within a configuration.
        You may override this setting for a particular policy by specifying the policy using the security_policy_id parameter.

        **Related API Endpoints**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/advanced-settings/evasive-path-match](https://techdocs.akamai.com/application-security/reference/put-evasive-path-match)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
        config_evasive_path_match = akamai.AppSecAdvancedSettingsEvasivePathMatch("configEvasivePathMatch",
            config_id=configuration.config_id,
            enable_path_match=True)
        # USE CASE: user wants to override the evasive path match setting for a security policy
        policy_override = akamai.AppSecAdvancedSettingsEvasivePathMatch("policyOverride",
            config_id=configuration.config_id,
            security_policy_id=var["security_policy_id"],
            enable_path_match=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[bool] enable_path_match: Whether to enable path match.
        :param pulumi.Input[str] security_policy_id: The ID of a specific security policy to which the evasive path match setting should be applied. If not supplied, the indicated setting will be applied to all policies within the configuration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecAdvancedSettingsEvasivePathMatchArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        **Scopes**: Security configuration; security policy

        The `resource_akamai_appsec_advanced_settings_evasive_path_match` resource allows you to enable, disable, or update the evasive path match setting for a configuration.
        This setting determines whether fuzzy matching is used to make URL matching more inclusive.
        This operation applies at the configuration level, and therefore applies to all policies within a configuration.
        You may override this setting for a particular policy by specifying the policy using the security_policy_id parameter.

        **Related API Endpoints**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/advanced-settings/evasive-path-match](https://techdocs.akamai.com/application-security/reference/put-evasive-path-match)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
        config_evasive_path_match = akamai.AppSecAdvancedSettingsEvasivePathMatch("configEvasivePathMatch",
            config_id=configuration.config_id,
            enable_path_match=True)
        # USE CASE: user wants to override the evasive path match setting for a security policy
        policy_override = akamai.AppSecAdvancedSettingsEvasivePathMatch("policyOverride",
            config_id=configuration.config_id,
            security_policy_id=var["security_policy_id"],
            enable_path_match=True)
        ```

        :param str resource_name: The name of the resource.
        :param AppSecAdvancedSettingsEvasivePathMatchArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecAdvancedSettingsEvasivePathMatchArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 enable_path_match: Optional[pulumi.Input[bool]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppSecAdvancedSettingsEvasivePathMatchArgs.__new__(AppSecAdvancedSettingsEvasivePathMatchArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if enable_path_match is None and not opts.urn:
                raise TypeError("Missing required property 'enable_path_match'")
            __props__.__dict__["enable_path_match"] = enable_path_match
            __props__.__dict__["security_policy_id"] = security_policy_id
        super(AppSecAdvancedSettingsEvasivePathMatch, __self__).__init__(
            'akamai:index/appSecAdvancedSettingsEvasivePathMatch:AppSecAdvancedSettingsEvasivePathMatch',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            enable_path_match: Optional[pulumi.Input[bool]] = None,
            security_policy_id: Optional[pulumi.Input[str]] = None) -> 'AppSecAdvancedSettingsEvasivePathMatch':
        """
        Get an existing AppSecAdvancedSettingsEvasivePathMatch resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[bool] enable_path_match: Whether to enable path match.
        :param pulumi.Input[str] security_policy_id: The ID of a specific security policy to which the evasive path match setting should be applied. If not supplied, the indicated setting will be applied to all policies within the configuration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecAdvancedSettingsEvasivePathMatchState.__new__(_AppSecAdvancedSettingsEvasivePathMatchState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["enable_path_match"] = enable_path_match
        __props__.__dict__["security_policy_id"] = security_policy_id
        return AppSecAdvancedSettingsEvasivePathMatch(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        The ID of the security configuration to use.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="enablePathMatch")
    def enable_path_match(self) -> pulumi.Output[bool]:
        """
        Whether to enable path match.
        """
        return pulumi.get(self, "enable_path_match")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of a specific security policy to which the evasive path match setting should be applied. If not supplied, the indicated setting will be applied to all policies within the configuration.
        """
        return pulumi.get(self, "security_policy_id")

