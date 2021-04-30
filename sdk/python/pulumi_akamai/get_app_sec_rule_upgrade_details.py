# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetAppSecRuleUpgradeDetailsResult',
    'AwaitableGetAppSecRuleUpgradeDetailsResult',
    'get_app_sec_rule_upgrade_details',
]

@pulumi.output_type
class GetAppSecRuleUpgradeDetailsResult:
    """
    A collection of values returned by getAppSecRuleUpgradeDetails.
    """
    def __init__(__self__, config_id=None, id=None, json=None, output_text=None, security_policy_id=None, version=None):
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
    @pulumi.getter
    def json(self) -> str:
        """
        A JSON-formatted list of the changes (additions and deletions) to the rules for the specified security policy.
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter(name="outputText")
    def output_text(self) -> str:
        """
        A tabular display showing changes (additions and deletions) to the rules for the specified security policy.
        """
        return pulumi.get(self, "output_text")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> str:
        return pulumi.get(self, "security_policy_id")

    @property
    @pulumi.getter
    def version(self) -> int:
        return pulumi.get(self, "version")


class AwaitableGetAppSecRuleUpgradeDetailsResult(GetAppSecRuleUpgradeDetailsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecRuleUpgradeDetailsResult(
            config_id=self.config_id,
            id=self.id,
            json=self.json,
            output_text=self.output_text,
            security_policy_id=self.security_policy_id,
            version=self.version)


def get_app_sec_rule_upgrade_details(config_id: Optional[int] = None,
                                     security_policy_id: Optional[str] = None,
                                     version: Optional[int] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecRuleUpgradeDetailsResult:
    """
    Use the `getAppSecRuleUpgradeDetails` data source to retrieve information on changes to the KRS rule sets.

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
    upgrade_details = akamai.get_app_sec_rule_upgrade_details(config_id=configuration.config_id,
        version=configuration.latest_version,
        security_policy_id=var["security_policy_id"])
    pulumi.export("upgradeDetailsText", upgrade_details.output_text)
    pulumi.export("upgradeDetailsJson", upgrade_details.json)
    ```


    :param int config_id: The ID of the security configuration to use.
    :param str security_policy_id: The ID of the security policy to use.
    :param int version: The version number of the security configuration to use.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    __args__['securityPolicyId'] = security_policy_id
    __args__['version'] = version
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecRuleUpgradeDetails:getAppSecRuleUpgradeDetails', __args__, opts=opts, typ=GetAppSecRuleUpgradeDetailsResult).value

    return AwaitableGetAppSecRuleUpgradeDetailsResult(
        config_id=__ret__.config_id,
        id=__ret__.id,
        json=__ret__.json,
        output_text=__ret__.output_text,
        security_policy_id=__ret__.security_policy_id,
        version=__ret__.version)
