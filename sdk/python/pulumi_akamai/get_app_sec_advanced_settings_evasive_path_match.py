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
    'GetAppSecAdvancedSettingsEvasivePathMatchResult',
    'AwaitableGetAppSecAdvancedSettingsEvasivePathMatchResult',
    'get_app_sec_advanced_settings_evasive_path_match',
    'get_app_sec_advanced_settings_evasive_path_match_output',
]

@pulumi.output_type
class GetAppSecAdvancedSettingsEvasivePathMatchResult:
    """
    A collection of values returned by getAppSecAdvancedSettingsEvasivePathMatch.
    """
    def __init__(__self__, config_id=None, id=None, json=None, output_text=None, security_policy_id=None):
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
        if security_policy_id and not isinstance(security_policy_id, str):
            raise TypeError("Expected argument 'security_policy_id' to be a str")
        pulumi.set(__self__, "security_policy_id", security_policy_id)

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
        """
        A JSON-formatted list of information about the evasive path match settings.
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter(name="outputText")
    def output_text(self) -> str:
        """
        A tabular display showing the evasive path match settings.
        """
        return pulumi.get(self, "output_text")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> Optional[str]:
        return pulumi.get(self, "security_policy_id")


class AwaitableGetAppSecAdvancedSettingsEvasivePathMatchResult(GetAppSecAdvancedSettingsEvasivePathMatchResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecAdvancedSettingsEvasivePathMatchResult(
            config_id=self.config_id,
            id=self.id,
            json=self.json,
            output_text=self.output_text,
            security_policy_id=self.security_policy_id)


def get_app_sec_advanced_settings_evasive_path_match(config_id: Optional[int] = None,
                                                     security_policy_id: Optional[str] = None,
                                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecAdvancedSettingsEvasivePathMatchResult:
    """
    **Scopes**: Security configuration; security policy

    Use the `AppSecAdvancedSettingsEvasivePathMatch` data source to retrieve information about the evasive path match for a configuration. This operation applies at the configuration level, and therefore applies to all policies within a configuration. You may retrieve these settings for a particular policy by specifying the policy using the `security_policy_id` parameter. For more information, see [Get evasive path match setting](https://techdocs.akamai.com/application-security/reference/get-evasive-path-match).

    **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/advanced-settings/evasive-path-match](https://techdocs.akamai.com/application-security/reference/get-evasive-path-match)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
    evasive_path_match = akamai.get_app_sec_advanced_settings_evasive_path_match(config_id=configuration.config_id)
    pulumi.export("advancedSettingsEvasivePathMatchOutput", evasive_path_match.output_text)
    pulumi.export("advancedSettingsEvasivePathMatchJson", evasive_path_match.json)
    policy_override = akamai.get_app_sec_advanced_settings_evasive_path_match(config_id=configuration.config_id,
        security_policy_id=var["security_policy_id"])
    pulumi.export("advancedSettingsPolicyEvasivePathMatchOutput", policy_override.output_text)
    pulumi.export("advancedSettingsPolicyEvasivePathMatchJson", policy_override.json)
    ```


    :param int config_id: The configuration ID.
    :param str security_policy_id: The ID of the security policy to use.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    __args__['securityPolicyId'] = security_policy_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecAdvancedSettingsEvasivePathMatch:getAppSecAdvancedSettingsEvasivePathMatch', __args__, opts=opts, typ=GetAppSecAdvancedSettingsEvasivePathMatchResult).value

    return AwaitableGetAppSecAdvancedSettingsEvasivePathMatchResult(
        config_id=__ret__.config_id,
        id=__ret__.id,
        json=__ret__.json,
        output_text=__ret__.output_text,
        security_policy_id=__ret__.security_policy_id)


@_utilities.lift_output_func(get_app_sec_advanced_settings_evasive_path_match)
def get_app_sec_advanced_settings_evasive_path_match_output(config_id: Optional[pulumi.Input[int]] = None,
                                                            security_policy_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppSecAdvancedSettingsEvasivePathMatchResult]:
    """
    **Scopes**: Security configuration; security policy

    Use the `AppSecAdvancedSettingsEvasivePathMatch` data source to retrieve information about the evasive path match for a configuration. This operation applies at the configuration level, and therefore applies to all policies within a configuration. You may retrieve these settings for a particular policy by specifying the policy using the `security_policy_id` parameter. For more information, see [Get evasive path match setting](https://techdocs.akamai.com/application-security/reference/get-evasive-path-match).

    **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/advanced-settings/evasive-path-match](https://techdocs.akamai.com/application-security/reference/get-evasive-path-match)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
    evasive_path_match = akamai.get_app_sec_advanced_settings_evasive_path_match(config_id=configuration.config_id)
    pulumi.export("advancedSettingsEvasivePathMatchOutput", evasive_path_match.output_text)
    pulumi.export("advancedSettingsEvasivePathMatchJson", evasive_path_match.json)
    policy_override = akamai.get_app_sec_advanced_settings_evasive_path_match(config_id=configuration.config_id,
        security_policy_id=var["security_policy_id"])
    pulumi.export("advancedSettingsPolicyEvasivePathMatchOutput", policy_override.output_text)
    pulumi.export("advancedSettingsPolicyEvasivePathMatchJson", policy_override.json)
    ```


    :param int config_id: The configuration ID.
    :param str security_policy_id: The ID of the security policy to use.
    """
    ...
