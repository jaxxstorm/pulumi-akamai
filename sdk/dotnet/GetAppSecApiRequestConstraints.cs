// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecApiRequestConstraints
    {
        /// <summary>
        /// **Scopes**: Security policy; API endpoint
        /// 
        /// Returns information about API endpoint constraints and actions. 
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/api-request-constraints](https://techdocs.akamai.com/application-security/reference/get-api-request-constraints)
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
        ///         var configuration = Output.Create(Akamai.GetAppSecConfiguration.InvokeAsync(new Akamai.GetAppSecConfigurationArgs
        ///         {
        ///             Name = "Documentation",
        ///         }));
        ///         var apisRequestConstraints = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecApiRequestConstraints.InvokeAsync(new Akamai.GetAppSecApiRequestConstraintsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = "gms1_134637",
        ///         })));
        ///         this.ApisConstraintsText = apisRequestConstraints.Apply(apisRequestConstraints =&gt; apisRequestConstraints.OutputText);
        ///         this.ApisConstraintsJson = apisRequestConstraints.Apply(apisRequestConstraints =&gt; apisRequestConstraints.Json);
        ///         var apiRequestConstraints = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecApiRequestConstraints.InvokeAsync(new Akamai.GetAppSecApiRequestConstraintsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = "gms1_134637",
        ///             ApiId = 624913,
        ///         })));
        ///         this.ApiConstraintsText = apiRequestConstraints.Apply(apiRequestConstraints =&gt; apiRequestConstraints.OutputText);
        ///         this.ApiConstraintsJson = apiRequestConstraints.Apply(apiRequestConstraints =&gt; apiRequestConstraints.Json);
        ///     }
        /// 
        ///     [Output("apisConstraintsText")]
        ///     public Output&lt;string&gt; ApisConstraintsText { get; set; }
        ///     [Output("apisConstraintsJson")]
        ///     public Output&lt;string&gt; ApisConstraintsJson { get; set; }
        ///     [Output("apiConstraintsText")]
        ///     public Output&lt;string&gt; ApiConstraintsText { get; set; }
        ///     [Output("apiConstraintsJson")]
        ///     public Output&lt;string&gt; ApiConstraintsJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `json`. JSON-formatted list of information about the APIs, their constraints, and their actions.
        /// - `output_text`. Tabular report of the APIs, their constraints, and their actions.
        /// </summary>
        public static Task<GetAppSecApiRequestConstraintsResult> InvokeAsync(GetAppSecApiRequestConstraintsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecApiRequestConstraintsResult>("akamai:index/getAppSecApiRequestConstraints:getAppSecApiRequestConstraints", args ?? new GetAppSecApiRequestConstraintsArgs(), options.WithDefaults());

        /// <summary>
        /// **Scopes**: Security policy; API endpoint
        /// 
        /// Returns information about API endpoint constraints and actions. 
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/api-request-constraints](https://techdocs.akamai.com/application-security/reference/get-api-request-constraints)
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
        ///         var configuration = Output.Create(Akamai.GetAppSecConfiguration.InvokeAsync(new Akamai.GetAppSecConfigurationArgs
        ///         {
        ///             Name = "Documentation",
        ///         }));
        ///         var apisRequestConstraints = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecApiRequestConstraints.InvokeAsync(new Akamai.GetAppSecApiRequestConstraintsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = "gms1_134637",
        ///         })));
        ///         this.ApisConstraintsText = apisRequestConstraints.Apply(apisRequestConstraints =&gt; apisRequestConstraints.OutputText);
        ///         this.ApisConstraintsJson = apisRequestConstraints.Apply(apisRequestConstraints =&gt; apisRequestConstraints.Json);
        ///         var apiRequestConstraints = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecApiRequestConstraints.InvokeAsync(new Akamai.GetAppSecApiRequestConstraintsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = "gms1_134637",
        ///             ApiId = 624913,
        ///         })));
        ///         this.ApiConstraintsText = apiRequestConstraints.Apply(apiRequestConstraints =&gt; apiRequestConstraints.OutputText);
        ///         this.ApiConstraintsJson = apiRequestConstraints.Apply(apiRequestConstraints =&gt; apiRequestConstraints.Json);
        ///     }
        /// 
        ///     [Output("apisConstraintsText")]
        ///     public Output&lt;string&gt; ApisConstraintsText { get; set; }
        ///     [Output("apisConstraintsJson")]
        ///     public Output&lt;string&gt; ApisConstraintsJson { get; set; }
        ///     [Output("apiConstraintsText")]
        ///     public Output&lt;string&gt; ApiConstraintsText { get; set; }
        ///     [Output("apiConstraintsJson")]
        ///     public Output&lt;string&gt; ApiConstraintsJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `json`. JSON-formatted list of information about the APIs, their constraints, and their actions.
        /// - `output_text`. Tabular report of the APIs, their constraints, and their actions.
        /// </summary>
        public static Output<GetAppSecApiRequestConstraintsResult> Invoke(GetAppSecApiRequestConstraintsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecApiRequestConstraintsResult>("akamai:index/getAppSecApiRequestConstraints:getAppSecApiRequestConstraints", args ?? new GetAppSecApiRequestConstraintsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppSecApiRequestConstraintsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the API endpoint you want to return constraint information for.
        /// </summary>
        [Input("apiId")]
        public int? ApiId { get; set; }

        /// <summary>
        /// . Unique identifier of the security configuration associated with the API constraints.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// . Unique identifier of the security policy associated with the API constraints.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public string SecurityPolicyId { get; set; } = null!;

        public GetAppSecApiRequestConstraintsArgs()
        {
        }
    }

    public sealed class GetAppSecApiRequestConstraintsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the API endpoint you want to return constraint information for.
        /// </summary>
        [Input("apiId")]
        public Input<int>? ApiId { get; set; }

        /// <summary>
        /// . Unique identifier of the security configuration associated with the API constraints.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Unique identifier of the security policy associated with the API constraints.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public Input<string> SecurityPolicyId { get; set; } = null!;

        public GetAppSecApiRequestConstraintsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecApiRequestConstraintsResult
    {
        public readonly int? ApiId;
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Json;
        public readonly string OutputText;
        public readonly string SecurityPolicyId;

        [OutputConstructor]
        private GetAppSecApiRequestConstraintsResult(
            int? apiId,

            int configId,

            string id,

            string json,

            string outputText,

            string securityPolicyId)
        {
            ApiId = apiId;
            ConfigId = configId;
            Id = id;
            Json = json;
            OutputText = outputText;
            SecurityPolicyId = securityPolicyId;
        }
    }
}
