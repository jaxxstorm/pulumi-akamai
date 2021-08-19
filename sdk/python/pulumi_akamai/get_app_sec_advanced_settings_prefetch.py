# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetAppSecAdvancedSettingsPrefetchResult',
    'AwaitableGetAppSecAdvancedSettingsPrefetchResult',
    'get_app_sec_advanced_settings_prefetch',
]

@pulumi.output_type
class GetAppSecAdvancedSettingsPrefetchResult:
    """
    A collection of values returned by getAppSecAdvancedSettingsPrefetch.
    """
    def __init__(__self__, config_id=None, id=None, json=None, output_text=None):
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
        A JSON-formatted list of information about the prefetch request settings.
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter(name="outputText")
    def output_text(self) -> str:
        """
        A tabular display showing the prefetch request settings.
        """
        return pulumi.get(self, "output_text")


class AwaitableGetAppSecAdvancedSettingsPrefetchResult(GetAppSecAdvancedSettingsPrefetchResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecAdvancedSettingsPrefetchResult(
            config_id=self.config_id,
            id=self.id,
            json=self.json,
            output_text=self.output_text)


def get_app_sec_advanced_settings_prefetch(config_id: Optional[int] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecAdvancedSettingsPrefetchResult:
    """
    Use the `AppSecAdvancedSettingsPrefetch` data source to retrieve information the prefetch request settings for a security configuration. The information available is described [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getprefetchrequestsforaconfiguration).

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
    prefetch = akamai.get_app_sec_advanced_settings_prefetch(config_id=configuration.config_id)
    pulumi.export("advancedSettingsPrefetchOutput", prefetch.output_text)
    pulumi.export("advancedSettingsPrefetchJson", prefetch.json)
    ```


    :param int config_id: The configuration ID.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecAdvancedSettingsPrefetch:getAppSecAdvancedSettingsPrefetch', __args__, opts=opts, typ=GetAppSecAdvancedSettingsPrefetchResult).value

    return AwaitableGetAppSecAdvancedSettingsPrefetchResult(
        config_id=__ret__.config_id,
        id=__ret__.id,
        json=__ret__.json,
        output_text=__ret__.output_text)