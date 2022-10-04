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
        /// **Scopes**: Security configuration; security policy
        /// 
        /// Returns information about the API endpoints associated with a security policy or configuration. 
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/api-endpoints](https://techdocs.akamai.com/application-security/reference/get-api-endpoints)
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
        ///             ApiName = "Contracts",
        ///             ConfigId = 58843,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `id_list`. List of API endpoint IDs.
        /// - `json`. JSON-formatted list of information about the API endpoints.
        /// - `output_text`. Tabular report showing the ID and name of the API endpoints.
        /// </summary>
        public static Task<GetAppSecApiEndpointsResult> InvokeAsync(GetAppSecApiEndpointsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecApiEndpointsResult>("akamai:index/getAppSecApiEndpoints:getAppSecApiEndpoints", args ?? new GetAppSecApiEndpointsArgs(), options.WithDefaults());

        /// <summary>
        /// **Scopes**: Security configuration; security policy
        /// 
        /// Returns information about the API endpoints associated with a security policy or configuration. 
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/api-endpoints](https://techdocs.akamai.com/application-security/reference/get-api-endpoints)
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
        ///             ApiName = "Contracts",
        ///             ConfigId = 58843,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `id_list`. List of API endpoint IDs.
        /// - `json`. JSON-formatted list of information about the API endpoints.
        /// - `output_text`. Tabular report showing the ID and name of the API endpoints.
        /// </summary>
        public static Output<GetAppSecApiEndpointsResult> Invoke(GetAppSecApiEndpointsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecApiEndpointsResult>("akamai:index/getAppSecApiEndpoints:getAppSecApiEndpoints", args ?? new GetAppSecApiEndpointsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppSecApiEndpointsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Name of the API endpoint you want to return information for. If not included, information is returned for all your API endpoints.
        /// </summary>
        [Input("apiName")]
        public string? ApiName { get; set; }

        /// <summary>
        /// . Unique identifier of the security configuration associated with the API endpoints.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// . Unique identifier of the security policy associated with the API endpoints. If not included, information is returned for all your security policies.
        /// </summary>
        [Input("securityPolicyId")]
        public string? SecurityPolicyId { get; set; }

        public GetAppSecApiEndpointsArgs()
        {
        }
    }

    public sealed class GetAppSecApiEndpointsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Name of the API endpoint you want to return information for. If not included, information is returned for all your API endpoints.
        /// </summary>
        [Input("apiName")]
        public Input<string>? ApiName { get; set; }

        /// <summary>
        /// . Unique identifier of the security configuration associated with the API endpoints.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Unique identifier of the security policy associated with the API endpoints. If not included, information is returned for all your security policies.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        public GetAppSecApiEndpointsInvokeArgs()
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
        public readonly ImmutableArray<int> IdLists;
        public readonly string Json;
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
