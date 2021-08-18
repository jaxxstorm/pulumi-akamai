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
    /// The `akamai.EdgeHostName` resource lets you configure a secure edge hostname. Your
    /// edge hostname determines how requests for your site, app, or content are mapped to
    /// Akamai edge servers.
    /// 
    /// An edge hostname is the CNAME target you use when directing your end user traffic to
    /// Akamai. Each hostname assigned to a property has a corresponding edge hostname.
    /// 
    /// Akamai supports three types of edge hostnames, depending on the level of security
    /// you need for your traffic: Standard TLS, Enhanced TLS, and Shared Certificate. When
    /// entering the `edge_hostname` attribute, you need to include a specific domain suffix
    /// for your edge hostname type:
    /// 
    /// | Edge hostname type | Domain suffix |
    /// |------|-------|
    /// | Enhanced TLS | edgekey.net |
    /// | Standard TLS | edgesuite.net |
    /// | Shared Cert | akamaized.net |
    /// 
    /// For example, if you use Standard TLS and have `www.example.com` as a hostname, your edge hostname would be `www.example.com.edgesuite.net`. If you wanted to use Enhanced TLS with the same hostname, your edge hostname would be `www.example.com.edgekey.net`. See the [Property Manager API (PAPI)](https://developer.akamai.com/api/core_features/property_manager/v1.html#createedgehostnames) for more information.
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
    ///         var provider_demo = new Akamai.EdgeHostName("provider-demo", new Akamai.EdgeHostNameArgs
    ///         {
    ///             ContractId = "ctr_1-AB123",
    ///             EdgeHostname = "www.example.org.edgesuite.net",
    ///             GroupId = "grp_123",
    ///             ProductId = "prd_Object_Delivery",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Argument reference
    /// 
    /// This resource supports these arguments:
    /// 
    /// * `name` - (Required) The name of the edge hostname.
    /// * `contract_id` - (Required) A contract's unique ID, including the `ctr_` prefix.
    /// * `group_id` - (Required) A group's unique ID, including the `grp_` prefix.
    /// * `product_id` - (Required) A product's unique ID, including the `prd_` prefix.
    /// * `edge_hostname` - (Required) One or more edge hostnames. The number of edge hostnames must be less than or equal to the number of public hostnames.
    /// * `certificate` - (Optional) Required only when creating an Enhanced TLS edge hostname. This argument sets the certificate enrollment ID. Edge hostnames for Enhanced TLS end in `edgekey.net`. You can retrieve this ID from the [Certificate Provisioning Service CLI](https://github.com/akamai/cli-cps) .
    /// * `ip_behavior` - (Required) Which version of the IP protocol to use: `IPV4` for version 4 only, `IPV6_PERFORMANCE` for version 6 only, or `IPV6_COMPLIANCE` for both 4 and 6. The default value is `IPV4`.
    /// * `contract` - (Deprecated) Replaced by `contract_id`. Maintained for legacy purposes.
    /// * `group` - (Deprecated) Replaced by `group_id`. Maintained for legacy purposes.
    /// * `product` - (Deprecated) Replaced by `product_id`. Maintained for legacy purposes.
    /// 
    /// ## Attributes reference
    /// 
    /// This resource returns this attribute:
    /// 
    /// * `ip_behavior` - Returns the IP protocol the hostname will use, either `IPV4` for version 4, IPV6_PERFORMANCE`for version 6, or`IPV6_COMPLIANCE` for both.
    /// 
    /// ## Import
    /// 
    /// Basic Usagehcl resource "akamai_edge_hostname" "example" {
    /// 
    /// # (resource arguments)
    /// 
    ///  } You can import Akamai edge hostnames using a comma-delimited string of edge hostname, contract ID, and group ID. You have to enter the values in this order:
    /// 
    /// `edge_hostname, contract_id, group_id` For example
    /// 
    /// ```sh
    ///  $ pulumi import akamai:index/edgeHostName:EdgeHostName example ehn_123,ctr_1-AB123,grp_123
    /// ```
    /// </summary>
    [AkamaiResourceType("akamai:index/edgeHostName:EdgeHostName")]
    public partial class EdgeHostName : Pulumi.CustomResource
    {
        [Output("certificate")]
        public Output<int?> Certificate { get; private set; } = null!;

        [Output("contract")]
        public Output<string> Contract { get; private set; } = null!;

        [Output("contractId")]
        public Output<string> ContractId { get; private set; } = null!;

        [Output("edgeHostname")]
        public Output<string> EdgeHostname { get; private set; } = null!;

        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        [Output("ipBehavior")]
        public Output<string> IpBehavior { get; private set; } = null!;

        [Output("product")]
        public Output<string> Product { get; private set; } = null!;

        [Output("productId")]
        public Output<string> ProductId { get; private set; } = null!;


        /// <summary>
        /// Create a EdgeHostName resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EdgeHostName(string name, EdgeHostNameArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/edgeHostName:EdgeHostName", name, args ?? new EdgeHostNameArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EdgeHostName(string name, Input<string> id, EdgeHostNameState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/edgeHostName:EdgeHostName", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "akamai:properties/edgeHostName:EdgeHostName"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EdgeHostName resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EdgeHostName Get(string name, Input<string> id, EdgeHostNameState? state = null, CustomResourceOptions? options = null)
        {
            return new EdgeHostName(name, id, state, options);
        }
    }

    public sealed class EdgeHostNameArgs : Pulumi.ResourceArgs
    {
        [Input("certificate")]
        public Input<int>? Certificate { get; set; }

        [Input("contract")]
        public Input<string>? Contract { get; set; }

        [Input("contractId")]
        public Input<string>? ContractId { get; set; }

        [Input("edgeHostname", required: true)]
        public Input<string> EdgeHostname { get; set; } = null!;

        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("ipBehavior", required: true)]
        public Input<string> IpBehavior { get; set; } = null!;

        [Input("product")]
        public Input<string>? Product { get; set; }

        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        public EdgeHostNameArgs()
        {
        }
    }

    public sealed class EdgeHostNameState : Pulumi.ResourceArgs
    {
        [Input("certificate")]
        public Input<int>? Certificate { get; set; }

        [Input("contract")]
        public Input<string>? Contract { get; set; }

        [Input("contractId")]
        public Input<string>? ContractId { get; set; }

        [Input("edgeHostname")]
        public Input<string>? EdgeHostname { get; set; }

        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("ipBehavior")]
        public Input<string>? IpBehavior { get; set; }

        [Input("product")]
        public Input<string>? Product { get; set; }

        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        public EdgeHostNameState()
        {
        }
    }
}
