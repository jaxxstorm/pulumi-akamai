# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetAppSecBypassNetworkListsResult',
    'AwaitableGetAppSecBypassNetworkListsResult',
    'get_app_sec_bypass_network_lists',
]

@pulumi.output_type
class GetAppSecBypassNetworkListsResult:
    """
    A collection of values returned by getAppSecBypassNetworkLists.
    """
    def __init__(__self__, bypass_network_lists=None, config_id=None, id=None, json=None, output_text=None):
        if bypass_network_lists and not isinstance(bypass_network_lists, list):
            raise TypeError("Expected argument 'bypass_network_lists' to be a list")
        pulumi.set(__self__, "bypass_network_lists", bypass_network_lists)
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
    @pulumi.getter(name="bypassNetworkLists")
    def bypass_network_lists(self) -> Sequence[str]:
        """
        A list of strings containing the network list IDs.
        """
        return pulumi.get(self, "bypass_network_lists")

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
        A JSON-formatted list of information about the bypass network lists.
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter(name="outputText")
    def output_text(self) -> str:
        """
        A tabular display showing the bypass network list information.
        """
        return pulumi.get(self, "output_text")


class AwaitableGetAppSecBypassNetworkListsResult(GetAppSecBypassNetworkListsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecBypassNetworkListsResult(
            bypass_network_lists=self.bypass_network_lists,
            config_id=self.config_id,
            id=self.id,
            json=self.json,
            output_text=self.output_text)


def get_app_sec_bypass_network_lists(config_id: Optional[int] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecBypassNetworkListsResult:
    """
    Use the `AppSecByPassNetworkList` data source to retrieve information about which network
    lists are used in the bypass network lists settings.  The information available is described
    [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getbypassnetworklistsforawapconfigversion).
    Note: this data source is only applicable to WAP (Web Application Protector) configurations.

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
    bypass_network_lists = akamai.get_app_sec_bypass_network_lists(config_id=configuration.config_id)
    pulumi.export("bypassNetworkListsOutput", bypass_network_lists.output_text)
    pulumi.export("bypassNetworkListsJson", bypass_network_lists.json)
    pulumi.export("bypassNetworkListsIdList", bypass_network_lists.bypass_network_lists)
    ```


    :param int config_id: The configuration ID to use.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecBypassNetworkLists:getAppSecBypassNetworkLists', __args__, opts=opts, typ=GetAppSecBypassNetworkListsResult).value

    return AwaitableGetAppSecBypassNetworkListsResult(
        bypass_network_lists=__ret__.bypass_network_lists,
        config_id=__ret__.config_id,
        id=__ret__.id,
        json=__ret__.json,
        output_text=__ret__.output_text)