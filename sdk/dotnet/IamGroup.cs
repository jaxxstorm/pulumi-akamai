// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    /// <summary>
    /// Use the `akamai.IamGroup` resource to list details about groups. Groups are organizational containers for the objects you use.  Groups can contain other groups, primary objects like properties, and secondary objects like edge hostnames or content provider (CP) codes.
    /// 
    /// ## Basic usage
    /// 
    /// This example returns the policy details based on the policy ID and optionally, a version:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Akamai = Pulumi.Akamai;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Akamai.IamGroup("example", new Akamai.IamGroupArgs
    ///         {
    ///             ParentGroupId = 12345,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Attributes reference
    /// 
    /// This resource returns this attribute:
    /// 
    /// * `sub_groups` - Sub-groups that are related to this group. Each identifier must be an integer.
    /// </summary>
    [AkamaiResourceType("akamai:index/iamGroup:IamGroup")]
    public partial class IamGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Human readable name for a group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for the parent group. Each identifier must be an integer.
        /// </summary>
        [Output("parentGroupId")]
        public Output<int> ParentGroupId { get; private set; } = null!;

        /// <summary>
        /// Subgroups IDs
        /// </summary>
        [Output("subGroups")]
        public Output<ImmutableArray<int>> SubGroups { get; private set; } = null!;


        /// <summary>
        /// Create a IamGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamGroup(string name, IamGroupArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/iamGroup:IamGroup", name, args ?? new IamGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamGroup(string name, Input<string> id, IamGroupState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/iamGroup:IamGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IamGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamGroup Get(string name, Input<string> id, IamGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new IamGroup(name, id, state, options);
        }
    }

    public sealed class IamGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human readable name for a group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A unique identifier for the parent group. Each identifier must be an integer.
        /// </summary>
        [Input("parentGroupId", required: true)]
        public Input<int> ParentGroupId { get; set; } = null!;

        public IamGroupArgs()
        {
        }
    }

    public sealed class IamGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human readable name for a group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A unique identifier for the parent group. Each identifier must be an integer.
        /// </summary>
        [Input("parentGroupId")]
        public Input<int>? ParentGroupId { get; set; }

        [Input("subGroups")]
        private InputList<int>? _subGroups;

        /// <summary>
        /// Subgroups IDs
        /// </summary>
        public InputList<int> SubGroups
        {
            get => _subGroups ?? (_subGroups = new InputList<int>());
            set => _subGroups = value;
        }

        public IamGroupState()
        {
        }
    }
}
