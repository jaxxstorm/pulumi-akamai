# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetPropertyResult',
    'AwaitableGetPropertyResult',
    'get_property',
    'get_property_output',
]

warnings.warn("""akamai.properties.getProperty has been deprecated in favor of akamai.getProperty""", DeprecationWarning)

@pulumi.output_type
class GetPropertyResult:
    """
    A collection of values returned by getProperty.
    """
    def __init__(__self__, id=None, name=None, rules=None, version=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if rules and not isinstance(rules, str):
            raise TypeError("Expected argument 'rules' to be a str")
        pulumi.set(__self__, "rules", rules)
        if version and not isinstance(version, int):
            raise TypeError("Expected argument 'version' to be a int")
        pulumi.set(__self__, "version", version)

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

    @property
    @pulumi.getter
    def rules(self) -> str:
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def version(self) -> Optional[int]:
        return pulumi.get(self, "version")


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


def get_property(name: Optional[str] = None,
                 version: Optional[int] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPropertyResult:
    """
    Use this data source to access information about an existing resource.
    """
    pulumi.log.warn("""get_property is deprecated: akamai.properties.getProperty has been deprecated in favor of akamai.getProperty""")
    __args__ = dict()
    __args__['name'] = name
    __args__['version'] = version
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:properties/getProperty:getProperty', __args__, opts=opts, typ=GetPropertyResult).value

    return AwaitableGetPropertyResult(
        id=__ret__.id,
        name=__ret__.name,
        rules=__ret__.rules,
        version=__ret__.version)


@_utilities.lift_output_func(get_property)
def get_property_output(name: Optional[pulumi.Input[str]] = None,
                        version: Optional[pulumi.Input[Optional[int]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPropertyResult]:
    """
    Use this data source to access information about an existing resource.
    """
    pulumi.log.warn("""get_property is deprecated: akamai.properties.getProperty has been deprecated in favor of akamai.getProperty""")
    ...
