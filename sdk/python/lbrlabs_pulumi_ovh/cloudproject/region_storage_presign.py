# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RegionStoragePresignArgs', 'RegionStoragePresign']

@pulumi.input_type
class RegionStoragePresignArgs:
    def __init__(__self__, *,
                 expire: pulumi.Input[int],
                 method: pulumi.Input[str],
                 object: pulumi.Input[str],
                 region_name: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RegionStoragePresign resource.
        :param pulumi.Input[int] expire: Define, in seconds, for how long your URL will be valid.
        :param pulumi.Input[str] method: The method you want to use to interact with your object. Can be either 'GET' or 'PUT'.
        :param pulumi.Input[str] object: The name of the object in your S3 bucket.
        :param pulumi.Input[str] region_name: The region in which your storage is located.
               Ex.: "GRA".
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] name: The name of your S3 storage container/bucket.
        """
        pulumi.set(__self__, "expire", expire)
        pulumi.set(__self__, "method", method)
        pulumi.set(__self__, "object", object)
        pulumi.set(__self__, "region_name", region_name)
        pulumi.set(__self__, "service_name", service_name)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def expire(self) -> pulumi.Input[int]:
        """
        Define, in seconds, for how long your URL will be valid.
        """
        return pulumi.get(self, "expire")

    @expire.setter
    def expire(self, value: pulumi.Input[int]):
        pulumi.set(self, "expire", value)

    @property
    @pulumi.getter
    def method(self) -> pulumi.Input[str]:
        """
        The method you want to use to interact with your object. Can be either 'GET' or 'PUT'.
        """
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: pulumi.Input[str]):
        pulumi.set(self, "method", value)

    @property
    @pulumi.getter
    def object(self) -> pulumi.Input[str]:
        """
        The name of the object in your S3 bucket.
        """
        return pulumi.get(self, "object")

    @object.setter
    def object(self, value: pulumi.Input[str]):
        pulumi.set(self, "object", value)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> pulumi.Input[str]:
        """
        The region in which your storage is located.
        Ex.: "GRA".
        """
        return pulumi.get(self, "region_name")

    @region_name.setter
    def region_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "region_name", value)

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
        The name of your S3 storage container/bucket.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _RegionStoragePresignState:
    def __init__(__self__, *,
                 expire: Optional[pulumi.Input[int]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object: Optional[pulumi.Input[str]] = None,
                 region_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RegionStoragePresign resources.
        :param pulumi.Input[int] expire: Define, in seconds, for how long your URL will be valid.
        :param pulumi.Input[str] method: The method you want to use to interact with your object. Can be either 'GET' or 'PUT'.
        :param pulumi.Input[str] name: The name of your S3 storage container/bucket.
        :param pulumi.Input[str] object: The name of the object in your S3 bucket.
        :param pulumi.Input[str] region_name: The region in which your storage is located.
               Ex.: "GRA".
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] url: Computed URL result.
        """
        if expire is not None:
            pulumi.set(__self__, "expire", expire)
        if method is not None:
            pulumi.set(__self__, "method", method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if object is not None:
            pulumi.set(__self__, "object", object)
        if region_name is not None:
            pulumi.set(__self__, "region_name", region_name)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def expire(self) -> Optional[pulumi.Input[int]]:
        """
        Define, in seconds, for how long your URL will be valid.
        """
        return pulumi.get(self, "expire")

    @expire.setter
    def expire(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expire", value)

    @property
    @pulumi.getter
    def method(self) -> Optional[pulumi.Input[str]]:
        """
        The method you want to use to interact with your object. Can be either 'GET' or 'PUT'.
        """
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "method", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of your S3 storage container/bucket.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def object(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the object in your S3 bucket.
        """
        return pulumi.get(self, "object")

    @object.setter
    def object(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object", value)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which your storage is located.
        Ex.: "GRA".
        """
        return pulumi.get(self, "region_name")

    @region_name.setter
    def region_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_name", value)

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
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        Computed URL result.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class RegionStoragePresign(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expire: Optional[pulumi.Input[int]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object: Optional[pulumi.Input[str]] = None,
                 region_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Generates a temporary presigned S3 URLs to download or upload an object.

        > __NOTE__ This resource is only compatible with the `High Performance - S3` solution for object storage.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh

        presigned_url_region_storage_presign = ovh.cloud_project.RegionStoragePresign("presignedUrlRegionStoragePresign",
            service_name="xxxxxxxxxxxxxxxxx",
            region_name="GRA",
            expire=3600,
            method="GET",
            object="an-object-in-the-bucket")
        pulumi.export("presignedUrl", presigned_url_region_storage_presign.url)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] expire: Define, in seconds, for how long your URL will be valid.
        :param pulumi.Input[str] method: The method you want to use to interact with your object. Can be either 'GET' or 'PUT'.
        :param pulumi.Input[str] name: The name of your S3 storage container/bucket.
        :param pulumi.Input[str] object: The name of the object in your S3 bucket.
        :param pulumi.Input[str] region_name: The region in which your storage is located.
               Ex.: "GRA".
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegionStoragePresignArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Generates a temporary presigned S3 URLs to download or upload an object.

        > __NOTE__ This resource is only compatible with the `High Performance - S3` solution for object storage.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh

        presigned_url_region_storage_presign = ovh.cloud_project.RegionStoragePresign("presignedUrlRegionStoragePresign",
            service_name="xxxxxxxxxxxxxxxxx",
            region_name="GRA",
            expire=3600,
            method="GET",
            object="an-object-in-the-bucket")
        pulumi.export("presignedUrl", presigned_url_region_storage_presign.url)
        ```

        :param str resource_name: The name of the resource.
        :param RegionStoragePresignArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegionStoragePresignArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expire: Optional[pulumi.Input[int]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object: Optional[pulumi.Input[str]] = None,
                 region_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegionStoragePresignArgs.__new__(RegionStoragePresignArgs)

            if expire is None and not opts.urn:
                raise TypeError("Missing required property 'expire'")
            __props__.__dict__["expire"] = expire
            if method is None and not opts.urn:
                raise TypeError("Missing required property 'method'")
            __props__.__dict__["method"] = method
            __props__.__dict__["name"] = name
            if object is None and not opts.urn:
                raise TypeError("Missing required property 'object'")
            __props__.__dict__["object"] = object
            if region_name is None and not opts.urn:
                raise TypeError("Missing required property 'region_name'")
            __props__.__dict__["region_name"] = region_name
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["url"] = None
        super(RegionStoragePresign, __self__).__init__(
            'ovh:CloudProject/regionStoragePresign:RegionStoragePresign',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            expire: Optional[pulumi.Input[int]] = None,
            method: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            object: Optional[pulumi.Input[str]] = None,
            region_name: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'RegionStoragePresign':
        """
        Get an existing RegionStoragePresign resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] expire: Define, in seconds, for how long your URL will be valid.
        :param pulumi.Input[str] method: The method you want to use to interact with your object. Can be either 'GET' or 'PUT'.
        :param pulumi.Input[str] name: The name of your S3 storage container/bucket.
        :param pulumi.Input[str] object: The name of the object in your S3 bucket.
        :param pulumi.Input[str] region_name: The region in which your storage is located.
               Ex.: "GRA".
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] url: Computed URL result.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RegionStoragePresignState.__new__(_RegionStoragePresignState)

        __props__.__dict__["expire"] = expire
        __props__.__dict__["method"] = method
        __props__.__dict__["name"] = name
        __props__.__dict__["object"] = object
        __props__.__dict__["region_name"] = region_name
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["url"] = url
        return RegionStoragePresign(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def expire(self) -> pulumi.Output[int]:
        """
        Define, in seconds, for how long your URL will be valid.
        """
        return pulumi.get(self, "expire")

    @property
    @pulumi.getter
    def method(self) -> pulumi.Output[str]:
        """
        The method you want to use to interact with your object. Can be either 'GET' or 'PUT'.
        """
        return pulumi.get(self, "method")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of your S3 storage container/bucket.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def object(self) -> pulumi.Output[str]:
        """
        The name of the object in your S3 bucket.
        """
        return pulumi.get(self, "object")

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> pulumi.Output[str]:
        """
        The region in which your storage is located.
        Ex.: "GRA".
        """
        return pulumi.get(self, "region_name")

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
    def url(self) -> pulumi.Output[str]:
        """
        Computed URL result.
        """
        return pulumi.get(self, "url")

