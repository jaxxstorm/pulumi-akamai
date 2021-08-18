# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetAppSecSelectableHostnamesResult',
    'AwaitableGetAppSecSelectableHostnamesResult',
    'get_app_sec_selectable_hostnames',
]

@pulumi.output_type
class GetAppSecSelectableHostnamesResult:
    """
    A collection of values returned by getAppSecSelectableHostnames.
    """
    def __init__(__self__, active_in_production=None, active_in_staging=None, config_id=None, contractid=None, groupid=None, hostnames=None, hostnames_json=None, id=None, output_text=None):
        if active_in_production and not isinstance(active_in_production, bool):
            raise TypeError("Expected argument 'active_in_production' to be a bool")
        pulumi.set(__self__, "active_in_production", active_in_production)
        if active_in_staging and not isinstance(active_in_staging, bool):
            raise TypeError("Expected argument 'active_in_staging' to be a bool")
        pulumi.set(__self__, "active_in_staging", active_in_staging)
        if config_id and not isinstance(config_id, int):
            raise TypeError("Expected argument 'config_id' to be a int")
        pulumi.set(__self__, "config_id", config_id)
        if contractid and not isinstance(contractid, str):
            raise TypeError("Expected argument 'contractid' to be a str")
        pulumi.set(__self__, "contractid", contractid)
        if groupid and not isinstance(groupid, int):
            raise TypeError("Expected argument 'groupid' to be a int")
        pulumi.set(__self__, "groupid", groupid)
        if hostnames and not isinstance(hostnames, list):
            raise TypeError("Expected argument 'hostnames' to be a list")
        pulumi.set(__self__, "hostnames", hostnames)
        if hostnames_json and not isinstance(hostnames_json, str):
            raise TypeError("Expected argument 'hostnames_json' to be a str")
        pulumi.set(__self__, "hostnames_json", hostnames_json)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_text and not isinstance(output_text, str):
            raise TypeError("Expected argument 'output_text' to be a str")
        pulumi.set(__self__, "output_text", output_text)

    @property
    @pulumi.getter(name="activeInProduction")
    def active_in_production(self) -> Optional[bool]:
        return pulumi.get(self, "active_in_production")

    @property
    @pulumi.getter(name="activeInStaging")
    def active_in_staging(self) -> Optional[bool]:
        return pulumi.get(self, "active_in_staging")

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[int]:
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter
    def contractid(self) -> Optional[str]:
        return pulumi.get(self, "contractid")

    @property
    @pulumi.getter
    def groupid(self) -> Optional[int]:
        return pulumi.get(self, "groupid")

    @property
    @pulumi.getter
    def hostnames(self) -> Sequence[str]:
        """
        The list of selectable hostnames.
        """
        return pulumi.get(self, "hostnames")

    @property
    @pulumi.getter(name="hostnamesJson")
    def hostnames_json(self) -> str:
        """
        The list of selectable hostnames in json format.
        """
        return pulumi.get(self, "hostnames_json")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="outputText")
    def output_text(self) -> str:
        """
        A tabular display of the selectable hostnames showing the name and config_id of the security configuration under which the host is protected in production, or '-' if the host is not protected in production.
        """
        return pulumi.get(self, "output_text")


class AwaitableGetAppSecSelectableHostnamesResult(GetAppSecSelectableHostnamesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppSecSelectableHostnamesResult(
            active_in_production=self.active_in_production,
            active_in_staging=self.active_in_staging,
            config_id=self.config_id,
            contractid=self.contractid,
            groupid=self.groupid,
            hostnames=self.hostnames,
            hostnames_json=self.hostnames_json,
            id=self.id,
            output_text=self.output_text)


def get_app_sec_selectable_hostnames(active_in_production: Optional[bool] = None,
                                     active_in_staging: Optional[bool] = None,
                                     config_id: Optional[int] = None,
                                     contractid: Optional[str] = None,
                                     groupid: Optional[int] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppSecSelectableHostnamesResult:
    """
    Use the `getAppSecSelectableHostnames` data source to retrieve the list of hostnames that may be protected under a given security configuration. You can specify the list to be retrieved either by supplying the name of a security configuration, or by supplying a group ID and contract ID.

    ## Example Usage

    Basic usage:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    configuration = akamai.get_app_sec_configuration(name=var["security_configuration"])
    selectable_hostnames_app_sec_selectable_hostnames = akamai.get_app_sec_selectable_hostnames(config_id=configuration.config_id)
    pulumi.export("selectableHostnames", selectable_hostnames_app_sec_selectable_hostnames.hostnames)
    pulumi.export("selectableHostnamesJson", selectable_hostnames_app_sec_selectable_hostnames.hostnames_json)
    pulumi.export("selectableHostnamesOutputText", selectable_hostnames_app_sec_selectable_hostnames.output_text)
    selectable_hostnames_for_create_configuration_app_sec_selectable_hostnames = akamai.get_app_sec_selectable_hostnames(contractid=var["contractid"],
        groupid=var["groupid"])
    pulumi.export("selectableHostnamesForCreateConfiguration", selectable_hostnames_for_create_configuration_app_sec_selectable_hostnames.hostnames)
    pulumi.export("selectableHostnamesForCreateConfigurationJson", selectable_hostnames_for_create_configuration_app_sec_selectable_hostnames.hostnames_json)
    pulumi.export("selectableHostnamesForCreateConfigurationOutputText", selectable_hostnames_for_create_configuration_app_sec_selectable_hostnames.output_text)
    ```


    :param int config_id: The ID of the security configuration to use.
    :param str contractid: The ID of the contract to use.
    :param int groupid: The ID of the group to use.
    """
    __args__ = dict()
    __args__['activeInProduction'] = active_in_production
    __args__['activeInStaging'] = active_in_staging
    __args__['configId'] = config_id
    __args__['contractid'] = contractid
    __args__['groupid'] = groupid
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:index/getAppSecSelectableHostnames:getAppSecSelectableHostnames', __args__, opts=opts, typ=GetAppSecSelectableHostnamesResult).value

    return AwaitableGetAppSecSelectableHostnamesResult(
        active_in_production=__ret__.active_in_production,
        active_in_staging=__ret__.active_in_staging,
        config_id=__ret__.config_id,
        contractid=__ret__.contractid,
        groupid=__ret__.groupid,
        hostnames=__ret__.hostnames,
        hostnames_json=__ret__.hostnames_json,
        id=__ret__.id,
        output_text=__ret__.output_text)
