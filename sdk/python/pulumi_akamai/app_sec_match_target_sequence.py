# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecMatchTargetSequenceArgs', 'AppSecMatchTargetSequence']

@pulumi.input_type
class AppSecMatchTargetSequenceArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 version: pulumi.Input[int],
                 match_target_sequence: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AppSecMatchTargetSequence resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[int] version: The version number of the security configuration to use.
        :param pulumi.Input[str] match_target_sequence: The name of a JSON file containing the sequence of all match targets defined for the specified security configuration version ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsequence)).
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "version", version)
        if match_target_sequence is not None:
            pulumi.set(__self__, "match_target_sequence", match_target_sequence)

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
    @pulumi.getter
    def version(self) -> pulumi.Input[int]:
        """
        The version number of the security configuration to use.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[int]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter(name="matchTargetSequence")
    def match_target_sequence(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a JSON file containing the sequence of all match targets defined for the specified security configuration version ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsequence)).
        """
        return pulumi.get(self, "match_target_sequence")

    @match_target_sequence.setter
    def match_target_sequence(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_target_sequence", value)


@pulumi.input_type
class _AppSecMatchTargetSequenceState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 match_target_sequence: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering AppSecMatchTargetSequence resources.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] match_target_sequence: The name of a JSON file containing the sequence of all match targets defined for the specified security configuration version ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsequence)).
        :param pulumi.Input[int] version: The version number of the security configuration to use.
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if match_target_sequence is not None:
            pulumi.set(__self__, "match_target_sequence", match_target_sequence)
        if version is not None:
            pulumi.set(__self__, "version", version)

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
    @pulumi.getter(name="matchTargetSequence")
    def match_target_sequence(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a JSON file containing the sequence of all match targets defined for the specified security configuration version ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsequence)).
        """
        return pulumi.get(self, "match_target_sequence")

    @match_target_sequence.setter
    def match_target_sequence(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_target_sequence", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        The version number of the security configuration to use.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class AppSecMatchTargetSequence(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 match_target_sequence: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        The `AppSecMatchTargetSequence` resource allows you to specify the order in which match targets are applied within a given security configuration version.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Akamai Tools")
        match_target_sequence = akamai.AppSecMatchTargetSequence("matchTargetSequence",
            config_id=configuration.config_id,
            version=configuration.latest_version,
            match_target_sequence=(lambda path: open(path).read())(f"{path['module']}/match_targets.json"))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] match_target_sequence: The name of a JSON file containing the sequence of all match targets defined for the specified security configuration version ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsequence)).
        :param pulumi.Input[int] version: The version number of the security configuration to use.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecMatchTargetSequenceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `AppSecMatchTargetSequence` resource allows you to specify the order in which match targets are applied within a given security configuration version.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Akamai Tools")
        match_target_sequence = akamai.AppSecMatchTargetSequence("matchTargetSequence",
            config_id=configuration.config_id,
            version=configuration.latest_version,
            match_target_sequence=(lambda path: open(path).read())(f"{path['module']}/match_targets.json"))
        ```

        :param str resource_name: The name of the resource.
        :param AppSecMatchTargetSequenceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecMatchTargetSequenceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 match_target_sequence: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
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
            __props__ = AppSecMatchTargetSequenceArgs.__new__(AppSecMatchTargetSequenceArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            __props__.__dict__["match_target_sequence"] = match_target_sequence
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__.__dict__["version"] = version
        super(AppSecMatchTargetSequence, __self__).__init__(
            'akamai:index/appSecMatchTargetSequence:AppSecMatchTargetSequence',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            match_target_sequence: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'AppSecMatchTargetSequence':
        """
        Get an existing AppSecMatchTargetSequence resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] match_target_sequence: The name of a JSON file containing the sequence of all match targets defined for the specified security configuration version ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsequence)).
        :param pulumi.Input[int] version: The version number of the security configuration to use.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecMatchTargetSequenceState.__new__(_AppSecMatchTargetSequenceState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["match_target_sequence"] = match_target_sequence
        __props__.__dict__["version"] = version
        return AppSecMatchTargetSequence(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        The ID of the security configuration to use.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="matchTargetSequence")
    def match_target_sequence(self) -> pulumi.Output[Optional[str]]:
        """
        The name of a JSON file containing the sequence of all match targets defined for the specified security configuration version ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsequence)).
        """
        return pulumi.get(self, "match_target_sequence")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        The version number of the security configuration to use.
        """
        return pulumi.get(self, "version")

