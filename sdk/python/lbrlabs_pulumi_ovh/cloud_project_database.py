# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['CloudProjectDatabaseArgs', 'CloudProjectDatabase']

@pulumi.input_type
class CloudProjectDatabaseArgs:
    def __init__(__self__, *,
                 engine: pulumi.Input[str],
                 flavor: pulumi.Input[str],
                 nodes: pulumi.Input[Sequence[pulumi.Input['CloudProjectDatabaseNodeArgs']]],
                 plan: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 version: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CloudProjectDatabase resource.
        :param pulumi.Input[str] engine: The database engine you want to deploy. To get a full list of available engine visit.
               [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        :param pulumi.Input[str] flavor: a valid OVH public cloud database flavor name in which the nodes will be started.
               Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
               You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
        :param pulumi.Input[Sequence[pulumi.Input['CloudProjectDatabaseNodeArgs']]] nodes: List of nodes object.
               Multi region cluster are not yet available, all node should be identical.
        :param pulumi.Input[str] plan: List of nodes object.
               Enum: "esential", "business", "enterprise".
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] version: The version of the engine in which the service should be deployed
        :param pulumi.Input[str] description: Small description of the database service.
        """
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "flavor", flavor)
        pulumi.set(__self__, "nodes", nodes)
        pulumi.set(__self__, "plan", plan)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "version", version)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Input[str]:
        """
        The database engine you want to deploy. To get a full list of available engine visit.
        [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: pulumi.Input[str]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter
    def flavor(self) -> pulumi.Input[str]:
        """
        a valid OVH public cloud database flavor name in which the nodes will be started.
        Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
        You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
        """
        return pulumi.get(self, "flavor")

    @flavor.setter
    def flavor(self, value: pulumi.Input[str]):
        pulumi.set(self, "flavor", value)

    @property
    @pulumi.getter
    def nodes(self) -> pulumi.Input[Sequence[pulumi.Input['CloudProjectDatabaseNodeArgs']]]:
        """
        List of nodes object.
        Multi region cluster are not yet available, all node should be identical.
        """
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: pulumi.Input[Sequence[pulumi.Input['CloudProjectDatabaseNodeArgs']]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Input[str]:
        """
        List of nodes object.
        Enum: "esential", "business", "enterprise".
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: pulumi.Input[str]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        """
        The version of the engine in which the service should be deployed
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[str]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Small description of the database service.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _CloudProjectDatabaseState:
    def __init__(__self__, *,
                 backup_time: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input['CloudProjectDatabaseEndpointArgs']]]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 flavor: Optional[pulumi.Input[str]] = None,
                 maintenance_time: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input['CloudProjectDatabaseNodeArgs']]]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CloudProjectDatabase resources.
        :param pulumi.Input[str] backup_time: Time on which backups start every day.
        :param pulumi.Input[str] created_at: Date of the creation of the cluster.
        :param pulumi.Input[str] description: Small description of the database service.
        :param pulumi.Input[Sequence[pulumi.Input['CloudProjectDatabaseEndpointArgs']]] endpoints: List of all endpoints objects of the service.
        :param pulumi.Input[str] engine: The database engine you want to deploy. To get a full list of available engine visit.
               [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        :param pulumi.Input[str] flavor: a valid OVH public cloud database flavor name in which the nodes will be started.
               Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
               You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
        :param pulumi.Input[str] maintenance_time: Time on which maintenances can start every day.
        :param pulumi.Input[str] network_type: Type of network of the cluster.
        :param pulumi.Input[Sequence[pulumi.Input['CloudProjectDatabaseNodeArgs']]] nodes: List of nodes object.
               Multi region cluster are not yet available, all node should be identical.
        :param pulumi.Input[str] plan: List of nodes object.
               Enum: "esential", "business", "enterprise".
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] status: Current status of the cluster.
        :param pulumi.Input[str] version: The version of the engine in which the service should be deployed
        """
        if backup_time is not None:
            pulumi.set(__self__, "backup_time", backup_time)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if endpoints is not None:
            pulumi.set(__self__, "endpoints", endpoints)
        if engine is not None:
            pulumi.set(__self__, "engine", engine)
        if flavor is not None:
            pulumi.set(__self__, "flavor", flavor)
        if maintenance_time is not None:
            pulumi.set(__self__, "maintenance_time", maintenance_time)
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if nodes is not None:
            pulumi.set(__self__, "nodes", nodes)
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="backupTime")
    def backup_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time on which backups start every day.
        """
        return pulumi.get(self, "backup_time")

    @backup_time.setter
    def backup_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_time", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date of the creation of the cluster.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Small description of the database service.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def endpoints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CloudProjectDatabaseEndpointArgs']]]]:
        """
        List of all endpoints objects of the service.
        """
        return pulumi.get(self, "endpoints")

    @endpoints.setter
    def endpoints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CloudProjectDatabaseEndpointArgs']]]]):
        pulumi.set(self, "endpoints", value)

    @property
    @pulumi.getter
    def engine(self) -> Optional[pulumi.Input[str]]:
        """
        The database engine you want to deploy. To get a full list of available engine visit.
        [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter
    def flavor(self) -> Optional[pulumi.Input[str]]:
        """
        a valid OVH public cloud database flavor name in which the nodes will be started.
        Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
        You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
        """
        return pulumi.get(self, "flavor")

    @flavor.setter
    def flavor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flavor", value)

    @property
    @pulumi.getter(name="maintenanceTime")
    def maintenance_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time on which maintenances can start every day.
        """
        return pulumi.get(self, "maintenance_time")

    @maintenance_time.setter
    def maintenance_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "maintenance_time", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of network of the cluster.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter
    def nodes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CloudProjectDatabaseNodeArgs']]]]:
        """
        List of nodes object.
        Multi region cluster are not yet available, all node should be identical.
        """
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CloudProjectDatabaseNodeArgs']]]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input[str]]:
        """
        List of nodes object.
        Enum: "esential", "business", "enterprise".
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Current status of the cluster.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the engine in which the service should be deployed
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class CloudProjectDatabase(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 flavor: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudProjectDatabaseNodeArgs']]]]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        To deploy a business PostgreSQL service with two nodes on public network:
        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh

        postgresql = ovh.CloudProjectDatabase("postgresql",
            description="my-first-postgresql",
            engine="postgresql",
            flavor="db1-15",
            nodes=[
                ovh.CloudProjectDatabaseNodeArgs(
                    region="GRA",
                ),
                ovh.CloudProjectDatabaseNodeArgs(
                    region="GRA",
                ),
            ],
            plan="business",
            service_name="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            version="14")
        ```

        To deploy an enterprise MongoDB service with three nodes on private network:
        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh

        mongodb = ovh.CloudProjectDatabase("mongodb",
            service_name=var["openstack_infos"]["project_id"],
            description="my-first-mongodb",
            engine="mongodb",
            version="5.0",
            plan="enterprise",
            nodes=[
                ovh.CloudProjectDatabaseNodeArgs(
                    region="SBG",
                    subnet_id="XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
                    network_id="XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
                ),
                ovh.CloudProjectDatabaseNodeArgs(
                    region="SBG",
                    subnet_id="XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
                    network_id="XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
                ),
                ovh.CloudProjectDatabaseNodeArgs(
                    region="SBG",
                    subnet_id="XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
                    network_id="XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
                ),
            ],
            flavor="db1-30")
        ```

        ## Import

        OVHcloud Managed database clusters can be imported using the `service_name`, `engine`, `id` of the cluster, separated by "/" E.g.,

        ```sh
         $ pulumi import ovh:index/cloudProjectDatabase:CloudProjectDatabase my_database_cluster <service_name>/<engine>/<id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Small description of the database service.
        :param pulumi.Input[str] engine: The database engine you want to deploy. To get a full list of available engine visit.
               [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        :param pulumi.Input[str] flavor: a valid OVH public cloud database flavor name in which the nodes will be started.
               Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
               You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudProjectDatabaseNodeArgs']]]] nodes: List of nodes object.
               Multi region cluster are not yet available, all node should be identical.
        :param pulumi.Input[str] plan: List of nodes object.
               Enum: "esential", "business", "enterprise".
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] version: The version of the engine in which the service should be deployed
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CloudProjectDatabaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        To deploy a business PostgreSQL service with two nodes on public network:
        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh

        postgresql = ovh.CloudProjectDatabase("postgresql",
            description="my-first-postgresql",
            engine="postgresql",
            flavor="db1-15",
            nodes=[
                ovh.CloudProjectDatabaseNodeArgs(
                    region="GRA",
                ),
                ovh.CloudProjectDatabaseNodeArgs(
                    region="GRA",
                ),
            ],
            plan="business",
            service_name="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            version="14")
        ```

        To deploy an enterprise MongoDB service with three nodes on private network:
        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh

        mongodb = ovh.CloudProjectDatabase("mongodb",
            service_name=var["openstack_infos"]["project_id"],
            description="my-first-mongodb",
            engine="mongodb",
            version="5.0",
            plan="enterprise",
            nodes=[
                ovh.CloudProjectDatabaseNodeArgs(
                    region="SBG",
                    subnet_id="XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
                    network_id="XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
                ),
                ovh.CloudProjectDatabaseNodeArgs(
                    region="SBG",
                    subnet_id="XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
                    network_id="XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
                ),
                ovh.CloudProjectDatabaseNodeArgs(
                    region="SBG",
                    subnet_id="XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
                    network_id="XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
                ),
            ],
            flavor="db1-30")
        ```

        ## Import

        OVHcloud Managed database clusters can be imported using the `service_name`, `engine`, `id` of the cluster, separated by "/" E.g.,

        ```sh
         $ pulumi import ovh:index/cloudProjectDatabase:CloudProjectDatabase my_database_cluster <service_name>/<engine>/<id>
        ```

        :param str resource_name: The name of the resource.
        :param CloudProjectDatabaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CloudProjectDatabaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 flavor: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudProjectDatabaseNodeArgs']]]]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CloudProjectDatabaseArgs.__new__(CloudProjectDatabaseArgs)

            __props__.__dict__["description"] = description
            if engine is None and not opts.urn:
                raise TypeError("Missing required property 'engine'")
            __props__.__dict__["engine"] = engine
            if flavor is None and not opts.urn:
                raise TypeError("Missing required property 'flavor'")
            __props__.__dict__["flavor"] = flavor
            if nodes is None and not opts.urn:
                raise TypeError("Missing required property 'nodes'")
            __props__.__dict__["nodes"] = nodes
            if plan is None and not opts.urn:
                raise TypeError("Missing required property 'plan'")
            __props__.__dict__["plan"] = plan
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__.__dict__["version"] = version
            __props__.__dict__["backup_time"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["endpoints"] = None
            __props__.__dict__["maintenance_time"] = None
            __props__.__dict__["network_type"] = None
            __props__.__dict__["status"] = None
        super(CloudProjectDatabase, __self__).__init__(
            'ovh:index/cloudProjectDatabase:CloudProjectDatabase',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_time: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudProjectDatabaseEndpointArgs']]]]] = None,
            engine: Optional[pulumi.Input[str]] = None,
            flavor: Optional[pulumi.Input[str]] = None,
            maintenance_time: Optional[pulumi.Input[str]] = None,
            network_type: Optional[pulumi.Input[str]] = None,
            nodes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudProjectDatabaseNodeArgs']]]]] = None,
            plan: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'CloudProjectDatabase':
        """
        Get an existing CloudProjectDatabase resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_time: Time on which backups start every day.
        :param pulumi.Input[str] created_at: Date of the creation of the cluster.
        :param pulumi.Input[str] description: Small description of the database service.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudProjectDatabaseEndpointArgs']]]] endpoints: List of all endpoints objects of the service.
        :param pulumi.Input[str] engine: The database engine you want to deploy. To get a full list of available engine visit.
               [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        :param pulumi.Input[str] flavor: a valid OVH public cloud database flavor name in which the nodes will be started.
               Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
               You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
        :param pulumi.Input[str] maintenance_time: Time on which maintenances can start every day.
        :param pulumi.Input[str] network_type: Type of network of the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudProjectDatabaseNodeArgs']]]] nodes: List of nodes object.
               Multi region cluster are not yet available, all node should be identical.
        :param pulumi.Input[str] plan: List of nodes object.
               Enum: "esential", "business", "enterprise".
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] status: Current status of the cluster.
        :param pulumi.Input[str] version: The version of the engine in which the service should be deployed
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CloudProjectDatabaseState.__new__(_CloudProjectDatabaseState)

        __props__.__dict__["backup_time"] = backup_time
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["endpoints"] = endpoints
        __props__.__dict__["engine"] = engine
        __props__.__dict__["flavor"] = flavor
        __props__.__dict__["maintenance_time"] = maintenance_time
        __props__.__dict__["network_type"] = network_type
        __props__.__dict__["nodes"] = nodes
        __props__.__dict__["plan"] = plan
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["status"] = status
        __props__.__dict__["version"] = version
        return CloudProjectDatabase(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupTime")
    def backup_time(self) -> pulumi.Output[str]:
        """
        Time on which backups start every day.
        """
        return pulumi.get(self, "backup_time")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Date of the creation of the cluster.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Small description of the database service.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def endpoints(self) -> pulumi.Output[Sequence['outputs.CloudProjectDatabaseEndpoint']]:
        """
        List of all endpoints objects of the service.
        """
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[str]:
        """
        The database engine you want to deploy. To get a full list of available engine visit.
        [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def flavor(self) -> pulumi.Output[str]:
        """
        a valid OVH public cloud database flavor name in which the nodes will be started.
        Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
        You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
        """
        return pulumi.get(self, "flavor")

    @property
    @pulumi.getter(name="maintenanceTime")
    def maintenance_time(self) -> pulumi.Output[str]:
        """
        Time on which maintenances can start every day.
        """
        return pulumi.get(self, "maintenance_time")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> pulumi.Output[str]:
        """
        Type of network of the cluster.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter
    def nodes(self) -> pulumi.Output[Sequence['outputs.CloudProjectDatabaseNode']]:
        """
        List of nodes object.
        Multi region cluster are not yet available, all node should be identical.
        """
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output[str]:
        """
        List of nodes object.
        Enum: "esential", "business", "enterprise".
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Current status of the cluster.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The version of the engine in which the service should be deployed
        """
        return pulumi.get(self, "version")

