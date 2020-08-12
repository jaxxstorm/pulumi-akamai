// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The provider type for the akamai package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'akamai';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }


    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        {
            inputs["dns"] = pulumi.output(args ? args.dns : undefined).apply(JSON.stringify);
            inputs["dnsSection"] = args ? args.dnsSection : undefined;
            inputs["edgerc"] = args ? args.edgerc : undefined;
            inputs["gtmSection"] = args ? args.gtmSection : undefined;
            inputs["gtms"] = pulumi.output(args ? args.gtms : undefined).apply(JSON.stringify);
            inputs["papiSection"] = args ? args.papiSection : undefined;
            inputs["properties"] = pulumi.output(args ? args.properties : undefined).apply(JSON.stringify);
            inputs["propertySection"] = args ? args.propertySection : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Provider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    readonly dns?: pulumi.Input<pulumi.Input<inputs.ProviderDn>[]>;
    readonly dnsSection?: pulumi.Input<string>;
    readonly edgerc?: pulumi.Input<string>;
    readonly gtmSection?: pulumi.Input<string>;
    readonly gtms?: pulumi.Input<pulumi.Input<inputs.ProviderGtm>[]>;
    readonly papiSection?: pulumi.Input<string>;
    readonly properties?: pulumi.Input<pulumi.Input<inputs.ProviderProperty>[]>;
    readonly propertySection?: pulumi.Input<string>;
}
