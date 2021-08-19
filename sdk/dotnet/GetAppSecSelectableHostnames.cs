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
        /// Use the `akamai.getAppSecSelectableHostnames` data source to retrieve the list of hostnames that may be protected under a given security configuration. You can specify the list to be retrieved either by supplying the name of a security configuration, or by supplying a group ID and contract ID.
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
        ///         var selectableHostnamesAppSecSelectableHostnames = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecSelectableHostnames.InvokeAsync(new Akamai.GetAppSecSelectableHostnamesArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.SelectableHostnames = selectableHostnamesAppSecSelectableHostnames.Apply(selectableHostnamesAppSecSelectableHostnames =&gt; selectableHostnamesAppSecSelectableHostnames.Hostnames);
        ///         this.SelectableHostnamesJson = selectableHostnamesAppSecSelectableHostnames.Apply(selectableHostnamesAppSecSelectableHostnames =&gt; selectableHostnamesAppSecSelectableHostnames.HostnamesJson);
        ///         this.SelectableHostnamesOutputText = selectableHostnamesAppSecSelectableHostnames.Apply(selectableHostnamesAppSecSelectableHostnames =&gt; selectableHostnamesAppSecSelectableHostnames.OutputText);
        ///         var selectableHostnamesForCreateConfigurationAppSecSelectableHostnames = Output.Create(Akamai.GetAppSecSelectableHostnames.InvokeAsync(new Akamai.GetAppSecSelectableHostnamesArgs
        ///         {
        ///             Contractid = @var.Contractid,
        ///             Groupid = @var.Groupid,
        ///         }));
        ///         this.SelectableHostnamesForCreateConfiguration = selectableHostnamesForCreateConfigurationAppSecSelectableHostnames.Apply(selectableHostnamesForCreateConfigurationAppSecSelectableHostnames =&gt; selectableHostnamesForCreateConfigurationAppSecSelectableHostnames.Hostnames);
        ///         this.SelectableHostnamesForCreateConfigurationJson = selectableHostnamesForCreateConfigurationAppSecSelectableHostnames.Apply(selectableHostnamesForCreateConfigurationAppSecSelectableHostnames =&gt; selectableHostnamesForCreateConfigurationAppSecSelectableHostnames.HostnamesJson);
        ///         this.SelectableHostnamesForCreateConfigurationOutputText = selectableHostnamesForCreateConfigurationAppSecSelectableHostnames.Apply(selectableHostnamesForCreateConfigurationAppSecSelectableHostnames =&gt; selectableHostnamesForCreateConfigurationAppSecSelectableHostnames.OutputText);
        ///     }
        /// 
        ///     [Output("selectableHostnames")]
        ///     public Output&lt;string&gt; SelectableHostnames { get; set; }
        ///     [Output("selectableHostnamesJson")]
        ///     public Output&lt;string&gt; SelectableHostnamesJson { get; set; }
        ///     [Output("selectableHostnamesOutputText")]
        ///     public Output&lt;string&gt; SelectableHostnamesOutputText { get; set; }
        ///     [Output("selectableHostnamesForCreateConfiguration")]
        ///     public Output&lt;string&gt; SelectableHostnamesForCreateConfiguration { get; set; }
        ///     [Output("selectableHostnamesForCreateConfigurationJson")]
        ///     public Output&lt;string&gt; SelectableHostnamesForCreateConfigurationJson { get; set; }
        ///     [Output("selectableHostnamesForCreateConfigurationOutputText")]
        ///     public Output&lt;string&gt; SelectableHostnamesForCreateConfigurationOutputText { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecSelectableHostnamesResult> InvokeAsync(GetAppSecSelectableHostnamesArgs? args = null, InvokeOptions? options = null)
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
        [Input("configId")]
        public int? ConfigId { get; set; }

        /// <summary>
        /// The ID of the contract to use.
        /// </summary>
        [Input("contractid")]
        public string? Contractid { get; set; }

        /// <summary>
        /// The ID of the group to use.
        /// </summary>
        [Input("groupid")]
        public int? Groupid { get; set; }

        public GetAppSecSelectableHostnamesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecSelectableHostnamesResult
    {
        public readonly bool? ActiveInProduction;
        public readonly bool? ActiveInStaging;
        public readonly int? ConfigId;
        public readonly string? Contractid;
        public readonly int? Groupid;
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

        [OutputConstructor]
        private GetAppSecSelectableHostnamesResult(
            bool? activeInProduction,

            bool? activeInStaging,

            int? configId,

            string? contractid,

            int? groupid,

            ImmutableArray<string> hostnames,

            string hostnamesJson,

            string id,

            string outputText)
        {
            ActiveInProduction = activeInProduction;
            ActiveInStaging = activeInStaging;
            ConfigId = configId;
            Contractid = contractid;
            Groupid = groupid;
            Hostnames = hostnames;
            HostnamesJson = hostnamesJson;
            Id = id;
            OutputText = outputText;
        }
    }
}
