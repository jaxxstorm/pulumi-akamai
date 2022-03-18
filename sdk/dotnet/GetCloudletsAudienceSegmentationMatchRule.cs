// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetCloudletsAudienceSegmentationMatchRule
    {
        /// <summary>
        /// Every policy version specifies the match rules that govern how the Cloudlet is used. Matches specify conditions that need to be met in the incoming request.
        /// 
        /// Use the `akamai.getCloudletsAudienceSegmentationMatchRule` data source to build a match rule JSON object for the Audience Segmentation Cloudlet.
        /// 
        /// ## Basic usage
        /// 
        /// This example returns the JSON-encoded rules for the Audience Segmentation Cloudlet:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Akamai.GetCloudletsAudienceSegmentationMatchRule.InvokeAsync(new Akamai.GetCloudletsAudienceSegmentationMatchRuleArgs
        ///         {
        ///             MatchRules = 
        ///             {
        ///                 new Akamai.Inputs.GetCloudletsAudienceSegmentationMatchRuleMatchRuleArgs
        ///                 {
        ///                     ForwardSettings = new Akamai.Inputs.GetCloudletsAudienceSegmentationMatchRuleMatchRuleForwardSettingsArgs
        ///                     {
        ///                         OriginId = "123",
        ///                         PathAndQs = "/test",
        ///                         UseIncomingQueryString = true,
        ///                     },
        ///                     Matches = 
        ///                     {
        ///                         new Akamai.Inputs.GetCloudletsAudienceSegmentationMatchRuleMatchRuleMatchArgs
        ///                         {
        ///                             MatchOperator = "contains",
        ///                             MatchType = "header",
        ///                             ObjectMatchValue = 
        ///                             {
        ///                                 
        ///                                 {
        ///                                     { "name", "cookie" },
        ///                                     { "options", 
        ///                                     {
        ///                                         { "value", 
        ///                                         {
        ///                                             "abcd",
        ///                                         } },
        ///                                     } },
        ///                                     { "type", "object" },
        ///                                 },
        ///                             },
        ///                         },
        ///                     },
        ///                     Name = "rule",
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
        public static Task<GetCloudletsAudienceSegmentationMatchRuleResult> InvokeAsync(GetCloudletsAudienceSegmentationMatchRuleArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudletsAudienceSegmentationMatchRuleResult>("akamai:index/getCloudletsAudienceSegmentationMatchRule:getCloudletsAudienceSegmentationMatchRule", args ?? new GetCloudletsAudienceSegmentationMatchRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Every policy version specifies the match rules that govern how the Cloudlet is used. Matches specify conditions that need to be met in the incoming request.
        /// 
        /// Use the `akamai.getCloudletsAudienceSegmentationMatchRule` data source to build a match rule JSON object for the Audience Segmentation Cloudlet.
        /// 
        /// ## Basic usage
        /// 
        /// This example returns the JSON-encoded rules for the Audience Segmentation Cloudlet:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Akamai.GetCloudletsAudienceSegmentationMatchRule.InvokeAsync(new Akamai.GetCloudletsAudienceSegmentationMatchRuleArgs
        ///         {
        ///             MatchRules = 
        ///             {
        ///                 new Akamai.Inputs.GetCloudletsAudienceSegmentationMatchRuleMatchRuleArgs
        ///                 {
        ///                     ForwardSettings = new Akamai.Inputs.GetCloudletsAudienceSegmentationMatchRuleMatchRuleForwardSettingsArgs
        ///                     {
        ///                         OriginId = "123",
        ///                         PathAndQs = "/test",
        ///                         UseIncomingQueryString = true,
        ///                     },
        ///                     Matches = 
        ///                     {
        ///                         new Akamai.Inputs.GetCloudletsAudienceSegmentationMatchRuleMatchRuleMatchArgs
        ///                         {
        ///                             MatchOperator = "contains",
        ///                             MatchType = "header",
        ///                             ObjectMatchValue = 
        ///                             {
        ///                                 
        ///                                 {
        ///                                     { "name", "cookie" },
        ///                                     { "options", 
        ///                                     {
        ///                                         { "value", 
        ///                                         {
        ///                                             "abcd",
        ///                                         } },
        ///                                     } },
        ///                                     { "type", "object" },
        ///                                 },
        ///                             },
        ///                         },
        ///                     },
        ///                     Name = "rule",
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
        public static Output<GetCloudletsAudienceSegmentationMatchRuleResult> Invoke(GetCloudletsAudienceSegmentationMatchRuleInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudletsAudienceSegmentationMatchRuleResult>("akamai:index/getCloudletsAudienceSegmentationMatchRule:getCloudletsAudienceSegmentationMatchRule", args ?? new GetCloudletsAudienceSegmentationMatchRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudletsAudienceSegmentationMatchRuleArgs : Pulumi.InvokeArgs
    {
        [Input("matchRules")]
        private List<Inputs.GetCloudletsAudienceSegmentationMatchRuleMatchRuleArgs>? _matchRules;

        /// <summary>
        /// - (Optional) A list of Cloudlet-specific match rules for a policy.
        /// </summary>
        public List<Inputs.GetCloudletsAudienceSegmentationMatchRuleMatchRuleArgs> MatchRules
        {
            get => _matchRules ?? (_matchRules = new List<Inputs.GetCloudletsAudienceSegmentationMatchRuleMatchRuleArgs>());
            set => _matchRules = value;
        }

        public GetCloudletsAudienceSegmentationMatchRuleArgs()
        {
        }
    }

    public sealed class GetCloudletsAudienceSegmentationMatchRuleInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("matchRules")]
        private InputList<Inputs.GetCloudletsAudienceSegmentationMatchRuleMatchRuleInputArgs>? _matchRules;

        /// <summary>
        /// - (Optional) A list of Cloudlet-specific match rules for a policy.
        /// </summary>
        public InputList<Inputs.GetCloudletsAudienceSegmentationMatchRuleMatchRuleInputArgs> MatchRules
        {
            get => _matchRules ?? (_matchRules = new InputList<Inputs.GetCloudletsAudienceSegmentationMatchRuleMatchRuleInputArgs>());
            set => _matchRules = value;
        }

        public GetCloudletsAudienceSegmentationMatchRuleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCloudletsAudienceSegmentationMatchRuleResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Json;
        public readonly ImmutableArray<Outputs.GetCloudletsAudienceSegmentationMatchRuleMatchRuleResult> MatchRules;

        [OutputConstructor]
        private GetCloudletsAudienceSegmentationMatchRuleResult(
            string id,

            string json,

            ImmutableArray<Outputs.GetCloudletsAudienceSegmentationMatchRuleMatchRuleResult> matchRules)
        {
            Id = id;
            Json = json;
            MatchRules = matchRules;
        }
    }
}