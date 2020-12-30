# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['AppSecSecurityPolicyClone']


class AppSecSecurityPolicyClone(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 create_from_security_policy: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 policy_prefix: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__['config_id'] = config_id
            if create_from_security_policy is None and not opts.urn:
                raise TypeError("Missing required property 'create_from_security_policy'")
            __props__['create_from_security_policy'] = create_from_security_policy
            __props__['policy_name'] = policy_name
            __props__['policy_prefix'] = policy_prefix
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__['version'] = version
            __props__['policy_id'] = None
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

        __props__ = dict()

        __props__["config_id"] = config_id
        __props__["create_from_security_policy"] = create_from_security_policy
        __props__["policy_id"] = policy_id
        __props__["policy_name"] = policy_name
        __props__["policy_prefix"] = policy_prefix
        __props__["version"] = version
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

