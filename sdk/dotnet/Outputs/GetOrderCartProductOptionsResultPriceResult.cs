// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace lbrlabs.Ovh.Outputs
{

    [OutputType]
    public sealed class GetOrderCartProductOptionsResultPriceResult
    {
        /// <summary>
        /// Capacities of the pricing (type of pricing)
        /// </summary>
        public readonly ImmutableArray<object> Capacities;
        /// <summary>
        /// Description of the pricing
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Duration for ordering the product
        /// </summary>
        public readonly string Duration;
        /// <summary>
        /// Interval of renewal
        /// </summary>
        public readonly int Interval;
        /// <summary>
        /// Maximum quantity that can be ordered
        /// </summary>
        public readonly int MaximumQuantity;
        /// <summary>
        /// Maximum repeat for renewal
        /// </summary>
        public readonly int MaximumRepeat;
        /// <summary>
        /// Minimum quantity that can be ordered
        /// </summary>
        public readonly int MinimumQuantity;
        /// <summary>
        /// Minimum repeat for renewal
        /// </summary>
        public readonly int MinimumRepeat;
        /// <summary>
        /// Price of the product in micro-centims
        /// </summary>
        public readonly int PriceInUcents;
        /// <summary>
        /// Price of the product (Price with its currency and textual representation)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrderCartProductOptionsResultPricePriceResult> Prices;
        /// <summary>
        /// Pricing model identifier
        /// </summary>
        public readonly string PricingMode;
        /// <summary>
        /// Pricing type
        /// </summary>
        public readonly string PricingType;

        [OutputConstructor]
        private GetOrderCartProductOptionsResultPriceResult(
            ImmutableArray<object> capacities,

            string description,

            string duration,

            int interval,

            int maximumQuantity,

            int maximumRepeat,

            int minimumQuantity,

            int minimumRepeat,

            int priceInUcents,

            ImmutableArray<Outputs.GetOrderCartProductOptionsResultPricePriceResult> prices,

            string pricingMode,

            string pricingType)
        {
            Capacities = capacities;
            Description = description;
            Duration = duration;
            Interval = interval;
            MaximumQuantity = maximumQuantity;
            MaximumRepeat = maximumRepeat;
            MinimumQuantity = minimumQuantity;
            MinimumRepeat = minimumRepeat;
            PriceInUcents = priceInUcents;
            Prices = prices;
            PricingMode = pricingMode;
            PricingType = pricingType;
        }
    }
}
