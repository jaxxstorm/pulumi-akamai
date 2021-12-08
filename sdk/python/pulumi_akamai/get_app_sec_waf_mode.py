# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetAppSecWafModeResult',
    'AwaitableGetAppSecWafModeResult',
    'get_app_sec_waf_mode',
    'get_app_sec_waf_mode_output',
]

@pulumi.output_type
class GetAppSecWafModeResult:
    """
    A collection of values returned by getAppSecWafMode.
    """
    def __init__(__self__, config_id=None, current_ruleset=None, eval_expiration_date=None, eval_ruleset=None, eval_status=None, id=None, json=None, mode=None, output_text=None, security_policy_id=None):
        if config_id and not isinstance(config_id, int):
            raise TypeError("Expected argument 'config_id' to be a int")
        pulumi.set(__self__, "config_id", config_id)
        if current_ruleset and not isinstance(current_ruleset, str):
            raise TypeError("Expected argument 'current_ruleset' to be a str")
        pulumi.set(__self__, "current_ruleset", current_ruleset)
        if eval_expiration_date and not isinstance(eval_expiration_date, str):
            raise TypeError("Expected argument 'eval_expiration_date' to be a str")
        pulumi.set(__self__, "eval_expiration_date", eval_expiration_date)
        if eval_ruleset and not isinstance(eval_ruleset, str):
            raise TypeError("Expected argument 'eval_ruleset' to be a str")
        pulumi.set(__self__, "eval_ruleset", eval_ruleset)
        if eval_status and not isinstance(eval_status, str):
            raise TypeError("Expected argument 'eval_status' to be a str")
        pulumi.set(__self__, "eval_status", eval_status)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if json and not isinstance(json, str):
            raise TypeError("Expected argument 'json' to be a str")
        pulumi.set(__self__, "json", json)
        if mode and not isinstance(mode, str):
            raise TypeError("Expected argument 'mode' to be a str")
        pulumi.set(__self__, "mode", mode)
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
    @pulumi.getter(name="currentRuleset")
    def current_ruleset(self) -> str:
        return pulumi.get(self, "current_ruleset")

    @property
    @pulumi.getter(name="evalExpirationDate")
    def eval_expiration_date(self) -> str:
        return pulumi.get(self, "eval_expiration_date")

    @property
    @pulumi.getter(name="evalRuleset")
    def eval_ruleset(self) -> str:
        return pulumi.get(self, "eval_ruleset")

    @property
    @pulumi.getter(name="evalStatus")
    def eval_status(self) -> str:
        return pulumi.get(self, "eval_status")

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
    @pulumi.getter
    def mode(self) -> str:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter(name="outputText")
    def output_text(self) -> str:
        return pulumi.get(self, "output_text")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> str:
        return pulumi.get(self, "security_policy_id")


class AwaitableGetAppSecWafModeResult(GetAppSecWafModeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecWafModeResult(
            config_id=self.config_id,
            current_ruleset=self.current_ruleset,
            eval_expiration_date=self.eval_expiration_date,
            eval_ruleset=self.eval_ruleset,
            eval_status=self.eval_status,
            id=self.id,
            json=self.json,
            mode=self.mode,
            output_text=self.output_text,
            security_policy_id=self.security_policy_id)


def get_app_sec_waf_mode(config_id: Optional[int] = None,
                         security_policy_id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecWafModeResult:
    """
    **Scopes**: Security policy

    Returns information about how the Kona Rule Set rules associated with a security configuration and security policy are updated. The WAF (Web Application Firewall) mode determines whether Kona Rule Sets are automatically updated as part of automated attack groups (`mode = AAG`) or whether you must periodically check for new rules and then manually update those rules yourself (`mode = KRS`).

    **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/mode](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getmode)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    waf_mode = akamai.get_app_sec_waf_mode(config_id=configuration.config_id,
        security_policy_id="gms1_134637")
    pulumi.export("wafModeMode", waf_mode.mode)
    pulumi.export("wafModeCurrentRuleset", waf_mode.current_ruleset)
    pulumi.export("wafModeEvalStatus", waf_mode.eval_status)
    pulumi.export("wafModeEvalRuleset", waf_mode.eval_ruleset)
    pulumi.export("wafModeEvalExpirationDate", waf_mode.eval_expiration_date)
    pulumi.export("wafModeText", waf_mode.output_text)
    pulumi.export("wafModeJson", waf_mode.json)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `mode`. Security policy mode, either **KRS** (update manually) or **AAG** (update automatically), For organizations running the Adaptive Security Engine (ASE) beta, you'll get back **ASE_AUTO** for automatic updates or **ASE_MANUAL** for manual updates. Please contact your Akamai representative to learn more about ASE.
    - `current_ruleset`. Current ruleset version and the ISO 8601 date the version was introduced.
    - `eval_status`. Specifies whether evaluation mode is enabled or disabled.
    - `eval_ruleset`. Evaluation ruleset version and the ISO 8601 date the evaluation began.
    - `eval_expiration_date`. ISO 8601 timestamp indicating when evaluation mode expires. Valid only if `eval_status` is set to **enabled**.
    - `output_text`. Tabular report of the mode information.
    - `json`. JSON-formatted list of the mode information.


    :param int config_id: . Unique identifier of the security configuration associated with the Kona Rule Set rules.
    :param str security_policy_id: . Unique identifier of the security policy associated with the Kona Rule Set rules.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    __args__['securityPolicyId'] = security_policy_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecWafMode:getAppSecWafMode', __args__, opts=opts, typ=GetAppSecWafModeResult).value

    return AwaitableGetAppSecWafModeResult(
        config_id=__ret__.config_id,
        current_ruleset=__ret__.current_ruleset,
        eval_expiration_date=__ret__.eval_expiration_date,
        eval_ruleset=__ret__.eval_ruleset,
        eval_status=__ret__.eval_status,
        id=__ret__.id,
        json=__ret__.json,
        mode=__ret__.mode,
        output_text=__ret__.output_text,
        security_policy_id=__ret__.security_policy_id)


@_utilities.lift_output_func(get_app_sec_waf_mode)
def get_app_sec_waf_mode_output(config_id: Optional[pulumi.Input[int]] = None,
                                security_policy_id: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppSecWafModeResult]:
    """
    **Scopes**: Security policy

    Returns information about how the Kona Rule Set rules associated with a security configuration and security policy are updated. The WAF (Web Application Firewall) mode determines whether Kona Rule Sets are automatically updated as part of automated attack groups (`mode = AAG`) or whether you must periodically check for new rules and then manually update those rules yourself (`mode = KRS`).

    **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/mode](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getmode)

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name="Documentation")
    waf_mode = akamai.get_app_sec_waf_mode(config_id=configuration.config_id,
        security_policy_id="gms1_134637")
    pulumi.export("wafModeMode", waf_mode.mode)
    pulumi.export("wafModeCurrentRuleset", waf_mode.current_ruleset)
    pulumi.export("wafModeEvalStatus", waf_mode.eval_status)
    pulumi.export("wafModeEvalRuleset", waf_mode.eval_ruleset)
    pulumi.export("wafModeEvalExpirationDate", waf_mode.eval_expiration_date)
    pulumi.export("wafModeText", waf_mode.output_text)
    pulumi.export("wafModeJson", waf_mode.json)
    ```
    ## Output Options

    The following options can be used to determine the information returned, and how that returned information is formatted:

    - `mode`. Security policy mode, either **KRS** (update manually) or **AAG** (update automatically), For organizations running the Adaptive Security Engine (ASE) beta, you'll get back **ASE_AUTO** for automatic updates or **ASE_MANUAL** for manual updates. Please contact your Akamai representative to learn more about ASE.
    - `current_ruleset`. Current ruleset version and the ISO 8601 date the version was introduced.
    - `eval_status`. Specifies whether evaluation mode is enabled or disabled.
    - `eval_ruleset`. Evaluation ruleset version and the ISO 8601 date the evaluation began.
    - `eval_expiration_date`. ISO 8601 timestamp indicating when evaluation mode expires. Valid only if `eval_status` is set to **enabled**.
    - `output_text`. Tabular report of the mode information.
    - `json`. JSON-formatted list of the mode information.


    :param int config_id: . Unique identifier of the security configuration associated with the Kona Rule Set rules.
    :param str security_policy_id: . Unique identifier of the security policy associated with the Kona Rule Set rules.
    """
    ...
