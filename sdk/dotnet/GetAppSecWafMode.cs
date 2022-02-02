// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecWafMode
    {
        /// <summary>
        /// **Scopes**: Security policy
        /// 
        /// Returns information about how the Kona Rule Set rules associated with a security configuration and security policy are updated. The WAF (Web Application Firewall) mode determines whether Kona Rule Sets are automatically updated as part of automated attack groups (`mode = AAG`) or whether you must periodically check for new rules and then manually update those rules yourself (`mode = KRS`).
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/mode](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getmode)
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
        ///             Name = "Documentation",
        ///         }));
        ///         var wafMode = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecWafMode.InvokeAsync(new Akamai.GetAppSecWafModeArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = "gms1_134637",
        ///         })));
        ///         this.WafModeMode = wafMode.Apply(wafMode =&gt; wafMode.Mode);
        ///         this.WafModeCurrentRuleset = wafMode.Apply(wafMode =&gt; wafMode.CurrentRuleset);
        ///         this.WafModeEvalStatus = wafMode.Apply(wafMode =&gt; wafMode.EvalStatus);
        ///         this.WafModeEvalRuleset = wafMode.Apply(wafMode =&gt; wafMode.EvalRuleset);
        ///         this.WafModeEvalExpirationDate = wafMode.Apply(wafMode =&gt; wafMode.EvalExpirationDate);
        ///         this.WafModeText = wafMode.Apply(wafMode =&gt; wafMode.OutputText);
        ///         this.WafModeJson = wafMode.Apply(wafMode =&gt; wafMode.Json);
        ///     }
        /// 
        ///     [Output("wafModeMode")]
        ///     public Output&lt;string&gt; WafModeMode { get; set; }
        ///     [Output("wafModeCurrentRuleset")]
        ///     public Output&lt;string&gt; WafModeCurrentRuleset { get; set; }
        ///     [Output("wafModeEvalStatus")]
        ///     public Output&lt;string&gt; WafModeEvalStatus { get; set; }
        ///     [Output("wafModeEvalRuleset")]
        ///     public Output&lt;string&gt; WafModeEvalRuleset { get; set; }
        ///     [Output("wafModeEvalExpirationDate")]
        ///     public Output&lt;string&gt; WafModeEvalExpirationDate { get; set; }
        ///     [Output("wafModeText")]
        ///     public Output&lt;string&gt; WafModeText { get; set; }
        ///     [Output("wafModeJson")]
        ///     public Output&lt;string&gt; WafModeJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `mode`. Security policy mode, either **KRS** (update manually) or **AAG** (update automatically), For organizations running the Adaptive Security Engine (ASE) beta, you'll get back **ASE_AUTO** for automatic updates or **ASE_MANUAL** for manual updates. Please contact your Akamai representative to learn more about ASE.
        /// - `current_ruleset`. Current ruleset version and the ISO 8601 date the version was introduced.
        /// - `eval_status`. Specifies whether evaluation mode is enabled or disabled.
        /// - `eval_ruleset`. Evaluation ruleset version and the ISO 8601 date the evaluation began.
        /// - `eval_expiration_date`. ISO 8601 timestamp indicating when evaluation mode expires. Valid only if `eval_status` is set to **enabled**.
        /// - `output_text`. Tabular report of the mode information.
        /// - `json`. JSON-formatted list of the mode information.
        /// </summary>
        public static Task<GetAppSecWafModeResult> InvokeAsync(GetAppSecWafModeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecWafModeResult>("akamai:index/getAppSecWafMode:getAppSecWafMode", args ?? new GetAppSecWafModeArgs(), options.WithDefaults());

        /// <summary>
        /// **Scopes**: Security policy
        /// 
        /// Returns information about how the Kona Rule Set rules associated with a security configuration and security policy are updated. The WAF (Web Application Firewall) mode determines whether Kona Rule Sets are automatically updated as part of automated attack groups (`mode = AAG`) or whether you must periodically check for new rules and then manually update those rules yourself (`mode = KRS`).
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/mode](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getmode)
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
        ///             Name = "Documentation",
        ///         }));
        ///         var wafMode = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecWafMode.InvokeAsync(new Akamai.GetAppSecWafModeArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = "gms1_134637",
        ///         })));
        ///         this.WafModeMode = wafMode.Apply(wafMode =&gt; wafMode.Mode);
        ///         this.WafModeCurrentRuleset = wafMode.Apply(wafMode =&gt; wafMode.CurrentRuleset);
        ///         this.WafModeEvalStatus = wafMode.Apply(wafMode =&gt; wafMode.EvalStatus);
        ///         this.WafModeEvalRuleset = wafMode.Apply(wafMode =&gt; wafMode.EvalRuleset);
        ///         this.WafModeEvalExpirationDate = wafMode.Apply(wafMode =&gt; wafMode.EvalExpirationDate);
        ///         this.WafModeText = wafMode.Apply(wafMode =&gt; wafMode.OutputText);
        ///         this.WafModeJson = wafMode.Apply(wafMode =&gt; wafMode.Json);
        ///     }
        /// 
        ///     [Output("wafModeMode")]
        ///     public Output&lt;string&gt; WafModeMode { get; set; }
        ///     [Output("wafModeCurrentRuleset")]
        ///     public Output&lt;string&gt; WafModeCurrentRuleset { get; set; }
        ///     [Output("wafModeEvalStatus")]
        ///     public Output&lt;string&gt; WafModeEvalStatus { get; set; }
        ///     [Output("wafModeEvalRuleset")]
        ///     public Output&lt;string&gt; WafModeEvalRuleset { get; set; }
        ///     [Output("wafModeEvalExpirationDate")]
        ///     public Output&lt;string&gt; WafModeEvalExpirationDate { get; set; }
        ///     [Output("wafModeText")]
        ///     public Output&lt;string&gt; WafModeText { get; set; }
        ///     [Output("wafModeJson")]
        ///     public Output&lt;string&gt; WafModeJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `mode`. Security policy mode, either **KRS** (update manually) or **AAG** (update automatically), For organizations running the Adaptive Security Engine (ASE) beta, you'll get back **ASE_AUTO** for automatic updates or **ASE_MANUAL** for manual updates. Please contact your Akamai representative to learn more about ASE.
        /// - `current_ruleset`. Current ruleset version and the ISO 8601 date the version was introduced.
        /// - `eval_status`. Specifies whether evaluation mode is enabled or disabled.
        /// - `eval_ruleset`. Evaluation ruleset version and the ISO 8601 date the evaluation began.
        /// - `eval_expiration_date`. ISO 8601 timestamp indicating when evaluation mode expires. Valid only if `eval_status` is set to **enabled**.
        /// - `output_text`. Tabular report of the mode information.
        /// - `json`. JSON-formatted list of the mode information.
        /// </summary>
        public static Output<GetAppSecWafModeResult> Invoke(GetAppSecWafModeInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecWafModeResult>("akamai:index/getAppSecWafMode:getAppSecWafMode", args ?? new GetAppSecWafModeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppSecWafModeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the Kona Rule Set rules.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// . Unique identifier of the security policy associated with the Kona Rule Set rules.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public string SecurityPolicyId { get; set; } = null!;

        public GetAppSecWafModeArgs()
        {
        }
    }

    public sealed class GetAppSecWafModeInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the Kona Rule Set rules.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Unique identifier of the security policy associated with the Kona Rule Set rules.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public Input<string> SecurityPolicyId { get; set; } = null!;

        public GetAppSecWafModeInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecWafModeResult
    {
        public readonly int ConfigId;
        public readonly string CurrentRuleset;
        public readonly string EvalExpirationDate;
        public readonly string EvalRuleset;
        public readonly string EvalStatus;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Json;
        public readonly string Mode;
        public readonly string OutputText;
        public readonly string SecurityPolicyId;

        [OutputConstructor]
        private GetAppSecWafModeResult(
            int configId,

            string currentRuleset,

            string evalExpirationDate,

            string evalRuleset,

            string evalStatus,

            string id,

            string json,

            string mode,

            string outputText,

            string securityPolicyId)
        {
            ConfigId = configId;
            CurrentRuleset = currentRuleset;
            EvalExpirationDate = evalExpirationDate;
            EvalRuleset = evalRuleset;
            EvalStatus = evalStatus;
            Id = id;
            Json = json;
            Mode = mode;
            OutputText = outputText;
            SecurityPolicyId = securityPolicyId;
        }
    }
}
