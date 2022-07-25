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
    'GetAppSecHostnameCoverageOverlappingResult',
    'AwaitableGetAppSecHostnameCoverageOverlappingResult',
    'get_app_sec_hostname_coverage_overlapping',
    'get_app_sec_hostname_coverage_overlapping_output',
]

@pulumi.output_type
class GetAppSecHostnameCoverageOverlappingResult:
    """
    A collection of values returned by getAppSecHostnameCoverageOverlapping.
    """
    def __init__(__self__, config_id=None, hostname=None, id=None, json=None, output_text=None):
        if config_id and not isinstance(config_id, int):
            raise TypeError("Expected argument 'config_id' to be a int")
        pulumi.set(__self__, "config_id", config_id)
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if json and not isinstance(json, str):
            raise TypeError("Expected argument 'json' to be a str")
        pulumi.set(__self__, "json", json)
        if output_text and not isinstance(output_text, str):
            raise TypeError("Expected argument 'output_text' to be a str")
        pulumi.set(__self__, "output_text", output_text)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> int:
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def json(self) -> str:
        return pulumi.get(self, "json")

    @property
    @pulumi.getter(name="outputText")
    def output_text(self) -> str:
        return pulumi.get(self, "output_text")


class AwaitableGetAppSecHostnameCoverageOverlappingResult(GetAppSecHostnameCoverageOverlappingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecHostnameCoverageOverlappingResult(
            config_id=self.config_id,
            hostname=self.hostname,
            id=self.id,
            json=self.json,
            output_text=self.output_text)


def get_app_sec_hostname_coverage_overlapping(config_id: Optional[int] = None,
                                              hostname: Optional[str] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecHostnameCoverageOverlappingResult:
    """
    **Scopes**: Security configuration; hostname

    Returns information about any other configuration versions that contain a hostname found in the current configuration version. The returned information is described in the [List hostname overlaps](https://developer.akamai.com/api/cloud_security/application_security/v1.html#gethostnamecoverageoverlapping) section of the Application Security API.

    **Related API Endpoint**:[/appsec/v1/configs/{configId}/versions/{versionNumber}/hostname-coverage/overlapping](https://techdocs.akamai.com/application-security/reference/get-hostname-coverage-overlapping)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    test = akamai.get_app_sec_hostname_coverage_overlapping(config_id=configuration.config_id,
        hostname="documentation.akamai.com")
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `json`. JSON-formatted list of the overlap information.
    - `output_text`. Tabular report of the overlap information.


    :param int config_id: . Unique identifier of the security configuration you want to return information for.
    :param str hostname: . Name of the host you want to return information for.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    __args__['hostname'] = hostname
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecHostnameCoverageOverlapping:getAppSecHostnameCoverageOverlapping', __args__, opts=opts, typ=GetAppSecHostnameCoverageOverlappingResult).value

    return AwaitableGetAppSecHostnameCoverageOverlappingResult(
        config_id=__ret__.config_id,
        hostname=__ret__.hostname,
        id=__ret__.id,
        json=__ret__.json,
        output_text=__ret__.output_text)


@_utilities.lift_output_func(get_app_sec_hostname_coverage_overlapping)
def get_app_sec_hostname_coverage_overlapping_output(config_id: Optional[pulumi.Input[int]] = None,
                                                     hostname: Optional[pulumi.Input[str]] = None,
                                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppSecHostnameCoverageOverlappingResult]:
    """
    **Scopes**: Security configuration; hostname

    Returns information about any other configuration versions that contain a hostname found in the current configuration version. The returned information is described in the [List hostname overlaps](https://developer.akamai.com/api/cloud_security/application_security/v1.html#gethostnamecoverageoverlapping) section of the Application Security API.

    **Related API Endpoint**:[/appsec/v1/configs/{configId}/versions/{versionNumber}/hostname-coverage/overlapping](https://techdocs.akamai.com/application-security/reference/get-hostname-coverage-overlapping)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    test = akamai.get_app_sec_hostname_coverage_overlapping(config_id=configuration.config_id,
        hostname="documentation.akamai.com")
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `json`. JSON-formatted list of the overlap information.
    - `output_text`. Tabular report of the overlap information.


    :param int config_id: . Unique identifier of the security configuration you want to return information for.
    :param str hostname: . Name of the host you want to return information for.
    """
    ...
