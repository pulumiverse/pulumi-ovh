// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.CloudProject.KubeIpRestrictionsArgs;
import com.pulumi.ovh.CloudProject.inputs.KubeIpRestrictionsState;
import com.pulumi.ovh.Utilities;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Apply IP restrictions to an OVHcloud Managed Kubernetes cluster.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.CloudProject.KubeIpRestrictions;
 * import com.pulumi.ovh.CloudProject.KubeIpRestrictionsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var vrackOnly = new KubeIpRestrictions(&#34;vrackOnly&#34;, KubeIpRestrictionsArgs.builder()        
 *             .ips(&#34;10.42.0.0/16&#34;)
 *             .kubeId(&#34;xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx&#34;)
 *             .serviceName(&#34;xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * OVHcloud Managed Kubernetes Service cluster IP restrictions can be imported using the `service_name` and the `id` of the cluster, separated by &#34;/&#34; E.g., bash
 * 
 * ```sh
 *  $ pulumi import ovh:CloudProject/kubeIpRestrictions:KubeIpRestrictions iprestrictions service_name/kube_id
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProject/kubeIpRestrictions:KubeIpRestrictions")
public class KubeIpRestrictions extends com.pulumi.resources.CustomResource {
    /**
     * List of CIDR authorized to interact with the managed Kubernetes cluster.
     * 
     */
    @Export(name="ips", type=List.class, parameters={String.class})
    private Output<List<String>> ips;

    /**
     * @return List of CIDR authorized to interact with the managed Kubernetes cluster.
     * 
     */
    public Output<List<String>> ips() {
        return this.ips;
    }
    /**
     * The id of the managed Kubernetes cluster.
     * 
     */
    @Export(name="kubeId", type=String.class, parameters={})
    private Output<String> kubeId;

    /**
     * @return The id of the managed Kubernetes cluster.
     * 
     */
    public Output<String> kubeId() {
        return this.kubeId;
    }
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Export(name="serviceName", type=String.class, parameters={})
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public KubeIpRestrictions(String name) {
        this(name, KubeIpRestrictionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KubeIpRestrictions(String name, KubeIpRestrictionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KubeIpRestrictions(String name, KubeIpRestrictionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/kubeIpRestrictions:KubeIpRestrictions", name, args == null ? KubeIpRestrictionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private KubeIpRestrictions(String name, Output<String> id, @Nullable KubeIpRestrictionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/kubeIpRestrictions:KubeIpRestrictions", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static KubeIpRestrictions get(String name, Output<String> id, @Nullable KubeIpRestrictionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new KubeIpRestrictions(name, id, state, options);
    }
}