# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecConfigurationRenameArgs', 'AppSecConfigurationRename']

@pulumi.input_type
class AppSecConfigurationRenameArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 description: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AppSecConfigurationRename resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuring being renamed.
        :param pulumi.Input[str] description: . Brief description of the security configuration.
        :param pulumi.Input[str] name: . New name for the security configuration.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        . Unique identifier of the security configuring being renamed.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        . Brief description of the security configuration.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        . New name for the security configuration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _AppSecConfigurationRenameState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AppSecConfigurationRename resources.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuring being renamed.
        :param pulumi.Input[str] description: . Brief description of the security configuration.
        :param pulumi.Input[str] name: . New name for the security configuration.
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        . Unique identifier of the security configuring being renamed.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        . Brief description of the security configuration.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        . New name for the security configuration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class AppSecConfigurationRename(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        **Scopes**: Security configuration

        Renames an existing security configuration.
        Note that you can change only the configuration name. You can't modify the ID assigned to a security configuration.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}](https://techdocs.akamai.com/application-security/reference/put-config)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration_app_sec_configuration = akamai.get_app_sec_configuration(name="Documentation")
        configuration_app_sec_configuration_rename = akamai.AppSecConfigurationRename("configurationAppSecConfigurationRename",
            config_id=configuration_app_sec_configuration.config_id,
            description="This configuration is by both the documentation team and the training team.")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuring being renamed.
        :param pulumi.Input[str] description: . Brief description of the security configuration.
        :param pulumi.Input[str] name: . New name for the security configuration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecConfigurationRenameArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        **Scopes**: Security configuration

        Renames an existing security configuration.
        Note that you can change only the configuration name. You can't modify the ID assigned to a security configuration.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}](https://techdocs.akamai.com/application-security/reference/put-config)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration_app_sec_configuration = akamai.get_app_sec_configuration(name="Documentation")
        configuration_app_sec_configuration_rename = akamai.AppSecConfigurationRename("configurationAppSecConfigurationRename",
            config_id=configuration_app_sec_configuration.config_id,
            description="This configuration is by both the documentation team and the training team.")
        ```

        :param str resource_name: The name of the resource.
        :param AppSecConfigurationRenameArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecConfigurationRenameArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppSecConfigurationRenameArgs.__new__(AppSecConfigurationRenameArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
        super(AppSecConfigurationRename, __self__).__init__(
            'akamai:index/appSecConfigurationRename:AppSecConfigurationRename',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'AppSecConfigurationRename':
        """
        Get an existing AppSecConfigurationRename resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuring being renamed.
        :param pulumi.Input[str] description: . Brief description of the security configuration.
        :param pulumi.Input[str] name: . New name for the security configuration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecConfigurationRenameState.__new__(_AppSecConfigurationRenameState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        return AppSecConfigurationRename(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        . Unique identifier of the security configuring being renamed.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        . Brief description of the security configuration.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        . New name for the security configuration.
        """
        return pulumi.get(self, "name")

