// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetPropertyProducts
    {
        /// <summary>
        /// Use the `akamai.getPropertyProducts` data source to list the products included on a contract. 
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// This example returns products associated with the [EdgeGrid client token](https://developer.akamai.com/getting-started/edgegrid) for a given contract:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         this.PropertyMatch = data.Akamai_property_products.My_example;
        ///     }
        /// 
        ///     [Output("propertyMatch")]
        ///     public Output&lt;string&gt; PropertyMatch { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Argument reference
        /// 
        /// This data source supports this argument:
        /// 
        /// * `contract_id` - (Required) A contract's unique ID, including the `ctr_` prefix. 
        /// 
        /// ## Attributes reference
        /// 
        /// This data source returns these attributes:
        /// 
        /// * `products` - A list of supported products for the contract, including:
        ///   * `product_id` - The product's unique ID, including the `prd_` prefix.
        ///   * `product_name` - A string containing the product name.
        /// </summary>
        public static Task<GetPropertyProductsResult> InvokeAsync(GetPropertyProductsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPropertyProductsResult>("akamai:index/getPropertyProducts:getPropertyProducts", args ?? new GetPropertyProductsArgs(), options.WithVersion());
    }


    public sealed class GetPropertyProductsArgs : Pulumi.InvokeArgs
    {
        [Input("contractId", required: true)]
        public string ContractId { get; set; } = null!;

        public GetPropertyProductsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPropertyProductsResult
    {
        public readonly string ContractId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetPropertyProductsProductResult> Products;

        [OutputConstructor]
        private GetPropertyProductsResult(
            string contractId,

            string id,

            ImmutableArray<Outputs.GetPropertyProductsProductResult> products)
        {
            ContractId = contractId;
            Id = id;
            Products = products;
        }
    }
}
