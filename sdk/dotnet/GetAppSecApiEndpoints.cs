// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecApiEndpoints
    {
        /// <summary>
        /// Use the `akamai.getAppSecApiEndpoints` data source to retrieve information about the API Endpoints associated with a security policy or configuration. The information available is described [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getapiendpoints).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic usage:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var apiEndpoints = Output.Create(Akamai.GetAppSecApiEndpoints.InvokeAsync(new Akamai.GetAppSecApiEndpointsArgs
        ///         {
        ///             ApiName = "TestEndpoint",
        ///             ConfigId = 43253,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecApiEndpointsResult> InvokeAsync(GetAppSecApiEndpointsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecApiEndpointsResult>("akamai:index/getAppSecApiEndpoints:getAppSecApiEndpoints", args ?? new GetAppSecApiEndpointsArgs(), options.WithVersion());
    }


    public sealed class GetAppSecApiEndpointsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of a specific endpoint.
        /// </summary>
        [Input("apiName")]
        public string? ApiName { get; set; }

        /// <summary>
        /// The configuration ID.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// The ID of the security policy to use.
        /// </summary>
        [Input("securityPolicyId")]
        public string? SecurityPolicyId { get; set; }

        public GetAppSecApiEndpointsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecApiEndpointsResult
    {
        public readonly string? ApiName;
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of IDs of the API endpoints.
        /// </summary>
        public readonly ImmutableArray<int> IdLists;
        /// <summary>
        /// A JSON-formatted list of information about the API endpoints.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// A tabular display showing the ID and name of the API endpoints.
        /// </summary>
        public readonly string OutputText;
        public readonly string? SecurityPolicyId;

        [OutputConstructor]
        private GetAppSecApiEndpointsResult(
            string? apiName,

            int configId,

            string id,

            ImmutableArray<int> idLists,

            string json,

            string outputText,

            string? securityPolicyId)
        {
            ApiName = apiName;
            ConfigId = configId;
            Id = id;
            IdLists = idLists;
            Json = json;
            OutputText = outputText;
            SecurityPolicyId = securityPolicyId;
        }
    }
}