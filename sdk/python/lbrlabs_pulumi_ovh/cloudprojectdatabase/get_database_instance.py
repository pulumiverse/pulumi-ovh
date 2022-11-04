# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDatabaseInstanceResult',
    'AwaitableGetDatabaseInstanceResult',
    'get_database_instance',
    'get_database_instance_output',
]

@pulumi.output_type
class GetDatabaseInstanceResult:
    """
    A collection of values returned by getDatabaseInstance.
    """
    def __init__(__self__, cluster_id=None, default=None, engine=None, id=None, name=None, service_name=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if default and not isinstance(default, bool):
            raise TypeError("Expected argument 'default' to be a bool")
        pulumi.set(__self__, "default", default)
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        pulumi.set(__self__, "engine", engine)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def default(self) -> bool:
        """
        Defines if the database has been created by default.
        """
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def engine(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the database.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        Current status of the database.
        """
        return pulumi.get(self, "service_name")


class AwaitableGetDatabaseInstanceResult(GetDatabaseInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseInstanceResult(
            cluster_id=self.cluster_id,
            default=self.default,
            engine=self.engine,
            id=self.id,
            name=self.name,
            service_name=self.service_name)


def get_database_instance(cluster_id: Optional[str] = None,
                          engine: Optional[str] = None,
                          name: Optional[str] = None,
                          service_name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseInstanceResult:
    """
    Use this data source to get information about a database of a database cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    database = ovh.CloudProjectDatabase.get_database_instance(service_name="XXX",
        engine="YYY",
        cluster_id="ZZZ",
        name="UUU")
    pulumi.export("databaseName", database.name)
    ```


    :param str cluster_id: Cluster ID
    :param str engine: The engine of the database cluster you want database information. To get a full list of available engine visit:
           [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
           Available engines:
           * `mysql`
           * `postgresql`
    :param str name: Name of the database.
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['engine'] = engine
    __args__['name'] = name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProjectDatabase/getDatabaseInstance:getDatabaseInstance', __args__, opts=opts, typ=GetDatabaseInstanceResult).value

    return AwaitableGetDatabaseInstanceResult(
        cluster_id=__ret__.cluster_id,
        default=__ret__.default,
        engine=__ret__.engine,
        id=__ret__.id,
        name=__ret__.name,
        service_name=__ret__.service_name)


@_utilities.lift_output_func(get_database_instance)
def get_database_instance_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                 engine: Optional[pulumi.Input[str]] = None,
                                 name: Optional[pulumi.Input[str]] = None,
                                 service_name: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseInstanceResult]:
    """
    Use this data source to get information about a database of a database cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    database = ovh.CloudProjectDatabase.get_database_instance(service_name="XXX",
        engine="YYY",
        cluster_id="ZZZ",
        name="UUU")
    pulumi.export("databaseName", database.name)
    ```


    :param str cluster_id: Cluster ID
    :param str engine: The engine of the database cluster you want database information. To get a full list of available engine visit:
           [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
           Available engines:
           * `mysql`
           * `postgresql`
    :param str name: Name of the database.
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...
