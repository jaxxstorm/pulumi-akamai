# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecConfigurationVersionCloneArgs', 'AppSecConfigurationVersionClone']

@pulumi.input_type
class AppSecConfigurationVersionCloneArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 create_from_version: pulumi.Input[int],
                 rule_update: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AppSecConfigurationVersionClone resource.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "create_from_version", create_from_version)
        if rule_update is not None:
            pulumi.set(__self__, "rule_update", rule_update)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="createFromVersion")
    def create_from_version(self) -> pulumi.Input[int]:
        return pulumi.get(self, "create_from_version")

    @create_from_version.setter
    def create_from_version(self, value: pulumi.Input[int]):
        pulumi.set(self, "create_from_version", value)

    @property
    @pulumi.getter(name="ruleUpdate")
    def rule_update(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "rule_update")

    @rule_update.setter
    def rule_update(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "rule_update", value)


@pulumi.input_type
class _AppSecConfigurationVersionCloneState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 create_from_version: Optional[pulumi.Input[int]] = None,
                 rule_update: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering AppSecConfigurationVersionClone resources.
        :param pulumi.Input[int] version: Version of cloned configuration
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if create_from_version is not None:
            pulumi.set(__self__, "create_from_version", create_from_version)
        if rule_update is not None:
            pulumi.set(__self__, "rule_update", rule_update)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="createFromVersion")
    def create_from_version(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "create_from_version")

    @create_from_version.setter
    def create_from_version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "create_from_version", value)

    @property
    @pulumi.getter(name="ruleUpdate")
    def rule_update(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "rule_update")

    @rule_update.setter
    def rule_update(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "rule_update", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        Version of cloned configuration
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class AppSecConfigurationVersionClone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 create_from_version: Optional[pulumi.Input[int]] = None,
                 rule_update: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a AppSecConfigurationVersionClone resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecConfigurationVersionCloneArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AppSecConfigurationVersionClone resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AppSecConfigurationVersionCloneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecConfigurationVersionCloneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 create_from_version: Optional[pulumi.Input[int]] = None,
                 rule_update: Optional[pulumi.Input[bool]] = None,
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
            __props__ = AppSecConfigurationVersionCloneArgs.__new__(AppSecConfigurationVersionCloneArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if create_from_version is None and not opts.urn:
                raise TypeError("Missing required property 'create_from_version'")
            __props__.__dict__["create_from_version"] = create_from_version
            __props__.__dict__["rule_update"] = rule_update
            __props__.__dict__["version"] = None
        super(AppSecConfigurationVersionClone, __self__).__init__(
            'akamai:index/appSecConfigurationVersionClone:AppSecConfigurationVersionClone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            create_from_version: Optional[pulumi.Input[int]] = None,
            rule_update: Optional[pulumi.Input[bool]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'AppSecConfigurationVersionClone':
        """
        Get an existing AppSecConfigurationVersionClone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] version: Version of cloned configuration
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecConfigurationVersionCloneState.__new__(_AppSecConfigurationVersionCloneState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["create_from_version"] = create_from_version
        __props__.__dict__["rule_update"] = rule_update
        __props__.__dict__["version"] = version
        return AppSecConfigurationVersionClone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="createFromVersion")
    def create_from_version(self) -> pulumi.Output[int]:
        return pulumi.get(self, "create_from_version")

    @property
    @pulumi.getter(name="ruleUpdate")
    def rule_update(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "rule_update")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        Version of cloned configuration
        """
        return pulumi.get(self, "version")

