// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package properties

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-akamai/sdk/v3/go/akamai"
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
	case "akamai:properties/cpCode:CpCode":
		r = &CpCode{}
	case "akamai:properties/edgeHostName:EdgeHostName":
		r = &EdgeHostName{}
	case "akamai:properties/property:Property":
		r = &Property{}
	case "akamai:properties/propertyActivation:PropertyActivation":
		r = &PropertyActivation{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := akamai.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"akamai",
		"properties/cpCode",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"properties/edgeHostName",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"properties/property",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"akamai",
		"properties/propertyActivation",
		&module{version},
	)
}
