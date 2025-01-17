// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetCpCode
    {
        public static Task<GetCpCodeResult> InvokeAsync(GetCpCodeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCpCodeResult>("akamai:index/getCpCode:getCpCode", args ?? new GetCpCodeArgs(), options.WithDefaults());

        public static Output<GetCpCodeResult> Invoke(GetCpCodeInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCpCodeResult>("akamai:index/getCpCode:getCpCode", args ?? new GetCpCodeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCpCodeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Replaced by `contract_id`. Maintained for legacy purposes.
        /// </summary>
        [Input("contract")]
        public string? Contract { get; set; }

        /// <summary>
        /// - (Required) A contract's unique ID, including the `ctr_` prefix.
        /// </summary>
        [Input("contractId")]
        public string? ContractId { get; set; }

        /// <summary>
        /// Replaced by `group_id`. Maintained for legacy purposes.
        /// </summary>
        [Input("group")]
        public string? Group { get; set; }

        /// <summary>
        /// The group's unique ID, including the `grp_` prefix.
        /// </summary>
        [Input("groupId")]
        public string? GroupId { get; set; }

        /// <summary>
        /// The name of the CP code.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetCpCodeArgs()
        {
        }
    }

    public sealed class GetCpCodeInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Replaced by `contract_id`. Maintained for legacy purposes.
        /// </summary>
        [Input("contract")]
        public Input<string>? Contract { get; set; }

        /// <summary>
        /// - (Required) A contract's unique ID, including the `ctr_` prefix.
        /// </summary>
        [Input("contractId")]
        public Input<string>? ContractId { get; set; }

        /// <summary>
        /// Replaced by `group_id`. Maintained for legacy purposes.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// The group's unique ID, including the `grp_` prefix.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the CP code.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetCpCodeInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCpCodeResult
    {
        public readonly string Contract;
        public readonly string ContractId;
        public readonly string Group;
        public readonly string GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly ImmutableArray<string> ProductIds;

        [OutputConstructor]
        private GetCpCodeResult(
            string contract,

            string contractId,

            string group,

            string groupId,

            string id,

            string name,

            ImmutableArray<string> productIds)
        {
            Contract = contract;
            ContractId = contractId;
            Group = group;
            GroupId = groupId;
            Id = id;
            Name = name;
            ProductIds = productIds;
        }
    }
}
