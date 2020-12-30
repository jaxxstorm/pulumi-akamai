// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecSelectableHostnames
    {
        /// <summary>
        /// Use the `akamai.getAppSecSelectableHostnames` data source to retrieve the list of hostnames that may be protected under a given security configuration version.
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
        ///             Name = "Akamai Tools",
        ///         }));
        ///         var selectableHostnamesAppSecSelectableHostnames = Output.Tuple(configuration, configuration).Apply(values =&gt;
        ///         {
        ///             var configuration = values.Item1;
        ///             var configuration1 = values.Item2;
        ///             return Output.Create(Akamai.GetAppSecSelectableHostnames.InvokeAsync(new Akamai.GetAppSecSelectableHostnamesArgs
        ///             {
        ///                 ConfigId = configuration.ConfigId,
        ///                 Version = configuration1.LatestVersion,
        ///             }));
        ///         });
        ///         this.SelectableHostnames = selectableHostnamesAppSecSelectableHostnames.Apply(selectableHostnamesAppSecSelectableHostnames =&gt; selectableHostnamesAppSecSelectableHostnames.Hostnames);
        ///         this.SelectableHostnamesJson = selectableHostnamesAppSecSelectableHostnames.Apply(selectableHostnamesAppSecSelectableHostnames =&gt; selectableHostnamesAppSecSelectableHostnames.HostnamesJson);
        ///         this.SelectableHostnamesOutputText = selectableHostnamesAppSecSelectableHostnames.Apply(selectableHostnamesAppSecSelectableHostnames =&gt; selectableHostnamesAppSecSelectableHostnames.OutputText);
        ///     }
        /// 
        ///     [Output("selectableHostnames")]
        ///     public Output&lt;string&gt; SelectableHostnames { get; set; }
        ///     [Output("selectableHostnamesJson")]
        ///     public Output&lt;string&gt; SelectableHostnamesJson { get; set; }
        ///     [Output("selectableHostnamesOutputText")]
        ///     public Output&lt;string&gt; SelectableHostnamesOutputText { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecSelectableHostnamesResult> InvokeAsync(GetAppSecSelectableHostnamesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecSelectableHostnamesResult>("akamai:index/getAppSecSelectableHostnames:getAppSecSelectableHostnames", args ?? new GetAppSecSelectableHostnamesArgs(), options.WithVersion());
    }


    public sealed class GetAppSecSelectableHostnamesArgs : Pulumi.InvokeArgs
    {
        [Input("activeInProduction")]
        public bool? ActiveInProduction { get; set; }

        [Input("activeInStaging")]
        public bool? ActiveInStaging { get; set; }

        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// The version number of the security configuration to use.
        /// </summary>
        [Input("version", required: true)]
        public int Version { get; set; }

        public GetAppSecSelectableHostnamesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecSelectableHostnamesResult
    {
        public readonly bool? ActiveInProduction;
        public readonly bool? ActiveInStaging;
        public readonly int ConfigId;
        /// <summary>
        /// The list of selectable hostnames.
        /// </summary>
        public readonly ImmutableArray<string> Hostnames;
        /// <summary>
        /// The list of selectable hostnames in json format.
        /// </summary>
        public readonly string HostnamesJson;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A tabular display of the selectable hostnames showing the name and config_id of the security configuration under which the host is protected in production, or '-' if the host is not protected in production.
        /// </summary>
        public readonly string OutputText;
        public readonly int Version;

        [OutputConstructor]
        private GetAppSecSelectableHostnamesResult(
            bool? activeInProduction,

            bool? activeInStaging,

            int configId,

            ImmutableArray<string> hostnames,

            string hostnamesJson,

            string id,

            string outputText,

            int version)
        {
            ActiveInProduction = activeInProduction;
            ActiveInStaging = activeInStaging;
            ConfigId = configId;
            Hostnames = hostnames;
            HostnamesJson = hostnamesJson;
            Id = id;
            OutputText = outputText;
            Version = version;
        }
    }
}