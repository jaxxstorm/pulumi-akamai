# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetAppSecReputationProfilesResult',
    'AwaitableGetAppSecReputationProfilesResult',
    'get_app_sec_reputation_profiles',
]

@pulumi.output_type
class GetAppSecReputationProfilesResult:
    """
    A collection of values returned by getAppSecReputationProfiles.
    """
    def __init__(__self__, config_id=None, id=None, json=None, output_text=None, reputation_profile_id=None):
        if config_id and not isinstance(config_id, int):
            raise TypeError("Expected argument 'config_id' to be a int")
        pulumi.set(__self__, "config_id", config_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if json and not isinstance(json, str):
            raise TypeError("Expected argument 'json' to be a str")
        pulumi.set(__self__, "json", json)
        if output_text and not isinstance(output_text, str):
            raise TypeError("Expected argument 'output_text' to be a str")
        pulumi.set(__self__, "output_text", output_text)
        if reputation_profile_id and not isinstance(reputation_profile_id, int):
            raise TypeError("Expected argument 'reputation_profile_id' to be a int")
        pulumi.set(__self__, "reputation_profile_id", reputation_profile_id)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> int:
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def json(self) -> str:
        """
        A JSON-formatted display of the details about the indicated reputation profile or profiles.
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter(name="outputText")
    def output_text(self) -> str:
        """
        A tabular display of the details about the indicated reputation profile or profiles.
        """
        return pulumi.get(self, "output_text")

    @property
    @pulumi.getter(name="reputationProfileId")
    def reputation_profile_id(self) -> Optional[int]:
        return pulumi.get(self, "reputation_profile_id")


class AwaitableGetAppSecReputationProfilesResult(GetAppSecReputationProfilesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecReputationProfilesResult(
            config_id=self.config_id,
            id=self.id,
            json=self.json,
            output_text=self.output_text,
            reputation_profile_id=self.reputation_profile_id)


def get_app_sec_reputation_profiles(config_id: Optional[int] = None,
                                    reputation_profile_id: Optional[int] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecReputationProfilesResult:
    """
    Use the `getAppSecReputationProfiles` data source to retrieve details about all reputation profiles, or a specific reputation profiles.

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
    reputation_profiles = akamai.get_app_sec_reputation_profiles(config_id=configuration.config_id)
    pulumi.export("reputationProfilesOutput", reputation_profiles.output_text)
    pulumi.export("reputationProfilesJson", reputation_profiles.json)
    reputation_profile = akamai.get_app_sec_reputation_profiles(config_id=configuration.config_id,
        reputation_profile_id=var["reputation_profile_id"])
    pulumi.export("reputationProfileJson", reputation_profile.json)
    pulumi.export("reputationProfileOutput", reputation_profile.output_text)
    ```


    :param int config_id: The ID of the security configuration to use.
    :param int reputation_profile_id: The ID of a given reputation profile. If not supplied, information about all reputation profiles is returned.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    __args__['reputationProfileId'] = reputation_profile_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecReputationProfiles:getAppSecReputationProfiles', __args__, opts=opts, typ=GetAppSecReputationProfilesResult).value

    return AwaitableGetAppSecReputationProfilesResult(
        config_id=__ret__.config_id,
        id=__ret__.id,
        json=__ret__.json,
        output_text=__ret__.output_text,
        reputation_profile_id=__ret__.reputation_profile_id)
