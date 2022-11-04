// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.CloudProject.KubeOidcArgs;
import com.pulumi.ovh.CloudProject.inputs.KubeOidcState;
import com.pulumi.ovh.Utilities;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Creates an OIDC configuration in an OVHcloud Managed Kubernetes cluster.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.CloudProject.KubeOidc;
 * import com.pulumi.ovh.CloudProject.KubeOidcArgs;
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
 *         var my_oidc = new KubeOidc(&#34;my-oidc&#34;, KubeOidcArgs.builder()        
 *             .serviceName(var_.projectid())
 *             .kubeId(ovh_cloud_project_kube.mykube().id())
 *             .clientId(&#34;xxx&#34;)
 *             .issuerUrl(&#34;https://ovh.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProject/kubeOidc:KubeOidc")
public class KubeOidc extends com.pulumi.resources.CustomResource {
    /**
     * The OIDC client ID.
     * 
     */
    @Export(name="clientId", type=String.class, parameters={})
    private Output<String> clientId;

    /**
     * @return The OIDC client ID.
     * 
     */
    public Output<String> clientId() {
        return this.clientId;
    }
    /**
     * The OIDC issuer url.
     * 
     */
    @Export(name="issuerUrl", type=String.class, parameters={})
    private Output<String> issuerUrl;

    /**
     * @return The OIDC issuer url.
     * 
     */
    public Output<String> issuerUrl() {
        return this.issuerUrl;
    }
    /**
     * The ID of the managed kubernetes cluster.
     * 
     */
    @Export(name="kubeId", type=String.class, parameters={})
    private Output<String> kubeId;

    /**
     * @return The ID of the managed kubernetes cluster.
     * 
     */
    public Output<String> kubeId() {
        return this.kubeId;
    }
    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Export(name="serviceName", type=String.class, parameters={})
    private Output<String> serviceName;

    /**
     * @return The ID of the public cloud project. If omitted,
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
    public KubeOidc(String name) {
        this(name, KubeOidcArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KubeOidc(String name, KubeOidcArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KubeOidc(String name, KubeOidcArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/kubeOidc:KubeOidc", name, args == null ? KubeOidcArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private KubeOidc(String name, Output<String> id, @Nullable KubeOidcState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/kubeOidc:KubeOidc", name, state, makeResourceOptions(options, id));
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
    public static KubeOidc get(String name, Output<String> id, @Nullable KubeOidcState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new KubeOidc(name, id, state, options);
    }
}