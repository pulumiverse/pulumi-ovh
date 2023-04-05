// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.CloudProject.RegionStoragePresignArgs;
import com.pulumi.ovh.CloudProject.inputs.RegionStoragePresignState;
import com.pulumi.ovh.Utilities;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Generates a temporary presigned S3 URLs to download or upload an object.
 * 
 * &gt; __NOTE__ This resource is only compatible with the `High Performance - S3` solution for object storage.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.CloudProject.RegionStoragePresign;
 * import com.pulumi.ovh.CloudProject.RegionStoragePresignArgs;
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
 *         var presignedUrlRegionStoragePresign = new RegionStoragePresign(&#34;presignedUrlRegionStoragePresign&#34;, RegionStoragePresignArgs.builder()        
 *             .serviceName(&#34;xxxxxxxxxxxxxxxxx&#34;)
 *             .regionName(&#34;GRA&#34;)
 *             .expire(3600)
 *             .method(&#34;GET&#34;)
 *             .object(&#34;an-object-in-the-bucket&#34;)
 *             .build());
 * 
 *         ctx.export(&#34;presignedUrl&#34;, presignedUrlRegionStoragePresign.url());
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProject/regionStoragePresign:RegionStoragePresign")
public class RegionStoragePresign extends com.pulumi.resources.CustomResource {
    /**
     * Define, in seconds, for how long your URL will be valid.
     * 
     */
    @Export(name="expire", refs={Integer.class}, tree="[0]")
    private Output<Integer> expire;

    /**
     * @return Define, in seconds, for how long your URL will be valid.
     * 
     */
    public Output<Integer> expire() {
        return this.expire;
    }
    /**
     * The method you want to use to interact with your object. Can be either &#39;GET&#39; or &#39;PUT&#39;.
     * 
     */
    @Export(name="method", refs={String.class}, tree="[0]")
    private Output<String> method;

    /**
     * @return The method you want to use to interact with your object. Can be either &#39;GET&#39; or &#39;PUT&#39;.
     * 
     */
    public Output<String> method() {
        return this.method;
    }
    /**
     * The name of your S3 storage container/bucket.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of your S3 storage container/bucket.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The name of the object in your S3 bucket.
     * 
     */
    @Export(name="object", refs={String.class}, tree="[0]")
    private Output<String> object;

    /**
     * @return The name of the object in your S3 bucket.
     * 
     */
    public Output<String> object() {
        return this.object;
    }
    /**
     * The region in which your storage is located.
     * Ex.: &#34;GRA&#34;.
     * 
     */
    @Export(name="regionName", refs={String.class}, tree="[0]")
    private Output<String> regionName;

    /**
     * @return The region in which your storage is located.
     * Ex.: &#34;GRA&#34;.
     * 
     */
    public Output<String> regionName() {
        return this.regionName;
    }
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
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
     * Computed URL result.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return Computed URL result.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RegionStoragePresign(String name) {
        this(name, RegionStoragePresignArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RegionStoragePresign(String name, RegionStoragePresignArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RegionStoragePresign(String name, RegionStoragePresignArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/regionStoragePresign:RegionStoragePresign", name, args == null ? RegionStoragePresignArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RegionStoragePresign(String name, Output<String> id, @Nullable RegionStoragePresignState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/regionStoragePresign:RegionStoragePresign", name, state, makeResourceOptions(options, id));
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
    public static RegionStoragePresign get(String name, Output<String> id, @Nullable RegionStoragePresignState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RegionStoragePresign(name, id, state, options);
    }
}
