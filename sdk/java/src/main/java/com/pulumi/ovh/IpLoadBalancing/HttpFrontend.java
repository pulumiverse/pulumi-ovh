// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.IpLoadBalancing;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.IpLoadBalancing.HttpFrontendArgs;
import com.pulumi.ovh.IpLoadBalancing.inputs.HttpFrontendState;
import com.pulumi.ovh.Utilities;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a backend HTTP server group (frontend) to be used by loadbalancing frontend(s)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.IpLoadBalancing.IpLoadBalancingFunctions;
 * import com.pulumi.ovh.IpLoadBalancing.inputs.GetIpLoadBalancingArgs;
 * import com.pulumi.ovh.IpLoadBalancing.HttpFarm;
 * import com.pulumi.ovh.IpLoadBalancing.HttpFarmArgs;
 * import com.pulumi.ovh.IpLoadBalancing.HttpFrontend;
 * import com.pulumi.ovh.IpLoadBalancing.HttpFrontendArgs;
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
 *         final var lb = IpLoadBalancingFunctions.getIpLoadBalancing(GetIpLoadBalancingArgs.builder()
 *             .serviceName(&#34;ip-1.2.3.4&#34;)
 *             .state(&#34;ok&#34;)
 *             .build());
 * 
 *         var farm80 = new HttpFarm(&#34;farm80&#34;, HttpFarmArgs.builder()        
 *             .displayName(&#34;ingress-8080-gra&#34;)
 *             .port(80)
 *             .serviceName(lb.applyValue(getIpLoadBalancingResult -&gt; getIpLoadBalancingResult.serviceName()))
 *             .zone(&#34;all&#34;)
 *             .build());
 * 
 *         var testfrontend = new HttpFrontend(&#34;testfrontend&#34;, HttpFrontendArgs.builder()        
 *             .defaultFarmId(farm80.id())
 *             .displayName(&#34;ingress-8080-gra&#34;)
 *             .port(&#34;80,443&#34;)
 *             .serviceName(lb.applyValue(getIpLoadBalancingResult -&gt; getIpLoadBalancingResult.serviceName()))
 *             .zone(&#34;all&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### With HTTP Header
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.IpLoadBalancing.IpLoadBalancingFunctions;
 * import com.pulumi.ovh.IpLoadBalancing.inputs.GetIpLoadBalancingArgs;
 * import com.pulumi.ovh.IpLoadBalancing.HttpFarm;
 * import com.pulumi.ovh.IpLoadBalancing.HttpFarmArgs;
 * import com.pulumi.ovh.IpLoadBalancing.HttpFrontend;
 * import com.pulumi.ovh.IpLoadBalancing.HttpFrontendArgs;
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
 *         final var lb = IpLoadBalancingFunctions.getIpLoadBalancing(GetIpLoadBalancingArgs.builder()
 *             .serviceName(&#34;ip-1.2.3.4&#34;)
 *             .state(&#34;ok&#34;)
 *             .build());
 * 
 *         var farm80 = new HttpFarm(&#34;farm80&#34;, HttpFarmArgs.builder()        
 *             .displayName(&#34;ingress-8080-gra&#34;)
 *             .port(80)
 *             .serviceName(lb.applyValue(getIpLoadBalancingResult -&gt; getIpLoadBalancingResult.serviceName()))
 *             .zone(&#34;all&#34;)
 *             .build());
 * 
 *         var testfrontend = new HttpFrontend(&#34;testfrontend&#34;, HttpFrontendArgs.builder()        
 *             .defaultFarmId(farm80.id())
 *             .displayName(&#34;ingress-8080-gra&#34;)
 *             .httpHeaders(            
 *                 &#34;X-Ip-Header %%ci&#34;,
 *                 &#34;X-Port-Header %%cp&#34;)
 *             .port(&#34;80,443&#34;)
 *             .serviceName(lb.applyValue(getIpLoadBalancingResult -&gt; getIpLoadBalancingResult.serviceName()))
 *             .zone(&#34;all&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="ovh:IpLoadBalancing/httpFrontend:HttpFrontend")
public class HttpFrontend extends com.pulumi.resources.CustomResource {
    /**
     * Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
     * 
     */
    @Export(name="allowedSources", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> allowedSources;

    /**
     * @return Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
     * 
     */
    public Output<Optional<List<String>>> allowedSources() {
        return Codegen.optional(this.allowedSources);
    }
    /**
     * Only attach frontend on these ip. No restriction if null. List of Ip blocks.
     * 
     */
    @Export(name="dedicatedIpfos", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> dedicatedIpfos;

    /**
     * @return Only attach frontend on these ip. No restriction if null. List of Ip blocks.
     * 
     */
    public Output<Optional<List<String>>> dedicatedIpfos() {
        return Codegen.optional(this.dedicatedIpfos);
    }
    /**
     * Default TCP Farm of your frontend
     * 
     */
    @Export(name="defaultFarmId", type=Integer.class, parameters={})
    private Output<Integer> defaultFarmId;

    /**
     * @return Default TCP Farm of your frontend
     * 
     */
    public Output<Integer> defaultFarmId() {
        return this.defaultFarmId;
    }
    /**
     * Default ssl served to your customer
     * 
     */
    @Export(name="defaultSslId", type=Integer.class, parameters={})
    private Output<Integer> defaultSslId;

    /**
     * @return Default ssl served to your customer
     * 
     */
    public Output<Integer> defaultSslId() {
        return this.defaultSslId;
    }
    /**
     * Disable your frontend. Default: &#39;false&#39;
     * 
     */
    @Export(name="disabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> disabled;

    /**
     * @return Disable your frontend. Default: &#39;false&#39;
     * 
     */
    public Output<Optional<Boolean>> disabled() {
        return Codegen.optional(this.disabled);
    }
    /**
     * Human readable name for your frontend, this field is for you
     * 
     */
    @Export(name="displayName", type=String.class, parameters={})
    private Output</* @Nullable */ String> displayName;

    /**
     * @return Human readable name for your frontend, this field is for you
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * HTTP Strict Transport Security. Default: &#39;false&#39;
     * 
     */
    @Export(name="hsts", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> hsts;

    /**
     * @return HTTP Strict Transport Security. Default: &#39;false&#39;
     * 
     */
    public Output<Optional<Boolean>> hsts() {
        return Codegen.optional(this.hsts);
    }
    /**
     * HTTP headers to add to the frontend. List of string.
     * 
     */
    @Export(name="httpHeaders", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> httpHeaders;

    /**
     * @return HTTP headers to add to the frontend. List of string.
     * 
     */
    public Output<Optional<List<String>>> httpHeaders() {
        return Codegen.optional(this.httpHeaders);
    }
    /**
     * Port(s) attached to your frontend. Supports single port (numerical value),
     * range (2 dash-delimited increasing ports) and comma-separated list of &#39;single port&#39;
     * and/or &#39;range&#39;. Each port must be in the [1;49151] range
     * 
     */
    @Export(name="port", type=String.class, parameters={})
    private Output<String> port;

    /**
     * @return Port(s) attached to your frontend. Supports single port (numerical value),
     * range (2 dash-delimited increasing ports) and comma-separated list of &#39;single port&#39;
     * and/or &#39;range&#39;. Each port must be in the [1;49151] range
     * 
     */
    public Output<String> port() {
        return this.port;
    }
    /**
     * Redirection HTTP&#39;
     * 
     */
    @Export(name="redirectLocation", type=String.class, parameters={})
    private Output</* @Nullable */ String> redirectLocation;

    /**
     * @return Redirection HTTP&#39;
     * 
     */
    public Output<Optional<String>> redirectLocation() {
        return Codegen.optional(this.redirectLocation);
    }
    /**
     * The internal name of your IP load balancing
     * 
     */
    @Export(name="serviceName", type=String.class, parameters={})
    private Output<String> serviceName;

    /**
     * @return The internal name of your IP load balancing
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * SSL deciphering. Default: &#39;false&#39;
     * 
     */
    @Export(name="ssl", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> ssl;

    /**
     * @return SSL deciphering. Default: &#39;false&#39;
     * 
     */
    public Output<Optional<Boolean>> ssl() {
        return Codegen.optional(this.ssl);
    }
    /**
     * Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
     * 
     */
    @Export(name="zone", type=String.class, parameters={})
    private Output<String> zone;

    /**
     * @return Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HttpFrontend(String name) {
        this(name, HttpFrontendArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HttpFrontend(String name, HttpFrontendArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HttpFrontend(String name, HttpFrontendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:IpLoadBalancing/httpFrontend:HttpFrontend", name, args == null ? HttpFrontendArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HttpFrontend(String name, Output<String> id, @Nullable HttpFrontendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:IpLoadBalancing/httpFrontend:HttpFrontend", name, state, makeResourceOptions(options, id));
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
    public static HttpFrontend get(String name, Output<String> id, @Nullable HttpFrontendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HttpFrontend(name, id, state, options);
    }
}
