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
        /// **Scopes**: Security configuration; contract; group
        /// 
        /// Returns the list of hostnames that can be (but aren't yet) protected by a security configuration. You can specify the set of hostnames to be retrieved either by supplying the name of a security configuration or by supplying an Akamai group ID and contract ID.
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/selectable-hostnames](https://techdocs.akamai.com/application-security/reference/get-selectable-hostnames)
        /// 
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `hostnames`. List of selectable hostnames.
        /// - `hostnames_json`. JSON-formatted list of selectable hostnames.
        /// - `output_text`. Tabular report of the selectable hostnames showing the name and config_id of the security configuration under which the host is protected in production.
        /// </summary>
        public static Task<GetAppSecSelectableHostnamesResult> InvokeAsync(GetAppSecSelectableHostnamesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecSelectableHostnamesResult>("akamai:index/getAppSecSelectableHostnames:getAppSecSelectableHostnames", args ?? new GetAppSecSelectableHostnamesArgs(), options.WithDefaults());

        /// <summary>
        /// **Scopes**: Security configuration; contract; group
        /// 
        /// Returns the list of hostnames that can be (but aren't yet) protected by a security configuration. You can specify the set of hostnames to be retrieved either by supplying the name of a security configuration or by supplying an Akamai group ID and contract ID.
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/selectable-hostnames](https://techdocs.akamai.com/application-security/reference/get-selectable-hostnames)
        /// 
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `hostnames`. List of selectable hostnames.
        /// - `hostnames_json`. JSON-formatted list of selectable hostnames.
        /// - `output_text`. Tabular report of the selectable hostnames showing the name and config_id of the security configuration under which the host is protected in production.
        /// </summary>
        public static Output<GetAppSecSelectableHostnamesResult> Invoke(GetAppSecSelectableHostnamesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecSelectableHostnamesResult>("akamai:index/getAppSecSelectableHostnames:getAppSecSelectableHostnames", args ?? new GetAppSecSelectableHostnamesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppSecSelectableHostnamesArgs : Pulumi.InvokeArgs
    {
        [Input("activeInProduction")]
        public bool? ActiveInProduction { get; set; }

        [Input("activeInStaging")]
        public bool? ActiveInStaging { get; set; }

        /// <summary>
        /// . Unique identifier of the security configuration you want to return hostname information for. If not included, information is returned for all your security configurations. Note that argument can't be used with either the `contractid` or the `groupid` arguments.
        /// </summary>
        [Input("configId")]
        public int? ConfigId { get; set; }

        /// <summary>
        /// . Unique identifier of the Akamai contract you want to return hostname information for. If not included, information is returned for all the Akamai contracts associated with your account. Note that this argument can't be used with the `config_id` argument.
        /// </summary>
        [Input("contractid")]
        public string? Contractid { get; set; }

        /// <summary>
        /// . Unique identifier of the contract group you want to return hostname information for. If not included, information is returned for all your contract groups. (Or, if you include the `contractid` argument, all the groups associated with the specified contract.) Note that this argument can't be used with the `config_id` argument.
        /// </summary>
        [Input("groupid")]
        public int? Groupid { get; set; }

        public GetAppSecSelectableHostnamesArgs()
        {
        }
    }

    public sealed class GetAppSecSelectableHostnamesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("activeInProduction")]
        public Input<bool>? ActiveInProduction { get; set; }

        [Input("activeInStaging")]
        public Input<bool>? ActiveInStaging { get; set; }

        /// <summary>
        /// . Unique identifier of the security configuration you want to return hostname information for. If not included, information is returned for all your security configurations. Note that argument can't be used with either the `contractid` or the `groupid` arguments.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// . Unique identifier of the Akamai contract you want to return hostname information for. If not included, information is returned for all the Akamai contracts associated with your account. Note that this argument can't be used with the `config_id` argument.
        /// </summary>
        [Input("contractid")]
        public Input<string>? Contractid { get; set; }

        /// <summary>
        /// . Unique identifier of the contract group you want to return hostname information for. If not included, information is returned for all your contract groups. (Or, if you include the `contractid` argument, all the groups associated with the specified contract.) Note that this argument can't be used with the `config_id` argument.
        /// </summary>
        [Input("groupid")]
        public Input<int>? Groupid { get; set; }

        public GetAppSecSelectableHostnamesInvokeArgs()
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
        public readonly ImmutableArray<string> Hostnames;
        public readonly string HostnamesJson;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
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
