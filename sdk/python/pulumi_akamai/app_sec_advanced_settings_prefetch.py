# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecAdvancedSettingsPrefetchArgs', 'AppSecAdvancedSettingsPrefetch']

@pulumi.input_type
class AppSecAdvancedSettingsPrefetchArgs:
    def __init__(__self__, *,
                 all_extensions: pulumi.Input[bool],
                 config_id: pulumi.Input[int],
                 enable_app_layer: pulumi.Input[bool],
                 enable_rate_controls: pulumi.Input[bool],
                 extensions: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a AppSecAdvancedSettingsPrefetch resource.
        :param pulumi.Input[bool] all_extensions: . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the prefetch settings being modified.
        :param pulumi.Input[bool] enable_app_layer: . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
        :param pulumi.Input[bool] enable_rate_controls: . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] extensions: . If `all_extensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won't be used with any file extensions not included in the array. If `all_extensions` is **true**, then this argument must be set to an empty array: **[]**.
        """
        pulumi.set(__self__, "all_extensions", all_extensions)
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "enable_app_layer", enable_app_layer)
        pulumi.set(__self__, "enable_rate_controls", enable_rate_controls)
        pulumi.set(__self__, "extensions", extensions)

    @property
    @pulumi.getter(name="allExtensions")
    def all_extensions(self) -> pulumi.Input[bool]:
        """
        . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
        """
        return pulumi.get(self, "all_extensions")

    @all_extensions.setter
    def all_extensions(self, value: pulumi.Input[bool]):
        pulumi.set(self, "all_extensions", value)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        . Unique identifier of the security configuration associated with the prefetch settings being modified.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="enableAppLayer")
    def enable_app_layer(self) -> pulumi.Input[bool]:
        """
        . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
        """
        return pulumi.get(self, "enable_app_layer")

    @enable_app_layer.setter
    def enable_app_layer(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enable_app_layer", value)

    @property
    @pulumi.getter(name="enableRateControls")
    def enable_rate_controls(self) -> pulumi.Input[bool]:
        """
        . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
        """
        return pulumi.get(self, "enable_rate_controls")

    @enable_rate_controls.setter
    def enable_rate_controls(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enable_rate_controls", value)

    @property
    @pulumi.getter
    def extensions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        . If `all_extensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won't be used with any file extensions not included in the array. If `all_extensions` is **true**, then this argument must be set to an empty array: **[]**.
        """
        return pulumi.get(self, "extensions")

    @extensions.setter
    def extensions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "extensions", value)


@pulumi.input_type
class _AppSecAdvancedSettingsPrefetchState:
    def __init__(__self__, *,
                 all_extensions: Optional[pulumi.Input[bool]] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 enable_app_layer: Optional[pulumi.Input[bool]] = None,
                 enable_rate_controls: Optional[pulumi.Input[bool]] = None,
                 extensions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering AppSecAdvancedSettingsPrefetch resources.
        :param pulumi.Input[bool] all_extensions: . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the prefetch settings being modified.
        :param pulumi.Input[bool] enable_app_layer: . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
        :param pulumi.Input[bool] enable_rate_controls: . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] extensions: . If `all_extensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won't be used with any file extensions not included in the array. If `all_extensions` is **true**, then this argument must be set to an empty array: **[]**.
        """
        if all_extensions is not None:
            pulumi.set(__self__, "all_extensions", all_extensions)
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if enable_app_layer is not None:
            pulumi.set(__self__, "enable_app_layer", enable_app_layer)
        if enable_rate_controls is not None:
            pulumi.set(__self__, "enable_rate_controls", enable_rate_controls)
        if extensions is not None:
            pulumi.set(__self__, "extensions", extensions)

    @property
    @pulumi.getter(name="allExtensions")
    def all_extensions(self) -> Optional[pulumi.Input[bool]]:
        """
        . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
        """
        return pulumi.get(self, "all_extensions")

    @all_extensions.setter
    def all_extensions(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "all_extensions", value)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        . Unique identifier of the security configuration associated with the prefetch settings being modified.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="enableAppLayer")
    def enable_app_layer(self) -> Optional[pulumi.Input[bool]]:
        """
        . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
        """
        return pulumi.get(self, "enable_app_layer")

    @enable_app_layer.setter
    def enable_app_layer(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_app_layer", value)

    @property
    @pulumi.getter(name="enableRateControls")
    def enable_rate_controls(self) -> Optional[pulumi.Input[bool]]:
        """
        . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
        """
        return pulumi.get(self, "enable_rate_controls")

    @enable_rate_controls.setter
    def enable_rate_controls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_rate_controls", value)

    @property
    @pulumi.getter
    def extensions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        . If `all_extensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won't be used with any file extensions not included in the array. If `all_extensions` is **true**, then this argument must be set to an empty array: **[]**.
        """
        return pulumi.get(self, "extensions")

    @extensions.setter
    def extensions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "extensions", value)


class AppSecAdvancedSettingsPrefetch(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 all_extensions: Optional[pulumi.Input[bool]] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 enable_app_layer: Optional[pulumi.Input[bool]] = None,
                 enable_rate_controls: Optional[pulumi.Input[bool]] = None,
                 extensions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        **Scopes**: Security configuration

        Enables inspection of internal requests (that is, requests between your origin servers and Akamai's edge servers).
        You can also use this resource to apply rate controls to prefetch requests.
        When prefetch is enabled, internal requests are inspected by your firewall the same way that external requests (requests that originate outside the firewall and outside Akamai's edge servers) are inspected.

        This operation applies at the security configuration level, meaning that the settings affect all the security policies in that configuration.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/advanced-settings/prefetch](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putprefetchrequestsforaconfiguration)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        prefetch = akamai.AppSecAdvancedSettingsPrefetch("prefetch",
            config_id=configuration.config_id,
            enable_app_layer=False,
            all_extensions=True,
            enable_rate_controls=False,
            extensions=[
                ".tiff",
                ".bmp",
                ".jpg",
                ".gif",
                ".png",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] all_extensions: . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the prefetch settings being modified.
        :param pulumi.Input[bool] enable_app_layer: . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
        :param pulumi.Input[bool] enable_rate_controls: . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] extensions: . If `all_extensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won't be used with any file extensions not included in the array. If `all_extensions` is **true**, then this argument must be set to an empty array: **[]**.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecAdvancedSettingsPrefetchArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        **Scopes**: Security configuration

        Enables inspection of internal requests (that is, requests between your origin servers and Akamai's edge servers).
        You can also use this resource to apply rate controls to prefetch requests.
        When prefetch is enabled, internal requests are inspected by your firewall the same way that external requests (requests that originate outside the firewall and outside Akamai's edge servers) are inspected.

        This operation applies at the security configuration level, meaning that the settings affect all the security policies in that configuration.

        **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/advanced-settings/prefetch](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putprefetchrequestsforaconfiguration)

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        configuration = akamai.get_app_sec_configuration(name="Documentation")
        prefetch = akamai.AppSecAdvancedSettingsPrefetch("prefetch",
            config_id=configuration.config_id,
            enable_app_layer=False,
            all_extensions=True,
            enable_rate_controls=False,
            extensions=[
                ".tiff",
                ".bmp",
                ".jpg",
                ".gif",
                ".png",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param AppSecAdvancedSettingsPrefetchArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecAdvancedSettingsPrefetchArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 all_extensions: Optional[pulumi.Input[bool]] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 enable_app_layer: Optional[pulumi.Input[bool]] = None,
                 enable_rate_controls: Optional[pulumi.Input[bool]] = None,
                 extensions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = AppSecAdvancedSettingsPrefetchArgs.__new__(AppSecAdvancedSettingsPrefetchArgs)

            if all_extensions is None and not opts.urn:
                raise TypeError("Missing required property 'all_extensions'")
            __props__.__dict__["all_extensions"] = all_extensions
            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if enable_app_layer is None and not opts.urn:
                raise TypeError("Missing required property 'enable_app_layer'")
            __props__.__dict__["enable_app_layer"] = enable_app_layer
            if enable_rate_controls is None and not opts.urn:
                raise TypeError("Missing required property 'enable_rate_controls'")
            __props__.__dict__["enable_rate_controls"] = enable_rate_controls
            if extensions is None and not opts.urn:
                raise TypeError("Missing required property 'extensions'")
            __props__.__dict__["extensions"] = extensions
        super(AppSecAdvancedSettingsPrefetch, __self__).__init__(
            'akamai:index/appSecAdvancedSettingsPrefetch:AppSecAdvancedSettingsPrefetch',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            all_extensions: Optional[pulumi.Input[bool]] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            enable_app_layer: Optional[pulumi.Input[bool]] = None,
            enable_rate_controls: Optional[pulumi.Input[bool]] = None,
            extensions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'AppSecAdvancedSettingsPrefetch':
        """
        Get an existing AppSecAdvancedSettingsPrefetch resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] all_extensions: . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration associated with the prefetch settings being modified.
        :param pulumi.Input[bool] enable_app_layer: . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
        :param pulumi.Input[bool] enable_rate_controls: . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] extensions: . If `all_extensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won't be used with any file extensions not included in the array. If `all_extensions` is **true**, then this argument must be set to an empty array: **[]**.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecAdvancedSettingsPrefetchState.__new__(_AppSecAdvancedSettingsPrefetchState)

        __props__.__dict__["all_extensions"] = all_extensions
        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["enable_app_layer"] = enable_app_layer
        __props__.__dict__["enable_rate_controls"] = enable_rate_controls
        __props__.__dict__["extensions"] = extensions
        return AppSecAdvancedSettingsPrefetch(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allExtensions")
    def all_extensions(self) -> pulumi.Output[bool]:
        """
        . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
        """
        return pulumi.get(self, "all_extensions")

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        . Unique identifier of the security configuration associated with the prefetch settings being modified.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="enableAppLayer")
    def enable_app_layer(self) -> pulumi.Output[bool]:
        """
        . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
        """
        return pulumi.get(self, "enable_app_layer")

    @property
    @pulumi.getter(name="enableRateControls")
    def enable_rate_controls(self) -> pulumi.Output[bool]:
        """
        . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
        """
        return pulumi.get(self, "enable_rate_controls")

    @property
    @pulumi.getter
    def extensions(self) -> pulumi.Output[Sequence[str]]:
        """
        . If `all_extensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won't be used with any file extensions not included in the array. If `all_extensions` is **true**, then this argument must be set to an empty array: **[]**.
        """
        return pulumi.get(self, "extensions")

