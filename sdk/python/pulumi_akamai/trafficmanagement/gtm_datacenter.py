# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['GtmDatacenterArgs', 'GtmDatacenter']

@pulumi.input_type
class GtmDatacenterArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 city: Optional[pulumi.Input[str]] = None,
                 clone_of: Optional[pulumi.Input[int]] = None,
                 cloud_server_host_header_override: Optional[pulumi.Input[bool]] = None,
                 cloud_server_targeting: Optional[pulumi.Input[bool]] = None,
                 continent: Optional[pulumi.Input[str]] = None,
                 country: Optional[pulumi.Input[str]] = None,
                 default_load_object: Optional[pulumi.Input['GtmDatacenterDefaultLoadObjectArgs']] = None,
                 latitude: Optional[pulumi.Input[float]] = None,
                 longitude: Optional[pulumi.Input[float]] = None,
                 nickname: Optional[pulumi.Input[str]] = None,
                 state_or_province: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a GtmDatacenter resource.
        """
        pulumi.set(__self__, "domain", domain)
        if city is not None:
            pulumi.set(__self__, "city", city)
        if clone_of is not None:
            pulumi.set(__self__, "clone_of", clone_of)
        if cloud_server_host_header_override is not None:
            pulumi.set(__self__, "cloud_server_host_header_override", cloud_server_host_header_override)
        if cloud_server_targeting is not None:
            pulumi.set(__self__, "cloud_server_targeting", cloud_server_targeting)
        if continent is not None:
            pulumi.set(__self__, "continent", continent)
        if country is not None:
            pulumi.set(__self__, "country", country)
        if default_load_object is not None:
            pulumi.set(__self__, "default_load_object", default_load_object)
        if latitude is not None:
            pulumi.set(__self__, "latitude", latitude)
        if longitude is not None:
            pulumi.set(__self__, "longitude", longitude)
        if nickname is not None:
            pulumi.set(__self__, "nickname", nickname)
        if state_or_province is not None:
            pulumi.set(__self__, "state_or_province", state_or_province)
        if wait_on_complete is not None:
            pulumi.set(__self__, "wait_on_complete", wait_on_complete)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def city(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "city")

    @city.setter
    def city(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "city", value)

    @property
    @pulumi.getter(name="cloneOf")
    def clone_of(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "clone_of")

    @clone_of.setter
    def clone_of(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "clone_of", value)

    @property
    @pulumi.getter(name="cloudServerHostHeaderOverride")
    def cloud_server_host_header_override(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "cloud_server_host_header_override")

    @cloud_server_host_header_override.setter
    def cloud_server_host_header_override(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cloud_server_host_header_override", value)

    @property
    @pulumi.getter(name="cloudServerTargeting")
    def cloud_server_targeting(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "cloud_server_targeting")

    @cloud_server_targeting.setter
    def cloud_server_targeting(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cloud_server_targeting", value)

    @property
    @pulumi.getter
    def continent(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "continent")

    @continent.setter
    def continent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "continent", value)

    @property
    @pulumi.getter
    def country(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "country")

    @country.setter
    def country(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country", value)

    @property
    @pulumi.getter(name="defaultLoadObject")
    def default_load_object(self) -> Optional[pulumi.Input['GtmDatacenterDefaultLoadObjectArgs']]:
        return pulumi.get(self, "default_load_object")

    @default_load_object.setter
    def default_load_object(self, value: Optional[pulumi.Input['GtmDatacenterDefaultLoadObjectArgs']]):
        pulumi.set(self, "default_load_object", value)

    @property
    @pulumi.getter
    def latitude(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "latitude")

    @latitude.setter
    def latitude(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "latitude", value)

    @property
    @pulumi.getter
    def longitude(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "longitude")

    @longitude.setter
    def longitude(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "longitude", value)

    @property
    @pulumi.getter
    def nickname(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "nickname")

    @nickname.setter
    def nickname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nickname", value)

    @property
    @pulumi.getter(name="stateOrProvince")
    def state_or_province(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "state_or_province")

    @state_or_province.setter
    def state_or_province(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state_or_province", value)

    @property
    @pulumi.getter(name="waitOnComplete")
    def wait_on_complete(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "wait_on_complete")

    @wait_on_complete.setter
    def wait_on_complete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_on_complete", value)


@pulumi.input_type
class _GtmDatacenterState:
    def __init__(__self__, *,
                 city: Optional[pulumi.Input[str]] = None,
                 clone_of: Optional[pulumi.Input[int]] = None,
                 cloud_server_host_header_override: Optional[pulumi.Input[bool]] = None,
                 cloud_server_targeting: Optional[pulumi.Input[bool]] = None,
                 continent: Optional[pulumi.Input[str]] = None,
                 country: Optional[pulumi.Input[str]] = None,
                 datacenter_id: Optional[pulumi.Input[int]] = None,
                 default_load_object: Optional[pulumi.Input['GtmDatacenterDefaultLoadObjectArgs']] = None,
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
                 wait_on_complete: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering GtmDatacenter resources.
        """
        if city is not None:
            pulumi.set(__self__, "city", city)
        if clone_of is not None:
            pulumi.set(__self__, "clone_of", clone_of)
        if cloud_server_host_header_override is not None:
            pulumi.set(__self__, "cloud_server_host_header_override", cloud_server_host_header_override)
        if cloud_server_targeting is not None:
            pulumi.set(__self__, "cloud_server_targeting", cloud_server_targeting)
        if continent is not None:
            pulumi.set(__self__, "continent", continent)
        if country is not None:
            pulumi.set(__self__, "country", country)
        if datacenter_id is not None:
            pulumi.set(__self__, "datacenter_id", datacenter_id)
        if default_load_object is not None:
            pulumi.set(__self__, "default_load_object", default_load_object)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if latitude is not None:
            pulumi.set(__self__, "latitude", latitude)
        if longitude is not None:
            pulumi.set(__self__, "longitude", longitude)
        if nickname is not None:
            pulumi.set(__self__, "nickname", nickname)
        if ping_interval is not None:
            pulumi.set(__self__, "ping_interval", ping_interval)
        if ping_packet_size is not None:
            pulumi.set(__self__, "ping_packet_size", ping_packet_size)
        if score_penalty is not None:
            pulumi.set(__self__, "score_penalty", score_penalty)
        if servermonitor_liveness_count is not None:
            pulumi.set(__self__, "servermonitor_liveness_count", servermonitor_liveness_count)
        if servermonitor_load_count is not None:
            pulumi.set(__self__, "servermonitor_load_count", servermonitor_load_count)
        if servermonitor_pool is not None:
            pulumi.set(__self__, "servermonitor_pool", servermonitor_pool)
        if state_or_province is not None:
            pulumi.set(__self__, "state_or_province", state_or_province)
        if virtual is not None:
            pulumi.set(__self__, "virtual", virtual)
        if wait_on_complete is not None:
            pulumi.set(__self__, "wait_on_complete", wait_on_complete)

    @property
    @pulumi.getter
    def city(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "city")

    @city.setter
    def city(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "city", value)

    @property
    @pulumi.getter(name="cloneOf")
    def clone_of(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "clone_of")

    @clone_of.setter
    def clone_of(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "clone_of", value)

    @property
    @pulumi.getter(name="cloudServerHostHeaderOverride")
    def cloud_server_host_header_override(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "cloud_server_host_header_override")

    @cloud_server_host_header_override.setter
    def cloud_server_host_header_override(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cloud_server_host_header_override", value)

    @property
    @pulumi.getter(name="cloudServerTargeting")
    def cloud_server_targeting(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "cloud_server_targeting")

    @cloud_server_targeting.setter
    def cloud_server_targeting(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cloud_server_targeting", value)

    @property
    @pulumi.getter
    def continent(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "continent")

    @continent.setter
    def continent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "continent", value)

    @property
    @pulumi.getter
    def country(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "country")

    @country.setter
    def country(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country", value)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "datacenter_id")

    @datacenter_id.setter
    def datacenter_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "datacenter_id", value)

    @property
    @pulumi.getter(name="defaultLoadObject")
    def default_load_object(self) -> Optional[pulumi.Input['GtmDatacenterDefaultLoadObjectArgs']]:
        return pulumi.get(self, "default_load_object")

    @default_load_object.setter
    def default_load_object(self, value: Optional[pulumi.Input['GtmDatacenterDefaultLoadObjectArgs']]):
        pulumi.set(self, "default_load_object", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def latitude(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "latitude")

    @latitude.setter
    def latitude(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "latitude", value)

    @property
    @pulumi.getter
    def longitude(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "longitude")

    @longitude.setter
    def longitude(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "longitude", value)

    @property
    @pulumi.getter
    def nickname(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "nickname")

    @nickname.setter
    def nickname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nickname", value)

    @property
    @pulumi.getter(name="pingInterval")
    def ping_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "ping_interval")

    @ping_interval.setter
    def ping_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ping_interval", value)

    @property
    @pulumi.getter(name="pingPacketSize")
    def ping_packet_size(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "ping_packet_size")

    @ping_packet_size.setter
    def ping_packet_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ping_packet_size", value)

    @property
    @pulumi.getter(name="scorePenalty")
    def score_penalty(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "score_penalty")

    @score_penalty.setter
    def score_penalty(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "score_penalty", value)

    @property
    @pulumi.getter(name="servermonitorLivenessCount")
    def servermonitor_liveness_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "servermonitor_liveness_count")

    @servermonitor_liveness_count.setter
    def servermonitor_liveness_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "servermonitor_liveness_count", value)

    @property
    @pulumi.getter(name="servermonitorLoadCount")
    def servermonitor_load_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "servermonitor_load_count")

    @servermonitor_load_count.setter
    def servermonitor_load_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "servermonitor_load_count", value)

    @property
    @pulumi.getter(name="servermonitorPool")
    def servermonitor_pool(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "servermonitor_pool")

    @servermonitor_pool.setter
    def servermonitor_pool(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "servermonitor_pool", value)

    @property
    @pulumi.getter(name="stateOrProvince")
    def state_or_province(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "state_or_province")

    @state_or_province.setter
    def state_or_province(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state_or_province", value)

    @property
    @pulumi.getter
    def virtual(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "virtual")

    @virtual.setter
    def virtual(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "virtual", value)

    @property
    @pulumi.getter(name="waitOnComplete")
    def wait_on_complete(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "wait_on_complete")

    @wait_on_complete.setter
    def wait_on_complete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_on_complete", value)


warnings.warn("""akamai.trafficmanagement.GtmDatacenter has been deprecated in favor of akamai.GtmDatacenter""", DeprecationWarning)


class GtmDatacenter(pulumi.CustomResource):
    warnings.warn("""akamai.trafficmanagement.GtmDatacenter has been deprecated in favor of akamai.GtmDatacenter""", DeprecationWarning)

    @overload
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
                 __props__=None):
        """
        Create a GtmDatacenter resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GtmDatacenterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a GtmDatacenter resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param GtmDatacenterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GtmDatacenterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
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
                 __props__=None):
        pulumi.log.warn("""GtmDatacenter is deprecated: akamai.trafficmanagement.GtmDatacenter has been deprecated in favor of akamai.GtmDatacenter""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GtmDatacenterArgs.__new__(GtmDatacenterArgs)

            __props__.__dict__["city"] = city
            __props__.__dict__["clone_of"] = clone_of
            __props__.__dict__["cloud_server_host_header_override"] = cloud_server_host_header_override
            __props__.__dict__["cloud_server_targeting"] = cloud_server_targeting
            __props__.__dict__["continent"] = continent
            __props__.__dict__["country"] = country
            __props__.__dict__["default_load_object"] = default_load_object
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            __props__.__dict__["latitude"] = latitude
            __props__.__dict__["longitude"] = longitude
            __props__.__dict__["nickname"] = nickname
            __props__.__dict__["state_or_province"] = state_or_province
            __props__.__dict__["wait_on_complete"] = wait_on_complete
            __props__.__dict__["datacenter_id"] = None
            __props__.__dict__["ping_interval"] = None
            __props__.__dict__["ping_packet_size"] = None
            __props__.__dict__["score_penalty"] = None
            __props__.__dict__["servermonitor_liveness_count"] = None
            __props__.__dict__["servermonitor_load_count"] = None
            __props__.__dict__["servermonitor_pool"] = None
            __props__.__dict__["virtual"] = None
        super(GtmDatacenter, __self__).__init__(
            'akamai:trafficmanagement/gtmDatacenter:GtmDatacenter',
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
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GtmDatacenterState.__new__(_GtmDatacenterState)

        __props__.__dict__["city"] = city
        __props__.__dict__["clone_of"] = clone_of
        __props__.__dict__["cloud_server_host_header_override"] = cloud_server_host_header_override
        __props__.__dict__["cloud_server_targeting"] = cloud_server_targeting
        __props__.__dict__["continent"] = continent
        __props__.__dict__["country"] = country
        __props__.__dict__["datacenter_id"] = datacenter_id
        __props__.__dict__["default_load_object"] = default_load_object
        __props__.__dict__["domain"] = domain
        __props__.__dict__["latitude"] = latitude
        __props__.__dict__["longitude"] = longitude
        __props__.__dict__["nickname"] = nickname
        __props__.__dict__["ping_interval"] = ping_interval
        __props__.__dict__["ping_packet_size"] = ping_packet_size
        __props__.__dict__["score_penalty"] = score_penalty
        __props__.__dict__["servermonitor_liveness_count"] = servermonitor_liveness_count
        __props__.__dict__["servermonitor_load_count"] = servermonitor_load_count
        __props__.__dict__["servermonitor_pool"] = servermonitor_pool
        __props__.__dict__["state_or_province"] = state_or_province
        __props__.__dict__["virtual"] = virtual
        __props__.__dict__["wait_on_complete"] = wait_on_complete
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
        return pulumi.get(self, "wait_on_complete")

