# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['GtmDatacenter']


class GtmDatacenter(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 city: Optional[pulumi.Input[str]] = None,
                 clone_of: Optional[pulumi.Input[int]] = None,
                 cloud_server_host_header_override: Optional[pulumi.Input[bool]] = None,
                 cloud_server_targeting: Optional[pulumi.Input[bool]] = None,
                 continent: Optional[pulumi.Input[str]] = None,
                 country: Optional[pulumi.Input[str]] = None,
                 default_load_object: Optional[pulumi.Input[pulumi.InputType['GtmDatacenterDefaultLoadObjectArgs']]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 latitude: Optional[pulumi.Input[float]] = None,
                 longitude: Optional[pulumi.Input[float]] = None,
                 nickname: Optional[pulumi.Input[str]] = None,
                 state_or_province: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `GtmDatacenter` provides the resource for creating, configuring and importing a gtm datacenter to integrate easily with your existing GTM infrastructure to provide a secure, high performance, highly available and scalable solution for Global Traffic Management. Note: Import requires an ID of the format: `existing_domain_name`:`existing_datacenter_id`

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        demo_datacenter = akamai.GtmDatacenter("demoDatacenter",
            domain="demo_domain.akadns.net",
            nickname="demo_datacenter")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] cloud_server_host_header_override: * `continent`
               * `country`
               * `latitude`
               * `longitude`
               * `state_or_province`
        :param pulumi.Input[str] domain: Domain name
        :param pulumi.Input[str] nickname: datacenter nickname
               * `default_load_object`
               * `load_object`
               * `load_object_port`
        :param pulumi.Input[bool] wait_on_complete: Wait for transaction to complete
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

            __props__['city'] = city
            __props__['clone_of'] = clone_of
            __props__['cloud_server_host_header_override'] = cloud_server_host_header_override
            __props__['cloud_server_targeting'] = cloud_server_targeting
            __props__['continent'] = continent
            __props__['country'] = country
            __props__['default_load_object'] = default_load_object
            if domain is None and not opts.urn:
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
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="akamai:trafficmanagement/gtmDatacenter:GtmDatacenter")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(GtmDatacenter, __self__).__init__(
            'akamai:index/gtmDatacenter:GtmDatacenter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            city: Optional[pulumi.Input[str]] = None,
            clone_of: Optional[pulumi.Input[int]] = None,
            cloud_server_host_header_override: Optional[pulumi.Input[bool]] = None,
            cloud_server_targeting: Optional[pulumi.Input[bool]] = None,
            continent: Optional[pulumi.Input[str]] = None,
            country: Optional[pulumi.Input[str]] = None,
            datacenter_id: Optional[pulumi.Input[int]] = None,
            default_load_object: Optional[pulumi.Input[pulumi.InputType['GtmDatacenterDefaultLoadObjectArgs']]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            latitude: Optional[pulumi.Input[float]] = None,
            longitude: Optional[pulumi.Input[float]] = None,
            nickname: Optional[pulumi.Input[str]] = None,
            ping_interval: Optional[pulumi.Input[int]] = None,
            ping_packet_size: Optional[pulumi.Input[int]] = None,
            score_penalty: Optional[pulumi.Input[int]] = None,
            servermonitor_liveness_count: Optional[pulumi.Input[int]] = None,
            servermonitor_load_count: Optional[pulumi.Input[int]] = None,
            servermonitor_pool: Optional[pulumi.Input[str]] = None,
            state_or_province: Optional[pulumi.Input[str]] = None,
            virtual: Optional[pulumi.Input[bool]] = None,
            wait_on_complete: Optional[pulumi.Input[bool]] = None) -> 'GtmDatacenter':
        """
        Get an existing GtmDatacenter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] cloud_server_host_header_override: * `continent`
               * `country`
               * `latitude`
               * `longitude`
               * `state_or_province`
        :param pulumi.Input[str] domain: Domain name
        :param pulumi.Input[str] nickname: datacenter nickname
               * `default_load_object`
               * `load_object`
               * `load_object_port`
        :param pulumi.Input[bool] wait_on_complete: Wait for transaction to complete
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

    @property
    @pulumi.getter
    def city(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "city")

    @property
    @pulumi.getter(name="cloneOf")
    def clone_of(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "clone_of")

    @property
    @pulumi.getter(name="cloudServerHostHeaderOverride")
    def cloud_server_host_header_override(self) -> pulumi.Output[Optional[bool]]:
        """
        * `continent`
        * `country`
        * `latitude`
        * `longitude`
        * `state_or_province`
        """
        return pulumi.get(self, "cloud_server_host_header_override")

    @property
    @pulumi.getter(name="cloudServerTargeting")
    def cloud_server_targeting(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "cloud_server_targeting")

    @property
    @pulumi.getter
    def continent(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "continent")

    @property
    @pulumi.getter
    def country(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "country")

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "datacenter_id")

    @property
    @pulumi.getter(name="defaultLoadObject")
    def default_load_object(self) -> pulumi.Output[Optional['outputs.GtmDatacenterDefaultLoadObject']]:
        return pulumi.get(self, "default_load_object")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        Domain name
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def latitude(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "latitude")

    @property
    @pulumi.getter
    def longitude(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "longitude")

    @property
    @pulumi.getter
    def nickname(self) -> pulumi.Output[Optional[str]]:
        """
        datacenter nickname
        * `default_load_object`
        * `load_object`
        * `load_object_port`
        """
        return pulumi.get(self, "nickname")

    @property
    @pulumi.getter(name="pingInterval")
    def ping_interval(self) -> pulumi.Output[int]:
        return pulumi.get(self, "ping_interval")

    @property
    @pulumi.getter(name="pingPacketSize")
    def ping_packet_size(self) -> pulumi.Output[int]:
        return pulumi.get(self, "ping_packet_size")

    @property
    @pulumi.getter(name="scorePenalty")
    def score_penalty(self) -> pulumi.Output[int]:
        return pulumi.get(self, "score_penalty")

    @property
    @pulumi.getter(name="servermonitorLivenessCount")
    def servermonitor_liveness_count(self) -> pulumi.Output[int]:
        return pulumi.get(self, "servermonitor_liveness_count")

    @property
    @pulumi.getter(name="servermonitorLoadCount")
    def servermonitor_load_count(self) -> pulumi.Output[int]:
        return pulumi.get(self, "servermonitor_load_count")

    @property
    @pulumi.getter(name="servermonitorPool")
    def servermonitor_pool(self) -> pulumi.Output[str]:
        return pulumi.get(self, "servermonitor_pool")

    @property
    @pulumi.getter(name="stateOrProvince")
    def state_or_province(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "state_or_province")

    @property
    @pulumi.getter
    def virtual(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "virtual")

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

