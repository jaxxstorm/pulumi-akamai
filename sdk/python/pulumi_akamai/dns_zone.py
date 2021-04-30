# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DnsZoneArgs', 'DnsZone']

@pulumi.input_type
class DnsZoneArgs:
    def __init__(__self__, *,
                 contract: pulumi.Input[str],
                 group: pulumi.Input[str],
                 type: pulumi.Input[str],
                 zone: pulumi.Input[str],
                 comment: Optional[pulumi.Input[str]] = None,
                 end_customer_id: Optional[pulumi.Input[str]] = None,
                 masters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sign_and_serve: Optional[pulumi.Input[bool]] = None,
                 sign_and_serve_algorithm: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 tsig_key: Optional[pulumi.Input['DnsZoneTsigKeyArgs']] = None):
        """
        The set of arguments for constructing a DnsZone resource.
        """
        pulumi.set(__self__, "contract", contract)
        pulumi.set(__self__, "group", group)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "zone", zone)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if end_customer_id is not None:
            pulumi.set(__self__, "end_customer_id", end_customer_id)
        if masters is not None:
            pulumi.set(__self__, "masters", masters)
        if sign_and_serve is not None:
            pulumi.set(__self__, "sign_and_serve", sign_and_serve)
        if sign_and_serve_algorithm is not None:
            pulumi.set(__self__, "sign_and_serve_algorithm", sign_and_serve_algorithm)
        if target is not None:
            pulumi.set(__self__, "target", target)
        if tsig_key is not None:
            pulumi.set(__self__, "tsig_key", tsig_key)

    @property
    @pulumi.getter
    def contract(self) -> pulumi.Input[str]:
        return pulumi.get(self, "contract")

    @contract.setter
    def contract(self, value: pulumi.Input[str]):
        pulumi.set(self, "contract", value)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Input[str]:
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input[str]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="endCustomerId")
    def end_customer_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "end_customer_id")

    @end_customer_id.setter
    def end_customer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_customer_id", value)

    @property
    @pulumi.getter
    def masters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "masters")

    @masters.setter
    def masters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "masters", value)

    @property
    @pulumi.getter(name="signAndServe")
    def sign_and_serve(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "sign_and_serve")

    @sign_and_serve.setter
    def sign_and_serve(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sign_and_serve", value)

    @property
    @pulumi.getter(name="signAndServeAlgorithm")
    def sign_and_serve_algorithm(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sign_and_serve_algorithm")

    @sign_and_serve_algorithm.setter
    def sign_and_serve_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sign_and_serve_algorithm", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter(name="tsigKey")
    def tsig_key(self) -> Optional[pulumi.Input['DnsZoneTsigKeyArgs']]:
        return pulumi.get(self, "tsig_key")

    @tsig_key.setter
    def tsig_key(self, value: Optional[pulumi.Input['DnsZoneTsigKeyArgs']]):
        pulumi.set(self, "tsig_key", value)


@pulumi.input_type
class _DnsZoneState:
    def __init__(__self__, *,
                 activation_state: Optional[pulumi.Input[str]] = None,
                 alias_count: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 contract: Optional[pulumi.Input[str]] = None,
                 end_customer_id: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 masters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sign_and_serve: Optional[pulumi.Input[bool]] = None,
                 sign_and_serve_algorithm: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 tsig_key: Optional[pulumi.Input['DnsZoneTsigKeyArgs']] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 version_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DnsZone resources.
        """
        if activation_state is not None:
            pulumi.set(__self__, "activation_state", activation_state)
        if alias_count is not None:
            pulumi.set(__self__, "alias_count", alias_count)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if contract is not None:
            pulumi.set(__self__, "contract", contract)
        if end_customer_id is not None:
            pulumi.set(__self__, "end_customer_id", end_customer_id)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if masters is not None:
            pulumi.set(__self__, "masters", masters)
        if sign_and_serve is not None:
            pulumi.set(__self__, "sign_and_serve", sign_and_serve)
        if sign_and_serve_algorithm is not None:
            pulumi.set(__self__, "sign_and_serve_algorithm", sign_and_serve_algorithm)
        if target is not None:
            pulumi.set(__self__, "target", target)
        if tsig_key is not None:
            pulumi.set(__self__, "tsig_key", tsig_key)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if version_id is not None:
            pulumi.set(__self__, "version_id", version_id)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="activationState")
    def activation_state(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "activation_state")

    @activation_state.setter
    def activation_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "activation_state", value)

    @property
    @pulumi.getter(name="aliasCount")
    def alias_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "alias_count")

    @alias_count.setter
    def alias_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "alias_count", value)

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
    @pulumi.getter(name="endCustomerId")
    def end_customer_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "end_customer_id")

    @end_customer_id.setter
    def end_customer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_customer_id", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def masters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "masters")

    @masters.setter
    def masters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "masters", value)

    @property
    @pulumi.getter(name="signAndServe")
    def sign_and_serve(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "sign_and_serve")

    @sign_and_serve.setter
    def sign_and_serve(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sign_and_serve", value)

    @property
    @pulumi.getter(name="signAndServeAlgorithm")
    def sign_and_serve_algorithm(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sign_and_serve_algorithm")

    @sign_and_serve_algorithm.setter
    def sign_and_serve_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sign_and_serve_algorithm", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter(name="tsigKey")
    def tsig_key(self) -> Optional[pulumi.Input['DnsZoneTsigKeyArgs']]:
        return pulumi.get(self, "tsig_key")

    @tsig_key.setter
    def tsig_key(self, value: Optional[pulumi.Input['DnsZoneTsigKeyArgs']]):
        pulumi.set(self, "tsig_key", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "version_id")

    @version_id.setter
    def version_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_id", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class DnsZone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 contract: Optional[pulumi.Input[str]] = None,
                 end_customer_id: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 masters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sign_and_serve: Optional[pulumi.Input[bool]] = None,
                 sign_and_serve_algorithm: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 tsig_key: Optional[pulumi.Input[pulumi.InputType['DnsZoneTsigKeyArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use the `DnsZone` resource to configure a DNS zone that integrates with your existing DNS infrastructure.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        demozone = akamai.DnsZone("demozone",
            comment="some comment",
            contract="ctr_1-AB123",
            group="100",
            masters=[
                "1.2.3.4",
                "1.2.3.5",
            ],
            sign_and_serve=False,
            type="secondary",
            zone="example.com")
        ```
        ## Argument reference

        This resource supports these arguments:

        * `comment` - (Required) A descriptive comment.
        * `contract` - (Required) The contract ID.
        * `group` - (Required) The currently selected group ID.
        * `zone` - (Required) The domain zone, encapsulating any nested subdomains.
        * `type` - (Required) Whether the zone is `primary`, `secondary`, or `alias`.
        * `masters` - (Required for `secondary` zones) The names or IP addresses of the nameservers that the zone data should be retrieved from.
        * `target` - (Required for `alias` zones) The name of the zone whose configuration this zone will copy.
        * `sign_and_serve` - (Optional) Whether DNSSEC Sign and Serve is enabled.
        * `sign_and_serve_algorithm` - (Optional) The algorithm used by Sign and Serve.
        * `tsig_key` - (Optional) The TSIG Key used in secure zone transfers. If used, requires these arguments:
            * `name` - The key name.
            * `algorithm` - The hashing algorithm.
            * `secret` - String known between transfer endpoints.
        * `end_customer_id` - (Optional) A free form identifier for the zone.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DnsZoneArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use the `DnsZone` resource to configure a DNS zone that integrates with your existing DNS infrastructure.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        demozone = akamai.DnsZone("demozone",
            comment="some comment",
            contract="ctr_1-AB123",
            group="100",
            masters=[
                "1.2.3.4",
                "1.2.3.5",
            ],
            sign_and_serve=False,
            type="secondary",
            zone="example.com")
        ```
        ## Argument reference

        This resource supports these arguments:

        * `comment` - (Required) A descriptive comment.
        * `contract` - (Required) The contract ID.
        * `group` - (Required) The currently selected group ID.
        * `zone` - (Required) The domain zone, encapsulating any nested subdomains.
        * `type` - (Required) Whether the zone is `primary`, `secondary`, or `alias`.
        * `masters` - (Required for `secondary` zones) The names or IP addresses of the nameservers that the zone data should be retrieved from.
        * `target` - (Required for `alias` zones) The name of the zone whose configuration this zone will copy.
        * `sign_and_serve` - (Optional) Whether DNSSEC Sign and Serve is enabled.
        * `sign_and_serve_algorithm` - (Optional) The algorithm used by Sign and Serve.
        * `tsig_key` - (Optional) The TSIG Key used in secure zone transfers. If used, requires these arguments:
            * `name` - The key name.
            * `algorithm` - The hashing algorithm.
            * `secret` - String known between transfer endpoints.
        * `end_customer_id` - (Optional) A free form identifier for the zone.

        :param str resource_name: The name of the resource.
        :param DnsZoneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DnsZoneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 contract: Optional[pulumi.Input[str]] = None,
                 end_customer_id: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 masters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sign_and_serve: Optional[pulumi.Input[bool]] = None,
                 sign_and_serve_algorithm: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 tsig_key: Optional[pulumi.Input[pulumi.InputType['DnsZoneTsigKeyArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
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
            __props__ = DnsZoneArgs.__new__(DnsZoneArgs)

            __props__.__dict__["comment"] = comment
            if contract is None and not opts.urn:
                raise TypeError("Missing required property 'contract'")
            __props__.__dict__["contract"] = contract
            __props__.__dict__["end_customer_id"] = end_customer_id
            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            __props__.__dict__["masters"] = masters
            __props__.__dict__["sign_and_serve"] = sign_and_serve
            __props__.__dict__["sign_and_serve_algorithm"] = sign_and_serve_algorithm
            __props__.__dict__["target"] = target
            __props__.__dict__["tsig_key"] = tsig_key
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
            __props__.__dict__["activation_state"] = None
            __props__.__dict__["alias_count"] = None
            __props__.__dict__["version_id"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="akamai:edgedns/dnsZone:DnsZone")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DnsZone, __self__).__init__(
            'akamai:index/dnsZone:DnsZone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            activation_state: Optional[pulumi.Input[str]] = None,
            alias_count: Optional[pulumi.Input[int]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            contract: Optional[pulumi.Input[str]] = None,
            end_customer_id: Optional[pulumi.Input[str]] = None,
            group: Optional[pulumi.Input[str]] = None,
            masters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            sign_and_serve: Optional[pulumi.Input[bool]] = None,
            sign_and_serve_algorithm: Optional[pulumi.Input[str]] = None,
            target: Optional[pulumi.Input[str]] = None,
            tsig_key: Optional[pulumi.Input[pulumi.InputType['DnsZoneTsigKeyArgs']]] = None,
            type: Optional[pulumi.Input[str]] = None,
            version_id: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'DnsZone':
        """
        Get an existing DnsZone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DnsZoneState.__new__(_DnsZoneState)

        __props__.__dict__["activation_state"] = activation_state
        __props__.__dict__["alias_count"] = alias_count
        __props__.__dict__["comment"] = comment
        __props__.__dict__["contract"] = contract
        __props__.__dict__["end_customer_id"] = end_customer_id
        __props__.__dict__["group"] = group
        __props__.__dict__["masters"] = masters
        __props__.__dict__["sign_and_serve"] = sign_and_serve
        __props__.__dict__["sign_and_serve_algorithm"] = sign_and_serve_algorithm
        __props__.__dict__["target"] = target
        __props__.__dict__["tsig_key"] = tsig_key
        __props__.__dict__["type"] = type
        __props__.__dict__["version_id"] = version_id
        __props__.__dict__["zone"] = zone
        return DnsZone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activationState")
    def activation_state(self) -> pulumi.Output[str]:
        return pulumi.get(self, "activation_state")

    @property
    @pulumi.getter(name="aliasCount")
    def alias_count(self) -> pulumi.Output[int]:
        return pulumi.get(self, "alias_count")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def contract(self) -> pulumi.Output[str]:
        return pulumi.get(self, "contract")

    @property
    @pulumi.getter(name="endCustomerId")
    def end_customer_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "end_customer_id")

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def masters(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "masters")

    @property
    @pulumi.getter(name="signAndServe")
    def sign_and_serve(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "sign_and_serve")

    @property
    @pulumi.getter(name="signAndServeAlgorithm")
    def sign_and_serve_algorithm(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "sign_and_serve_algorithm")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "target")

    @property
    @pulumi.getter(name="tsigKey")
    def tsig_key(self) -> pulumi.Output[Optional['outputs.DnsZoneTsigKey']]:
        return pulumi.get(self, "tsig_key")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "version_id")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        return pulumi.get(self, "zone")

