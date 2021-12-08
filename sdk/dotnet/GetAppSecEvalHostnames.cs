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
    public static class GetAppSecEvalHostnames
    {
        /// <summary>
        /// **Scopes**: Security configuration
        /// 
        /// Returns the evaluation hostnames for a configuration. In evaluation mode, you use evaluation hosts to monitor how well your configuration settings protects host traffic. (Note that the evaluation host isn't actually protected, and the host takes no action other than recording the actions it would have taken had it been on the production network).
        /// 
        /// Evaluation mode for hostnames is only available for organizations running Web Application Protector.
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/selected-hostnames/eval-hostnames](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getevaluationhostnames)
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
        ///         var evalHostnamesAppSecEvalHostnames = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecEvalHostnames.InvokeAsync(new Akamai.GetAppSecEvalHostnamesArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.EvalHostnames = evalHostnamesAppSecEvalHostnames.Apply(evalHostnamesAppSecEvalHostnames =&gt; evalHostnamesAppSecEvalHostnames.Hostnames);
        ///         this.EvalHostnamesOutput = evalHostnamesAppSecEvalHostnames.Apply(evalHostnamesAppSecEvalHostnames =&gt; evalHostnamesAppSecEvalHostnames.OutputText);
        ///         this.EvalHostnamesJson = evalHostnamesAppSecEvalHostnames.Apply(evalHostnamesAppSecEvalHostnames =&gt; evalHostnamesAppSecEvalHostnames.Json);
        ///     }
        /// 
        ///     [Output("evalHostnames")]
        ///     public Output&lt;string&gt; EvalHostnames { get; set; }
        ///     [Output("evalHostnamesOutput")]
        ///     public Output&lt;string&gt; EvalHostnamesOutput { get; set; }
        ///     [Output("evalHostnamesJson")]
        ///     public Output&lt;string&gt; EvalHostnamesJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `hostnames`. List of evaluation hostnames.
        /// - `json`. JSON-formatted list of evaluation hostnames.
        /// - `output_text`. Tabular report showing evaluation hostnames.
        /// </summary>
        public static Task<GetAppSecEvalHostnamesResult> InvokeAsync(GetAppSecEvalHostnamesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecEvalHostnamesResult>("akamai:index/getAppSecEvalHostnames:getAppSecEvalHostnames", args ?? new GetAppSecEvalHostnamesArgs(), options.WithVersion());

        /// <summary>
        /// **Scopes**: Security configuration
        /// 
        /// Returns the evaluation hostnames for a configuration. In evaluation mode, you use evaluation hosts to monitor how well your configuration settings protects host traffic. (Note that the evaluation host isn't actually protected, and the host takes no action other than recording the actions it would have taken had it been on the production network).
        /// 
        /// Evaluation mode for hostnames is only available for organizations running Web Application Protector.
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/selected-hostnames/eval-hostnames](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getevaluationhostnames)
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
        ///         var evalHostnamesAppSecEvalHostnames = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecEvalHostnames.InvokeAsync(new Akamai.GetAppSecEvalHostnamesArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.EvalHostnames = evalHostnamesAppSecEvalHostnames.Apply(evalHostnamesAppSecEvalHostnames =&gt; evalHostnamesAppSecEvalHostnames.Hostnames);
        ///         this.EvalHostnamesOutput = evalHostnamesAppSecEvalHostnames.Apply(evalHostnamesAppSecEvalHostnames =&gt; evalHostnamesAppSecEvalHostnames.OutputText);
        ///         this.EvalHostnamesJson = evalHostnamesAppSecEvalHostnames.Apply(evalHostnamesAppSecEvalHostnames =&gt; evalHostnamesAppSecEvalHostnames.Json);
        ///     }
        /// 
        ///     [Output("evalHostnames")]
        ///     public Output&lt;string&gt; EvalHostnames { get; set; }
        ///     [Output("evalHostnamesOutput")]
        ///     public Output&lt;string&gt; EvalHostnamesOutput { get; set; }
        ///     [Output("evalHostnamesJson")]
        ///     public Output&lt;string&gt; EvalHostnamesJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `hostnames`. List of evaluation hostnames.
        /// - `json`. JSON-formatted list of evaluation hostnames.
        /// - `output_text`. Tabular report showing evaluation hostnames.
        /// </summary>
        public static Output<GetAppSecEvalHostnamesResult> Invoke(GetAppSecEvalHostnamesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecEvalHostnamesResult>("akamai:index/getAppSecEvalHostnames:getAppSecEvalHostnames", args ?? new GetAppSecEvalHostnamesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetAppSecEvalHostnamesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration running in evaluation mode.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        public GetAppSecEvalHostnamesArgs()
        {
        }
    }

    public sealed class GetAppSecEvalHostnamesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration running in evaluation mode.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        public GetAppSecEvalHostnamesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecEvalHostnamesResult
    {
        public readonly int ConfigId;
        public readonly ImmutableArray<string> Hostnames;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Json;
        public readonly string OutputText;

        [OutputConstructor]
        private GetAppSecEvalHostnamesResult(
            int configId,

            ImmutableArray<string> hostnames,

            string id,

            string json,

            string outputText)
        {
            ConfigId = configId;
            Hostnames = hostnames;
            Id = id;
            Json = json;
            OutputText = outputText;
        }
    }
}
