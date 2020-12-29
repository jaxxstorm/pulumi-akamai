# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'GetDnsRecordSetResult',
    'AwaitableGetDnsRecordSetResult',
    'get_dns_record_set',
]

@pulumi.output_type
class GetDnsRecordSetResult:
    """
    A collection of values returned by getDnsRecordSet.
    """
    def __init__(__self__, host=None, id=None, rdatas=None, record_type=None, zone=None):
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if rdatas and not isinstance(rdatas, list):
            raise TypeError("Expected argument 'rdatas' to be a list")
        pulumi.set(__self__, "rdatas", rdatas)
        if record_type and not isinstance(record_type, str):
            raise TypeError("Expected argument 'record_type' to be a str")
        pulumi.set(__self__, "record_type", record_type)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def host(self) -> str:
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def rdatas(self) -> Sequence[str]:
        return pulumi.get(self, "rdatas")

    @property
    @pulumi.getter(name="recordType")
    def record_type(self) -> str:
        return pulumi.get(self, "record_type")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetDnsRecordSetResult(GetDnsRecordSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDnsRecordSetResult(
            host=self.host,
            id=self.id,
            rdatas=self.rdatas,
            record_type=self.record_type,
            zone=self.zone)


def get_dns_record_set(host: Optional[str] = None,
                       record_type: Optional[str] = None,
                       zone: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDnsRecordSetResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['host'] = host
    __args__['recordType'] = record_type
    __args__['zone'] = zone
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:index/getDnsRecordSet:getDnsRecordSet', __args__, opts=opts, typ=GetDnsRecordSetResult).value

    return AwaitableGetDnsRecordSetResult(
        host=__ret__.host,
        id=__ret__.id,
        rdatas=__ret__.rdatas,
        record_type=__ret__.record_type,
        zone=__ret__.zone)
