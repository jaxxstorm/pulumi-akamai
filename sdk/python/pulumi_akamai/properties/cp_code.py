# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['CpCode']

warnings.warn("""akamai.properties.CpCode has been deprecated in favor of akamai.CpCode""", DeprecationWarning)


class CpCode(pulumi.CustomResource):
    warnings.warn("""akamai.properties.CpCode has been deprecated in favor of akamai.CpCode""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 contract: Optional[pulumi.Input[str]] = None,
                 contract_id: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 product: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a CpCode resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        pulumi.log.warn("""CpCode is deprecated: akamai.properties.CpCode has been deprecated in favor of akamai.CpCode""")
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

            if contract is not None and not opts.urn:
                warnings.warn("""use \"contract_id\" attribute instead""", DeprecationWarning)
                pulumi.log.warn("""contract is deprecated: use \"contract_id\" attribute instead""")
            __props__['contract'] = contract
            __props__['contract_id'] = contract_id
            if group is not None and not opts.urn:
                warnings.warn("""use \"group_id\" attribute instead""", DeprecationWarning)
                pulumi.log.warn("""group is deprecated: use \"group_id\" attribute instead""")
            __props__['group'] = group
            __props__['group_id'] = group_id
            __props__['name'] = name
            if product is None and not opts.urn:
                raise TypeError("Missing required property 'product'")
            __props__['product'] = product
        super(CpCode, __self__).__init__(
            'akamai:properties/cpCode:CpCode',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            contract: Optional[pulumi.Input[str]] = None,
            contract_id: Optional[pulumi.Input[str]] = None,
            group: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            product: Optional[pulumi.Input[str]] = None) -> 'CpCode':
        """
        Get an existing CpCode resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["contract"] = contract
        __props__["contract_id"] = contract_id
        __props__["group"] = group
        __props__["group_id"] = group_id
        __props__["name"] = name
        __props__["product"] = product
        return CpCode(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def contract(self) -> pulumi.Output[str]:
        return pulumi.get(self, "contract")

    @property
    @pulumi.getter(name="contractId")
    def contract_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "contract_id")

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        return pulumi.get(self, "group")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def product(self) -> pulumi.Output[str]:
        return pulumi.get(self, "product")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

