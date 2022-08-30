# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetCloudProjectDatabaseUsersResult',
    'AwaitableGetCloudProjectDatabaseUsersResult',
    'get_cloud_project_database_users',
    'get_cloud_project_database_users_output',
]

@pulumi.output_type
class GetCloudProjectDatabaseUsersResult:
    """
    A collection of values returned by getCloudProjectDatabaseUsers.
    """
    def __init__(__self__, cluster_id=None, engine=None, id=None, service_name=None, user_ids=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        pulumi.set(__self__, "engine", engine)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if user_ids and not isinstance(user_ids, list):
            raise TypeError("Expected argument 'user_ids' to be a list")
        pulumi.set(__self__, "user_ids", user_ids)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def engine(self) -> str:
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="userIds")
    def user_ids(self) -> Sequence[str]:
        """
        The list of users ids of the database cluster associated with the project.
        """
        return pulumi.get(self, "user_ids")


class AwaitableGetCloudProjectDatabaseUsersResult(GetCloudProjectDatabaseUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudProjectDatabaseUsersResult(
            cluster_id=self.cluster_id,
            engine=self.engine,
            id=self.id,
            service_name=self.service_name,
            user_ids=self.user_ids)


def get_cloud_project_database_users(cluster_id: Optional[str] = None,
                                     engine: Optional[str] = None,
                                     service_name: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudProjectDatabaseUsersResult:
    """
    Use this data source to get the list of users of a database cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    users = ovh.get_cloud_project_database_users(service_name="XXXX",
        engine="YYYY",
        cluster_id="ZZZ")
    pulumi.export("userIds", users.user_ids)
    ```


    :param str cluster_id: Cluster ID
    :param str engine: The engine of the database cluster you want to list users. To get a full list of available engine visit.
           [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['engine'] = engine
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:index/getCloudProjectDatabaseUsers:getCloudProjectDatabaseUsers', __args__, opts=opts, typ=GetCloudProjectDatabaseUsersResult).value

    return AwaitableGetCloudProjectDatabaseUsersResult(
        cluster_id=__ret__.cluster_id,
        engine=__ret__.engine,
        id=__ret__.id,
        service_name=__ret__.service_name,
        user_ids=__ret__.user_ids)


@_utilities.lift_output_func(get_cloud_project_database_users)
def get_cloud_project_database_users_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                            engine: Optional[pulumi.Input[str]] = None,
                                            service_name: Optional[pulumi.Input[str]] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCloudProjectDatabaseUsersResult]:
    """
    Use this data source to get the list of users of a database cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    users = ovh.get_cloud_project_database_users(service_name="XXXX",
        engine="YYYY",
        cluster_id="ZZZ")
    pulumi.export("userIds", users.user_ids)
    ```


    :param str cluster_id: Cluster ID
    :param str engine: The engine of the database cluster you want to list users. To get a full list of available engine visit.
           [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...