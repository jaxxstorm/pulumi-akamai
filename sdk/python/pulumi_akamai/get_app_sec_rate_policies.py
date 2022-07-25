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
    'GetAppSecRatePoliciesResult',
    'AwaitableGetAppSecRatePoliciesResult',
    'get_app_sec_rate_policies',
    'get_app_sec_rate_policies_output',
]

@pulumi.output_type
class GetAppSecRatePoliciesResult:
    """
    A collection of values returned by getAppSecRatePolicies.
    """
    def __init__(__self__, config_id=None, id=None, json=None, output_text=None, rate_policy_id=None):
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
        if rate_policy_id and not isinstance(rate_policy_id, int):
            raise TypeError("Expected argument 'rate_policy_id' to be a int")
        pulumi.set(__self__, "rate_policy_id", rate_policy_id)

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
    @pulumi.getter(name="ratePolicyId")
    def rate_policy_id(self) -> Optional[int]:
        return pulumi.get(self, "rate_policy_id")


class AwaitableGetAppSecRatePoliciesResult(GetAppSecRatePoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecRatePoliciesResult(
            config_id=self.config_id,
            id=self.id,
            json=self.json,
            output_text=self.output_text,
            rate_policy_id=self.rate_policy_id)


def get_app_sec_rate_policies(config_id: Optional[int] = None,
                              rate_policy_id: Optional[int] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecRatePoliciesResult:
    """
    **Scopes**: Security configuration; rate policy

    Returns information about your rate policies. Rate polices help you monitor and moderate the number and rate of all the requests you receive; in turn, this helps you prevent your website from being overwhelmed by a dramatic, and unexpected, surge in traffic.

    **Related API Endpoint:** [/appsec/v1/configs/{configId}/versions/{versionNumber}/rate-policies](https://techdocs.akamai.com/application-security/reference/get-rate-policies)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    rate_policies = akamai.get_app_sec_rate_policies(config_id=configuration.config_id)
    pulumi.export("ratePoliciesOutput", rate_policies.output_text)
    pulumi.export("ratePoliciesJson", rate_policies.json)
    rate_policy = akamai.get_app_sec_rate_policies(config_id=configuration.config_id,
        rate_policy_id=122149)
    pulumi.export("ratePolicyJson", rate_policy.json)
    pulumi.export("ratePolicyOutput", rate_policy.output_text)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `output_text`. Tabular report showing the ID and name of the rate policies.
    - `json`. JSON-formatted list of the rate policy information.


    :param int config_id: . Unique identifier of the security configuration associated with the rate policies.
    :param int rate_policy_id: . Unique identifier of the rate policy you want to return information for. If not included, information is returned for all your rate policies.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    __args__['ratePolicyId'] = rate_policy_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecRatePolicies:getAppSecRatePolicies', __args__, opts=opts, typ=GetAppSecRatePoliciesResult).value

    return AwaitableGetAppSecRatePoliciesResult(
        config_id=__ret__.config_id,
        id=__ret__.id,
        json=__ret__.json,
        output_text=__ret__.output_text,
        rate_policy_id=__ret__.rate_policy_id)


@_utilities.lift_output_func(get_app_sec_rate_policies)
def get_app_sec_rate_policies_output(config_id: Optional[pulumi.Input[int]] = None,
                                     rate_policy_id: Optional[pulumi.Input[Optional[int]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppSecRatePoliciesResult]:
    """
    **Scopes**: Security configuration; rate policy

    Returns information about your rate policies. Rate polices help you monitor and moderate the number and rate of all the requests you receive; in turn, this helps you prevent your website from being overwhelmed by a dramatic, and unexpected, surge in traffic.

    **Related API Endpoint:** [/appsec/v1/configs/{configId}/versions/{versionNumber}/rate-policies](https://techdocs.akamai.com/application-security/reference/get-rate-policies)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    rate_policies = akamai.get_app_sec_rate_policies(config_id=configuration.config_id)
    pulumi.export("ratePoliciesOutput", rate_policies.output_text)
    pulumi.export("ratePoliciesJson", rate_policies.json)
    rate_policy = akamai.get_app_sec_rate_policies(config_id=configuration.config_id,
        rate_policy_id=122149)
    pulumi.export("ratePolicyJson", rate_policy.json)
    pulumi.export("ratePolicyOutput", rate_policy.output_text)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `output_text`. Tabular report showing the ID and name of the rate policies.
    - `json`. JSON-formatted list of the rate policy information.


    :param int config_id: . Unique identifier of the security configuration associated with the rate policies.
    :param int rate_policy_id: . Unique identifier of the rate policy you want to return information for. If not included, information is returned for all your rate policies.
    """
    ...
