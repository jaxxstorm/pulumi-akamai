// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Once you complete the Let's Encrypt challenges, optionally use the `akamai.CpsDvValidation` resource to send the acknowledgement to CPS and inform it that tokens are ready for validation. You can also wait for CPS to check for the tokens, which it does on a regular schedule. Next, CPS automatically deploys the certificate on Staging, and eventually on the Production network.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const example = new akamai.CpsDvValidation("example", {
 *     enrollmentId: akamai_cps_dv_enrollment.example.id,
 *     sans: akamai_cps_dv_enrollment.example.sans,
 * });
 * ```
 * ## Attributes reference
 *
 * * `status` - The status of certificate validation.
 */
export class CpsDvValidation extends pulumi.CustomResource {
    /**
     * Get an existing CpsDvValidation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CpsDvValidationState, opts?: pulumi.CustomResourceOptions): CpsDvValidation {
        return new CpsDvValidation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/cpsDvValidation:CpsDvValidation';

    /**
     * Returns true if the given object is an instance of CpsDvValidation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CpsDvValidation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CpsDvValidation.__pulumiType;
    }

    /**
     * Unique identifier for the DV certificate enrollment.
     */
    public readonly enrollmentId!: pulumi.Output<number>;
    /**
     * The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
     */
    public readonly sans!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a CpsDvValidation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CpsDvValidationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CpsDvValidationArgs | CpsDvValidationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CpsDvValidationState | undefined;
            resourceInputs["enrollmentId"] = state ? state.enrollmentId : undefined;
            resourceInputs["sans"] = state ? state.sans : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as CpsDvValidationArgs | undefined;
            if ((!args || args.enrollmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enrollmentId'");
            }
            resourceInputs["enrollmentId"] = args ? args.enrollmentId : undefined;
            resourceInputs["sans"] = args ? args.sans : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CpsDvValidation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CpsDvValidation resources.
 */
export interface CpsDvValidationState {
    /**
     * Unique identifier for the DV certificate enrollment.
     */
    enrollmentId?: pulumi.Input<number>;
    /**
     * The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
     */
    sans?: pulumi.Input<pulumi.Input<string>[]>;
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CpsDvValidation resource.
 */
export interface CpsDvValidationArgs {
    /**
     * Unique identifier for the DV certificate enrollment.
     */
    enrollmentId: pulumi.Input<number>;
    /**
     * The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
     */
    sans?: pulumi.Input<pulumi.Input<string>[]>;
}
