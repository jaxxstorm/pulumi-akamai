// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface DnsZoneTsigKey {
    algorithm: string;
    /**
     * key name
     * * `algorithm`
     * * `secret`
     */
    name: string;
    secret: string;
}

export interface GetContractsContract {
    contractId: string;
    contractTypeName: string;
}

export interface GetGroupsGroup {
    contractIds: string[];
    groupId: string;
    groupName: string;
    parentGroupId: string;
}

export interface GetPropertiesProperty {
    contractId: string;
    groupId: string;
    latestVersion: number;
    note: string;
    productId: string;
    productionVersion: number;
    propertyId: string;
    propertyName: string;
    ruleFormat: string;
    stagingVersion: number;
}

export interface GetPropertyProductsProduct {
    productId: string;
    productName: string;
}

export interface GetPropertyRulesTemplateVariable {
    name: string;
    type?: string;
    value: string;
}

export interface GtmAsmapAssignment {
    asNumbers: number[];
    datacenterId: number;
    nickname: string;
}

export interface GtmAsmapDefaultDatacenter {
    datacenterId: number;
    nickname?: string;
}

export interface GtmCidrmapAssignment {
    blocks?: string[];
    datacenterId: number;
    nickname: string;
}

export interface GtmCidrmapDefaultDatacenter {
    datacenterId: number;
    nickname?: string;
}

export interface GtmDatacenterDefaultLoadObject {
    loadObject?: string;
    loadObjectPort?: number;
    /**
     * * `city`
     * * `cloneOf`
     */
    loadServers?: string[];
}

export interface GtmGeomapAssignment {
    countries?: string[];
    datacenterId: number;
    nickname: string;
}

export interface GtmGeomapDefaultDatacenter {
    datacenterId: number;
    nickname?: string;
}

export interface GtmPropertyLivenessTest {
    answersRequired?: boolean;
    /**
     * * `errorPenalty`
     */
    disableNonstandardPortWarning?: boolean;
    disabled?: boolean;
    errorPenalty?: number;
    httpError3xx?: boolean;
    httpError4xx?: boolean;
    httpError5xx?: boolean;
    /**
     * `name`
     * `value`
     */
    httpHeaders?: outputs.GtmPropertyLivenessTestHttpHeader[];
    /**
     * Liveness test name
     * * `testInterval`
     * * `testObjectProtocol`
     * * `testTimeout`
     */
    name: string;
    peerCertificateVerification?: boolean;
    /**
     * * `requestString`
     * * `resourceType`
     * * `responseString`
     * * `sslClientCertificate`
     * * `sslClientPrivateKey`
     * * `testObject`
     * * `testObjectPassword`
     * * `testObjectPort`
     * * `testObjectUsername`
     * * `timeoutPenalty`
     */
    recursionRequested?: boolean;
    requestString?: string;
    resourceType?: string;
    responseString?: string;
    sslClientCertificate?: string;
    sslClientPrivateKey?: string;
    testInterval: number;
    testObject: string;
    testObjectPassword?: string;
    testObjectPort?: number;
    testObjectProtocol: string;
    testObjectUsername?: string;
    testTimeout: number;
    timeoutPenalty?: number;
}

export interface GtmPropertyLivenessTestHttpHeader {
    /**
     * Liveness test name
     * * `testInterval`
     * * `testObjectProtocol`
     * * `testTimeout`
     */
    name?: string;
    value?: string;
}

export interface GtmPropertyStaticRrSet {
    rdatas?: string[];
    ttl?: number;
    /**
     * Property type  
     * * `scoreAggregationType`
     */
    type?: string;
}

export interface GtmPropertyTrafficTarget {
    datacenterId?: number;
    /**
     * * `weight`
     */
    enabled?: boolean;
    handoutCname?: string;
    /**
     * Liveness test name
     * * `testInterval`
     * * `testObjectProtocol`
     * * `testTimeout`
     */
    name?: string;
    servers?: string[];
    weight?: number;
}

export interface GtmResourceResourceInstance {
    datacenterId: number;
    loadObject?: string;
    loadObjectPort?: number;
    loadServers?: string[];
    /**
     * * `hostHeader`
     * * `leastSquaresDecay`
     * * `upperBound`
     * * `description`
     * * `leaderString`
     * * `constrainedProperty`
     * * `loadImbalancePercent`
     * * `maxUMultiplicativeIncrement`
     * * `decayRate`
     */
    useDefaultLoadObject?: boolean;
}

export interface PropertyOrigin {
    cacheKeyHostname?: string;
    compress?: boolean;
    enableTrueClientIp?: boolean;
    forwardHostname?: string;
    hostname?: string;
    port?: number;
}

export interface PropertyRuleError {
    behaviorName?: string;
    detail?: string;
    errorLocation?: string;
    instance?: string;
    statusCode?: number;
    title?: string;
    type?: string;
}

export interface PropertyRuleWarning {
    behaviorName?: string;
    detail?: string;
    errorLocation?: string;
    instance?: string;
    statusCode?: number;
    title?: string;
    type?: string;
}

export interface PropertyVariablesVariable {
    variables?: outputs.PropertyVariablesVariableVariable[];
}

export interface PropertyVariablesVariableVariable {
    description?: string;
    hidden: boolean;
    name: string;
    sensitive: boolean;
    value?: string;
}

export interface ProviderAppsec {
    accessToken?: string;
    accountKey?: string;
    clientSecret?: string;
    clientToken?: string;
    host?: string;
    maxBody?: number;
}

export interface ProviderConfig {
    accessToken?: string;
    accountKey?: string;
    clientSecret?: string;
    clientToken?: string;
    host?: string;
    maxBody?: number;
}

export interface ProviderDns {
    accessToken?: string;
    accountKey?: string;
    clientSecret?: string;
    clientToken?: string;
    host?: string;
    maxBody?: number;
}

export interface ProviderGtm {
    accessToken?: string;
    accountKey?: string;
    clientSecret?: string;
    clientToken?: string;
    host?: string;
    maxBody?: number;
}

export interface ProviderProperty {
    accessToken?: string;
    accountKey?: string;
    clientSecret?: string;
    clientToken?: string;
    host?: string;
    maxBody?: number;
}
export namespace config {
    export interface Appsecs {
        accessToken?: string;
        accountKey?: string;
        clientSecret?: string;
        clientToken?: string;
        host?: string;
        maxBody?: number;
    }

    export interface Config {
        accessToken?: string;
        accountKey?: string;
        clientSecret?: string;
        clientToken?: string;
        host?: string;
        maxBody?: number;
    }

    export interface Dns {
        accessToken?: string;
        accountKey?: string;
        clientSecret?: string;
        clientToken?: string;
        host?: string;
        maxBody?: number;
    }

    export interface Gtm {
        accessToken?: string;
        accountKey?: string;
        clientSecret?: string;
        clientToken?: string;
        host?: string;
        maxBody?: number;
    }

    export interface Property {
        accessToken?: string;
        accountKey?: string;
        clientSecret?: string;
        clientToken?: string;
        host?: string;
        maxBody?: number;
    }
}

export namespace edgedns {
    export interface DnsZoneTsigKey {
        algorithm: string;
        name: string;
        secret: string;
    }
}

export namespace properties {
    export interface PropertyOrigin {
        cacheKeyHostname?: string;
        compress?: boolean;
        enableTrueClientIp?: boolean;
        forwardHostname?: string;
        hostname?: string;
        port?: number;
    }

    export interface PropertyRuleError {
        behaviorName?: string;
        detail?: string;
        errorLocation?: string;
        instance?: string;
        statusCode?: number;
        title?: string;
        type?: string;
    }

    export interface PropertyRuleWarning {
        behaviorName?: string;
        detail?: string;
        errorLocation?: string;
        instance?: string;
        statusCode?: number;
        title?: string;
        type?: string;
    }

    export interface PropertyVariablesVariable {
        variables?: outputs.properties.PropertyVariablesVariableVariable[];
    }

    export interface PropertyVariablesVariableVariable {
        description?: string;
        hidden: boolean;
        name: string;
        sensitive: boolean;
        value?: string;
    }
}

export namespace trafficmanagement {
    export interface GtmASmapAssignment {
        asNumbers: number[];
        datacenterId: number;
        nickname: string;
    }

    export interface GtmASmapDefaultDatacenter {
        datacenterId: number;
        nickname?: string;
    }

    export interface GtmCidrmapAssignment {
        blocks?: string[];
        datacenterId: number;
        nickname: string;
    }

    export interface GtmCidrmapDefaultDatacenter {
        datacenterId: number;
        nickname?: string;
    }

    export interface GtmDatacenterDefaultLoadObject {
        loadObject?: string;
        loadObjectPort?: number;
        loadServers?: string[];
    }

    export interface GtmGeomapAssignment {
        countries?: string[];
        datacenterId: number;
        nickname: string;
    }

    export interface GtmGeomapDefaultDatacenter {
        datacenterId: number;
        nickname?: string;
    }

    export interface GtmPropertyLivenessTest {
        answersRequired?: boolean;
        disableNonstandardPortWarning?: boolean;
        disabled?: boolean;
        errorPenalty?: number;
        httpError3xx?: boolean;
        httpError4xx?: boolean;
        httpError5xx?: boolean;
        httpHeaders?: outputs.trafficmanagement.GtmPropertyLivenessTestHttpHeader[];
        name: string;
        peerCertificateVerification?: boolean;
        recursionRequested?: boolean;
        requestString?: string;
        resourceType?: string;
        responseString?: string;
        sslClientCertificate?: string;
        sslClientPrivateKey?: string;
        testInterval: number;
        testObject: string;
        testObjectPassword?: string;
        testObjectPort?: number;
        testObjectProtocol: string;
        testObjectUsername?: string;
        testTimeout: number;
        timeoutPenalty?: number;
    }

    export interface GtmPropertyLivenessTestHttpHeader {
        name?: string;
        value?: string;
    }

    export interface GtmPropertyStaticRrSet {
        rdatas?: string[];
        ttl?: number;
        type?: string;
    }

    export interface GtmPropertyTrafficTarget {
        datacenterId?: number;
        enabled?: boolean;
        handoutCname?: string;
        name?: string;
        servers?: string[];
        weight?: number;
    }

    export interface GtmResourceResourceInstance {
        datacenterId: number;
        loadObject?: string;
        loadObjectPort?: number;
        loadServers?: string[];
        useDefaultLoadObject?: boolean;
    }
}
