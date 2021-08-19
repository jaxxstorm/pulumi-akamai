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
    /// Use the `akamai.AppSecSecurityPolicy` resource to create a new security policy.
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
    ///             Name = @var.Security_configuration,
    ///         }));
    ///         var securityPolicyCreateAppSecSecurityPolicy = new Akamai.AppSecSecurityPolicy("securityPolicyCreateAppSecSecurityPolicy", new Akamai.AppSecSecurityPolicyArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             DefaultSettings = @var.Default_settings,
    ///             SecurityPolicyName = @var.Policy_name,
    ///             SecurityPolicyPrefix = @var.Policy_prefix,
    ///         });
    ///         this.SecurityPolicyCreate = securityPolicyCreateAppSecSecurityPolicy.SecurityPolicyId;
    ///     }
    /// 
    ///     [Output("securityPolicyCreate")]
    ///     public Output&lt;string&gt; SecurityPolicyCreate { get; set; }
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecSecurityPolicy:AppSecSecurityPolicy")]
    public partial class AppSecSecurityPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The configuration ID to use.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// The ID of the security policy to clone from.
        /// </summary>
        [Output("createFromSecurityPolicyId")]
        public Output<string?> CreateFromSecurityPolicyId { get; private set; } = null!;

        /// <summary>
        /// Whether the new policy should use the default settings. If not supplied, defaults to true.
        /// </summary>
        [Output("defaultSettings")]
        public Output<bool?> DefaultSettings { get; private set; } = null!;

        /// <summary>
        /// The ID of the newly created security policy.
        /// </summary>
        [Output("securityPolicyId")]
        public Output<string> SecurityPolicyId { get; private set; } = null!;

        /// <summary>
        /// The name of the new security policy.
        /// </summary>
        [Output("securityPolicyName")]
        public Output<string> SecurityPolicyName { get; private set; } = null!;

        /// <summary>
        /// The four-character alphanumeric string prefix for the policy ID.
        /// </summary>
        [Output("securityPolicyPrefix")]
        public Output<string> SecurityPolicyPrefix { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecSecurityPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecSecurityPolicy(string name, AppSecSecurityPolicyArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecSecurityPolicy:AppSecSecurityPolicy", name, args ?? new AppSecSecurityPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecSecurityPolicy(string name, Input<string> id, AppSecSecurityPolicyState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecSecurityPolicy:AppSecSecurityPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecSecurityPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecSecurityPolicy Get(string name, Input<string> id, AppSecSecurityPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecSecurityPolicy(name, id, state, options);
        }
    }

    public sealed class AppSecSecurityPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration ID to use.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// The ID of the security policy to clone from.
        /// </summary>
        [Input("createFromSecurityPolicyId")]
        public Input<string>? CreateFromSecurityPolicyId { get; set; }

        /// <summary>
        /// Whether the new policy should use the default settings. If not supplied, defaults to true.
        /// </summary>
        [Input("defaultSettings")]
        public Input<bool>? DefaultSettings { get; set; }

        /// <summary>
        /// The name of the new security policy.
        /// </summary>
        [Input("securityPolicyName", required: true)]
        public Input<string> SecurityPolicyName { get; set; } = null!;

        /// <summary>
        /// The four-character alphanumeric string prefix for the policy ID.
        /// </summary>
        [Input("securityPolicyPrefix", required: true)]
        public Input<string> SecurityPolicyPrefix { get; set; } = null!;

        public AppSecSecurityPolicyArgs()
        {
        }
    }

    public sealed class AppSecSecurityPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration ID to use.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// The ID of the security policy to clone from.
        /// </summary>
        [Input("createFromSecurityPolicyId")]
        public Input<string>? CreateFromSecurityPolicyId { get; set; }

        /// <summary>
        /// Whether the new policy should use the default settings. If not supplied, defaults to true.
        /// </summary>
        [Input("defaultSettings")]
        public Input<bool>? DefaultSettings { get; set; }

        /// <summary>
        /// The ID of the newly created security policy.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        /// <summary>
        /// The name of the new security policy.
        /// </summary>
        [Input("securityPolicyName")]
        public Input<string>? SecurityPolicyName { get; set; }

        /// <summary>
        /// The four-character alphanumeric string prefix for the policy ID.
        /// </summary>
        [Input("securityPolicyPrefix")]
        public Input<string>? SecurityPolicyPrefix { get; set; }

        public AppSecSecurityPolicyState()
        {
        }
    }
}