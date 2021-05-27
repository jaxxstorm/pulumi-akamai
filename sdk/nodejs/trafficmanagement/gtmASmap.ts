// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * @deprecated akamai.trafficmanagement.GtmASmap has been deprecated in favor of akamai.GtmAsmap
 */
export class GtmASmap extends pulumi.CustomResource {
    /**
     * Get an existing GtmASmap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GtmASmapState, opts?: pulumi.CustomResourceOptions): GtmASmap {
        pulumi.log.warn("GtmASmap is deprecated: akamai.trafficmanagement.GtmASmap has been deprecated in favor of akamai.GtmAsmap")
        return new GtmASmap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:trafficmanagement/gtmASmap:GtmASmap';

    /**
     * Returns true if the given object is an instance of GtmASmap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GtmASmap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GtmASmap.__pulumiType;
    }

    public readonly assignments!: pulumi.Output<outputs.trafficmanagement.GtmASmapAssignment[] | undefined>;
    public readonly defaultDatacenter!: pulumi.Output<outputs.trafficmanagement.GtmASmapDefaultDatacenter>;
    public readonly domain!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly waitOnComplete!: pulumi.Output<boolean | undefined>;

    /**
     * Create a GtmASmap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated akamai.trafficmanagement.GtmASmap has been deprecated in favor of akamai.GtmAsmap */
    constructor(name: string, args: GtmASmapArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated akamai.trafficmanagement.GtmASmap has been deprecated in favor of akamai.GtmAsmap */
    constructor(name: string, argsOrState?: GtmASmapArgs | GtmASmapState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("GtmASmap is deprecated: akamai.trafficmanagement.GtmASmap has been deprecated in favor of akamai.GtmAsmap")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GtmASmapState | undefined;
            inputs["assignments"] = state ? state.assignments : undefined;
            inputs["defaultDatacenter"] = state ? state.defaultDatacenter : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["waitOnComplete"] = state ? state.waitOnComplete : undefined;
        } else {
            const args = argsOrState as GtmASmapArgs | undefined;
            if ((!args || args.defaultDatacenter === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultDatacenter'");
            }
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            inputs["assignments"] = args ? args.assignments : undefined;
            inputs["defaultDatacenter"] = args ? args.defaultDatacenter : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["waitOnComplete"] = args ? args.waitOnComplete : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GtmASmap.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GtmASmap resources.
 */
export interface GtmASmapState {
    assignments?: pulumi.Input<pulumi.Input<inputs.trafficmanagement.GtmASmapAssignment>[]>;
    defaultDatacenter?: pulumi.Input<inputs.trafficmanagement.GtmASmapDefaultDatacenter>;
    domain?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    waitOnComplete?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GtmASmap resource.
 */
export interface GtmASmapArgs {
    assignments?: pulumi.Input<pulumi.Input<inputs.trafficmanagement.GtmASmapAssignment>[]>;
    defaultDatacenter: pulumi.Input<inputs.trafficmanagement.GtmASmapDefaultDatacenter>;
    domain: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    waitOnComplete?: pulumi.Input<boolean>;
}
