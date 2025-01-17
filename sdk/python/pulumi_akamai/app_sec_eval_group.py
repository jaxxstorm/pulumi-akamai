# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecEvalGroupArgs', 'AppSecEvalGroup']

@pulumi.input_type
class AppSecEvalGroupArgs:
    def __init__(__self__, *,
                 attack_group: pulumi.Input[str],
                 attack_group_action: pulumi.Input[str],
                 config_id: pulumi.Input[int],
                 security_policy_id: pulumi.Input[str],
                 condition_exception: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AppSecEvalGroup resource.
        :param pulumi.Input[str] attack_group: . Unique identifier of the evaluation attack group being modified.
        :param pulumi.Input[str] attack_group_action: . Action to be taken any time the attack group is triggered. Allowed values are:
               - **alert**. Record the event.
               - **deny**. Block the request.
               - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
               - **none**. Take no action.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration where evaluation is taking place.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the evaluation process.
        :param pulumi.Input[str] condition_exception: . Path to a JSON file containing properties and property values for the attack group.
        """
        pulumi.set(__self__, "attack_group", attack_group)
        pulumi.set(__self__, "attack_group_action", attack_group_action)
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "security_policy_id", security_policy_id)
        if condition_exception is not None:
            pulumi.set(__self__, "condition_exception", condition_exception)

    @property
    @pulumi.getter(name="attackGroup")
    def attack_group(self) -> pulumi.Input[str]:
        """
        . Unique identifier of the evaluation attack group being modified.
        """
        return pulumi.get(self, "attack_group")

    @attack_group.setter
    def attack_group(self, value: pulumi.Input[str]):
        pulumi.set(self, "attack_group", value)

    @property
    @pulumi.getter(name="attackGroupAction")
    def attack_group_action(self) -> pulumi.Input[str]:
        """
        . Action to be taken any time the attack group is triggered. Allowed values are:
        - **alert**. Record the event.
        - **deny**. Block the request.
        - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
        - **none**. Take no action.
        """
        return pulumi.get(self, "attack_group_action")

    @attack_group_action.setter
    def attack_group_action(self, value: pulumi.Input[str]):
        pulumi.set(self, "attack_group_action", value)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        . Unique identifier of the security configuration where evaluation is taking place.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

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
    @pulumi.getter(name="conditionException")
    def condition_exception(self) -> Optional[pulumi.Input[str]]:
        """
        . Path to a JSON file containing properties and property values for the attack group.
        """
        return pulumi.get(self, "condition_exception")

    @condition_exception.setter
    def condition_exception(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "condition_exception", value)


@pulumi.input_type
class _AppSecEvalGroupState:
    def __init__(__self__, *,
                 attack_group: Optional[pulumi.Input[str]] = None,
                 attack_group_action: Optional[pulumi.Input[str]] = None,
                 condition_exception: Optional[pulumi.Input[str]] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AppSecEvalGroup resources.
        :param pulumi.Input[str] attack_group: . Unique identifier of the evaluation attack group being modified.
        :param pulumi.Input[str] attack_group_action: . Action to be taken any time the attack group is triggered. Allowed values are:
               - **alert**. Record the event.
               - **deny**. Block the request.
               - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
               - **none**. Take no action.
        :param pulumi.Input[str] condition_exception: . Path to a JSON file containing properties and property values for the attack group.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration where evaluation is taking place.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the evaluation process.
        """
        if attack_group is not None:
            pulumi.set(__self__, "attack_group", attack_group)
        if attack_group_action is not None:
            pulumi.set(__self__, "attack_group_action", attack_group_action)
        if condition_exception is not None:
            pulumi.set(__self__, "condition_exception", condition_exception)
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if security_policy_id is not None:
            pulumi.set(__self__, "security_policy_id", security_policy_id)

    @property
    @pulumi.getter(name="attackGroup")
    def attack_group(self) -> Optional[pulumi.Input[str]]:
        """
        . Unique identifier of the evaluation attack group being modified.
        """
        return pulumi.get(self, "attack_group")

    @attack_group.setter
    def attack_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attack_group", value)

    @property
    @pulumi.getter(name="attackGroupAction")
    def attack_group_action(self) -> Optional[pulumi.Input[str]]:
        """
        . Action to be taken any time the attack group is triggered. Allowed values are:
        - **alert**. Record the event.
        - **deny**. Block the request.
        - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
        - **none**. Take no action.
        """
        return pulumi.get(self, "attack_group_action")

    @attack_group_action.setter
    def attack_group_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attack_group_action", value)

    @property
    @pulumi.getter(name="conditionException")
    def condition_exception(self) -> Optional[pulumi.Input[str]]:
        """
        . Path to a JSON file containing properties and property values for the attack group.
        """
        return pulumi.get(self, "condition_exception")

    @condition_exception.setter
    def condition_exception(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "condition_exception", value)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        . Unique identifier of the security configuration where evaluation is taking place.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

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


class AppSecEvalGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attack_group: Optional[pulumi.Input[str]] = None,
                 attack_group_action: Optional[pulumi.Input[str]] = None,
                 condition_exception: Optional[pulumi.Input[str]] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        **Scopes**: Evaluation attack group

        Modifies the action and the conditions and exceptions for an evaluation mode attack group.

        Note that this resource is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        eval_attack_group = akamai.AppSecEvalGroup("evalAttackGroup",
            config_id=configuration.config_id,
            security_policy_id="gms1_134637",
            attack_group="SQL",
            attack_group_action="deny",
            condition_exception=(lambda path: open(path).read())(f"{path['module']}/condition_exception.json"))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attack_group: . Unique identifier of the evaluation attack group being modified.
        :param pulumi.Input[str] attack_group_action: . Action to be taken any time the attack group is triggered. Allowed values are:
               - **alert**. Record the event.
               - **deny**. Block the request.
               - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
               - **none**. Take no action.
        :param pulumi.Input[str] condition_exception: . Path to a JSON file containing properties and property values for the attack group.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration where evaluation is taking place.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the evaluation process.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecEvalGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        **Scopes**: Evaluation attack group

        Modifies the action and the conditions and exceptions for an evaluation mode attack group.

        Note that this resource is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        eval_attack_group = akamai.AppSecEvalGroup("evalAttackGroup",
            config_id=configuration.config_id,
            security_policy_id="gms1_134637",
            attack_group="SQL",
            attack_group_action="deny",
            condition_exception=(lambda path: open(path).read())(f"{path['module']}/condition_exception.json"))
        ```

        :param str resource_name: The name of the resource.
        :param AppSecEvalGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecEvalGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attack_group: Optional[pulumi.Input[str]] = None,
                 attack_group_action: Optional[pulumi.Input[str]] = None,
                 condition_exception: Optional[pulumi.Input[str]] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppSecEvalGroupArgs.__new__(AppSecEvalGroupArgs)

            if attack_group is None and not opts.urn:
                raise TypeError("Missing required property 'attack_group'")
            __props__.__dict__["attack_group"] = attack_group
            if attack_group_action is None and not opts.urn:
                raise TypeError("Missing required property 'attack_group_action'")
            __props__.__dict__["attack_group_action"] = attack_group_action
            __props__.__dict__["condition_exception"] = condition_exception
            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if security_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_policy_id'")
            __props__.__dict__["security_policy_id"] = security_policy_id
        super(AppSecEvalGroup, __self__).__init__(
            'akamai:index/appSecEvalGroup:AppSecEvalGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attack_group: Optional[pulumi.Input[str]] = None,
            attack_group_action: Optional[pulumi.Input[str]] = None,
            condition_exception: Optional[pulumi.Input[str]] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            security_policy_id: Optional[pulumi.Input[str]] = None) -> 'AppSecEvalGroup':
        """
        Get an existing AppSecEvalGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attack_group: . Unique identifier of the evaluation attack group being modified.
        :param pulumi.Input[str] attack_group_action: . Action to be taken any time the attack group is triggered. Allowed values are:
               - **alert**. Record the event.
               - **deny**. Block the request.
               - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
               - **none**. Take no action.
        :param pulumi.Input[str] condition_exception: . Path to a JSON file containing properties and property values for the attack group.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration where evaluation is taking place.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the evaluation process.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecEvalGroupState.__new__(_AppSecEvalGroupState)

        __props__.__dict__["attack_group"] = attack_group
        __props__.__dict__["attack_group_action"] = attack_group_action
        __props__.__dict__["condition_exception"] = condition_exception
        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["security_policy_id"] = security_policy_id
        return AppSecEvalGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attackGroup")
    def attack_group(self) -> pulumi.Output[str]:
        """
        . Unique identifier of the evaluation attack group being modified.
        """
        return pulumi.get(self, "attack_group")

    @property
    @pulumi.getter(name="attackGroupAction")
    def attack_group_action(self) -> pulumi.Output[str]:
        """
        . Action to be taken any time the attack group is triggered. Allowed values are:
        - **alert**. Record the event.
        - **deny**. Block the request.
        - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
        - **none**. Take no action.
        """
        return pulumi.get(self, "attack_group_action")

    @property
    @pulumi.getter(name="conditionException")
    def condition_exception(self) -> pulumi.Output[Optional[str]]:
        """
        . Path to a JSON file containing properties and property values for the attack group.
        """
        return pulumi.get(self, "condition_exception")

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        . Unique identifier of the security configuration where evaluation is taking place.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Output[str]:
        """
        . Unique identifier of the security policy associated with the evaluation process.
        """
        return pulumi.get(self, "security_policy_id")

