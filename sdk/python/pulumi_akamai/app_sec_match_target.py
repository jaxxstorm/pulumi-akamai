# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecMatchTargetArgs', 'AppSecMatchTarget']

@pulumi.input_type
class AppSecMatchTargetArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 match_target: pulumi.Input[str]):
        """
        The set of arguments for constructing a AppSecMatchTarget resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the match target being modified.
        :param pulumi.Input[str] match_target: . Path to a JSON file containing one or more match target definitions.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "match_target", match_target)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        . Unique identifier of the security configuration associated with the match target being modified.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="matchTarget")
    def match_target(self) -> pulumi.Input[str]:
        """
        . Path to a JSON file containing one or more match target definitions.
        """
        return pulumi.get(self, "match_target")

    @match_target.setter
    def match_target(self, value: pulumi.Input[str]):
        pulumi.set(self, "match_target", value)


@pulumi.input_type
class _AppSecMatchTargetState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 match_target: Optional[pulumi.Input[str]] = None,
                 match_target_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering AppSecMatchTarget resources.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the match target being modified.
        :param pulumi.Input[str] match_target: . Path to a JSON file containing one or more match target definitions.
        :param pulumi.Input[int] match_target_id: Unique identifier of the match target
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if match_target is not None:
            pulumi.set(__self__, "match_target", match_target)
        if match_target_id is not None:
            pulumi.set(__self__, "match_target_id", match_target_id)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        . Unique identifier of the security configuration associated with the match target being modified.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="matchTarget")
    def match_target(self) -> Optional[pulumi.Input[str]]:
        """
        . Path to a JSON file containing one or more match target definitions.
        """
        return pulumi.get(self, "match_target")

    @match_target.setter
    def match_target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_target", value)

    @property
    @pulumi.getter(name="matchTargetId")
    def match_target_id(self) -> Optional[pulumi.Input[int]]:
        """
        Unique identifier of the match target
        """
        return pulumi.get(self, "match_target_id")

    @match_target_id.setter
    def match_target_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "match_target_id", value)


class AppSecMatchTarget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 match_target: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        **Scopes**: Security configuration

        Creates a match target associated with a security configuration. Match targets determine which security policy should apply to an API, hostname or path.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/match-targets](https://techdocs.akamai.com/application-security/reference/post-match-targets)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        match_target = akamai.AppSecMatchTarget("matchTarget",
            config_id=configuration.config_id,
            match_target=(lambda path: open(path).read())(f"{path['module']}/match_targets.json"))
        ```
        ## Output Options

        In addition to the arguments above, the following attribute is exported:

        - `match_target_id`. ID of the match target.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the match target being modified.
        :param pulumi.Input[str] match_target: . Path to a JSON file containing one or more match target definitions.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecMatchTargetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        **Scopes**: Security configuration

        Creates a match target associated with a security configuration. Match targets determine which security policy should apply to an API, hostname or path.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/match-targets](https://techdocs.akamai.com/application-security/reference/post-match-targets)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        match_target = akamai.AppSecMatchTarget("matchTarget",
            config_id=configuration.config_id,
            match_target=(lambda path: open(path).read())(f"{path['module']}/match_targets.json"))
        ```
        ## Output Options

        In addition to the arguments above, the following attribute is exported:

        - `match_target_id`. ID of the match target.

        :param str resource_name: The name of the resource.
        :param AppSecMatchTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecMatchTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 match_target: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppSecMatchTargetArgs.__new__(AppSecMatchTargetArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if match_target is None and not opts.urn:
                raise TypeError("Missing required property 'match_target'")
            __props__.__dict__["match_target"] = match_target
            __props__.__dict__["match_target_id"] = None
        super(AppSecMatchTarget, __self__).__init__(
            'akamai:index/appSecMatchTarget:AppSecMatchTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            match_target: Optional[pulumi.Input[str]] = None,
            match_target_id: Optional[pulumi.Input[int]] = None) -> 'AppSecMatchTarget':
        """
        Get an existing AppSecMatchTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the match target being modified.
        :param pulumi.Input[str] match_target: . Path to a JSON file containing one or more match target definitions.
        :param pulumi.Input[int] match_target_id: Unique identifier of the match target
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecMatchTargetState.__new__(_AppSecMatchTargetState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["match_target"] = match_target
        __props__.__dict__["match_target_id"] = match_target_id
        return AppSecMatchTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        . Unique identifier of the security configuration associated with the match target being modified.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="matchTarget")
    def match_target(self) -> pulumi.Output[str]:
        """
        . Path to a JSON file containing one or more match target definitions.
        """
        return pulumi.get(self, "match_target")

    @property
    @pulumi.getter(name="matchTargetId")
    def match_target_id(self) -> pulumi.Output[int]:
        """
        Unique identifier of the match target
        """
        return pulumi.get(self, "match_target_id")

