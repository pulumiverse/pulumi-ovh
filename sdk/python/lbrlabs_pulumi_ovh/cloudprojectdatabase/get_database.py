# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetDatabaseResult',
    'AwaitableGetDatabaseResult',
    'get_database',
    'get_database_output',
]

@pulumi.output_type
class GetDatabaseResult:
    """
    A collection of values returned by getDatabase.
    """
    def __init__(__self__, advanced_configuration=None, backup_time=None, created_at=None, description=None, disk_size=None, disk_type=None, endpoints=None, engine=None, flavor=None, id=None, kafka_rest_api=None, maintenance_time=None, network_type=None, nodes=None, opensearch_acls_enabled=None, plan=None, service_name=None, status=None, version=None):
        if advanced_configuration and not isinstance(advanced_configuration, dict):
            raise TypeError("Expected argument 'advanced_configuration' to be a dict")
        pulumi.set(__self__, "advanced_configuration", advanced_configuration)
        if backup_time and not isinstance(backup_time, str):
            raise TypeError("Expected argument 'backup_time' to be a str")
        pulumi.set(__self__, "backup_time", backup_time)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_size and not isinstance(disk_size, int):
            raise TypeError("Expected argument 'disk_size' to be a int")
        pulumi.set(__self__, "disk_size", disk_size)
        if disk_type and not isinstance(disk_type, str):
            raise TypeError("Expected argument 'disk_type' to be a str")
        pulumi.set(__self__, "disk_type", disk_type)
        if endpoints and not isinstance(endpoints, list):
            raise TypeError("Expected argument 'endpoints' to be a list")
        pulumi.set(__self__, "endpoints", endpoints)
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        pulumi.set(__self__, "engine", engine)
        if flavor and not isinstance(flavor, str):
            raise TypeError("Expected argument 'flavor' to be a str")
        pulumi.set(__self__, "flavor", flavor)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kafka_rest_api and not isinstance(kafka_rest_api, bool):
            raise TypeError("Expected argument 'kafka_rest_api' to be a bool")
        pulumi.set(__self__, "kafka_rest_api", kafka_rest_api)
        if maintenance_time and not isinstance(maintenance_time, str):
            raise TypeError("Expected argument 'maintenance_time' to be a str")
        pulumi.set(__self__, "maintenance_time", maintenance_time)
        if network_type and not isinstance(network_type, str):
            raise TypeError("Expected argument 'network_type' to be a str")
        pulumi.set(__self__, "network_type", network_type)
        if nodes and not isinstance(nodes, list):
            raise TypeError("Expected argument 'nodes' to be a list")
        pulumi.set(__self__, "nodes", nodes)
        if opensearch_acls_enabled and not isinstance(opensearch_acls_enabled, bool):
            raise TypeError("Expected argument 'opensearch_acls_enabled' to be a bool")
        pulumi.set(__self__, "opensearch_acls_enabled", opensearch_acls_enabled)
        if plan and not isinstance(plan, str):
            raise TypeError("Expected argument 'plan' to be a str")
        pulumi.set(__self__, "plan", plan)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="advancedConfiguration")
    def advanced_configuration(self) -> Mapping[str, str]:
        """
        Advanced configuration key / value.
        """
        return pulumi.get(self, "advanced_configuration")

    @property
    @pulumi.getter(name="backupTime")
    def backup_time(self) -> str:
        """
        Time on which backups start every day.
        """
        return pulumi.get(self, "backup_time")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Date of the creation of the cluster.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Small description of the database service.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskSize")
    def disk_size(self) -> int:
        """
        The disk size (in GB) of the database service.
        """
        return pulumi.get(self, "disk_size")

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> str:
        """
        The disk type of the database service.
        """
        return pulumi.get(self, "disk_type")

    @property
    @pulumi.getter
    def endpoints(self) -> Sequence['outputs.GetDatabaseEndpointResult']:
        """
        List of all endpoints objects of the service.
        """
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter
    def engine(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def flavor(self) -> str:
        """
        A valid OVHcloud public cloud database flavor name in which the nodes will be started.
        """
        return pulumi.get(self, "flavor")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="kafkaRestApi")
    def kafka_rest_api(self) -> bool:
        """
        Defines whether the REST API is enabled on a kafka cluster.
        """
        return pulumi.get(self, "kafka_rest_api")

    @property
    @pulumi.getter(name="maintenanceTime")
    def maintenance_time(self) -> str:
        """
        Time on which maintenances can start every day.
        """
        return pulumi.get(self, "maintenance_time")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> str:
        """
        Type of network of the cluster.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter
    def nodes(self) -> Sequence['outputs.GetDatabaseNodeResult']:
        """
        List of nodes object.
        """
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter(name="opensearchAclsEnabled")
    def opensearch_acls_enabled(self) -> bool:
        return pulumi.get(self, "opensearch_acls_enabled")

    @property
    @pulumi.getter
    def plan(self) -> str:
        """
        Plan of the cluster.
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Current status of the cluster.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The version of the engine in which the service should be deployed
        """
        return pulumi.get(self, "version")


class AwaitableGetDatabaseResult(GetDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseResult(
            advanced_configuration=self.advanced_configuration,
            backup_time=self.backup_time,
            created_at=self.created_at,
            description=self.description,
            disk_size=self.disk_size,
            disk_type=self.disk_type,
            endpoints=self.endpoints,
            engine=self.engine,
            flavor=self.flavor,
            id=self.id,
            kafka_rest_api=self.kafka_rest_api,
            maintenance_time=self.maintenance_time,
            network_type=self.network_type,
            nodes=self.nodes,
            opensearch_acls_enabled=self.opensearch_acls_enabled,
            plan=self.plan,
            service_name=self.service_name,
            status=self.status,
            version=self.version)


def get_database(engine: Optional[str] = None,
                 id: Optional[str] = None,
                 service_name: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseResult:
    """
    Use this data source to get the managed database of a public cloud project.

    ## Example Usage

    To get information of a database cluster service:

    ```python
    import pulumi
    import pulumi_ovh as ovh

    db = ovh.CloudProjectDatabase.get_database(service_name="XXXXXX",
        engine="YYYY",
        id="ZZZZ")
    pulumi.export("clusterId", db.id)
    ```


    :param str engine: The database engine you want to get information. To get a full list of available engine visit:
           [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
    :param str id: Cluster ID
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['engine'] = engine
    __args__['id'] = id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProjectDatabase/getDatabase:getDatabase', __args__, opts=opts, typ=GetDatabaseResult).value

    return AwaitableGetDatabaseResult(
        advanced_configuration=__ret__.advanced_configuration,
        backup_time=__ret__.backup_time,
        created_at=__ret__.created_at,
        description=__ret__.description,
        disk_size=__ret__.disk_size,
        disk_type=__ret__.disk_type,
        endpoints=__ret__.endpoints,
        engine=__ret__.engine,
        flavor=__ret__.flavor,
        id=__ret__.id,
        kafka_rest_api=__ret__.kafka_rest_api,
        maintenance_time=__ret__.maintenance_time,
        network_type=__ret__.network_type,
        nodes=__ret__.nodes,
        opensearch_acls_enabled=__ret__.opensearch_acls_enabled,
        plan=__ret__.plan,
        service_name=__ret__.service_name,
        status=__ret__.status,
        version=__ret__.version)


@_utilities.lift_output_func(get_database)
def get_database_output(engine: Optional[pulumi.Input[str]] = None,
                        id: Optional[pulumi.Input[str]] = None,
                        service_name: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseResult]:
    """
    Use this data source to get the managed database of a public cloud project.

    ## Example Usage

    To get information of a database cluster service:

    ```python
    import pulumi
    import pulumi_ovh as ovh

    db = ovh.CloudProjectDatabase.get_database(service_name="XXXXXX",
        engine="YYYY",
        id="ZZZZ")
    pulumi.export("clusterId", db.id)
    ```


    :param str engine: The database engine you want to get information. To get a full list of available engine visit:
           [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
    :param str id: Cluster ID
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...
