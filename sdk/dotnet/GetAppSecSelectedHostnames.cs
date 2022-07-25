// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecSelectedHostnames
    {
        /// <summary>
        /// **Scopes**: Security configuration
        /// 
        /// Returns a list of the hostnames currently protected by the specified security configuration.
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/selected-hostnames](https://techdocs.akamai.com/application-security/reference/get-selected-hostnames)
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
        ///         var selectedHostnamesAppSecSelectedHostnames = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecSelectedHostnames.InvokeAsync(new Akamai.GetAppSecSelectedHostnamesArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.SelectedHostnames = selectedHostnamesAppSecSelectedHostnames.Apply(selectedHostnamesAppSecSelectedHostnames =&gt; selectedHostnamesAppSecSelectedHostnames.Hostnames);
        ///         this.SelectedHostnamesJson = selectedHostnamesAppSecSelectedHostnames.Apply(selectedHostnamesAppSecSelectedHostnames =&gt; selectedHostnamesAppSecSelectedHostnames.HostnamesJson);
        ///         this.SelectedHostnamesOutputText = selectedHostnamesAppSecSelectedHostnames.Apply(selectedHostnamesAppSecSelectedHostnames =&gt; selectedHostnamesAppSecSelectedHostnames.OutputText);
        ///     }
        /// 
        ///     [Output("selectedHostnames")]
        ///     public Output&lt;string&gt; SelectedHostnames { get; set; }
        ///     [Output("selectedHostnamesJson")]
        ///     public Output&lt;string&gt; SelectedHostnamesJson { get; set; }
        ///     [Output("selectedHostnamesOutputText")]
        ///     public Output&lt;string&gt; SelectedHostnamesOutputText { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `hostnames`. List of selected hostnames.
        /// - `hostnames_json`. JSON-formatted list of selected hostnames.
        /// - `output_text`. Tabular report of the selected hostnames.
        /// </summary>
        public static Task<GetAppSecSelectedHostnamesResult> InvokeAsync(GetAppSecSelectedHostnamesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecSelectedHostnamesResult>("akamai:index/getAppSecSelectedHostnames:getAppSecSelectedHostnames", args ?? new GetAppSecSelectedHostnamesArgs(), options.WithDefaults());

        /// <summary>
        /// **Scopes**: Security configuration
        /// 
        /// Returns a list of the hostnames currently protected by the specified security configuration.
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/selected-hostnames](https://techdocs.akamai.com/application-security/reference/get-selected-hostnames)
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
        ///         var selectedHostnamesAppSecSelectedHostnames = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecSelectedHostnames.InvokeAsync(new Akamai.GetAppSecSelectedHostnamesArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.SelectedHostnames = selectedHostnamesAppSecSelectedHostnames.Apply(selectedHostnamesAppSecSelectedHostnames =&gt; selectedHostnamesAppSecSelectedHostnames.Hostnames);
        ///         this.SelectedHostnamesJson = selectedHostnamesAppSecSelectedHostnames.Apply(selectedHostnamesAppSecSelectedHostnames =&gt; selectedHostnamesAppSecSelectedHostnames.HostnamesJson);
        ///         this.SelectedHostnamesOutputText = selectedHostnamesAppSecSelectedHostnames.Apply(selectedHostnamesAppSecSelectedHostnames =&gt; selectedHostnamesAppSecSelectedHostnames.OutputText);
        ///     }
        /// 
        ///     [Output("selectedHostnames")]
        ///     public Output&lt;string&gt; SelectedHostnames { get; set; }
        ///     [Output("selectedHostnamesJson")]
        ///     public Output&lt;string&gt; SelectedHostnamesJson { get; set; }
        ///     [Output("selectedHostnamesOutputText")]
        ///     public Output&lt;string&gt; SelectedHostnamesOutputText { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `hostnames`. List of selected hostnames.
        /// - `hostnames_json`. JSON-formatted list of selected hostnames.
        /// - `output_text`. Tabular report of the selected hostnames.
        /// </summary>
        public static Output<GetAppSecSelectedHostnamesResult> Invoke(GetAppSecSelectedHostnamesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecSelectedHostnamesResult>("akamai:index/getAppSecSelectedHostnames:getAppSecSelectedHostnames", args ?? new GetAppSecSelectedHostnamesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppSecSelectedHostnamesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the protected hosts.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        public GetAppSecSelectedHostnamesArgs()
        {
        }
    }

    public sealed class GetAppSecSelectedHostnamesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the protected hosts.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        public GetAppSecSelectedHostnamesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecSelectedHostnamesResult
    {
        public readonly int ConfigId;
        public readonly ImmutableArray<string> Hostnames;
        public readonly string HostnamesJson;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OutputText;

        [OutputConstructor]
        private GetAppSecSelectedHostnamesResult(
            int configId,

            ImmutableArray<string> hostnames,

            string hostnamesJson,

            string id,

            string outputText)
        {
            ConfigId = configId;
            Hostnames = hostnames;
            HostnamesJson = hostnamesJson;
            Id = id;
            OutputText = outputText;
        }
    }
}
