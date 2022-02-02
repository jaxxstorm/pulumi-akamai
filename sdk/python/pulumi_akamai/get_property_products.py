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
    'GetPropertyProductsResult',
    'AwaitableGetPropertyProductsResult',
    'get_property_products',
    'get_property_products_output',
]

@pulumi.output_type
class GetPropertyProductsResult:
    """
    A collection of values returned by getPropertyProducts.
    """
    def __init__(__self__, contract_id=None, id=None, products=None):
        if contract_id and not isinstance(contract_id, str):
            raise TypeError("Expected argument 'contract_id' to be a str")
        pulumi.set(__self__, "contract_id", contract_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if products and not isinstance(products, list):
            raise TypeError("Expected argument 'products' to be a list")
        pulumi.set(__self__, "products", products)

    @property
    @pulumi.getter(name="contractId")
    def contract_id(self) -> str:
        return pulumi.get(self, "contract_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def products(self) -> Sequence['outputs.GetPropertyProductsProductResult']:
        return pulumi.get(self, "products")


class AwaitableGetPropertyProductsResult(GetPropertyProductsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPropertyProductsResult(
            contract_id=self.contract_id,
            id=self.id,
            products=self.products)


def get_property_products(contract_id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPropertyProductsResult:
    """
    Use the `get_property_products` data source to list the products included on a contract.

    ## Example Usage

    This example returns products associated with the [EdgeGrid client token](https://developer.akamai.com/getting-started/edgegrid) for a given contract:

    ```python
    import pulumi

    pulumi.export("propertyMatch", data["akamai_property_products"]["my-example"])
    ```
    ## Attributes reference

    This data source returns these attributes:

    * `products` - A list of supported products for the contract, including:
      * `product_id` - The product's unique ID, including the `prd_` prefix.
      * `product_name` - A string containing the product name.


    :param str contract_id: - (Required) A contract's unique ID, including the `ctr_` prefix.
    """
    __args__ = dict()
    __args__['contractId'] = contract_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('akamai:index/getPropertyProducts:getPropertyProducts', __args__, opts=opts, typ=GetPropertyProductsResult).value

    return AwaitableGetPropertyProductsResult(
        contract_id=__ret__.contract_id,
        id=__ret__.id,
        products=__ret__.products)


@_utilities.lift_output_func(get_property_products)
def get_property_products_output(contract_id: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPropertyProductsResult]:
    """
    Use the `get_property_products` data source to list the products included on a contract.

    ## Example Usage

    This example returns products associated with the [EdgeGrid client token](https://developer.akamai.com/getting-started/edgegrid) for a given contract:

    ```python
    import pulumi

    pulumi.export("propertyMatch", data["akamai_property_products"]["my-example"])
    ```
    ## Attributes reference

    This data source returns these attributes:

    * `products` - A list of supported products for the contract, including:
      * `product_id` - The product's unique ID, including the `prd_` prefix.
      * `product_name` - A string containing the product name.


    :param str contract_id: - (Required) A contract's unique ID, including the `ctr_` prefix.
    """
    ...
