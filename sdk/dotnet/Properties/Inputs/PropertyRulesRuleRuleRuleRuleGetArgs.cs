// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Properties.Inputs
{

    public sealed class PropertyRulesRuleRuleRuleRuleGetArgs : Pulumi.ResourceArgs
    {
        [Input("behaviors")]
        private InputList<Inputs.PropertyRulesRuleRuleRuleRuleBehaviorGetArgs>? _behaviors;
        public InputList<Inputs.PropertyRulesRuleRuleRuleRuleBehaviorGetArgs> Behaviors
        {
            get => _behaviors ?? (_behaviors = new InputList<Inputs.PropertyRulesRuleRuleRuleRuleBehaviorGetArgs>());
            set => _behaviors = value;
        }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("criteriaMatch")]
        public Input<string>? CriteriaMatch { get; set; }

        [Input("criterias")]
        private InputList<Inputs.PropertyRulesRuleRuleRuleRuleCriteriaGetArgs>? _criterias;
        public InputList<Inputs.PropertyRulesRuleRuleRuleRuleCriteriaGetArgs> Criterias
        {
            get => _criterias ?? (_criterias = new InputList<Inputs.PropertyRulesRuleRuleRuleRuleCriteriaGetArgs>());
            set => _criterias = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("rules")]
        private InputList<Inputs.PropertyRulesRuleRuleRuleRuleRuleGetArgs>? _rules;
        public InputList<Inputs.PropertyRulesRuleRuleRuleRuleRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.PropertyRulesRuleRuleRuleRuleRuleGetArgs>());
            set => _rules = value;
        }

        public PropertyRulesRuleRuleRuleRuleGetArgs()
        {
        }
    }
}
