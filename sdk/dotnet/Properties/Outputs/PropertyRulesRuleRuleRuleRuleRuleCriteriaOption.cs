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
    public sealed class PropertyRulesRuleRuleRuleRuleRuleCriteriaOption
    {
        public readonly string Key;
        public readonly string? Value;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private PropertyRulesRuleRuleRuleRuleRuleCriteriaOption(
            string key,

            string? value,

            ImmutableArray<string> values)
        {
            Key = key;
            Value = value;
            Values = values;
        }
    }
}
