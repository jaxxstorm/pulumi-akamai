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
    'GetIamRolesResult',
    'AwaitableGetIamRolesResult',
    'get_iam_roles',
]

@pulumi.output_type
class GetIamRolesResult:
    """
    A collection of values returned by getIamRoles.
    """
    def __init__(__self__, id=None, roles=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def roles(self) -> Sequence['outputs.GetIamRolesRoleResult']:
        return pulumi.get(self, "roles")


class AwaitableGetIamRolesResult(GetIamRolesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIamRolesResult(
            id=self.id,
            roles=self.roles)


def get_iam_roles(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIamRolesResult:
    """
    Use `get_iam_roles` to list roles for the current account and contract type. The account and contract type are determined by the access tokens in your API client. Use the `roleId` from this data source to construct the `auth_grants_json` when creating or updating a user's auth grants.

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    my_roles = akamai.get_iam_roles()
    pulumi.export("roles", my_roles)
    ```
    ## Attributes reference

    These attributes are returned:

    * `roles` — A list of roles.

    [API Reference](https://techdocs.akamai.com/iam-api/reference/get-roles)
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('akamai:index/getIamRoles:getIamRoles', __args__, opts=opts, typ=GetIamRolesResult).value

    return AwaitableGetIamRolesResult(
        id=__ret__.id,
        roles=__ret__.roles)
