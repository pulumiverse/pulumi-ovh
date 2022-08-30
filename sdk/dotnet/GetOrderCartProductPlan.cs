// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetOrderCartProductPlan
    {
        /// <summary>
        /// Use this data source to retrieve information of order cart product plan.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mycart = Ovh.GetOrderCart.Invoke(new()
        ///     {
        ///         OvhSubsidiary = "fr",
        ///         Description = "my cart",
        ///     });
        /// 
        ///     var plan = Ovh.GetOrderCartProductPlan.Invoke(new()
        ///     {
        ///         CartId = mycart.Apply(getOrderCartResult =&gt; getOrderCartResult.Id),
        ///         PriceCapacity = "renew",
        ///         Product = "cloud",
        ///         PlanCode = "project",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOrderCartProductPlanResult> InvokeAsync(GetOrderCartProductPlanArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrderCartProductPlanResult>("ovh:index/getOrderCartProductPlan:getOrderCartProductPlan", args ?? new GetOrderCartProductPlanArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information of order cart product plan.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mycart = Ovh.GetOrderCart.Invoke(new()
        ///     {
        ///         OvhSubsidiary = "fr",
        ///         Description = "my cart",
        ///     });
        /// 
        ///     var plan = Ovh.GetOrderCartProductPlan.Invoke(new()
        ///     {
        ///         CartId = mycart.Apply(getOrderCartResult =&gt; getOrderCartResult.Id),
        ///         PriceCapacity = "renew",
        ///         Product = "cloud",
        ///         PlanCode = "project",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetOrderCartProductPlanResult> Invoke(GetOrderCartProductPlanInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOrderCartProductPlanResult>("ovh:index/getOrderCartProductPlan:getOrderCartProductPlan", args ?? new GetOrderCartProductPlanInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrderCartProductPlanArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cart identifier
        /// </summary>
        [Input("cartId", required: true)]
        public string CartId { get; set; } = null!;

        /// <summary>
        /// Catalog name
        /// </summary>
        [Input("catalogName")]
        public string? CatalogName { get; set; }

        /// <summary>
        /// Product offer identifier
        /// </summary>
        [Input("planCode", required: true)]
        public string PlanCode { get; set; } = null!;

        /// <summary>
        /// Capacity of the pricing (type of pricing)
        /// </summary>
        [Input("priceCapacity", required: true)]
        public string PriceCapacity { get; set; } = null!;

        /// <summary>
        /// Product
        /// </summary>
        [Input("product", required: true)]
        public string Product { get; set; } = null!;

        public GetOrderCartProductPlanArgs()
        {
        }
        public static new GetOrderCartProductPlanArgs Empty => new GetOrderCartProductPlanArgs();
    }

    public sealed class GetOrderCartProductPlanInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cart identifier
        /// </summary>
        [Input("cartId", required: true)]
        public Input<string> CartId { get; set; } = null!;

        /// <summary>
        /// Catalog name
        /// </summary>
        [Input("catalogName")]
        public Input<string>? CatalogName { get; set; }

        /// <summary>
        /// Product offer identifier
        /// </summary>
        [Input("planCode", required: true)]
        public Input<string> PlanCode { get; set; } = null!;

        /// <summary>
        /// Capacity of the pricing (type of pricing)
        /// </summary>
        [Input("priceCapacity", required: true)]
        public Input<string> PriceCapacity { get; set; } = null!;

        /// <summary>
        /// Product
        /// </summary>
        [Input("product", required: true)]
        public Input<string> Product { get; set; } = null!;

        public GetOrderCartProductPlanInvokeArgs()
        {
        }
        public static new GetOrderCartProductPlanInvokeArgs Empty => new GetOrderCartProductPlanInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrderCartProductPlanResult
    {
        public readonly string CartId;
        public readonly string? CatalogName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Product offer identifier
        /// </summary>
        public readonly string PlanCode;
        public readonly string PriceCapacity;
        /// <summary>
        /// Prices of the product offer
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrderCartProductPlanPriceResult> Prices;
        public readonly string Product;
        /// <summary>
        /// Name of the product
        /// </summary>
        public readonly string ProductName;
        /// <summary>
        /// Product type
        /// </summary>
        public readonly string ProductType;
        /// <summary>
        /// Selected Price according to capacity
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrderCartProductPlanSelectedPriceResult> SelectedPrices;

        [OutputConstructor]
        private GetOrderCartProductPlanResult(
            string cartId,

            string? catalogName,

            string id,

            string planCode,

            string priceCapacity,

            ImmutableArray<Outputs.GetOrderCartProductPlanPriceResult> prices,

            string product,

            string productName,

            string productType,

            ImmutableArray<Outputs.GetOrderCartProductPlanSelectedPriceResult> selectedPrices)
        {
            CartId = cartId;
            CatalogName = catalogName;
            Id = id;
            PlanCode = planCode;
            PriceCapacity = priceCapacity;
            Prices = prices;
            Product = product;
            ProductName = productName;
            ProductType = productType;
            SelectedPrices = selectedPrices;
        }
    }
}
