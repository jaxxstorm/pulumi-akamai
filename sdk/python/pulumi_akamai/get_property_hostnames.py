# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetPropertyHostnamesResult',
    'AwaitableGetPropertyHostnamesResult',
    'get_property_hostnames',
]

@pulumi.output_type
class GetPropertyHostnamesResult:
    """
    A collection of values returned by getPropertyHostnames.
    """
    def __init__(__self__, contract_id=None, group_id=None, hostnames=None, id=None, property_id=None, version=None):
        if contract_id and not isinstance(contract_id, str):
            raise TypeError("Expected argument 'contract_id' to be a str")
        pulumi.set(__self__, "contract_id", contract_id)
        if group_id and not isinstance(group_id, str):
            raise TypeError("Expected argument 'group_id' to be a str")
        pulumi.set(__self__, "group_id", group_id)
        if hostnames and not isinstance(hostnames, list):
            raise TypeError("Expected argument 'hostnames' to be a list")
        pulumi.set(__self__, "hostnames", hostnames)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if property_id and not isinstance(property_id, str):
            raise TypeError("Expected argument 'property_id' to be a str")
        pulumi.set(__self__, "property_id", property_id)
        if version and not isinstance(version, int):
            raise TypeError("Expected argument 'version' to be a int")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="contractId")
    def contract_id(self) -> str:
        return pulumi.get(self, "contract_id")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> str:
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def hostnames(self) -> Sequence['outputs.GetPropertyHostnamesHostnameResult']:
        return pulumi.get(self, "hostnames")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="propertyId")
    def property_id(self) -> str:
        return pulumi.get(self, "property_id")

    @property
    @pulumi.getter
    def version(self) -> int:
        return pulumi.get(self, "version")


class AwaitableGetPropertyHostnamesResult(GetPropertyHostnamesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPropertyHostnamesResult(
            contract_id=self.contract_id,
            group_id=self.group_id,
            hostnames=self.hostnames,
            id=self.id,
            property_id=self.property_id,
            version=self.version)


def get_property_hostnames(contract_id: Optional[str] = None,
                           group_id: Optional[str] = None,
                           property_id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPropertyHostnamesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['contractId'] = contract_id
    __args__['groupId'] = group_id
    __args__['propertyId'] = property_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:index/getPropertyHostnames:getPropertyHostnames', __args__, opts=opts, typ=GetPropertyHostnamesResult).value

    return AwaitableGetPropertyHostnamesResult(
        contract_id=__ret__.contract_id,
        group_id=__ret__.group_id,
        hostnames=__ret__.hostnames,
        id=__ret__.id,
        property_id=__ret__.property_id,
        version=__ret__.version)