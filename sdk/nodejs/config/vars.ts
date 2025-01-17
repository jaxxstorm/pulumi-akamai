// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("akamai");

export declare const appsecSection: string | undefined;
Object.defineProperty(exports, "appsecSection", {
    get() {
        return __config.get("appsecSection");
    },
    enumerable: true,
});

export declare const appsecs: outputs.config.Appsecs[] | undefined;
Object.defineProperty(exports, "appsecs", {
    get() {
        return __config.getObject<outputs.config.Appsecs[]>("appsecs");
    },
    enumerable: true,
});

export declare const cacheEnabled: boolean | undefined;
Object.defineProperty(exports, "cacheEnabled", {
    get() {
        return __config.getObject<boolean>("cacheEnabled");
    },
    enumerable: true,
});

export declare const config: outputs.config.Config | undefined;
Object.defineProperty(exports, "config", {
    get() {
        return __config.getObject<outputs.config.Config>("config");
    },
    enumerable: true,
});

/**
 * The section of the edgerc file to use for configuration
 */
export declare const configSection: string | undefined;
Object.defineProperty(exports, "configSection", {
    get() {
        return __config.get("configSection");
    },
    enumerable: true,
});

export declare const dns: outputs.config.Dns | undefined;
Object.defineProperty(exports, "dns", {
    get() {
        return __config.getObject<outputs.config.Dns>("dns");
    },
    enumerable: true,
});

export declare const dnsSection: string | undefined;
Object.defineProperty(exports, "dnsSection", {
    get() {
        return __config.get("dnsSection");
    },
    enumerable: true,
});

export declare const edgerc: string | undefined;
Object.defineProperty(exports, "edgerc", {
    get() {
        return __config.get("edgerc");
    },
    enumerable: true,
});

export declare const gtm: outputs.config.Gtm | undefined;
Object.defineProperty(exports, "gtm", {
    get() {
        return __config.getObject<outputs.config.Gtm>("gtm");
    },
    enumerable: true,
});

export declare const gtmSection: string | undefined;
Object.defineProperty(exports, "gtmSection", {
    get() {
        return __config.get("gtmSection");
    },
    enumerable: true,
});

export declare const networklistSection: string | undefined;
Object.defineProperty(exports, "networklistSection", {
    get() {
        return __config.get("networklistSection");
    },
    enumerable: true,
});

export declare const networks: outputs.config.Networks[] | undefined;
Object.defineProperty(exports, "networks", {
    get() {
        return __config.getObject<outputs.config.Networks[]>("networks");
    },
    enumerable: true,
});

export declare const papiSection: string | undefined;
Object.defineProperty(exports, "papiSection", {
    get() {
        return __config.get("papiSection");
    },
    enumerable: true,
});

export declare const property: outputs.config.Property | undefined;
Object.defineProperty(exports, "property", {
    get() {
        return __config.getObject<outputs.config.Property>("property");
    },
    enumerable: true,
});

export declare const propertySection: string | undefined;
Object.defineProperty(exports, "propertySection", {
    get() {
        return __config.get("propertySection");
    },
    enumerable: true,
});

