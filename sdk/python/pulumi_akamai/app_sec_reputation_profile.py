# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecReputationProfileArgs', 'AppSecReputationProfile']

@pulumi.input_type
class AppSecReputationProfileArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 reputation_profile: pulumi.Input[str]):
        """
        The set of arguments for constructing a AppSecReputationProfile resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] reputation_profile: The name of a file containing a JSON-formatted definition of the reputation profile. ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postreputationprofiles))
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "reputation_profile", reputation_profile)

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
    @pulumi.getter(name="reputationProfile")
    def reputation_profile(self) -> pulumi.Input[str]:
        """
        The name of a file containing a JSON-formatted definition of the reputation profile. ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postreputationprofiles))
        """
        return pulumi.get(self, "reputation_profile")

    @reputation_profile.setter
    def reputation_profile(self, value: pulumi.Input[str]):
        pulumi.set(self, "reputation_profile", value)


@pulumi.input_type
class _AppSecReputationProfileState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 reputation_profile: Optional[pulumi.Input[str]] = None,
                 reputation_profile_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering AppSecReputationProfile resources.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] reputation_profile: The name of a file containing a JSON-formatted definition of the reputation profile. ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postreputationprofiles))
        :param pulumi.Input[int] reputation_profile_id: The ID of the newly created or modified reputation profile.
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if reputation_profile is not None:
            pulumi.set(__self__, "reputation_profile", reputation_profile)
        if reputation_profile_id is not None:
            pulumi.set(__self__, "reputation_profile_id", reputation_profile_id)

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
    @pulumi.getter(name="reputationProfile")
    def reputation_profile(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a file containing a JSON-formatted definition of the reputation profile. ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postreputationprofiles))
        """
        return pulumi.get(self, "reputation_profile")

    @reputation_profile.setter
    def reputation_profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reputation_profile", value)

    @property
    @pulumi.getter(name="reputationProfileId")
    def reputation_profile_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the newly created or modified reputation profile.
        """
        return pulumi.get(self, "reputation_profile_id")

    @reputation_profile_id.setter
    def reputation_profile_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reputation_profile_id", value)


class AppSecReputationProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 reputation_profile: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use the `AppSecReputationProfile` resource to create or modify a reputation profile for a specific security configuration.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
        # USE CASE: user wants to create a reputation profile for a given configuration and version, using a JSON definition
        reputation_profile = akamai.AppSecReputationProfile("reputationProfile",
            config_id=configuration.config_id,
            reputation_profile=(lambda path: open(path).read())(f"{path['module']}/reputation_profile.json"))
        pulumi.export("reputationProfileId", akamai_appsec_reputation_profile["reputation_profile_id"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] reputation_profile: The name of a file containing a JSON-formatted definition of the reputation profile. ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postreputationprofiles))
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecReputationProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use the `AppSecReputationProfile` resource to create or modify a reputation profile for a specific security configuration.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
        # USE CASE: user wants to create a reputation profile for a given configuration and version, using a JSON definition
        reputation_profile = akamai.AppSecReputationProfile("reputationProfile",
            config_id=configuration.config_id,
            reputation_profile=(lambda path: open(path).read())(f"{path['module']}/reputation_profile.json"))
        pulumi.export("reputationProfileId", akamai_appsec_reputation_profile["reputation_profile_id"])
        ```

        :param str resource_name: The name of the resource.
        :param AppSecReputationProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecReputationProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 reputation_profile: Optional[pulumi.Input[str]] = None,
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
            __props__ = AppSecReputationProfileArgs.__new__(AppSecReputationProfileArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if reputation_profile is None and not opts.urn:
                raise TypeError("Missing required property 'reputation_profile'")
            __props__.__dict__["reputation_profile"] = reputation_profile
            __props__.__dict__["reputation_profile_id"] = None
        super(AppSecReputationProfile, __self__).__init__(
            'akamai:index/appSecReputationProfile:AppSecReputationProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            reputation_profile: Optional[pulumi.Input[str]] = None,
            reputation_profile_id: Optional[pulumi.Input[int]] = None) -> 'AppSecReputationProfile':
        """
        Get an existing AppSecReputationProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] reputation_profile: The name of a file containing a JSON-formatted definition of the reputation profile. ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postreputationprofiles))
        :param pulumi.Input[int] reputation_profile_id: The ID of the newly created or modified reputation profile.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecReputationProfileState.__new__(_AppSecReputationProfileState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["reputation_profile"] = reputation_profile
        __props__.__dict__["reputation_profile_id"] = reputation_profile_id
        return AppSecReputationProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        The ID of the security configuration to use.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="reputationProfile")
    def reputation_profile(self) -> pulumi.Output[str]:
        """
        The name of a file containing a JSON-formatted definition of the reputation profile. ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postreputationprofiles))
        """
        return pulumi.get(self, "reputation_profile")

    @property
    @pulumi.getter(name="reputationProfileId")
    def reputation_profile_id(self) -> pulumi.Output[int]:
        """
        The ID of the newly created or modified reputation profile.
        """
        return pulumi.get(self, "reputation_profile_id")

