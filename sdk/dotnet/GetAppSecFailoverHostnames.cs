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
    public static class GetAppSecFailoverHostnames
    {
        /// <summary>
        /// Use the `akamai.getAppSecFailoverHostnames` data source to retrieve a list of the failover hostnames in a configuration. The information available is described [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getfailoverhostnames).
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
        ///         var failoverHostnamesAppSecFailoverHostnames = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecFailoverHostnames.InvokeAsync(new Akamai.GetAppSecFailoverHostnamesArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.FailoverHostnames = failoverHostnamesAppSecFailoverHostnames.Apply(failoverHostnamesAppSecFailoverHostnames =&gt; failoverHostnamesAppSecFailoverHostnames.Hostnames);
        ///         this.FailoverHostnamesOutput = failoverHostnamesAppSecFailoverHostnames.Apply(failoverHostnamesAppSecFailoverHostnames =&gt; failoverHostnamesAppSecFailoverHostnames.OutputText);
        ///         this.FailoverHostnamesJson = failoverHostnamesAppSecFailoverHostnames.Apply(failoverHostnamesAppSecFailoverHostnames =&gt; failoverHostnamesAppSecFailoverHostnames.Json);
        ///     }
        /// 
        ///     [Output("failoverHostnames")]
        ///     public Output&lt;string&gt; FailoverHostnames { get; set; }
        ///     [Output("failoverHostnamesOutput")]
        ///     public Output&lt;string&gt; FailoverHostnamesOutput { get; set; }
        ///     [Output("failoverHostnamesJson")]
        ///     public Output&lt;string&gt; FailoverHostnamesJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecFailoverHostnamesResult> InvokeAsync(GetAppSecFailoverHostnamesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecFailoverHostnamesResult>("akamai:index/getAppSecFailoverHostnames:getAppSecFailoverHostnames", args ?? new GetAppSecFailoverHostnamesArgs(), options.WithVersion());

        /// <summary>
        /// Use the `akamai.getAppSecFailoverHostnames` data source to retrieve a list of the failover hostnames in a configuration. The information available is described [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getfailoverhostnames).
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
        ///         var failoverHostnamesAppSecFailoverHostnames = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecFailoverHostnames.InvokeAsync(new Akamai.GetAppSecFailoverHostnamesArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.FailoverHostnames = failoverHostnamesAppSecFailoverHostnames.Apply(failoverHostnamesAppSecFailoverHostnames =&gt; failoverHostnamesAppSecFailoverHostnames.Hostnames);
        ///         this.FailoverHostnamesOutput = failoverHostnamesAppSecFailoverHostnames.Apply(failoverHostnamesAppSecFailoverHostnames =&gt; failoverHostnamesAppSecFailoverHostnames.OutputText);
        ///         this.FailoverHostnamesJson = failoverHostnamesAppSecFailoverHostnames.Apply(failoverHostnamesAppSecFailoverHostnames =&gt; failoverHostnamesAppSecFailoverHostnames.Json);
        ///     }
        /// 
        ///     [Output("failoverHostnames")]
        ///     public Output&lt;string&gt; FailoverHostnames { get; set; }
        ///     [Output("failoverHostnamesOutput")]
        ///     public Output&lt;string&gt; FailoverHostnamesOutput { get; set; }
        ///     [Output("failoverHostnamesJson")]
        ///     public Output&lt;string&gt; FailoverHostnamesJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAppSecFailoverHostnamesResult> Invoke(GetAppSecFailoverHostnamesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecFailoverHostnamesResult>("akamai:index/getAppSecFailoverHostnames:getAppSecFailoverHostnames", args ?? new GetAppSecFailoverHostnamesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetAppSecFailoverHostnamesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        public GetAppSecFailoverHostnamesArgs()
        {
        }
    }

    public sealed class GetAppSecFailoverHostnamesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        public GetAppSecFailoverHostnamesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecFailoverHostnamesResult
    {
        public readonly int ConfigId;
        /// <summary>
        /// A list of the failover hostnames.
        /// </summary>
        public readonly ImmutableArray<string> Hostnames;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A JSON-formatted list of the failover hostnames.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// A tabular display showing the failover hostnames.
        /// </summary>
        public readonly string OutputText;

        [OutputConstructor]
        private GetAppSecFailoverHostnamesResult(
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
