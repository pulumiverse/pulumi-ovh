// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Inputs
{

    public sealed class DedicatedServiceInstallTaskDetailsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Template change log details.
        /// </summary>
        [Input("changeLog")]
        public Input<string>? ChangeLog { get; set; }

        /// <summary>
        /// Set up the server using the provided hostname instead of the default hostname.
        /// </summary>
        [Input("customHostname")]
        public Input<string>? CustomHostname { get; set; }

        /// <summary>
        /// Disk group id.
        /// </summary>
        [Input("diskGroupId")]
        public Input<int>? DiskGroupId { get; set; }

        /// <summary>
        /// set to true to install RTM.
        /// </summary>
        [Input("installRtm")]
        public Input<bool>? InstallRtm { get; set; }

        /// <summary>
        /// set to true to install sql server (Windows template only).
        /// </summary>
        [Input("installSqlServer")]
        public Input<bool>? InstallSqlServer { get; set; }

        /// <summary>
        /// language.
        /// </summary>
        [Input("language")]
        public Input<string>? Language { get; set; }

        /// <summary>
        /// set to true to disable RAID.
        /// </summary>
        [Input("noRaid")]
        public Input<bool>? NoRaid { get; set; }

        /// <summary>
        /// Indicate the URL where your postinstall customisation script is located.
        /// </summary>
        [Input("postInstallationScriptLink")]
        public Input<string>? PostInstallationScriptLink { get; set; }

        /// <summary>
        /// Indicate the string returned by your postinstall customisation script on successful execution. Advice: your script should return a unique validation string in case of succes. A good example is 'loh1Xee7eo OK OK OK UGh8Ang1Gu'.
        /// </summary>
        [Input("postInstallationScriptReturn")]
        public Input<string>? PostInstallationScriptReturn { get; set; }

        /// <summary>
        /// set to true to make a hardware raid reset.
        /// </summary>
        [Input("resetHwRaid")]
        public Input<bool>? ResetHwRaid { get; set; }

        /// <summary>
        /// soft raid devices.
        /// </summary>
        [Input("softRaidDevices")]
        public Input<int>? SoftRaidDevices { get; set; }

        /// <summary>
        /// Name of the ssh key that should be installed. Password login will be disabled.
        /// </summary>
        [Input("sshKeyName")]
        public Input<string>? SshKeyName { get; set; }

        /// <summary>
        /// Use the distribution's native kernel instead of the recommended OVH Kernel.
        /// </summary>
        [Input("useDistribKernel")]
        public Input<bool>? UseDistribKernel { get; set; }

        /// <summary>
        /// set to true to use SPLA.
        /// </summary>
        [Input("useSpla")]
        public Input<bool>? UseSpla { get; set; }

        public DedicatedServiceInstallTaskDetailsGetArgs()
        {
        }
        public static new DedicatedServiceInstallTaskDetailsGetArgs Empty => new DedicatedServiceInstallTaskDetailsGetArgs();
    }
}
