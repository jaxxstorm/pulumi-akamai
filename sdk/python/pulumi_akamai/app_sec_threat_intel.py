# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecThreatIntelArgs', 'AppSecThreatIntel']

@pulumi.input_type
class AppSecThreatIntelArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 security_policy_id: pulumi.Input[str],
                 threat_intel: pulumi.Input[str]):
        """
        The set of arguments for constructing a AppSecThreatIntel resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
        :param pulumi.Input[str] threat_intel: . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "security_policy_id", security_policy_id)
        pulumi.set(__self__, "threat_intel", threat_intel)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Input[str]:
        """
        . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_policy_id", value)

    @property
    @pulumi.getter(name="threatIntel")
    def threat_intel(self) -> pulumi.Input[str]:
        """
        . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
        """
        return pulumi.get(self, "threat_intel")

    @threat_intel.setter
    def threat_intel(self, value: pulumi.Input[str]):
        pulumi.set(self, "threat_intel", value)


@pulumi.input_type
class _AppSecThreatIntelState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 threat_intel: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AppSecThreatIntel resources.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
        :param pulumi.Input[str] threat_intel: . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if security_policy_id is not None:
            pulumi.set(__self__, "security_policy_id", security_policy_id)
        if threat_intel is not None:
            pulumi.set(__self__, "threat_intel", threat_intel)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_policy_id", value)

    @property
    @pulumi.getter(name="threatIntel")
    def threat_intel(self) -> Optional[pulumi.Input[str]]:
        """
        . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
        """
        return pulumi.get(self, "threat_intel")

    @threat_intel.setter
    def threat_intel(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "threat_intel", value)


class AppSecThreatIntel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 threat_intel: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        **Scopes**: Security policy

        Enables or disables threat intelligence for a security policy. This resource is only available to organizations running the Adaptive Security Engine (ASE) beta Please contact your Akamai representative for more information.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rules/threat-intel](https://techdocs.akamai.com/application-security/reference/put-rules-threat-intel)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        threat_intel = akamai.AppSecThreatIntel("threatIntel",
            config_id=configuration.config_id,
            security_policy_id="gms1_134637",
            threat_intel="on")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
        :param pulumi.Input[str] threat_intel: . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecThreatIntelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        **Scopes**: Security policy

        Enables or disables threat intelligence for a security policy. This resource is only available to organizations running the Adaptive Security Engine (ASE) beta Please contact your Akamai representative for more information.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rules/threat-intel](https://techdocs.akamai.com/application-security/reference/put-rules-threat-intel)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        threat_intel = akamai.AppSecThreatIntel("threatIntel",
            config_id=configuration.config_id,
            security_policy_id="gms1_134637",
            threat_intel="on")
        ```

        :param str resource_name: The name of the resource.
        :param AppSecThreatIntelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecThreatIntelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 threat_intel: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppSecThreatIntelArgs.__new__(AppSecThreatIntelArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if security_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_policy_id'")
            __props__.__dict__["security_policy_id"] = security_policy_id
            if threat_intel is None and not opts.urn:
                raise TypeError("Missing required property 'threat_intel'")
            __props__.__dict__["threat_intel"] = threat_intel
        super(AppSecThreatIntel, __self__).__init__(
            'akamai:index/appSecThreatIntel:AppSecThreatIntel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            security_policy_id: Optional[pulumi.Input[str]] = None,
            threat_intel: Optional[pulumi.Input[str]] = None) -> 'AppSecThreatIntel':
        """
        Get an existing AppSecThreatIntel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
        :param pulumi.Input[str] threat_intel: . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecThreatIntelState.__new__(_AppSecThreatIntelState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["security_policy_id"] = security_policy_id
        __props__.__dict__["threat_intel"] = threat_intel
        return AppSecThreatIntel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        . Unique identifier of the security configuration associated with the threat intelligence protection settings being modified.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Output[str]:
        """
        . Unique identifier of the security policy associated with the threat intelligence protection settings being modified.
        """
        return pulumi.get(self, "security_policy_id")

    @property
    @pulumi.getter(name="threatIntel")
    def threat_intel(self) -> pulumi.Output[str]:
        """
        . Set to `on` to enable threat intelligence protection; set to **off** to disable threat intelligence protection.
        """
        return pulumi.get(self, "threat_intel")

