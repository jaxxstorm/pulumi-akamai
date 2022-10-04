# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetGroupResult',
    'AwaitableGetGroupResult',
    'get_group',
    'get_group_output',
]

@pulumi.output_type
class GetGroupResult:
    """
    A collection of values returned by getGroup.
    """
    def __init__(__self__, contract=None, contract_id=None, group_name=None, id=None, name=None):
        if contract and not isinstance(contract, str):
            raise TypeError("Expected argument 'contract' to be a str")
        if contract is not None:
            warnings.warn("""The setting \"contract\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""contract is deprecated: The setting \"contract\" has been deprecated.""")

        pulumi.set(__self__, "contract", contract)
        if contract_id and not isinstance(contract_id, str):
            raise TypeError("Expected argument 'contract_id' to be a str")
        pulumi.set(__self__, "contract_id", contract_id)
        if group_name and not isinstance(group_name, str):
            raise TypeError("Expected argument 'group_name' to be a str")
        pulumi.set(__self__, "group_name", group_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        if name is not None:
            warnings.warn("""The setting \"name\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""name is deprecated: The setting \"name\" has been deprecated.""")

        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def contract(self) -> str:
        return pulumi.get(self, "contract")

    @property
    @pulumi.getter(name="contractId")
    def contract_id(self) -> str:
        return pulumi.get(self, "contract_id")

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

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


class AwaitableGetGroupResult(GetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupResult(
            contract=self.contract,
            contract_id=self.contract_id,
            group_name=self.group_name,
            id=self.id,
            name=self.name)


def get_group(contract: Optional[str] = None,
              contract_id: Optional[str] = None,
              group_name: Optional[str] = None,
              name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupResult:
    """
    Use the `get_group` data source to get a group by name.

    Each account features a hierarchy of groups, which control access to your
    Akamai configurations and help consolidate reporting functions, typically
    mapping to an organizational hierarchy. Using either Control Center or the
    [Identity Management: User Administration API](https://techdocs.akamai.com/iam-api/reference/api),
    account administrators can assign properties to specific groups, each with
    its own set of users and accompanying roles.

    ## Attributes reference

    This data source returns this attribute:

    * `id` - The group's unique ID, including the `grp_` prefix.


    :param str contract: Replaced by `contract_id`. Maintained for legacy purposes.
    :param str contract_id: - (Required) A contract's unique ID, including the `ctr_` prefix.
    :param str group_name: The group name.
    :param str name: Replaced by `group_name`. Maintained for legacy purposes.
    """
    __args__ = dict()
    __args__['contract'] = contract
    __args__['contractId'] = contract_id
    __args__['groupName'] = group_name
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('akamai:index/getGroup:getGroup', __args__, opts=opts, typ=GetGroupResult).value

    return AwaitableGetGroupResult(
        contract=__ret__.contract,
        contract_id=__ret__.contract_id,
        group_name=__ret__.group_name,
        id=__ret__.id,
        name=__ret__.name)


@_utilities.lift_output_func(get_group)
def get_group_output(contract: Optional[pulumi.Input[Optional[str]]] = None,
                     contract_id: Optional[pulumi.Input[Optional[str]]] = None,
                     group_name: Optional[pulumi.Input[Optional[str]]] = None,
                     name: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGroupResult]:
    """
    Use the `get_group` data source to get a group by name.

    Each account features a hierarchy of groups, which control access to your
    Akamai configurations and help consolidate reporting functions, typically
    mapping to an organizational hierarchy. Using either Control Center or the
    [Identity Management: User Administration API](https://techdocs.akamai.com/iam-api/reference/api),
    account administrators can assign properties to specific groups, each with
    its own set of users and accompanying roles.

    ## Attributes reference

    This data source returns this attribute:

    * `id` - The group's unique ID, including the `grp_` prefix.


    :param str contract: Replaced by `contract_id`. Maintained for legacy purposes.
    :param str contract_id: - (Required) A contract's unique ID, including the `ctr_` prefix.
    :param str group_name: The group name.
    :param str name: Replaced by `group_name`. Maintained for legacy purposes.
    """
    ...
