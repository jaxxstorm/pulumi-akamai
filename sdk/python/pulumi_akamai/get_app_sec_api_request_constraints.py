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
    'GetAppSecApiRequestConstraintsResult',
    'AwaitableGetAppSecApiRequestConstraintsResult',
    'get_app_sec_api_request_constraints',
    'get_app_sec_api_request_constraints_output',
]

@pulumi.output_type
class GetAppSecApiRequestConstraintsResult:
    """
    A collection of values returned by getAppSecApiRequestConstraints.
    """
    def __init__(__self__, api_id=None, config_id=None, id=None, json=None, output_text=None, security_policy_id=None):
        if api_id and not isinstance(api_id, int):
            raise TypeError("Expected argument 'api_id' to be a int")
        pulumi.set(__self__, "api_id", api_id)
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
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[int]:
        return pulumi.get(self, "api_id")

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


class AwaitableGetAppSecApiRequestConstraintsResult(GetAppSecApiRequestConstraintsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecApiRequestConstraintsResult(
            api_id=self.api_id,
            config_id=self.config_id,
            id=self.id,
            json=self.json,
            output_text=self.output_text,
            security_policy_id=self.security_policy_id)


def get_app_sec_api_request_constraints(api_id: Optional[int] = None,
                                        config_id: Optional[int] = None,
                                        security_policy_id: Optional[str] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecApiRequestConstraintsResult:
    """
    **Scopes**: Security policy; API endpoint

    Returns information about API endpoint constraints and actions. The returned information is described in the [API Constraints members](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getapirequestconstraints) section of the Application Security API.

    **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/api-request-constraints](https://techdocs.akamai.com/application-security/reference/get-api-request-constraints)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    apis_request_constraints = akamai.get_app_sec_api_request_constraints(config_id=configuration.config_id,
        security_policy_id="gms1_134637")
    pulumi.export("apisConstraintsText", apis_request_constraints.output_text)
    pulumi.export("apisConstraintsJson", apis_request_constraints.json)
    api_request_constraints = akamai.get_app_sec_api_request_constraints(config_id=configuration.config_id,
        security_policy_id="gms1_134637",
        api_id=624913)
    pulumi.export("apiConstraintsText", api_request_constraints.output_text)
    pulumi.export("apiConstraintsJson", api_request_constraints.json)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `json`. JSON-formatted list of information about the APIs, their constraints, and their actions.
    - `output_text`. Tabular report of the APIs, their constraints, and their actions.


    :param int api_id: . Unique identifier of the API endpoint you want to return constraint information for.
    :param int config_id: . Unique identifier of the security configuration associated with the API constraints.
    :param str security_policy_id: . Unique identifier of the security policy associated with the API constraints.
    """
    __args__ = dict()
    __args__['apiId'] = api_id
    __args__['configId'] = config_id
    __args__['securityPolicyId'] = security_policy_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecApiRequestConstraints:getAppSecApiRequestConstraints', __args__, opts=opts, typ=GetAppSecApiRequestConstraintsResult).value

    return AwaitableGetAppSecApiRequestConstraintsResult(
        api_id=__ret__.api_id,
        config_id=__ret__.config_id,
        id=__ret__.id,
        json=__ret__.json,
        output_text=__ret__.output_text,
        security_policy_id=__ret__.security_policy_id)


@_utilities.lift_output_func(get_app_sec_api_request_constraints)
def get_app_sec_api_request_constraints_output(api_id: Optional[pulumi.Input[Optional[int]]] = None,
                                               config_id: Optional[pulumi.Input[int]] = None,
                                               security_policy_id: Optional[pulumi.Input[str]] = None,
                                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppSecApiRequestConstraintsResult]:
    """
    **Scopes**: Security policy; API endpoint

    Returns information about API endpoint constraints and actions. The returned information is described in the [API Constraints members](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getapirequestconstraints) section of the Application Security API.

    **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/api-request-constraints](https://techdocs.akamai.com/application-security/reference/get-api-request-constraints)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    apis_request_constraints = akamai.get_app_sec_api_request_constraints(config_id=configuration.config_id,
        security_policy_id="gms1_134637")
    pulumi.export("apisConstraintsText", apis_request_constraints.output_text)
    pulumi.export("apisConstraintsJson", apis_request_constraints.json)
    api_request_constraints = akamai.get_app_sec_api_request_constraints(config_id=configuration.config_id,
        security_policy_id="gms1_134637",
        api_id=624913)
    pulumi.export("apiConstraintsText", api_request_constraints.output_text)
    pulumi.export("apiConstraintsJson", api_request_constraints.json)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `json`. JSON-formatted list of information about the APIs, their constraints, and their actions.
    - `output_text`. Tabular report of the APIs, their constraints, and their actions.


    :param int api_id: . Unique identifier of the API endpoint you want to return constraint information for.
    :param int config_id: . Unique identifier of the security configuration associated with the API constraints.
    :param str security_policy_id: . Unique identifier of the security policy associated with the API constraints.
    """
    ...
