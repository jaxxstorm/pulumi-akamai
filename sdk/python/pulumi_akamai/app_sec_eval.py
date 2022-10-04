# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecEvalArgs', 'AppSecEval']

@pulumi.input_type
class AppSecEvalArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 eval_operation: pulumi.Input[str],
                 security_policy_id: pulumi.Input[str],
                 eval_mode: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AppSecEval resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
        :param pulumi.Input[str] eval_operation: . Evaluation mode operation. Allowed values are:
               - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
               - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
               - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
               - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
               - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the evaluation process.
        :param pulumi.Input[str] eval_mode: . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "eval_operation", eval_operation)
        pulumi.set(__self__, "security_policy_id", security_policy_id)
        if eval_mode is not None:
            pulumi.set(__self__, "eval_mode", eval_mode)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="evalOperation")
    def eval_operation(self) -> pulumi.Input[str]:
        """
        . Evaluation mode operation. Allowed values are:
        - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
        - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
        - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
        - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
        - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
        """
        return pulumi.get(self, "eval_operation")

    @eval_operation.setter
    def eval_operation(self, value: pulumi.Input[str]):
        pulumi.set(self, "eval_operation", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Input[str]:
        """
        . Unique identifier of the security policy associated with the evaluation process.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_policy_id", value)

    @property
    @pulumi.getter(name="evalMode")
    def eval_mode(self) -> Optional[pulumi.Input[str]]:
        """
        . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
        """
        return pulumi.get(self, "eval_mode")

    @eval_mode.setter
    def eval_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eval_mode", value)


@pulumi.input_type
class _AppSecEvalState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 current_ruleset: Optional[pulumi.Input[str]] = None,
                 eval_mode: Optional[pulumi.Input[str]] = None,
                 eval_operation: Optional[pulumi.Input[str]] = None,
                 eval_status: Optional[pulumi.Input[str]] = None,
                 evaluating_ruleset: Optional[pulumi.Input[str]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AppSecEval resources.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
        :param pulumi.Input[str] current_ruleset: Versioning information for the Kona Rule Set currently in use in production
        :param pulumi.Input[str] eval_mode: . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
        :param pulumi.Input[str] eval_operation: . Evaluation mode operation. Allowed values are:
               - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
               - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
               - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
               - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
               - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
        :param pulumi.Input[str] eval_status: Whether an evaluation is currently in progress
        :param pulumi.Input[str] evaluating_ruleset: Versioning information for the Kona Rule Set being evaluated
        :param pulumi.Input[str] expiration_date: Date when the evaluation period ends
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the evaluation process.
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if current_ruleset is not None:
            pulumi.set(__self__, "current_ruleset", current_ruleset)
        if eval_mode is not None:
            pulumi.set(__self__, "eval_mode", eval_mode)
        if eval_operation is not None:
            pulumi.set(__self__, "eval_operation", eval_operation)
        if eval_status is not None:
            pulumi.set(__self__, "eval_status", eval_status)
        if evaluating_ruleset is not None:
            pulumi.set(__self__, "evaluating_ruleset", evaluating_ruleset)
        if expiration_date is not None:
            pulumi.set(__self__, "expiration_date", expiration_date)
        if security_policy_id is not None:
            pulumi.set(__self__, "security_policy_id", security_policy_id)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="currentRuleset")
    def current_ruleset(self) -> Optional[pulumi.Input[str]]:
        """
        Versioning information for the Kona Rule Set currently in use in production
        """
        return pulumi.get(self, "current_ruleset")

    @current_ruleset.setter
    def current_ruleset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "current_ruleset", value)

    @property
    @pulumi.getter(name="evalMode")
    def eval_mode(self) -> Optional[pulumi.Input[str]]:
        """
        . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
        """
        return pulumi.get(self, "eval_mode")

    @eval_mode.setter
    def eval_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eval_mode", value)

    @property
    @pulumi.getter(name="evalOperation")
    def eval_operation(self) -> Optional[pulumi.Input[str]]:
        """
        . Evaluation mode operation. Allowed values are:
        - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
        - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
        - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
        - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
        - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
        """
        return pulumi.get(self, "eval_operation")

    @eval_operation.setter
    def eval_operation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eval_operation", value)

    @property
    @pulumi.getter(name="evalStatus")
    def eval_status(self) -> Optional[pulumi.Input[str]]:
        """
        Whether an evaluation is currently in progress
        """
        return pulumi.get(self, "eval_status")

    @eval_status.setter
    def eval_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eval_status", value)

    @property
    @pulumi.getter(name="evaluatingRuleset")
    def evaluating_ruleset(self) -> Optional[pulumi.Input[str]]:
        """
        Versioning information for the Kona Rule Set being evaluated
        """
        return pulumi.get(self, "evaluating_ruleset")

    @evaluating_ruleset.setter
    def evaluating_ruleset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "evaluating_ruleset", value)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[pulumi.Input[str]]:
        """
        Date when the evaluation period ends
        """
        return pulumi.get(self, "expiration_date")

    @expiration_date.setter
    def expiration_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_date", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        . Unique identifier of the security policy associated with the evaluation process.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_policy_id", value)


class AppSecEval(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 eval_mode: Optional[pulumi.Input[str]] = None,
                 eval_operation: Optional[pulumi.Input[str]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        **Scopes**: Security policy

        Issues an evaluation mode command (`Start`, `Stop`, `Restart`, `Update`, or `Complete`) to a security configuration.
        Evaluation mode is used for testing and fine-tuning your Kona Rule Set rules and configuration settings.
        In evaluation mode rules are triggered by events, but the only thing those rules do is record the actions they *would* have taken had the event occurred on the production network.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/eval](https://techdocs.akamai.com/application-security/reference/post-policy-eval)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        eval_operation = akamai.AppSecEval("evalOperation",
            config_id=configuration.config_id,
            security_policy_id="gms1_134637",
            eval_operation="START")
        pulumi.export("evalModeEvaluatingRuleset", eval_operation.evaluating_ruleset)
        pulumi.export("evalModeExpirationDate", eval_operation.expiration_date)
        pulumi.export("evalModeCurrentRuleset", eval_operation.current_ruleset)
        pulumi.export("evalModeStatus", eval_operation.eval_status)
        ```
        ## Output Options

        The following options can be used to determine the information returned, and how that returned information is formatted:

        - `evaluating_ruleset`. Versioning information for the Kona Rule Set being evaluated.
        - `expiration_date`. Date when the evaluation period ends.
        - `current_ruleset`. Versioning information for the Kona Rule Set currently in use on the production network.
        - `eval_status`. If **true**, an evaluation is currently in progress; if **false**, evaluation is either paused or is not running.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
        :param pulumi.Input[str] eval_mode: . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
        :param pulumi.Input[str] eval_operation: . Evaluation mode operation. Allowed values are:
               - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
               - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
               - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
               - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
               - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the evaluation process.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecEvalArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        **Scopes**: Security policy

        Issues an evaluation mode command (`Start`, `Stop`, `Restart`, `Update`, or `Complete`) to a security configuration.
        Evaluation mode is used for testing and fine-tuning your Kona Rule Set rules and configuration settings.
        In evaluation mode rules are triggered by events, but the only thing those rules do is record the actions they *would* have taken had the event occurred on the production network.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/eval](https://techdocs.akamai.com/application-security/reference/post-policy-eval)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        eval_operation = akamai.AppSecEval("evalOperation",
            config_id=configuration.config_id,
            security_policy_id="gms1_134637",
            eval_operation="START")
        pulumi.export("evalModeEvaluatingRuleset", eval_operation.evaluating_ruleset)
        pulumi.export("evalModeExpirationDate", eval_operation.expiration_date)
        pulumi.export("evalModeCurrentRuleset", eval_operation.current_ruleset)
        pulumi.export("evalModeStatus", eval_operation.eval_status)
        ```
        ## Output Options

        The following options can be used to determine the information returned, and how that returned information is formatted:

        - `evaluating_ruleset`. Versioning information for the Kona Rule Set being evaluated.
        - `expiration_date`. Date when the evaluation period ends.
        - `current_ruleset`. Versioning information for the Kona Rule Set currently in use on the production network.
        - `eval_status`. If **true**, an evaluation is currently in progress; if **false**, evaluation is either paused or is not running.

        :param str resource_name: The name of the resource.
        :param AppSecEvalArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecEvalArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 eval_mode: Optional[pulumi.Input[str]] = None,
                 eval_operation: Optional[pulumi.Input[str]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppSecEvalArgs.__new__(AppSecEvalArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            __props__.__dict__["eval_mode"] = eval_mode
            if eval_operation is None and not opts.urn:
                raise TypeError("Missing required property 'eval_operation'")
            __props__.__dict__["eval_operation"] = eval_operation
            if security_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_policy_id'")
            __props__.__dict__["security_policy_id"] = security_policy_id
            __props__.__dict__["current_ruleset"] = None
            __props__.__dict__["eval_status"] = None
            __props__.__dict__["evaluating_ruleset"] = None
            __props__.__dict__["expiration_date"] = None
        super(AppSecEval, __self__).__init__(
            'akamai:index/appSecEval:AppSecEval',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            current_ruleset: Optional[pulumi.Input[str]] = None,
            eval_mode: Optional[pulumi.Input[str]] = None,
            eval_operation: Optional[pulumi.Input[str]] = None,
            eval_status: Optional[pulumi.Input[str]] = None,
            evaluating_ruleset: Optional[pulumi.Input[str]] = None,
            expiration_date: Optional[pulumi.Input[str]] = None,
            security_policy_id: Optional[pulumi.Input[str]] = None) -> 'AppSecEval':
        """
        Get an existing AppSecEval resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
        :param pulumi.Input[str] current_ruleset: Versioning information for the Kona Rule Set currently in use in production
        :param pulumi.Input[str] eval_mode: . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
        :param pulumi.Input[str] eval_operation: . Evaluation mode operation. Allowed values are:
               - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
               - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
               - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
               - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
               - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
        :param pulumi.Input[str] eval_status: Whether an evaluation is currently in progress
        :param pulumi.Input[str] evaluating_ruleset: Versioning information for the Kona Rule Set being evaluated
        :param pulumi.Input[str] expiration_date: Date when the evaluation period ends
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the evaluation process.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecEvalState.__new__(_AppSecEvalState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["current_ruleset"] = current_ruleset
        __props__.__dict__["eval_mode"] = eval_mode
        __props__.__dict__["eval_operation"] = eval_operation
        __props__.__dict__["eval_status"] = eval_status
        __props__.__dict__["evaluating_ruleset"] = evaluating_ruleset
        __props__.__dict__["expiration_date"] = expiration_date
        __props__.__dict__["security_policy_id"] = security_policy_id
        return AppSecEval(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="currentRuleset")
    def current_ruleset(self) -> pulumi.Output[str]:
        """
        Versioning information for the Kona Rule Set currently in use in production
        """
        return pulumi.get(self, "current_ruleset")

    @property
    @pulumi.getter(name="evalMode")
    def eval_mode(self) -> pulumi.Output[Optional[str]]:
        """
        . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
        """
        return pulumi.get(self, "eval_mode")

    @property
    @pulumi.getter(name="evalOperation")
    def eval_operation(self) -> pulumi.Output[str]:
        """
        . Evaluation mode operation. Allowed values are:
        - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
        - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
        - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
        - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
        - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
        """
        return pulumi.get(self, "eval_operation")

    @property
    @pulumi.getter(name="evalStatus")
    def eval_status(self) -> pulumi.Output[str]:
        """
        Whether an evaluation is currently in progress
        """
        return pulumi.get(self, "eval_status")

    @property
    @pulumi.getter(name="evaluatingRuleset")
    def evaluating_ruleset(self) -> pulumi.Output[str]:
        """
        Versioning information for the Kona Rule Set being evaluated
        """
        return pulumi.get(self, "evaluating_ruleset")

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> pulumi.Output[str]:
        """
        Date when the evaluation period ends
        """
        return pulumi.get(self, "expiration_date")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Output[str]:
        """
        . Unique identifier of the security policy associated with the evaluation process.
        """
        return pulumi.get(self, "security_policy_id")

