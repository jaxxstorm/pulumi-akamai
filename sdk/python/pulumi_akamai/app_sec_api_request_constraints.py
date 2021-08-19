# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
        :param pulumi.Input[str] action: The action to assign to API request constraints: either `alert`, `deny`, or `none`.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] security_policy_id: The ID of the security policy to use.
        :param pulumi.Input[int] api_endpoint_id: The ID of the API endpoint to use.
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
        The action to assign to API request constraints: either `alert`, `deny`, or `none`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        The ID of the security configuration to use.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Input[str]:
        """
        The ID of the security policy to use.
        """
        return pulumi.get(self, "security_policy_id")

    @security_policy_id.setter
    def security_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_policy_id", value)

    @property
    @pulumi.getter(name="apiEndpointId")
    def api_endpoint_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the API endpoint to use.
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
        :param pulumi.Input[str] action: The action to assign to API request constraints: either `alert`, `deny`, or `none`.
        :param pulumi.Input[int] api_endpoint_id: The ID of the API endpoint to use.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] security_policy_id: The ID of the security policy to use.
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
        The action to assign to API request constraints: either `alert`, `deny`, or `none`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="apiEndpointId")
    def api_endpoint_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the API endpoint to use.
        """
        return pulumi.get(self, "api_endpoint_id")

    @api_endpoint_id.setter
    def api_endpoint_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "api_endpoint_id", value)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the security configuration to use.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the security policy to use.
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
        The `resource_akamai_appsec_api_request_constraints` resource allows you to update what action to take when the API request constraint triggers. This operation modifies an individual API constraint action. To use this operation, use the `getAppSecApiEndpoints` data source to list one or all API endpoints, and use the ID of the selected endpoint. Use the `action` paameter to specify how the alert should be handled.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The action to assign to API request constraints: either `alert`, `deny`, or `none`.
        :param pulumi.Input[int] api_endpoint_id: The ID of the API endpoint to use.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] security_policy_id: The ID of the security policy to use.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecApiRequestConstraintsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `resource_akamai_appsec_api_request_constraints` resource allows you to update what action to take when the API request constraint triggers. This operation modifies an individual API constraint action. To use this operation, use the `getAppSecApiEndpoints` data source to list one or all API endpoints, and use the ID of the selected endpoint. Use the `action` paameter to specify how the alert should be handled.

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
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
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
        :param pulumi.Input[str] action: The action to assign to API request constraints: either `alert`, `deny`, or `none`.
        :param pulumi.Input[int] api_endpoint_id: The ID of the API endpoint to use.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] security_policy_id: The ID of the security policy to use.
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
        The action to assign to API request constraints: either `alert`, `deny`, or `none`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="apiEndpointId")
    def api_endpoint_id(self) -> pulumi.Output[Optional[int]]:
        """
        The ID of the API endpoint to use.
        """
        return pulumi.get(self, "api_endpoint_id")

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        The ID of the security configuration to use.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Output[str]:
        """
        The ID of the security policy to use.
        """
        return pulumi.get(self, "security_policy_id")
