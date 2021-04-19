# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['GtmCidrmapArgs', 'GtmCidrmap']

@pulumi.input_type
class GtmCidrmapArgs:
    def __init__(__self__, *,
                 default_datacenter: pulumi.Input['GtmCidrmapDefaultDatacenterArgs'],
                 domain: pulumi.Input[str],
                 assignments: Optional[pulumi.Input[Sequence[pulumi.Input['GtmCidrmapAssignmentArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a GtmCidrmap resource.
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
    def default_datacenter(self) -> pulumi.Input['GtmCidrmapDefaultDatacenterArgs']:
        return pulumi.get(self, "default_datacenter")

    @default_datacenter.setter
    def default_datacenter(self, value: pulumi.Input['GtmCidrmapDefaultDatacenterArgs']):
        pulumi.set(self, "default_datacenter", value)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def assignments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GtmCidrmapAssignmentArgs']]]]:
        return pulumi.get(self, "assignments")

    @assignments.setter
    def assignments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GtmCidrmapAssignmentArgs']]]]):
        pulumi.set(self, "assignments", value)

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


@pulumi.input_type
class _GtmCidrmapState:
    def __init__(__self__, *,
                 assignments: Optional[pulumi.Input[Sequence[pulumi.Input['GtmCidrmapAssignmentArgs']]]] = None,
                 default_datacenter: Optional[pulumi.Input['GtmCidrmapDefaultDatacenterArgs']] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering GtmCidrmap resources.
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
    def assignments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GtmCidrmapAssignmentArgs']]]]:
        return pulumi.get(self, "assignments")

    @assignments.setter
    def assignments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GtmCidrmapAssignmentArgs']]]]):
        pulumi.set(self, "assignments", value)

    @property
    @pulumi.getter(name="defaultDatacenter")
    def default_datacenter(self) -> Optional[pulumi.Input['GtmCidrmapDefaultDatacenterArgs']]:
        return pulumi.get(self, "default_datacenter")

    @default_datacenter.setter
    def default_datacenter(self, value: Optional[pulumi.Input['GtmCidrmapDefaultDatacenterArgs']]):
        pulumi.set(self, "default_datacenter", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

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


warnings.warn("""akamai.trafficmanagement.GtmCidrmap has been deprecated in favor of akamai.GtmCidrmap""", DeprecationWarning)


class GtmCidrmap(pulumi.CustomResource):
    warnings.warn("""akamai.trafficmanagement.GtmCidrmap has been deprecated in favor of akamai.GtmCidrmap""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assignments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GtmCidrmapAssignmentArgs']]]]] = None,
                 default_datacenter: Optional[pulumi.Input[pulumi.InputType['GtmCidrmapDefaultDatacenterArgs']]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a GtmCidrmap resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GtmCidrmapArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a GtmCidrmap resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param GtmCidrmapArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GtmCidrmapArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assignments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GtmCidrmapAssignmentArgs']]]]] = None,
                 default_datacenter: Optional[pulumi.Input[pulumi.InputType['GtmCidrmapDefaultDatacenterArgs']]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 wait_on_complete: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        pulumi.log.warn("""GtmCidrmap is deprecated: akamai.trafficmanagement.GtmCidrmap has been deprecated in favor of akamai.GtmCidrmap""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GtmCidrmapArgs.__new__(GtmCidrmapArgs)

            __props__.__dict__["assignments"] = assignments
            if default_datacenter is None and not opts.urn:
                raise TypeError("Missing required property 'default_datacenter'")
            __props__.__dict__["default_datacenter"] = default_datacenter
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            __props__.__dict__["name"] = name
            __props__.__dict__["wait_on_complete"] = wait_on_complete
        super(GtmCidrmap, __self__).__init__(
            'akamai:trafficmanagement/gtmCidrmap:GtmCidrmap',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            assignments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GtmCidrmapAssignmentArgs']]]]] = None,
            default_datacenter: Optional[pulumi.Input[pulumi.InputType['GtmCidrmapDefaultDatacenterArgs']]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            wait_on_complete: Optional[pulumi.Input[bool]] = None) -> 'GtmCidrmap':
        """
        Get an existing GtmCidrmap resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GtmCidrmapState.__new__(_GtmCidrmapState)

        __props__.__dict__["assignments"] = assignments
        __props__.__dict__["default_datacenter"] = default_datacenter
        __props__.__dict__["domain"] = domain
        __props__.__dict__["name"] = name
        __props__.__dict__["wait_on_complete"] = wait_on_complete
        return GtmCidrmap(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def assignments(self) -> pulumi.Output[Optional[Sequence['outputs.GtmCidrmapAssignment']]]:
        return pulumi.get(self, "assignments")

    @property
    @pulumi.getter(name="defaultDatacenter")
    def default_datacenter(self) -> pulumi.Output['outputs.GtmCidrmapDefaultDatacenter']:
        return pulumi.get(self, "default_datacenter")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="waitOnComplete")
    def wait_on_complete(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "wait_on_complete")

