// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./appSecActivations";
export * from "./appSecAdvancedSettingsEvasivePathMatch";
export * from "./appSecAdvancedSettingsLogging";
export * from "./appSecAdvancedSettingsPragmaHeader";
export * from "./appSecAdvancedSettingsPrefetch";
export * from "./appSecApiConstraintsProtection";
export * from "./appSecApiRequestConstraints";
export * from "./appSecAttackGroup";
export * from "./appSecByPassNetworkList";
export * from "./appSecConfiguration";
export * from "./appSecConfigurationRename";
export * from "./appSecCustomDeny";
export * from "./appSecCustomRule";
export * from "./appSecCustomRuleAction";
export * from "./appSecEval";
export * from "./appSecEvalGroup";
export * from "./appSecEvalRule";
export * from "./appSecIPGeo";
export * from "./appSecIPGeoProtection";
export * from "./appSecMatchTarget";
export * from "./appSecMatchTargetSequence";
export * from "./appSecPenaltyBox";
export * from "./appSecRatePolicy";
export * from "./appSecRatePolicyAction";
export * from "./appSecRateProtection";
export * from "./appSecReputationProfile";
export * from "./appSecReputationProfileAction";
export * from "./appSecReputationProfileAnalysis";
export * from "./appSecReputationProtection";
export * from "./appSecRule";
export * from "./appSecRuleUpgrade";
export * from "./appSecSecurityPolicy";
export * from "./appSecSecurityPolicyRename";
export * from "./appSecSelectedHostnames";
export * from "./appSecSiemSettings";
export * from "./appSecSlowPost";
export * from "./appSecSlowPostProtection";
export * from "./appSecThreatIntel";
export * from "./appSecVersionNodes";
export * from "./appSecWafMode";
export * from "./appSecWafProtection";
export * from "./appSecWapSelectedHostnames";
export * from "./cloudletsApplicationLoadBalancer";
export * from "./cloudletsApplicationLoadBalancerActivation";
export * from "./cloudletsPolicy";
export * from "./cloudletsPolicyActivation";
export * from "./cpCode";
export * from "./cpsDvEnrollment";
export * from "./cpsDvValidation";
export * from "./datastream";
export * from "./dnsRecord";
export * from "./dnsZone";
export * from "./edgeHostName";
export * from "./edgeKv";
export * from "./edgeWorker";
export * from "./edgeWorkersActivation";
export * from "./getAppSecAdvancedSettingsEvasivePathMatch";
export * from "./getAppSecAdvancedSettingsLogging";
export * from "./getAppSecAdvancedSettingsPragmaHeader";
export * from "./getAppSecAdvancedSettingsPrefetch";
export * from "./getAppSecApiEndpoints";
export * from "./getAppSecApiRequestConstraints";
export * from "./getAppSecAttackGroups";
export * from "./getAppSecBypassNetworkLists";
export * from "./getAppSecConfiguration";
export * from "./getAppSecConfigurationVersion";
export * from "./getAppSecContractsGroups";
export * from "./getAppSecCustomDeny";
export * from "./getAppSecCustomRuleActions";
export * from "./getAppSecCustomRules";
export * from "./getAppSecEval";
export * from "./getAppSecEvalGroups";
export * from "./getAppSecEvalRules";
export * from "./getAppSecExportConfiguration";
export * from "./getAppSecFailoverHostnames";
export * from "./getAppSecHostnameCoverage";
export * from "./getAppSecHostnameCoverageMatchTargets";
export * from "./getAppSecHostnameCoverageOverlapping";
export * from "./getAppSecIPGeo";
export * from "./getAppSecMatchTargets";
export * from "./getAppSecPenaltyBox";
export * from "./getAppSecRatePolicies";
export * from "./getAppSecRatePolicyActions";
export * from "./getAppSecReputationProfileActions";
export * from "./getAppSecReputationProfileAnalysis";
export * from "./getAppSecReputationProfiles";
export * from "./getAppSecRuleUpgradeDetails";
export * from "./getAppSecRules";
export * from "./getAppSecSecurityPolicy";
export * from "./getAppSecSecurityPolicyProtections";
export * from "./getAppSecSelectableHostnames";
export * from "./getAppSecSelectedHostnames";
export * from "./getAppSecSiemDefinitions";
export * from "./getAppSecSiemSettings";
export * from "./getAppSecSlowPost";
export * from "./getAppSecThreatIntel";
export * from "./getAppSecTuningRecommendations";
export * from "./getAppSecVersionNotes";
export * from "./getAppSecWafMode";
export * from "./getAppSecWapSelectedHostnames";
export * from "./getAuthoritiesSet";
export * from "./getCloudletsApiPrioritizationMatchRule";
export * from "./getCloudletsApplicationLoadBalancer";
export * from "./getCloudletsApplicationLoadBalancerMatchRule";
export * from "./getCloudletsAudienceSegmentationMatchRule";
export * from "./getCloudletsEdgeRedirectorMatchRule";
export * from "./getCloudletsForwardRewriteMatchRule";
export * from "./getCloudletsPhasedReleaseMatchRule";
export * from "./getCloudletsPolicy";
export * from "./getCloudletsVisitorPrioritizationMatchRule";
export * from "./getContract";
export * from "./getContracts";
export * from "./getCpCode";
export * from "./getDatastreamActivationHistory";
export * from "./getDatastreamDatasetFields";
export * from "./getDnsRecordSet";
export * from "./getEdgeWorkersPropertyRules";
export * from "./getEdgeWorkersResourceTier";
export * from "./getGroup";
export * from "./getGroups";
export * from "./getGtmDefaultDatacenter";
export * from "./getNetworkLists";
export * from "./getProperties";
export * from "./getProperty";
export * from "./getPropertyHostnames";
export * from "./getPropertyProducts";
export * from "./getPropertyRuleFormats";
export * from "./getPropertyRules";
export * from "./getPropertyRulesTemplate";
export * from "./gtmAsmap";
export * from "./gtmCidrmap";
export * from "./gtmDatacenter";
export * from "./gtmDomain";
export * from "./gtmGeomap";
export * from "./gtmProperty";
export * from "./gtmResource";
export * from "./networkList";
export * from "./networkListActivations";
export * from "./networkListDescription";
export * from "./networkListSubscription";
export * from "./property";
export * from "./propertyActivation";
export * from "./propertyVariables";
export * from "./provider";

// Export sub-modules:
import * as config from "./config";
import * as edgedns from "./edgedns";
import * as properties from "./properties";
import * as trafficmanagement from "./trafficmanagement";
import * as types from "./types";

export {
    config,
    edgedns,
    properties,
    trafficmanagement,
    types,
};

// Import resources to register:
import { AppSecActivations } from "./appSecActivations";
import { AppSecAdvancedSettingsEvasivePathMatch } from "./appSecAdvancedSettingsEvasivePathMatch";
import { AppSecAdvancedSettingsLogging } from "./appSecAdvancedSettingsLogging";
import { AppSecAdvancedSettingsPragmaHeader } from "./appSecAdvancedSettingsPragmaHeader";
import { AppSecAdvancedSettingsPrefetch } from "./appSecAdvancedSettingsPrefetch";
import { AppSecApiConstraintsProtection } from "./appSecApiConstraintsProtection";
import { AppSecApiRequestConstraints } from "./appSecApiRequestConstraints";
import { AppSecAttackGroup } from "./appSecAttackGroup";
import { AppSecByPassNetworkList } from "./appSecByPassNetworkList";
import { AppSecConfiguration } from "./appSecConfiguration";
import { AppSecConfigurationRename } from "./appSecConfigurationRename";
import { AppSecCustomDeny } from "./appSecCustomDeny";
import { AppSecCustomRule } from "./appSecCustomRule";
import { AppSecCustomRuleAction } from "./appSecCustomRuleAction";
import { AppSecEval } from "./appSecEval";
import { AppSecEvalGroup } from "./appSecEvalGroup";
import { AppSecEvalRule } from "./appSecEvalRule";
import { AppSecIPGeo } from "./appSecIPGeo";
import { AppSecIPGeoProtection } from "./appSecIPGeoProtection";
import { AppSecMatchTarget } from "./appSecMatchTarget";
import { AppSecMatchTargetSequence } from "./appSecMatchTargetSequence";
import { AppSecPenaltyBox } from "./appSecPenaltyBox";
import { AppSecRatePolicy } from "./appSecRatePolicy";
import { AppSecRatePolicyAction } from "./appSecRatePolicyAction";
import { AppSecRateProtection } from "./appSecRateProtection";
import { AppSecReputationProfile } from "./appSecReputationProfile";
import { AppSecReputationProfileAction } from "./appSecReputationProfileAction";
import { AppSecReputationProfileAnalysis } from "./appSecReputationProfileAnalysis";
import { AppSecReputationProtection } from "./appSecReputationProtection";
import { AppSecRule } from "./appSecRule";
import { AppSecRuleUpgrade } from "./appSecRuleUpgrade";
import { AppSecSecurityPolicy } from "./appSecSecurityPolicy";
import { AppSecSecurityPolicyRename } from "./appSecSecurityPolicyRename";
import { AppSecSelectedHostnames } from "./appSecSelectedHostnames";
import { AppSecSiemSettings } from "./appSecSiemSettings";
import { AppSecSlowPost } from "./appSecSlowPost";
import { AppSecSlowPostProtection } from "./appSecSlowPostProtection";
import { AppSecThreatIntel } from "./appSecThreatIntel";
import { AppSecVersionNodes } from "./appSecVersionNodes";
import { AppSecWafMode } from "./appSecWafMode";
import { AppSecWafProtection } from "./appSecWafProtection";
import { AppSecWapSelectedHostnames } from "./appSecWapSelectedHostnames";
import { CloudletsApplicationLoadBalancer } from "./cloudletsApplicationLoadBalancer";
import { CloudletsApplicationLoadBalancerActivation } from "./cloudletsApplicationLoadBalancerActivation";
import { CloudletsPolicy } from "./cloudletsPolicy";
import { CloudletsPolicyActivation } from "./cloudletsPolicyActivation";
import { CpCode } from "./cpCode";
import { CpsDvEnrollment } from "./cpsDvEnrollment";
import { CpsDvValidation } from "./cpsDvValidation";
import { Datastream } from "./datastream";
import { DnsRecord } from "./dnsRecord";
import { DnsZone } from "./dnsZone";
import { EdgeHostName } from "./edgeHostName";
import { EdgeKv } from "./edgeKv";
import { EdgeWorker } from "./edgeWorker";
import { EdgeWorkersActivation } from "./edgeWorkersActivation";
import { GtmAsmap } from "./gtmAsmap";
import { GtmCidrmap } from "./gtmCidrmap";
import { GtmDatacenter } from "./gtmDatacenter";
import { GtmDomain } from "./gtmDomain";
import { GtmGeomap } from "./gtmGeomap";
import { GtmProperty } from "./gtmProperty";
import { GtmResource } from "./gtmResource";
import { NetworkList } from "./networkList";
import { NetworkListActivations } from "./networkListActivations";
import { NetworkListDescription } from "./networkListDescription";
import { NetworkListSubscription } from "./networkListSubscription";
import { Property } from "./property";
import { PropertyActivation } from "./propertyActivation";
import { PropertyVariables } from "./propertyVariables";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "akamai:index/appSecActivations:AppSecActivations":
                return new AppSecActivations(name, <any>undefined, { urn })
            case "akamai:index/appSecAdvancedSettingsEvasivePathMatch:AppSecAdvancedSettingsEvasivePathMatch":
                return new AppSecAdvancedSettingsEvasivePathMatch(name, <any>undefined, { urn })
            case "akamai:index/appSecAdvancedSettingsLogging:AppSecAdvancedSettingsLogging":
                return new AppSecAdvancedSettingsLogging(name, <any>undefined, { urn })
            case "akamai:index/appSecAdvancedSettingsPragmaHeader:AppSecAdvancedSettingsPragmaHeader":
                return new AppSecAdvancedSettingsPragmaHeader(name, <any>undefined, { urn })
            case "akamai:index/appSecAdvancedSettingsPrefetch:AppSecAdvancedSettingsPrefetch":
                return new AppSecAdvancedSettingsPrefetch(name, <any>undefined, { urn })
            case "akamai:index/appSecApiConstraintsProtection:AppSecApiConstraintsProtection":
                return new AppSecApiConstraintsProtection(name, <any>undefined, { urn })
            case "akamai:index/appSecApiRequestConstraints:AppSecApiRequestConstraints":
                return new AppSecApiRequestConstraints(name, <any>undefined, { urn })
            case "akamai:index/appSecAttackGroup:AppSecAttackGroup":
                return new AppSecAttackGroup(name, <any>undefined, { urn })
            case "akamai:index/appSecByPassNetworkList:AppSecByPassNetworkList":
                return new AppSecByPassNetworkList(name, <any>undefined, { urn })
            case "akamai:index/appSecConfiguration:AppSecConfiguration":
                return new AppSecConfiguration(name, <any>undefined, { urn })
            case "akamai:index/appSecConfigurationRename:AppSecConfigurationRename":
                return new AppSecConfigurationRename(name, <any>undefined, { urn })
            case "akamai:index/appSecCustomDeny:AppSecCustomDeny":
                return new AppSecCustomDeny(name, <any>undefined, { urn })
            case "akamai:index/appSecCustomRule:AppSecCustomRule":
                return new AppSecCustomRule(name, <any>undefined, { urn })
            case "akamai:index/appSecCustomRuleAction:AppSecCustomRuleAction":
                return new AppSecCustomRuleAction(name, <any>undefined, { urn })
            case "akamai:index/appSecEval:AppSecEval":
                return new AppSecEval(name, <any>undefined, { urn })
            case "akamai:index/appSecEvalGroup:AppSecEvalGroup":
                return new AppSecEvalGroup(name, <any>undefined, { urn })
            case "akamai:index/appSecEvalRule:AppSecEvalRule":
                return new AppSecEvalRule(name, <any>undefined, { urn })
            case "akamai:index/appSecIPGeo:AppSecIPGeo":
                return new AppSecIPGeo(name, <any>undefined, { urn })
            case "akamai:index/appSecIPGeoProtection:AppSecIPGeoProtection":
                return new AppSecIPGeoProtection(name, <any>undefined, { urn })
            case "akamai:index/appSecMatchTarget:AppSecMatchTarget":
                return new AppSecMatchTarget(name, <any>undefined, { urn })
            case "akamai:index/appSecMatchTargetSequence:AppSecMatchTargetSequence":
                return new AppSecMatchTargetSequence(name, <any>undefined, { urn })
            case "akamai:index/appSecPenaltyBox:AppSecPenaltyBox":
                return new AppSecPenaltyBox(name, <any>undefined, { urn })
            case "akamai:index/appSecRatePolicy:AppSecRatePolicy":
                return new AppSecRatePolicy(name, <any>undefined, { urn })
            case "akamai:index/appSecRatePolicyAction:AppSecRatePolicyAction":
                return new AppSecRatePolicyAction(name, <any>undefined, { urn })
            case "akamai:index/appSecRateProtection:AppSecRateProtection":
                return new AppSecRateProtection(name, <any>undefined, { urn })
            case "akamai:index/appSecReputationProfile:AppSecReputationProfile":
                return new AppSecReputationProfile(name, <any>undefined, { urn })
            case "akamai:index/appSecReputationProfileAction:AppSecReputationProfileAction":
                return new AppSecReputationProfileAction(name, <any>undefined, { urn })
            case "akamai:index/appSecReputationProfileAnalysis:AppSecReputationProfileAnalysis":
                return new AppSecReputationProfileAnalysis(name, <any>undefined, { urn })
            case "akamai:index/appSecReputationProtection:AppSecReputationProtection":
                return new AppSecReputationProtection(name, <any>undefined, { urn })
            case "akamai:index/appSecRule:AppSecRule":
                return new AppSecRule(name, <any>undefined, { urn })
            case "akamai:index/appSecRuleUpgrade:AppSecRuleUpgrade":
                return new AppSecRuleUpgrade(name, <any>undefined, { urn })
            case "akamai:index/appSecSecurityPolicy:AppSecSecurityPolicy":
                return new AppSecSecurityPolicy(name, <any>undefined, { urn })
            case "akamai:index/appSecSecurityPolicyRename:AppSecSecurityPolicyRename":
                return new AppSecSecurityPolicyRename(name, <any>undefined, { urn })
            case "akamai:index/appSecSelectedHostnames:AppSecSelectedHostnames":
                return new AppSecSelectedHostnames(name, <any>undefined, { urn })
            case "akamai:index/appSecSiemSettings:AppSecSiemSettings":
                return new AppSecSiemSettings(name, <any>undefined, { urn })
            case "akamai:index/appSecSlowPost:AppSecSlowPost":
                return new AppSecSlowPost(name, <any>undefined, { urn })
            case "akamai:index/appSecSlowPostProtection:AppSecSlowPostProtection":
                return new AppSecSlowPostProtection(name, <any>undefined, { urn })
            case "akamai:index/appSecThreatIntel:AppSecThreatIntel":
                return new AppSecThreatIntel(name, <any>undefined, { urn })
            case "akamai:index/appSecVersionNodes:AppSecVersionNodes":
                return new AppSecVersionNodes(name, <any>undefined, { urn })
            case "akamai:index/appSecWafMode:AppSecWafMode":
                return new AppSecWafMode(name, <any>undefined, { urn })
            case "akamai:index/appSecWafProtection:AppSecWafProtection":
                return new AppSecWafProtection(name, <any>undefined, { urn })
            case "akamai:index/appSecWapSelectedHostnames:AppSecWapSelectedHostnames":
                return new AppSecWapSelectedHostnames(name, <any>undefined, { urn })
            case "akamai:index/cloudletsApplicationLoadBalancer:CloudletsApplicationLoadBalancer":
                return new CloudletsApplicationLoadBalancer(name, <any>undefined, { urn })
            case "akamai:index/cloudletsApplicationLoadBalancerActivation:CloudletsApplicationLoadBalancerActivation":
                return new CloudletsApplicationLoadBalancerActivation(name, <any>undefined, { urn })
            case "akamai:index/cloudletsPolicy:CloudletsPolicy":
                return new CloudletsPolicy(name, <any>undefined, { urn })
            case "akamai:index/cloudletsPolicyActivation:CloudletsPolicyActivation":
                return new CloudletsPolicyActivation(name, <any>undefined, { urn })
            case "akamai:index/cpCode:CpCode":
                return new CpCode(name, <any>undefined, { urn })
            case "akamai:index/cpsDvEnrollment:CpsDvEnrollment":
                return new CpsDvEnrollment(name, <any>undefined, { urn })
            case "akamai:index/cpsDvValidation:CpsDvValidation":
                return new CpsDvValidation(name, <any>undefined, { urn })
            case "akamai:index/datastream:Datastream":
                return new Datastream(name, <any>undefined, { urn })
            case "akamai:index/dnsRecord:DnsRecord":
                return new DnsRecord(name, <any>undefined, { urn })
            case "akamai:index/dnsZone:DnsZone":
                return new DnsZone(name, <any>undefined, { urn })
            case "akamai:index/edgeHostName:EdgeHostName":
                return new EdgeHostName(name, <any>undefined, { urn })
            case "akamai:index/edgeKv:EdgeKv":
                return new EdgeKv(name, <any>undefined, { urn })
            case "akamai:index/edgeWorker:EdgeWorker":
                return new EdgeWorker(name, <any>undefined, { urn })
            case "akamai:index/edgeWorkersActivation:EdgeWorkersActivation":
                return new EdgeWorkersActivation(name, <any>undefined, { urn })
            case "akamai:index/gtmAsmap:GtmAsmap":
                return new GtmAsmap(name, <any>undefined, { urn })
            case "akamai:index/gtmCidrmap:GtmCidrmap":
                return new GtmCidrmap(name, <any>undefined, { urn })
            case "akamai:index/gtmDatacenter:GtmDatacenter":
                return new GtmDatacenter(name, <any>undefined, { urn })
            case "akamai:index/gtmDomain:GtmDomain":
                return new GtmDomain(name, <any>undefined, { urn })
            case "akamai:index/gtmGeomap:GtmGeomap":
                return new GtmGeomap(name, <any>undefined, { urn })
            case "akamai:index/gtmProperty:GtmProperty":
                return new GtmProperty(name, <any>undefined, { urn })
            case "akamai:index/gtmResource:GtmResource":
                return new GtmResource(name, <any>undefined, { urn })
            case "akamai:index/networkList:NetworkList":
                return new NetworkList(name, <any>undefined, { urn })
            case "akamai:index/networkListActivations:NetworkListActivations":
                return new NetworkListActivations(name, <any>undefined, { urn })
            case "akamai:index/networkListDescription:NetworkListDescription":
                return new NetworkListDescription(name, <any>undefined, { urn })
            case "akamai:index/networkListSubscription:NetworkListSubscription":
                return new NetworkListSubscription(name, <any>undefined, { urn })
            case "akamai:index/property:Property":
                return new Property(name, <any>undefined, { urn })
            case "akamai:index/propertyActivation:PropertyActivation":
                return new PropertyActivation(name, <any>undefined, { urn })
            case "akamai:index/propertyVariables:PropertyVariables":
                return new PropertyVariables(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("akamai", "index/appSecActivations", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecAdvancedSettingsEvasivePathMatch", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecAdvancedSettingsLogging", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecAdvancedSettingsPragmaHeader", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecAdvancedSettingsPrefetch", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecApiConstraintsProtection", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecApiRequestConstraints", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecAttackGroup", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecByPassNetworkList", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecConfiguration", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecConfigurationRename", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecCustomDeny", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecCustomRule", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecCustomRuleAction", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecEval", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecEvalGroup", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecEvalRule", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecIPGeo", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecIPGeoProtection", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecMatchTarget", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecMatchTargetSequence", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecPenaltyBox", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecRatePolicy", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecRatePolicyAction", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecRateProtection", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecReputationProfile", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecReputationProfileAction", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecReputationProfileAnalysis", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecReputationProtection", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecRule", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecRuleUpgrade", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecSecurityPolicy", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecSecurityPolicyRename", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecSelectedHostnames", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecSiemSettings", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecSlowPost", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecSlowPostProtection", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecThreatIntel", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecVersionNodes", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecWafMode", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecWafProtection", _module)
pulumi.runtime.registerResourceModule("akamai", "index/appSecWapSelectedHostnames", _module)
pulumi.runtime.registerResourceModule("akamai", "index/cloudletsApplicationLoadBalancer", _module)
pulumi.runtime.registerResourceModule("akamai", "index/cloudletsApplicationLoadBalancerActivation", _module)
pulumi.runtime.registerResourceModule("akamai", "index/cloudletsPolicy", _module)
pulumi.runtime.registerResourceModule("akamai", "index/cloudletsPolicyActivation", _module)
pulumi.runtime.registerResourceModule("akamai", "index/cpCode", _module)
pulumi.runtime.registerResourceModule("akamai", "index/cpsDvEnrollment", _module)
pulumi.runtime.registerResourceModule("akamai", "index/cpsDvValidation", _module)
pulumi.runtime.registerResourceModule("akamai", "index/datastream", _module)
pulumi.runtime.registerResourceModule("akamai", "index/dnsRecord", _module)
pulumi.runtime.registerResourceModule("akamai", "index/dnsZone", _module)
pulumi.runtime.registerResourceModule("akamai", "index/edgeHostName", _module)
pulumi.runtime.registerResourceModule("akamai", "index/edgeKv", _module)
pulumi.runtime.registerResourceModule("akamai", "index/edgeWorker", _module)
pulumi.runtime.registerResourceModule("akamai", "index/edgeWorkersActivation", _module)
pulumi.runtime.registerResourceModule("akamai", "index/gtmAsmap", _module)
pulumi.runtime.registerResourceModule("akamai", "index/gtmCidrmap", _module)
pulumi.runtime.registerResourceModule("akamai", "index/gtmDatacenter", _module)
pulumi.runtime.registerResourceModule("akamai", "index/gtmDomain", _module)
pulumi.runtime.registerResourceModule("akamai", "index/gtmGeomap", _module)
pulumi.runtime.registerResourceModule("akamai", "index/gtmProperty", _module)
pulumi.runtime.registerResourceModule("akamai", "index/gtmResource", _module)
pulumi.runtime.registerResourceModule("akamai", "index/networkList", _module)
pulumi.runtime.registerResourceModule("akamai", "index/networkListActivations", _module)
pulumi.runtime.registerResourceModule("akamai", "index/networkListDescription", _module)
pulumi.runtime.registerResourceModule("akamai", "index/networkListSubscription", _module)
pulumi.runtime.registerResourceModule("akamai", "index/property", _module)
pulumi.runtime.registerResourceModule("akamai", "index/propertyActivation", _module)
pulumi.runtime.registerResourceModule("akamai", "index/propertyVariables", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("akamai", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:akamai") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
