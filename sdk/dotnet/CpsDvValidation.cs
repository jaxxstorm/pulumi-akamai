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
    /// Once you complete the Let's Encrypt challenges, optionally use the `akamai.CpsDvValidation` resource to send the acknowledgement to CPS and inform it that tokens are ready for validation. You can also wait for CPS to check for the tokens, which it does on a regular schedule. Next, CPS automatically deploys the certificate on Staging, and eventually on the Production network.
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
    ///         var example = new Akamai.CpsDvValidation("example", new Akamai.CpsDvValidationArgs
    ///         {
    ///             EnrollmentId = akamai_cps_dv_enrollment.Example.Id,
    ///             Sans = akamai_cps_dv_enrollment.Example.Sans,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Attributes reference
    /// 
    /// * `status` - The status of certificate validation.
    /// </summary>
    [AkamaiResourceType("akamai:index/cpsDvValidation:CpsDvValidation")]
    public partial class CpsDvValidation : Pulumi.CustomResource
    {
        /// <summary>
        /// Unique identifier for the DV certificate enrollment.
        /// </summary>
        [Output("enrollmentId")]
        public Output<int> EnrollmentId { get; private set; } = null!;

        /// <summary>
        /// The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
        /// </summary>
        [Output("sans")]
        public Output<ImmutableArray<string>> Sans { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a CpsDvValidation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CpsDvValidation(string name, CpsDvValidationArgs args, CustomResourceOptions? options = null)
            : base("akamai:index/cpsDvValidation:CpsDvValidation", name, args ?? new CpsDvValidationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CpsDvValidation(string name, Input<string> id, CpsDvValidationState? state = null, CustomResourceOptions? options = null)
            : base("akamai:index/cpsDvValidation:CpsDvValidation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CpsDvValidation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CpsDvValidation Get(string name, Input<string> id, CpsDvValidationState? state = null, CustomResourceOptions? options = null)
        {
            return new CpsDvValidation(name, id, state, options);
        }
    }

    public sealed class CpsDvValidationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier for the DV certificate enrollment.
        /// </summary>
        [Input("enrollmentId", required: true)]
        public Input<int> EnrollmentId { get; set; } = null!;

        [Input("sans")]
        private InputList<string>? _sans;

        /// <summary>
        /// The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
        /// </summary>
        public InputList<string> Sans
        {
            get => _sans ?? (_sans = new InputList<string>());
            set => _sans = value;
        }

        public CpsDvValidationArgs()
        {
        }
    }

    public sealed class CpsDvValidationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier for the DV certificate enrollment.
        /// </summary>
        [Input("enrollmentId")]
        public Input<int>? EnrollmentId { get; set; }

        [Input("sans")]
        private InputList<string>? _sans;

        /// <summary>
        /// The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
        /// </summary>
        public InputList<string> Sans
        {
            get => _sans ?? (_sans = new InputList<string>());
            set => _sans = value;
        }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public CpsDvValidationState()
        {
        }
    }
}
