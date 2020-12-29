# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from ._inputs import *

__all__ = ['Provider']


class Provider(pulumi.ProviderResource):
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
                 papi_section: Optional[pulumi.Input[str]] = None,
                 property: Optional[pulumi.Input[pulumi.InputType['ProviderPropertyArgs']]] = None,
                 property_section: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The provider type for the akamai package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_section: The section of the edgerc file to use for configuration
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

            if appsec_section is not None and not opts.urn:
                warnings.warn("""The setting \"appsec_section\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("appsec_section is deprecated: The setting \"appsec_section\" has been deprecated.")
            __props__['appsec_section'] = appsec_section
            if appsecs is not None and not opts.urn:
                warnings.warn("""The setting \"appsec\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("appsecs is deprecated: The setting \"appsec\" has been deprecated.")
            __props__['appsecs'] = pulumi.Output.from_input(appsecs).apply(pulumi.runtime.to_json) if appsecs is not None else None
            __props__['cache_enabled'] = pulumi.Output.from_input(cache_enabled).apply(pulumi.runtime.to_json) if cache_enabled is not None else None
            __props__['config'] = pulumi.Output.from_input(config).apply(pulumi.runtime.to_json) if config is not None else None
            __props__['config_section'] = config_section
            if dns is not None and not opts.urn:
                warnings.warn("""The setting \"dns\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("dns is deprecated: The setting \"dns\" has been deprecated.")
            __props__['dns'] = pulumi.Output.from_input(dns).apply(pulumi.runtime.to_json) if dns is not None else None
            if dns_section is not None and not opts.urn:
                warnings.warn("""The setting \"dns_section\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("dns_section is deprecated: The setting \"dns_section\" has been deprecated.")
            __props__['dns_section'] = dns_section
            __props__['edgerc'] = edgerc
            if gtm is not None and not opts.urn:
                warnings.warn("""The setting \"gtm\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("gtm is deprecated: The setting \"gtm\" has been deprecated.")
            __props__['gtm'] = pulumi.Output.from_input(gtm).apply(pulumi.runtime.to_json) if gtm is not None else None
            if gtm_section is not None and not opts.urn:
                warnings.warn("""The setting \"gtm_section\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("gtm_section is deprecated: The setting \"gtm_section\" has been deprecated.")
            __props__['gtm_section'] = gtm_section
            if papi_section is not None and not opts.urn:
                warnings.warn("""The setting \"papi_section\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("papi_section is deprecated: The setting \"papi_section\" has been deprecated.")
            __props__['papi_section'] = papi_section
            if property is not None and not opts.urn:
                warnings.warn("""The setting \"property\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("property is deprecated: The setting \"property\" has been deprecated.")
            __props__['property'] = pulumi.Output.from_input(property).apply(pulumi.runtime.to_json) if property is not None else None
            if property_section is not None and not opts.urn:
                warnings.warn("""The setting \"property_section\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("property_section is deprecated: The setting \"property_section\" has been deprecated.")
            __props__['property_section'] = property_section
        super(Provider, __self__).__init__(
            'akamai',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

