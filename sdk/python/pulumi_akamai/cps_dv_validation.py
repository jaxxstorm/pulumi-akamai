# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['CpsDvValidationArgs', 'CpsDvValidation']

@pulumi.input_type
class CpsDvValidationArgs:
    def __init__(__self__, *,
                 enrollment_id: pulumi.Input[int],
                 sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a CpsDvValidation resource.
        :param pulumi.Input[int] enrollment_id: Unique identifier for the DV certificate enrollment.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] sans: The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
        """
        pulumi.set(__self__, "enrollment_id", enrollment_id)
        if sans is not None:
            pulumi.set(__self__, "sans", sans)

    @property
    @pulumi.getter(name="enrollmentId")
    def enrollment_id(self) -> pulumi.Input[int]:
        """
        Unique identifier for the DV certificate enrollment.
        """
        return pulumi.get(self, "enrollment_id")

    @enrollment_id.setter
    def enrollment_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "enrollment_id", value)

    @property
    @pulumi.getter
    def sans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
        """
        return pulumi.get(self, "sans")

    @sans.setter
    def sans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "sans", value)


@pulumi.input_type
class _CpsDvValidationState:
    def __init__(__self__, *,
                 enrollment_id: Optional[pulumi.Input[int]] = None,
                 sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CpsDvValidation resources.
        :param pulumi.Input[int] enrollment_id: Unique identifier for the DV certificate enrollment.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] sans: The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
        """
        if enrollment_id is not None:
            pulumi.set(__self__, "enrollment_id", enrollment_id)
        if sans is not None:
            pulumi.set(__self__, "sans", sans)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="enrollmentId")
    def enrollment_id(self) -> Optional[pulumi.Input[int]]:
        """
        Unique identifier for the DV certificate enrollment.
        """
        return pulumi.get(self, "enrollment_id")

    @enrollment_id.setter
    def enrollment_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "enrollment_id", value)

    @property
    @pulumi.getter
    def sans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
        """
        return pulumi.get(self, "sans")

    @sans.setter
    def sans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "sans", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class CpsDvValidation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enrollment_id: Optional[pulumi.Input[int]] = None,
                 sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Once you complete the Let's Encrypt challenges, optionally use the `CpsDvValidation` resource to send the acknowledgement to CPS and inform it that tokens are ready for validation. You can also wait for CPS to check for the tokens, which it does on a regular schedule. Next, CPS automatically deploys the certificate on Staging, and eventually on the Production network.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        example = akamai.CpsDvValidation("example",
            enrollment_id=akamai_cps_dv_enrollment["example"]["id"],
            sans=akamai_cps_dv_enrollment["example"]["sans"])
        ```
        ## Attributes reference

        * `status` - The status of certificate validation.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] enrollment_id: Unique identifier for the DV certificate enrollment.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] sans: The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CpsDvValidationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Once you complete the Let's Encrypt challenges, optionally use the `CpsDvValidation` resource to send the acknowledgement to CPS and inform it that tokens are ready for validation. You can also wait for CPS to check for the tokens, which it does on a regular schedule. Next, CPS automatically deploys the certificate on Staging, and eventually on the Production network.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        example = akamai.CpsDvValidation("example",
            enrollment_id=akamai_cps_dv_enrollment["example"]["id"],
            sans=akamai_cps_dv_enrollment["example"]["sans"])
        ```
        ## Attributes reference

        * `status` - The status of certificate validation.

        :param str resource_name: The name of the resource.
        :param CpsDvValidationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CpsDvValidationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enrollment_id: Optional[pulumi.Input[int]] = None,
                 sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CpsDvValidationArgs.__new__(CpsDvValidationArgs)

            if enrollment_id is None and not opts.urn:
                raise TypeError("Missing required property 'enrollment_id'")
            __props__.__dict__["enrollment_id"] = enrollment_id
            __props__.__dict__["sans"] = sans
            __props__.__dict__["status"] = None
        super(CpsDvValidation, __self__).__init__(
            'akamai:index/cpsDvValidation:CpsDvValidation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enrollment_id: Optional[pulumi.Input[int]] = None,
            sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'CpsDvValidation':
        """
        Get an existing CpsDvValidation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] enrollment_id: Unique identifier for the DV certificate enrollment.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] sans: The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CpsDvValidationState.__new__(_CpsDvValidationState)

        __props__.__dict__["enrollment_id"] = enrollment_id
        __props__.__dict__["sans"] = sans
        __props__.__dict__["status"] = status
        return CpsDvValidation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="enrollmentId")
    def enrollment_id(self) -> pulumi.Output[int]:
        """
        Unique identifier for the DV certificate enrollment.
        """
        return pulumi.get(self, "enrollment_id")

    @property
    @pulumi.getter
    def sans(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
        """
        return pulumi.get(self, "sans")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

