// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetGroups
    {
        /// <summary>
        /// Use the `akamai.getGroups` data source to list groups associated with the [EdgeGrid API client token](https://techdocs.akamai.com/developer/docs/authenticate-with-edgegrid) you're using.
        /// 
        /// ## Basic usage
        /// 
        /// Return groups associated with the EdgeGrid API client token you're using:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var my_example = Output.Create(Akamai.GetGroups.InvokeAsync());
        ///         this.PropertyMatch = my_example;
        ///     }
        /// 
        ///     [Output("propertyMatch")]
        ///     public Output&lt;string&gt; PropertyMatch { get; set; }
        /// }
        /// ```
        /// 
        /// ## Attributes reference
        /// 
        /// This data source returns these attributes:
        /// 
        /// * `groups` - A list of supported groups, with the following attributes:
        ///   * `group_id` - A group's unique ID, including the `grp_` prefix.
        ///   * `group_name` - The name of the group.
        ///   * `parent_group_id` - The ID of the parent group, if applicable.
        ///   * `contract_ids` - An array of strings listing the contract IDs for each group.
        /// </summary>
        public static Task<GetGroupsResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupsResult>("akamai:index/getGroups:getGroups", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetGroupsResult
    {
        public readonly ImmutableArray<Outputs.GetGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetGroupsResult(
            ImmutableArray<Outputs.GetGroupsGroupResult> groups,

            string id)
        {
            Groups = groups;
            Id = id;
        }
    }
}
