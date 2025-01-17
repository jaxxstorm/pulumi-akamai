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
    'GetAppSecAttackGroupsResult',
    'AwaitableGetAppSecAttackGroupsResult',
    'get_app_sec_attack_groups',
    'get_app_sec_attack_groups_output',
]

@pulumi.output_type
class GetAppSecAttackGroupsResult:
    """
    A collection of values returned by getAppSecAttackGroups.
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


class AwaitableGetAppSecAttackGroupsResult(GetAppSecAttackGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecAttackGroupsResult(
            attack_group=self.attack_group,
            attack_group_action=self.attack_group_action,
            condition_exception=self.condition_exception,
            config_id=self.config_id,
            id=self.id,
            json=self.json,
            output_text=self.output_text,
            security_policy_id=self.security_policy_id)


def get_app_sec_attack_groups(attack_group: Optional[str] = None,
                              config_id: Optional[int] = None,
                              security_policy_id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecAttackGroupsResult:
    """
    **Scopes**: Security policy; attack group

    Returns the action and the condition-exception information for an attack group or set of attack groups. Attack groups are collections of Kona Rule Set rules used to streamline the management of website protections.

    **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/attack-groups](https://techdocs.akamai.com/application-security/reference/get-policy-attack-groups)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    attack_group = akamai.get_app_sec_attack_groups(config_id=configuration.config_id,
        security_policy_id="gms1_134637",
        attack_group="SQL")
    pulumi.export("attackGroupAction", attack_group.attack_group_action)
    pulumi.export("conditionException", attack_group.condition_exception)
    pulumi.export("json", attack_group.json)
    pulumi.export("outputText", attack_group.output_text)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `attack_group_action`. Action taken anytime the attack group is triggered. This information is returned only when a single attack group is retrieved. Valid values are:
      - **alert**. The event is recorded.
      - **deny**. The request is blocked.
      - **deny_custom_{custom_deny_id}**. The action defined by the custom deny is taken.
      - **none**. No action is taken.
    - `condition_exception`. Conditions and exceptions assigned to the attack group. This information is returned only when a single attack group is retrieved.
    - `json`. JSON-formatted list of the action and the condition-exception information for the attack group. This information is returned only when a single attack group is retrieved.
    - `output_text`. Tabular report showing the attack group's action as well as Boolean values indicating whether conditions and exceptions have been configured for the group.


    :param str attack_group: . Unique name of the attack group you want to return information for. If not included, information is returned for all your attack groups.
    :param int config_id: . Unique identifier of the security configuration associated with the attack group.
    :param str security_policy_id: . Unique identifier of the security policy associated with the attack group.
    """
    __args__ = dict()
    __args__['attackGroup'] = attack_group
    __args__['configId'] = config_id
    __args__['securityPolicyId'] = security_policy_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecAttackGroups:getAppSecAttackGroups', __args__, opts=opts, typ=GetAppSecAttackGroupsResult).value

    return AwaitableGetAppSecAttackGroupsResult(
        attack_group=__ret__.attack_group,
        attack_group_action=__ret__.attack_group_action,
        condition_exception=__ret__.condition_exception,
        config_id=__ret__.config_id,
        id=__ret__.id,
        json=__ret__.json,
        output_text=__ret__.output_text,
        security_policy_id=__ret__.security_policy_id)


@_utilities.lift_output_func(get_app_sec_attack_groups)
def get_app_sec_attack_groups_output(attack_group: Optional[pulumi.Input[Optional[str]]] = None,
                                     config_id: Optional[pulumi.Input[int]] = None,
                                     security_policy_id: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppSecAttackGroupsResult]:
    """
    **Scopes**: Security policy; attack group

    Returns the action and the condition-exception information for an attack group or set of attack groups. Attack groups are collections of Kona Rule Set rules used to streamline the management of website protections.

    **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/attack-groups](https://techdocs.akamai.com/application-security/reference/get-policy-attack-groups)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    attack_group = akamai.get_app_sec_attack_groups(config_id=configuration.config_id,
        security_policy_id="gms1_134637",
        attack_group="SQL")
    pulumi.export("attackGroupAction", attack_group.attack_group_action)
    pulumi.export("conditionException", attack_group.condition_exception)
    pulumi.export("json", attack_group.json)
    pulumi.export("outputText", attack_group.output_text)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `attack_group_action`. Action taken anytime the attack group is triggered. This information is returned only when a single attack group is retrieved. Valid values are:
      - **alert**. The event is recorded.
      - **deny**. The request is blocked.
      - **deny_custom_{custom_deny_id}**. The action defined by the custom deny is taken.
      - **none**. No action is taken.
    - `condition_exception`. Conditions and exceptions assigned to the attack group. This information is returned only when a single attack group is retrieved.
    - `json`. JSON-formatted list of the action and the condition-exception information for the attack group. This information is returned only when a single attack group is retrieved.
    - `output_text`. Tabular report showing the attack group's action as well as Boolean values indicating whether conditions and exceptions have been configured for the group.


    :param str attack_group: . Unique name of the attack group you want to return information for. If not included, information is returned for all your attack groups.
    :param int config_id: . Unique identifier of the security configuration associated with the attack group.
    :param str security_policy_id: . Unique identifier of the security policy associated with the attack group.
    """
    ...
