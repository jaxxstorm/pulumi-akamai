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
    'GetAppSecSelectedHostnamesResult',
    'AwaitableGetAppSecSelectedHostnamesResult',
    'get_app_sec_selected_hostnames',
    'get_app_sec_selected_hostnames_output',
]

@pulumi.output_type
class GetAppSecSelectedHostnamesResult:
    """
    A collection of values returned by getAppSecSelectedHostnames.
    """
    def __init__(__self__, config_id=None, hostnames=None, hostnames_json=None, id=None, output_text=None):
        if config_id and not isinstance(config_id, int):
            raise TypeError("Expected argument 'config_id' to be a int")
        pulumi.set(__self__, "config_id", config_id)
        if hostnames and not isinstance(hostnames, list):
            raise TypeError("Expected argument 'hostnames' to be a list")
        pulumi.set(__self__, "hostnames", hostnames)
        if hostnames_json and not isinstance(hostnames_json, str):
            raise TypeError("Expected argument 'hostnames_json' to be a str")
        pulumi.set(__self__, "hostnames_json", hostnames_json)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
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
    @pulumi.getter(name="hostnamesJson")
    def hostnames_json(self) -> str:
        return pulumi.get(self, "hostnames_json")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="outputText")
    def output_text(self) -> str:
        return pulumi.get(self, "output_text")


class AwaitableGetAppSecSelectedHostnamesResult(GetAppSecSelectedHostnamesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecSelectedHostnamesResult(
            config_id=self.config_id,
            hostnames=self.hostnames,
            hostnames_json=self.hostnames_json,
            id=self.id,
            output_text=self.output_text)


def get_app_sec_selected_hostnames(config_id: Optional[int] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecSelectedHostnamesResult:
    """
    **Scopes**: Security configuration

    Returns a list of the hostnames currently protected by the specified security configuration.

    **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/selected-hostnames](https://techdocs.akamai.com/application-security/reference/get-selected-hostnames)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    selected_hostnames_app_sec_selected_hostnames = akamai.get_app_sec_selected_hostnames(config_id=configuration.config_id)
    pulumi.export("selectedHostnames", selected_hostnames_app_sec_selected_hostnames.hostnames)
    pulumi.export("selectedHostnamesJson", selected_hostnames_app_sec_selected_hostnames.hostnames_json)
    pulumi.export("selectedHostnamesOutputText", selected_hostnames_app_sec_selected_hostnames.output_text)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `hostnames`. List of selected hostnames.
    - `hostnames_json`. JSON-formatted list of selected hostnames.
    - `output_text`. Tabular report of the selected hostnames.


    :param int config_id: . Unique identifier of the security configuration associated with the protected hosts.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecSelectedHostnames:getAppSecSelectedHostnames', __args__, opts=opts, typ=GetAppSecSelectedHostnamesResult).value

    return AwaitableGetAppSecSelectedHostnamesResult(
        config_id=__ret__.config_id,
        hostnames=__ret__.hostnames,
        hostnames_json=__ret__.hostnames_json,
        id=__ret__.id,
        output_text=__ret__.output_text)


@_utilities.lift_output_func(get_app_sec_selected_hostnames)
def get_app_sec_selected_hostnames_output(config_id: Optional[pulumi.Input[int]] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppSecSelectedHostnamesResult]:
    """
    **Scopes**: Security configuration

    Returns a list of the hostnames currently protected by the specified security configuration.

    **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/selected-hostnames](https://techdocs.akamai.com/application-security/reference/get-selected-hostnames)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    selected_hostnames_app_sec_selected_hostnames = akamai.get_app_sec_selected_hostnames(config_id=configuration.config_id)
    pulumi.export("selectedHostnames", selected_hostnames_app_sec_selected_hostnames.hostnames)
    pulumi.export("selectedHostnamesJson", selected_hostnames_app_sec_selected_hostnames.hostnames_json)
    pulumi.export("selectedHostnamesOutputText", selected_hostnames_app_sec_selected_hostnames.output_text)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `hostnames`. List of selected hostnames.
    - `hostnames_json`. JSON-formatted list of selected hostnames.
    - `output_text`. Tabular report of the selected hostnames.


    :param int config_id: . Unique identifier of the security configuration associated with the protected hosts.
    """
    ...
