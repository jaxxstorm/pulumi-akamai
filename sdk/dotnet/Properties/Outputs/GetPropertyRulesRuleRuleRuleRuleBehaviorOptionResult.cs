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
    public sealed class GetPropertyRulesRuleRuleRuleRuleBehaviorOptionResult
    {
        /// <summary>
        /// — (Required) The option name.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// — (Optional) A single value for the option.
        /// </summary>
        public readonly string? Value;
        /// <summary>
        /// — (Optional) An array of values for the option.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetPropertyRulesRuleRuleRuleRuleBehaviorOptionResult(
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
