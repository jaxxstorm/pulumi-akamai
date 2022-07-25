# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecSelectedHostnamesArgs', 'AppSecSelectedHostnames']

@pulumi.input_type
class AppSecSelectedHostnamesArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 hostnames: pulumi.Input[Sequence[pulumi.Input[str]]],
                 mode: pulumi.Input[str]):
        """
        The set of arguments for constructing a AppSecSelectedHostnames resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the hostnames.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hostnames: . JSON array of hostnames to be added or removed from the protected hosts list.
        :param pulumi.Input[str] mode: . Indicates how the `hostnames` array is to be applied. Allowed values are:
               - **APPEND**. Hosts listed in the `hostnames` array are added to the current list of selected hostnames.
               - **REPLACE**. Hosts listed in the `hostnames`  array overwrite the current list of selected hostnames: the “old” hostnames are replaced by the specified set of hostnames.
               - **REMOVE**, Hosts listed in the `hostnames` array are removed from the current list of select hostnames.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "hostnames", hostnames)
        pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        . Unique identifier of the security configuration associated with the hostnames.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter
    def hostnames(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        . JSON array of hostnames to be added or removed from the protected hosts list.
        """
        return pulumi.get(self, "hostnames")

    @hostnames.setter
    def hostnames(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "hostnames", value)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[str]:
        """
        . Indicates how the `hostnames` array is to be applied. Allowed values are:
        - **APPEND**. Hosts listed in the `hostnames` array are added to the current list of selected hostnames.
        - **REPLACE**. Hosts listed in the `hostnames`  array overwrite the current list of selected hostnames: the “old” hostnames are replaced by the specified set of hostnames.
        - **REMOVE**, Hosts listed in the `hostnames` array are removed from the current list of select hostnames.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "mode", value)


@pulumi.input_type
class _AppSecSelectedHostnamesState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AppSecSelectedHostnames resources.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the hostnames.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hostnames: . JSON array of hostnames to be added or removed from the protected hosts list.
        :param pulumi.Input[str] mode: . Indicates how the `hostnames` array is to be applied. Allowed values are:
               - **APPEND**. Hosts listed in the `hostnames` array are added to the current list of selected hostnames.
               - **REPLACE**. Hosts listed in the `hostnames`  array overwrite the current list of selected hostnames: the “old” hostnames are replaced by the specified set of hostnames.
               - **REMOVE**, Hosts listed in the `hostnames` array are removed from the current list of select hostnames.
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if hostnames is not None:
            pulumi.set(__self__, "hostnames", hostnames)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        . Unique identifier of the security configuration associated with the hostnames.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter
    def hostnames(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        . JSON array of hostnames to be added or removed from the protected hosts list.
        """
        return pulumi.get(self, "hostnames")

    @hostnames.setter
    def hostnames(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "hostnames", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        . Indicates how the `hostnames` array is to be applied. Allowed values are:
        - **APPEND**. Hosts listed in the `hostnames` array are added to the current list of selected hostnames.
        - **REPLACE**. Hosts listed in the `hostnames`  array overwrite the current list of selected hostnames: the “old” hostnames are replaced by the specified set of hostnames.
        - **REMOVE**, Hosts listed in the `hostnames` array are removed from the current list of select hostnames.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)


class AppSecSelectedHostnames(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        **Scopes**: Security configuration

        Modifies the list of hostnames protected under by a security configuration.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/selected-hostnames](https://techdocs.akamai.com/application-security/reference/put-selected-hostnames-per-config)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        appsecselectedhostnames = akamai.AppSecSelectedHostnames("appsecselectedhostnames",
            config_id=configuration.config_id,
            hostnames=["example.com"],
            mode="APPEND")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the hostnames.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hostnames: . JSON array of hostnames to be added or removed from the protected hosts list.
        :param pulumi.Input[str] mode: . Indicates how the `hostnames` array is to be applied. Allowed values are:
               - **APPEND**. Hosts listed in the `hostnames` array are added to the current list of selected hostnames.
               - **REPLACE**. Hosts listed in the `hostnames`  array overwrite the current list of selected hostnames: the “old” hostnames are replaced by the specified set of hostnames.
               - **REMOVE**, Hosts listed in the `hostnames` array are removed from the current list of select hostnames.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecSelectedHostnamesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        **Scopes**: Security configuration

        Modifies the list of hostnames protected under by a security configuration.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/selected-hostnames](https://techdocs.akamai.com/application-security/reference/put-selected-hostnames-per-config)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        appsecselectedhostnames = akamai.AppSecSelectedHostnames("appsecselectedhostnames",
            config_id=configuration.config_id,
            hostnames=["example.com"],
            mode="APPEND")
        ```

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
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppSecSelectedHostnamesArgs.__new__(AppSecSelectedHostnamesArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if hostnames is None and not opts.urn:
                raise TypeError("Missing required property 'hostnames'")
            __props__.__dict__["hostnames"] = hostnames
            if mode is None and not opts.urn:
                raise TypeError("Missing required property 'mode'")
            __props__.__dict__["mode"] = mode
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
            mode: Optional[pulumi.Input[str]] = None) -> 'AppSecSelectedHostnames':
        """
        Get an existing AppSecSelectedHostnames resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the hostnames.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hostnames: . JSON array of hostnames to be added or removed from the protected hosts list.
        :param pulumi.Input[str] mode: . Indicates how the `hostnames` array is to be applied. Allowed values are:
               - **APPEND**. Hosts listed in the `hostnames` array are added to the current list of selected hostnames.
               - **REPLACE**. Hosts listed in the `hostnames`  array overwrite the current list of selected hostnames: the “old” hostnames are replaced by the specified set of hostnames.
               - **REMOVE**, Hosts listed in the `hostnames` array are removed from the current list of select hostnames.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecSelectedHostnamesState.__new__(_AppSecSelectedHostnamesState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["hostnames"] = hostnames
        __props__.__dict__["mode"] = mode
        return AppSecSelectedHostnames(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        . Unique identifier of the security configuration associated with the hostnames.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter
    def hostnames(self) -> pulumi.Output[Sequence[str]]:
        """
        . JSON array of hostnames to be added or removed from the protected hosts list.
        """
        return pulumi.get(self, "hostnames")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        """
        . Indicates how the `hostnames` array is to be applied. Allowed values are:
        - **APPEND**. Hosts listed in the `hostnames` array are added to the current list of selected hostnames.
        - **REPLACE**. Hosts listed in the `hostnames`  array overwrite the current list of selected hostnames: the “old” hostnames are replaced by the specified set of hostnames.
        - **REMOVE**, Hosts listed in the `hostnames` array are removed from the current list of select hostnames.
        """
        return pulumi.get(self, "mode")

