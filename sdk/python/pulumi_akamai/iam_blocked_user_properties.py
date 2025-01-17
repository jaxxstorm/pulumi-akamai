# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IamBlockedUserPropertiesArgs', 'IamBlockedUserProperties']

@pulumi.input_type
class IamBlockedUserPropertiesArgs:
    def __init__(__self__, *,
                 blocked_properties: pulumi.Input[Sequence[pulumi.Input[int]]],
                 group_id: pulumi.Input[int],
                 identity_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a IamBlockedUserProperties resource.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] blocked_properties: List of properties to block for a user. The property IDs must be an integer.
        :param pulumi.Input[int] group_id: A unique identifier for a group. Each identifier must be an integer.
        :param pulumi.Input[str] identity_id: A unique identifier that corresponds to a user's actual profile or client ID. Each identifier must be a string.
        """
        pulumi.set(__self__, "blocked_properties", blocked_properties)
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "identity_id", identity_id)

    @property
    @pulumi.getter(name="blockedProperties")
    def blocked_properties(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        List of properties to block for a user. The property IDs must be an integer.
        """
        return pulumi.get(self, "blocked_properties")

    @blocked_properties.setter
    def blocked_properties(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "blocked_properties", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[int]:
        """
        A unique identifier for a group. Each identifier must be an integer.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="identityId")
    def identity_id(self) -> pulumi.Input[str]:
        """
        A unique identifier that corresponds to a user's actual profile or client ID. Each identifier must be a string.
        """
        return pulumi.get(self, "identity_id")

    @identity_id.setter
    def identity_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "identity_id", value)


@pulumi.input_type
class _IamBlockedUserPropertiesState:
    def __init__(__self__, *,
                 blocked_properties: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 group_id: Optional[pulumi.Input[int]] = None,
                 identity_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IamBlockedUserProperties resources.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] blocked_properties: List of properties to block for a user. The property IDs must be an integer.
        :param pulumi.Input[int] group_id: A unique identifier for a group. Each identifier must be an integer.
        :param pulumi.Input[str] identity_id: A unique identifier that corresponds to a user's actual profile or client ID. Each identifier must be a string.
        """
        if blocked_properties is not None:
            pulumi.set(__self__, "blocked_properties", blocked_properties)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if identity_id is not None:
            pulumi.set(__self__, "identity_id", identity_id)

    @property
    @pulumi.getter(name="blockedProperties")
    def blocked_properties(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        List of properties to block for a user. The property IDs must be an integer.
        """
        return pulumi.get(self, "blocked_properties")

    @blocked_properties.setter
    def blocked_properties(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "blocked_properties", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[int]]:
        """
        A unique identifier for a group. Each identifier must be an integer.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="identityId")
    def identity_id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique identifier that corresponds to a user's actual profile or client ID. Each identifier must be a string.
        """
        return pulumi.get(self, "identity_id")

    @identity_id.setter
    def identity_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_id", value)


class IamBlockedUserProperties(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocked_properties: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 group_id: Optional[pulumi.Input[int]] = None,
                 identity_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use the `IamBlockedUserProperties` resource to remove or grant access to properties. Administrators can block a user's access to any property, overriding any available role already assigned to that user.

        ## Basic usage

        This example returns the policy details based on the policy ID and optionally, a version:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        example = akamai.IamBlockedUserProperties("example",
            blocked_properties=[
                1,
                2,
                3,
                4,
                5,
            ],
            group_id=12345,
            identity_id="A-B-123456")
        ```

        ## Attributes reference

        This resource doesn't return any attributes.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] blocked_properties: List of properties to block for a user. The property IDs must be an integer.
        :param pulumi.Input[int] group_id: A unique identifier for a group. Each identifier must be an integer.
        :param pulumi.Input[str] identity_id: A unique identifier that corresponds to a user's actual profile or client ID. Each identifier must be a string.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IamBlockedUserPropertiesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use the `IamBlockedUserProperties` resource to remove or grant access to properties. Administrators can block a user's access to any property, overriding any available role already assigned to that user.

        ## Basic usage

        This example returns the policy details based on the policy ID and optionally, a version:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        example = akamai.IamBlockedUserProperties("example",
            blocked_properties=[
                1,
                2,
                3,
                4,
                5,
            ],
            group_id=12345,
            identity_id="A-B-123456")
        ```

        ## Attributes reference

        This resource doesn't return any attributes.

        :param str resource_name: The name of the resource.
        :param IamBlockedUserPropertiesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamBlockedUserPropertiesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocked_properties: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 group_id: Optional[pulumi.Input[int]] = None,
                 identity_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamBlockedUserPropertiesArgs.__new__(IamBlockedUserPropertiesArgs)

            if blocked_properties is None and not opts.urn:
                raise TypeError("Missing required property 'blocked_properties'")
            __props__.__dict__["blocked_properties"] = blocked_properties
            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            if identity_id is None and not opts.urn:
                raise TypeError("Missing required property 'identity_id'")
            __props__.__dict__["identity_id"] = identity_id
        super(IamBlockedUserProperties, __self__).__init__(
            'akamai:index/iamBlockedUserProperties:IamBlockedUserProperties',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            blocked_properties: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            group_id: Optional[pulumi.Input[int]] = None,
            identity_id: Optional[pulumi.Input[str]] = None) -> 'IamBlockedUserProperties':
        """
        Get an existing IamBlockedUserProperties resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] blocked_properties: List of properties to block for a user. The property IDs must be an integer.
        :param pulumi.Input[int] group_id: A unique identifier for a group. Each identifier must be an integer.
        :param pulumi.Input[str] identity_id: A unique identifier that corresponds to a user's actual profile or client ID. Each identifier must be a string.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamBlockedUserPropertiesState.__new__(_IamBlockedUserPropertiesState)

        __props__.__dict__["blocked_properties"] = blocked_properties
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["identity_id"] = identity_id
        return IamBlockedUserProperties(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="blockedProperties")
    def blocked_properties(self) -> pulumi.Output[Sequence[int]]:
        """
        List of properties to block for a user. The property IDs must be an integer.
        """
        return pulumi.get(self, "blocked_properties")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[int]:
        """
        A unique identifier for a group. Each identifier must be an integer.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="identityId")
    def identity_id(self) -> pulumi.Output[str]:
        """
        A unique identifier that corresponds to a user's actual profile or client ID. Each identifier must be a string.
        """
        return pulumi.get(self, "identity_id")

