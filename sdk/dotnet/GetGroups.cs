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
        public static Task<GetGroupsResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupsResult>("akamai:index/getGroups:getGroups", InvokeArgs.Empty, options.WithVersion());
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
