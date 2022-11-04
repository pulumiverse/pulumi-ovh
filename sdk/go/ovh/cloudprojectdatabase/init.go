// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh"
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
	case "ovh:CloudProjectDatabase/databaseInstance:DatabaseInstance":
		r = &DatabaseInstance{}
	case "ovh:CloudProjectDatabase/integration:Integration":
		r = &Integration{}
	case "ovh:CloudProjectDatabase/ipRestriction:IpRestriction":
		r = &IpRestriction{}
	case "ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl":
		r = &KafkaAcl{}
	case "ovh:CloudProjectDatabase/kafkaTopic:KafkaTopic":
		r = &KafkaTopic{}
	case "ovh:CloudProjectDatabase/m3DbNamespace:M3DbNamespace":
		r = &M3DbNamespace{}
	case "ovh:CloudProjectDatabase/m3DbUser:M3DbUser":
		r = &M3DbUser{}
	case "ovh:CloudProjectDatabase/mongoDbUser:MongoDbUser":
		r = &MongoDbUser{}
	case "ovh:CloudProjectDatabase/opensearchPattern:OpensearchPattern":
		r = &OpensearchPattern{}
	case "ovh:CloudProjectDatabase/opensearchUser:OpensearchUser":
		r = &OpensearchUser{}
	case "ovh:CloudProjectDatabase/postgresSqlUser:PostgresSqlUser":
		r = &PostgresSqlUser{}
	case "ovh:CloudProjectDatabase/redisUser:RedisUser":
		r = &RedisUser{}
	case "ovh:CloudProjectDatabase/user:User":
		r = &User{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := ovh.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"ovh",
		"CloudProjectDatabase/databaseInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"CloudProjectDatabase/integration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"CloudProjectDatabase/ipRestriction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"CloudProjectDatabase/kafkaAcl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"CloudProjectDatabase/kafkaTopic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"CloudProjectDatabase/m3DbNamespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"CloudProjectDatabase/m3DbUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"CloudProjectDatabase/mongoDbUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"CloudProjectDatabase/opensearchPattern",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"CloudProjectDatabase/opensearchUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"CloudProjectDatabase/postgresSqlUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"CloudProjectDatabase/redisUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"CloudProjectDatabase/user",
		&module{version},
	)
}
