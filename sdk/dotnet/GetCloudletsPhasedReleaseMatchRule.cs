// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetCloudletsPhasedReleaseMatchRule
    {
        /// <summary>
        /// Every policy version specifies the match rules that govern how the Cloudlet is used. Matches specify conditions that need to be met in the incoming request.
        /// 
        /// Use the `akamai.getCloudletsPhasedReleaseMatchRule` data source to build a match rule JSON object for the Phased Release Cloudlet.
        /// 
        /// ## Basic usage
        /// 
        /// This example returns the JSON-encoded rules for the Phased Release Cloudlet:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Akamai.GetCloudletsPhasedReleaseMatchRule.InvokeAsync(new Akamai.GetCloudletsPhasedReleaseMatchRuleArgs
        ///         {
        ///             MatchRules = 
        ///             {
        ///                 new Akamai.Inputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleArgs
        ///                 {
        ///                     End = 1645037845,
        ///                     ForwardSettings = new Akamai.Inputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsArgs
        ///                     {
        ///                         OriginId = "1234",
        ///                         Percent = 100,
        ///                     },
        ///                     Matches = 
        ///                     {
        ///                         new Akamai.Inputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchArgs
        ///                         {
        ///                             CaseSensitive = false,
        ///                             CheckIps = "CONNECTING_IP XFF_HEADERS",
        ///                             MatchOperator = "equals",
        ///                             MatchType = "header",
        ///                             Negate = false,
        ///                             ObjectMatchValue = 
        ///                             {
        ///                                 
        ///                                 {
        ///                                     { "name", "Content-Type" },
        ///                                     { "options", 
        ///                                     {
        ///                                         { "value", 
        ///                                         {
        ///                                             "application/json",
        ///                                         } },
        ///                                     } },
        ///                                     { "type", "object" },
        ///                                 },
        ///                             },
        ///                         },
        ///                     },
        ///                     Name = "rule",
        ///                     Start = 1644865045,
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// ## Attributes reference
        /// 
        /// This data source returns these attributes:
        /// 
        /// * `type` - The type of Cloudlet the rule is for.
        /// * `json` - A `match_rules` JSON structure generated from the API schema that defines the rules for this policy.
        /// </summary>
        public static Task<GetCloudletsPhasedReleaseMatchRuleResult> InvokeAsync(GetCloudletsPhasedReleaseMatchRuleArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudletsPhasedReleaseMatchRuleResult>("akamai:index/getCloudletsPhasedReleaseMatchRule:getCloudletsPhasedReleaseMatchRule", args ?? new GetCloudletsPhasedReleaseMatchRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Every policy version specifies the match rules that govern how the Cloudlet is used. Matches specify conditions that need to be met in the incoming request.
        /// 
        /// Use the `akamai.getCloudletsPhasedReleaseMatchRule` data source to build a match rule JSON object for the Phased Release Cloudlet.
        /// 
        /// ## Basic usage
        /// 
        /// This example returns the JSON-encoded rules for the Phased Release Cloudlet:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Akamai.GetCloudletsPhasedReleaseMatchRule.InvokeAsync(new Akamai.GetCloudletsPhasedReleaseMatchRuleArgs
        ///         {
        ///             MatchRules = 
        ///             {
        ///                 new Akamai.Inputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleArgs
        ///                 {
        ///                     End = 1645037845,
        ///                     ForwardSettings = new Akamai.Inputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsArgs
        ///                     {
        ///                         OriginId = "1234",
        ///                         Percent = 100,
        ///                     },
        ///                     Matches = 
        ///                     {
        ///                         new Akamai.Inputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchArgs
        ///                         {
        ///                             CaseSensitive = false,
        ///                             CheckIps = "CONNECTING_IP XFF_HEADERS",
        ///                             MatchOperator = "equals",
        ///                             MatchType = "header",
        ///                             Negate = false,
        ///                             ObjectMatchValue = 
        ///                             {
        ///                                 
        ///                                 {
        ///                                     { "name", "Content-Type" },
        ///                                     { "options", 
        ///                                     {
        ///                                         { "value", 
        ///                                         {
        ///                                             "application/json",
        ///                                         } },
        ///                                     } },
        ///                                     { "type", "object" },
        ///                                 },
        ///                             },
        ///                         },
        ///                     },
        ///                     Name = "rule",
        ///                     Start = 1644865045,
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// ## Attributes reference
        /// 
        /// This data source returns these attributes:
        /// 
        /// * `type` - The type of Cloudlet the rule is for.
        /// * `json` - A `match_rules` JSON structure generated from the API schema that defines the rules for this policy.
        /// </summary>
        public static Output<GetCloudletsPhasedReleaseMatchRuleResult> Invoke(GetCloudletsPhasedReleaseMatchRuleInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudletsPhasedReleaseMatchRuleResult>("akamai:index/getCloudletsPhasedReleaseMatchRule:getCloudletsPhasedReleaseMatchRule", args ?? new GetCloudletsPhasedReleaseMatchRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudletsPhasedReleaseMatchRuleArgs : Pulumi.InvokeArgs
    {
        [Input("matchRules")]
        private List<Inputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleArgs>? _matchRules;

        /// <summary>
        /// - (Optional) A list of Cloudlet-specific match rules for a policy.
        /// </summary>
        public List<Inputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleArgs> MatchRules
        {
            get => _matchRules ?? (_matchRules = new List<Inputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleArgs>());
            set => _matchRules = value;
        }

        public GetCloudletsPhasedReleaseMatchRuleArgs()
        {
        }
    }

    public sealed class GetCloudletsPhasedReleaseMatchRuleInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("matchRules")]
        private InputList<Inputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleInputArgs>? _matchRules;

        /// <summary>
        /// - (Optional) A list of Cloudlet-specific match rules for a policy.
        /// </summary>
        public InputList<Inputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleInputArgs> MatchRules
        {
            get => _matchRules ?? (_matchRules = new InputList<Inputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleInputArgs>());
            set => _matchRules = value;
        }

        public GetCloudletsPhasedReleaseMatchRuleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCloudletsPhasedReleaseMatchRuleResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Json;
        public readonly ImmutableArray<Outputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleResult> MatchRules;

        [OutputConstructor]
        private GetCloudletsPhasedReleaseMatchRuleResult(
            string id,

            string json,

            ImmutableArray<Outputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleResult> matchRules)
        {
            Id = id;
            Json = json;
            MatchRules = matchRules;
        }
    }
}