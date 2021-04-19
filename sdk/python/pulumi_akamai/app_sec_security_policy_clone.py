# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecSecurityPolicyCloneArgs', 'AppSecSecurityPolicyClone']

@pulumi.input_type
class AppSecSecurityPolicyCloneArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 create_from_security_policy: pulumi.Input[str],
                 version: pulumi.Input[int],
                 policy_name: Optional[pulumi.Input[str]] = None,
                 policy_prefix: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AppSecSecurityPolicyClone resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] create_from_security_policy: The ID of the security policy to clone.
        :param pulumi.Input[str] policy_name: The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.
        :param pulumi.Input[str] policy_prefix: The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "create_from_security_policy", create_from_security_policy)
        pulumi.set(__self__, "version", version)
        if policy_name is not None:
            pulumi.set(__self__, "policy_name", policy_name)
        if policy_prefix is not None:
            pulumi.set(__self__, "policy_prefix", policy_prefix)

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
    @pulumi.getter(name="createFromSecurityPolicy")
    def create_from_security_policy(self) -> pulumi.Input[str]:
        """
        The ID of the security policy to clone.
        """
        return pulumi.get(self, "create_from_security_policy")

    @create_from_security_policy.setter
    def create_from_security_policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "create_from_security_policy", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[int]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[int]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.
        """
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_name", value)

    @property
    @pulumi.getter(name="policyPrefix")
    def policy_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated.
        """
        return pulumi.get(self, "policy_prefix")

    @policy_prefix.setter
    def policy_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_prefix", value)


@pulumi.input_type
class _AppSecSecurityPolicyCloneState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 create_from_security_policy: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 policy_prefix: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering AppSecSecurityPolicyClone resources.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] create_from_security_policy: The ID of the security policy to clone.
        :param pulumi.Input[str] policy_id: The ID of the new security policy.
        :param pulumi.Input[str] policy_name: The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.
        :param pulumi.Input[str] policy_prefix: The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated.
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if create_from_security_policy is not None:
            pulumi.set(__self__, "create_from_security_policy", create_from_security_policy)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if policy_name is not None:
            pulumi.set(__self__, "policy_name", policy_name)
        if policy_prefix is not None:
            pulumi.set(__self__, "policy_prefix", policy_prefix)
        if version is not None:
            pulumi.set(__self__, "version", version)

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
    @pulumi.getter(name="createFromSecurityPolicy")
    def create_from_security_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the security policy to clone.
        """
        return pulumi.get(self, "create_from_security_policy")

    @create_from_security_policy.setter
    def create_from_security_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_from_security_policy", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the new security policy.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.
        """
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_name", value)

    @property
    @pulumi.getter(name="policyPrefix")
    def policy_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated.
        """
        return pulumi.get(self, "policy_prefix")

    @policy_prefix.setter
    def policy_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_prefix", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class AppSecSecurityPolicyClone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 create_from_security_policy: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 policy_prefix: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        The `AppSecSecurityPolicyClone` resource allows you to create a new security policy by cloning an existing policy.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Akamai Tools")
        security_policy_clone_app_sec_security_policy_clone = akamai.AppSecSecurityPolicyClone("securityPolicyCloneAppSecSecurityPolicyClone",
            config_id=configuration.config_id,
            version=configuration.latest_version,
            create_from_security_policy="crAP_75829",
            policy_name="Test Policy 1",
            policy_prefix="TP1")
        pulumi.export("securityPolicyClone", security_policy_clone_app_sec_security_policy_clone.policy_id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] create_from_security_policy: The ID of the security policy to clone.
        :param pulumi.Input[str] policy_name: The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.
        :param pulumi.Input[str] policy_prefix: The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecSecurityPolicyCloneArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `AppSecSecurityPolicyClone` resource allows you to create a new security policy by cloning an existing policy.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Akamai Tools")
        security_policy_clone_app_sec_security_policy_clone = akamai.AppSecSecurityPolicyClone("securityPolicyCloneAppSecSecurityPolicyClone",
            config_id=configuration.config_id,
            version=configuration.latest_version,
            create_from_security_policy="crAP_75829",
            policy_name="Test Policy 1",
            policy_prefix="TP1")
        pulumi.export("securityPolicyClone", security_policy_clone_app_sec_security_policy_clone.policy_id)
        ```

        :param str resource_name: The name of the resource.
        :param AppSecSecurityPolicyCloneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecSecurityPolicyCloneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 create_from_security_policy: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 policy_prefix: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
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
            __props__ = AppSecSecurityPolicyCloneArgs.__new__(AppSecSecurityPolicyCloneArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if create_from_security_policy is None and not opts.urn:
                raise TypeError("Missing required property 'create_from_security_policy'")
            __props__.__dict__["create_from_security_policy"] = create_from_security_policy
            __props__.__dict__["policy_name"] = policy_name
            __props__.__dict__["policy_prefix"] = policy_prefix
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__.__dict__["version"] = version
            __props__.__dict__["policy_id"] = None
        super(AppSecSecurityPolicyClone, __self__).__init__(
            'akamai:index/appSecSecurityPolicyClone:AppSecSecurityPolicyClone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            create_from_security_policy: Optional[pulumi.Input[str]] = None,
            policy_id: Optional[pulumi.Input[str]] = None,
            policy_name: Optional[pulumi.Input[str]] = None,
            policy_prefix: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'AppSecSecurityPolicyClone':
        """
        Get an existing AppSecSecurityPolicyClone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] create_from_security_policy: The ID of the security policy to clone.
        :param pulumi.Input[str] policy_id: The ID of the new security policy.
        :param pulumi.Input[str] policy_name: The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.
        :param pulumi.Input[str] policy_prefix: The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecSecurityPolicyCloneState.__new__(_AppSecSecurityPolicyCloneState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["create_from_security_policy"] = create_from_security_policy
        __props__.__dict__["policy_id"] = policy_id
        __props__.__dict__["policy_name"] = policy_name
        __props__.__dict__["policy_prefix"] = policy_prefix
        __props__.__dict__["version"] = version
        return AppSecSecurityPolicyClone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        The ID of the security configuration to use.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="createFromSecurityPolicy")
    def create_from_security_policy(self) -> pulumi.Output[str]:
        """
        The ID of the security policy to clone.
        """
        return pulumi.get(self, "create_from_security_policy")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        The ID of the new security policy.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.
        """
        return pulumi.get(self, "policy_name")

    @property
    @pulumi.getter(name="policyPrefix")
    def policy_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated.
        """
        return pulumi.get(self, "policy_prefix")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        return pulumi.get(self, "version")

