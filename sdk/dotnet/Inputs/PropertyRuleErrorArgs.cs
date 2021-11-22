// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Inputs
{

    public sealed class PropertyRuleErrorArgs : Pulumi.ResourceArgs
    {
        [Input("behaviorName")]
        public Input<string>? BehaviorName { get; set; }

        [Input("detail")]
        public Input<string>? Detail { get; set; }

        [Input("errorLocation")]
        public Input<string>? ErrorLocation { get; set; }

        [Input("instance")]
        public Input<string>? Instance { get; set; }

        [Input("statusCode")]
        public Input<int>? StatusCode { get; set; }

        [Input("title")]
        public Input<string>? Title { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public PropertyRuleErrorArgs()
        {
        }
    }
}
