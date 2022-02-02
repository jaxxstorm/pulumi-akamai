// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * @deprecated akamai.edgedns.DnsZone has been deprecated in favor of akamai.DnsZone
 */
export class DnsZone extends pulumi.CustomResource {
    /**
     * Get an existing DnsZone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DnsZoneState, opts?: pulumi.CustomResourceOptions): DnsZone {
        pulumi.log.warn("DnsZone is deprecated: akamai.edgedns.DnsZone has been deprecated in favor of akamai.DnsZone")
        return new DnsZone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:edgedns/dnsZone:DnsZone';

    /**
     * Returns true if the given object is an instance of DnsZone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DnsZone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DnsZone.__pulumiType;
    }

    public /*out*/ readonly activationState!: pulumi.Output<string>;
    public /*out*/ readonly aliasCount!: pulumi.Output<number>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly contract!: pulumi.Output<string>;
    public readonly endCustomerId!: pulumi.Output<string | undefined>;
    public readonly group!: pulumi.Output<string | undefined>;
    public readonly masters!: pulumi.Output<string[] | undefined>;
    public readonly signAndServe!: pulumi.Output<boolean | undefined>;
    public readonly signAndServeAlgorithm!: pulumi.Output<string | undefined>;
    public readonly target!: pulumi.Output<string | undefined>;
    public readonly tsigKey!: pulumi.Output<outputs.edgedns.DnsZoneTsigKey | undefined>;
    public readonly type!: pulumi.Output<string>;
    public /*out*/ readonly versionId!: pulumi.Output<string>;
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a DnsZone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated akamai.edgedns.DnsZone has been deprecated in favor of akamai.DnsZone */
    constructor(name: string, args: DnsZoneArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated akamai.edgedns.DnsZone has been deprecated in favor of akamai.DnsZone */
    constructor(name: string, argsOrState?: DnsZoneArgs | DnsZoneState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DnsZone is deprecated: akamai.edgedns.DnsZone has been deprecated in favor of akamai.DnsZone")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DnsZoneState | undefined;
            resourceInputs["activationState"] = state ? state.activationState : undefined;
            resourceInputs["aliasCount"] = state ? state.aliasCount : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["contract"] = state ? state.contract : undefined;
            resourceInputs["endCustomerId"] = state ? state.endCustomerId : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["masters"] = state ? state.masters : undefined;
            resourceInputs["signAndServe"] = state ? state.signAndServe : undefined;
            resourceInputs["signAndServeAlgorithm"] = state ? state.signAndServeAlgorithm : undefined;
            resourceInputs["target"] = state ? state.target : undefined;
            resourceInputs["tsigKey"] = state ? state.tsigKey : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["versionId"] = state ? state.versionId : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as DnsZoneArgs | undefined;
            if ((!args || args.contract === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contract'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["contract"] = args ? args.contract : undefined;
            resourceInputs["endCustomerId"] = args ? args.endCustomerId : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["masters"] = args ? args.masters : undefined;
            resourceInputs["signAndServe"] = args ? args.signAndServe : undefined;
            resourceInputs["signAndServeAlgorithm"] = args ? args.signAndServeAlgorithm : undefined;
            resourceInputs["target"] = args ? args.target : undefined;
            resourceInputs["tsigKey"] = args ? args.tsigKey : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["activationState"] = undefined /*out*/;
            resourceInputs["aliasCount"] = undefined /*out*/;
            resourceInputs["versionId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DnsZone.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DnsZone resources.
 */
export interface DnsZoneState {
    activationState?: pulumi.Input<string>;
    aliasCount?: pulumi.Input<number>;
    comment?: pulumi.Input<string>;
    contract?: pulumi.Input<string>;
    endCustomerId?: pulumi.Input<string>;
    group?: pulumi.Input<string>;
    masters?: pulumi.Input<pulumi.Input<string>[]>;
    signAndServe?: pulumi.Input<boolean>;
    signAndServeAlgorithm?: pulumi.Input<string>;
    target?: pulumi.Input<string>;
    tsigKey?: pulumi.Input<inputs.edgedns.DnsZoneTsigKey>;
    type?: pulumi.Input<string>;
    versionId?: pulumi.Input<string>;
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DnsZone resource.
 */
export interface DnsZoneArgs {
    comment?: pulumi.Input<string>;
    contract: pulumi.Input<string>;
    endCustomerId?: pulumi.Input<string>;
    group?: pulumi.Input<string>;
    masters?: pulumi.Input<pulumi.Input<string>[]>;
    signAndServe?: pulumi.Input<boolean>;
    signAndServeAlgorithm?: pulumi.Input<string>;
    target?: pulumi.Input<string>;
    tsigKey?: pulumi.Input<inputs.edgedns.DnsZoneTsigKey>;
    type: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}
