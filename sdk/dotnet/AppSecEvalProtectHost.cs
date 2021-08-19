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
    /// The `resource_akamai_appsec_eval_protect_host` resource allows you to move hostnames that you are evaluating to active protection. When you move a hostname from the evaluation hostnames list, it’s added to your security policy as a protected hostname.
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
    ///         var evalHostnames = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecEvalHostnames.InvokeAsync(new Akamai.GetAppSecEvalHostnamesArgs
    ///         {
    ///             ConfigId = configuration.ConfigId,
    ///         })));
    ///         var protectHost = new Akamai.AppSecEvalProtectHost("protectHost", new Akamai.AppSecEvalProtectHostArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             Hostnames = evalHostnames.Apply(evalHostnames =&gt; evalHostnames.Hostnames),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecEvalProtectHost:AppSecEvalProtectHost")]
    public partial class AppSecEvalProtectHost : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// The evaluation hostnames to be moved to active protection.
        /// </summary>
        [Output("hostnames")]
        public Output<ImmutableArray<string>> Hostnames { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecEvalProtectHost resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecEvalProtectHost(string name, AppSecEvalProtectHostArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecEvalProtectHost:AppSecEvalProtectHost", name, args ?? new AppSecEvalProtectHostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecEvalProtectHost(string name, Input<string> id, AppSecEvalProtectHostState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecEvalProtectHost:AppSecEvalProtectHost", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecEvalProtectHost resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecEvalProtectHost Get(string name, Input<string> id, AppSecEvalProtectHostState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecEvalProtectHost(name, id, state, options);
        }
    }

    public sealed class AppSecEvalProtectHostArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        [Input("hostnames", required: true)]
        private InputList<string>? _hostnames;

        /// <summary>
        /// The evaluation hostnames to be moved to active protection.
        /// </summary>
        public InputList<string> Hostnames
        {
            get => _hostnames ?? (_hostnames = new InputList<string>());
            set => _hostnames = value;
        }

        public AppSecEvalProtectHostArgs()
        {
        }
    }

    public sealed class AppSecEvalProtectHostState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        [Input("hostnames")]
        private InputList<string>? _hostnames;

        /// <summary>
        /// The evaluation hostnames to be moved to active protection.
        /// </summary>
        public InputList<string> Hostnames
        {
            get => _hostnames ?? (_hostnames = new InputList<string>());
            set => _hostnames = value;
        }

        public AppSecEvalProtectHostState()
        {
        }
    }
}