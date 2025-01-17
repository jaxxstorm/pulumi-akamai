// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Properties
{
    /// <summary>
    /// The `akamai.CpCode` resource lets you create or reuse content provider (CP) codes.  CP codes track web traffic handled by Akamai servers. Akamai gives you a CP code when you purchase a product. You need this code when you activate associated properties.
    /// 
    /// You can create additional CP codes to support more detailed billing and reporting functions.
    /// 
    /// By default, the Akamai Provider uses your existing CP code instead of creating a new one.
    /// 
    /// ## Example Usage
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
    ///         var cpCode = new Akamai.CpCode("cpCode", new Akamai.CpCodeArgs
    ///         {
    ///             ContractId = akamai_contract.Contract.Id,
    ///             GroupId = akamai_group.Group.Id,
    ///             ProductId = "prd_Object_Delivery",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Here's a real-life example that includes other data sources as dependencies:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Akamai = Pulumi.Akamai;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var groupName = "example group name";
    ///         var cpcodeName = "My CP Code";
    ///         var exampleContract = Output.Create(Akamai.GetContract.InvokeAsync(new Akamai.GetContractArgs
    ///         {
    ///             GroupName = groupName,
    ///         }));
    ///         var exampleGroup = exampleContract.Apply(exampleContract =&gt; Output.Create(Akamai.GetGroup.InvokeAsync(new Akamai.GetGroupArgs
    ///         {
    ///             Name = groupName,
    ///             ContractId = exampleContract.Id,
    ///         })));
    ///         var exampleCp = new Akamai.CpCode("exampleCp", new Akamai.CpCodeArgs
    ///         {
    ///             GroupId = exampleGroup.Apply(exampleGroup =&gt; exampleGroup.Id),
    ///             ContractId = exampleContract.Apply(exampleContract =&gt; exampleContract.Id),
    ///             ProductId = "prd_Object_Delivery",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Attributes reference
    /// 
    /// * `id` - The ID of the CP code.
    /// 
    /// ## Import
    /// 
    /// Basic Usagehcl resource "akamai_cp_code" "example" {
    /// 
    /// # (resource arguments)
    /// 
    ///  } You can import your Akamai CP codes using a comma-delimited string of the CP code, contract, and group IDs. You have to enter the IDs in this order`cpcode_id,contract_id,group_id` For example
    /// 
    /// ```sh
    ///  $ pulumi import akamai:properties/cpCode:CpCode example cpc_123,ctr_1-AB123,grp_123
    /// ```
    /// </summary>
    [Obsolete(@"akamai.properties.CpCode has been deprecated in favor of akamai.CpCode")]
    [AkamaiResourceType("akamai:properties/cpCode:CpCode")]
    public partial class CpCode : Pulumi.CustomResource
    {
        /// <summary>
        /// Replaced by `contract_id`. Maintained for legacy purposes.
        /// </summary>
        [Output("contract")]
        public Output<string> Contract { get; private set; } = null!;

        /// <summary>
        /// - (Required) A contract's unique ID, including the `ctr_` prefix.
        /// </summary>
        [Output("contractId")]
        public Output<string> ContractId { get; private set; } = null!;

        /// <summary>
        /// Replaced by `group_id`. Maintained for legacy purposes.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// - (Required) A group's unique ID, including the `grp_` prefix.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// - (Required) A descriptive label for the CP code. If you're creating a new CP code, the name can't include commas, underscores, quotes, or any of these special characters: ^ # %.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Replaced by `product_id`. Maintained for legacy purposes.
        /// </summary>
        [Output("product")]
        public Output<string> Product { get; private set; } = null!;

        [Output("productId")]
        public Output<string> ProductId { get; private set; } = null!;


        /// <summary>
        /// Create a CpCode resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CpCode(string name, CpCodeArgs? args = null, CustomResourceOptions? options = null)
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
        /// - (Required) A group's unique ID, including the `grp_` prefix.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// - (Required) A descriptive label for the CP code. If you're creating a new CP code, the name can't include commas, underscores, quotes, or any of these special characters: ^ # %.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Replaced by `product_id`. Maintained for legacy purposes.
        /// </summary>
        [Input("product")]
        public Input<string>? Product { get; set; }

        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        public CpCodeArgs()
        {
        }
    }

    public sealed class CpCodeState : Pulumi.ResourceArgs
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
        /// - (Required) A group's unique ID, including the `grp_` prefix.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// - (Required) A descriptive label for the CP code. If you're creating a new CP code, the name can't include commas, underscores, quotes, or any of these special characters: ^ # %.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Replaced by `product_id`. Maintained for legacy purposes.
        /// </summary>
        [Input("product")]
        public Input<string>? Product { get; set; }

        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        public CpCodeState()
        {
        }
    }
}
