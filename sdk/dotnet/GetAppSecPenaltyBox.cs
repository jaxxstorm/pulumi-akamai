// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecPenaltyBox
    {
        /// <summary>
        /// Use the `akamai.AppSecPenaltyBox` data source to retrieve the penalty box settings for a specified security policy.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic usage:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var configuration = Output.Create(Akamai.GetAppSecConfiguration.InvokeAsync(new Akamai.GetAppSecConfigurationArgs
        ///         {
        ///             Name = @var.Security_configuration,
        ///         }));
        ///         var penaltyBox = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecPenaltyBox.InvokeAsync(new Akamai.GetAppSecPenaltyBoxArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = @var.Security_policy_id,
        ///         })));
        ///         this.PenaltyBoxAction = penaltyBox.Apply(penaltyBox =&gt; penaltyBox.Action);
        ///         this.PenaltyBoxEnabled = penaltyBox.Apply(penaltyBox =&gt; penaltyBox.Enabled);
        ///         this.PenaltyBoxText = penaltyBox.Apply(penaltyBox =&gt; penaltyBox.OutputText);
        ///     }
        /// 
        ///     [Output("penaltyBoxAction")]
        ///     public Output&lt;string&gt; PenaltyBoxAction { get; set; }
        ///     [Output("penaltyBoxEnabled")]
        ///     public Output&lt;string&gt; PenaltyBoxEnabled { get; set; }
        ///     [Output("penaltyBoxText")]
        ///     public Output&lt;string&gt; PenaltyBoxText { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecPenaltyBoxResult> InvokeAsync(GetAppSecPenaltyBoxArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecPenaltyBoxResult>("akamai:index/getAppSecPenaltyBox:getAppSecPenaltyBox", args ?? new GetAppSecPenaltyBoxArgs(), options.WithVersion());
    }


    public sealed class GetAppSecPenaltyBoxArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// The ID of the security policy to use.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public string SecurityPolicyId { get; set; } = null!;

        public GetAppSecPenaltyBoxArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecPenaltyBoxResult
    {
        /// <summary>
        /// The action for the penalty box: `alert`, `deny`, or `none`.
        /// </summary>
        public readonly string Action;
        public readonly int ConfigId;
        /// <summary>
        /// Either `true` or `false`, indicating whether penalty box protection is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A tabular display of the `action` and `enabled` information.
        /// </summary>
        public readonly string OutputText;
        public readonly string SecurityPolicyId;

        [OutputConstructor]
        private GetAppSecPenaltyBoxResult(
            string action,

            int configId,

            bool enabled,

            string id,

            string outputText,

            string securityPolicyId)
        {
            Action = action;
            ConfigId = configId;
            Enabled = enabled;
            Id = id;
            OutputText = outputText;
            SecurityPolicyId = securityPolicyId;
        }
    }
}
