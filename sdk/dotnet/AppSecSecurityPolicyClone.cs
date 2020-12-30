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
    /// The `akamai.AppSecSecurityPolicyClone` resource allows you to create a new security policy by cloning an existing policy.
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
    ///         var configuration = Output.Create(Akamai.GetAppSecConfiguration.InvokeAsync(new Akamai.GetAppSecConfigurationArgs
    ///         {
    ///             Name = "Akamai Tools",
    ///         }));
    ///         var securityPolicyCloneAppSecSecurityPolicyClone = new Akamai.AppSecSecurityPolicyClone("securityPolicyCloneAppSecSecurityPolicyClone", new Akamai.AppSecSecurityPolicyCloneArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             Version = configuration.Apply(configuration =&gt; configuration.LatestVersion),
    ///             CreateFromSecurityPolicy = "crAP_75829",
    ///             PolicyName = "Test Policy 1",
    ///             PolicyPrefix = "TP1",
    ///         });
    ///         this.SecurityPolicyClone = securityPolicyCloneAppSecSecurityPolicyClone.PolicyId;
    ///     }
    /// 
    ///     [Output("securityPolicyClone")]
    ///     public Output&lt;string&gt; SecurityPolicyClone { get; set; }
    /// }
    /// ```
    /// </summary>
    public partial class AppSecSecurityPolicyClone : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// The ID of the security policy to clone.
        /// </summary>
        [Output("createFromSecurityPolicy")]
        public Output<string> CreateFromSecurityPolicy { get; private set; } = null!;

        /// <summary>
        /// The ID of the new security policy.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        /// <summary>
        /// The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.
        /// </summary>
        [Output("policyName")]
        public Output<string?> PolicyName { get; private set; } = null!;

        /// <summary>
        /// The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated.
        /// </summary>
        [Output("policyPrefix")]
        public Output<string?> PolicyPrefix { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecSecurityPolicyClone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecSecurityPolicyClone(string name, AppSecSecurityPolicyCloneArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecSecurityPolicyClone:AppSecSecurityPolicyClone", name, args ?? new AppSecSecurityPolicyCloneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecSecurityPolicyClone(string name, Input<string> id, AppSecSecurityPolicyCloneState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecSecurityPolicyClone:AppSecSecurityPolicyClone", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecSecurityPolicyClone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecSecurityPolicyClone Get(string name, Input<string> id, AppSecSecurityPolicyCloneState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecSecurityPolicyClone(name, id, state, options);
        }
    }

    public sealed class AppSecSecurityPolicyCloneArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// The ID of the security policy to clone.
        /// </summary>
        [Input("createFromSecurityPolicy", required: true)]
        public Input<string> CreateFromSecurityPolicy { get; set; } = null!;

        /// <summary>
        /// The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        /// <summary>
        /// The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated.
        /// </summary>
        [Input("policyPrefix")]
        public Input<string>? PolicyPrefix { get; set; }

        [Input("version", required: true)]
        public Input<int> Version { get; set; } = null!;

        public AppSecSecurityPolicyCloneArgs()
        {
        }
    }

    public sealed class AppSecSecurityPolicyCloneState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// The ID of the security policy to clone.
        /// </summary>
        [Input("createFromSecurityPolicy")]
        public Input<string>? CreateFromSecurityPolicy { get; set; }

        /// <summary>
        /// The ID of the new security policy.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        /// <summary>
        /// The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated.
        /// </summary>
        [Input("policyPrefix")]
        public Input<string>? PolicyPrefix { get; set; }

        [Input("version")]
        public Input<int>? Version { get; set; }

        public AppSecSecurityPolicyCloneState()
        {
        }
    }
}