# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .app_sec_activations import *
from .app_sec_configuration_version_clone import *
from .app_sec_custom_rule import *
from .app_sec_custom_rule_action import *
from .app_sec_match_target import *
from .app_sec_match_target_sequence import *
from .app_sec_security_policy_clone import *
from .app_sec_selected_hostnames import *
from .cp_code import *
from .dns_record import *
from .dns_zone import *
from .edge_host_name import *
from .get_app_sec_configuration import *
from .get_app_sec_configuration_version import *
from .get_app_sec_custom_rule_actions import *
from .get_app_sec_custom_rules import *
from .get_app_sec_export_configuration import *
from .get_app_sec_match_targets import *
from .get_app_sec_security_policy import *
from .get_app_sec_selectable_hostnames import *
from .get_app_sec_selected_hostnames import *
from .get_authorities_set import *
from .get_contract import *
from .get_contracts import *
from .get_cp_code import *
from .get_dns_record_set import *
from .get_group import *
from .get_groups import *
from .get_gtm_default_datacenter import *
from .get_properties import *
from .get_property import *
from .get_property_products import *
from .get_property_rule_formats import *
from .get_property_rules import *
from .get_property_rules_template import *
from .gtm_asmap import *
from .gtm_cidrmap import *
from .gtm_datacenter import *
from .gtm_domain import *
from .gtm_geomap import *
from .gtm_property import *
from .gtm_resource import *
from .property import *
from .property_activation import *
from .property_variables import *
from .provider import *
from ._inputs import *
from . import outputs

# Make subpackages available:
from . import (
    config,
    edgedns,
    properties,
    trafficmanagement,
)

def _register_module():
    import pulumi
    from . import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "akamai:index/appSecActivations:AppSecActivations":
                return AppSecActivations(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/appSecConfigurationVersionClone:AppSecConfigurationVersionClone":
                return AppSecConfigurationVersionClone(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/appSecCustomRule:AppSecCustomRule":
                return AppSecCustomRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/appSecCustomRuleAction:AppSecCustomRuleAction":
                return AppSecCustomRuleAction(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/appSecMatchTarget:AppSecMatchTarget":
                return AppSecMatchTarget(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/appSecMatchTargetSequence:AppSecMatchTargetSequence":
                return AppSecMatchTargetSequence(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/appSecSecurityPolicyClone:AppSecSecurityPolicyClone":
                return AppSecSecurityPolicyClone(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/appSecSelectedHostnames:AppSecSelectedHostnames":
                return AppSecSelectedHostnames(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/cpCode:CpCode":
                return CpCode(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/dnsRecord:DnsRecord":
                return DnsRecord(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/dnsZone:DnsZone":
                return DnsZone(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/edgeHostName:EdgeHostName":
                return EdgeHostName(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/gtmAsmap:GtmAsmap":
                return GtmAsmap(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/gtmCidrmap:GtmCidrmap":
                return GtmCidrmap(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/gtmDatacenter:GtmDatacenter":
                return GtmDatacenter(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/gtmDomain:GtmDomain":
                return GtmDomain(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/gtmGeomap:GtmGeomap":
                return GtmGeomap(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/gtmProperty:GtmProperty":
                return GtmProperty(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/gtmResource:GtmResource":
                return GtmResource(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/property:Property":
                return Property(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/propertyActivation:PropertyActivation":
                return PropertyActivation(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "akamai:index/propertyVariables:PropertyVariables":
                return PropertyVariables(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("akamai", "index/appSecActivations", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/appSecConfigurationVersionClone", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/appSecCustomRule", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/appSecCustomRuleAction", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/appSecMatchTarget", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/appSecMatchTargetSequence", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/appSecSecurityPolicyClone", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/appSecSelectedHostnames", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/cpCode", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/dnsRecord", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/dnsZone", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/edgeHostName", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/gtmAsmap", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/gtmCidrmap", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/gtmDatacenter", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/gtmDomain", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/gtmGeomap", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/gtmProperty", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/gtmResource", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/property", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/propertyActivation", _module_instance)
    pulumi.runtime.register_resource_module("akamai", "index/propertyVariables", _module_instance)


    class Package(pulumi.runtime.ResourcePackage):
        _version = _utilities.get_semver_version()

        def version(self):
            return Package._version

        def construct_provider(self, name: str, typ: str, urn: str) -> pulumi.ProviderResource:
            if typ != "pulumi:providers:akamai":
                raise Exception(f"unknown provider type {typ}")
            return Provider(name, pulumi.ResourceOptions(urn=urn))


    pulumi.runtime.register_resource_package("akamai", Package())

_register_module()
