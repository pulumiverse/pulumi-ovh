// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Ovh.Outputs
{

    [OutputType]
    public sealed class GetOrderCartProductOptionsPlanSelectedPricePriceResult
    {
        /// <summary>
        /// Currency code
        /// </summary>
        public readonly string CurrencyCode;
        /// <summary>
        /// Textual representation
        /// </summary>
        public readonly string Text;
        /// <summary>
        /// The effective price
        /// </summary>
        public readonly double Value;

        [OutputConstructor]
        private GetOrderCartProductOptionsPlanSelectedPricePriceResult(
            string currencyCode,

            string text,

            double value)
        {
            CurrencyCode = currencyCode;
            Text = text;
            Value = value;
        }
    }
}
