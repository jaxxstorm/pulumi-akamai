# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetPropertyResult:
    """
    A collection of values returned by getProperty.
    """
    def __init__(__self__, id=None, name=None, rules=None, version=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if rules and not isinstance(rules, str):
            raise TypeError("Expected argument 'rules' to be a str")
        __self__.rules = rules
        if version and not isinstance(version, float):
            raise TypeError("Expected argument 'version' to be a float")
        __self__.version = version


class AwaitableGetPropertyResult(GetPropertyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPropertyResult(
            id=self.id,
            name=self.name,
            rules=self.rules,
            version=self.version)


def get_property(name=None, version=None, opts=None):
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['version'] = version
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:properties/getProperty:getProperty', __args__, opts=opts).value

    return AwaitableGetPropertyResult(
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        rules=__ret__.get('rules'),
        version=__ret__.get('version'))