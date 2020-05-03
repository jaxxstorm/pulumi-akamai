// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Properties.Inputs
{

    public sealed class PropertyRulesRuleArgs : Pulumi.ResourceArgs
    {
        [Input("behaviors")]
        private InputList<Inputs.PropertyRulesRuleBehaviorArgs>? _behaviors;
        public InputList<Inputs.PropertyRulesRuleBehaviorArgs> Behaviors
        {
            get => _behaviors ?? (_behaviors = new InputList<Inputs.PropertyRulesRuleBehaviorArgs>());
            set => _behaviors = value;
        }

        [Input("criteriaMatch")]
        public Input<string>? CriteriaMatch { get; set; }

        [Input("isSecure")]
        public Input<bool>? IsSecure { get; set; }

        [Input("rules")]
        private InputList<Inputs.PropertyRulesRuleRuleArgs>? _rules;
        public InputList<Inputs.PropertyRulesRuleRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.PropertyRulesRuleRuleArgs>());
            set => _rules = value;
        }

        [Input("variables")]
        private InputList<Inputs.PropertyRulesRuleVariableArgs>? _variables;
        public InputList<Inputs.PropertyRulesRuleVariableArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.PropertyRulesRuleVariableArgs>());
            set => _variables = value;
        }

        public PropertyRulesRuleArgs()
        {
        }
    }
}
