// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Inputs
{

    public sealed class IpServicePlanOptionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Catalog name
        /// </summary>
        [Input("catalogName")]
        public Input<string>? CatalogName { get; set; }

        [Input("configurations")]
        private InputList<Inputs.IpServicePlanOptionConfigurationGetArgs>? _configurations;

        /// <summary>
        /// Representation of a configuration item for personalizing product
        /// </summary>
        public InputList<Inputs.IpServicePlanOptionConfigurationGetArgs> Configurations
        {
            get => _configurations ?? (_configurations = new InputList<Inputs.IpServicePlanOptionConfigurationGetArgs>());
            set => _configurations = value;
        }

        /// <summary>
        /// duration
        /// </summary>
        [Input("duration", required: true)]
        public Input<string> Duration { get; set; } = null!;

        /// <summary>
        /// Plan code
        /// </summary>
        [Input("planCode", required: true)]
        public Input<string> PlanCode { get; set; } = null!;

        /// <summary>
        /// Pricing model identifier
        /// </summary>
        [Input("pricingMode", required: true)]
        public Input<string> PricingMode { get; set; } = null!;

        public IpServicePlanOptionGetArgs()
        {
        }
        public static new IpServicePlanOptionGetArgs Empty => new IpServicePlanOptionGetArgs();
    }
}
