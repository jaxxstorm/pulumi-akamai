// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetCloudletsApiPrioritizationMatchRule
    {
        /// <summary>
        /// Every policy version specifies the match rules that govern how the Cloudlet is used. Matches specify conditions that need to be met in the incoming request.
        /// 
        /// Use the `akamai.getCloudletsApiPrioritizationMatchRule` data source to build a match rule JSON object for the API Prioritization Cloudlet.
        /// 
        /// ## Basic usage
        /// 
        /// This example returns the JSON-encoded rules for the API Prioritization Cloudlet:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Akamai.GetCloudletsApiPrioritizationMatchRule.InvokeAsync(new Akamai.GetCloudletsApiPrioritizationMatchRuleArgs
        ///         {
        ///             MatchRules = 
        ///             {
        ///                 new Akamai.Inputs.GetCloudletsApiPrioritizationMatchRuleMatchRuleArgs
        ///                 {
        ///                     Disabled = false,
        ///                     End = 1645037845,
        ///                     MatchUrl = "example.com",
        ///                     Matches = 
        ///                     {
        ///                         new Akamai.Inputs.GetCloudletsApiPrioritizationMatchRuleMatchRuleMatchArgs
        ///                         {
        ///                             CaseSensitive = true,
        ///                             MatchOperator = "equals",
        ///                             MatchType = "method",
        ///                             Negate = false,
        ///                             ObjectMatchValue = 
        ///                             {
        ///                                 
        ///                                 {
        ///                                     { "type", "simple" },
        ///                                     { "value", 
        ///                                     {
        ///                                         "POST",
        ///                                     } },
        ///                                 },
        ///                             },
        ///                         },
        ///                     },
        ///                     Name = "rule",
        ///                     PassThroughPercent = 10,
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
        public static Task<GetCloudletsApiPrioritizationMatchRuleResult> InvokeAsync(GetCloudletsApiPrioritizationMatchRuleArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudletsApiPrioritizationMatchRuleResult>("akamai:index/getCloudletsApiPrioritizationMatchRule:getCloudletsApiPrioritizationMatchRule", args ?? new GetCloudletsApiPrioritizationMatchRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Every policy version specifies the match rules that govern how the Cloudlet is used. Matches specify conditions that need to be met in the incoming request.
        /// 
        /// Use the `akamai.getCloudletsApiPrioritizationMatchRule` data source to build a match rule JSON object for the API Prioritization Cloudlet.
        /// 
        /// ## Basic usage
        /// 
        /// This example returns the JSON-encoded rules for the API Prioritization Cloudlet:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Akamai.GetCloudletsApiPrioritizationMatchRule.InvokeAsync(new Akamai.GetCloudletsApiPrioritizationMatchRuleArgs
        ///         {
        ///             MatchRules = 
        ///             {
        ///                 new Akamai.Inputs.GetCloudletsApiPrioritizationMatchRuleMatchRuleArgs
        ///                 {
        ///                     Disabled = false,
        ///                     End = 1645037845,
        ///                     MatchUrl = "example.com",
        ///                     Matches = 
        ///                     {
        ///                         new Akamai.Inputs.GetCloudletsApiPrioritizationMatchRuleMatchRuleMatchArgs
        ///                         {
        ///                             CaseSensitive = true,
        ///                             MatchOperator = "equals",
        ///                             MatchType = "method",
        ///                             Negate = false,
        ///                             ObjectMatchValue = 
        ///                             {
        ///                                 
        ///                                 {
        ///                                     { "type", "simple" },
        ///                                     { "value", 
        ///                                     {
        ///                                         "POST",
        ///                                     } },
        ///                                 },
        ///                             },
        ///                         },
        ///                     },
        ///                     Name = "rule",
        ///                     PassThroughPercent = 10,
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
        public static Output<GetCloudletsApiPrioritizationMatchRuleResult> Invoke(GetCloudletsApiPrioritizationMatchRuleInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudletsApiPrioritizationMatchRuleResult>("akamai:index/getCloudletsApiPrioritizationMatchRule:getCloudletsApiPrioritizationMatchRule", args ?? new GetCloudletsApiPrioritizationMatchRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudletsApiPrioritizationMatchRuleArgs : Pulumi.InvokeArgs
    {
        [Input("matchRules")]
        private List<Inputs.GetCloudletsApiPrioritizationMatchRuleMatchRuleArgs>? _matchRules;

        /// <summary>
        /// - (Optional) A list of Cloudlet-specific match rules for a policy.
        /// </summary>
        public List<Inputs.GetCloudletsApiPrioritizationMatchRuleMatchRuleArgs> MatchRules
        {
            get => _matchRules ?? (_matchRules = new List<Inputs.GetCloudletsApiPrioritizationMatchRuleMatchRuleArgs>());
            set => _matchRules = value;
        }

        public GetCloudletsApiPrioritizationMatchRuleArgs()
        {
        }
    }

    public sealed class GetCloudletsApiPrioritizationMatchRuleInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("matchRules")]
        private InputList<Inputs.GetCloudletsApiPrioritizationMatchRuleMatchRuleInputArgs>? _matchRules;

        /// <summary>
        /// - (Optional) A list of Cloudlet-specific match rules for a policy.
        /// </summary>
        public InputList<Inputs.GetCloudletsApiPrioritizationMatchRuleMatchRuleInputArgs> MatchRules
        {
            get => _matchRules ?? (_matchRules = new InputList<Inputs.GetCloudletsApiPrioritizationMatchRuleMatchRuleInputArgs>());
            set => _matchRules = value;
        }

        public GetCloudletsApiPrioritizationMatchRuleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCloudletsApiPrioritizationMatchRuleResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Json;
        public readonly ImmutableArray<Outputs.GetCloudletsApiPrioritizationMatchRuleMatchRuleResult> MatchRules;

        [OutputConstructor]
        private GetCloudletsApiPrioritizationMatchRuleResult(
            string id,

            string json,

            ImmutableArray<Outputs.GetCloudletsApiPrioritizationMatchRuleMatchRuleResult> matchRules)
        {
            Id = id;
            Json = json;
            MatchRules = matchRules;
        }
    }
}