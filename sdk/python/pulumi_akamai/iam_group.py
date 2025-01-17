# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IamGroupArgs', 'IamGroup']

@pulumi.input_type
class IamGroupArgs:
    def __init__(__self__, *,
                 parent_group_id: pulumi.Input[int],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IamGroup resource.
        :param pulumi.Input[int] parent_group_id: A unique identifier for the parent group. Each identifier must be an integer.
        :param pulumi.Input[str] name: Human readable name for a group.
        """
        pulumi.set(__self__, "parent_group_id", parent_group_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="parentGroupId")
    def parent_group_id(self) -> pulumi.Input[int]:
        """
        A unique identifier for the parent group. Each identifier must be an integer.
        """
        return pulumi.get(self, "parent_group_id")

    @parent_group_id.setter
    def parent_group_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "parent_group_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Human readable name for a group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _IamGroupState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_group_id: Optional[pulumi.Input[int]] = None,
                 sub_groups: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None):
        """
        Input properties used for looking up and filtering IamGroup resources.
        :param pulumi.Input[str] name: Human readable name for a group.
        :param pulumi.Input[int] parent_group_id: A unique identifier for the parent group. Each identifier must be an integer.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] sub_groups: Subgroups IDs
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_group_id is not None:
            pulumi.set(__self__, "parent_group_id", parent_group_id)
        if sub_groups is not None:
            pulumi.set(__self__, "sub_groups", sub_groups)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Human readable name for a group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentGroupId")
    def parent_group_id(self) -> Optional[pulumi.Input[int]]:
        """
        A unique identifier for the parent group. Each identifier must be an integer.
        """
        return pulumi.get(self, "parent_group_id")

    @parent_group_id.setter
    def parent_group_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "parent_group_id", value)

    @property
    @pulumi.getter(name="subGroups")
    def sub_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        Subgroups IDs
        """
        return pulumi.get(self, "sub_groups")

    @sub_groups.setter
    def sub_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "sub_groups", value)


class IamGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_group_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Use the `IamGroup` resource to list details about groups. Groups are organizational containers for the objects you use.  Groups can contain other groups, primary objects like properties, and secondary objects like edge hostnames or content provider (CP) codes.

        ## Basic usage

        This example returns the policy details based on the policy ID and optionally, a version:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        example = akamai.IamGroup("example", parent_group_id=12345)
        ```

        ## Attributes reference

        This resource returns this attribute:

        * `sub_groups` - Sub-groups that are related to this group. Each identifier must be an integer.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Human readable name for a group.
        :param pulumi.Input[int] parent_group_id: A unique identifier for the parent group. Each identifier must be an integer.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IamGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use the `IamGroup` resource to list details about groups. Groups are organizational containers for the objects you use.  Groups can contain other groups, primary objects like properties, and secondary objects like edge hostnames or content provider (CP) codes.

        ## Basic usage

        This example returns the policy details based on the policy ID and optionally, a version:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        example = akamai.IamGroup("example", parent_group_id=12345)
        ```

        ## Attributes reference

        This resource returns this attribute:

        * `sub_groups` - Sub-groups that are related to this group. Each identifier must be an integer.

        :param str resource_name: The name of the resource.
        :param IamGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_group_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamGroupArgs.__new__(IamGroupArgs)

            __props__.__dict__["name"] = name
            if parent_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'parent_group_id'")
            __props__.__dict__["parent_group_id"] = parent_group_id
            __props__.__dict__["sub_groups"] = None
        super(IamGroup, __self__).__init__(
            'akamai:index/iamGroup:IamGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent_group_id: Optional[pulumi.Input[int]] = None,
            sub_groups: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None) -> 'IamGroup':
        """
        Get an existing IamGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Human readable name for a group.
        :param pulumi.Input[int] parent_group_id: A unique identifier for the parent group. Each identifier must be an integer.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] sub_groups: Subgroups IDs
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamGroupState.__new__(_IamGroupState)

        __props__.__dict__["name"] = name
        __props__.__dict__["parent_group_id"] = parent_group_id
        __props__.__dict__["sub_groups"] = sub_groups
        return IamGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Human readable name for a group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentGroupId")
    def parent_group_id(self) -> pulumi.Output[int]:
        """
        A unique identifier for the parent group. Each identifier must be an integer.
        """
        return pulumi.get(self, "parent_group_id")

    @property
    @pulumi.getter(name="subGroups")
    def sub_groups(self) -> pulumi.Output[Sequence[int]]:
        """
        Subgroups IDs
        """
        return pulumi.get(self, "sub_groups")

