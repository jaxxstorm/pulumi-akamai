// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security policy; API endpoint
 *
 * Returns information about API endpoint constraints and actions. The returned information is described in the [API Constraints members](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getapirequestconstraints) section of the Application Security API.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/api-request-constraints](https://techdocs.akamai.com/application-security/reference/get-api-request-constraints)
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const configuration = akamai.getAppSecConfiguration({
 *     name: "Documentation",
 * });
 * const apisRequestConstraints = configuration.then(configuration => akamai.getAppSecApiRequestConstraints({
 *     configId: configuration.configId,
 *     securityPolicyId: "gms1_134637",
 * }));
 * export const apisConstraintsText = apisRequestConstraints.then(apisRequestConstraints => apisRequestConstraints.outputText);
 * export const apisConstraintsJson = apisRequestConstraints.then(apisRequestConstraints => apisRequestConstraints.json);
 * const apiRequestConstraints = configuration.then(configuration => akamai.getAppSecApiRequestConstraints({
 *     configId: configuration.configId,
 *     securityPolicyId: "gms1_134637",
 *     apiId: 624913,
 * }));
 * export const apiConstraintsText = apiRequestConstraints.then(apiRequestConstraints => apiRequestConstraints.outputText);
 * export const apiConstraintsJson = apiRequestConstraints.then(apiRequestConstraints => apiRequestConstraints.json);
 * ```
 * ## Output Options
 *
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 *
 * - `json`. JSON-formatted list of information about the APIs, their constraints, and their actions.
 * - `outputText`. Tabular report of the APIs, their constraints, and their actions.
 */
export function getAppSecApiRequestConstraints(args: GetAppSecApiRequestConstraintsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecApiRequestConstraintsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("akamai:index/getAppSecApiRequestConstraints:getAppSecApiRequestConstraints", {
        "apiId": args.apiId,
        "configId": args.configId,
        "securityPolicyId": args.securityPolicyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecApiRequestConstraints.
 */
export interface GetAppSecApiRequestConstraintsArgs {
    /**
     * . Unique identifier of the API endpoint you want to return constraint information for.
     */
    apiId?: number;
    /**
     * . Unique identifier of the security configuration associated with the API constraints.
     */
    configId: number;
    /**
     * . Unique identifier of the security policy associated with the API constraints.
     */
    securityPolicyId: string;
}

/**
 * A collection of values returned by getAppSecApiRequestConstraints.
 */
export interface GetAppSecApiRequestConstraintsResult {
    readonly apiId?: number;
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly json: string;
    readonly outputText: string;
    readonly securityPolicyId: string;
}

export function getAppSecApiRequestConstraintsOutput(args: GetAppSecApiRequestConstraintsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecApiRequestConstraintsResult> {
    return pulumi.output(args).apply(a => getAppSecApiRequestConstraints(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecApiRequestConstraints.
 */
export interface GetAppSecApiRequestConstraintsOutputArgs {
    /**
     * . Unique identifier of the API endpoint you want to return constraint information for.
     */
    apiId?: pulumi.Input<number>;
    /**
     * . Unique identifier of the security configuration associated with the API constraints.
     */
    configId: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the API constraints.
     */
    securityPolicyId: pulumi.Input<string>;
}
