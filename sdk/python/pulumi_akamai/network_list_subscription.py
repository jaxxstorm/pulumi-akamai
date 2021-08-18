# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['NetworkListSubscriptionArgs', 'NetworkListSubscription']

@pulumi.input_type
class NetworkListSubscriptionArgs:
    def __init__(__self__, *,
                 network_lists: pulumi.Input[Sequence[pulumi.Input[str]]],
                 recipients: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a NetworkListSubscription resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] network_lists: A list containing one or more IDs of the network lists to which the indicated email
               addresses should be subscribed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] recipients: A bracketed, comma-separated list of email addresses that will be notified of changes to any
               of the specified network lists.
        """
        pulumi.set(__self__, "network_lists", network_lists)
        pulumi.set(__self__, "recipients", recipients)

    @property
    @pulumi.getter(name="networkLists")
    def network_lists(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list containing one or more IDs of the network lists to which the indicated email
        addresses should be subscribed.
        """
        return pulumi.get(self, "network_lists")

    @network_lists.setter
    def network_lists(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "network_lists", value)

    @property
    @pulumi.getter
    def recipients(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A bracketed, comma-separated list of email addresses that will be notified of changes to any
        of the specified network lists.
        """
        return pulumi.get(self, "recipients")

    @recipients.setter
    def recipients(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "recipients", value)


@pulumi.input_type
class _NetworkListSubscriptionState:
    def __init__(__self__, *,
                 network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 recipients: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering NetworkListSubscription resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] network_lists: A list containing one or more IDs of the network lists to which the indicated email
               addresses should be subscribed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] recipients: A bracketed, comma-separated list of email addresses that will be notified of changes to any
               of the specified network lists.
        """
        if network_lists is not None:
            pulumi.set(__self__, "network_lists", network_lists)
        if recipients is not None:
            pulumi.set(__self__, "recipients", recipients)

    @property
    @pulumi.getter(name="networkLists")
    def network_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list containing one or more IDs of the network lists to which the indicated email
        addresses should be subscribed.
        """
        return pulumi.get(self, "network_lists")

    @network_lists.setter
    def network_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "network_lists", value)

    @property
    @pulumi.getter
    def recipients(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A bracketed, comma-separated list of email addresses that will be notified of changes to any
        of the specified network lists.
        """
        return pulumi.get(self, "recipients")

    @recipients.setter
    def recipients(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "recipients", value)


class NetworkListSubscription(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 recipients: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Use the `NetworkListSubscription` resource to specify a set of email addresses to be notified of changes to any
        of a set of network lists.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        network_lists_filter = akamai.get_network_lists(name=var["network_list"])
        subscribe = akamai.NetworkListSubscription("subscribe",
            network_lists=network_lists_filter.lists,
            recipients=["user@example.com"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] network_lists: A list containing one or more IDs of the network lists to which the indicated email
               addresses should be subscribed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] recipients: A bracketed, comma-separated list of email addresses that will be notified of changes to any
               of the specified network lists.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkListSubscriptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use the `NetworkListSubscription` resource to specify a set of email addresses to be notified of changes to any
        of a set of network lists.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        network_lists_filter = akamai.get_network_lists(name=var["network_list"])
        subscribe = akamai.NetworkListSubscription("subscribe",
            network_lists=network_lists_filter.lists,
            recipients=["user@example.com"])
        ```

        :param str resource_name: The name of the resource.
        :param NetworkListSubscriptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkListSubscriptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 recipients: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = NetworkListSubscriptionArgs.__new__(NetworkListSubscriptionArgs)

            if network_lists is None and not opts.urn:
                raise TypeError("Missing required property 'network_lists'")
            __props__.__dict__["network_lists"] = network_lists
            if recipients is None and not opts.urn:
                raise TypeError("Missing required property 'recipients'")
            __props__.__dict__["recipients"] = recipients
        super(NetworkListSubscription, __self__).__init__(
            'akamai:index/networkListSubscription:NetworkListSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            network_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            recipients: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'NetworkListSubscription':
        """
        Get an existing NetworkListSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] network_lists: A list containing one or more IDs of the network lists to which the indicated email
               addresses should be subscribed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] recipients: A bracketed, comma-separated list of email addresses that will be notified of changes to any
               of the specified network lists.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkListSubscriptionState.__new__(_NetworkListSubscriptionState)

        __props__.__dict__["network_lists"] = network_lists
        __props__.__dict__["recipients"] = recipients
        return NetworkListSubscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="networkLists")
    def network_lists(self) -> pulumi.Output[Sequence[str]]:
        """
        A list containing one or more IDs of the network lists to which the indicated email
        addresses should be subscribed.
        """
        return pulumi.get(self, "network_lists")

    @property
    @pulumi.getter
    def recipients(self) -> pulumi.Output[Sequence[str]]:
        """
        A bracketed, comma-separated list of email addresses that will be notified of changes to any
        of the specified network lists.
        """
        return pulumi.get(self, "recipients")

