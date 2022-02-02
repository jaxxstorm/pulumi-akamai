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

__all__ = ['GtmAsmapArgs', 'GtmAsmap']

@pulumi.input_type
class GtmAsmapArgs:
    def __init__(__self__, *,
                 default_datacenter: pulumi.Input['GtmAsmapDefaultDatacenterArgs'],
                 domain: pulumi.Input[str],
                 assignments: Optional[pulumi.Input[Sequence[pulumi.Input['GtmAsmapAssignmentArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a GtmAsmap resource.
        :param pulumi.Input['GtmAsmapDefaultDatacenterArgs'] default_datacenter: A placeholder for all other AS zones not found in these AS zones. Requires these additional arguments:
        :param pulumi.Input[str] domain: The GTM Domain name for the AS map.
        :param pulumi.Input[Sequence[pulumi.Input['GtmAsmapAssignmentArgs']]] assignments: Contains information about the AS zone groupings of AS IDs. You can have multiple entries with this argument. If used, requires these arguments:
        :param pulumi.Input[str] name: A descriptive label for the AS map. Properties set up for  AS mapping can use this as reference.
        :param pulumi.Input[bool] wait_on_complete: A boolean that, if `true`, waits for transaction to complete.
        """
        pulumi.set(__self__, "default_datacenter", default_datacenter)
        pulumi.set(__self__, "domain", domain)
        if assignments is not None:
            pulumi.set(__self__, "assignments", assignments)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if wait_on_complete is not None:
            pulumi.set(__self__, "wait_on_complete", wait_on_complete)

    @property
    @pulumi.getter(name="defaultDatacenter")
    def default_datacenter(self) -> pulumi.Input['GtmAsmapDefaultDatacenterArgs']:
        """
        A placeholder for all other AS zones not found in these AS zones. Requires these additional arguments:
        """
        return pulumi.get(self, "default_datacenter")

    @default_datacenter.setter
    def default_datacenter(self, value: pulumi.Input['GtmAsmapDefaultDatacenterArgs']):
        pulumi.set(self, "default_datacenter", value)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        The GTM Domain name for the AS map.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def assignments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GtmAsmapAssignmentArgs']]]]:
        """
        Contains information about the AS zone groupings of AS IDs. You can have multiple entries with this argument. If used, requires these arguments:
        """
        return pulumi.get(self, "assignments")

    @assignments.setter
    def assignments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GtmAsmapAssignmentArgs']]]]):
        pulumi.set(self, "assignments", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A descriptive label for the AS map. Properties set up for  AS mapping can use this as reference.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="waitOnComplete")
    def wait_on_complete(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean that, if `true`, waits for transaction to complete.
        """
        return pulumi.get(self, "wait_on_complete")

    @wait_on_complete.setter
    def wait_on_complete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_on_complete", value)


@pulumi.input_type
class _GtmAsmapState:
    def __init__(__self__, *,
                 assignments: Optional[pulumi.Input[Sequence[pulumi.Input['GtmAsmapAssignmentArgs']]]] = None,
                 default_datacenter: Optional[pulumi.Input['GtmAsmapDefaultDatacenterArgs']] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering GtmAsmap resources.
        :param pulumi.Input[Sequence[pulumi.Input['GtmAsmapAssignmentArgs']]] assignments: Contains information about the AS zone groupings of AS IDs. You can have multiple entries with this argument. If used, requires these arguments:
        :param pulumi.Input['GtmAsmapDefaultDatacenterArgs'] default_datacenter: A placeholder for all other AS zones not found in these AS zones. Requires these additional arguments:
        :param pulumi.Input[str] domain: The GTM Domain name for the AS map.
        :param pulumi.Input[str] name: A descriptive label for the AS map. Properties set up for  AS mapping can use this as reference.
        :param pulumi.Input[bool] wait_on_complete: A boolean that, if `true`, waits for transaction to complete.
        """
        if assignments is not None:
            pulumi.set(__self__, "assignments", assignments)
        if default_datacenter is not None:
            pulumi.set(__self__, "default_datacenter", default_datacenter)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if wait_on_complete is not None:
            pulumi.set(__self__, "wait_on_complete", wait_on_complete)

    @property
    @pulumi.getter
    def assignments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GtmAsmapAssignmentArgs']]]]:
        """
        Contains information about the AS zone groupings of AS IDs. You can have multiple entries with this argument. If used, requires these arguments:
        """
        return pulumi.get(self, "assignments")

    @assignments.setter
    def assignments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GtmAsmapAssignmentArgs']]]]):
        pulumi.set(self, "assignments", value)

    @property
    @pulumi.getter(name="defaultDatacenter")
    def default_datacenter(self) -> Optional[pulumi.Input['GtmAsmapDefaultDatacenterArgs']]:
        """
        A placeholder for all other AS zones not found in these AS zones. Requires these additional arguments:
        """
        return pulumi.get(self, "default_datacenter")

    @default_datacenter.setter
    def default_datacenter(self, value: Optional[pulumi.Input['GtmAsmapDefaultDatacenterArgs']]):
        pulumi.set(self, "default_datacenter", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The GTM Domain name for the AS map.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A descriptive label for the AS map. Properties set up for  AS mapping can use this as reference.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="waitOnComplete")
    def wait_on_complete(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean that, if `true`, waits for transaction to complete.
        """
        return pulumi.get(self, "wait_on_complete")

    @wait_on_complete.setter
    def wait_on_complete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_on_complete", value)


class GtmAsmap(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assignments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GtmAsmapAssignmentArgs']]]]] = None,
                 default_datacenter: Optional[pulumi.Input[pulumi.InputType['GtmAsmapDefaultDatacenterArgs']]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Use the `GtmAsmap` resource to create, configure, and import a GTM Autonomous System (AS) map. AS mapping lets you configure a GTM property that returns a CNAME based on the AS number associated with the requester's IP address.

        You can reuse maps for multiple properties or create new ones. AS maps split the Internet into multiple AS block zones. Properties that use AS maps can specify handout integers for each zone. AS mapping lets you configure a property that directs users to a specific environment or to the origin.

        > **Note** Import requires an ID with this format: `existing_domain_name`:`existing_map_name`.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        demo_asmap = akamai.GtmAsmap("demoAsmap",
            default_datacenter=akamai.GtmAsmapDefaultDatacenterArgs(
                datacenter_id=5400,
                nickname="All Other AS numbers",
            ),
            domain="demo_domain.akadns.net")
        ```
        ## Schema reference

        You can download the GTM AS Map backing schema from the [Global Traffic Management API](https://developer.akamai.com/api/web_performance/global_traffic_management/v1.html#asmap) page.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GtmAsmapAssignmentArgs']]]] assignments: Contains information about the AS zone groupings of AS IDs. You can have multiple entries with this argument. If used, requires these arguments:
        :param pulumi.Input[pulumi.InputType['GtmAsmapDefaultDatacenterArgs']] default_datacenter: A placeholder for all other AS zones not found in these AS zones. Requires these additional arguments:
        :param pulumi.Input[str] domain: The GTM Domain name for the AS map.
        :param pulumi.Input[str] name: A descriptive label for the AS map. Properties set up for  AS mapping can use this as reference.
        :param pulumi.Input[bool] wait_on_complete: A boolean that, if `true`, waits for transaction to complete.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GtmAsmapArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use the `GtmAsmap` resource to create, configure, and import a GTM Autonomous System (AS) map. AS mapping lets you configure a GTM property that returns a CNAME based on the AS number associated with the requester's IP address.

        You can reuse maps for multiple properties or create new ones. AS maps split the Internet into multiple AS block zones. Properties that use AS maps can specify handout integers for each zone. AS mapping lets you configure a property that directs users to a specific environment or to the origin.

        > **Note** Import requires an ID with this format: `existing_domain_name`:`existing_map_name`.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        demo_asmap = akamai.GtmAsmap("demoAsmap",
            default_datacenter=akamai.GtmAsmapDefaultDatacenterArgs(
                datacenter_id=5400,
                nickname="All Other AS numbers",
            ),
            domain="demo_domain.akadns.net")
        ```
        ## Schema reference

        You can download the GTM AS Map backing schema from the [Global Traffic Management API](https://developer.akamai.com/api/web_performance/global_traffic_management/v1.html#asmap) page.

        :param str resource_name: The name of the resource.
        :param GtmAsmapArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GtmAsmapArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assignments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GtmAsmapAssignmentArgs']]]]] = None,
                 default_datacenter: Optional[pulumi.Input[pulumi.InputType['GtmAsmapDefaultDatacenterArgs']]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None,
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
            __props__ = GtmAsmapArgs.__new__(GtmAsmapArgs)

            __props__.__dict__["assignments"] = assignments
            if default_datacenter is None and not opts.urn:
                raise TypeError("Missing required property 'default_datacenter'")
            __props__.__dict__["default_datacenter"] = default_datacenter
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            __props__.__dict__["name"] = name
            __props__.__dict__["wait_on_complete"] = wait_on_complete
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="akamai:trafficmanagement/gtmASmap:GtmASmap")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(GtmAsmap, __self__).__init__(
            'akamai:index/gtmAsmap:GtmAsmap',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            assignments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GtmAsmapAssignmentArgs']]]]] = None,
            default_datacenter: Optional[pulumi.Input[pulumi.InputType['GtmAsmapDefaultDatacenterArgs']]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            wait_on_complete: Optional[pulumi.Input[bool]] = None) -> 'GtmAsmap':
        """
        Get an existing GtmAsmap resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GtmAsmapAssignmentArgs']]]] assignments: Contains information about the AS zone groupings of AS IDs. You can have multiple entries with this argument. If used, requires these arguments:
        :param pulumi.Input[pulumi.InputType['GtmAsmapDefaultDatacenterArgs']] default_datacenter: A placeholder for all other AS zones not found in these AS zones. Requires these additional arguments:
        :param pulumi.Input[str] domain: The GTM Domain name for the AS map.
        :param pulumi.Input[str] name: A descriptive label for the AS map. Properties set up for  AS mapping can use this as reference.
        :param pulumi.Input[bool] wait_on_complete: A boolean that, if `true`, waits for transaction to complete.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GtmAsmapState.__new__(_GtmAsmapState)

        __props__.__dict__["assignments"] = assignments
        __props__.__dict__["default_datacenter"] = default_datacenter
        __props__.__dict__["domain"] = domain
        __props__.__dict__["name"] = name
        __props__.__dict__["wait_on_complete"] = wait_on_complete
        return GtmAsmap(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def assignments(self) -> pulumi.Output[Optional[Sequence['outputs.GtmAsmapAssignment']]]:
        """
        Contains information about the AS zone groupings of AS IDs. You can have multiple entries with this argument. If used, requires these arguments:
        """
        return pulumi.get(self, "assignments")

    @property
    @pulumi.getter(name="defaultDatacenter")
    def default_datacenter(self) -> pulumi.Output['outputs.GtmAsmapDefaultDatacenter']:
        """
        A placeholder for all other AS zones not found in these AS zones. Requires these additional arguments:
        """
        return pulumi.get(self, "default_datacenter")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        The GTM Domain name for the AS map.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A descriptive label for the AS map. Properties set up for  AS mapping can use this as reference.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="waitOnComplete")
    def wait_on_complete(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean that, if `true`, waits for transaction to complete.
        """
        return pulumi.get(self, "wait_on_complete")

