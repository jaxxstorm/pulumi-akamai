// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Outputs
{

    [OutputType]
    public sealed class GetPropertyRulesTemplateVariableResult
    {
        /// <summary>
        /// The name of the variable used in the template.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of variable: `string`, `number`, `bool`, or `jsonBlock`.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The value of the variable passed as a string.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetPropertyRulesTemplateVariableResult(
            string name,

            string? type,

            string value)
        {
            Name = name;
            Type = type;
            Value = value;
        }
    }
}
