// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Properties.Inputs
{

    public sealed class GetPropertyRulesRuleRuleBehaviorArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// — (Required) The name of the behavior.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("options")]
        private List<Inputs.GetPropertyRulesRuleRuleBehaviorOptionArgs>? _options;

        /// <summary>
        /// — (Optional) One or more options for the behavior.
        /// </summary>
        public List<Inputs.GetPropertyRulesRuleRuleBehaviorOptionArgs> Options
        {
            get => _options ?? (_options = new List<Inputs.GetPropertyRulesRuleRuleBehaviorOptionArgs>());
            set => _options = value;
        }

        public GetPropertyRulesRuleRuleBehaviorArgs()
        {
        }
    }
}
