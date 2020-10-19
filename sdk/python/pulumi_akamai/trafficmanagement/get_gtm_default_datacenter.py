# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetGtmDefaultDatacenterResult',
    'AwaitableGetGtmDefaultDatacenterResult',
    'get_gtm_default_datacenter',
]

@pulumi.output_type
class GetGtmDefaultDatacenterResult:
    """
    A collection of values returned by getGtmDefaultDatacenter.
    """
    def __init__(__self__, datacenter=None, datacenter_id=None, domain=None, id=None, nickname=None):
        if datacenter and not isinstance(datacenter, int):
            raise TypeError("Expected argument 'datacenter' to be a int")
        pulumi.set(__self__, "datacenter", datacenter)
        if datacenter_id and not isinstance(datacenter_id, int):
            raise TypeError("Expected argument 'datacenter_id' to be a int")
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if nickname and not isinstance(nickname, str):
            raise TypeError("Expected argument 'nickname' to be a str")
        pulumi.set(__self__, "nickname", nickname)

    @property
    @pulumi.getter
    def datacenter(self) -> Optional[int]:
        return pulumi.get(self, "datacenter")

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> int:
        return pulumi.get(self, "datacenter_id")

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def nickname(self) -> str:
        return pulumi.get(self, "nickname")


class AwaitableGetGtmDefaultDatacenterResult(GetGtmDefaultDatacenterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGtmDefaultDatacenterResult(
            datacenter=self.datacenter,
            datacenter_id=self.datacenter_id,
            domain=self.domain,
            id=self.id,
            nickname=self.nickname)


def get_gtm_default_datacenter(datacenter: Optional[int] = None,
                               domain: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGtmDefaultDatacenterResult:
    """
    Use `trafficmanagement.getGtmDefaultDatacenter` data source to retrieve default datacenter id and nickname.

    ## Example Usage


    :param int datacenter: — (Optional. Default 5400)
    :param str domain: — (Required)
    """
    __args__ = dict()
    __args__['datacenter'] = datacenter
    __args__['domain'] = domain
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:trafficmanagement/getGtmDefaultDatacenter:getGtmDefaultDatacenter', __args__, opts=opts, typ=GetGtmDefaultDatacenterResult).value

    return AwaitableGetGtmDefaultDatacenterResult(
        datacenter=__ret__.datacenter,
        datacenter_id=__ret__.datacenter_id,
        domain=__ret__.domain,
        id=__ret__.id,
        nickname=__ret__.nickname)
