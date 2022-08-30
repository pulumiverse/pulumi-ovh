// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Outputs
{

    [OutputType]
    public sealed class CloudProjectUserRole
    {
        /// <summary>
        /// A description associated with the user.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// id of the role
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// name of the role
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// list of permissions associated with the role
        /// </summary>
        public readonly ImmutableArray<string> Permissions;

        [OutputConstructor]
        private CloudProjectUserRole(
            string? description,

            string? id,

            string? name,

            ImmutableArray<string> permissions)
        {
            Description = description;
            Id = id;
            Name = name;
            Permissions = permissions;
        }
    }
}
