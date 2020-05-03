# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GtmCidrmap(pulumi.CustomResource):
    assignments: pulumi.Output[list]
    """
    — (multiple allowed)
    * `datacenter_id`
    * `nickname`

      * `blocks` (`list`) - — (List)
      * `datacenter_id` (`float`)
      * `nickname` (`str`)
    """
    default_datacenter: pulumi.Output[dict]
    domain: pulumi.Output[str]
    """
    — Domain name 
    """
    name: pulumi.Output[str]
    """
    — Resource name
    * `default_datacenter`
    * `datacenter_id`
    * `nickname`
    """
    wait_on_complete: pulumi.Output[bool]
    """
    — (Boolean, Default: true) Wait for transaction to complete
    """
    def __init__(__self__, resource_name, opts=None, assignments=None, default_datacenter=None, domain=None, name=None, wait_on_complete=None, __props__=None, __name__=None, __opts__=None):
        """
        `TrafficManagement.GtmCidrmap` provides the resource for creating, configuring and importing a gtm Cidr Map to integrate easily with your existing GTM infrastructure to provide a secure, high performance, highly available and scalable solution for Global Traffic Management. Note: Import requires an ID of the format: `existing_domain_name`:`existing_map_name`



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] assignments: — (multiple allowed)
               * `datacenter_id`
               * `nickname`
        :param pulumi.Input[str] domain: — Domain name 
        :param pulumi.Input[str] name: — Resource name
               * `default_datacenter`
               * `datacenter_id`
               * `nickname`
        :param pulumi.Input[bool] wait_on_complete: — (Boolean, Default: true) Wait for transaction to complete

        The **assignments** object supports the following:

          * `blocks` (`pulumi.Input[list]`) - — (List)
          * `datacenter_id` (`pulumi.Input[float]`)
          * `nickname` (`pulumi.Input[str]`)

        The **default_datacenter** object supports the following:

          * `datacenter_id` (`pulumi.Input[float]`)
          * `nickname` (`pulumi.Input[str]`)
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['assignments'] = assignments
            if default_datacenter is None:
                raise TypeError("Missing required property 'default_datacenter'")
            __props__['default_datacenter'] = default_datacenter
            if domain is None:
                raise TypeError("Missing required property 'domain'")
            __props__['domain'] = domain
            __props__['name'] = name
            __props__['wait_on_complete'] = wait_on_complete
        super(GtmCidrmap, __self__).__init__(
            'akamai:TrafficManagement/gtmCidrmap:GtmCidrmap',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, assignments=None, default_datacenter=None, domain=None, name=None, wait_on_complete=None):
        """
        Get an existing GtmCidrmap resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] assignments: — (multiple allowed)
               * `datacenter_id`
               * `nickname`
        :param pulumi.Input[str] domain: — Domain name 
        :param pulumi.Input[str] name: — Resource name
               * `default_datacenter`
               * `datacenter_id`
               * `nickname`
        :param pulumi.Input[bool] wait_on_complete: — (Boolean, Default: true) Wait for transaction to complete

        The **assignments** object supports the following:

          * `blocks` (`pulumi.Input[list]`) - — (List)
          * `datacenter_id` (`pulumi.Input[float]`)
          * `nickname` (`pulumi.Input[str]`)

        The **default_datacenter** object supports the following:

          * `datacenter_id` (`pulumi.Input[float]`)
          * `nickname` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["assignments"] = assignments
        __props__["default_datacenter"] = default_datacenter
        __props__["domain"] = domain
        __props__["name"] = name
        __props__["wait_on_complete"] = wait_on_complete
        return GtmCidrmap(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

