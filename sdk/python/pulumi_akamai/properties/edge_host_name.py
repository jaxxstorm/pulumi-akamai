# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EdgeHostNameArgs', 'EdgeHostName']

@pulumi.input_type
class EdgeHostNameArgs:
    def __init__(__self__, *,
                 edge_hostname: pulumi.Input[str],
                 ip_behavior: pulumi.Input[str],
                 certificate: Optional[pulumi.Input[int]] = None,
                 contract: Optional[pulumi.Input[str]] = None,
                 contract_id: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 product: Optional[pulumi.Input[str]] = None,
                 product_id: Optional[pulumi.Input[str]] = None,
                 status_update_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 use_cases: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EdgeHostName resource.
        :param pulumi.Input[str] edge_hostname: One or more edge hostnames. The number of edge hostnames must be less than or equal to the number of public hostnames.
        :param pulumi.Input[str] ip_behavior: Which version of the IP protocol to use: `IPV4` for version 4 only, `IPV6_PERFORMANCE` for version 6 only, or `IPV6_COMPLIANCE` for both 4 and 6.
        :param pulumi.Input[int] certificate: Required only when creating an Enhanced TLS edge hostname. This argument sets the certificate enrollment ID. Edge hostnames for Enhanced TLS end in `edgekey.net`. You can retrieve this ID from the [Certificate Provisioning Service CLI](https://github.com/akamai/cli-cps) .
        :param pulumi.Input[str] contract: Replaced by `contract_id`. Maintained for legacy purposes.
        :param pulumi.Input[str] contract_id: A contract's unique ID, including the `ctr_` prefix.
        :param pulumi.Input[str] group: Replaced by `group_id`. Maintained for legacy purposes.
        :param pulumi.Input[str] group_id: A group's unique ID, including the `grp_` prefix.
        :param pulumi.Input[str] product: Replaced by `product_id`. Maintained for legacy purposes.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] status_update_emails: Email address that should receive updates on the IP behavior update request. Required for update operation.
        :param pulumi.Input[str] use_cases: A JSON encoded list of use cases.
        """
        pulumi.set(__self__, "edge_hostname", edge_hostname)
        pulumi.set(__self__, "ip_behavior", ip_behavior)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if contract is not None:
            warnings.warn("""The setting \"contract\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""contract is deprecated: The setting \"contract\" has been deprecated.""")
        if contract is not None:
            pulumi.set(__self__, "contract", contract)
        if contract_id is not None:
            pulumi.set(__self__, "contract_id", contract_id)
        if group is not None:
            warnings.warn("""The setting \"group\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""group is deprecated: The setting \"group\" has been deprecated.""")
        if group is not None:
            pulumi.set(__self__, "group", group)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if product is not None:
            warnings.warn("""The setting \"product\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""product is deprecated: The setting \"product\" has been deprecated.""")
        if product is not None:
            pulumi.set(__self__, "product", product)
        if product_id is not None:
            pulumi.set(__self__, "product_id", product_id)
        if status_update_emails is not None:
            pulumi.set(__self__, "status_update_emails", status_update_emails)
        if use_cases is not None:
            pulumi.set(__self__, "use_cases", use_cases)

    @property
    @pulumi.getter(name="edgeHostname")
    def edge_hostname(self) -> pulumi.Input[str]:
        """
        One or more edge hostnames. The number of edge hostnames must be less than or equal to the number of public hostnames.
        """
        return pulumi.get(self, "edge_hostname")

    @edge_hostname.setter
    def edge_hostname(self, value: pulumi.Input[str]):
        pulumi.set(self, "edge_hostname", value)

    @property
    @pulumi.getter(name="ipBehavior")
    def ip_behavior(self) -> pulumi.Input[str]:
        """
        Which version of the IP protocol to use: `IPV4` for version 4 only, `IPV6_PERFORMANCE` for version 6 only, or `IPV6_COMPLIANCE` for both 4 and 6.
        """
        return pulumi.get(self, "ip_behavior")

    @ip_behavior.setter
    def ip_behavior(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_behavior", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[int]]:
        """
        Required only when creating an Enhanced TLS edge hostname. This argument sets the certificate enrollment ID. Edge hostnames for Enhanced TLS end in `edgekey.net`. You can retrieve this ID from the [Certificate Provisioning Service CLI](https://github.com/akamai/cli-cps) .
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter
    def contract(self) -> Optional[pulumi.Input[str]]:
        """
        Replaced by `contract_id`. Maintained for legacy purposes.
        """
        return pulumi.get(self, "contract")

    @contract.setter
    def contract(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "contract", value)

    @property
    @pulumi.getter(name="contractId")
    def contract_id(self) -> Optional[pulumi.Input[str]]:
        """
        A contract's unique ID, including the `ctr_` prefix.
        """
        return pulumi.get(self, "contract_id")

    @contract_id.setter
    def contract_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "contract_id", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        Replaced by `group_id`. Maintained for legacy purposes.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        A group's unique ID, including the `grp_` prefix.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter
    def product(self) -> Optional[pulumi.Input[str]]:
        """
        Replaced by `product_id`. Maintained for legacy purposes.
        """
        return pulumi.get(self, "product")

    @product.setter
    def product(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product", value)

    @property
    @pulumi.getter(name="productId")
    def product_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "product_id")

    @product_id.setter
    def product_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product_id", value)

    @property
    @pulumi.getter(name="statusUpdateEmails")
    def status_update_emails(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Email address that should receive updates on the IP behavior update request. Required for update operation.
        """
        return pulumi.get(self, "status_update_emails")

    @status_update_emails.setter
    def status_update_emails(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "status_update_emails", value)

    @property
    @pulumi.getter(name="useCases")
    def use_cases(self) -> Optional[pulumi.Input[str]]:
        """
        A JSON encoded list of use cases.
        """
        return pulumi.get(self, "use_cases")

    @use_cases.setter
    def use_cases(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "use_cases", value)


@pulumi.input_type
class _EdgeHostNameState:
    def __init__(__self__, *,
                 certificate: Optional[pulumi.Input[int]] = None,
                 contract: Optional[pulumi.Input[str]] = None,
                 contract_id: Optional[pulumi.Input[str]] = None,
                 edge_hostname: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 ip_behavior: Optional[pulumi.Input[str]] = None,
                 product: Optional[pulumi.Input[str]] = None,
                 product_id: Optional[pulumi.Input[str]] = None,
                 status_update_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 use_cases: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EdgeHostName resources.
        :param pulumi.Input[int] certificate: Required only when creating an Enhanced TLS edge hostname. This argument sets the certificate enrollment ID. Edge hostnames for Enhanced TLS end in `edgekey.net`. You can retrieve this ID from the [Certificate Provisioning Service CLI](https://github.com/akamai/cli-cps) .
        :param pulumi.Input[str] contract: Replaced by `contract_id`. Maintained for legacy purposes.
        :param pulumi.Input[str] contract_id: A contract's unique ID, including the `ctr_` prefix.
        :param pulumi.Input[str] edge_hostname: One or more edge hostnames. The number of edge hostnames must be less than or equal to the number of public hostnames.
        :param pulumi.Input[str] group: Replaced by `group_id`. Maintained for legacy purposes.
        :param pulumi.Input[str] group_id: A group's unique ID, including the `grp_` prefix.
        :param pulumi.Input[str] ip_behavior: Which version of the IP protocol to use: `IPV4` for version 4 only, `IPV6_PERFORMANCE` for version 6 only, or `IPV6_COMPLIANCE` for both 4 and 6.
        :param pulumi.Input[str] product: Replaced by `product_id`. Maintained for legacy purposes.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] status_update_emails: Email address that should receive updates on the IP behavior update request. Required for update operation.
        :param pulumi.Input[str] use_cases: A JSON encoded list of use cases.
        """
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if contract is not None:
            warnings.warn("""The setting \"contract\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""contract is deprecated: The setting \"contract\" has been deprecated.""")
        if contract is not None:
            pulumi.set(__self__, "contract", contract)
        if contract_id is not None:
            pulumi.set(__self__, "contract_id", contract_id)
        if edge_hostname is not None:
            pulumi.set(__self__, "edge_hostname", edge_hostname)
        if group is not None:
            warnings.warn("""The setting \"group\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""group is deprecated: The setting \"group\" has been deprecated.""")
        if group is not None:
            pulumi.set(__self__, "group", group)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if ip_behavior is not None:
            pulumi.set(__self__, "ip_behavior", ip_behavior)
        if product is not None:
            warnings.warn("""The setting \"product\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""product is deprecated: The setting \"product\" has been deprecated.""")
        if product is not None:
            pulumi.set(__self__, "product", product)
        if product_id is not None:
            pulumi.set(__self__, "product_id", product_id)
        if status_update_emails is not None:
            pulumi.set(__self__, "status_update_emails", status_update_emails)
        if use_cases is not None:
            pulumi.set(__self__, "use_cases", use_cases)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[int]]:
        """
        Required only when creating an Enhanced TLS edge hostname. This argument sets the certificate enrollment ID. Edge hostnames for Enhanced TLS end in `edgekey.net`. You can retrieve this ID from the [Certificate Provisioning Service CLI](https://github.com/akamai/cli-cps) .
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter
    def contract(self) -> Optional[pulumi.Input[str]]:
        """
        Replaced by `contract_id`. Maintained for legacy purposes.
        """
        return pulumi.get(self, "contract")

    @contract.setter
    def contract(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "contract", value)

    @property
    @pulumi.getter(name="contractId")
    def contract_id(self) -> Optional[pulumi.Input[str]]:
        """
        A contract's unique ID, including the `ctr_` prefix.
        """
        return pulumi.get(self, "contract_id")

    @contract_id.setter
    def contract_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "contract_id", value)

    @property
    @pulumi.getter(name="edgeHostname")
    def edge_hostname(self) -> Optional[pulumi.Input[str]]:
        """
        One or more edge hostnames. The number of edge hostnames must be less than or equal to the number of public hostnames.
        """
        return pulumi.get(self, "edge_hostname")

    @edge_hostname.setter
    def edge_hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edge_hostname", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        Replaced by `group_id`. Maintained for legacy purposes.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        A group's unique ID, including the `grp_` prefix.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="ipBehavior")
    def ip_behavior(self) -> Optional[pulumi.Input[str]]:
        """
        Which version of the IP protocol to use: `IPV4` for version 4 only, `IPV6_PERFORMANCE` for version 6 only, or `IPV6_COMPLIANCE` for both 4 and 6.
        """
        return pulumi.get(self, "ip_behavior")

    @ip_behavior.setter
    def ip_behavior(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_behavior", value)

    @property
    @pulumi.getter
    def product(self) -> Optional[pulumi.Input[str]]:
        """
        Replaced by `product_id`. Maintained for legacy purposes.
        """
        return pulumi.get(self, "product")

    @product.setter
    def product(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product", value)

    @property
    @pulumi.getter(name="productId")
    def product_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "product_id")

    @product_id.setter
    def product_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product_id", value)

    @property
    @pulumi.getter(name="statusUpdateEmails")
    def status_update_emails(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Email address that should receive updates on the IP behavior update request. Required for update operation.
        """
        return pulumi.get(self, "status_update_emails")

    @status_update_emails.setter
    def status_update_emails(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "status_update_emails", value)

    @property
    @pulumi.getter(name="useCases")
    def use_cases(self) -> Optional[pulumi.Input[str]]:
        """
        A JSON encoded list of use cases.
        """
        return pulumi.get(self, "use_cases")

    @use_cases.setter
    def use_cases(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "use_cases", value)


warnings.warn("""akamai.properties.EdgeHostName has been deprecated in favor of akamai.EdgeHostName""", DeprecationWarning)


class EdgeHostName(pulumi.CustomResource):
    warnings.warn("""akamai.properties.EdgeHostName has been deprecated in favor of akamai.EdgeHostName""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[int]] = None,
                 contract: Optional[pulumi.Input[str]] = None,
                 contract_id: Optional[pulumi.Input[str]] = None,
                 edge_hostname: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 ip_behavior: Optional[pulumi.Input[str]] = None,
                 product: Optional[pulumi.Input[str]] = None,
                 product_id: Optional[pulumi.Input[str]] = None,
                 status_update_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 use_cases: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `EdgeHostName` resource lets you configure a secure edge hostname. Your
        edge hostname determines how requests for your site, app, or content are mapped to
        Akamai edge servers.

        An edge hostname is the CNAME target you use when directing your end user traffic to
        Akamai. Each hostname assigned to a property has a corresponding edge hostname.

        Akamai supports three types of edge hostnames, depending on the level of security
        you need for your traffic: Standard TLS, Enhanced TLS, and Shared Certificate. When
        entering the `edge_hostname` attribute, you need to include a specific domain suffix
        for your edge hostname type:

        | Edge hostname type | Domain suffix |
        |------|-------|
        | Enhanced TLS | edgekey.net |
        | Standard TLS | edgesuite.net |
        | Shared Cert | akamaized.net |

        For example, if you use Standard TLS and have `www.example.com` as a hostname, your edge hostname would be `www.example.com.edgesuite.net`. If you wanted to use Enhanced TLS with the same hostname, your edge hostname would be `www.example.com.edgekey.net`. See the [Property Manager API (PAPI)](https://developer.akamai.com/api/core_features/property_manager/v1.html#createedgehostnames) for more information.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        provider_demo = akamai.EdgeHostName("provider-demo",
            contract_id="ctr_1-AB123",
            edge_hostname="www.example.org.edgesuite.net",
            group_id="grp_123",
            ip_behavior="IPV4",
            product_id="prd_Object_Delivery")
        ```
        ## Attributes reference

        This resource returns this attribute:

        * `ip_behavior` - Returns the IP protocol the hostname will use, either `IPV4` for version 4, IPV6_PERFORMANCE`for version 6, or`IPV6_COMPLIANCE` for both.

        ## Import

        Basic Usagehcl resource "akamai_edge_hostname" "example" {

        # (resource arguments) } You can import Akamai edge hostnames using a comma-delimited string of edge hostname, contract ID, and group ID. You have to enter the values in this order:

        `edge_hostname, contract_id, group_id` For example

        ```sh
         $ pulumi import akamai:properties/edgeHostName:EdgeHostName example ehn_123,ctr_1-AB123,grp_123
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] certificate: Required only when creating an Enhanced TLS edge hostname. This argument sets the certificate enrollment ID. Edge hostnames for Enhanced TLS end in `edgekey.net`. You can retrieve this ID from the [Certificate Provisioning Service CLI](https://github.com/akamai/cli-cps) .
        :param pulumi.Input[str] contract: Replaced by `contract_id`. Maintained for legacy purposes.
        :param pulumi.Input[str] contract_id: A contract's unique ID, including the `ctr_` prefix.
        :param pulumi.Input[str] edge_hostname: One or more edge hostnames. The number of edge hostnames must be less than or equal to the number of public hostnames.
        :param pulumi.Input[str] group: Replaced by `group_id`. Maintained for legacy purposes.
        :param pulumi.Input[str] group_id: A group's unique ID, including the `grp_` prefix.
        :param pulumi.Input[str] ip_behavior: Which version of the IP protocol to use: `IPV4` for version 4 only, `IPV6_PERFORMANCE` for version 6 only, or `IPV6_COMPLIANCE` for both 4 and 6.
        :param pulumi.Input[str] product: Replaced by `product_id`. Maintained for legacy purposes.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] status_update_emails: Email address that should receive updates on the IP behavior update request. Required for update operation.
        :param pulumi.Input[str] use_cases: A JSON encoded list of use cases.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EdgeHostNameArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `EdgeHostName` resource lets you configure a secure edge hostname. Your
        edge hostname determines how requests for your site, app, or content are mapped to
        Akamai edge servers.

        An edge hostname is the CNAME target you use when directing your end user traffic to
        Akamai. Each hostname assigned to a property has a corresponding edge hostname.

        Akamai supports three types of edge hostnames, depending on the level of security
        you need for your traffic: Standard TLS, Enhanced TLS, and Shared Certificate. When
        entering the `edge_hostname` attribute, you need to include a specific domain suffix
        for your edge hostname type:

        | Edge hostname type | Domain suffix |
        |------|-------|
        | Enhanced TLS | edgekey.net |
        | Standard TLS | edgesuite.net |
        | Shared Cert | akamaized.net |

        For example, if you use Standard TLS and have `www.example.com` as a hostname, your edge hostname would be `www.example.com.edgesuite.net`. If you wanted to use Enhanced TLS with the same hostname, your edge hostname would be `www.example.com.edgekey.net`. See the [Property Manager API (PAPI)](https://developer.akamai.com/api/core_features/property_manager/v1.html#createedgehostnames) for more information.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        provider_demo = akamai.EdgeHostName("provider-demo",
            contract_id="ctr_1-AB123",
            edge_hostname="www.example.org.edgesuite.net",
            group_id="grp_123",
            ip_behavior="IPV4",
            product_id="prd_Object_Delivery")
        ```
        ## Attributes reference

        This resource returns this attribute:

        * `ip_behavior` - Returns the IP protocol the hostname will use, either `IPV4` for version 4, IPV6_PERFORMANCE`for version 6, or`IPV6_COMPLIANCE` for both.

        ## Import

        Basic Usagehcl resource "akamai_edge_hostname" "example" {

        # (resource arguments) } You can import Akamai edge hostnames using a comma-delimited string of edge hostname, contract ID, and group ID. You have to enter the values in this order:

        `edge_hostname, contract_id, group_id` For example

        ```sh
         $ pulumi import akamai:properties/edgeHostName:EdgeHostName example ehn_123,ctr_1-AB123,grp_123
        ```

        :param str resource_name: The name of the resource.
        :param EdgeHostNameArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EdgeHostNameArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[int]] = None,
                 contract: Optional[pulumi.Input[str]] = None,
                 contract_id: Optional[pulumi.Input[str]] = None,
                 edge_hostname: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 ip_behavior: Optional[pulumi.Input[str]] = None,
                 product: Optional[pulumi.Input[str]] = None,
                 product_id: Optional[pulumi.Input[str]] = None,
                 status_update_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 use_cases: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""EdgeHostName is deprecated: akamai.properties.EdgeHostName has been deprecated in favor of akamai.EdgeHostName""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EdgeHostNameArgs.__new__(EdgeHostNameArgs)

            __props__.__dict__["certificate"] = certificate
            if contract is not None and not opts.urn:
                warnings.warn("""The setting \"contract\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("""contract is deprecated: The setting \"contract\" has been deprecated.""")
            __props__.__dict__["contract"] = contract
            __props__.__dict__["contract_id"] = contract_id
            if edge_hostname is None and not opts.urn:
                raise TypeError("Missing required property 'edge_hostname'")
            __props__.__dict__["edge_hostname"] = edge_hostname
            if group is not None and not opts.urn:
                warnings.warn("""The setting \"group\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("""group is deprecated: The setting \"group\" has been deprecated.""")
            __props__.__dict__["group"] = group
            __props__.__dict__["group_id"] = group_id
            if ip_behavior is None and not opts.urn:
                raise TypeError("Missing required property 'ip_behavior'")
            __props__.__dict__["ip_behavior"] = ip_behavior
            if product is not None and not opts.urn:
                warnings.warn("""The setting \"product\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("""product is deprecated: The setting \"product\" has been deprecated.""")
            __props__.__dict__["product"] = product
            __props__.__dict__["product_id"] = product_id
            __props__.__dict__["status_update_emails"] = status_update_emails
            __props__.__dict__["use_cases"] = use_cases
        super(EdgeHostName, __self__).__init__(
            'akamai:properties/edgeHostName:EdgeHostName',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            certificate: Optional[pulumi.Input[int]] = None,
            contract: Optional[pulumi.Input[str]] = None,
            contract_id: Optional[pulumi.Input[str]] = None,
            edge_hostname: Optional[pulumi.Input[str]] = None,
            group: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            ip_behavior: Optional[pulumi.Input[str]] = None,
            product: Optional[pulumi.Input[str]] = None,
            product_id: Optional[pulumi.Input[str]] = None,
            status_update_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            use_cases: Optional[pulumi.Input[str]] = None) -> 'EdgeHostName':
        """
        Get an existing EdgeHostName resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] certificate: Required only when creating an Enhanced TLS edge hostname. This argument sets the certificate enrollment ID. Edge hostnames for Enhanced TLS end in `edgekey.net`. You can retrieve this ID from the [Certificate Provisioning Service CLI](https://github.com/akamai/cli-cps) .
        :param pulumi.Input[str] contract: Replaced by `contract_id`. Maintained for legacy purposes.
        :param pulumi.Input[str] contract_id: A contract's unique ID, including the `ctr_` prefix.
        :param pulumi.Input[str] edge_hostname: One or more edge hostnames. The number of edge hostnames must be less than or equal to the number of public hostnames.
        :param pulumi.Input[str] group: Replaced by `group_id`. Maintained for legacy purposes.
        :param pulumi.Input[str] group_id: A group's unique ID, including the `grp_` prefix.
        :param pulumi.Input[str] ip_behavior: Which version of the IP protocol to use: `IPV4` for version 4 only, `IPV6_PERFORMANCE` for version 6 only, or `IPV6_COMPLIANCE` for both 4 and 6.
        :param pulumi.Input[str] product: Replaced by `product_id`. Maintained for legacy purposes.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] status_update_emails: Email address that should receive updates on the IP behavior update request. Required for update operation.
        :param pulumi.Input[str] use_cases: A JSON encoded list of use cases.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EdgeHostNameState.__new__(_EdgeHostNameState)

        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["contract"] = contract
        __props__.__dict__["contract_id"] = contract_id
        __props__.__dict__["edge_hostname"] = edge_hostname
        __props__.__dict__["group"] = group
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["ip_behavior"] = ip_behavior
        __props__.__dict__["product"] = product
        __props__.__dict__["product_id"] = product_id
        __props__.__dict__["status_update_emails"] = status_update_emails
        __props__.__dict__["use_cases"] = use_cases
        return EdgeHostName(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[Optional[int]]:
        """
        Required only when creating an Enhanced TLS edge hostname. This argument sets the certificate enrollment ID. Edge hostnames for Enhanced TLS end in `edgekey.net`. You can retrieve this ID from the [Certificate Provisioning Service CLI](https://github.com/akamai/cli-cps) .
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter
    def contract(self) -> pulumi.Output[str]:
        """
        Replaced by `contract_id`. Maintained for legacy purposes.
        """
        return pulumi.get(self, "contract")

    @property
    @pulumi.getter(name="contractId")
    def contract_id(self) -> pulumi.Output[str]:
        """
        A contract's unique ID, including the `ctr_` prefix.
        """
        return pulumi.get(self, "contract_id")

    @property
    @pulumi.getter(name="edgeHostname")
    def edge_hostname(self) -> pulumi.Output[str]:
        """
        One or more edge hostnames. The number of edge hostnames must be less than or equal to the number of public hostnames.
        """
        return pulumi.get(self, "edge_hostname")

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        """
        Replaced by `group_id`. Maintained for legacy purposes.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        A group's unique ID, including the `grp_` prefix.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="ipBehavior")
    def ip_behavior(self) -> pulumi.Output[str]:
        """
        Which version of the IP protocol to use: `IPV4` for version 4 only, `IPV6_PERFORMANCE` for version 6 only, or `IPV6_COMPLIANCE` for both 4 and 6.
        """
        return pulumi.get(self, "ip_behavior")

    @property
    @pulumi.getter
    def product(self) -> pulumi.Output[str]:
        """
        Replaced by `product_id`. Maintained for legacy purposes.
        """
        return pulumi.get(self, "product")

    @property
    @pulumi.getter(name="productId")
    def product_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "product_id")

    @property
    @pulumi.getter(name="statusUpdateEmails")
    def status_update_emails(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Email address that should receive updates on the IP behavior update request. Required for update operation.
        """
        return pulumi.get(self, "status_update_emails")

    @property
    @pulumi.getter(name="useCases")
    def use_cases(self) -> pulumi.Output[Optional[str]]:
        """
        A JSON encoded list of use cases.
        """
        return pulumi.get(self, "use_cases")

