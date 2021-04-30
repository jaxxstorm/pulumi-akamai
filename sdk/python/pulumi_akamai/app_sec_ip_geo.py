# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecIPGeoArgs', 'AppSecIPGeo']

@pulumi.input_type
class AppSecIPGeoArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 mode: pulumi.Input[str],
                 security_policy_id: pulumi.Input[str],
                 version: pulumi.Input[int],
                 exception_ip_network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 geo_network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a AppSecIPGeo resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[str] mode: The mode to use for IP/Geo firewall blocking: `block` to block specific IPs, geographies or network lists, or `allow` to allow specific IPs or geographies to be let through while blocking the rest.
        :param pulumi.Input[str] security_policy_id: The ID of the security policy to use.
        :param pulumi.Input[int] version: The version number of the security configuration to use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exception_ip_network_lists: The network lists to be allowed regardless of `mode`, `geo_network_lists`, and `ip_network_lists` parameters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] geo_network_lists: The network lists to be blocked or allowed geographically.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_network_lists: The network lists to be blocked or allowd by IP address.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "mode", mode)
        pulumi.set(__self__, "security_policy_id", security_policy_id)
        pulumi.set(__self__, "version", version)
        if exception_ip_network_lists is not None:
            pulumi.set(__self__, "exception_ip_network_lists", exception_ip_network_lists)
        if geo_network_lists is not None:
            pulumi.set(__self__, "geo_network_lists", geo_network_lists)
        if ip_network_lists is not None:
            pulumi.set(__self__, "ip_network_lists", ip_network_lists)

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
    @pulumi.getter
    def mode(self) -> pulumi.Input[str]:
        """
        The mode to use for IP/Geo firewall blocking: `block` to block specific IPs, geographies or network lists, or `allow` to allow specific IPs or geographies to be let through while blocking the rest.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "mode", value)

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
    @pulumi.getter
    def version(self) -> pulumi.Input[int]:
        """
        The version number of the security configuration to use.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[int]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter(name="exceptionIpNetworkLists")
    def exception_ip_network_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The network lists to be allowed regardless of `mode`, `geo_network_lists`, and `ip_network_lists` parameters.
        """
        return pulumi.get(self, "exception_ip_network_lists")

    @exception_ip_network_lists.setter
    def exception_ip_network_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "exception_ip_network_lists", value)

    @property
    @pulumi.getter(name="geoNetworkLists")
    def geo_network_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The network lists to be blocked or allowed geographically.
        """
        return pulumi.get(self, "geo_network_lists")

    @geo_network_lists.setter
    def geo_network_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "geo_network_lists", value)

    @property
    @pulumi.getter(name="ipNetworkLists")
    def ip_network_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The network lists to be blocked or allowd by IP address.
        """
        return pulumi.get(self, "ip_network_lists")

    @ip_network_lists.setter
    def ip_network_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_network_lists", value)


@pulumi.input_type
class _AppSecIPGeoState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[int]] = None,
                 exception_ip_network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 geo_network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering AppSecIPGeo resources.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exception_ip_network_lists: The network lists to be allowed regardless of `mode`, `geo_network_lists`, and `ip_network_lists` parameters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] geo_network_lists: The network lists to be blocked or allowed geographically.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_network_lists: The network lists to be blocked or allowd by IP address.
        :param pulumi.Input[str] mode: The mode to use for IP/Geo firewall blocking: `block` to block specific IPs, geographies or network lists, or `allow` to allow specific IPs or geographies to be let through while blocking the rest.
        :param pulumi.Input[str] security_policy_id: The ID of the security policy to use.
        :param pulumi.Input[int] version: The version number of the security configuration to use.
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if exception_ip_network_lists is not None:
            pulumi.set(__self__, "exception_ip_network_lists", exception_ip_network_lists)
        if geo_network_lists is not None:
            pulumi.set(__self__, "geo_network_lists", geo_network_lists)
        if ip_network_lists is not None:
            pulumi.set(__self__, "ip_network_lists", ip_network_lists)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if security_policy_id is not None:
            pulumi.set(__self__, "security_policy_id", security_policy_id)
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
    @pulumi.getter(name="exceptionIpNetworkLists")
    def exception_ip_network_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The network lists to be allowed regardless of `mode`, `geo_network_lists`, and `ip_network_lists` parameters.
        """
        return pulumi.get(self, "exception_ip_network_lists")

    @exception_ip_network_lists.setter
    def exception_ip_network_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "exception_ip_network_lists", value)

    @property
    @pulumi.getter(name="geoNetworkLists")
    def geo_network_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The network lists to be blocked or allowed geographically.
        """
        return pulumi.get(self, "geo_network_lists")

    @geo_network_lists.setter
    def geo_network_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "geo_network_lists", value)

    @property
    @pulumi.getter(name="ipNetworkLists")
    def ip_network_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The network lists to be blocked or allowd by IP address.
        """
        return pulumi.get(self, "ip_network_lists")

    @ip_network_lists.setter
    def ip_network_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_network_lists", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        The mode to use for IP/Geo firewall blocking: `block` to block specific IPs, geographies or network lists, or `allow` to allow specific IPs or geographies to be let through while blocking the rest.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

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

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        The version number of the security configuration to use.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class AppSecIPGeo(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 exception_ip_network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 geo_network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Use the `AppSecIPGeo` resource to update the method and which network lists to use for IP/Geo firewall blocking.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
        # USE CASE: user wants to update the IP/GEO firewall mode to "block specific IPs/Subnets and Geos" and update the IP list, GEO list & Exception list
        ip_geo_block = akamai.AppSecIPGeo("ipGeoBlock",
            config_id=configuration.config_id,
            version=configuration.latest_version,
            security_policy_id=var["security_policy_id1"],
            mode=var["block"],
            geo_network_lists=var["geo_network_lists"],
            ip_network_lists=var["ip_network_lists"],
            exception_ip_network_lists=var["exception_ip_network_lists"])
        # USE CASE: user wants to update the IP/GEO firewall mode to "block all traffic except IPs/Subnets in block exceptions" and update the Exception list
        ip_geo_allow = akamai.AppSecIPGeo("ipGeoAllow",
            config_id=configuration.config_id,
            version=configuration.latest_version,
            security_policy_id=var["security_policy_id2"],
            mode=var["allow"],
            exception_ip_network_lists=var["exception_ip_network_lists"])
        pulumi.export("ipGeoModeBlock", ip_geo_block.mode)
        pulumi.export("blockGeoNetworkLists", ip_geo_block.geo_network_lists)
        pulumi.export("blockIpNetworkLists", ip_geo_block.ip_network_lists)
        pulumi.export("blockExceptionIpNetworkLists", ip_geo_block.exception_ip_network_lists)
        pulumi.export("ipGeoModeAllow", ip_geo_allow.mode)
        pulumi.export("allowExceptionIpNetworkLists", ip_geo_allow.exception_ip_network_lists)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exception_ip_network_lists: The network lists to be allowed regardless of `mode`, `geo_network_lists`, and `ip_network_lists` parameters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] geo_network_lists: The network lists to be blocked or allowed geographically.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_network_lists: The network lists to be blocked or allowd by IP address.
        :param pulumi.Input[str] mode: The mode to use for IP/Geo firewall blocking: `block` to block specific IPs, geographies or network lists, or `allow` to allow specific IPs or geographies to be let through while blocking the rest.
        :param pulumi.Input[str] security_policy_id: The ID of the security policy to use.
        :param pulumi.Input[int] version: The version number of the security configuration to use.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecIPGeoArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use the `AppSecIPGeo` resource to update the method and which network lists to use for IP/Geo firewall blocking.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
        # USE CASE: user wants to update the IP/GEO firewall mode to "block specific IPs/Subnets and Geos" and update the IP list, GEO list & Exception list
        ip_geo_block = akamai.AppSecIPGeo("ipGeoBlock",
            config_id=configuration.config_id,
            version=configuration.latest_version,
            security_policy_id=var["security_policy_id1"],
            mode=var["block"],
            geo_network_lists=var["geo_network_lists"],
            ip_network_lists=var["ip_network_lists"],
            exception_ip_network_lists=var["exception_ip_network_lists"])
        # USE CASE: user wants to update the IP/GEO firewall mode to "block all traffic except IPs/Subnets in block exceptions" and update the Exception list
        ip_geo_allow = akamai.AppSecIPGeo("ipGeoAllow",
            config_id=configuration.config_id,
            version=configuration.latest_version,
            security_policy_id=var["security_policy_id2"],
            mode=var["allow"],
            exception_ip_network_lists=var["exception_ip_network_lists"])
        pulumi.export("ipGeoModeBlock", ip_geo_block.mode)
        pulumi.export("blockGeoNetworkLists", ip_geo_block.geo_network_lists)
        pulumi.export("blockIpNetworkLists", ip_geo_block.ip_network_lists)
        pulumi.export("blockExceptionIpNetworkLists", ip_geo_block.exception_ip_network_lists)
        pulumi.export("ipGeoModeAllow", ip_geo_allow.mode)
        pulumi.export("allowExceptionIpNetworkLists", ip_geo_allow.exception_ip_network_lists)
        ```

        :param str resource_name: The name of the resource.
        :param AppSecIPGeoArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecIPGeoArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 exception_ip_network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 geo_network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 security_policy_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = AppSecIPGeoArgs.__new__(AppSecIPGeoArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            __props__.__dict__["exception_ip_network_lists"] = exception_ip_network_lists
            __props__.__dict__["geo_network_lists"] = geo_network_lists
            __props__.__dict__["ip_network_lists"] = ip_network_lists
            if mode is None and not opts.urn:
                raise TypeError("Missing required property 'mode'")
            __props__.__dict__["mode"] = mode
            if security_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_policy_id'")
            __props__.__dict__["security_policy_id"] = security_policy_id
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__.__dict__["version"] = version
        super(AppSecIPGeo, __self__).__init__(
            'akamai:index/appSecIPGeo:AppSecIPGeo',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            exception_ip_network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            geo_network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ip_network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            security_policy_id: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'AppSecIPGeo':
        """
        Get an existing AppSecIPGeo resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] config_id: The ID of the security configuration to use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exception_ip_network_lists: The network lists to be allowed regardless of `mode`, `geo_network_lists`, and `ip_network_lists` parameters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] geo_network_lists: The network lists to be blocked or allowed geographically.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_network_lists: The network lists to be blocked or allowd by IP address.
        :param pulumi.Input[str] mode: The mode to use for IP/Geo firewall blocking: `block` to block specific IPs, geographies or network lists, or `allow` to allow specific IPs or geographies to be let through while blocking the rest.
        :param pulumi.Input[str] security_policy_id: The ID of the security policy to use.
        :param pulumi.Input[int] version: The version number of the security configuration to use.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecIPGeoState.__new__(_AppSecIPGeoState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["exception_ip_network_lists"] = exception_ip_network_lists
        __props__.__dict__["geo_network_lists"] = geo_network_lists
        __props__.__dict__["ip_network_lists"] = ip_network_lists
        __props__.__dict__["mode"] = mode
        __props__.__dict__["security_policy_id"] = security_policy_id
        __props__.__dict__["version"] = version
        return AppSecIPGeo(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        The ID of the security configuration to use.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="exceptionIpNetworkLists")
    def exception_ip_network_lists(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The network lists to be allowed regardless of `mode`, `geo_network_lists`, and `ip_network_lists` parameters.
        """
        return pulumi.get(self, "exception_ip_network_lists")

    @property
    @pulumi.getter(name="geoNetworkLists")
    def geo_network_lists(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The network lists to be blocked or allowed geographically.
        """
        return pulumi.get(self, "geo_network_lists")

    @property
    @pulumi.getter(name="ipNetworkLists")
    def ip_network_lists(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The network lists to be blocked or allowd by IP address.
        """
        return pulumi.get(self, "ip_network_lists")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        """
        The mode to use for IP/Geo firewall blocking: `block` to block specific IPs, geographies or network lists, or `allow` to allow specific IPs or geographies to be let through while blocking the rest.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter(name="securityPolicyId")
    def security_policy_id(self) -> pulumi.Output[str]:
        """
        The ID of the security policy to use.
        """
        return pulumi.get(self, "security_policy_id")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        The version number of the security configuration to use.
        """
        return pulumi.get(self, "version")

