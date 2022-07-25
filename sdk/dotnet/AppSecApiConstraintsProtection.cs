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
    /// **Scopes**: Security policy
    /// 
    /// Enables or disables API constraints protection. These constraints specify the action to be taken when designated API endpoints are invoked.
    /// 
    /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/protections](https://techdocs.akamai.com/application-security/reference/put-policy-protections)
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
    ///             Name = "Documentation",
    ///         }));
    ///         var protection = new Akamai.AppSecApiConstraintsProtection("protection", new Akamai.AppSecApiConstraintsProtectionArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             SecurityPolicyId = "gms1_134637",
    ///             Enabled = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Output Options
    /// 
    /// The following options can be used to determine the information returned, and how that returned information is formatted:
    /// 
    /// - `output_text`. Tabular report showing the current protection settings.
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecApiConstraintsProtection:AppSecApiConstraintsProtection")]
    public partial class AppSecApiConstraintsProtection : Pulumi.CustomResource
    {
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Text Export representation
        /// </summary>
        [Output("outputText")]
        public Output<string> OutputText { get; private set; } = null!;

        [Output("securityPolicyId")]
        public Output<string> SecurityPolicyId { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecApiConstraintsProtection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecApiConstraintsProtection(string name, AppSecApiConstraintsProtectionArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecApiConstraintsProtection:AppSecApiConstraintsProtection", name, args ?? new AppSecApiConstraintsProtectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecApiConstraintsProtection(string name, Input<string> id, AppSecApiConstraintsProtectionState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecApiConstraintsProtection:AppSecApiConstraintsProtection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecApiConstraintsProtection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecApiConstraintsProtection Get(string name, Input<string> id, AppSecApiConstraintsProtectionState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecApiConstraintsProtection(name, id, state, options);
        }
    }

    public sealed class AppSecApiConstraintsProtectionArgs : Pulumi.ResourceArgs
    {
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("securityPolicyId", required: true)]
        public Input<string> SecurityPolicyId { get; set; } = null!;

        public AppSecApiConstraintsProtectionArgs()
        {
        }
    }

    public sealed class AppSecApiConstraintsProtectionState : Pulumi.ResourceArgs
    {
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Text Export representation
        /// </summary>
        [Input("outputText")]
        public Input<string>? OutputText { get; set; }

        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        public AppSecApiConstraintsProtectionState()
        {
        }
    }
}
