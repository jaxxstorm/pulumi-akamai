// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecAdvancedSettingsLogging
    {
        /// <summary>
        /// Use the `akamai.AppSecAdvancedSettingsLogging` data source to retrieve information about the HTTP header logging controls for a configuration. This operation applies at the configuration level, and therefore applies to all policies within a configuration. You may retrieve these settings for a particular policy by specifying the policy using the security_policy_id parameter. The information available is described [here](https://developer.akamai.com/api/cloud_security/application_security/v1.html#gethttpheaderloggingforaconfiguration).
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
        ///         var logging = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecAdvancedSettingsLogging.InvokeAsync(new Akamai.GetAppSecAdvancedSettingsLoggingArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.AdvancedSettingsLoggingOutput = logging.Apply(logging =&gt; logging.OutputText);
        ///         this.AdvancedSettingsLoggingJson = logging.Apply(logging =&gt; logging.Json);
        ///         var policyOverride = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecAdvancedSettingsLogging.InvokeAsync(new Akamai.GetAppSecAdvancedSettingsLoggingArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = @var.Security_policy_id,
        ///         })));
        ///         this.AdvancedSettingsPolicyLoggingOutput = policyOverride.Apply(policyOverride =&gt; policyOverride.OutputText);
        ///         this.AdvancedSettingsPolicyLoggingJson = policyOverride.Apply(policyOverride =&gt; policyOverride.Json);
        ///     }
        /// 
        ///     [Output("advancedSettingsLoggingOutput")]
        ///     public Output&lt;string&gt; AdvancedSettingsLoggingOutput { get; set; }
        ///     [Output("advancedSettingsLoggingJson")]
        ///     public Output&lt;string&gt; AdvancedSettingsLoggingJson { get; set; }
        ///     [Output("advancedSettingsPolicyLoggingOutput")]
        ///     public Output&lt;string&gt; AdvancedSettingsPolicyLoggingOutput { get; set; }
        ///     [Output("advancedSettingsPolicyLoggingJson")]
        ///     public Output&lt;string&gt; AdvancedSettingsPolicyLoggingJson { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecAdvancedSettingsLoggingResult> InvokeAsync(GetAppSecAdvancedSettingsLoggingArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecAdvancedSettingsLoggingResult>("akamai:index/getAppSecAdvancedSettingsLogging:getAppSecAdvancedSettingsLogging", args ?? new GetAppSecAdvancedSettingsLoggingArgs(), options.WithVersion());
    }


    public sealed class GetAppSecAdvancedSettingsLoggingArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The configuration ID.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// The ID of the security policy to use.
        /// </summary>
        [Input("securityPolicyId")]
        public string? SecurityPolicyId { get; set; }

        public GetAppSecAdvancedSettingsLoggingArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecAdvancedSettingsLoggingResult
    {
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A JSON-formatted list of information about the logging settings.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// A tabular display showing the logging settings.
        /// </summary>
        public readonly string OutputText;
        public readonly string? SecurityPolicyId;

        [OutputConstructor]
        private GetAppSecAdvancedSettingsLoggingResult(
            int configId,

            string id,

            string json,

            string outputText,

            string? securityPolicyId)
        {
            ConfigId = configId;
            Id = id;
            Json = json;
            OutputText = outputText;
            SecurityPolicyId = securityPolicyId;
        }
    }
}
