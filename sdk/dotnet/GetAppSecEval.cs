// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecEval
    {
        public static Task<GetAppSecEvalResult> InvokeAsync(GetAppSecEvalArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecEvalResult>("akamai:index/getAppSecEval:getAppSecEval", args ?? new GetAppSecEvalArgs(), options.WithDefaults());

        public static Output<GetAppSecEvalResult> Invoke(GetAppSecEvalInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecEvalResult>("akamai:index/getAppSecEval:getAppSecEval", args ?? new GetAppSecEvalInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppSecEvalArgs : Pulumi.InvokeArgs
    {
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        [Input("securityPolicyId", required: true)]
        public string SecurityPolicyId { get; set; } = null!;

        public GetAppSecEvalArgs()
        {
        }
    }

    public sealed class GetAppSecEvalInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        [Input("securityPolicyId", required: true)]
        public Input<string> SecurityPolicyId { get; set; } = null!;

        public GetAppSecEvalInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecEvalResult
    {
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OutputText;
        public readonly string SecurityPolicyId;

        [OutputConstructor]
        private GetAppSecEvalResult(
            int configId,

            string id,

            string outputText,

            string securityPolicyId)
        {
            ConfigId = configId;
            Id = id;
            OutputText = outputText;
            SecurityPolicyId = securityPolicyId;
        }
    }
}
