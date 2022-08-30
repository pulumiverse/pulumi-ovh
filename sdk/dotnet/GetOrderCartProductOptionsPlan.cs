// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetOrderCartProductOptionsPlan
    {
        /// <summary>
        /// Use this data source to retrieve information of order cart product options plan.
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
        ///     var plan = Ovh.GetOrderCartProductOptionsPlan.Invoke(new()
        ///     {
        ///         CartId = mycart.Apply(getOrderCartResult =&gt; getOrderCartResult.Id),
        ///         PriceCapacity = "renew",
        ///         Product = "cloud",
        ///         PlanCode = "project",
        ///         OptionsPlanCode = "vrack",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOrderCartProductOptionsPlanResult> InvokeAsync(GetOrderCartProductOptionsPlanArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrderCartProductOptionsPlanResult>("ovh:index/getOrderCartProductOptionsPlan:getOrderCartProductOptionsPlan", args ?? new GetOrderCartProductOptionsPlanArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information of order cart product options plan.
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
        ///     var plan = Ovh.GetOrderCartProductOptionsPlan.Invoke(new()
        ///     {
        ///         CartId = mycart.Apply(getOrderCartResult =&gt; getOrderCartResult.Id),
        ///         PriceCapacity = "renew",
        ///         Product = "cloud",
        ///         PlanCode = "project",
        ///         OptionsPlanCode = "vrack",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetOrderCartProductOptionsPlanResult> Invoke(GetOrderCartProductOptionsPlanInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOrderCartProductOptionsPlanResult>("ovh:index/getOrderCartProductOptionsPlan:getOrderCartProductOptionsPlan", args ?? new GetOrderCartProductOptionsPlanInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrderCartProductOptionsPlanArgs : global::Pulumi.InvokeArgs
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
        /// options plan code.
        /// </summary>
        [Input("optionsPlanCode", required: true)]
        public string OptionsPlanCode { get; set; } = null!;

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

        public GetOrderCartProductOptionsPlanArgs()
        {
        }
        public static new GetOrderCartProductOptionsPlanArgs Empty => new GetOrderCartProductOptionsPlanArgs();
    }

    public sealed class GetOrderCartProductOptionsPlanInvokeArgs : global::Pulumi.InvokeArgs
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
        /// options plan code.
        /// </summary>
        [Input("optionsPlanCode", required: true)]
        public Input<string> OptionsPlanCode { get; set; } = null!;

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

        public GetOrderCartProductOptionsPlanInvokeArgs()
        {
        }
        public static new GetOrderCartProductOptionsPlanInvokeArgs Empty => new GetOrderCartProductOptionsPlanInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrderCartProductOptionsPlanResult
    {
        public readonly string CartId;
        public readonly string? CatalogName;
        /// <summary>
        /// Define if options of this family are exclusive with each other
        /// </summary>
        public readonly bool Exclusive;
        /// <summary>
        /// Option family
        /// </summary>
        public readonly string Family;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Define if an option of this family is mandatory
        /// </summary>
        public readonly bool Mandatory;
        public readonly string OptionsPlanCode;
        /// <summary>
        /// Product offer identifier
        /// </summary>
        public readonly string PlanCode;
        public readonly string PriceCapacity;
        /// <summary>
        /// Prices of the product offer
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrderCartProductOptionsPlanPriceResult> Prices;
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
        public readonly ImmutableArray<Outputs.GetOrderCartProductOptionsPlanSelectedPriceResult> SelectedPrices;

        [OutputConstructor]
        private GetOrderCartProductOptionsPlanResult(
            string cartId,

            string? catalogName,

            bool exclusive,

            string family,

            string id,

            bool mandatory,

            string optionsPlanCode,

            string planCode,

            string priceCapacity,

            ImmutableArray<Outputs.GetOrderCartProductOptionsPlanPriceResult> prices,

            string product,

            string productName,

            string productType,

            ImmutableArray<Outputs.GetOrderCartProductOptionsPlanSelectedPriceResult> selectedPrices)
        {
            CartId = cartId;
            CatalogName = catalogName;
            Exclusive = exclusive;
            Family = family;
            Id = id;
            Mandatory = mandatory;
            OptionsPlanCode = optionsPlanCode;
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
