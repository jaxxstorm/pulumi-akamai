# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecApiRequestConstraintsArgs', 'AppSecApiRequestConstraints']

@pulumi.input_type
class AppSecApiRequestConstraintsArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 config_id: pulumi.Input[int],
                 security_policy_id: pulumi.Input[str],
                 api_endpoint_id: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a AppSecApiRequestConstraints resource.
        :param pulumi.Input[str] action: . Action to assign to the API request constraint. Allowed values are:
               - **alert**, Record the event.
               - **deny**. Block the request.
               - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
               - **none**. Take no action.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the API request constraint settings being modified.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the API request constraint settings being modified.
        :param pulumi.Input[int] api_endpoint_id: . ID of the API endpoint the constraint will be assigned to.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "security_policy_id", security_policy_id)
        if api_endpoint_id is not None:
            pulumi.set(__self__, "api_endpoint_id", api_endpoint_id)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        . Action to assign to the API request constraint. Allowed values are:
        - **alert**, Record the event.
        - **deny**. Block the request.
        - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
        - **none**. Take no action.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        . Unique identifier of the security configuration associated with the API request constraint settings being modified.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Input[str]:
        """
        . Unique identifier of the security policy associated with the API request constraint settings being modified.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_policy_id", value)

    @property
    @pulumi.getter(name="apiEndpointId")
    def api_endpoint_id(self) -> Optional[pulumi.Input[int]]:
        """
        . ID of the API endpoint the constraint will be assigned to.
        """
        return pulumi.get(self, "api_endpoint_id")

    @api_endpoint_id.setter
    def api_endpoint_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "api_endpoint_id", value)


@pulumi.input_type
class _AppSecApiRequestConstraintsState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 api_endpoint_id: Optional[pulumi.Input[int]] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AppSecApiRequestConstraints resources.
        :param pulumi.Input[str] action: . Action to assign to the API request constraint. Allowed values are:
               - **alert**, Record the event.
               - **deny**. Block the request.
               - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
               - **none**. Take no action.
        :param pulumi.Input[int] api_endpoint_id: . ID of the API endpoint the constraint will be assigned to.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the API request constraint settings being modified.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the API request constraint settings being modified.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if api_endpoint_id is not None:
            pulumi.set(__self__, "api_endpoint_id", api_endpoint_id)
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if security_policy_id is not None:
            pulumi.set(__self__, "security_policy_id", security_policy_id)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        . Action to assign to the API request constraint. Allowed values are:
        - **alert**, Record the event.
        - **deny**. Block the request.
        - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
        - **none**. Take no action.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="apiEndpointId")
    def api_endpoint_id(self) -> Optional[pulumi.Input[int]]:
        """
        . ID of the API endpoint the constraint will be assigned to.
        """
        return pulumi.get(self, "api_endpoint_id")

    @api_endpoint_id.setter
    def api_endpoint_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "api_endpoint_id", value)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        . Unique identifier of the security configuration associated with the API request constraint settings being modified.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        . Unique identifier of the security policy associated with the API request constraint settings being modified.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_policy_id", value)


class AppSecApiRequestConstraints(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 api_endpoint_id: Optional[pulumi.Input[int]] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        api_endpoint = akamai.get_app_sec_api_endpoints(config_id=configuration.config_id,
            security_policy_id="gms1_134637",
            api_name="Contracts")
        api_request_constraints = akamai.AppSecApiRequestConstraints("apiRequestConstraints",
            config_id=configuration.config_id,
            security_policy_id="gms1_134637",
            api_endpoint_id=api_endpoint.id,
            action="alert")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: . Action to assign to the API request constraint. Allowed values are:
               - **alert**, Record the event.
               - **deny**. Block the request.
               - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
               - **none**. Take no action.
        :param pulumi.Input[int] api_endpoint_id: . ID of the API endpoint the constraint will be assigned to.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the API request constraint settings being modified.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the API request constraint settings being modified.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecApiRequestConstraintsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        api_endpoint = akamai.get_app_sec_api_endpoints(config_id=configuration.config_id,
            security_policy_id="gms1_134637",
            api_name="Contracts")
        api_request_constraints = akamai.AppSecApiRequestConstraints("apiRequestConstraints",
            config_id=configuration.config_id,
            security_policy_id="gms1_134637",
            api_endpoint_id=api_endpoint.id,
            action="alert")
        ```

        :param str resource_name: The name of the resource.
        :param AppSecApiRequestConstraintsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecApiRequestConstraintsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 api_endpoint_id: Optional[pulumi.Input[int]] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppSecApiRequestConstraintsArgs.__new__(AppSecApiRequestConstraintsArgs)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            __props__.__dict__["api_endpoint_id"] = api_endpoint_id
            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if security_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_policy_id'")
            __props__.__dict__["security_policy_id"] = security_policy_id
        super(AppSecApiRequestConstraints, __self__).__init__(
            'akamai:index/appSecApiRequestConstraints:AppSecApiRequestConstraints',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            api_endpoint_id: Optional[pulumi.Input[int]] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            security_policy_id: Optional[pulumi.Input[str]] = None) -> 'AppSecApiRequestConstraints':
        """
        Get an existing AppSecApiRequestConstraints resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: . Action to assign to the API request constraint. Allowed values are:
               - **alert**, Record the event.
               - **deny**. Block the request.
               - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
               - **none**. Take no action.
        :param pulumi.Input[int] api_endpoint_id: . ID of the API endpoint the constraint will be assigned to.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the API request constraint settings being modified.
        :param pulumi.Input[str] security_policy_id: . Unique identifier of the security policy associated with the API request constraint settings being modified.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecApiRequestConstraintsState.__new__(_AppSecApiRequestConstraintsState)

        __props__.__dict__["action"] = action
        __props__.__dict__["api_endpoint_id"] = api_endpoint_id
        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["security_policy_id"] = security_policy_id
        return AppSecApiRequestConstraints(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        . Action to assign to the API request constraint. Allowed values are:
        - **alert**, Record the event.
        - **deny**. Block the request.
        - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
        - **none**. Take no action.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="apiEndpointId")
    def api_endpoint_id(self) -> pulumi.Output[Optional[int]]:
        """
        . ID of the API endpoint the constraint will be assigned to.
        """
        return pulumi.get(self, "api_endpoint_id")

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        . Unique identifier of the security configuration associated with the API request constraint settings being modified.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Output[str]:
        """
        . Unique identifier of the security policy associated with the API request constraint settings being modified.
        """
        return pulumi.get(self, "security_policy_id")

