# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetAppSecPenaltyBoxResult',
    'AwaitableGetAppSecPenaltyBoxResult',
    'get_app_sec_penalty_box',
    'get_app_sec_penalty_box_output',
]

@pulumi.output_type
class GetAppSecPenaltyBoxResult:
    """
    A collection of values returned by getAppSecPenaltyBox.
    """
    def __init__(__self__, action=None, config_id=None, enabled=None, id=None, output_text=None, security_policy_id=None):
        if action and not isinstance(action, str):
            raise TypeError("Expected argument 'action' to be a str")
        pulumi.set(__self__, "action", action)
        if config_id and not isinstance(config_id, int):
            raise TypeError("Expected argument 'config_id' to be a int")
        pulumi.set(__self__, "config_id", config_id)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_text and not isinstance(output_text, str):
            raise TypeError("Expected argument 'output_text' to be a str")
        pulumi.set(__self__, "output_text", output_text)
        if security_policy_id and not isinstance(security_policy_id, str):
            raise TypeError("Expected argument 'security_policy_id' to be a str")
        pulumi.set(__self__, "security_policy_id", security_policy_id)

    @property
    @pulumi.getter
    def action(self) -> str:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> int:
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

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

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> str:
        return pulumi.get(self, "security_policy_id")


class AwaitableGetAppSecPenaltyBoxResult(GetAppSecPenaltyBoxResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecPenaltyBoxResult(
            action=self.action,
            config_id=self.config_id,
            enabled=self.enabled,
            id=self.id,
            output_text=self.output_text,
            security_policy_id=self.security_policy_id)


def get_app_sec_penalty_box(config_id: Optional[int] = None,
                            security_policy_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecPenaltyBoxResult:
    """
    **Scopes**: Security policy

    Returns penalty box settings for the specified security policy. When using automated attack groups, and when the penalty box is enabled, clients that trigger an attack group are placed in the “penalty box.” That means that, for the next 10 minutes, all requests from that client are ignored.

    **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/penalty-box](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getpenaltybox)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    penalty_box = akamai.get_app_sec_penalty_box(config_id=configuration.config_id,
        security_policy_id="gms1_134637")
    pulumi.export("penaltyBoxAction", penalty_box.action)
    pulumi.export("penaltyBoxEnabled", penalty_box.enabled)
    pulumi.export("penaltyBoxText", penalty_box.output_text)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `action`. Action taken any time the penalty box is triggered. Valid values are:
      - **alert**. Record the event.
      - **deny**. The request is blocked.
      - **deny_custom_{custom_deny_id}**. The action defined by the custom deny is taken.
      - **none**. Take no action.
    - `enabled`. If **true**, penalty box protection is enabled. If **false**, penalty box protection is disabled.
    - `output_text`. Tabular report of penalty box protection settings.


    :param int config_id: . Unique identifier of the security configuration associated with the penalty box settings.
    :param str security_policy_id: . Unique identifier of the security policy associated with the penalty box settings.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    __args__['securityPolicyId'] = security_policy_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecPenaltyBox:getAppSecPenaltyBox', __args__, opts=opts, typ=GetAppSecPenaltyBoxResult).value

    return AwaitableGetAppSecPenaltyBoxResult(
        action=__ret__.action,
        config_id=__ret__.config_id,
        enabled=__ret__.enabled,
        id=__ret__.id,
        output_text=__ret__.output_text,
        security_policy_id=__ret__.security_policy_id)


@_utilities.lift_output_func(get_app_sec_penalty_box)
def get_app_sec_penalty_box_output(config_id: Optional[pulumi.Input[int]] = None,
                                   security_policy_id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppSecPenaltyBoxResult]:
    """
    **Scopes**: Security policy

    Returns penalty box settings for the specified security policy. When using automated attack groups, and when the penalty box is enabled, clients that trigger an attack group are placed in the “penalty box.” That means that, for the next 10 minutes, all requests from that client are ignored.

    **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/penalty-box](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getpenaltybox)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    penalty_box = akamai.get_app_sec_penalty_box(config_id=configuration.config_id,
        security_policy_id="gms1_134637")
    pulumi.export("penaltyBoxAction", penalty_box.action)
    pulumi.export("penaltyBoxEnabled", penalty_box.enabled)
    pulumi.export("penaltyBoxText", penalty_box.output_text)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `action`. Action taken any time the penalty box is triggered. Valid values are:
      - **alert**. Record the event.
      - **deny**. The request is blocked.
      - **deny_custom_{custom_deny_id}**. The action defined by the custom deny is taken.
      - **none**. Take no action.
    - `enabled`. If **true**, penalty box protection is enabled. If **false**, penalty box protection is disabled.
    - `output_text`. Tabular report of penalty box protection settings.


    :param int config_id: . Unique identifier of the security configuration associated with the penalty box settings.
    :param str security_policy_id: . Unique identifier of the security policy associated with the penalty box settings.
    """
    ...
