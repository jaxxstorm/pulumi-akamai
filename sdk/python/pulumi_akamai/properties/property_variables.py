# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['PropertyVariables']

warnings.warn("""akamai.properties.PropertyVariables has been deprecated in favor of akamai.PropertyVariables""", DeprecationWarning)


class PropertyVariables(pulumi.CustomResource):
    warnings.warn("""akamai.properties.PropertyVariables has been deprecated in favor of akamai.PropertyVariables""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PropertyVariablesVariableArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a PropertyVariables resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        pulumi.log.warn("""PropertyVariables is deprecated: akamai.properties.PropertyVariables has been deprecated in favor of akamai.PropertyVariables""")
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

            if variables is not None and not opts.urn:
                warnings.warn("""resource \"akamai_property_variables\" is no longer supported - See Akamai Terraform Upgrade Guide""", DeprecationWarning)
                pulumi.log.warn("""variables is deprecated: resource \"akamai_property_variables\" is no longer supported - See Akamai Terraform Upgrade Guide""")
            __props__['variables'] = variables
            __props__['json'] = None
        super(PropertyVariables, __self__).__init__(
            'akamai:properties/propertyVariables:PropertyVariables',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            json: Optional[pulumi.Input[str]] = None,
            variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PropertyVariablesVariableArgs']]]]] = None) -> 'PropertyVariables':
        """
        Get an existing PropertyVariables resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] json: JSON variables representation
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["json"] = json
        __props__["variables"] = variables
        return PropertyVariables(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def json(self) -> pulumi.Output[str]:
        """
        JSON variables representation
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter
    def variables(self) -> pulumi.Output[Optional[Sequence['outputs.PropertyVariablesVariable']]]:
        return pulumi.get(self, "variables")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

