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
    /// Modifies the method used for firewall blocking, and manages the network lists used for IP/Geo firewall blocking.
    /// 
    /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/ip-geo-firewall](https://techdocs.akamai.com/application-security/reference/put-policy-ip-geo-firewall)
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
    ///         var ipGeoBlock = new Akamai.AppSecIPGeo("ipGeoBlock", new Akamai.AppSecIPGeoArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             SecurityPolicyId = "gms1_134637",
    ///             Mode = "block",
    ///             GeoNetworkLists = 
    ///             {
    ///                 "06038_GEO_TEST",
    ///             },
    ///             IpNetworkLists = 
    ///             {
    ///                 "56921_TEST",
    ///             },
    ///             ExceptionIpNetworkLists = 
    ///             {
    ///                 "07126_EXCEPTION_TEST",
    ///             },
    ///         });
    ///         // USE CASE: User wants to update the IP/Geo firewall mode and update the exception list.
    ///         var ipGeoAllow = new Akamai.AppSecIPGeo("ipGeoAllow", new Akamai.AppSecIPGeoArgs
    ///         {
    ///             ConfigId = configuration.Apply(configuration =&gt; configuration.ConfigId),
    ///             SecurityPolicyId = "gms1-090334",
    ///             Mode = "allow",
    ///             ExceptionIpNetworkLists = 
    ///             {
    ///                 "07126_EXCEPTION_TEST",
    ///             },
    ///         });
    ///         this.IpGeoModeBlock = ipGeoBlock.Mode;
    ///         this.BlockGeoNetworkLists = ipGeoBlock.GeoNetworkLists;
    ///         this.BlockIpNetworkLists = ipGeoBlock.IpNetworkLists;
    ///         this.BlockExceptionIpNetworkLists = ipGeoBlock.ExceptionIpNetworkLists;
    ///         this.IpGeoModeAllow = ipGeoAllow.Mode;
    ///         this.AllowExceptionIpNetworkLists = ipGeoAllow.ExceptionIpNetworkLists;
    ///     }
    /// 
    ///     [Output("ipGeoModeBlock")]
    ///     public Output&lt;string&gt; IpGeoModeBlock { get; set; }
    ///     [Output("blockGeoNetworkLists")]
    ///     public Output&lt;string&gt; BlockGeoNetworkLists { get; set; }
    ///     [Output("blockIpNetworkLists")]
    ///     public Output&lt;string&gt; BlockIpNetworkLists { get; set; }
    ///     [Output("blockExceptionIpNetworkLists")]
    ///     public Output&lt;string&gt; BlockExceptionIpNetworkLists { get; set; }
    ///     [Output("ipGeoModeAllow")]
    ///     public Output&lt;string&gt; IpGeoModeAllow { get; set; }
    ///     [Output("allowExceptionIpNetworkLists")]
    ///     public Output&lt;string&gt; AllowExceptionIpNetworkLists { get; set; }
    /// }
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/appSecIPGeo:AppSecIPGeo")]
    public partial class AppSecIPGeo : Pulumi.CustomResource
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the IP/Geo lists being modified.
        /// </summary>
        [Output("configId")]
        public Output<int> ConfigId { get; private set; } = null!;

        /// <summary>
        /// . JSON array of network lists that are always allowed to pass through the firewall, regardless of the value of any other setting.
        /// </summary>
        [Output("exceptionIpNetworkLists")]
        public Output<ImmutableArray<string>> ExceptionIpNetworkLists { get; private set; } = null!;

        /// <summary>
        /// . JSON array of geographic network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
        /// </summary>
        [Output("geoNetworkLists")]
        public Output<ImmutableArray<string>> GeoNetworkLists { get; private set; } = null!;

        /// <summary>
        /// . JSON array of IP network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
        /// </summary>
        [Output("ipNetworkLists")]
        public Output<ImmutableArray<string>> IpNetworkLists { get; private set; } = null!;

        /// <summary>
        /// . Set to **block** to prevent the specified network lists from being allowed through the firewall: all other entities will be allowed to pass through the firewall. Set to **allow** to allow the specified network lists to pass through the firewall; all other entities will be prevented from passing through the firewall.
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// . Unique identifier of the security policy associated with the IP/Geo lists being modified.
        /// </summary>
        [Output("securityPolicyId")]
        public Output<string> SecurityPolicyId { get; private set; } = null!;


        /// <summary>
        /// Create a AppSecIPGeo resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppSecIPGeo(string name, AppSecIPGeoArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/appSecIPGeo:AppSecIPGeo", name, args ?? new AppSecIPGeoArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppSecIPGeo(string name, Input<string> id, AppSecIPGeoState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/appSecIPGeo:AppSecIPGeo", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppSecIPGeo resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppSecIPGeo Get(string name, Input<string> id, AppSecIPGeoState? state = null, CustomResourceOptions? options = null)
        {
            return new AppSecIPGeo(name, id, state, options);
        }
    }

    public sealed class AppSecIPGeoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the IP/Geo lists being modified.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        [Input("exceptionIpNetworkLists")]
        private InputList<string>? _exceptionIpNetworkLists;

        /// <summary>
        /// . JSON array of network lists that are always allowed to pass through the firewall, regardless of the value of any other setting.
        /// </summary>
        public InputList<string> ExceptionIpNetworkLists
        {
            get => _exceptionIpNetworkLists ?? (_exceptionIpNetworkLists = new InputList<string>());
            set => _exceptionIpNetworkLists = value;
        }

        [Input("geoNetworkLists")]
        private InputList<string>? _geoNetworkLists;

        /// <summary>
        /// . JSON array of geographic network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
        /// </summary>
        public InputList<string> GeoNetworkLists
        {
            get => _geoNetworkLists ?? (_geoNetworkLists = new InputList<string>());
            set => _geoNetworkLists = value;
        }

        [Input("ipNetworkLists")]
        private InputList<string>? _ipNetworkLists;

        /// <summary>
        /// . JSON array of IP network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
        /// </summary>
        public InputList<string> IpNetworkLists
        {
            get => _ipNetworkLists ?? (_ipNetworkLists = new InputList<string>());
            set => _ipNetworkLists = value;
        }

        /// <summary>
        /// . Set to **block** to prevent the specified network lists from being allowed through the firewall: all other entities will be allowed to pass through the firewall. Set to **allow** to allow the specified network lists to pass through the firewall; all other entities will be prevented from passing through the firewall.
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        /// <summary>
        /// . Unique identifier of the security policy associated with the IP/Geo lists being modified.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public Input<string> SecurityPolicyId { get; set; } = null!;

        public AppSecIPGeoArgs()
        {
        }
    }

    public sealed class AppSecIPGeoState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the IP/Geo lists being modified.
        /// </summary>
        [Input("configId")]
        public Input<int>? ConfigId { get; set; }

        [Input("exceptionIpNetworkLists")]
        private InputList<string>? _exceptionIpNetworkLists;

        /// <summary>
        /// . JSON array of network lists that are always allowed to pass through the firewall, regardless of the value of any other setting.
        /// </summary>
        public InputList<string> ExceptionIpNetworkLists
        {
            get => _exceptionIpNetworkLists ?? (_exceptionIpNetworkLists = new InputList<string>());
            set => _exceptionIpNetworkLists = value;
        }

        [Input("geoNetworkLists")]
        private InputList<string>? _geoNetworkLists;

        /// <summary>
        /// . JSON array of geographic network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
        /// </summary>
        public InputList<string> GeoNetworkLists
        {
            get => _geoNetworkLists ?? (_geoNetworkLists = new InputList<string>());
            set => _geoNetworkLists = value;
        }

        [Input("ipNetworkLists")]
        private InputList<string>? _ipNetworkLists;

        /// <summary>
        /// . JSON array of IP network lists that, depending on the value of the `mode` argument, will be blocked or allowed through the firewall.
        /// </summary>
        public InputList<string> IpNetworkLists
        {
            get => _ipNetworkLists ?? (_ipNetworkLists = new InputList<string>());
            set => _ipNetworkLists = value;
        }

        /// <summary>
        /// . Set to **block** to prevent the specified network lists from being allowed through the firewall: all other entities will be allowed to pass through the firewall. Set to **allow** to allow the specified network lists to pass through the firewall; all other entities will be prevented from passing through the firewall.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// . Unique identifier of the security policy associated with the IP/Geo lists being modified.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        public AppSecIPGeoState()
        {
        }
    }
}
