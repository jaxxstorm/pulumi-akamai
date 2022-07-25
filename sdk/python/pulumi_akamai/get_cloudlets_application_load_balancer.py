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

__all__ = [
    'GetCloudletsApplicationLoadBalancerResult',
    'AwaitableGetCloudletsApplicationLoadBalancerResult',
    'get_cloudlets_application_load_balancer',
    'get_cloudlets_application_load_balancer_output',
]

@pulumi.output_type
class GetCloudletsApplicationLoadBalancerResult:
    """
    A collection of values returned by getCloudletsApplicationLoadBalancer.
    """
    def __init__(__self__, balancing_type=None, created_by=None, created_date=None, data_centers=None, deleted=None, description=None, id=None, immutable=None, last_modified_by=None, last_modified_date=None, liveness_settings=None, origin_id=None, type=None, version=None, warnings=None):
        if balancing_type and not isinstance(balancing_type, str):
            raise TypeError("Expected argument 'balancing_type' to be a str")
        pulumi.set(__self__, "balancing_type", balancing_type)
        if created_by and not isinstance(created_by, str):
            raise TypeError("Expected argument 'created_by' to be a str")
        pulumi.set(__self__, "created_by", created_by)
        if created_date and not isinstance(created_date, str):
            raise TypeError("Expected argument 'created_date' to be a str")
        pulumi.set(__self__, "created_date", created_date)
        if data_centers and not isinstance(data_centers, list):
            raise TypeError("Expected argument 'data_centers' to be a list")
        pulumi.set(__self__, "data_centers", data_centers)
        if deleted and not isinstance(deleted, bool):
            raise TypeError("Expected argument 'deleted' to be a bool")
        pulumi.set(__self__, "deleted", deleted)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if immutable and not isinstance(immutable, bool):
            raise TypeError("Expected argument 'immutable' to be a bool")
        pulumi.set(__self__, "immutable", immutable)
        if last_modified_by and not isinstance(last_modified_by, str):
            raise TypeError("Expected argument 'last_modified_by' to be a str")
        pulumi.set(__self__, "last_modified_by", last_modified_by)
        if last_modified_date and not isinstance(last_modified_date, str):
            raise TypeError("Expected argument 'last_modified_date' to be a str")
        pulumi.set(__self__, "last_modified_date", last_modified_date)
        if liveness_settings and not isinstance(liveness_settings, list):
            raise TypeError("Expected argument 'liveness_settings' to be a list")
        pulumi.set(__self__, "liveness_settings", liveness_settings)
        if origin_id and not isinstance(origin_id, str):
            raise TypeError("Expected argument 'origin_id' to be a str")
        pulumi.set(__self__, "origin_id", origin_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if version and not isinstance(version, int):
            raise TypeError("Expected argument 'version' to be a int")
        pulumi.set(__self__, "version", version)
        if warnings and not isinstance(warnings, str):
            raise TypeError("Expected argument 'warnings' to be a str")
        pulumi.set(__self__, "warnings", warnings)

    @property
    @pulumi.getter(name="balancingType")
    def balancing_type(self) -> str:
        return pulumi.get(self, "balancing_type")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> str:
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> str:
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter(name="dataCenters")
    def data_centers(self) -> Sequence['outputs.GetCloudletsApplicationLoadBalancerDataCenterResult']:
        return pulumi.get(self, "data_centers")

    @property
    @pulumi.getter
    def deleted(self) -> bool:
        return pulumi.get(self, "deleted")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def immutable(self) -> bool:
        return pulumi.get(self, "immutable")

    @property
    @pulumi.getter(name="lastModifiedBy")
    def last_modified_by(self) -> str:
        return pulumi.get(self, "last_modified_by")

    @property
    @pulumi.getter(name="lastModifiedDate")
    def last_modified_date(self) -> str:
        return pulumi.get(self, "last_modified_date")

    @property
    @pulumi.getter(name="livenessSettings")
    def liveness_settings(self) -> Sequence['outputs.GetCloudletsApplicationLoadBalancerLivenessSettingResult']:
        return pulumi.get(self, "liveness_settings")

    @property
    @pulumi.getter(name="originId")
    def origin_id(self) -> str:
        return pulumi.get(self, "origin_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> Optional[int]:
        return pulumi.get(self, "version")

    @property
    @pulumi.getter
    def warnings(self) -> str:
        return pulumi.get(self, "warnings")


class AwaitableGetCloudletsApplicationLoadBalancerResult(GetCloudletsApplicationLoadBalancerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudletsApplicationLoadBalancerResult(
            balancing_type=self.balancing_type,
            created_by=self.created_by,
            created_date=self.created_date,
            data_centers=self.data_centers,
            deleted=self.deleted,
            description=self.description,
            id=self.id,
            immutable=self.immutable,
            last_modified_by=self.last_modified_by,
            last_modified_date=self.last_modified_date,
            liveness_settings=self.liveness_settings,
            origin_id=self.origin_id,
            type=self.type,
            version=self.version,
            warnings=self.warnings)


def get_cloudlets_application_load_balancer(origin_id: Optional[str] = None,
                                            version: Optional[int] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudletsApplicationLoadBalancerResult:
    """
    Use the `CloudletsApplicationLoadBalancer` data source to list details about the Application Load Balancer configuration with a specified policy version, or latest if not specified.

    ## Basic usage

    This example returns the load balancing configuration details based on the origin ID and optionally, a version:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    example = akamai.get_cloudlets_application_load_balancer(origin_id="alb_test_1",
        version=1)
    ```

    ## Attributes reference

    This data source returns these attributes:

    * `description` - The description of the load balancing configuration.
    * `type` - The type of Conditional Origin. `APPLICATION_LOAD_BALANCER` is the only supported value.
    * `balancing_type` - The type of load balancing being performed, either `WEIGHTED` or `PERFORMANCE`.
    * `created_by` - The name of the user who created this load balancing configuration.
    * `created_date` - The date, in ISO 8601 format, when this load balancing configuration was created.
    * `deleted` - Whether the Conditional Origin version has been deleted. If `false`, you can use this version again.
    * `immutable` - Whether you can edit the load balancing version. The default setting for this member is false. It automatically becomes true when the load balancing version is activated for the first time.
    * `last_modified_by` - The user who last modified the load balancing configuration.
    * `last_modified_date` - The date, in ISO 8601 format, when the initial load balancing configuration was last modified.
    * `warnings` - A list of warnings that occured during the activation of the load balancing configuration.
    * `data_centers` - Specifies the Conditional Origins being used as data centers for an Application Load Balancer implementation. Only Conditional Origins with an origin type of `CUSTOMER` or `NETSTORAGE` can be used as data centers in an Application Load Balancer configuration.
      * `city` - The city in which the data center is located.
      * `cloud_server_host_header_override` - Whether the cloud server host header is overridden.
      * `cloud_service` - Whether this datacenter is a cloud service.
      * `continent` - The code of the continent on which the data center is located. See [Continent Codes](https://control.akamai.com/dl/edgescape/continentCodes.csv) for a list of valid codes.
      * `country` - The country in which the data center is located. See [Country Codes](https://control.akamai.com/dl/edgescape/cc2continent.csv) for a list of valid codes.
      * `hostname` - The name of the host that can be used as a Conditional Origin. This should match the `hostname` value defined for this datacenter in Property Manager.
      * `latitude` - The latitude value for the data center. This member supports six decimal places of precision.
      * `liveness_hosts` - A list of the origin servers used to poll the data centers in an Application Load Balancer configuration. These servers support basic HTTP polling.
      * `longitude` - The longitude value for the data center. This member supports six decimal places of precision.
      * `origin_id` - The ID of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
      * `percent` - The percent of traffic that is sent to the data center. The total for all data centers must equal 100%.
      * `state_or_province` - The state, province, or region where the data center is located.
    * `liveness_settings` - Specifies the health of each load balanced data center defined in the data center list.
      * `host_header` - The Host header for the liveness HTTP request.
      * `additional_headers` - Maps additional case-insensitive HTTP header names included to the liveness testing requests.
      * `interval` - The frequency of liveness tests. Defaults to 60 seconds, minimum is 10 seconds.
      * `path` - The path to the test object used for liveness testing. The function of the test object is to help determine whether the data center is functioning.
      * `peer_certificate_verification` - Whether to validate the origin certificate for an HTTPS request.
      * `port` - The port for the test object. The default port is 80, which is standard for HTTP. Enter 443 if you are using HTTPS.
      * `protocol` - The protocol or scheme for the database, either `HTTP` or `HTTPS`.
      * `request_string` - The request used for TCP and TCPS tests.
      * `response_string` - The response used for TCP and TCPS tests.
      * `status_3xx_failure` - If `true`, marks the liveness test as failed when the request returns a 3xx (redirection) status code.
      * `status_4xx_failure` - If `true`, marks the liveness test as failed when the request returns a 4xx (client error) status code.
      * `status_5xx_failure` - If `true`, marks the liveness test as failed when the request returns a 5xx (server error) status code.
      * `timeout` - The number of seconds the system waits before failing the liveness test.


    :param str origin_id: - (Required) A unique identifier for the Conditional Origin that supports the load balancing configuration. The Conditional Origin type must be set to `APPLICATION_LOAD_BALANCER` in the `origin` behavior. See property rules for more information.
    :param int version: - (Optional) The version number of the load balancing configuration.
    """
    __args__ = dict()
    __args__['originId'] = origin_id
    __args__['version'] = version
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('akamai:index/getCloudletsApplicationLoadBalancer:getCloudletsApplicationLoadBalancer', __args__, opts=opts, typ=GetCloudletsApplicationLoadBalancerResult).value

    return AwaitableGetCloudletsApplicationLoadBalancerResult(
        balancing_type=__ret__.balancing_type,
        created_by=__ret__.created_by,
        created_date=__ret__.created_date,
        data_centers=__ret__.data_centers,
        deleted=__ret__.deleted,
        description=__ret__.description,
        id=__ret__.id,
        immutable=__ret__.immutable,
        last_modified_by=__ret__.last_modified_by,
        last_modified_date=__ret__.last_modified_date,
        liveness_settings=__ret__.liveness_settings,
        origin_id=__ret__.origin_id,
        type=__ret__.type,
        version=__ret__.version,
        warnings=__ret__.warnings)


@_utilities.lift_output_func(get_cloudlets_application_load_balancer)
def get_cloudlets_application_load_balancer_output(origin_id: Optional[pulumi.Input[str]] = None,
                                                   version: Optional[pulumi.Input[Optional[int]]] = None,
                                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCloudletsApplicationLoadBalancerResult]:
    """
    Use the `CloudletsApplicationLoadBalancer` data source to list details about the Application Load Balancer configuration with a specified policy version, or latest if not specified.

    ## Basic usage

    This example returns the load balancing configuration details based on the origin ID and optionally, a version:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    example = akamai.get_cloudlets_application_load_balancer(origin_id="alb_test_1",
        version=1)
    ```

    ## Attributes reference

    This data source returns these attributes:

    * `description` - The description of the load balancing configuration.
    * `type` - The type of Conditional Origin. `APPLICATION_LOAD_BALANCER` is the only supported value.
    * `balancing_type` - The type of load balancing being performed, either `WEIGHTED` or `PERFORMANCE`.
    * `created_by` - The name of the user who created this load balancing configuration.
    * `created_date` - The date, in ISO 8601 format, when this load balancing configuration was created.
    * `deleted` - Whether the Conditional Origin version has been deleted. If `false`, you can use this version again.
    * `immutable` - Whether you can edit the load balancing version. The default setting for this member is false. It automatically becomes true when the load balancing version is activated for the first time.
    * `last_modified_by` - The user who last modified the load balancing configuration.
    * `last_modified_date` - The date, in ISO 8601 format, when the initial load balancing configuration was last modified.
    * `warnings` - A list of warnings that occured during the activation of the load balancing configuration.
    * `data_centers` - Specifies the Conditional Origins being used as data centers for an Application Load Balancer implementation. Only Conditional Origins with an origin type of `CUSTOMER` or `NETSTORAGE` can be used as data centers in an Application Load Balancer configuration.
      * `city` - The city in which the data center is located.
      * `cloud_server_host_header_override` - Whether the cloud server host header is overridden.
      * `cloud_service` - Whether this datacenter is a cloud service.
      * `continent` - The code of the continent on which the data center is located. See [Continent Codes](https://control.akamai.com/dl/edgescape/continentCodes.csv) for a list of valid codes.
      * `country` - The country in which the data center is located. See [Country Codes](https://control.akamai.com/dl/edgescape/cc2continent.csv) for a list of valid codes.
      * `hostname` - The name of the host that can be used as a Conditional Origin. This should match the `hostname` value defined for this datacenter in Property Manager.
      * `latitude` - The latitude value for the data center. This member supports six decimal places of precision.
      * `liveness_hosts` - A list of the origin servers used to poll the data centers in an Application Load Balancer configuration. These servers support basic HTTP polling.
      * `longitude` - The longitude value for the data center. This member supports six decimal places of precision.
      * `origin_id` - The ID of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
      * `percent` - The percent of traffic that is sent to the data center. The total for all data centers must equal 100%.
      * `state_or_province` - The state, province, or region where the data center is located.
    * `liveness_settings` - Specifies the health of each load balanced data center defined in the data center list.
      * `host_header` - The Host header for the liveness HTTP request.
      * `additional_headers` - Maps additional case-insensitive HTTP header names included to the liveness testing requests.
      * `interval` - The frequency of liveness tests. Defaults to 60 seconds, minimum is 10 seconds.
      * `path` - The path to the test object used for liveness testing. The function of the test object is to help determine whether the data center is functioning.
      * `peer_certificate_verification` - Whether to validate the origin certificate for an HTTPS request.
      * `port` - The port for the test object. The default port is 80, which is standard for HTTP. Enter 443 if you are using HTTPS.
      * `protocol` - The protocol or scheme for the database, either `HTTP` or `HTTPS`.
      * `request_string` - The request used for TCP and TCPS tests.
      * `response_string` - The response used for TCP and TCPS tests.
      * `status_3xx_failure` - If `true`, marks the liveness test as failed when the request returns a 3xx (redirection) status code.
      * `status_4xx_failure` - If `true`, marks the liveness test as failed when the request returns a 4xx (client error) status code.
      * `status_5xx_failure` - If `true`, marks the liveness test as failed when the request returns a 5xx (server error) status code.
      * `timeout` - The number of seconds the system waits before failing the liveness test.


    :param str origin_id: - (Required) A unique identifier for the Conditional Origin that supports the load balancing configuration. The Conditional Origin type must be set to `APPLICATION_LOAD_BALANCER` in the `origin` behavior. See property rules for more information.
    :param int version: - (Optional) The version number of the load balancing configuration.
    """
    ...
