# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecEvalHostnamesArgs', 'AppSecEvalHostnames']

@pulumi.input_type
class AppSecEvalHostnamesArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 hostnames: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a AppSecEvalHostnames resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration in evaluation mode.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hostnames: . JSON array of hostnames to be used in the evaluation process. Note that this list replaces your existing list of evaluation hosts.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "hostnames", hostnames)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        . Unique identifier of the security configuration in evaluation mode.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter
    def hostnames(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        . JSON array of hostnames to be used in the evaluation process. Note that this list replaces your existing list of evaluation hosts.
        """
        return pulumi.get(self, "hostnames")

    @hostnames.setter
    def hostnames(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "hostnames", value)


@pulumi.input_type
class _AppSecEvalHostnamesState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering AppSecEvalHostnames resources.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration in evaluation mode.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hostnames: . JSON array of hostnames to be used in the evaluation process. Note that this list replaces your existing list of evaluation hosts.
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if hostnames is not None:
            pulumi.set(__self__, "hostnames", hostnames)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        . Unique identifier of the security configuration in evaluation mode.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter
    def hostnames(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        . JSON array of hostnames to be used in the evaluation process. Note that this list replaces your existing list of evaluation hosts.
        """
        return pulumi.get(self, "hostnames")

    @hostnames.setter
    def hostnames(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "hostnames", value)


class AppSecEvalHostnames(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        **Scopes**: Security configuration

        Modifies the list of hostnames evaluated while a security configuration is in evaluation mode.
        During evaluation mode, hosts take no action of any kind when responding to traffic.
        Instead, these hosts simply maintain a record of the actions they *would* have taken if they had been responding to live traffic in your production network.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/selected-hostnames/eval-hostnames](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putevaluationhostnames)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        eval_hostnames = akamai.AppSecEvalHostnames("evalHostnames",
            config_id=configuration.config_id,
            hostnames=[
                "documentation.akamai.com",
                "training.akamai.com",
                "videos.akamai.com",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration in evaluation mode.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hostnames: . JSON array of hostnames to be used in the evaluation process. Note that this list replaces your existing list of evaluation hosts.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecEvalHostnamesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        **Scopes**: Security configuration

        Modifies the list of hostnames evaluated while a security configuration is in evaluation mode.
        During evaluation mode, hosts take no action of any kind when responding to traffic.
        Instead, these hosts simply maintain a record of the actions they *would* have taken if they had been responding to live traffic in your production network.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/selected-hostnames/eval-hostnames](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putevaluationhostnames)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        eval_hostnames = akamai.AppSecEvalHostnames("evalHostnames",
            config_id=configuration.config_id,
            hostnames=[
                "documentation.akamai.com",
                "training.akamai.com",
                "videos.akamai.com",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param AppSecEvalHostnamesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecEvalHostnamesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = AppSecEvalHostnamesArgs.__new__(AppSecEvalHostnamesArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if hostnames is None and not opts.urn:
                raise TypeError("Missing required property 'hostnames'")
            __props__.__dict__["hostnames"] = hostnames
        super(AppSecEvalHostnames, __self__).__init__(
            'akamai:index/appSecEvalHostnames:AppSecEvalHostnames',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'AppSecEvalHostnames':
        """
        Get an existing AppSecEvalHostnames resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration in evaluation mode.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hostnames: . JSON array of hostnames to be used in the evaluation process. Note that this list replaces your existing list of evaluation hosts.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecEvalHostnamesState.__new__(_AppSecEvalHostnamesState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["hostnames"] = hostnames
        return AppSecEvalHostnames(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        . Unique identifier of the security configuration in evaluation mode.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter
    def hostnames(self) -> pulumi.Output[Sequence[str]]:
        """
        . JSON array of hostnames to be used in the evaluation process. Note that this list replaces your existing list of evaluation hosts.
        """
        return pulumi.get(self, "hostnames")

