# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'PropertyActivationRuleError',
    'PropertyActivationRuleWarning',
    'PropertyHostname',
    'PropertyHostnameCertStatus',
    'PropertyOrigin',
    'PropertyRuleError',
    'PropertyRuleWarning',
]

@pulumi.output_type
class PropertyActivationRuleError(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "behaviorName":
            suggest = "behavior_name"
        elif key == "errorLocation":
            suggest = "error_location"
        elif key == "statusCode":
            suggest = "status_code"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PropertyActivationRuleError. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PropertyActivationRuleError.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PropertyActivationRuleError.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 behavior_name: Optional[str] = None,
                 detail: Optional[str] = None,
                 error_location: Optional[str] = None,
                 instance: Optional[str] = None,
                 status_code: Optional[int] = None,
                 title: Optional[str] = None,
                 type: Optional[str] = None):
        if behavior_name is not None:
            pulumi.set(__self__, "behavior_name", behavior_name)
        if detail is not None:
            pulumi.set(__self__, "detail", detail)
        if error_location is not None:
            pulumi.set(__self__, "error_location", error_location)
        if instance is not None:
            pulumi.set(__self__, "instance", instance)
        if status_code is not None:
            pulumi.set(__self__, "status_code", status_code)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="behaviorName")
    def behavior_name(self) -> Optional[str]:
        return pulumi.get(self, "behavior_name")

    @property
    @pulumi.getter
    def detail(self) -> Optional[str]:
        return pulumi.get(self, "detail")

    @property
    @pulumi.getter(name="errorLocation")
    def error_location(self) -> Optional[str]:
        return pulumi.get(self, "error_location")

    @property
    @pulumi.getter
    def instance(self) -> Optional[str]:
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter(name="statusCode")
    def status_code(self) -> Optional[int]:
        return pulumi.get(self, "status_code")

    @property
    @pulumi.getter
    def title(self) -> Optional[str]:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")


@pulumi.output_type
class PropertyActivationRuleWarning(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "behaviorName":
            suggest = "behavior_name"
        elif key == "errorLocation":
            suggest = "error_location"
        elif key == "statusCode":
            suggest = "status_code"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PropertyActivationRuleWarning. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PropertyActivationRuleWarning.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PropertyActivationRuleWarning.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 behavior_name: Optional[str] = None,
                 detail: Optional[str] = None,
                 error_location: Optional[str] = None,
                 instance: Optional[str] = None,
                 status_code: Optional[int] = None,
                 title: Optional[str] = None,
                 type: Optional[str] = None):
        if behavior_name is not None:
            pulumi.set(__self__, "behavior_name", behavior_name)
        if detail is not None:
            pulumi.set(__self__, "detail", detail)
        if error_location is not None:
            pulumi.set(__self__, "error_location", error_location)
        if instance is not None:
            pulumi.set(__self__, "instance", instance)
        if status_code is not None:
            pulumi.set(__self__, "status_code", status_code)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="behaviorName")
    def behavior_name(self) -> Optional[str]:
        return pulumi.get(self, "behavior_name")

    @property
    @pulumi.getter
    def detail(self) -> Optional[str]:
        return pulumi.get(self, "detail")

    @property
    @pulumi.getter(name="errorLocation")
    def error_location(self) -> Optional[str]:
        return pulumi.get(self, "error_location")

    @property
    @pulumi.getter
    def instance(self) -> Optional[str]:
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter(name="statusCode")
    def status_code(self) -> Optional[int]:
        return pulumi.get(self, "status_code")

    @property
    @pulumi.getter
    def title(self) -> Optional[str]:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")


@pulumi.output_type
class PropertyHostname(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "certProvisioningType":
            suggest = "cert_provisioning_type"
        elif key == "cnameFrom":
            suggest = "cname_from"
        elif key == "cnameTo":
            suggest = "cname_to"
        elif key == "certStatuses":
            suggest = "cert_statuses"
        elif key == "cnameType":
            suggest = "cname_type"
        elif key == "edgeHostnameId":
            suggest = "edge_hostname_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PropertyHostname. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PropertyHostname.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PropertyHostname.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cert_provisioning_type: str,
                 cname_from: str,
                 cname_to: str,
                 cert_statuses: Optional[Sequence['outputs.PropertyHostnameCertStatus']] = None,
                 cname_type: Optional[str] = None,
                 edge_hostname_id: Optional[str] = None):
        """
        :param str cert_provisioning_type: The certificate's provisioning type, either the default `CPS_MANAGED` type for the custom certificates you provision with the [Certificate Provisioning System (CPS)](https://techdocs.akamai.com/cps/docs), or `DEFAULT` for certificates provisioned automatically.
        :param str cname_from: A string containing the original origin's hostname. For example, `"example.org"`.
        :param str cname_to: A string containing the hostname for edge content. For example,  `"example.org.edgesuite.net"`.
        """
        pulumi.set(__self__, "cert_provisioning_type", cert_provisioning_type)
        pulumi.set(__self__, "cname_from", cname_from)
        pulumi.set(__self__, "cname_to", cname_to)
        if cert_statuses is not None:
            pulumi.set(__self__, "cert_statuses", cert_statuses)
        if cname_type is not None:
            pulumi.set(__self__, "cname_type", cname_type)
        if edge_hostname_id is not None:
            pulumi.set(__self__, "edge_hostname_id", edge_hostname_id)

    @property
    @pulumi.getter(name="certProvisioningType")
    def cert_provisioning_type(self) -> str:
        """
        The certificate's provisioning type, either the default `CPS_MANAGED` type for the custom certificates you provision with the [Certificate Provisioning System (CPS)](https://techdocs.akamai.com/cps/docs), or `DEFAULT` for certificates provisioned automatically.
        """
        return pulumi.get(self, "cert_provisioning_type")

    @property
    @pulumi.getter(name="cnameFrom")
    def cname_from(self) -> str:
        """
        A string containing the original origin's hostname. For example, `"example.org"`.
        """
        return pulumi.get(self, "cname_from")

    @property
    @pulumi.getter(name="cnameTo")
    def cname_to(self) -> str:
        """
        A string containing the hostname for edge content. For example,  `"example.org.edgesuite.net"`.
        """
        return pulumi.get(self, "cname_to")

    @property
    @pulumi.getter(name="certStatuses")
    def cert_statuses(self) -> Optional[Sequence['outputs.PropertyHostnameCertStatus']]:
        return pulumi.get(self, "cert_statuses")

    @property
    @pulumi.getter(name="cnameType")
    def cname_type(self) -> Optional[str]:
        return pulumi.get(self, "cname_type")

    @property
    @pulumi.getter(name="edgeHostnameId")
    def edge_hostname_id(self) -> Optional[str]:
        return pulumi.get(self, "edge_hostname_id")


@pulumi.output_type
class PropertyHostnameCertStatus(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "productionStatus":
            suggest = "production_status"
        elif key == "stagingStatus":
            suggest = "staging_status"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PropertyHostnameCertStatus. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PropertyHostnameCertStatus.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PropertyHostnameCertStatus.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 hostname: Optional[str] = None,
                 production_status: Optional[str] = None,
                 staging_status: Optional[str] = None,
                 target: Optional[str] = None):
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if production_status is not None:
            pulumi.set(__self__, "production_status", production_status)
        if staging_status is not None:
            pulumi.set(__self__, "staging_status", staging_status)
        if target is not None:
            pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[str]:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="productionStatus")
    def production_status(self) -> Optional[str]:
        return pulumi.get(self, "production_status")

    @property
    @pulumi.getter(name="stagingStatus")
    def staging_status(self) -> Optional[str]:
        return pulumi.get(self, "staging_status")

    @property
    @pulumi.getter
    def target(self) -> Optional[str]:
        return pulumi.get(self, "target")


@pulumi.output_type
class PropertyOrigin(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "cacheKeyHostname":
            suggest = "cache_key_hostname"
        elif key == "enableTrueClientIp":
            suggest = "enable_true_client_ip"
        elif key == "forwardHostname":
            suggest = "forward_hostname"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PropertyOrigin. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PropertyOrigin.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PropertyOrigin.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cache_key_hostname: Optional[str] = None,
                 compress: Optional[bool] = None,
                 enable_true_client_ip: Optional[bool] = None,
                 forward_hostname: Optional[str] = None,
                 hostname: Optional[str] = None,
                 port: Optional[int] = None):
        if cache_key_hostname is not None:
            pulumi.set(__self__, "cache_key_hostname", cache_key_hostname)
        if compress is not None:
            pulumi.set(__self__, "compress", compress)
        if enable_true_client_ip is not None:
            pulumi.set(__self__, "enable_true_client_ip", enable_true_client_ip)
        if forward_hostname is not None:
            pulumi.set(__self__, "forward_hostname", forward_hostname)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="cacheKeyHostname")
    def cache_key_hostname(self) -> Optional[str]:
        return pulumi.get(self, "cache_key_hostname")

    @property
    @pulumi.getter
    def compress(self) -> Optional[bool]:
        return pulumi.get(self, "compress")

    @property
    @pulumi.getter(name="enableTrueClientIp")
    def enable_true_client_ip(self) -> Optional[bool]:
        return pulumi.get(self, "enable_true_client_ip")

    @property
    @pulumi.getter(name="forwardHostname")
    def forward_hostname(self) -> Optional[str]:
        return pulumi.get(self, "forward_hostname")

    @property
    @pulumi.getter
    def hostname(self) -> Optional[str]:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        return pulumi.get(self, "port")


@pulumi.output_type
class PropertyRuleError(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "behaviorName":
            suggest = "behavior_name"
        elif key == "errorLocation":
            suggest = "error_location"
        elif key == "statusCode":
            suggest = "status_code"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PropertyRuleError. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PropertyRuleError.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PropertyRuleError.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 behavior_name: Optional[str] = None,
                 detail: Optional[str] = None,
                 error_location: Optional[str] = None,
                 instance: Optional[str] = None,
                 status_code: Optional[int] = None,
                 title: Optional[str] = None,
                 type: Optional[str] = None):
        if behavior_name is not None:
            pulumi.set(__self__, "behavior_name", behavior_name)
        if detail is not None:
            pulumi.set(__self__, "detail", detail)
        if error_location is not None:
            pulumi.set(__self__, "error_location", error_location)
        if instance is not None:
            pulumi.set(__self__, "instance", instance)
        if status_code is not None:
            pulumi.set(__self__, "status_code", status_code)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="behaviorName")
    def behavior_name(self) -> Optional[str]:
        return pulumi.get(self, "behavior_name")

    @property
    @pulumi.getter
    def detail(self) -> Optional[str]:
        return pulumi.get(self, "detail")

    @property
    @pulumi.getter(name="errorLocation")
    def error_location(self) -> Optional[str]:
        return pulumi.get(self, "error_location")

    @property
    @pulumi.getter
    def instance(self) -> Optional[str]:
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter(name="statusCode")
    def status_code(self) -> Optional[int]:
        return pulumi.get(self, "status_code")

    @property
    @pulumi.getter
    def title(self) -> Optional[str]:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")


@pulumi.output_type
class PropertyRuleWarning(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "behaviorName":
            suggest = "behavior_name"
        elif key == "errorLocation":
            suggest = "error_location"
        elif key == "statusCode":
            suggest = "status_code"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PropertyRuleWarning. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PropertyRuleWarning.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PropertyRuleWarning.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 behavior_name: Optional[str] = None,
                 detail: Optional[str] = None,
                 error_location: Optional[str] = None,
                 instance: Optional[str] = None,
                 status_code: Optional[int] = None,
                 title: Optional[str] = None,
                 type: Optional[str] = None):
        if behavior_name is not None:
            pulumi.set(__self__, "behavior_name", behavior_name)
        if detail is not None:
            pulumi.set(__self__, "detail", detail)
        if error_location is not None:
            pulumi.set(__self__, "error_location", error_location)
        if instance is not None:
            pulumi.set(__self__, "instance", instance)
        if status_code is not None:
            pulumi.set(__self__, "status_code", status_code)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="behaviorName")
    def behavior_name(self) -> Optional[str]:
        return pulumi.get(self, "behavior_name")

    @property
    @pulumi.getter
    def detail(self) -> Optional[str]:
        return pulumi.get(self, "detail")

    @property
    @pulumi.getter(name="errorLocation")
    def error_location(self) -> Optional[str]:
        return pulumi.get(self, "error_location")

    @property
    @pulumi.getter
    def instance(self) -> Optional[str]:
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter(name="statusCode")
    def status_code(self) -> Optional[int]:
        return pulumi.get(self, "status_code")

    @property
    @pulumi.getter
    def title(self) -> Optional[str]:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")


