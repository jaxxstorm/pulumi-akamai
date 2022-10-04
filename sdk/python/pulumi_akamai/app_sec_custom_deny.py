# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecCustomDenyArgs', 'AppSecCustomDeny']

@pulumi.input_type
class AppSecCustomDenyArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 custom_deny: pulumi.Input[str]):
        """
        The set of arguments for constructing a AppSecCustomDeny resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the custom deny.
        :param pulumi.Input[str] custom_deny: . Path to a JSON file containing properties and property values for the custom deny.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "custom_deny", custom_deny)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        . Unique identifier of the security configuration associated with the custom deny.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="customDeny")
    def custom_deny(self) -> pulumi.Input[str]:
        """
        . Path to a JSON file containing properties and property values for the custom deny.
        """
        return pulumi.get(self, "custom_deny")

    @custom_deny.setter
    def custom_deny(self, value: pulumi.Input[str]):
        pulumi.set(self, "custom_deny", value)


@pulumi.input_type
class _AppSecCustomDenyState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 custom_deny: Optional[pulumi.Input[str]] = None,
                 custom_deny_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AppSecCustomDeny resources.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the custom deny.
        :param pulumi.Input[str] custom_deny: . Path to a JSON file containing properties and property values for the custom deny.
        :param pulumi.Input[str] custom_deny_id: custom_deny_id
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if custom_deny is not None:
            pulumi.set(__self__, "custom_deny", custom_deny)
        if custom_deny_id is not None:
            pulumi.set(__self__, "custom_deny_id", custom_deny_id)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        . Unique identifier of the security configuration associated with the custom deny.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="customDeny")
    def custom_deny(self) -> Optional[pulumi.Input[str]]:
        """
        . Path to a JSON file containing properties and property values for the custom deny.
        """
        return pulumi.get(self, "custom_deny")

    @custom_deny.setter
    def custom_deny(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_deny", value)

    @property
    @pulumi.getter(name="customDenyId")
    def custom_deny_id(self) -> Optional[pulumi.Input[str]]:
        """
        custom_deny_id
        """
        return pulumi.get(self, "custom_deny_id")

    @custom_deny_id.setter
    def custom_deny_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_deny_id", value)


class AppSecCustomDeny(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 custom_deny: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        **Scopes**: Custom deny

        Modifies a custom deny action. Custom denies enable you to craft your own error message or redirect pages for use when HTTP requests are denied.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/custom-deny](https://techdocs.akamai.com/application-security/reference/get-custom-deny-actions)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        custom_deny = akamai.AppSecCustomDeny("customDeny",
            config_id=configuration.config_id,
            custom_deny=(lambda path: open(path).read())(f"{path['module']}/custom_deny.json"))
        pulumi.export("customDenyId", custom_deny.custom_deny_id)
        ```
        ## Output Options

        The following options can be used to determine the information returned, and how that returned information is formatted:

        - `custom_deny_id`. ID of the new custom deny action.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the custom deny.
        :param pulumi.Input[str] custom_deny: . Path to a JSON file containing properties and property values for the custom deny.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecCustomDenyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        **Scopes**: Custom deny

        Modifies a custom deny action. Custom denies enable you to craft your own error message or redirect pages for use when HTTP requests are denied.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/custom-deny](https://techdocs.akamai.com/application-security/reference/get-custom-deny-actions)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        custom_deny = akamai.AppSecCustomDeny("customDeny",
            config_id=configuration.config_id,
            custom_deny=(lambda path: open(path).read())(f"{path['module']}/custom_deny.json"))
        pulumi.export("customDenyId", custom_deny.custom_deny_id)
        ```
        ## Output Options

        The following options can be used to determine the information returned, and how that returned information is formatted:

        - `custom_deny_id`. ID of the new custom deny action.

        :param str resource_name: The name of the resource.
        :param AppSecCustomDenyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecCustomDenyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 custom_deny: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppSecCustomDenyArgs.__new__(AppSecCustomDenyArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if custom_deny is None and not opts.urn:
                raise TypeError("Missing required property 'custom_deny'")
            __props__.__dict__["custom_deny"] = custom_deny
            __props__.__dict__["custom_deny_id"] = None
        super(AppSecCustomDeny, __self__).__init__(
            'akamai:index/appSecCustomDeny:AppSecCustomDeny',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            custom_deny: Optional[pulumi.Input[str]] = None,
            custom_deny_id: Optional[pulumi.Input[str]] = None) -> 'AppSecCustomDeny':
        """
        Get an existing AppSecCustomDeny resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the custom deny.
        :param pulumi.Input[str] custom_deny: . Path to a JSON file containing properties and property values for the custom deny.
        :param pulumi.Input[str] custom_deny_id: custom_deny_id
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecCustomDenyState.__new__(_AppSecCustomDenyState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["custom_deny"] = custom_deny
        __props__.__dict__["custom_deny_id"] = custom_deny_id
        return AppSecCustomDeny(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        . Unique identifier of the security configuration associated with the custom deny.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="customDeny")
    def custom_deny(self) -> pulumi.Output[str]:
        """
        . Path to a JSON file containing properties and property values for the custom deny.
        """
        return pulumi.get(self, "custom_deny")

    @property
    @pulumi.getter(name="customDenyId")
    def custom_deny_id(self) -> pulumi.Output[str]:
        """
        custom_deny_id
        """
        return pulumi.get(self, "custom_deny_id")

