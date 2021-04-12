# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = [
    'GetContractResult',
    'AwaitableGetContractResult',
    'get_contract',
]

@pulumi.output_type
class GetContractResult:
    """
    A collection of values returned by getContract.
    """
    def __init__(__self__, group=None, group_id=None, group_name=None, id=None):
        if group and not isinstance(group, str):
            raise TypeError("Expected argument 'group' to be a str")
        if group is not None:
            warnings.warn("""The setting \"group\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""group is deprecated: The setting \"group\" has been deprecated.""")

        pulumi.set(__self__, "group", group)
        if group_id and not isinstance(group_id, str):
            raise TypeError("Expected argument 'group_id' to be a str")
        pulumi.set(__self__, "group_id", group_id)
        if group_name and not isinstance(group_name, str):
            raise TypeError("Expected argument 'group_name' to be a str")
        pulumi.set(__self__, "group_name", group_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def group(self) -> Optional[str]:
        return pulumi.get(self, "group")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> str:
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> str:
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetContractResult(GetContractResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContractResult(
            group=self.group,
            group_id=self.group_id,
            group_name=self.group_name,
            id=self.id)


def get_contract(group: Optional[str] = None,
                 group_id: Optional[str] = None,
                 group_name: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContractResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['group'] = group
    __args__['groupId'] = group_id
    __args__['groupName'] = group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:index/getContract:getContract', __args__, opts=opts, typ=GetContractResult).value

    return AwaitableGetContractResult(
        group=__ret__.group,
        group_id=__ret__.group_id,
        group_name=__ret__.group_name,
        id=__ret__.id)
