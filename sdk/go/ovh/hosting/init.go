// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hosting

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/internal"
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
	case "ovh:Hosting/privateDatabase:PrivateDatabase":
		r = &PrivateDatabase{}
	case "ovh:Hosting/privateDatabaseAllowlist:PrivateDatabaseAllowlist":
		r = &PrivateDatabaseAllowlist{}
	case "ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb":
		r = &PrivateDatabaseDb{}
	case "ovh:Hosting/privateDatabaseUser:PrivateDatabaseUser":
		r = &PrivateDatabaseUser{}
	case "ovh:Hosting/privateDatabaseUserGrant:PrivateDatabaseUserGrant":
		r = &PrivateDatabaseUserGrant{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"ovh",
		"Hosting/privateDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Hosting/privateDatabaseAllowlist",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Hosting/privateDatabaseDb",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Hosting/privateDatabaseUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Hosting/privateDatabaseUserGrant",
		&module{version},
	)
}
