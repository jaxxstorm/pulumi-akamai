// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Scopes**: Security configuration; security policy
 *
 * Returns information about the API endpoints associated with a security policy or configuration. The returned information is described in the [Endpoint members](https://developer.akamai.com/api/cloud_security/application_security/v1.html#apiendpoint) section of the Application Security API documentation.
 *
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/api-endpoints](https://developer.akamai.com/api/cloud_security/application_security/v1.html#getapiendpoints)
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const apiEndpoints = pulumi.output(akamai.getAppSecApiEndpoints({
 *     apiName: "Contracts",
 *     configId: 58843,
 * }));
 * ```
 * ## Output Options
 *
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 *
 * - `idList`. List of API endpoint IDs.
 * - `json`. JSON-formatted list of information about the API endpoints.
 * - `outputText`. Tabular report showing the ID and name of the API endpoints.
 */
export function getAppSecApiEndpoints(args: GetAppSecApiEndpointsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppSecApiEndpointsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("akamai:index/getAppSecApiEndpoints:getAppSecApiEndpoints", {
        "apiName": args.apiName,
        "configId": args.configId,
        "securityPolicyId": args.securityPolicyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppSecApiEndpoints.
 */
export interface GetAppSecApiEndpointsArgs {
    /**
     * . Name of the API endpoint you want to return information for. If not included, information is returned for all your API endpoints.
     */
    apiName?: string;
    /**
     * . Unique identifier of the security configuration associated with the API endpoints.
     */
    configId: number;
    /**
     * . Unique identifier of the security policy associated with the API endpoints. If not included, information is returned for all your security policies.
     */
    securityPolicyId?: string;
}

/**
 * A collection of values returned by getAppSecApiEndpoints.
 */
export interface GetAppSecApiEndpointsResult {
    readonly apiName?: string;
    readonly configId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly idLists: number[];
    readonly json: string;
    readonly outputText: string;
    readonly securityPolicyId?: string;
}

export function getAppSecApiEndpointsOutput(args: GetAppSecApiEndpointsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppSecApiEndpointsResult> {
    return pulumi.output(args).apply(a => getAppSecApiEndpoints(a, opts))
}

/**
 * A collection of arguments for invoking getAppSecApiEndpoints.
 */
export interface GetAppSecApiEndpointsOutputArgs {
    /**
     * . Name of the API endpoint you want to return information for. If not included, information is returned for all your API endpoints.
     */
    apiName?: pulumi.Input<string>;
    /**
     * . Unique identifier of the security configuration associated with the API endpoints.
     */
    configId: pulumi.Input<number>;
    /**
     * . Unique identifier of the security policy associated with the API endpoints. If not included, information is returned for all your security policies.
     */
    securityPolicyId?: pulumi.Input<string>;
}
