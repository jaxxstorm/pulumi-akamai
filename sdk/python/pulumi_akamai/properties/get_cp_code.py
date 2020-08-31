# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetCpCodeResult',
    'AwaitableGetCpCodeResult',
    'get_cp_code',
]

@pulumi.output_type
class GetCpCodeResult:
    """
    A collection of values returned by getCpCode.
    """
    def __init__(__self__, contract=None, group=None, id=None, name=None):
        if contract and not isinstance(contract, str):
            raise TypeError("Expected argument 'contract' to be a str")
        pulumi.set(__self__, "contract", contract)
        if group and not isinstance(group, str):
            raise TypeError("Expected argument 'group' to be a str")
        pulumi.set(__self__, "group", group)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def contract(self) -> str:
        return pulumi.get(self, "contract")

    @property
    @pulumi.getter
    def group(self) -> str:
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


class AwaitableGetCpCodeResult(GetCpCodeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCpCodeResult(
            contract=self.contract,
            group=self.group,
            id=self.id,
            name=self.name)


def get_cp_code(contract: Optional[str] = None,
                group: Optional[str] = None,
                name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCpCodeResult:
    """
    Use `properties.CpCode` data source to retrieve a group id.

    ## Example Usage


    :param str contract: — (Required) The contract ID
    :param str group: — (Required) The group ID
    :param str name: — (Required) The CP code name.
    """
    __args__ = dict()
    __args__['contract'] = contract
    __args__['group'] = group
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:properties/getCpCode:getCpCode', __args__, opts=opts, typ=GetCpCodeResult).value

    return AwaitableGetCpCodeResult(
        contract=__ret__.contract,
        group=__ret__.group,
        id=__ret__.id,
        name=__ret__.name)
