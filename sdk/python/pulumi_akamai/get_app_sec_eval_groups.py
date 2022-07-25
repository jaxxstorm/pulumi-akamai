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
    'GetAppSecEvalGroupsResult',
    'AwaitableGetAppSecEvalGroupsResult',
    'get_app_sec_eval_groups',
    'get_app_sec_eval_groups_output',
]

@pulumi.output_type
class GetAppSecEvalGroupsResult:
    """
    A collection of values returned by getAppSecEvalGroups.
    """
    def __init__(__self__, attack_group=None, attack_group_action=None, condition_exception=None, config_id=None, id=None, json=None, output_text=None, security_policy_id=None):
        if attack_group and not isinstance(attack_group, str):
            raise TypeError("Expected argument 'attack_group' to be a str")
        pulumi.set(__self__, "attack_group", attack_group)
        if attack_group_action and not isinstance(attack_group_action, str):
            raise TypeError("Expected argument 'attack_group_action' to be a str")
        pulumi.set(__self__, "attack_group_action", attack_group_action)
        if condition_exception and not isinstance(condition_exception, str):
            raise TypeError("Expected argument 'condition_exception' to be a str")
        pulumi.set(__self__, "condition_exception", condition_exception)
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
    @pulumi.getter(name="attackGroup")
    def attack_group(self) -> Optional[str]:
        return pulumi.get(self, "attack_group")

    @property
    @pulumi.getter(name="attackGroupAction")
    def attack_group_action(self) -> str:
        return pulumi.get(self, "attack_group_action")

    @property
    @pulumi.getter(name="conditionException")
    def condition_exception(self) -> str:
        return pulumi.get(self, "condition_exception")

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

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> str:
        return pulumi.get(self, "security_policy_id")


class AwaitableGetAppSecEvalGroupsResult(GetAppSecEvalGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecEvalGroupsResult(
            attack_group=self.attack_group,
            attack_group_action=self.attack_group_action,
            condition_exception=self.condition_exception,
            config_id=self.config_id,
            id=self.id,
            json=self.json,
            output_text=self.output_text,
            security_policy_id=self.security_policy_id)


def get_app_sec_eval_groups(attack_group: Optional[str] = None,
                            config_id: Optional[int] = None,
                            security_policy_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecEvalGroupsResult:
    """
    Use this data source to access information about an existing resource.

    :param str attack_group: . Unique identifier of the evaluation attack group you want to return information for. If not included, information is returned for all your evaluation attack groups.
    :param int config_id: . Unique identifier of the security configuration associated with the evaluation attack group.
    :param str security_policy_id: . Unique identifier of the security policy associated with the evaluation attack group.
    """
    __args__ = dict()
    __args__['attackGroup'] = attack_group
    __args__['configId'] = config_id
    __args__['securityPolicyId'] = security_policy_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecEvalGroups:getAppSecEvalGroups', __args__, opts=opts, typ=GetAppSecEvalGroupsResult).value

    return AwaitableGetAppSecEvalGroupsResult(
        attack_group=__ret__.attack_group,
        attack_group_action=__ret__.attack_group_action,
        condition_exception=__ret__.condition_exception,
        config_id=__ret__.config_id,
        id=__ret__.id,
        json=__ret__.json,
        output_text=__ret__.output_text,
        security_policy_id=__ret__.security_policy_id)


@_utilities.lift_output_func(get_app_sec_eval_groups)
def get_app_sec_eval_groups_output(attack_group: Optional[pulumi.Input[Optional[str]]] = None,
                                   config_id: Optional[pulumi.Input[int]] = None,
                                   security_policy_id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppSecEvalGroupsResult]:
    """
    Use this data source to access information about an existing resource.

    :param str attack_group: . Unique identifier of the evaluation attack group you want to return information for. If not included, information is returned for all your evaluation attack groups.
    :param int config_id: . Unique identifier of the security configuration associated with the evaluation attack group.
    :param str security_policy_id: . Unique identifier of the security policy associated with the evaluation attack group.
    """
    ...
