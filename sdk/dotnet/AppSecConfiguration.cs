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
    /// **Scopes**: Contract and group
    /// 
    /// Creates a new WAP (Web Application Protector) or KSD (Kona Site Defender) security configuration. KSD security configurations start out empty (i.e., unconfigured), while WAP configurations are created using preset values. The contract referenced in the request body determines the type of configuration you can create.
    /// 
    /// In addition to manually creating a new configuration, you can use the `create_from_config_id` argument to clone an existing configuration.
    /// 
    /// **Related API Endpoint**: [/appsec/v1/configs](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postconfigurations)
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
    ///         var selectableHostnames = Output.Create(Akamai.GetAppSecSelectableHostnames.InvokeAsync(new Akamai.GetAppSecSelectableHostnamesArgs
    ///         {
    ///             ConfigId = "Documentation",
    ///         }));
    ///         var createConfig = new Akamai.AppSecConfiguration("createConfig", new Akamai.AppSecConfigurationArgs
    ///         {
    ///             Description = "This configuration is used as a testing environment for the documentation team.",
    ///             ContractId = "5-2WA382",
    ///             GroupId = 12198,
    ///             HostNames = 
    ///             {
    ///                 "documentation.akamai.com",
    ///                 "training.akamai.com",
    ///             },
    ///         });
    ///         this.CreateConfigId = createConfig.ConfigId;
    ///         var cloneConfig = new Akamai.AppSecConfiguration("cloneConfig", new Akamai.AppSecConfigurationArgs
    ///         {
    ///             Description = "This configuration is used as a testing environment for the documentation team.",
    ///             CreateFromConfigId = data.Akamai_appsec_configuration.Configuration.Config_id,
    ///             CreateFromVersion = data.Akamai_appsec_configuration.Configuration.Latest_version,
    ///             ContractId = "5-2WA382",
    ///             GroupId = 12198,
    ///             HostNames = selectableHostnames.Apply(selectableHostnames =&gt; selectableHostnames.Hostnames),
    ///         });
    ///         this.CloneConfigId = cloneConfig.ConfigId;
    ///     }
    /// 
    ///     [Output("createConfigId")]
    ///     public Output&lt;string&gt; CreateConfigId { get; set; }
    ///     [Output("cloneConfigId")]
    ///     public Output&lt;string&gt; CloneConfigId { get; set; }
    /// }
    /// ```
    /// ## Output Options
    /// 
    /// The following options can be used to determine the information returned, and how that returned information is formatted:
    /// 
    /// - `config_id`. ID of the new security configuration.
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecConfiguration:AppSecConfiguration")]
    public partial class AppSecConfiguration : Pulumi.CustomResource
    {
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// . Unique identifier of the Akamai contract t associated with the new configuration.
        /// </summary>
        [Output("contractId")]
        public Output<string> ContractId { get; private set; } = null!;

        /// <summary>
        /// . Unique identifier of the existing configuration being cloned in order to create the new configuration.
        /// </summary>
        [Output("createFromConfigId")]
        public Output<int?> CreateFromConfigId { get; private set; } = null!;

        /// <summary>
        /// . Version number of the security configuration being cloned.
        /// </summary>
        [Output("createFromVersion")]
        public Output<int?> CreateFromVersion { get; private set; } = null!;

        /// <summary>
        /// . Brief description of the new configuration.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// . Unique identifier of the contract group associated with the new configuration.
        /// </summary>
        [Output("groupId")]
        public Output<int> GroupId { get; private set; } = null!;

        /// <summary>
        /// . JSON array containing the hostnames to be protected by the new configuration. You must specify at least one hostname in order to create a new configuration.
        /// </summary>
        [Output("hostNames")]
        public Output<ImmutableArray<string>> HostNames { get; private set; } = null!;

        /// <summary>
        /// . Name of the new configuration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecConfiguration(string name, AppSecConfigurationArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecConfiguration:AppSecConfiguration", name, args ?? new AppSecConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecConfiguration(string name, Input<string> id, AppSecConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecConfiguration:AppSecConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecConfiguration Get(string name, Input<string> id, AppSecConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecConfiguration(name, id, state, options);
        }
    }

    public sealed class AppSecConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the Akamai contract t associated with the new configuration.
        /// </summary>
        [Input("contractId", required: true)]
        public Input<string> ContractId { get; set; } = null!;

        /// <summary>
        /// . Unique identifier of the existing configuration being cloned in order to create the new configuration.
        /// </summary>
        [Input("createFromConfigId")]
        public Input<int>? CreateFromConfigId { get; set; }

        /// <summary>
        /// . Version number of the security configuration being cloned.
        /// </summary>
        [Input("createFromVersion")]
        public Input<int>? CreateFromVersion { get; set; }

        /// <summary>
        /// . Brief description of the new configuration.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// . Unique identifier of the contract group associated with the new configuration.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<int> GroupId { get; set; } = null!;

        [Input("hostNames", required: true)]
        private InputList<string>? _hostNames;

        /// <summary>
        /// . JSON array containing the hostnames to be protected by the new configuration. You must specify at least one hostname in order to create a new configuration.
        /// </summary>
        public InputList<string> HostNames
        {
            get => _hostNames ?? (_hostNames = new InputList<string>());
            set => _hostNames = value;
        }

        /// <summary>
        /// . Name of the new configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AppSecConfigurationArgs()
        {
        }
    }

    public sealed class AppSecConfigurationState : Pulumi.ResourceArgs
    {
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        /// <summary>
        /// . Unique identifier of the Akamai contract t associated with the new configuration.
        /// </summary>
        [Input("contractId")]
        public Input<string>? ContractId { get; set; }

        /// <summary>
        /// . Unique identifier of the existing configuration being cloned in order to create the new configuration.
        /// </summary>
        [Input("createFromConfigId")]
        public Input<int>? CreateFromConfigId { get; set; }

        /// <summary>
        /// . Version number of the security configuration being cloned.
        /// </summary>
        [Input("createFromVersion")]
        public Input<int>? CreateFromVersion { get; set; }

        /// <summary>
        /// . Brief description of the new configuration.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// . Unique identifier of the contract group associated with the new configuration.
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        [Input("hostNames")]
        private InputList<string>? _hostNames;

        /// <summary>
        /// . JSON array containing the hostnames to be protected by the new configuration. You must specify at least one hostname in order to create a new configuration.
        /// </summary>
        public InputList<string> HostNames
        {
            get => _hostNames ?? (_hostNames = new InputList<string>());
            set => _hostNames = value;
        }

        /// <summary>
        /// . Name of the new configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AppSecConfigurationState()
        {
        }
    }
}
