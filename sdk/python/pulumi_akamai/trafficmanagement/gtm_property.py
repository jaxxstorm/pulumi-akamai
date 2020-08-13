# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GtmProperty(pulumi.CustomResource):
    backup_cname: pulumi.Output[str]
    backup_ip: pulumi.Output[str]
    balance_by_download_score: pulumi.Output[bool]
    """
    — (Boolean)
    * `static_ttl`
    * `unreachable_threshold`
    * `health_multiplier`
    * `dynamic_ttl`
    * `max_unreachable_penalty`
    * `map_name`
    * `load_imbalance_percentage`
    * `health_max`
    * `cname`
    * `comments`
    * `ghost_demand_reporting`
    * `min_live_fraction`
    """
    cname: pulumi.Output[str]
    comments: pulumi.Output[str]
    domain: pulumi.Output[str]
    """
    — Domain name
    """
    dynamic_ttl: pulumi.Output[float]
    failback_delay: pulumi.Output[float]
    failover_delay: pulumi.Output[float]
    ghost_demand_reporting: pulumi.Output[bool]
    handout_limit: pulumi.Output[float]
    handout_mode: pulumi.Output[str]
    health_max: pulumi.Output[float]
    health_multiplier: pulumi.Output[float]
    health_threshold: pulumi.Output[float]
    ipv6: pulumi.Output[bool]
    """
    — (Boolean)
    * `stickiness_bonus_percentage`
    * `stickiness_bonus_constant`
    * `health_threshold`
    """
    liveness_tests: pulumi.Output[list]
    """
    — (multiple allowed)

      * `answersRequired` (`bool`) - — (Boolean)
      * `disableNonstandardPortWarning` (`bool`) - — (Boolean)
        * `error_penalty`
      * `disabled` (`bool`) - — (Boolean)
      * `errorPenalty` (`float`)
      * `httpError3xx` (`bool`) - — (Boolean)
      * `httpError4xx` (`bool`) - — (Boolean)
      * `httpError5xx` (`bool`) - — (Boolean)
      * `httpHeaders` (`list`) - — (multiple allowed)
        `name`
        `value`
        * `name` (`str`) - — Liveness test name
          * `test_interval`
          * `test_object_protocol`
          * `test_timeout`
        * `value` (`str`)

      * `name` (`str`) - — Liveness test name
        * `test_interval`
        * `test_object_protocol`
        * `test_timeout`
      * `peerCertificateVerification` (`bool`) - — (Boolean)
      * `recursionRequested` (`bool`) - — (Boolean)
        * `request_string`
        * `resource_type`
        * `response_string`
        * `ssl_client_certificate`
        * `ssl_client_private_key`
        * `test_object`
        * `test_object_password`
        * `test_object_port`
        * `test_object_username`
        * `timeout_penalty`
      * `requestString` (`str`)
      * `resourceType` (`str`)
      * `responseString` (`str`)
      * `sslClientCertificate` (`str`)
      * `sslClientPrivateKey` (`str`)
      * `testInterval` (`float`)
      * `testObject` (`str`)
      * `testObjectPassword` (`str`)
      * `testObjectPort` (`float`)
      * `testObjectProtocol` (`str`)
      * `testObjectUsername` (`str`)
      * `testTimeout` (`float`)
      * `timeoutPenalty` (`float`)
    """
    load_imbalance_percentage: pulumi.Output[float]
    map_name: pulumi.Output[str]
    max_unreachable_penalty: pulumi.Output[float]
    min_live_fraction: pulumi.Output[float]
    name: pulumi.Output[str]
    """
    — Liveness test name
    * `test_interval`
    * `test_object_protocol`
    * `test_timeout`
    """
    score_aggregation_type: pulumi.Output[str]
    static_rr_sets: pulumi.Output[list]
    """
    — (multiple allowed)
    * `type`
    * `ttl`

      * `rdatas` (`list`) - — (List)
      * `ttl` (`float`)
      * `type` (`str`) - — Property type  
        * `score_aggregation_type`
    """
    static_ttl: pulumi.Output[float]
    stickiness_bonus_constant: pulumi.Output[float]
    stickiness_bonus_percentage: pulumi.Output[float]
    traffic_targets: pulumi.Output[list]
    """
    — (multiple allowed)
    * `datacenter_id`

      * `datacenter_id` (`float`)
      * `enabled` (`bool`) - — (Boolean)
        * `weight`
      * `handoutCname` (`str`)
      * `name` (`str`) - — Liveness test name
        * `test_interval`
        * `test_object_protocol`
        * `test_timeout`
      * `servers` (`list`) - — (List)
      * `weight` (`float`)
    """
    type: pulumi.Output[str]
    """
    — Property type  
    * `score_aggregation_type`
    """
    unreachable_threshold: pulumi.Output[float]
    use_computed_targets: pulumi.Output[bool]
    """
    — (Boolean)
    * `backup_ip`
    """
    wait_on_complete: pulumi.Output[bool]
    """
    — (Boolean, Default: true) Wait for transaction to complete
    * `failover_delay`
    * `failback_delay`
    """
    weighted_hash_bits_for_ipv4: pulumi.Output[float]
    weighted_hash_bits_for_ipv6: pulumi.Output[float]
    def __init__(__self__, resource_name, opts=None, backup_cname=None, backup_ip=None, balance_by_download_score=None, cname=None, comments=None, domain=None, dynamic_ttl=None, failback_delay=None, failover_delay=None, ghost_demand_reporting=None, handout_limit=None, handout_mode=None, health_max=None, health_multiplier=None, health_threshold=None, ipv6=None, liveness_tests=None, load_imbalance_percentage=None, map_name=None, max_unreachable_penalty=None, min_live_fraction=None, name=None, score_aggregation_type=None, static_rr_sets=None, static_ttl=None, stickiness_bonus_constant=None, stickiness_bonus_percentage=None, traffic_targets=None, type=None, unreachable_threshold=None, use_computed_targets=None, wait_on_complete=None, __props__=None, __name__=None, __opts__=None):
        """
        `trafficmanagement.GtmProperty` provides the resource for creating, configuring and importing a gtm property to integrate easily with your existing GTM infrastructure to provide a secure, high performance, highly available and scalable solution for Global Traffic Management. Note: Import requires an ID of the format: `existing_domain_name`:`existing_property_name`

        ## Example Usage
        ### Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        demo_property = akamai.trafficmanagement.GtmProperty("demoProperty",
            domain="demo_domain.akadns.net",
            handout_limit=5,
            handout_mode="normal",
            score_aggregation_type="median",
            traffic_targets=[{
                "datacenter_id": 3131,
            }],
            type="weighted-round-robin")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] balance_by_download_score: — (Boolean)
               * `static_ttl`
               * `unreachable_threshold`
               * `health_multiplier`
               * `dynamic_ttl`
               * `max_unreachable_penalty`
               * `map_name`
               * `load_imbalance_percentage`
               * `health_max`
               * `cname`
               * `comments`
               * `ghost_demand_reporting`
               * `min_live_fraction`
        :param pulumi.Input[str] domain: — Domain name
        :param pulumi.Input[bool] ipv6: — (Boolean)
               * `stickiness_bonus_percentage`
               * `stickiness_bonus_constant`
               * `health_threshold`
        :param pulumi.Input[list] liveness_tests: — (multiple allowed)
        :param pulumi.Input[str] name: — Liveness test name
               * `test_interval`
               * `test_object_protocol`
               * `test_timeout`
        :param pulumi.Input[list] static_rr_sets: — (multiple allowed)
               * `type`
               * `ttl`
        :param pulumi.Input[list] traffic_targets: — (multiple allowed)
               * `datacenter_id`
        :param pulumi.Input[str] type: — Property type  
               * `score_aggregation_type`
        :param pulumi.Input[bool] use_computed_targets: — (Boolean)
               * `backup_ip`
        :param pulumi.Input[bool] wait_on_complete: — (Boolean, Default: true) Wait for transaction to complete
               * `failover_delay`
               * `failback_delay`

        The **liveness_tests** object supports the following:

          * `answersRequired` (`pulumi.Input[bool]`) - — (Boolean)
          * `disableNonstandardPortWarning` (`pulumi.Input[bool]`) - — (Boolean)
            * `error_penalty`
          * `disabled` (`pulumi.Input[bool]`) - — (Boolean)
          * `errorPenalty` (`pulumi.Input[float]`)
          * `httpError3xx` (`pulumi.Input[bool]`) - — (Boolean)
          * `httpError4xx` (`pulumi.Input[bool]`) - — (Boolean)
          * `httpError5xx` (`pulumi.Input[bool]`) - — (Boolean)
          * `httpHeaders` (`pulumi.Input[list]`) - — (multiple allowed)
            `name`
            `value`
            * `name` (`pulumi.Input[str]`) - — Liveness test name
              * `test_interval`
              * `test_object_protocol`
              * `test_timeout`
            * `value` (`pulumi.Input[str]`)

          * `name` (`pulumi.Input[str]`) - — Liveness test name
            * `test_interval`
            * `test_object_protocol`
            * `test_timeout`
          * `peerCertificateVerification` (`pulumi.Input[bool]`) - — (Boolean)
          * `recursionRequested` (`pulumi.Input[bool]`) - — (Boolean)
            * `request_string`
            * `resource_type`
            * `response_string`
            * `ssl_client_certificate`
            * `ssl_client_private_key`
            * `test_object`
            * `test_object_password`
            * `test_object_port`
            * `test_object_username`
            * `timeout_penalty`
          * `requestString` (`pulumi.Input[str]`)
          * `resourceType` (`pulumi.Input[str]`)
          * `responseString` (`pulumi.Input[str]`)
          * `sslClientCertificate` (`pulumi.Input[str]`)
          * `sslClientPrivateKey` (`pulumi.Input[str]`)
          * `testInterval` (`pulumi.Input[float]`)
          * `testObject` (`pulumi.Input[str]`)
          * `testObjectPassword` (`pulumi.Input[str]`)
          * `testObjectPort` (`pulumi.Input[float]`)
          * `testObjectProtocol` (`pulumi.Input[str]`)
          * `testObjectUsername` (`pulumi.Input[str]`)
          * `testTimeout` (`pulumi.Input[float]`)
          * `timeoutPenalty` (`pulumi.Input[float]`)

        The **static_rr_sets** object supports the following:

          * `rdatas` (`pulumi.Input[list]`) - — (List)
          * `ttl` (`pulumi.Input[float]`)
          * `type` (`pulumi.Input[str]`) - — Property type  
            * `score_aggregation_type`

        The **traffic_targets** object supports the following:

          * `datacenter_id` (`pulumi.Input[float]`)
          * `enabled` (`pulumi.Input[bool]`) - — (Boolean)
            * `weight`
          * `handoutCname` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`) - — Liveness test name
            * `test_interval`
            * `test_object_protocol`
            * `test_timeout`
          * `servers` (`pulumi.Input[list]`) - — (List)
          * `weight` (`pulumi.Input[float]`)
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

            __props__['backup_cname'] = backup_cname
            __props__['backup_ip'] = backup_ip
            __props__['balance_by_download_score'] = balance_by_download_score
            __props__['cname'] = cname
            __props__['comments'] = comments
            if domain is None:
                raise TypeError("Missing required property 'domain'")
            __props__['domain'] = domain
            __props__['dynamic_ttl'] = dynamic_ttl
            __props__['failback_delay'] = failback_delay
            __props__['failover_delay'] = failover_delay
            __props__['ghost_demand_reporting'] = ghost_demand_reporting
            if handout_limit is None:
                raise TypeError("Missing required property 'handout_limit'")
            __props__['handout_limit'] = handout_limit
            if handout_mode is None:
                raise TypeError("Missing required property 'handout_mode'")
            __props__['handout_mode'] = handout_mode
            __props__['health_max'] = health_max
            __props__['health_multiplier'] = health_multiplier
            __props__['health_threshold'] = health_threshold
            __props__['ipv6'] = ipv6
            __props__['liveness_tests'] = liveness_tests
            __props__['load_imbalance_percentage'] = load_imbalance_percentage
            __props__['map_name'] = map_name
            __props__['max_unreachable_penalty'] = max_unreachable_penalty
            __props__['min_live_fraction'] = min_live_fraction
            __props__['name'] = name
            if score_aggregation_type is None:
                raise TypeError("Missing required property 'score_aggregation_type'")
            __props__['score_aggregation_type'] = score_aggregation_type
            __props__['static_rr_sets'] = static_rr_sets
            __props__['static_ttl'] = static_ttl
            __props__['stickiness_bonus_constant'] = stickiness_bonus_constant
            __props__['stickiness_bonus_percentage'] = stickiness_bonus_percentage
            if traffic_targets is None:
                raise TypeError("Missing required property 'traffic_targets'")
            __props__['traffic_targets'] = traffic_targets
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['unreachable_threshold'] = unreachable_threshold
            __props__['use_computed_targets'] = use_computed_targets
            __props__['wait_on_complete'] = wait_on_complete
            __props__['weighted_hash_bits_for_ipv4'] = None
            __props__['weighted_hash_bits_for_ipv6'] = None
        super(GtmProperty, __self__).__init__(
            'akamai:trafficmanagement/gtmProperty:GtmProperty',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, backup_cname=None, backup_ip=None, balance_by_download_score=None, cname=None, comments=None, domain=None, dynamic_ttl=None, failback_delay=None, failover_delay=None, ghost_demand_reporting=None, handout_limit=None, handout_mode=None, health_max=None, health_multiplier=None, health_threshold=None, ipv6=None, liveness_tests=None, load_imbalance_percentage=None, map_name=None, max_unreachable_penalty=None, min_live_fraction=None, name=None, score_aggregation_type=None, static_rr_sets=None, static_ttl=None, stickiness_bonus_constant=None, stickiness_bonus_percentage=None, traffic_targets=None, type=None, unreachable_threshold=None, use_computed_targets=None, wait_on_complete=None, weighted_hash_bits_for_ipv4=None, weighted_hash_bits_for_ipv6=None):
        """
        Get an existing GtmProperty resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] balance_by_download_score: — (Boolean)
               * `static_ttl`
               * `unreachable_threshold`
               * `health_multiplier`
               * `dynamic_ttl`
               * `max_unreachable_penalty`
               * `map_name`
               * `load_imbalance_percentage`
               * `health_max`
               * `cname`
               * `comments`
               * `ghost_demand_reporting`
               * `min_live_fraction`
        :param pulumi.Input[str] domain: — Domain name
        :param pulumi.Input[bool] ipv6: — (Boolean)
               * `stickiness_bonus_percentage`
               * `stickiness_bonus_constant`
               * `health_threshold`
        :param pulumi.Input[list] liveness_tests: — (multiple allowed)
        :param pulumi.Input[str] name: — Liveness test name
               * `test_interval`
               * `test_object_protocol`
               * `test_timeout`
        :param pulumi.Input[list] static_rr_sets: — (multiple allowed)
               * `type`
               * `ttl`
        :param pulumi.Input[list] traffic_targets: — (multiple allowed)
               * `datacenter_id`
        :param pulumi.Input[str] type: — Property type  
               * `score_aggregation_type`
        :param pulumi.Input[bool] use_computed_targets: — (Boolean)
               * `backup_ip`
        :param pulumi.Input[bool] wait_on_complete: — (Boolean, Default: true) Wait for transaction to complete
               * `failover_delay`
               * `failback_delay`

        The **liveness_tests** object supports the following:

          * `answersRequired` (`pulumi.Input[bool]`) - — (Boolean)
          * `disableNonstandardPortWarning` (`pulumi.Input[bool]`) - — (Boolean)
            * `error_penalty`
          * `disabled` (`pulumi.Input[bool]`) - — (Boolean)
          * `errorPenalty` (`pulumi.Input[float]`)
          * `httpError3xx` (`pulumi.Input[bool]`) - — (Boolean)
          * `httpError4xx` (`pulumi.Input[bool]`) - — (Boolean)
          * `httpError5xx` (`pulumi.Input[bool]`) - — (Boolean)
          * `httpHeaders` (`pulumi.Input[list]`) - — (multiple allowed)
            `name`
            `value`
            * `name` (`pulumi.Input[str]`) - — Liveness test name
              * `test_interval`
              * `test_object_protocol`
              * `test_timeout`
            * `value` (`pulumi.Input[str]`)

          * `name` (`pulumi.Input[str]`) - — Liveness test name
            * `test_interval`
            * `test_object_protocol`
            * `test_timeout`
          * `peerCertificateVerification` (`pulumi.Input[bool]`) - — (Boolean)
          * `recursionRequested` (`pulumi.Input[bool]`) - — (Boolean)
            * `request_string`
            * `resource_type`
            * `response_string`
            * `ssl_client_certificate`
            * `ssl_client_private_key`
            * `test_object`
            * `test_object_password`
            * `test_object_port`
            * `test_object_username`
            * `timeout_penalty`
          * `requestString` (`pulumi.Input[str]`)
          * `resourceType` (`pulumi.Input[str]`)
          * `responseString` (`pulumi.Input[str]`)
          * `sslClientCertificate` (`pulumi.Input[str]`)
          * `sslClientPrivateKey` (`pulumi.Input[str]`)
          * `testInterval` (`pulumi.Input[float]`)
          * `testObject` (`pulumi.Input[str]`)
          * `testObjectPassword` (`pulumi.Input[str]`)
          * `testObjectPort` (`pulumi.Input[float]`)
          * `testObjectProtocol` (`pulumi.Input[str]`)
          * `testObjectUsername` (`pulumi.Input[str]`)
          * `testTimeout` (`pulumi.Input[float]`)
          * `timeoutPenalty` (`pulumi.Input[float]`)

        The **static_rr_sets** object supports the following:

          * `rdatas` (`pulumi.Input[list]`) - — (List)
          * `ttl` (`pulumi.Input[float]`)
          * `type` (`pulumi.Input[str]`) - — Property type  
            * `score_aggregation_type`

        The **traffic_targets** object supports the following:

          * `datacenter_id` (`pulumi.Input[float]`)
          * `enabled` (`pulumi.Input[bool]`) - — (Boolean)
            * `weight`
          * `handoutCname` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`) - — Liveness test name
            * `test_interval`
            * `test_object_protocol`
            * `test_timeout`
          * `servers` (`pulumi.Input[list]`) - — (List)
          * `weight` (`pulumi.Input[float]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backup_cname"] = backup_cname
        __props__["backup_ip"] = backup_ip
        __props__["balance_by_download_score"] = balance_by_download_score
        __props__["cname"] = cname
        __props__["comments"] = comments
        __props__["domain"] = domain
        __props__["dynamic_ttl"] = dynamic_ttl
        __props__["failback_delay"] = failback_delay
        __props__["failover_delay"] = failover_delay
        __props__["ghost_demand_reporting"] = ghost_demand_reporting
        __props__["handout_limit"] = handout_limit
        __props__["handout_mode"] = handout_mode
        __props__["health_max"] = health_max
        __props__["health_multiplier"] = health_multiplier
        __props__["health_threshold"] = health_threshold
        __props__["ipv6"] = ipv6
        __props__["liveness_tests"] = liveness_tests
        __props__["load_imbalance_percentage"] = load_imbalance_percentage
        __props__["map_name"] = map_name
        __props__["max_unreachable_penalty"] = max_unreachable_penalty
        __props__["min_live_fraction"] = min_live_fraction
        __props__["name"] = name
        __props__["score_aggregation_type"] = score_aggregation_type
        __props__["static_rr_sets"] = static_rr_sets
        __props__["static_ttl"] = static_ttl
        __props__["stickiness_bonus_constant"] = stickiness_bonus_constant
        __props__["stickiness_bonus_percentage"] = stickiness_bonus_percentage
        __props__["traffic_targets"] = traffic_targets
        __props__["type"] = type
        __props__["unreachable_threshold"] = unreachable_threshold
        __props__["use_computed_targets"] = use_computed_targets
        __props__["wait_on_complete"] = wait_on_complete
        __props__["weighted_hash_bits_for_ipv4"] = weighted_hash_bits_for_ipv4
        __props__["weighted_hash_bits_for_ipv6"] = weighted_hash_bits_for_ipv6
        return GtmProperty(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
