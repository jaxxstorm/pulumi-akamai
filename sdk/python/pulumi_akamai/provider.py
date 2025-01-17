# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._inputs import *

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 appsec_section: Optional[pulumi.Input[str]] = None,
                 appsecs: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderAppsecArgs']]]] = None,
                 cache_enabled: Optional[pulumi.Input[bool]] = None,
                 config: Optional[pulumi.Input['ProviderConfigArgs']] = None,
                 config_section: Optional[pulumi.Input[str]] = None,
                 dns: Optional[pulumi.Input['ProviderDnsArgs']] = None,
                 dns_section: Optional[pulumi.Input[str]] = None,
                 edgerc: Optional[pulumi.Input[str]] = None,
                 gtm: Optional[pulumi.Input['ProviderGtmArgs']] = None,
                 gtm_section: Optional[pulumi.Input[str]] = None,
                 networklist_section: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderNetworkArgs']]]] = None,
                 papi_section: Optional[pulumi.Input[str]] = None,
                 property: Optional[pulumi.Input['ProviderPropertyArgs']] = None,
                 property_section: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] config_section: The section of the edgerc file to use for configuration
        """
        if appsec_section is not None:
            warnings.warn("""The setting \"appsec_section\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""appsec_section is deprecated: The setting \"appsec_section\" has been deprecated.""")
        if appsec_section is not None:
            pulumi.set(__self__, "appsec_section", appsec_section)
        if appsecs is not None:
            warnings.warn("""The setting \"appsec\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""appsecs is deprecated: The setting \"appsec\" has been deprecated.""")
        if appsecs is not None:
            pulumi.set(__self__, "appsecs", appsecs)
        if cache_enabled is not None:
            pulumi.set(__self__, "cache_enabled", cache_enabled)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if config_section is not None:
            pulumi.set(__self__, "config_section", config_section)
        if dns is not None:
            warnings.warn("""The setting \"dns\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""dns is deprecated: The setting \"dns\" has been deprecated.""")
        if dns is not None:
            pulumi.set(__self__, "dns", dns)
        if dns_section is not None:
            warnings.warn("""The setting \"dns_section\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""dns_section is deprecated: The setting \"dns_section\" has been deprecated.""")
        if dns_section is not None:
            pulumi.set(__self__, "dns_section", dns_section)
        if edgerc is not None:
            pulumi.set(__self__, "edgerc", edgerc)
        if gtm is not None:
            warnings.warn("""The setting \"gtm\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""gtm is deprecated: The setting \"gtm\" has been deprecated.""")
        if gtm is not None:
            pulumi.set(__self__, "gtm", gtm)
        if gtm_section is not None:
            warnings.warn("""The setting \"gtm_section\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""gtm_section is deprecated: The setting \"gtm_section\" has been deprecated.""")
        if gtm_section is not None:
            pulumi.set(__self__, "gtm_section", gtm_section)
        if networklist_section is not None:
            warnings.warn("""The setting \"networklist_section\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""networklist_section is deprecated: The setting \"networklist_section\" has been deprecated.""")
        if networklist_section is not None:
            pulumi.set(__self__, "networklist_section", networklist_section)
        if networks is not None:
            pulumi.set(__self__, "networks", networks)
        if papi_section is not None:
            warnings.warn("""The setting \"papi_section\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""papi_section is deprecated: The setting \"papi_section\" has been deprecated.""")
        if papi_section is not None:
            pulumi.set(__self__, "papi_section", papi_section)
        if property is not None:
            warnings.warn("""The setting \"property\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""property is deprecated: The setting \"property\" has been deprecated.""")
        if property is not None:
            pulumi.set(__self__, "property", property)
        if property_section is not None:
            warnings.warn("""The setting \"property_section\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""property_section is deprecated: The setting \"property_section\" has been deprecated.""")
        if property_section is not None:
            pulumi.set(__self__, "property_section", property_section)

    @property
    @pulumi.getter(name="appsecSection")
    def appsec_section(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "appsec_section")

    @appsec_section.setter
    def appsec_section(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "appsec_section", value)

    @property
    @pulumi.getter
    def appsecs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProviderAppsecArgs']]]]:
        return pulumi.get(self, "appsecs")

    @appsecs.setter
    def appsecs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderAppsecArgs']]]]):
        pulumi.set(self, "appsecs", value)

    @property
    @pulumi.getter(name="cacheEnabled")
    def cache_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "cache_enabled")

    @cache_enabled.setter
    def cache_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cache_enabled", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input['ProviderConfigArgs']]:
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input['ProviderConfigArgs']]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="configSection")
    def config_section(self) -> Optional[pulumi.Input[str]]:
        """
        The section of the edgerc file to use for configuration
        """
        return pulumi.get(self, "config_section")

    @config_section.setter
    def config_section(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_section", value)

    @property
    @pulumi.getter
    def dns(self) -> Optional[pulumi.Input['ProviderDnsArgs']]:
        return pulumi.get(self, "dns")

    @dns.setter
    def dns(self, value: Optional[pulumi.Input['ProviderDnsArgs']]):
        pulumi.set(self, "dns", value)

    @property
    @pulumi.getter(name="dnsSection")
    def dns_section(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dns_section")

    @dns_section.setter
    def dns_section(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_section", value)

    @property
    @pulumi.getter
    def edgerc(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "edgerc")

    @edgerc.setter
    def edgerc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edgerc", value)

    @property
    @pulumi.getter
    def gtm(self) -> Optional[pulumi.Input['ProviderGtmArgs']]:
        return pulumi.get(self, "gtm")

    @gtm.setter
    def gtm(self, value: Optional[pulumi.Input['ProviderGtmArgs']]):
        pulumi.set(self, "gtm", value)

    @property
    @pulumi.getter(name="gtmSection")
    def gtm_section(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "gtm_section")

    @gtm_section.setter
    def gtm_section(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gtm_section", value)

    @property
    @pulumi.getter(name="networklistSection")
    def networklist_section(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "networklist_section")

    @networklist_section.setter
    def networklist_section(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "networklist_section", value)

    @property
    @pulumi.getter
    def networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProviderNetworkArgs']]]]:
        return pulumi.get(self, "networks")

    @networks.setter
    def networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderNetworkArgs']]]]):
        pulumi.set(self, "networks", value)

    @property
    @pulumi.getter(name="papiSection")
    def papi_section(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "papi_section")

    @papi_section.setter
    def papi_section(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "papi_section", value)

    @property
    @pulumi.getter(name="propertySection")
    def property_section(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "property_section")

    @property_section.setter
    def property_section(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "property_section", value)

    @property
    @pulumi.getter
    def property(self) -> Optional[pulumi.Input['ProviderPropertyArgs']]:
        return pulumi.get(self, "property")

    @property.setter
    def property(self, value: Optional[pulumi.Input['ProviderPropertyArgs']]):
        pulumi.set(self, "property", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 appsec_section: Optional[pulumi.Input[str]] = None,
                 appsecs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProviderAppsecArgs']]]]] = None,
                 cache_enabled: Optional[pulumi.Input[bool]] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['ProviderConfigArgs']]] = None,
                 config_section: Optional[pulumi.Input[str]] = None,
                 dns: Optional[pulumi.Input[pulumi.InputType['ProviderDnsArgs']]] = None,
                 dns_section: Optional[pulumi.Input[str]] = None,
                 edgerc: Optional[pulumi.Input[str]] = None,
                 gtm: Optional[pulumi.Input[pulumi.InputType['ProviderGtmArgs']]] = None,
                 gtm_section: Optional[pulumi.Input[str]] = None,
                 networklist_section: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProviderNetworkArgs']]]]] = None,
                 papi_section: Optional[pulumi.Input[str]] = None,
                 property: Optional[pulumi.Input[pulumi.InputType['ProviderPropertyArgs']]] = None,
                 property_section: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the akamai package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_section: The section of the edgerc file to use for configuration
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the akamai package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 appsec_section: Optional[pulumi.Input[str]] = None,
                 appsecs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProviderAppsecArgs']]]]] = None,
                 cache_enabled: Optional[pulumi.Input[bool]] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['ProviderConfigArgs']]] = None,
                 config_section: Optional[pulumi.Input[str]] = None,
                 dns: Optional[pulumi.Input[pulumi.InputType['ProviderDnsArgs']]] = None,
                 dns_section: Optional[pulumi.Input[str]] = None,
                 edgerc: Optional[pulumi.Input[str]] = None,
                 gtm: Optional[pulumi.Input[pulumi.InputType['ProviderGtmArgs']]] = None,
                 gtm_section: Optional[pulumi.Input[str]] = None,
                 networklist_section: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProviderNetworkArgs']]]]] = None,
                 papi_section: Optional[pulumi.Input[str]] = None,
                 property: Optional[pulumi.Input[pulumi.InputType['ProviderPropertyArgs']]] = None,
                 property_section: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if appsec_section is not None and not opts.urn:
                warnings.warn("""The setting \"appsec_section\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("""appsec_section is deprecated: The setting \"appsec_section\" has been deprecated.""")
            __props__.__dict__["appsec_section"] = appsec_section
            if appsecs is not None and not opts.urn:
                warnings.warn("""The setting \"appsec\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("""appsecs is deprecated: The setting \"appsec\" has been deprecated.""")
            __props__.__dict__["appsecs"] = pulumi.Output.from_input(appsecs).apply(pulumi.runtime.to_json) if appsecs is not None else None
            __props__.__dict__["cache_enabled"] = pulumi.Output.from_input(cache_enabled).apply(pulumi.runtime.to_json) if cache_enabled is not None else None
            __props__.__dict__["config"] = pulumi.Output.from_input(config).apply(pulumi.runtime.to_json) if config is not None else None
            __props__.__dict__["config_section"] = config_section
            if dns is not None and not opts.urn:
                warnings.warn("""The setting \"dns\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("""dns is deprecated: The setting \"dns\" has been deprecated.""")
            __props__.__dict__["dns"] = pulumi.Output.from_input(dns).apply(pulumi.runtime.to_json) if dns is not None else None
            if dns_section is not None and not opts.urn:
                warnings.warn("""The setting \"dns_section\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("""dns_section is deprecated: The setting \"dns_section\" has been deprecated.""")
            __props__.__dict__["dns_section"] = dns_section
            __props__.__dict__["edgerc"] = edgerc
            if gtm is not None and not opts.urn:
                warnings.warn("""The setting \"gtm\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("""gtm is deprecated: The setting \"gtm\" has been deprecated.""")
            __props__.__dict__["gtm"] = pulumi.Output.from_input(gtm).apply(pulumi.runtime.to_json) if gtm is not None else None
            if gtm_section is not None and not opts.urn:
                warnings.warn("""The setting \"gtm_section\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("""gtm_section is deprecated: The setting \"gtm_section\" has been deprecated.""")
            __props__.__dict__["gtm_section"] = gtm_section
            if networklist_section is not None and not opts.urn:
                warnings.warn("""The setting \"networklist_section\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("""networklist_section is deprecated: The setting \"networklist_section\" has been deprecated.""")
            __props__.__dict__["networklist_section"] = networklist_section
            __props__.__dict__["networks"] = pulumi.Output.from_input(networks).apply(pulumi.runtime.to_json) if networks is not None else None
            if papi_section is not None and not opts.urn:
                warnings.warn("""The setting \"papi_section\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("""papi_section is deprecated: The setting \"papi_section\" has been deprecated.""")
            __props__.__dict__["papi_section"] = papi_section
            if property is not None and not opts.urn:
                warnings.warn("""The setting \"property\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("""property is deprecated: The setting \"property\" has been deprecated.""")
            __props__.__dict__["property"] = pulumi.Output.from_input(property).apply(pulumi.runtime.to_json) if property is not None else None
            if property_section is not None and not opts.urn:
                warnings.warn("""The setting \"property_section\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("""property_section is deprecated: The setting \"property_section\" has been deprecated.""")
            __props__.__dict__["property_section"] = property_section
        super(Provider, __self__).__init__(
            'akamai',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="appsecSection")
    def appsec_section(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "appsec_section")

    @property
    @pulumi.getter(name="configSection")
    def config_section(self) -> pulumi.Output[Optional[str]]:
        """
        The section of the edgerc file to use for configuration
        """
        return pulumi.get(self, "config_section")

    @property
    @pulumi.getter(name="dnsSection")
    def dns_section(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dns_section")

    @property
    @pulumi.getter
    def edgerc(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "edgerc")

    @property
    @pulumi.getter(name="gtmSection")
    def gtm_section(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "gtm_section")

    @property
    @pulumi.getter(name="networklistSection")
    def networklist_section(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "networklist_section")

    @property
    @pulumi.getter(name="papiSection")
    def papi_section(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "papi_section")

    @property
    @pulumi.getter(name="propertySection")
    def property_section(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "property_section")

