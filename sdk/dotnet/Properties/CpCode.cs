// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Properties
{
    [Obsolete(@"akamai.properties.CpCode has been deprecated in favor of akamai.CpCode")]
    [AkamaiResourceType("akamai:properties/cpCode:CpCode")]
    public partial class CpCode : Pulumi.CustomResource
    {
        [Output("contract")]
        public Output<string> Contract { get; private set; } = null!;

        [Output("contractId")]
        public Output<string> ContractId { get; private set; } = null!;

        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("product")]
        public Output<string> Product { get; private set; } = null!;


        /// <summary>
        /// Create a CpCode resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CpCode(string name, CpCodeArgs args, CustomResourceOptions? options = null)
            : base("akamai:properties/cpCode:CpCode", name, args ?? new CpCodeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CpCode(string name, Input<string> id, CpCodeState? state = null, CustomResourceOptions? options = null)
            : base("akamai:properties/cpCode:CpCode", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CpCode resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CpCode Get(string name, Input<string> id, CpCodeState? state = null, CustomResourceOptions? options = null)
        {
            return new CpCode(name, id, state, options);
        }
    }

    public sealed class CpCodeArgs : Pulumi.ResourceArgs
    {
        [Input("contract")]
        public Input<string>? Contract { get; set; }

        [Input("contractId")]
        public Input<string>? ContractId { get; set; }

        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("product", required: true)]
        public Input<string> Product { get; set; } = null!;

        public CpCodeArgs()
        {
        }
    }

    public sealed class CpCodeState : Pulumi.ResourceArgs
    {
        [Input("contract")]
        public Input<string>? Contract { get; set; }

        [Input("contractId")]
        public Input<string>? ContractId { get; set; }

        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("product")]
        public Input<string>? Product { get; set; }

        public CpCodeState()
        {
        }
    }
}
