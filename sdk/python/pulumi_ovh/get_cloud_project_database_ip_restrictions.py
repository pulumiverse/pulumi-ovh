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
    'GetCloudProjectDatabaseIpRestrictionsResult',
    'AwaitableGetCloudProjectDatabaseIpRestrictionsResult',
    'get_cloud_project_database_ip_restrictions',
    'get_cloud_project_database_ip_restrictions_output',
]

@pulumi.output_type
class GetCloudProjectDatabaseIpRestrictionsResult:
    """
    A collection of values returned by getCloudProjectDatabaseIpRestrictions.
    """
    def __init__(__self__, cluster_id=None, engine=None, id=None, ips=None, service_name=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        pulumi.set(__self__, "engine", engine)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ips and not isinstance(ips, list):
            raise TypeError("Expected argument 'ips' to be a list")
        pulumi.set(__self__, "ips", ips)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

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
    @pulumi.getter
    def ips(self) -> Sequence[str]:
        """
        The list of IP restriction of the database associated with the project.
        """
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")


class AwaitableGetCloudProjectDatabaseIpRestrictionsResult(GetCloudProjectDatabaseIpRestrictionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudProjectDatabaseIpRestrictionsResult(
            cluster_id=self.cluster_id,
            engine=self.engine,
            id=self.id,
            ips=self.ips,
            service_name=self.service_name)


def get_cloud_project_database_ip_restrictions(cluster_id: Optional[str] = None,
                                               engine: Optional[str] = None,
                                               service_name: Optional[str] = None,
                                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudProjectDatabaseIpRestrictionsResult:
    """
    Use the list of IP restrictions associated with a public cloud project.

    ## Example Usage

    To get the list of IP restriction on a database cluster service:
    ```python
    import pulumi
    import pulumi_ovh as ovh

    iprestrictions = ovh.get_cloud_project_database_ip_restrictions(service_name="XXXXXX",
        engine="YYYY",
        cluster_id="ZZZZ")
    pulumi.export("ips", iprestrictions.ips)
    ```


    :param str cluster_id: Cluster ID
    :param str engine: The engine of the database cluster you want to list IP restrictions. To get a full list of available engine visit.
           [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['engine'] = engine
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:index/getCloudProjectDatabaseIpRestrictions:getCloudProjectDatabaseIpRestrictions', __args__, opts=opts, typ=GetCloudProjectDatabaseIpRestrictionsResult).value

    return AwaitableGetCloudProjectDatabaseIpRestrictionsResult(
        cluster_id=__ret__.cluster_id,
        engine=__ret__.engine,
        id=__ret__.id,
        ips=__ret__.ips,
        service_name=__ret__.service_name)


@_utilities.lift_output_func(get_cloud_project_database_ip_restrictions)
def get_cloud_project_database_ip_restrictions_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                                      engine: Optional[pulumi.Input[str]] = None,
                                                      service_name: Optional[pulumi.Input[str]] = None,
                                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCloudProjectDatabaseIpRestrictionsResult]:
    """
    Use the list of IP restrictions associated with a public cloud project.

    ## Example Usage

    To get the list of IP restriction on a database cluster service:
    ```python
    import pulumi
    import pulumi_ovh as ovh

    iprestrictions = ovh.get_cloud_project_database_ip_restrictions(service_name="XXXXXX",
        engine="YYYY",
        cluster_id="ZZZZ")
    pulumi.export("ips", iprestrictions.ips)
    ```


    :param str cluster_id: Cluster ID
    :param str engine: The engine of the database cluster you want to list IP restrictions. To get a full list of available engine visit.
           [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...
