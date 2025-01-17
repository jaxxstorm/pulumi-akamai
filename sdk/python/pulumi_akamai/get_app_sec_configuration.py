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
    'GetAppSecConfigurationResult',
    'AwaitableGetAppSecConfigurationResult',
    'get_app_sec_configuration',
    'get_app_sec_configuration_output',
]

@pulumi.output_type
class GetAppSecConfigurationResult:
    """
    A collection of values returned by getAppSecConfiguration.
    """
    def __init__(__self__, config_id=None, id=None, latest_version=None, name=None, output_text=None, production_version=None, staging_version=None):
        if config_id and not isinstance(config_id, int):
            raise TypeError("Expected argument 'config_id' to be a int")
        pulumi.set(__self__, "config_id", config_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if latest_version and not isinstance(latest_version, int):
            raise TypeError("Expected argument 'latest_version' to be a int")
        pulumi.set(__self__, "latest_version", latest_version)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if output_text and not isinstance(output_text, str):
            raise TypeError("Expected argument 'output_text' to be a str")
        pulumi.set(__self__, "output_text", output_text)
        if production_version and not isinstance(production_version, int):
            raise TypeError("Expected argument 'production_version' to be a int")
        pulumi.set(__self__, "production_version", production_version)
        if staging_version and not isinstance(staging_version, int):
            raise TypeError("Expected argument 'staging_version' to be a int")
        pulumi.set(__self__, "staging_version", staging_version)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> int:
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="latestVersion")
    def latest_version(self) -> int:
        return pulumi.get(self, "latest_version")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outputText")
    def output_text(self) -> str:
        return pulumi.get(self, "output_text")

    @property
    @pulumi.getter(name="productionVersion")
    def production_version(self) -> int:
        return pulumi.get(self, "production_version")

    @property
    @pulumi.getter(name="stagingVersion")
    def staging_version(self) -> int:
        return pulumi.get(self, "staging_version")


class AwaitableGetAppSecConfigurationResult(GetAppSecConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecConfigurationResult(
            config_id=self.config_id,
            id=self.id,
            latest_version=self.latest_version,
            name=self.name,
            output_text=self.output_text,
            production_version=self.production_version,
            staging_version=self.staging_version)


def get_app_sec_configuration(name: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecConfigurationResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: . Name of the security configuration you want to return information for. If not included, information is returned for all your security configurations.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecConfiguration:getAppSecConfiguration', __args__, opts=opts, typ=GetAppSecConfigurationResult).value

    return AwaitableGetAppSecConfigurationResult(
        config_id=__ret__.config_id,
        id=__ret__.id,
        latest_version=__ret__.latest_version,
        name=__ret__.name,
        output_text=__ret__.output_text,
        production_version=__ret__.production_version,
        staging_version=__ret__.staging_version)


@_utilities.lift_output_func(get_app_sec_configuration)
def get_app_sec_configuration_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppSecConfigurationResult]:
    """
    Use this data source to access information about an existing resource.

    :param str name: . Name of the security configuration you want to return information for. If not included, information is returned for all your security configurations.
    """
    ...
