// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.Order
{
    public static class GetCartProduct
    {
        /// <summary>
        /// Use this data source to retrieve information of order cart product products.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mycart = Ovh.Order.GetCart.Invoke(new()
        ///     {
        ///         OvhSubsidiary = "fr",
        ///         Description = "my cart",
        ///     });
        /// 
        ///     var plans = Ovh.Order.GetCartProduct.Invoke(new()
        ///     {
        ///         CartId = mycart.Apply(getCartResult =&gt; getCartResult.Id),
        ///         Product = "...",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCartProductResult> InvokeAsync(GetCartProductArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCartProductResult>("ovh:Order/getCartProduct:getCartProduct", args ?? new GetCartProductArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information of order cart product products.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mycart = Ovh.Order.GetCart.Invoke(new()
        ///     {
        ///         OvhSubsidiary = "fr",
        ///         Description = "my cart",
        ///     });
        /// 
        ///     var plans = Ovh.Order.GetCartProduct.Invoke(new()
        ///     {
        ///         CartId = mycart.Apply(getCartResult =&gt; getCartResult.Id),
        ///         Product = "...",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCartProductResult> Invoke(GetCartProductInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCartProductResult>("ovh:Order/getCartProduct:getCartProduct", args ?? new GetCartProductInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCartProductArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cart identifier
        /// </summary>
        [Input("cartId", required: true)]
        public string CartId { get; set; } = null!;

        /// <summary>
        /// product
        /// </summary>
        [Input("product", required: true)]
        public string Product { get; set; } = null!;

        public GetCartProductArgs()
        {
        }
        public static new GetCartProductArgs Empty => new GetCartProductArgs();
    }

    public sealed class GetCartProductInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cart identifier
        /// </summary>
        [Input("cartId", required: true)]
        public Input<string> CartId { get; set; } = null!;

        /// <summary>
        /// product
        /// </summary>
        [Input("product", required: true)]
        public Input<string> Product { get; set; } = null!;

        public GetCartProductInvokeArgs()
        {
        }
        public static new GetCartProductInvokeArgs Empty => new GetCartProductInvokeArgs();
    }


    [OutputType]
    public sealed class GetCartProductResult
    {
        public readonly string CartId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Product;
        /// <summary>
        /// products results
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCartProductResultResult> Results;

        [OutputConstructor]
        private GetCartProductResult(
            string cartId,

            string id,

            string product,

            ImmutableArray<Outputs.GetCartProductResultResult> results)
        {
            CartId = cartId;
            Id = id;
            Product = product;
            Results = results;
        }
    }
}
