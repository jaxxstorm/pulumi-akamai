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
    'GetAppSecFailoverHostnamesResult',
    'AwaitableGetAppSecFailoverHostnamesResult',
    'get_app_sec_failover_hostnames',
    'get_app_sec_failover_hostnames_output',
]

@pulumi.output_type
class GetAppSecFailoverHostnamesResult:
    """
    A collection of values returned by getAppSecFailoverHostnames.
    """
    def __init__(__self__, config_id=None, hostnames=None, id=None, json=None, output_text=None):
        if config_id and not isinstance(config_id, int):
            raise TypeError("Expected argument 'config_id' to be a int")
        pulumi.set(__self__, "config_id", config_id)
        if hostnames and not isinstance(hostnames, list):
            raise TypeError("Expected argument 'hostnames' to be a list")
        pulumi.set(__self__, "hostnames", hostnames)
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
    def hostnames(self) -> Sequence[str]:
        return pulumi.get(self, "hostnames")

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


class AwaitableGetAppSecFailoverHostnamesResult(GetAppSecFailoverHostnamesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecFailoverHostnamesResult(
            config_id=self.config_id,
            hostnames=self.hostnames,
            id=self.id,
            json=self.json,
            output_text=self.output_text)


def get_app_sec_failover_hostnames(config_id: Optional[int] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecFailoverHostnamesResult:
    """
    **Scopes**: Security configuration

    Returns a list of the failover hostnames in a configuration.

    **Related API Endpoint**: [/appsec/v1/configs/{configId}/failover-hostnames](https://techdocs.akamai.com/application-security/reference/get-failover-hostnames)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    failover_hostnames_app_sec_failover_hostnames = akamai.get_app_sec_failover_hostnames(config_id=configuration.config_id)
    pulumi.export("failoverHostnames", failover_hostnames_app_sec_failover_hostnames.hostnames)
    pulumi.export("failoverHostnamesOutput", failover_hostnames_app_sec_failover_hostnames.output_text)
    pulumi.export("failoverHostnamesJson", failover_hostnames_app_sec_failover_hostnames.json)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `hostnames`. List of the failover hostnames.
    - `json`. JSON-formatted list of the failover hostnames.


    :param int config_id: . Unique identifier of the security configuration associated with the failover hosts.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecFailoverHostnames:getAppSecFailoverHostnames', __args__, opts=opts, typ=GetAppSecFailoverHostnamesResult).value

    return AwaitableGetAppSecFailoverHostnamesResult(
        config_id=__ret__.config_id,
        hostnames=__ret__.hostnames,
        id=__ret__.id,
        json=__ret__.json,
        output_text=__ret__.output_text)


@_utilities.lift_output_func(get_app_sec_failover_hostnames)
def get_app_sec_failover_hostnames_output(config_id: Optional[pulumi.Input[int]] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppSecFailoverHostnamesResult]:
    """
    **Scopes**: Security configuration

    Returns a list of the failover hostnames in a configuration.

    **Related API Endpoint**: [/appsec/v1/configs/{configId}/failover-hostnames](https://techdocs.akamai.com/application-security/reference/get-failover-hostnames)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    failover_hostnames_app_sec_failover_hostnames = akamai.get_app_sec_failover_hostnames(config_id=configuration.config_id)
    pulumi.export("failoverHostnames", failover_hostnames_app_sec_failover_hostnames.hostnames)
    pulumi.export("failoverHostnamesOutput", failover_hostnames_app_sec_failover_hostnames.output_text)
    pulumi.export("failoverHostnamesJson", failover_hostnames_app_sec_failover_hostnames.json)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `hostnames`. List of the failover hostnames.
    - `json`. JSON-formatted list of the failover hostnames.


    :param int config_id: . Unique identifier of the security configuration associated with the failover hosts.
    """
    ...
