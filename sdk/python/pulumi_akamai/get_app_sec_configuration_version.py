# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'GetAppSecConfigurationVersionResult',
    'AwaitableGetAppSecConfigurationVersionResult',
    'get_app_sec_configuration_version',
]

@pulumi.output_type
class GetAppSecConfigurationVersionResult:
    """
    A collection of values returned by getAppSecConfigurationVersion.
    """
    def __init__(__self__, config_id=None, id=None, latest_version=None, output_text=None, production_status=None, staging_status=None, version=None):
        if config_id and not isinstance(config_id, int):
            raise TypeError("Expected argument 'config_id' to be a int")
        pulumi.set(__self__, "config_id", config_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if latest_version and not isinstance(latest_version, int):
            raise TypeError("Expected argument 'latest_version' to be a int")
        pulumi.set(__self__, "latest_version", latest_version)
        if output_text and not isinstance(output_text, str):
            raise TypeError("Expected argument 'output_text' to be a str")
        pulumi.set(__self__, "output_text", output_text)
        if production_status and not isinstance(production_status, str):
            raise TypeError("Expected argument 'production_status' to be a str")
        pulumi.set(__self__, "production_status", production_status)
        if staging_status and not isinstance(staging_status, str):
            raise TypeError("Expected argument 'staging_status' to be a str")
        pulumi.set(__self__, "staging_status", staging_status)
        if version and not isinstance(version, int):
            raise TypeError("Expected argument 'version' to be a int")
        pulumi.set(__self__, "version", version)

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
        """
        The last version of the security configuration created.
        """
        return pulumi.get(self, "latest_version")

    @property
    @pulumi.getter(name="outputText")
    def output_text(self) -> str:
        """
        A tabular display showing the following information about all versions of the security configuration: version number, staging status, and production status.
        """
        return pulumi.get(self, "output_text")

    @property
    @pulumi.getter(name="productionStatus")
    def production_status(self) -> str:
        """
        The status of the specified version in production: "Active", "Inactive", or "Deactivated". Returned only if `version` was specified.
        """
        return pulumi.get(self, "production_status")

    @property
    @pulumi.getter(name="stagingStatus")
    def staging_status(self) -> str:
        """
        The status of the specified version in staging: "Active", "Inactive", or "Deactivated". Returned only if `version` was specified.
        """
        return pulumi.get(self, "staging_status")

    @property
    @pulumi.getter
    def version(self) -> Optional[int]:
        return pulumi.get(self, "version")


class AwaitableGetAppSecConfigurationVersionResult(GetAppSecConfigurationVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecConfigurationVersionResult(
            config_id=self.config_id,
            id=self.id,
            latest_version=self.latest_version,
            output_text=self.output_text,
            production_status=self.production_status,
            staging_status=self.staging_status,
            version=self.version)


def get_app_sec_configuration_version(config_id: Optional[int] = None,
                                      version: Optional[int] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecConfigurationVersionResult:
    """
    Use the `getAppSecConfigurationVersion` data source to retrieve information about the versions of a security configuration, or about a specific version.

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    specific_configuration = akamai.get_app_sec_configuration(name="Akamai Tools")
    versions = akamai.get_app_sec_configuration_version(config_id=specific_configuration.config_id)
    pulumi.export("versionsOutputText", versions.output_text)
    pulumi.export("versionsLatest", versions.latest_version)
    specific_version = akamai.get_app_sec_configuration_version(config_id=specific_configuration.config_id,
        version=42)
    pulumi.export("specificVersionVersion", specific_version.version)
    pulumi.export("specificVersionStaging", specific_version.staging_status)
    pulumi.export("specificVersionProduction", specific_version.production_status)
    ```


    :param int config_id: The ID of the security configuration to use.
    :param int version: The version number of the security configuration to use. If not supplied, information about all versions of the specified security configuration is returned.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    __args__['version'] = version
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecConfigurationVersion:getAppSecConfigurationVersion', __args__, opts=opts, typ=GetAppSecConfigurationVersionResult).value

    return AwaitableGetAppSecConfigurationVersionResult(
        config_id=__ret__.config_id,
        id=__ret__.id,
        latest_version=__ret__.latest_version,
        output_text=__ret__.output_text,
        production_status=__ret__.production_status,
        staging_status=__ret__.staging_status,
        version=__ret__.version)
