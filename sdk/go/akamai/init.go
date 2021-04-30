// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "akamai:index/appSecActivations:AppSecActivations":
		r = &AppSecActivations{}
	case "akamai:index/appSecAdvancedSettingsLogging:AppSecAdvancedSettingsLogging":
		r = &AppSecAdvancedSettingsLogging{}
	case "akamai:index/appSecAdvancedSettingsPrefetch:AppSecAdvancedSettingsPrefetch":
		r = &AppSecAdvancedSettingsPrefetch{}
	case "akamai:index/appSecApiRequestConstraints:AppSecApiRequestConstraints":
		r = &AppSecApiRequestConstraints{}
	case "akamai:index/appSecAttackGroupAction:AppSecAttackGroupAction":
		r = &AppSecAttackGroupAction{}
	case "akamai:index/appSecAttackGroupActionConditionException:AppSecAttackGroupActionConditionException":
		r = &AppSecAttackGroupActionConditionException{}
	case "akamai:index/appSecByPassNetworkList:AppSecByPassNetworkList":
		r = &AppSecByPassNetworkList{}
	case "akamai:index/appSecConfiguration:AppSecConfiguration":
		r = &AppSecConfiguration{}
	case "akamai:index/appSecConfigurationClone:AppSecConfigurationClone":
		r = &AppSecConfigurationClone{}
	case "akamai:index/appSecConfigurationRename:AppSecConfigurationRename":
		r = &AppSecConfigurationRename{}
	case "akamai:index/appSecConfigurationVersionClone:AppSecConfigurationVersionClone":
		r = &AppSecConfigurationVersionClone{}
	case "akamai:index/appSecCustomDeny:AppSecCustomDeny":
		r = &AppSecCustomDeny{}
	case "akamai:index/appSecCustomRule:AppSecCustomRule":
		r = &AppSecCustomRule{}
	case "akamai:index/appSecCustomRuleAction:AppSecCustomRuleAction":
		r = &AppSecCustomRuleAction{}
	case "akamai:index/appSecEval:AppSecEval":
		r = &AppSecEval{}
	case "akamai:index/appSecEvalHostnames:AppSecEvalHostnames":
		r = &AppSecEvalHostnames{}
	case "akamai:index/appSecEvalProtectHost:AppSecEvalProtectHost":
		r = &AppSecEvalProtectHost{}
	case "akamai:index/appSecEvalRuleAction:AppSecEvalRuleAction":
		r = &AppSecEvalRuleAction{}
	case "akamai:index/appSecEvalRuleConditionException:AppSecEvalRuleConditionException":
		r = &AppSecEvalRuleConditionException{}
	case "akamai:index/appSecIPGeo:AppSecIPGeo":
		r = &AppSecIPGeo{}
	case "akamai:index/appSecMatchTarget:AppSecMatchTarget":
		r = &AppSecMatchTarget{}
	case "akamai:index/appSecMatchTargetSequence:AppSecMatchTargetSequence":
		r = &AppSecMatchTargetSequence{}
	case "akamai:index/appSecPenaltyBox:AppSecPenaltyBox":
		r = &AppSecPenaltyBox{}
	case "akamai:index/appSecRatePolicy:AppSecRatePolicy":
		r = &AppSecRatePolicy{}
	case "akamai:index/appSecRatePolicyAction:AppSecRatePolicyAction":
		r = &AppSecRatePolicyAction{}
	case "akamai:index/appSecRateProtection:AppSecRateProtection":
		r = &AppSecRateProtection{}
	case "akamai:index/appSecReputationProfile:AppSecReputationProfile":
		r = &AppSecReputationProfile{}
	case "akamai:index/appSecReputationProfileAction:AppSecReputationProfileAction":
		r = &AppSecReputationProfileAction{}
	case "akamai:index/appSecReputationProfileAnalysis:AppSecReputationProfileAnalysis":
		r = &AppSecReputationProfileAnalysis{}
	case "akamai:index/appSecReputationProtection:AppSecReputationProtection":
		r = &AppSecReputationProtection{}
	case "akamai:index/appSecRuleAction:AppSecRuleAction":
		r = &AppSecRuleAction{}
	case "akamai:index/appSecRuleConditionException:AppSecRuleConditionException":
		r = &AppSecRuleConditionException{}
	case "akamai:index/appSecRuleUpgrade:AppSecRuleUpgrade":
		r = &AppSecRuleUpgrade{}
	case "akamai:index/appSecSecurityPolicy:AppSecSecurityPolicy":
		r = &AppSecSecurityPolicy{}
	case "akamai:index/appSecSecurityPolicyClone:AppSecSecurityPolicyClone":
		r = &AppSecSecurityPolicyClone{}
	case "akamai:index/appSecSecurityPolicyProtections:AppSecSecurityPolicyProtections":
		r = &AppSecSecurityPolicyProtections{}
	case "akamai:index/appSecSecurityPolicyRename:AppSecSecurityPolicyRename":
		r = &AppSecSecurityPolicyRename{}
	case "akamai:index/appSecSelectedHostnames:AppSecSelectedHostnames":
		r = &AppSecSelectedHostnames{}
	case "akamai:index/appSecSiemSettings:AppSecSiemSettings":
		r = &AppSecSiemSettings{}
	case "akamai:index/appSecSlowPost:AppSecSlowPost":
		r = &AppSecSlowPost{}
	case "akamai:index/appSecSlowPostProtection:AppSecSlowPostProtection":
		r = &AppSecSlowPostProtection{}
	case "akamai:index/appSecVersionNodes:AppSecVersionNodes":
		r = &AppSecVersionNodes{}
	case "akamai:index/appSecWafMode:AppSecWafMode":
		r = &AppSecWafMode{}
	case "akamai:index/appSecWafProtection:AppSecWafProtection":
		r = &AppSecWafProtection{}
	case "akamai:index/cpCode:CpCode":
		r = &CpCode{}
	case "akamai:index/dnsRecord:DnsRecord":
		r = &DnsRecord{}
	case "akamai:index/dnsZone:DnsZone":
		r = &DnsZone{}
	case "akamai:index/edgeHostName:EdgeHostName":
		r = &EdgeHostName{}
	case "akamai:index/gtmAsmap:GtmAsmap":
		r = &GtmAsmap{}
	case "akamai:index/gtmCidrmap:GtmCidrmap":
		r = &GtmCidrmap{}
	case "akamai:index/gtmDatacenter:GtmDatacenter":
		r = &GtmDatacenter{}
	case "akamai:index/gtmDomain:GtmDomain":
		r = &GtmDomain{}
	case "akamai:index/gtmGeomap:GtmGeomap":
		r = &GtmGeomap{}
	case "akamai:index/gtmProperty:GtmProperty":
		r = &GtmProperty{}
	case "akamai:index/gtmResource:GtmResource":
		r = &GtmResource{}
	case "akamai:index/property:Property":
		r = &Property{}
	case "akamai:index/propertyActivation:PropertyActivation":
		r = &PropertyActivation{}
	case "akamai:index/propertyVariables:PropertyVariables":
		r = &PropertyVariables{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:akamai" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecActivations",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecAdvancedSettingsLogging",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecAdvancedSettingsPrefetch",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecApiRequestConstraints",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecAttackGroupAction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecAttackGroupActionConditionException",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecByPassNetworkList",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecConfigurationClone",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecConfigurationRename",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecConfigurationVersionClone",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecCustomDeny",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecCustomRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecCustomRuleAction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecEval",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecEvalHostnames",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecEvalProtectHost",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecEvalRuleAction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecEvalRuleConditionException",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecIPGeo",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecMatchTarget",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecMatchTargetSequence",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecPenaltyBox",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecRatePolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecRatePolicyAction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecRateProtection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecReputationProfile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecReputationProfileAction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecReputationProfileAnalysis",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecReputationProtection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecRuleAction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecRuleConditionException",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecRuleUpgrade",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecSecurityPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecSecurityPolicyClone",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecSecurityPolicyProtections",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecSecurityPolicyRename",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecSelectedHostnames",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecSiemSettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecSlowPost",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecSlowPostProtection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecVersionNodes",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecWafMode",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/appSecWafProtection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/cpCode",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/dnsRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/dnsZone",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/edgeHostName",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/gtmAsmap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/gtmCidrmap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/gtmDatacenter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/gtmDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/gtmGeomap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/gtmProperty",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/gtmResource",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/property",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/propertyActivation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"index/propertyVariables",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"akamai",
		&pkg{version},
	)
}
