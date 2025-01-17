// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetEdgeWorkersResourceTier
    {
        /// <summary>
        /// Use the `akamai.getEdgeWorkersResourceTier` data source to list the available resource tiers for a specific contract ID. The resource tier defines the resource consumption [limits](https://techdocs.akamai.com/edgeworkers/docs/resource-tier-limitations) for an EdgeWorker ID.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// This example returns the resource tier fields for an EdgeWorker ID:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Akamai.GetEdgeWorkersResourceTier.InvokeAsync(new Akamai.GetEdgeWorkersResourceTierArgs
        ///         {
        ///             ContractId = "1-ABC",
        ///             ResourceTierName = "Basic Compute",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Attributes reference
        /// 
        /// This data source returns these attributes:
        /// 
        /// * `resource_tier_id` - Unique identifier of the resource tier.
        /// </summary>
        public static Task<GetEdgeWorkersResourceTierResult> InvokeAsync(GetEdgeWorkersResourceTierArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEdgeWorkersResourceTierResult>("akamai:index/getEdgeWorkersResourceTier:getEdgeWorkersResourceTier", args ?? new GetEdgeWorkersResourceTierArgs(), options.WithDefaults());

        /// <summary>
        /// Use the `akamai.getEdgeWorkersResourceTier` data source to list the available resource tiers for a specific contract ID. The resource tier defines the resource consumption [limits](https://techdocs.akamai.com/edgeworkers/docs/resource-tier-limitations) for an EdgeWorker ID.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// This example returns the resource tier fields for an EdgeWorker ID:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Akamai.GetEdgeWorkersResourceTier.InvokeAsync(new Akamai.GetEdgeWorkersResourceTierArgs
        ///         {
        ///             ContractId = "1-ABC",
        ///             ResourceTierName = "Basic Compute",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Attributes reference
        /// 
        /// This data source returns these attributes:
        /// 
        /// * `resource_tier_id` - Unique identifier of the resource tier.
        /// </summary>
        public static Output<GetEdgeWorkersResourceTierResult> Invoke(GetEdgeWorkersResourceTierInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEdgeWorkersResourceTierResult>("akamai:index/getEdgeWorkersResourceTier:getEdgeWorkersResourceTier", args ?? new GetEdgeWorkersResourceTierInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEdgeWorkersResourceTierArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier of a contract.
        /// </summary>
        [Input("contractId", required: true)]
        public string ContractId { get; set; } = null!;

        /// <summary>
        /// Unique name of the resource tier.
        /// </summary>
        [Input("resourceTierName", required: true)]
        public string ResourceTierName { get; set; } = null!;

        public GetEdgeWorkersResourceTierArgs()
        {
        }
    }

    public sealed class GetEdgeWorkersResourceTierInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier of a contract.
        /// </summary>
        [Input("contractId", required: true)]
        public Input<string> ContractId { get; set; } = null!;

        /// <summary>
        /// Unique name of the resource tier.
        /// </summary>
        [Input("resourceTierName", required: true)]
        public Input<string> ResourceTierName { get; set; } = null!;

        public GetEdgeWorkersResourceTierInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEdgeWorkersResourceTierResult
    {
        public readonly string ContractId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int ResourceTierId;
        public readonly string ResourceTierName;

        [OutputConstructor]
        private GetEdgeWorkersResourceTierResult(
            string contractId,

            string id,

            int resourceTierId,

            string resourceTierName)
        {
            ContractId = contractId;
            Id = id;
            ResourceTierId = resourceTierId;
            ResourceTierName = resourceTierName;
        }
    }
}
