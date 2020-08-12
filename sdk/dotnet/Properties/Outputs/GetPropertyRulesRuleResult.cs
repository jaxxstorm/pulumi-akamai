// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Properties.Outputs
{

    [OutputType]
    public sealed class GetPropertyRulesRuleResult
    {
        /// <summary>
        /// — (Optional) One or more behaviors to apply to requests that match.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPropertyRulesRuleBehaviorResult> Behaviors;
        public readonly string? CriteriaMatch;
        /// <summary>
        /// — (Optional) Whether the property is a secure (Enhanced TLS) property or not (top-level only).
        /// </summary>
        public readonly bool? IsSecure;
        /// <summary>
        /// — (Optional) Child rules (may be nested five levels deep).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPropertyRulesRuleRuleResult> Rules;
        public readonly ImmutableArray<Outputs.GetPropertyRulesRuleVariableResult> Variables;

        [OutputConstructor]
        private GetPropertyRulesRuleResult(
            ImmutableArray<Outputs.GetPropertyRulesRuleBehaviorResult> behaviors,

            string? criteriaMatch,

            bool? isSecure,

            ImmutableArray<Outputs.GetPropertyRulesRuleRuleResult> rules,

            ImmutableArray<Outputs.GetPropertyRulesRuleVariableResult> variables)
        {
            Behaviors = behaviors;
            CriteriaMatch = criteriaMatch;
            IsSecure = isSecure;
            Rules = rules;
            Variables = variables;
        }
    }
}
