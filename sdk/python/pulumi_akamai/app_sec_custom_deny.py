# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] custom_deny: The JSON-formatted definition of the custom deny action ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#63df3de3)).
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "custom_deny", custom_deny)

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
    @pulumi.getter(name="customDeny")
    def custom_deny(self) -> pulumi.Input[str]:
        """
        The JSON-formatted definition of the custom deny action ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#63df3de3)).
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
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] custom_deny: The JSON-formatted definition of the custom deny action ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#63df3de3)).
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
        The ID of the security configuration to use.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="customDeny")
    def custom_deny(self) -> Optional[pulumi.Input[str]]:
        """
        The JSON-formatted definition of the custom deny action ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#63df3de3)).
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
        The `resource_akamai_appsec_custom_deny` resource allows you to create a new custom deny action for a specific configuration.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
        custom_deny = akamai.AppSecCustomDeny("customDeny",
            config_id=configuration.config_id,
            custom_deny=(lambda path: open(path).read())(f"{path['module']}/custom_deny.json"))
        pulumi.export("customDenyId", custom_deny.custom_deny_id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] custom_deny: The JSON-formatted definition of the custom deny action ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#63df3de3)).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecCustomDenyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `resource_akamai_appsec_custom_deny` resource allows you to create a new custom deny action for a specific configuration.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
        custom_deny = akamai.AppSecCustomDeny("customDeny",
            config_id=configuration.config_id,
            custom_deny=(lambda path: open(path).read())(f"{path['module']}/custom_deny.json"))
        pulumi.export("customDenyId", custom_deny.custom_deny_id)
        ```

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
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
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
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] custom_deny: The JSON-formatted definition of the custom deny action ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#63df3de3)).
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
        The ID of the security configuration to use.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="customDeny")
    def custom_deny(self) -> pulumi.Output[str]:
        """
        The JSON-formatted definition of the custom deny action ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#63df3de3)).
        """
        return pulumi.get(self, "custom_deny")

    @property
    @pulumi.getter(name="customDenyId")
    def custom_deny_id(self) -> pulumi.Output[str]:
        """
        custom_deny_id
        """
        return pulumi.get(self, "custom_deny_id")

