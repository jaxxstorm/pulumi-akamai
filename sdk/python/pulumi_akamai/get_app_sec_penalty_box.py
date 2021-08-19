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
        """
        The action for the penalty box: `alert`, `deny`, or `none`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> int:
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Either `true` or `false`, indicating whether penalty box protection is enabled.
        """
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
        """
        A tabular display of the `action` and `enabled` information.
        """
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
    Use the `AppSecPenaltyBox` data source to retrieve the penalty box settings for a specified security policy.

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
    penalty_box = akamai.get_app_sec_penalty_box(config_id=configuration.config_id,
        security_policy_id=var["security_policy_id"])
    pulumi.export("penaltyBoxAction", penalty_box.action)
    pulumi.export("penaltyBoxEnabled", penalty_box.enabled)
    pulumi.export("penaltyBoxText", penalty_box.output_text)
    ```


    :param int config_id: The ID of the security configuration to use.
    :param str security_policy_id: The ID of the security policy to use.
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