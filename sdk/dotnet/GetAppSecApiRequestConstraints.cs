// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Akamai
{
    public static class GetAppSecApiRequestConstraints
    {
        /// <summary>
        /// Use the `akamai.AppSecApiRequestConstraints` data source to retrieve a list of APIs with their constraints and associated actions, or the constraints and actions for a particular API. The information available is described [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getapirequestconstraints).
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
        ///             Name = @var.Security_configuration,
        ///         }));
        ///         var apisRequestConstraints = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecApiRequestConstraints.InvokeAsync(new Akamai.GetAppSecApiRequestConstraintsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = @var.Security_policy_id,
        ///         })));
        ///         this.ApisConstraintsText = apisRequestConstraints.Apply(apisRequestConstraints =&gt; apisRequestConstraints.OutputText);
        ///         this.ApisConstraintsJson = apisRequestConstraints.Apply(apisRequestConstraints =&gt; apisRequestConstraints.Json);
        ///         var apiRequestConstraints = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecApiRequestConstraints.InvokeAsync(new Akamai.GetAppSecApiRequestConstraintsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = @var.Security_policy_id,
        ///             ApiId = @var.Api_id,
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
        /// </summary>
        public static Task<GetAppSecApiRequestConstraintsResult> InvokeAsync(GetAppSecApiRequestConstraintsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecApiRequestConstraintsResult>("akamai:index/getAppSecApiRequestConstraints:getAppSecApiRequestConstraints", args ?? new GetAppSecApiRequestConstraintsArgs(), options.WithVersion());

        /// <summary>
        /// Use the `akamai.AppSecApiRequestConstraints` data source to retrieve a list of APIs with their constraints and associated actions, or the constraints and actions for a particular API. The information available is described [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getapirequestconstraints).
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
        ///             Name = @var.Security_configuration,
        ///         }));
        ///         var apisRequestConstraints = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecApiRequestConstraints.InvokeAsync(new Akamai.GetAppSecApiRequestConstraintsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = @var.Security_policy_id,
        ///         })));
        ///         this.ApisConstraintsText = apisRequestConstraints.Apply(apisRequestConstraints =&gt; apisRequestConstraints.OutputText);
        ///         this.ApisConstraintsJson = apisRequestConstraints.Apply(apisRequestConstraints =&gt; apisRequestConstraints.Json);
        ///         var apiRequestConstraints = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecApiRequestConstraints.InvokeAsync(new Akamai.GetAppSecApiRequestConstraintsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = @var.Security_policy_id,
        ///             ApiId = @var.Api_id,
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
        /// </summary>
        public static Output<GetAppSecApiRequestConstraintsResult> Invoke(GetAppSecApiRequestConstraintsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecApiRequestConstraintsResult>("akamai:index/getAppSecApiRequestConstraints:getAppSecApiRequestConstraints", args ?? new GetAppSecApiRequestConstraintsInvokeArgs(), options.WithVersion());
    }


    public sealed class GetAppSecApiRequestConstraintsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of a specific API for which to retrieve constraint information.
        /// </summary>
        [Input("apiId")]
        public int? ApiId { get; set; }

        /// <summary>
        /// The configuration ID to use.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// The ID of the security policy to use.
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
        /// The ID of a specific API for which to retrieve constraint information.
        /// </summary>
        [Input("apiId")]
        public Input<int>? ApiId { get; set; }

        /// <summary>
        /// The configuration ID to use.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// The ID of the security policy to use.
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
        /// <summary>
        /// A JSON-formatted list of information about the APIs and their constraints and actions.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// A tabular display showing the APIs and their constraints and actions.
        /// </summary>
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
