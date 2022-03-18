# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetDatastreamActivationHistoryResult',
    'AwaitableGetDatastreamActivationHistoryResult',
    'get_datastream_activation_history',
    'get_datastream_activation_history_output',
]

@pulumi.output_type
class GetDatastreamActivationHistoryResult:
    """
    A collection of values returned by getDatastreamActivationHistory.
    """
    def __init__(__self__, activations=None, id=None, stream_id=None):
        if activations and not isinstance(activations, list):
            raise TypeError("Expected argument 'activations' to be a list")
        pulumi.set(__self__, "activations", activations)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if stream_id and not isinstance(stream_id, int):
            raise TypeError("Expected argument 'stream_id' to be a int")
        pulumi.set(__self__, "stream_id", stream_id)

    @property
    @pulumi.getter
    def activations(self) -> Sequence['outputs.GetDatastreamActivationHistoryActivationResult']:
        return pulumi.get(self, "activations")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="streamId")
    def stream_id(self) -> int:
        return pulumi.get(self, "stream_id")


class AwaitableGetDatastreamActivationHistoryResult(GetDatastreamActivationHistoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatastreamActivationHistoryResult(
            activations=self.activations,
            id=self.id,
            stream_id=self.stream_id)


def get_datastream_activation_history(stream_id: Optional[int] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatastreamActivationHistoryResult:
    """
    Use the `get_datastream_activation_history` data source to list detailed information about the activation status changes for all versions of a stream.

    ## Example Usage

    This example returns the activation history for a provided stream ID:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    ds = akamai.get_datastream_activation_history(stream_id=12345)
    pulumi.export("dsHistoryStreamId", ds.stream_id)
    pulumi.export("dsHistoryActivations", ds.activations)
    ```
    ## Attributes reference

    This data source returns these attributes:

    * `activations` - Detailed information about an activation status change for a version of a stream, including:
      * `created_by` - The user who activated or deactivated the stream.
      * `created_date` - The date and time of an activation status change.
      * `stream_id` - A stream's unique identifier.
      * `stream_version_id` - A stream version's unique identifier.
      * `is_active` -	Whether the version of the stream is active.


    :param int stream_id: - (Required) A stream's unique identifier.
    """
    __args__ = dict()
    __args__['streamId'] = stream_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:index/getDatastreamActivationHistory:getDatastreamActivationHistory', __args__, opts=opts, typ=GetDatastreamActivationHistoryResult).value

    return AwaitableGetDatastreamActivationHistoryResult(
        activations=__ret__.activations,
        id=__ret__.id,
        stream_id=__ret__.stream_id)


@_utilities.lift_output_func(get_datastream_activation_history)
def get_datastream_activation_history_output(stream_id: Optional[pulumi.Input[int]] = None,
                                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatastreamActivationHistoryResult]:
    """
    Use the `get_datastream_activation_history` data source to list detailed information about the activation status changes for all versions of a stream.

    ## Example Usage

    This example returns the activation history for a provided stream ID:

    ```python
    import pulumi
    import pulumi_akamai as akamai

    ds = akamai.get_datastream_activation_history(stream_id=12345)
    pulumi.export("dsHistoryStreamId", ds.stream_id)
    pulumi.export("dsHistoryActivations", ds.activations)
    ```
    ## Attributes reference

    This data source returns these attributes:

    * `activations` - Detailed information about an activation status change for a version of a stream, including:
      * `created_by` - The user who activated or deactivated the stream.
      * `created_date` - The date and time of an activation status change.
      * `stream_id` - A stream's unique identifier.
      * `stream_version_id` - A stream version's unique identifier.
      * `is_active` -	Whether the version of the stream is active.


    :param int stream_id: - (Required) A stream's unique identifier.
    """
    ...