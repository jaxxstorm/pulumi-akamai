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
    'GetAppSecAdvancedSettingsPrefetchResult',
    'AwaitableGetAppSecAdvancedSettingsPrefetchResult',
    'get_app_sec_advanced_settings_prefetch',
    'get_app_sec_advanced_settings_prefetch_output',
]

@pulumi.output_type
class GetAppSecAdvancedSettingsPrefetchResult:
    """
    A collection of values returned by getAppSecAdvancedSettingsPrefetch.
    """
    def __init__(__self__, config_id=None, id=None, json=None, output_text=None):
        if config_id and not isinstance(config_id, int):
            raise TypeError("Expected argument 'config_id' to be a int")
        pulumi.set(__self__, "config_id", config_id)
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


class AwaitableGetAppSecAdvancedSettingsPrefetchResult(GetAppSecAdvancedSettingsPrefetchResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecAdvancedSettingsPrefetchResult(
            config_id=self.config_id,
            id=self.id,
            json=self.json,
            output_text=self.output_text)


def get_app_sec_advanced_settings_prefetch(config_id: Optional[int] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecAdvancedSettingsPrefetchResult:
    """
    **Scopes**: Security configuration

    Returns information about your prefetch request settings. By default, Web Application Firewall inspects only external requests — requests originating outside of your firewall or Akamai's edge servers. When prefetch is enabled, requests between your origin servers and Akamai's edge servers can also be inspected by the firewall.

    **Related** **API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/advanced-settings/prefetch](https://techdocs.akamai.com/application-security/reference/get-advanced-settings-prefetch)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    prefetch = akamai.get_app_sec_advanced_settings_prefetch(config_id=configuration.config_id)
    pulumi.export("advancedSettingsPrefetchOutput", prefetch.output_text)
    pulumi.export("advancedSettingsPrefetchJson", prefetch.json)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `json`. JSON-formatted list of information about the prefetch request settings.
    - `output_text`. Tabular report showing the prefetch request settings.


    :param int config_id: . Unique identifier of the security configuration associated with the prefetch settings.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecAdvancedSettingsPrefetch:getAppSecAdvancedSettingsPrefetch', __args__, opts=opts, typ=GetAppSecAdvancedSettingsPrefetchResult).value

    return AwaitableGetAppSecAdvancedSettingsPrefetchResult(
        config_id=__ret__.config_id,
        id=__ret__.id,
        json=__ret__.json,
        output_text=__ret__.output_text)


@_utilities.lift_output_func(get_app_sec_advanced_settings_prefetch)
def get_app_sec_advanced_settings_prefetch_output(config_id: Optional[pulumi.Input[int]] = None,
                                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppSecAdvancedSettingsPrefetchResult]:
    """
    **Scopes**: Security configuration

    Returns information about your prefetch request settings. By default, Web Application Firewall inspects only external requests — requests originating outside of your firewall or Akamai's edge servers. When prefetch is enabled, requests between your origin servers and Akamai's edge servers can also be inspected by the firewall.

    **Related** **API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/advanced-settings/prefetch](https://techdocs.akamai.com/application-security/reference/get-advanced-settings-prefetch)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    prefetch = akamai.get_app_sec_advanced_settings_prefetch(config_id=configuration.config_id)
    pulumi.export("advancedSettingsPrefetchOutput", prefetch.output_text)
    pulumi.export("advancedSettingsPrefetchJson", prefetch.json)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `json`. JSON-formatted list of information about the prefetch request settings.
    - `output_text`. Tabular report showing the prefetch request settings.


    :param int config_id: . Unique identifier of the security configuration associated with the prefetch settings.
    """
    ...
