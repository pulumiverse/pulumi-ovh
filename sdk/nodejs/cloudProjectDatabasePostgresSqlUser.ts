// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates an user for a postgresql cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@lbrlabs/pulumi-ovh";
 * import * as ovh from "@pulumi/ovh";
 *
 * const postgresql = ovh.getCloudProjectDatabase({
 *     serviceName: "XXXX",
 *     engine: "postgresql",
 *     clusterId: "ZZZZ",
 * });
 * const user = new ovh.CloudProjectDatabasePostgresSqlUser("user", {
 *     serviceName: ovh_cloud_project_database.postgresql.service_name,
 *     clusterId: ovh_cloud_project_database.postgresql.id,
 *     roles: ["replication"],
 * });
 * ```
 *
 * ## Import
 *
 * OVHcloud Managed postgresql clusters users can be imported using the `service_name`, `cluster_id` and `id` of the user, separated by "/" E.g.,
 *
 * ```sh
 *  $ pulumi import ovh:index/cloudProjectDatabasePostgresSqlUser:CloudProjectDatabasePostgresSqlUser my_user <service_name>/<cluster_id>/<id>
 * ```
 */
export class CloudProjectDatabasePostgresSqlUser extends pulumi.CustomResource {
    /**
     * Get an existing CloudProjectDatabasePostgresSqlUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CloudProjectDatabasePostgresSqlUserState, opts?: pulumi.CustomResourceOptions): CloudProjectDatabasePostgresSqlUser {
        return new CloudProjectDatabasePostgresSqlUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/cloudProjectDatabasePostgresSqlUser:CloudProjectDatabasePostgresSqlUser';

    /**
     * Returns true if the given object is an instance of CloudProjectDatabasePostgresSqlUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudProjectDatabasePostgresSqlUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudProjectDatabasePostgresSqlUser.__pulumiType;
    }

    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Date of the creation of the user.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Name of the user.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Password of the user.
     */
    public /*out*/ readonly password!: pulumi.Output<string>;
    /**
     * Roles the user belongs to. Possible values:
     * * `["replication"]`
     */
    public readonly roles!: pulumi.Output<string[] | undefined>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Current status of the user.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a CloudProjectDatabasePostgresSqlUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudProjectDatabasePostgresSqlUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CloudProjectDatabasePostgresSqlUserArgs | CloudProjectDatabasePostgresSqlUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CloudProjectDatabasePostgresSqlUserState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as CloudProjectDatabasePostgresSqlUserArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["password"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CloudProjectDatabasePostgresSqlUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CloudProjectDatabasePostgresSqlUser resources.
 */
export interface CloudProjectDatabasePostgresSqlUserState {
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Date of the creation of the user.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Name of the user.
     */
    name?: pulumi.Input<string>;
    /**
     * Password of the user.
     */
    password?: pulumi.Input<string>;
    /**
     * Roles the user belongs to. Possible values:
     * * `["replication"]`
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Current status of the user.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CloudProjectDatabasePostgresSqlUser resource.
 */
export interface CloudProjectDatabasePostgresSqlUserArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Name of the user.
     */
    name?: pulumi.Input<string>;
    /**
     * Roles the user belongs to. Possible values:
     * * `["replication"]`
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
