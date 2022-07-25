# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['CloudletsApplicationLoadBalancerArgs', 'CloudletsApplicationLoadBalancer']

@pulumi.input_type
class CloudletsApplicationLoadBalancerArgs:
    def __init__(__self__, *,
                 data_centers: pulumi.Input[Sequence[pulumi.Input['CloudletsApplicationLoadBalancerDataCenterArgs']]],
                 origin_id: pulumi.Input[str],
                 balancing_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 liveness_settings: Optional[pulumi.Input['CloudletsApplicationLoadBalancerLivenessSettingsArgs']] = None):
        """
        The set of arguments for constructing a CloudletsApplicationLoadBalancer resource.
        :param pulumi.Input[Sequence[pulumi.Input['CloudletsApplicationLoadBalancerDataCenterArgs']]] data_centers: Specifies the Conditional Origins being used as data centers for an Application Load Balancer implementation. Only Conditional Origins with an origin type of `CUSTOMER` or `NETSTORAGE` can be used as data centers in an Application Load Balancer configuration.
        :param pulumi.Input[str] origin_id: The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
        :param pulumi.Input[str] balancing_type: The type of load balancing being performed, either `WEIGHTED` or `PERFORMANCE`.
        :param pulumi.Input[str] description: The description of the load balancing configuration.
        :param pulumi.Input['CloudletsApplicationLoadBalancerLivenessSettingsArgs'] liveness_settings: Specifies the health of each load balanced data center defined in the data center list.
        """
        pulumi.set(__self__, "data_centers", data_centers)
        pulumi.set(__self__, "origin_id", origin_id)
        if balancing_type is not None:
            pulumi.set(__self__, "balancing_type", balancing_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if liveness_settings is not None:
            pulumi.set(__self__, "liveness_settings", liveness_settings)

    @property
    @pulumi.getter(name="dataCenters")
    def data_centers(self) -> pulumi.Input[Sequence[pulumi.Input['CloudletsApplicationLoadBalancerDataCenterArgs']]]:
        """
        Specifies the Conditional Origins being used as data centers for an Application Load Balancer implementation. Only Conditional Origins with an origin type of `CUSTOMER` or `NETSTORAGE` can be used as data centers in an Application Load Balancer configuration.
        """
        return pulumi.get(self, "data_centers")

    @data_centers.setter
    def data_centers(self, value: pulumi.Input[Sequence[pulumi.Input['CloudletsApplicationLoadBalancerDataCenterArgs']]]):
        pulumi.set(self, "data_centers", value)

    @property
    @pulumi.getter(name="originId")
    def origin_id(self) -> pulumi.Input[str]:
        """
        The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
        """
        return pulumi.get(self, "origin_id")

    @origin_id.setter
    def origin_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "origin_id", value)

    @property
    @pulumi.getter(name="balancingType")
    def balancing_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of load balancing being performed, either `WEIGHTED` or `PERFORMANCE`.
        """
        return pulumi.get(self, "balancing_type")

    @balancing_type.setter
    def balancing_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "balancing_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the load balancing configuration.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="livenessSettings")
    def liveness_settings(self) -> Optional[pulumi.Input['CloudletsApplicationLoadBalancerLivenessSettingsArgs']]:
        """
        Specifies the health of each load balanced data center defined in the data center list.
        """
        return pulumi.get(self, "liveness_settings")

    @liveness_settings.setter
    def liveness_settings(self, value: Optional[pulumi.Input['CloudletsApplicationLoadBalancerLivenessSettingsArgs']]):
        pulumi.set(self, "liveness_settings", value)


@pulumi.input_type
class _CloudletsApplicationLoadBalancerState:
    def __init__(__self__, *,
                 balancing_type: Optional[pulumi.Input[str]] = None,
                 data_centers: Optional[pulumi.Input[Sequence[pulumi.Input['CloudletsApplicationLoadBalancerDataCenterArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 liveness_settings: Optional[pulumi.Input['CloudletsApplicationLoadBalancerLivenessSettingsArgs']] = None,
                 origin_id: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 warnings: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CloudletsApplicationLoadBalancer resources.
        :param pulumi.Input[str] balancing_type: The type of load balancing being performed, either `WEIGHTED` or `PERFORMANCE`.
        :param pulumi.Input[Sequence[pulumi.Input['CloudletsApplicationLoadBalancerDataCenterArgs']]] data_centers: Specifies the Conditional Origins being used as data centers for an Application Load Balancer implementation. Only Conditional Origins with an origin type of `CUSTOMER` or `NETSTORAGE` can be used as data centers in an Application Load Balancer configuration.
        :param pulumi.Input[str] description: The description of the load balancing configuration.
        :param pulumi.Input['CloudletsApplicationLoadBalancerLivenessSettingsArgs'] liveness_settings: Specifies the health of each load balanced data center defined in the data center list.
        :param pulumi.Input[str] origin_id: The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
        :param pulumi.Input[int] version: The version number of the load balancing configuration.
        :param pulumi.Input[str] warnings: A list of warnings that occurred during the activation of the load balancing configuration.
        """
        if balancing_type is not None:
            pulumi.set(__self__, "balancing_type", balancing_type)
        if data_centers is not None:
            pulumi.set(__self__, "data_centers", data_centers)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if liveness_settings is not None:
            pulumi.set(__self__, "liveness_settings", liveness_settings)
        if origin_id is not None:
            pulumi.set(__self__, "origin_id", origin_id)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if warnings is not None:
            pulumi.set(__self__, "warnings", warnings)

    @property
    @pulumi.getter(name="balancingType")
    def balancing_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of load balancing being performed, either `WEIGHTED` or `PERFORMANCE`.
        """
        return pulumi.get(self, "balancing_type")

    @balancing_type.setter
    def balancing_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "balancing_type", value)

    @property
    @pulumi.getter(name="dataCenters")
    def data_centers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CloudletsApplicationLoadBalancerDataCenterArgs']]]]:
        """
        Specifies the Conditional Origins being used as data centers for an Application Load Balancer implementation. Only Conditional Origins with an origin type of `CUSTOMER` or `NETSTORAGE` can be used as data centers in an Application Load Balancer configuration.
        """
        return pulumi.get(self, "data_centers")

    @data_centers.setter
    def data_centers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CloudletsApplicationLoadBalancerDataCenterArgs']]]]):
        pulumi.set(self, "data_centers", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the load balancing configuration.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="livenessSettings")
    def liveness_settings(self) -> Optional[pulumi.Input['CloudletsApplicationLoadBalancerLivenessSettingsArgs']]:
        """
        Specifies the health of each load balanced data center defined in the data center list.
        """
        return pulumi.get(self, "liveness_settings")

    @liveness_settings.setter
    def liveness_settings(self, value: Optional[pulumi.Input['CloudletsApplicationLoadBalancerLivenessSettingsArgs']]):
        pulumi.set(self, "liveness_settings", value)

    @property
    @pulumi.getter(name="originId")
    def origin_id(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
        """
        return pulumi.get(self, "origin_id")

    @origin_id.setter
    def origin_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "origin_id", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        The version number of the load balancing configuration.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter
    def warnings(self) -> Optional[pulumi.Input[str]]:
        """
        A list of warnings that occurred during the activation of the load balancing configuration.
        """
        return pulumi.get(self, "warnings")

    @warnings.setter
    def warnings(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "warnings", value)


class CloudletsApplicationLoadBalancer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 balancing_type: Optional[pulumi.Input[str]] = None,
                 data_centers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudletsApplicationLoadBalancerDataCenterArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 liveness_settings: Optional[pulumi.Input[pulumi.InputType['CloudletsApplicationLoadBalancerLivenessSettingsArgs']]] = None,
                 origin_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use the `CloudletsApplicationLoadBalancer` resource to create the Application Load Balancer Cloudlet configuration. The Application Load Balancer Cloudlet provides intelligent, scalable traffic management across physical, virtual, and cloud-hosted data centers without requiring the origin to send load feedback. This Cloudlet can automatically detect load conditions and route traffic to the optimal data source while maintaining custom routing policies and consistent visitor session behavior for your visitors.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        example = akamai.CloudletsApplicationLoadBalancer("example",
            balancing_type="WEIGHTED",
            data_centers=[akamai.CloudletsApplicationLoadBalancerDataCenterArgs(
                city="Boston",
                cloud_server_host_header_override=False,
                cloud_service=True,
                continent="NA",
                country="US",
                hostname="example-hostname",
                latitude=102.78108,
                liveness_hosts=["example"],
                longitude=-116.07064,
                origin_id="alb_test_1",
                percent=100,
                state_or_province="MA",
            )],
            description="application_load_balancer description",
            liveness_settings=akamai.CloudletsApplicationLoadBalancerLivenessSettingsArgs(
                additional_headers={
                    "additionalHeaders": "123",
                },
                host_header="header",
                interval=10,
                path="/status",
                port=1234,
                protocol="HTTP",
                request_string="test_request_string",
                response_string="test_response_string",
                timeout=60,
            ),
            origin_id="alb_test_1")
        ```

        ## Import

        Basic usagehcl resource "akamai_cloudlets_application_load_balancer" "example" {

        # (resource arguments)

         } You can import your Akamai Application Load Balancer configuration using an origin ID. For example

        ```sh
         $ pulumi import akamai:index/cloudletsApplicationLoadBalancer:CloudletsApplicationLoadBalancer example alb_test_1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] balancing_type: The type of load balancing being performed, either `WEIGHTED` or `PERFORMANCE`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudletsApplicationLoadBalancerDataCenterArgs']]]] data_centers: Specifies the Conditional Origins being used as data centers for an Application Load Balancer implementation. Only Conditional Origins with an origin type of `CUSTOMER` or `NETSTORAGE` can be used as data centers in an Application Load Balancer configuration.
        :param pulumi.Input[str] description: The description of the load balancing configuration.
        :param pulumi.Input[pulumi.InputType['CloudletsApplicationLoadBalancerLivenessSettingsArgs']] liveness_settings: Specifies the health of each load balanced data center defined in the data center list.
        :param pulumi.Input[str] origin_id: The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CloudletsApplicationLoadBalancerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use the `CloudletsApplicationLoadBalancer` resource to create the Application Load Balancer Cloudlet configuration. The Application Load Balancer Cloudlet provides intelligent, scalable traffic management across physical, virtual, and cloud-hosted data centers without requiring the origin to send load feedback. This Cloudlet can automatically detect load conditions and route traffic to the optimal data source while maintaining custom routing policies and consistent visitor session behavior for your visitors.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        example = akamai.CloudletsApplicationLoadBalancer("example",
            balancing_type="WEIGHTED",
            data_centers=[akamai.CloudletsApplicationLoadBalancerDataCenterArgs(
                city="Boston",
                cloud_server_host_header_override=False,
                cloud_service=True,
                continent="NA",
                country="US",
                hostname="example-hostname",
                latitude=102.78108,
                liveness_hosts=["example"],
                longitude=-116.07064,
                origin_id="alb_test_1",
                percent=100,
                state_or_province="MA",
            )],
            description="application_load_balancer description",
            liveness_settings=akamai.CloudletsApplicationLoadBalancerLivenessSettingsArgs(
                additional_headers={
                    "additionalHeaders": "123",
                },
                host_header="header",
                interval=10,
                path="/status",
                port=1234,
                protocol="HTTP",
                request_string="test_request_string",
                response_string="test_response_string",
                timeout=60,
            ),
            origin_id="alb_test_1")
        ```

        ## Import

        Basic usagehcl resource "akamai_cloudlets_application_load_balancer" "example" {

        # (resource arguments)

         } You can import your Akamai Application Load Balancer configuration using an origin ID. For example

        ```sh
         $ pulumi import akamai:index/cloudletsApplicationLoadBalancer:CloudletsApplicationLoadBalancer example alb_test_1
        ```

        :param str resource_name: The name of the resource.
        :param CloudletsApplicationLoadBalancerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CloudletsApplicationLoadBalancerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 balancing_type: Optional[pulumi.Input[str]] = None,
                 data_centers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudletsApplicationLoadBalancerDataCenterArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 liveness_settings: Optional[pulumi.Input[pulumi.InputType['CloudletsApplicationLoadBalancerLivenessSettingsArgs']]] = None,
                 origin_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CloudletsApplicationLoadBalancerArgs.__new__(CloudletsApplicationLoadBalancerArgs)

            __props__.__dict__["balancing_type"] = balancing_type
            if data_centers is None and not opts.urn:
                raise TypeError("Missing required property 'data_centers'")
            __props__.__dict__["data_centers"] = data_centers
            __props__.__dict__["description"] = description
            __props__.__dict__["liveness_settings"] = liveness_settings
            if origin_id is None and not opts.urn:
                raise TypeError("Missing required property 'origin_id'")
            __props__.__dict__["origin_id"] = origin_id
            __props__.__dict__["version"] = None
            __props__.__dict__["warnings"] = None
        super(CloudletsApplicationLoadBalancer, __self__).__init__(
            'akamai:index/cloudletsApplicationLoadBalancer:CloudletsApplicationLoadBalancer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            balancing_type: Optional[pulumi.Input[str]] = None,
            data_centers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudletsApplicationLoadBalancerDataCenterArgs']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            liveness_settings: Optional[pulumi.Input[pulumi.InputType['CloudletsApplicationLoadBalancerLivenessSettingsArgs']]] = None,
            origin_id: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None,
            warnings: Optional[pulumi.Input[str]] = None) -> 'CloudletsApplicationLoadBalancer':
        """
        Get an existing CloudletsApplicationLoadBalancer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] balancing_type: The type of load balancing being performed, either `WEIGHTED` or `PERFORMANCE`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudletsApplicationLoadBalancerDataCenterArgs']]]] data_centers: Specifies the Conditional Origins being used as data centers for an Application Load Balancer implementation. Only Conditional Origins with an origin type of `CUSTOMER` or `NETSTORAGE` can be used as data centers in an Application Load Balancer configuration.
        :param pulumi.Input[str] description: The description of the load balancing configuration.
        :param pulumi.Input[pulumi.InputType['CloudletsApplicationLoadBalancerLivenessSettingsArgs']] liveness_settings: Specifies the health of each load balanced data center defined in the data center list.
        :param pulumi.Input[str] origin_id: The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
        :param pulumi.Input[int] version: The version number of the load balancing configuration.
        :param pulumi.Input[str] warnings: A list of warnings that occurred during the activation of the load balancing configuration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CloudletsApplicationLoadBalancerState.__new__(_CloudletsApplicationLoadBalancerState)

        __props__.__dict__["balancing_type"] = balancing_type
        __props__.__dict__["data_centers"] = data_centers
        __props__.__dict__["description"] = description
        __props__.__dict__["liveness_settings"] = liveness_settings
        __props__.__dict__["origin_id"] = origin_id
        __props__.__dict__["version"] = version
        __props__.__dict__["warnings"] = warnings
        return CloudletsApplicationLoadBalancer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="balancingType")
    def balancing_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of load balancing being performed, either `WEIGHTED` or `PERFORMANCE`.
        """
        return pulumi.get(self, "balancing_type")

    @property
    @pulumi.getter(name="dataCenters")
    def data_centers(self) -> pulumi.Output[Sequence['outputs.CloudletsApplicationLoadBalancerDataCenter']]:
        """
        Specifies the Conditional Origins being used as data centers for an Application Load Balancer implementation. Only Conditional Origins with an origin type of `CUSTOMER` or `NETSTORAGE` can be used as data centers in an Application Load Balancer configuration.
        """
        return pulumi.get(self, "data_centers")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the load balancing configuration.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="livenessSettings")
    def liveness_settings(self) -> pulumi.Output[Optional['outputs.CloudletsApplicationLoadBalancerLivenessSettings']]:
        """
        Specifies the health of each load balanced data center defined in the data center list.
        """
        return pulumi.get(self, "liveness_settings")

    @property
    @pulumi.getter(name="originId")
    def origin_id(self) -> pulumi.Output[str]:
        """
        The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
        """
        return pulumi.get(self, "origin_id")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        The version number of the load balancing configuration.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter
    def warnings(self) -> pulumi.Output[str]:
        """
        A list of warnings that occurred during the activation of the load balancing configuration.
        """
        return pulumi.get(self, "warnings")

