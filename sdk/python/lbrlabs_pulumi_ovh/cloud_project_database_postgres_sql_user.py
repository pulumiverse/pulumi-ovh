# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['CloudProjectDatabasePostgresSqlUserArgs', 'CloudProjectDatabasePostgresSqlUser']

@pulumi.input_type
class CloudProjectDatabasePostgresSqlUserArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a CloudProjectDatabasePostgresSqlUser resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] name: Name of the user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Roles the user belongs to. Possible values:
               * `["replication"]`
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "service_name", service_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

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
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the user.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Roles the user belongs to. Possible values:
        * `["replication"]`
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "roles", value)


@pulumi.input_type
class _CloudProjectDatabasePostgresSqlUserState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CloudProjectDatabasePostgresSqlUser resources.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] created_at: Date of the creation of the user.
        :param pulumi.Input[str] name: Name of the user.
        :param pulumi.Input[str] password: Password of the user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Roles the user belongs to. Possible values:
               * `["replication"]`
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] status: Current status of the user.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date of the creation of the user.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the user.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password of the user.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Roles the user belongs to. Possible values:
        * `["replication"]`
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "roles", value)

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
        Current status of the user.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class CloudProjectDatabasePostgresSqlUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an user for a postgresql cluster associated with a public cloud project.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh
        import pulumi_ovh as ovh

        postgresql = ovh.get_cloud_project_database(service_name="XXXX",
            engine="postgresql",
            cluster_id="ZZZZ")
        user = ovh.CloudProjectDatabasePostgresSqlUser("user",
            service_name=ovh_cloud_project_database["postgresql"]["service_name"],
            cluster_id=ovh_cloud_project_database["postgresql"]["id"],
            roles=["replication"])
        ```

        ## Import

        OVHcloud Managed postgresql clusters users can be imported using the `service_name`, `cluster_id` and `id` of the user, separated by "/" E.g.,

        ```sh
         $ pulumi import ovh:index/cloudProjectDatabasePostgresSqlUser:CloudProjectDatabasePostgresSqlUser my_user <service_name>/<cluster_id>/<id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] name: Name of the user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Roles the user belongs to. Possible values:
               * `["replication"]`
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CloudProjectDatabasePostgresSqlUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an user for a postgresql cluster associated with a public cloud project.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh
        import pulumi_ovh as ovh

        postgresql = ovh.get_cloud_project_database(service_name="XXXX",
            engine="postgresql",
            cluster_id="ZZZZ")
        user = ovh.CloudProjectDatabasePostgresSqlUser("user",
            service_name=ovh_cloud_project_database["postgresql"]["service_name"],
            cluster_id=ovh_cloud_project_database["postgresql"]["id"],
            roles=["replication"])
        ```

        ## Import

        OVHcloud Managed postgresql clusters users can be imported using the `service_name`, `cluster_id` and `id` of the user, separated by "/" E.g.,

        ```sh
         $ pulumi import ovh:index/cloudProjectDatabasePostgresSqlUser:CloudProjectDatabasePostgresSqlUser my_user <service_name>/<cluster_id>/<id>
        ```

        :param str resource_name: The name of the resource.
        :param CloudProjectDatabasePostgresSqlUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CloudProjectDatabasePostgresSqlUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CloudProjectDatabasePostgresSqlUserArgs.__new__(CloudProjectDatabasePostgresSqlUserArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["name"] = name
            __props__.__dict__["roles"] = roles
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["created_at"] = None
            __props__.__dict__["password"] = None
            __props__.__dict__["status"] = None
        super(CloudProjectDatabasePostgresSqlUser, __self__).__init__(
            'ovh:index/cloudProjectDatabasePostgresSqlUser:CloudProjectDatabasePostgresSqlUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'CloudProjectDatabasePostgresSqlUser':
        """
        Get an existing CloudProjectDatabasePostgresSqlUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] created_at: Date of the creation of the user.
        :param pulumi.Input[str] name: Name of the user.
        :param pulumi.Input[str] password: Password of the user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Roles the user belongs to. Possible values:
               * `["replication"]`
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] status: Current status of the user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CloudProjectDatabasePostgresSqlUserState.__new__(_CloudProjectDatabasePostgresSqlUserState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["roles"] = roles
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["status"] = status
        return CloudProjectDatabasePostgresSqlUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Date of the creation of the user.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the user.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Password of the user.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Roles the user belongs to. Possible values:
        * `["replication"]`
        """
        return pulumi.get(self, "roles")

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
        Current status of the user.
        """
        return pulumi.get(self, "status")

