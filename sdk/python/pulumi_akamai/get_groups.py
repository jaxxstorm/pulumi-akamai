# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetGroupsResult',
    'AwaitableGetGroupsResult',
    'get_groups',
]

@pulumi.output_type
class GetGroupsResult:
    """
    A collection of values returned by getGroups.
    """
    def __init__(__self__, groups=None, id=None):
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetGroupsGroupResult']:
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetGroupsResult(GetGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupsResult(
            groups=self.groups,
            id=self.id)


def get_groups(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupsResult:
    """
    Use the `get_groups` data source to list groups associated with the [EdgeGrid API client token](https://techdocs.akamai.com/developer/docs/authenticate-with-edgegrid) you're using.

    ## Basic usage

    Return groups associated with the EdgeGrid API client token you're using:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    my_example = akamai.get_groups()
    pulumi.export("propertyMatch", my_example)
    ```

    ## Attributes reference

    This data source returns these attributes:

    * `groups` - A list of supported groups, with the following attributes:
      * `group_id` - A group's unique ID, including the `grp_` prefix.
      * `group_name` - The name of the group.
      * `parent_group_id` - The ID of the parent group, if applicable.
      * `contract_ids` - An array of strings listing the contract IDs for each group.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('akamai:index/getGroups:getGroups', __args__, opts=opts, typ=GetGroupsResult).value

    return AwaitableGetGroupsResult(
        groups=__ret__.groups,
        id=__ret__.id)
