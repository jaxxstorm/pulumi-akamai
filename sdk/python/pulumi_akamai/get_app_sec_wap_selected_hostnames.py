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
    'GetAppSecWapSelectedHostnamesResult',
    'AwaitableGetAppSecWapSelectedHostnamesResult',
    'get_app_sec_wap_selected_hostnames',
    'get_app_sec_wap_selected_hostnames_output',
]

@pulumi.output_type
class GetAppSecWapSelectedHostnamesResult:
    """
    A collection of values returned by getAppSecWapSelectedHostnames.
    """
    def __init__(__self__, config_id=None, evaluated_hosts=None, id=None, json=None, match_targets=None, output_text=None, protected_hosts=None, security_policy_id=None, selected_hosts=None):
        if config_id and not isinstance(config_id, int):
            raise TypeError("Expected argument 'config_id' to be a int")
        pulumi.set(__self__, "config_id", config_id)
        if evaluated_hosts and not isinstance(evaluated_hosts, list):
            raise TypeError("Expected argument 'evaluated_hosts' to be a list")
        pulumi.set(__self__, "evaluated_hosts", evaluated_hosts)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if json and not isinstance(json, str):
            raise TypeError("Expected argument 'json' to be a str")
        pulumi.set(__self__, "json", json)
        if match_targets and not isinstance(match_targets, str):
            raise TypeError("Expected argument 'match_targets' to be a str")
        pulumi.set(__self__, "match_targets", match_targets)
        if output_text and not isinstance(output_text, str):
            raise TypeError("Expected argument 'output_text' to be a str")
        pulumi.set(__self__, "output_text", output_text)
        if protected_hosts and not isinstance(protected_hosts, list):
            raise TypeError("Expected argument 'protected_hosts' to be a list")
        pulumi.set(__self__, "protected_hosts", protected_hosts)
        if security_policy_id and not isinstance(security_policy_id, str):
            raise TypeError("Expected argument 'security_policy_id' to be a str")
        pulumi.set(__self__, "security_policy_id", security_policy_id)
        if selected_hosts and not isinstance(selected_hosts, list):
            raise TypeError("Expected argument 'selected_hosts' to be a list")
        pulumi.set(__self__, "selected_hosts", selected_hosts)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> int:
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="evaluatedHosts")
    def evaluated_hosts(self) -> Sequence[str]:
        return pulumi.get(self, "evaluated_hosts")

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
    @pulumi.getter(name="matchTargets")
    def match_targets(self) -> str:
        return pulumi.get(self, "match_targets")

    @property
    @pulumi.getter(name="outputText")
    def output_text(self) -> str:
        return pulumi.get(self, "output_text")

    @property
    @pulumi.getter(name="protectedHosts")
    def protected_hosts(self) -> Sequence[str]:
        return pulumi.get(self, "protected_hosts")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> str:
        return pulumi.get(self, "security_policy_id")

    @property
    @pulumi.getter(name="selectedHosts")
    def selected_hosts(self) -> Sequence[str]:
        return pulumi.get(self, "selected_hosts")


class AwaitableGetAppSecWapSelectedHostnamesResult(GetAppSecWapSelectedHostnamesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecWapSelectedHostnamesResult(
            config_id=self.config_id,
            evaluated_hosts=self.evaluated_hosts,
            id=self.id,
            json=self.json,
            match_targets=self.match_targets,
            output_text=self.output_text,
            protected_hosts=self.protected_hosts,
            security_policy_id=self.security_policy_id,
            selected_hosts=self.selected_hosts)


def get_app_sec_wap_selected_hostnames(config_id: Optional[int] = None,
                                       security_policy_id: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecWapSelectedHostnamesResult:
    """
    **Scopes**: Security policy

    Returns hostnames currently protected or being evaluated by a configuration and security policy.
    This resource is available only to organizations running Web Application Protector (WAP).
    Note that the WAP selected hostnames feature is currently in beta.
    Please contact your Akamai representative for more information.

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    wap_selected_hostnames = akamai.get_app_sec_wap_selected_hostnames(config_id=configuration.config_id,
        security_policy_id="gms1_134637")
    pulumi.export("protectedHostnames", wap_selected_hostnames.protected_hosts)
    pulumi.export("evaluatedHostnames", wap_selected_hostnames.evaluated_hosts)
    ```
    ## Output Options

    The following options can be used to determine the information returned and how that returned information is formatted:

    - `protected_hostnames`. List of hostnames currently protected under the security configuration and security policy.
    - `evaluated_hostnames`. List of hostnames currently being evaluated under the security configuration and security policy.
    - `hostnames_json`. JSON-formatted report of the protected and evaluated hostnames.
    - `output_text`. Tabular reports of the protected and evaluated hostnames.


    :param int config_id: . Unique identifier of the security configuration associated with the hostnames.
    :param str security_policy_id: . Unique identifier of the security policy associated with the hostnames.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    __args__['securityPolicyId'] = security_policy_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecWapSelectedHostnames:getAppSecWapSelectedHostnames', __args__, opts=opts, typ=GetAppSecWapSelectedHostnamesResult).value

    return AwaitableGetAppSecWapSelectedHostnamesResult(
        config_id=__ret__.config_id,
        evaluated_hosts=__ret__.evaluated_hosts,
        id=__ret__.id,
        json=__ret__.json,
        match_targets=__ret__.match_targets,
        output_text=__ret__.output_text,
        protected_hosts=__ret__.protected_hosts,
        security_policy_id=__ret__.security_policy_id,
        selected_hosts=__ret__.selected_hosts)


@_utilities.lift_output_func(get_app_sec_wap_selected_hostnames)
def get_app_sec_wap_selected_hostnames_output(config_id: Optional[pulumi.Input[int]] = None,
                                              security_policy_id: Optional[pulumi.Input[str]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppSecWapSelectedHostnamesResult]:
    """
    **Scopes**: Security policy

    Returns hostnames currently protected or being evaluated by a configuration and security policy.
    This resource is available only to organizations running Web Application Protector (WAP).
    Note that the WAP selected hostnames feature is currently in beta.
    Please contact your Akamai representative for more information.

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    wap_selected_hostnames = akamai.get_app_sec_wap_selected_hostnames(config_id=configuration.config_id,
        security_policy_id="gms1_134637")
    pulumi.export("protectedHostnames", wap_selected_hostnames.protected_hosts)
    pulumi.export("evaluatedHostnames", wap_selected_hostnames.evaluated_hosts)
    ```
    ## Output Options

    The following options can be used to determine the information returned and how that returned information is formatted:

    - `protected_hostnames`. List of hostnames currently protected under the security configuration and security policy.
    - `evaluated_hostnames`. List of hostnames currently being evaluated under the security configuration and security policy.
    - `hostnames_json`. JSON-formatted report of the protected and evaluated hostnames.
    - `output_text`. Tabular reports of the protected and evaluated hostnames.


    :param int config_id: . Unique identifier of the security configuration associated with the hostnames.
    :param str security_policy_id: . Unique identifier of the security policy associated with the hostnames.
    """
    ...
