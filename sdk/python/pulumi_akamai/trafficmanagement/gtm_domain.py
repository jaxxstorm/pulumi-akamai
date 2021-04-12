# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['GtmDomainArgs', 'GtmDomain']

@pulumi.input_type
class GtmDomainArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 cname_coalescing_enabled: Optional[pulumi.Input[bool]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 contract: Optional[pulumi.Input[str]] = None,
                 default_error_penalty: Optional[pulumi.Input[int]] = None,
                 default_ssl_client_certificate: Optional[pulumi.Input[str]] = None,
                 default_ssl_client_private_key: Optional[pulumi.Input[str]] = None,
                 default_timeout_penalty: Optional[pulumi.Input[int]] = None,
                 email_notification_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 end_user_mapping_enabled: Optional[pulumi.Input[bool]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 load_feedback: Optional[pulumi.Input[bool]] = None,
                 load_imbalance_percentage: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a GtmDomain resource.
        """
        pulumi.set(__self__, "type", type)
        if cname_coalescing_enabled is not None:
            pulumi.set(__self__, "cname_coalescing_enabled", cname_coalescing_enabled)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if contract is not None:
            pulumi.set(__self__, "contract", contract)
        if default_error_penalty is not None:
            pulumi.set(__self__, "default_error_penalty", default_error_penalty)
        if default_ssl_client_certificate is not None:
            pulumi.set(__self__, "default_ssl_client_certificate", default_ssl_client_certificate)
        if default_ssl_client_private_key is not None:
            pulumi.set(__self__, "default_ssl_client_private_key", default_ssl_client_private_key)
        if default_timeout_penalty is not None:
            pulumi.set(__self__, "default_timeout_penalty", default_timeout_penalty)
        if email_notification_lists is not None:
            pulumi.set(__self__, "email_notification_lists", email_notification_lists)
        if end_user_mapping_enabled is not None:
            pulumi.set(__self__, "end_user_mapping_enabled", end_user_mapping_enabled)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if load_feedback is not None:
            pulumi.set(__self__, "load_feedback", load_feedback)
        if load_imbalance_percentage is not None:
            pulumi.set(__self__, "load_imbalance_percentage", load_imbalance_percentage)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if wait_on_complete is not None:
            pulumi.set(__self__, "wait_on_complete", wait_on_complete)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="cnameCoalescingEnabled")
    def cname_coalescing_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "cname_coalescing_enabled")

    @cname_coalescing_enabled.setter
    def cname_coalescing_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cname_coalescing_enabled", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def contract(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "contract")

    @contract.setter
    def contract(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "contract", value)

    @property
    @pulumi.getter(name="defaultErrorPenalty")
    def default_error_penalty(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "default_error_penalty")

    @default_error_penalty.setter
    def default_error_penalty(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_error_penalty", value)

    @property
    @pulumi.getter(name="defaultSslClientCertificate")
    def default_ssl_client_certificate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "default_ssl_client_certificate")

    @default_ssl_client_certificate.setter
    def default_ssl_client_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_ssl_client_certificate", value)

    @property
    @pulumi.getter(name="defaultSslClientPrivateKey")
    def default_ssl_client_private_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "default_ssl_client_private_key")

    @default_ssl_client_private_key.setter
    def default_ssl_client_private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_ssl_client_private_key", value)

    @property
    @pulumi.getter(name="defaultTimeoutPenalty")
    def default_timeout_penalty(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "default_timeout_penalty")

    @default_timeout_penalty.setter
    def default_timeout_penalty(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_timeout_penalty", value)

    @property
    @pulumi.getter(name="emailNotificationLists")
    def email_notification_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "email_notification_lists")

    @email_notification_lists.setter
    def email_notification_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "email_notification_lists", value)

    @property
    @pulumi.getter(name="endUserMappingEnabled")
    def end_user_mapping_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "end_user_mapping_enabled")

    @end_user_mapping_enabled.setter
    def end_user_mapping_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "end_user_mapping_enabled", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="loadFeedback")
    def load_feedback(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "load_feedback")

    @load_feedback.setter
    def load_feedback(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "load_feedback", value)

    @property
    @pulumi.getter(name="loadImbalancePercentage")
    def load_imbalance_percentage(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "load_imbalance_percentage")

    @load_imbalance_percentage.setter
    def load_imbalance_percentage(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "load_imbalance_percentage", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="waitOnComplete")
    def wait_on_complete(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "wait_on_complete")

    @wait_on_complete.setter
    def wait_on_complete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_on_complete", value)


warnings.warn("""akamai.trafficmanagement.GtmDomain has been deprecated in favor of akamai.GtmDomain""", DeprecationWarning)


class GtmDomain(pulumi.CustomResource):
    warnings.warn("""akamai.trafficmanagement.GtmDomain has been deprecated in favor of akamai.GtmDomain""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cname_coalescing_enabled: Optional[pulumi.Input[bool]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 contract: Optional[pulumi.Input[str]] = None,
                 default_error_penalty: Optional[pulumi.Input[int]] = None,
                 default_ssl_client_certificate: Optional[pulumi.Input[str]] = None,
                 default_ssl_client_private_key: Optional[pulumi.Input[str]] = None,
                 default_timeout_penalty: Optional[pulumi.Input[int]] = None,
                 email_notification_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 end_user_mapping_enabled: Optional[pulumi.Input[bool]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 load_feedback: Optional[pulumi.Input[bool]] = None,
                 load_imbalance_percentage: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a GtmDomain resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GtmDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a GtmDomain resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param GtmDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GtmDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cname_coalescing_enabled: Optional[pulumi.Input[bool]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 contract: Optional[pulumi.Input[str]] = None,
                 default_error_penalty: Optional[pulumi.Input[int]] = None,
                 default_ssl_client_certificate: Optional[pulumi.Input[str]] = None,
                 default_ssl_client_private_key: Optional[pulumi.Input[str]] = None,
                 default_timeout_penalty: Optional[pulumi.Input[int]] = None,
                 email_notification_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 end_user_mapping_enabled: Optional[pulumi.Input[bool]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 load_feedback: Optional[pulumi.Input[bool]] = None,
                 load_imbalance_percentage: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        pulumi.log.warn("""GtmDomain is deprecated: akamai.trafficmanagement.GtmDomain has been deprecated in favor of akamai.GtmDomain""")
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

            __props__['cname_coalescing_enabled'] = cname_coalescing_enabled
            __props__['comment'] = comment
            __props__['contract'] = contract
            __props__['default_error_penalty'] = default_error_penalty
            __props__['default_ssl_client_certificate'] = default_ssl_client_certificate
            __props__['default_ssl_client_private_key'] = default_ssl_client_private_key
            __props__['default_timeout_penalty'] = default_timeout_penalty
            __props__['email_notification_lists'] = email_notification_lists
            __props__['end_user_mapping_enabled'] = end_user_mapping_enabled
            __props__['group'] = group
            __props__['load_feedback'] = load_feedback
            __props__['load_imbalance_percentage'] = load_imbalance_percentage
            __props__['name'] = name
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['wait_on_complete'] = wait_on_complete
            __props__['default_health_max'] = None
            __props__['default_health_multiplier'] = None
            __props__['default_health_threshold'] = None
            __props__['default_max_unreachable_penalty'] = None
            __props__['default_unreachable_threshold'] = None
            __props__['map_update_interval'] = None
            __props__['max_properties'] = None
            __props__['max_resources'] = None
            __props__['max_test_timeout'] = None
            __props__['max_ttl'] = None
            __props__['min_pingable_region_fraction'] = None
            __props__['min_test_interval'] = None
            __props__['min_ttl'] = None
            __props__['ping_interval'] = None
            __props__['ping_packet_size'] = None
            __props__['round_robin_prefix'] = None
            __props__['servermonitor_liveness_count'] = None
            __props__['servermonitor_load_count'] = None
            __props__['servermonitor_pool'] = None
        super(GtmDomain, __self__).__init__(
            'akamai:trafficmanagement/gtmDomain:GtmDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cname_coalescing_enabled: Optional[pulumi.Input[bool]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            contract: Optional[pulumi.Input[str]] = None,
            default_error_penalty: Optional[pulumi.Input[int]] = None,
            default_health_max: Optional[pulumi.Input[float]] = None,
            default_health_multiplier: Optional[pulumi.Input[float]] = None,
            default_health_threshold: Optional[pulumi.Input[float]] = None,
            default_max_unreachable_penalty: Optional[pulumi.Input[int]] = None,
            default_ssl_client_certificate: Optional[pulumi.Input[str]] = None,
            default_ssl_client_private_key: Optional[pulumi.Input[str]] = None,
            default_timeout_penalty: Optional[pulumi.Input[int]] = None,
            default_unreachable_threshold: Optional[pulumi.Input[float]] = None,
            email_notification_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            end_user_mapping_enabled: Optional[pulumi.Input[bool]] = None,
            group: Optional[pulumi.Input[str]] = None,
            load_feedback: Optional[pulumi.Input[bool]] = None,
            load_imbalance_percentage: Optional[pulumi.Input[float]] = None,
            map_update_interval: Optional[pulumi.Input[int]] = None,
            max_properties: Optional[pulumi.Input[int]] = None,
            max_resources: Optional[pulumi.Input[int]] = None,
            max_test_timeout: Optional[pulumi.Input[float]] = None,
            max_ttl: Optional[pulumi.Input[int]] = None,
            min_pingable_region_fraction: Optional[pulumi.Input[float]] = None,
            min_test_interval: Optional[pulumi.Input[int]] = None,
            min_ttl: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            ping_interval: Optional[pulumi.Input[int]] = None,
            ping_packet_size: Optional[pulumi.Input[int]] = None,
            round_robin_prefix: Optional[pulumi.Input[str]] = None,
            servermonitor_liveness_count: Optional[pulumi.Input[int]] = None,
            servermonitor_load_count: Optional[pulumi.Input[int]] = None,
            servermonitor_pool: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            wait_on_complete: Optional[pulumi.Input[bool]] = None) -> 'GtmDomain':
        """
        Get an existing GtmDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cname_coalescing_enabled"] = cname_coalescing_enabled
        __props__["comment"] = comment
        __props__["contract"] = contract
        __props__["default_error_penalty"] = default_error_penalty
        __props__["default_health_max"] = default_health_max
        __props__["default_health_multiplier"] = default_health_multiplier
        __props__["default_health_threshold"] = default_health_threshold
        __props__["default_max_unreachable_penalty"] = default_max_unreachable_penalty
        __props__["default_ssl_client_certificate"] = default_ssl_client_certificate
        __props__["default_ssl_client_private_key"] = default_ssl_client_private_key
        __props__["default_timeout_penalty"] = default_timeout_penalty
        __props__["default_unreachable_threshold"] = default_unreachable_threshold
        __props__["email_notification_lists"] = email_notification_lists
        __props__["end_user_mapping_enabled"] = end_user_mapping_enabled
        __props__["group"] = group
        __props__["load_feedback"] = load_feedback
        __props__["load_imbalance_percentage"] = load_imbalance_percentage
        __props__["map_update_interval"] = map_update_interval
        __props__["max_properties"] = max_properties
        __props__["max_resources"] = max_resources
        __props__["max_test_timeout"] = max_test_timeout
        __props__["max_ttl"] = max_ttl
        __props__["min_pingable_region_fraction"] = min_pingable_region_fraction
        __props__["min_test_interval"] = min_test_interval
        __props__["min_ttl"] = min_ttl
        __props__["name"] = name
        __props__["ping_interval"] = ping_interval
        __props__["ping_packet_size"] = ping_packet_size
        __props__["round_robin_prefix"] = round_robin_prefix
        __props__["servermonitor_liveness_count"] = servermonitor_liveness_count
        __props__["servermonitor_load_count"] = servermonitor_load_count
        __props__["servermonitor_pool"] = servermonitor_pool
        __props__["type"] = type
        __props__["wait_on_complete"] = wait_on_complete
        return GtmDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cnameCoalescingEnabled")
    def cname_coalescing_enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "cname_coalescing_enabled")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def contract(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "contract")

    @property
    @pulumi.getter(name="defaultErrorPenalty")
    def default_error_penalty(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "default_error_penalty")

    @property
    @pulumi.getter(name="defaultHealthMax")
    def default_health_max(self) -> pulumi.Output[float]:
        return pulumi.get(self, "default_health_max")

    @property
    @pulumi.getter(name="defaultHealthMultiplier")
    def default_health_multiplier(self) -> pulumi.Output[float]:
        return pulumi.get(self, "default_health_multiplier")

    @property
    @pulumi.getter(name="defaultHealthThreshold")
    def default_health_threshold(self) -> pulumi.Output[float]:
        return pulumi.get(self, "default_health_threshold")

    @property
    @pulumi.getter(name="defaultMaxUnreachablePenalty")
    def default_max_unreachable_penalty(self) -> pulumi.Output[int]:
        return pulumi.get(self, "default_max_unreachable_penalty")

    @property
    @pulumi.getter(name="defaultSslClientCertificate")
    def default_ssl_client_certificate(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "default_ssl_client_certificate")

    @property
    @pulumi.getter(name="defaultSslClientPrivateKey")
    def default_ssl_client_private_key(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "default_ssl_client_private_key")

    @property
    @pulumi.getter(name="defaultTimeoutPenalty")
    def default_timeout_penalty(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "default_timeout_penalty")

    @property
    @pulumi.getter(name="defaultUnreachableThreshold")
    def default_unreachable_threshold(self) -> pulumi.Output[float]:
        return pulumi.get(self, "default_unreachable_threshold")

    @property
    @pulumi.getter(name="emailNotificationLists")
    def email_notification_lists(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "email_notification_lists")

    @property
    @pulumi.getter(name="endUserMappingEnabled")
    def end_user_mapping_enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "end_user_mapping_enabled")

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "group")

    @property
    @pulumi.getter(name="loadFeedback")
    def load_feedback(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "load_feedback")

    @property
    @pulumi.getter(name="loadImbalancePercentage")
    def load_imbalance_percentage(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "load_imbalance_percentage")

    @property
    @pulumi.getter(name="mapUpdateInterval")
    def map_update_interval(self) -> pulumi.Output[int]:
        return pulumi.get(self, "map_update_interval")

    @property
    @pulumi.getter(name="maxProperties")
    def max_properties(self) -> pulumi.Output[int]:
        return pulumi.get(self, "max_properties")

    @property
    @pulumi.getter(name="maxResources")
    def max_resources(self) -> pulumi.Output[int]:
        return pulumi.get(self, "max_resources")

    @property
    @pulumi.getter(name="maxTestTimeout")
    def max_test_timeout(self) -> pulumi.Output[float]:
        return pulumi.get(self, "max_test_timeout")

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> pulumi.Output[int]:
        return pulumi.get(self, "max_ttl")

    @property
    @pulumi.getter(name="minPingableRegionFraction")
    def min_pingable_region_fraction(self) -> pulumi.Output[float]:
        return pulumi.get(self, "min_pingable_region_fraction")

    @property
    @pulumi.getter(name="minTestInterval")
    def min_test_interval(self) -> pulumi.Output[int]:
        return pulumi.get(self, "min_test_interval")

    @property
    @pulumi.getter(name="minTtl")
    def min_ttl(self) -> pulumi.Output[int]:
        return pulumi.get(self, "min_ttl")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pingInterval")
    def ping_interval(self) -> pulumi.Output[int]:
        return pulumi.get(self, "ping_interval")

    @property
    @pulumi.getter(name="pingPacketSize")
    def ping_packet_size(self) -> pulumi.Output[int]:
        return pulumi.get(self, "ping_packet_size")

    @property
    @pulumi.getter(name="roundRobinPrefix")
    def round_robin_prefix(self) -> pulumi.Output[str]:
        return pulumi.get(self, "round_robin_prefix")

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
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="waitOnComplete")
    def wait_on_complete(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "wait_on_complete")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

