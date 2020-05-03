# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GtmDatacenter(pulumi.CustomResource):
    city: pulumi.Output[str]
    clone_of: pulumi.Output[float]
    cloud_server_host_header_override: pulumi.Output[bool]
    """
    — (Boolean)
    * `continent`
    * `country`
    * `latitude`
    * `longitude`
    * `state_or_province`
    """
    cloud_server_targeting: pulumi.Output[bool]
    """
    — (Boolean)
    """
    continent: pulumi.Output[str]
    country: pulumi.Output[str]
    datacenter_id: pulumi.Output[float]
    default_load_object: pulumi.Output[dict]
    domain: pulumi.Output[str]
    """
    — Domain name 
    """
    latitude: pulumi.Output[float]
    longitude: pulumi.Output[float]
    nickname: pulumi.Output[str]
    """
    — datacenter nickname
    * `default_load_object`
    * `load_object`
    * `load_object_port`
    """
    ping_interval: pulumi.Output[float]
    ping_packet_size: pulumi.Output[float]
    score_penalty: pulumi.Output[float]
    servermonitor_liveness_count: pulumi.Output[float]
    servermonitor_load_count: pulumi.Output[float]
    servermonitor_pool: pulumi.Output[str]
    state_or_province: pulumi.Output[str]
    virtual: pulumi.Output[bool]
    """
    — (Boolean)
    """
    wait_on_complete: pulumi.Output[bool]
    """
    — (Boolean, Default: true) Wait for transaction to complete
    """
    def __init__(__self__, resource_name, opts=None, city=None, clone_of=None, cloud_server_host_header_override=None, cloud_server_targeting=None, continent=None, country=None, default_load_object=None, domain=None, latitude=None, longitude=None, nickname=None, state_or_province=None, wait_on_complete=None, __props__=None, __name__=None, __opts__=None):
        """
        `TrafficManagement.GtmDatacenter` provides the resource for creating, configuring and importing a gtm datacenter to integrate easily with your existing GTM infrastructure to provide a secure, high performance, highly available and scalable solution for Global Traffic Management. Note: Import requires an ID of the format: `existing_domain_name`:`existing_datacenter_id`



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] cloud_server_host_header_override: — (Boolean)
               * `continent`
               * `country`
               * `latitude`
               * `longitude`
               * `state_or_province`
        :param pulumi.Input[bool] cloud_server_targeting: — (Boolean)
        :param pulumi.Input[str] domain: — Domain name 
        :param pulumi.Input[str] nickname: — datacenter nickname
               * `default_load_object`
               * `load_object`
               * `load_object_port`
        :param pulumi.Input[bool] wait_on_complete: — (Boolean, Default: true) Wait for transaction to complete

        The **default_load_object** object supports the following:

          * `loadObject` (`pulumi.Input[str]`)
          * `loadObjectPort` (`pulumi.Input[float]`)
          * `loadServers` (`pulumi.Input[list]`) - — (List)
            * `city`
            * `clone_of`
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

            __props__['city'] = city
            __props__['clone_of'] = clone_of
            __props__['cloud_server_host_header_override'] = cloud_server_host_header_override
            __props__['cloud_server_targeting'] = cloud_server_targeting
            __props__['continent'] = continent
            __props__['country'] = country
            __props__['default_load_object'] = default_load_object
            if domain is None:
                raise TypeError("Missing required property 'domain'")
            __props__['domain'] = domain
            __props__['latitude'] = latitude
            __props__['longitude'] = longitude
            __props__['nickname'] = nickname
            __props__['state_or_province'] = state_or_province
            __props__['wait_on_complete'] = wait_on_complete
            __props__['datacenter_id'] = None
            __props__['ping_interval'] = None
            __props__['ping_packet_size'] = None
            __props__['score_penalty'] = None
            __props__['servermonitor_liveness_count'] = None
            __props__['servermonitor_load_count'] = None
            __props__['servermonitor_pool'] = None
            __props__['virtual'] = None
        super(GtmDatacenter, __self__).__init__(
            'akamai:TrafficManagement/gtmDatacenter:GtmDatacenter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, city=None, clone_of=None, cloud_server_host_header_override=None, cloud_server_targeting=None, continent=None, country=None, datacenter_id=None, default_load_object=None, domain=None, latitude=None, longitude=None, nickname=None, ping_interval=None, ping_packet_size=None, score_penalty=None, servermonitor_liveness_count=None, servermonitor_load_count=None, servermonitor_pool=None, state_or_province=None, virtual=None, wait_on_complete=None):
        """
        Get an existing GtmDatacenter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] cloud_server_host_header_override: — (Boolean)
               * `continent`
               * `country`
               * `latitude`
               * `longitude`
               * `state_or_province`
        :param pulumi.Input[bool] cloud_server_targeting: — (Boolean)
        :param pulumi.Input[str] domain: — Domain name 
        :param pulumi.Input[str] nickname: — datacenter nickname
               * `default_load_object`
               * `load_object`
               * `load_object_port`
        :param pulumi.Input[bool] virtual: — (Boolean)
        :param pulumi.Input[bool] wait_on_complete: — (Boolean, Default: true) Wait for transaction to complete

        The **default_load_object** object supports the following:

          * `loadObject` (`pulumi.Input[str]`)
          * `loadObjectPort` (`pulumi.Input[float]`)
          * `loadServers` (`pulumi.Input[list]`) - — (List)
            * `city`
            * `clone_of`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["city"] = city
        __props__["clone_of"] = clone_of
        __props__["cloud_server_host_header_override"] = cloud_server_host_header_override
        __props__["cloud_server_targeting"] = cloud_server_targeting
        __props__["continent"] = continent
        __props__["country"] = country
        __props__["datacenter_id"] = datacenter_id
        __props__["default_load_object"] = default_load_object
        __props__["domain"] = domain
        __props__["latitude"] = latitude
        __props__["longitude"] = longitude
        __props__["nickname"] = nickname
        __props__["ping_interval"] = ping_interval
        __props__["ping_packet_size"] = ping_packet_size
        __props__["score_penalty"] = score_penalty
        __props__["servermonitor_liveness_count"] = servermonitor_liveness_count
        __props__["servermonitor_load_count"] = servermonitor_load_count
        __props__["servermonitor_pool"] = servermonitor_pool
        __props__["state_or_province"] = state_or_province
        __props__["virtual"] = virtual
        __props__["wait_on_complete"] = wait_on_complete
        return GtmDatacenter(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

