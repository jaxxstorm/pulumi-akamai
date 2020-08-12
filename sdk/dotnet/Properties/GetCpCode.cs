// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Properties
{
    public static class GetCpCode
    {
        /// <summary>
        /// Use `akamai.Properties.CpCode` data source to retrieve a group id.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCpCodeResult> InvokeAsync(GetCpCodeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCpCodeResult>("akamai:Properties/getCpCode:getCpCode", args ?? new GetCpCodeArgs(), options.WithVersion());
    }


    public sealed class GetCpCodeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// — (Required) The contract ID
        /// </summary>
        [Input("contract", required: true)]
        public string Contract { get; set; } = null!;

        /// <summary>
        /// — (Required) The group ID
        /// </summary>
        [Input("group", required: true)]
        public string Group { get; set; } = null!;

        /// <summary>
        /// — (Required) The CP code name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetCpCodeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCpCodeResult
    {
        public readonly string Contract;
        public readonly string Group;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetCpCodeResult(
            string contract,

            string group,

            string id,

            string name)
        {
            Contract = contract;
            Group = group;
            Id = id;
            Name = name;
        }
    }
}
