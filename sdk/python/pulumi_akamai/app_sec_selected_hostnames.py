# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = ['AppSecSelectedHostnamesArgs', 'AppSecSelectedHostnames']

@pulumi.input_type
class AppSecSelectedHostnamesArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 hostnames: pulumi.Input[Sequence[pulumi.Input[str]]],
                 mode: pulumi.Input[str],
                 version: pulumi.Input[int]):
        """
        The set of arguments for constructing a AppSecSelectedHostnames resource.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "hostnames", hostnames)
        pulumi.set(__self__, "mode", mode)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter
    def hostnames(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "hostnames")

    @hostnames.setter
    def hostnames(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "hostnames", value)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[str]:
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[int]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[int]):
        pulumi.set(self, "version", value)


class AppSecSelectedHostnames(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a AppSecSelectedHostnames resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecSelectedHostnamesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AppSecSelectedHostnames resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AppSecSelectedHostnamesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecSelectedHostnamesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__['config_id'] = config_id
            if hostnames is None and not opts.urn:
                raise TypeError("Missing required property 'hostnames'")
            __props__['hostnames'] = hostnames
            if mode is None and not opts.urn:
                raise TypeError("Missing required property 'mode'")
            __props__['mode'] = mode
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__['version'] = version
        super(AppSecSelectedHostnames, __self__).__init__(
            'akamai:index/appSecSelectedHostnames:AppSecSelectedHostnames',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'AppSecSelectedHostnames':
        """
        Get an existing AppSecSelectedHostnames resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["config_id"] = config_id
        __props__["hostnames"] = hostnames
        __props__["mode"] = mode
        __props__["version"] = version
        return AppSecSelectedHostnames(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter
    def hostnames(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "hostnames")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

