// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetIamGrantableRoles
    {
        /// <summary>
        /// List which grantable roles you can include in a new custom role or add to an existing custom role.
        /// 
        /// ## Basic usage
        /// 
        /// This example returns the available roles to grant to users:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Akamai.GetIamGrantableRoles.InvokeAsync());
        ///         this.AkaGrantableRolesCount = data.Akamai_iam_grantable_roles.Test.Grantable_roles.Length;
        ///         this.AkaGrantableRoles = data.Akamai_iam_grantable_roles.Test;
        ///     }
        /// 
        ///     [Output("akaGrantableRolesCount")]
        ///     public Output&lt;string&gt; AkaGrantableRolesCount { get; set; }
        ///     [Output("akaGrantableRoles")]
        ///     public Output&lt;string&gt; AkaGrantableRoles { get; set; }
        /// }
        /// ```
        /// 
        /// ## Attributes reference
        /// 
        /// This resource returns this attribute:
        /// 
        /// * `grantable_roles` - Lists which grantable roles you can include in a new custom role or add to an existing custom role.
        ///   * `granted_role_id` - Granted role ID.
        ///   * `name` - Granted role name.
        ///   * `description` - Granted role description.
        /// </summary>
        public static Task<GetIamGrantableRolesResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIamGrantableRolesResult>("akamai:index/getIamGrantableRoles:getIamGrantableRoles", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetIamGrantableRolesResult
    {
        public readonly ImmutableArray<Outputs.GetIamGrantableRolesGrantableRoleResult> GrantableRoles;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetIamGrantableRolesResult(
            ImmutableArray<Outputs.GetIamGrantableRolesGrantableRoleResult> grantableRoles,

            string id)
        {
            GrantableRoles = grantableRoles;
            Id = id;
        }
    }
}
