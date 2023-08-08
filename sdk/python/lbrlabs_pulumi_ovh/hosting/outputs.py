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
    'PrivateDatabaseOrder',
    'PrivateDatabaseOrderDetail',
    'PrivateDatabasePlan',
    'PrivateDatabasePlanConfiguration',
    'PrivateDatabasePlanOption',
    'PrivateDatabasePlanOptionConfiguration',
    'GetPrivateDatabaseDbUserResult',
    'GetPrivateDatabaseUserDatabaseResult',
]

@pulumi.output_type
class PrivateDatabaseOrder(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "expirationDate":
            suggest = "expiration_date"
        elif key == "orderId":
            suggest = "order_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PrivateDatabaseOrder. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PrivateDatabaseOrder.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PrivateDatabaseOrder.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 date: Optional[str] = None,
                 details: Optional[Sequence['outputs.PrivateDatabaseOrderDetail']] = None,
                 expiration_date: Optional[str] = None,
                 order_id: Optional[int] = None):
        """
        :param str date: date
        :param Sequence['PrivateDatabaseOrderDetailArgs'] details: Information about a Bill entry
        :param str expiration_date: expiration date
        :param int order_id: order id
        """
        if date is not None:
            pulumi.set(__self__, "date", date)
        if details is not None:
            pulumi.set(__self__, "details", details)
        if expiration_date is not None:
            pulumi.set(__self__, "expiration_date", expiration_date)
        if order_id is not None:
            pulumi.set(__self__, "order_id", order_id)

    @property
    @pulumi.getter
    def date(self) -> Optional[str]:
        """
        date
        """
        return pulumi.get(self, "date")

    @property
    @pulumi.getter
    def details(self) -> Optional[Sequence['outputs.PrivateDatabaseOrderDetail']]:
        """
        Information about a Bill entry
        """
        return pulumi.get(self, "details")

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[str]:
        """
        expiration date
        """
        return pulumi.get(self, "expiration_date")

    @property
    @pulumi.getter(name="orderId")
    def order_id(self) -> Optional[int]:
        """
        order id
        """
        return pulumi.get(self, "order_id")


@pulumi.output_type
class PrivateDatabaseOrderDetail(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "orderDetailId":
            suggest = "order_detail_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PrivateDatabaseOrderDetail. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PrivateDatabaseOrderDetail.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PrivateDatabaseOrderDetail.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 description: Optional[str] = None,
                 domain: Optional[str] = None,
                 order_detail_id: Optional[int] = None,
                 quantity: Optional[str] = None):
        """
        :param str description: Custom description on your privatedatabase order.
        :param str domain: expiration date
        :param int order_detail_id: order detail id
        :param str quantity: quantity
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if order_detail_id is not None:
            pulumi.set(__self__, "order_detail_id", order_detail_id)
        if quantity is not None:
            pulumi.set(__self__, "quantity", quantity)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Custom description on your privatedatabase order.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def domain(self) -> Optional[str]:
        """
        expiration date
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="orderDetailId")
    def order_detail_id(self) -> Optional[int]:
        """
        order detail id
        """
        return pulumi.get(self, "order_detail_id")

    @property
    @pulumi.getter
    def quantity(self) -> Optional[str]:
        """
        quantity
        """
        return pulumi.get(self, "quantity")


@pulumi.output_type
class PrivateDatabasePlan(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "planCode":
            suggest = "plan_code"
        elif key == "pricingMode":
            suggest = "pricing_mode"
        elif key == "catalogName":
            suggest = "catalog_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PrivateDatabasePlan. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PrivateDatabasePlan.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PrivateDatabasePlan.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 duration: str,
                 plan_code: str,
                 pricing_mode: str,
                 catalog_name: Optional[str] = None,
                 configurations: Optional[Sequence['outputs.PrivateDatabasePlanConfiguration']] = None):
        """
        :param str duration: duration.
        :param str plan_code: Plan code.
        :param str pricing_mode: Pricing model identifier
        :param str catalog_name: Catalog name
        :param Sequence['PrivateDatabasePlanConfigurationArgs'] configurations: Representation of a configuration item for personalizing product
        """
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "plan_code", plan_code)
        pulumi.set(__self__, "pricing_mode", pricing_mode)
        if catalog_name is not None:
            pulumi.set(__self__, "catalog_name", catalog_name)
        if configurations is not None:
            pulumi.set(__self__, "configurations", configurations)

    @property
    @pulumi.getter
    def duration(self) -> str:
        """
        duration.
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter(name="planCode")
    def plan_code(self) -> str:
        """
        Plan code.
        """
        return pulumi.get(self, "plan_code")

    @property
    @pulumi.getter(name="pricingMode")
    def pricing_mode(self) -> str:
        """
        Pricing model identifier
        """
        return pulumi.get(self, "pricing_mode")

    @property
    @pulumi.getter(name="catalogName")
    def catalog_name(self) -> Optional[str]:
        """
        Catalog name
        """
        return pulumi.get(self, "catalog_name")

    @property
    @pulumi.getter
    def configurations(self) -> Optional[Sequence['outputs.PrivateDatabasePlanConfiguration']]:
        """
        Representation of a configuration item for personalizing product
        """
        return pulumi.get(self, "configurations")


@pulumi.output_type
class PrivateDatabasePlanConfiguration(dict):
    def __init__(__self__, *,
                 label: str,
                 value: str):
        """
        :param str label: Identifier of the resource
        :param str value: Path to the resource in API.OVH.COM
               
               Plan order valid values can be found on OVHcloud [APIv6](https://api.ovh.com/console/#/hosting/privateDatabase/availableOrderCapacities~GET)
        """
        pulumi.set(__self__, "label", label)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def label(self) -> str:
        """
        Identifier of the resource
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Path to the resource in API.OVH.COM

        Plan order valid values can be found on OVHcloud [APIv6](https://api.ovh.com/console/#/hosting/privateDatabase/availableOrderCapacities~GET)
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class PrivateDatabasePlanOption(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "planCode":
            suggest = "plan_code"
        elif key == "pricingMode":
            suggest = "pricing_mode"
        elif key == "catalogName":
            suggest = "catalog_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PrivateDatabasePlanOption. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PrivateDatabasePlanOption.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PrivateDatabasePlanOption.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 duration: str,
                 plan_code: str,
                 pricing_mode: str,
                 catalog_name: Optional[str] = None,
                 configurations: Optional[Sequence['outputs.PrivateDatabasePlanOptionConfiguration']] = None):
        """
        :param str duration: duration.
        :param str plan_code: Plan code.
        :param str pricing_mode: Pricing model identifier
        :param str catalog_name: Catalog name
        :param Sequence['PrivateDatabasePlanOptionConfigurationArgs'] configurations: Representation of a configuration item for personalizing product
        """
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "plan_code", plan_code)
        pulumi.set(__self__, "pricing_mode", pricing_mode)
        if catalog_name is not None:
            pulumi.set(__self__, "catalog_name", catalog_name)
        if configurations is not None:
            pulumi.set(__self__, "configurations", configurations)

    @property
    @pulumi.getter
    def duration(self) -> str:
        """
        duration.
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter(name="planCode")
    def plan_code(self) -> str:
        """
        Plan code.
        """
        return pulumi.get(self, "plan_code")

    @property
    @pulumi.getter(name="pricingMode")
    def pricing_mode(self) -> str:
        """
        Pricing model identifier
        """
        return pulumi.get(self, "pricing_mode")

    @property
    @pulumi.getter(name="catalogName")
    def catalog_name(self) -> Optional[str]:
        """
        Catalog name
        """
        return pulumi.get(self, "catalog_name")

    @property
    @pulumi.getter
    def configurations(self) -> Optional[Sequence['outputs.PrivateDatabasePlanOptionConfiguration']]:
        """
        Representation of a configuration item for personalizing product
        """
        return pulumi.get(self, "configurations")


@pulumi.output_type
class PrivateDatabasePlanOptionConfiguration(dict):
    def __init__(__self__, *,
                 label: str,
                 value: str):
        """
        :param str label: Identifier of the resource
        :param str value: Path to the resource in API.OVH.COM
               
               Plan order valid values can be found on OVHcloud [APIv6](https://api.ovh.com/console/#/hosting/privateDatabase/availableOrderCapacities~GET)
        """
        pulumi.set(__self__, "label", label)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def label(self) -> str:
        """
        Identifier of the resource
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Path to the resource in API.OVH.COM

        Plan order valid values can be found on OVHcloud [APIv6](https://api.ovh.com/console/#/hosting/privateDatabase/availableOrderCapacities~GET)
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class GetPrivateDatabaseDbUserResult(dict):
    def __init__(__self__, *,
                 grant_type: str,
                 user_name: str):
        """
        :param str grant_type: Grant of this user for this database
        :param str user_name: User's name granted on this database
        """
        pulumi.set(__self__, "grant_type", grant_type)
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="grantType")
    def grant_type(self) -> str:
        """
        Grant of this user for this database
        """
        return pulumi.get(self, "grant_type")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        User's name granted on this database
        """
        return pulumi.get(self, "user_name")


@pulumi.output_type
class GetPrivateDatabaseUserDatabaseResult(dict):
    def __init__(__self__, *,
                 database_name: str,
                 grant_type: str):
        """
        :param str database_name: Database's name linked to this user
        :param str grant_type: Grant of this user for this database
        """
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "grant_type", grant_type)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> str:
        """
        Database's name linked to this user
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="grantType")
    def grant_type(self) -> str:
        """
        Grant of this user for this database
        """
        return pulumi.get(self, "grant_type")


