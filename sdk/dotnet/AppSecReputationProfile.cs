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
    /// Creates or modifies a reputation profile.
    /// Reputation profiles grade the security risk of an IP address based on previous activities associated with that address.
    /// Depending on the reputation score and how your configuration has been set up, requests from a specific IP address can trigger an alert or even be blocked.
    /// 
    /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/reputation-profiles](https://techdocs.akamai.com/application-security/reference/put-reputation-profile)
    /// 
    /// ## Example Usage
    /// 
    /// Basic usage:
    /// 
    /// ```csharp
    /// using System.IO;
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
    ///         var reputationProfile = new Akamai.AppSecReputationProfile("reputationProfile", new Akamai.AppSecReputationProfileArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             ReputationProfile = File.ReadAllText($"{path.Module}/reputation_profile.json"),
    ///         });
    ///         this.ReputationProfileId = reputationProfile.ReputationProfileId;
    ///     }
    /// 
    ///     [Output("reputationProfileId")]
    ///     public Output&lt;string&gt; ReputationProfileId { get; set; }
    /// }
    /// ```
    /// ## Output Options
    /// 
    /// The following options can be used to determine the information returned, and how that returned information is formatted:
    /// 
    /// - `reputation_profile_id`. ID of the newly-created or newly-modified reputation profile.
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecReputationProfile:AppSecReputationProfile")]
    public partial class AppSecReputationProfile : Pulumi.CustomResource
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the reputation profile being modified.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// . Path to a JSON file containing a definition of the reputation profile.
        /// </summary>
        [Output("reputationProfile")]
        public Output<string> ReputationProfile { get; private set; } = null!;

        /// <summary>
        /// Unique identifer of the reputation profile
        /// </summary>
        [Output("reputationProfileId")]
        public Output<int> ReputationProfileId { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecReputationProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecReputationProfile(string name, AppSecReputationProfileArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecReputationProfile:AppSecReputationProfile", name, args ?? new AppSecReputationProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecReputationProfile(string name, Input<string> id, AppSecReputationProfileState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecReputationProfile:AppSecReputationProfile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecReputationProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecReputationProfile Get(string name, Input<string> id, AppSecReputationProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecReputationProfile(name, id, state, options);
        }
    }

    public sealed class AppSecReputationProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the reputation profile being modified.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Path to a JSON file containing a definition of the reputation profile.
        /// </summary>
        [Input("reputationProfile", required: true)]
        public Input<string> ReputationProfile { get; set; } = null!;

        public AppSecReputationProfileArgs()
        {
        }
    }

    public sealed class AppSecReputationProfileState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the reputation profile being modified.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// . Path to a JSON file containing a definition of the reputation profile.
        /// </summary>
        [Input("reputationProfile")]
        public Input<string>? ReputationProfile { get; set; }

        /// <summary>
        /// Unique identifer of the reputation profile
        /// </summary>
        [Input("reputationProfileId")]
        public Input<int>? ReputationProfileId { get; set; }

        public AppSecReputationProfileState()
        {
        }
    }
}
