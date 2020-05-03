# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetContractResult:
    """
    A collection of values returned by getContract.
    """
    def __init__(__self__, group=None, id=None):
        if group and not isinstance(group, str):
            raise TypeError("Expected argument 'group' to be a str")
        __self__.group = group
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
class AwaitableGetContractResult(GetContractResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContractResult(
            group=self.group,
            id=self.id)

def get_contract(group=None,opts=None):
    """
    Use `.getContract` data source to retrieve a group id.




    :param str group: — (Optional) The group within which the contract can be found.
    """
    __args__ = dict()


    __args__['group'] = group
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:index/getContract:getContract', __args__, opts=opts).value

    return AwaitableGetContractResult(
        group=__ret__.get('group'),
        id=__ret__.get('id'))
