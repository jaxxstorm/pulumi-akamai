# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['GtmCidrmapArgs', 'GtmCidrmap']

@pulumi.input_type
class GtmCidrmapArgs:
    def __init__(__self__, *,
                 default_datacenter: pulumi.Input['GtmCidrmapDefaultDatacenterArgs'],
                 domain: pulumi.Input[str],
                 assignments: Optional[pulumi.Input[Sequence[pulumi.Input['GtmCidrmapAssignmentArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a GtmCidrmap resource.
        :param pulumi.Input[str] domain: Domain name
        :param pulumi.Input[Sequence[pulumi.Input['GtmCidrmapAssignmentArgs']]] assignments: * `datacenter_id`
               * `nickname`
        :param pulumi.Input[str] name: Resource name
               * `default_datacenter`
               * `datacenter_id`
               * `nickname`
        :param pulumi.Input[bool] wait_on_complete: Wait for transaction to complete
        """
        pulumi.set(__self__, "default_datacenter", default_datacenter)
        pulumi.set(__self__, "domain", domain)
        if assignments is not None:
            pulumi.set(__self__, "assignments", assignments)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if wait_on_complete is not None:
            pulumi.set(__self__, "wait_on_complete", wait_on_complete)

    @property
    @pulumi.getter(name="defaultDatacenter")
    def default_datacenter(self) -> pulumi.Input['GtmCidrmapDefaultDatacenterArgs']:
        return pulumi.get(self, "default_datacenter")

    @default_datacenter.setter
    def default_datacenter(self, value: pulumi.Input['GtmCidrmapDefaultDatacenterArgs']):
        pulumi.set(self, "default_datacenter", value)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        Domain name
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def assignments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GtmCidrmapAssignmentArgs']]]]:
        """
        * `datacenter_id`
        * `nickname`
        """
        return pulumi.get(self, "assignments")

    @assignments.setter
    def assignments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GtmCidrmapAssignmentArgs']]]]):
        pulumi.set(self, "assignments", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Resource name
        * `default_datacenter`
        * `datacenter_id`
        * `nickname`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="waitOnComplete")
    def wait_on_complete(self) -> Optional[pulumi.Input[bool]]:
        """
        Wait for transaction to complete
        """
        return pulumi.get(self, "wait_on_complete")

    @wait_on_complete.setter
    def wait_on_complete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_on_complete", value)


class GtmCidrmap(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assignments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GtmCidrmapAssignmentArgs']]]]] = None,
                 default_datacenter: Optional[pulumi.Input[pulumi.InputType['GtmCidrmapDefaultDatacenterArgs']]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `GtmCidrmap` provides the resource for creating, configuring and importing a gtm Cidr Map to integrate easily with your existing GTM infrastructure to provide a secure, high performance, highly available and scalable solution for Global Traffic Management. Note: Import requires an ID of the format: `existing_domain_name`:`existing_map_name`

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        demo_cidrmap = akamai.GtmCidrmap("demoCidrmap",
            default_datacenter=akamai.GtmCidrmapDefaultDatacenterArgs(
                datacenter_id=5400,
                nickname="All Other CIDR Blocks",
            ),
            domain="demo_domain.akadns.net")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GtmCidrmapAssignmentArgs']]]] assignments: * `datacenter_id`
               * `nickname`
        :param pulumi.Input[str] domain: Domain name
        :param pulumi.Input[str] name: Resource name
               * `default_datacenter`
               * `datacenter_id`
               * `nickname`
        :param pulumi.Input[bool] wait_on_complete: Wait for transaction to complete
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GtmCidrmapArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `GtmCidrmap` provides the resource for creating, configuring and importing a gtm Cidr Map to integrate easily with your existing GTM infrastructure to provide a secure, high performance, highly available and scalable solution for Global Traffic Management. Note: Import requires an ID of the format: `existing_domain_name`:`existing_map_name`

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        demo_cidrmap = akamai.GtmCidrmap("demoCidrmap",
            default_datacenter=akamai.GtmCidrmapDefaultDatacenterArgs(
                datacenter_id=5400,
                nickname="All Other CIDR Blocks",
            ),
            domain="demo_domain.akadns.net")
        ```

        :param str resource_name: The name of the resource.
        :param GtmCidrmapArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GtmCidrmapArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assignments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GtmCidrmapAssignmentArgs']]]]] = None,
                 default_datacenter: Optional[pulumi.Input[pulumi.InputType['GtmCidrmapDefaultDatacenterArgs']]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['assignments'] = assignments
            if default_datacenter is None and not opts.urn:
                raise TypeError("Missing required property 'default_datacenter'")
            __props__['default_datacenter'] = default_datacenter
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__['domain'] = domain
            __props__['name'] = name
            __props__['wait_on_complete'] = wait_on_complete
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="akamai:trafficmanagement/gtmCidrmap:GtmCidrmap")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(GtmCidrmap, __self__).__init__(
            'akamai:index/gtmCidrmap:GtmCidrmap',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            assignments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GtmCidrmapAssignmentArgs']]]]] = None,
            default_datacenter: Optional[pulumi.Input[pulumi.InputType['GtmCidrmapDefaultDatacenterArgs']]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            wait_on_complete: Optional[pulumi.Input[bool]] = None) -> 'GtmCidrmap':
        """
        Get an existing GtmCidrmap resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GtmCidrmapAssignmentArgs']]]] assignments: * `datacenter_id`
               * `nickname`
        :param pulumi.Input[str] domain: Domain name
        :param pulumi.Input[str] name: Resource name
               * `default_datacenter`
               * `datacenter_id`
               * `nickname`
        :param pulumi.Input[bool] wait_on_complete: Wait for transaction to complete
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["assignments"] = assignments
        __props__["default_datacenter"] = default_datacenter
        __props__["domain"] = domain
        __props__["name"] = name
        __props__["wait_on_complete"] = wait_on_complete
        return GtmCidrmap(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def assignments(self) -> pulumi.Output[Optional[Sequence['outputs.GtmCidrmapAssignment']]]:
        """
        * `datacenter_id`
        * `nickname`
        """
        return pulumi.get(self, "assignments")

    @property
    @pulumi.getter(name="defaultDatacenter")
    def default_datacenter(self) -> pulumi.Output['outputs.GtmCidrmapDefaultDatacenter']:
        return pulumi.get(self, "default_datacenter")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        Domain name
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name
        * `default_datacenter`
        * `datacenter_id`
        * `nickname`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="waitOnComplete")
    def wait_on_complete(self) -> pulumi.Output[Optional[bool]]:
        """
        Wait for transaction to complete
        """
        return pulumi.get(self, "wait_on_complete")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

