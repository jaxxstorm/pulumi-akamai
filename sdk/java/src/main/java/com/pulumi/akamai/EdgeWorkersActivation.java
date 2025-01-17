// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.akamai.EdgeWorkersActivationArgs;
import com.pulumi.akamai.Utilities;
import com.pulumi.akamai.inputs.EdgeWorkersActivationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Use the `akamai.EdgeWorkersActivation` resource to activate a specific EdgeWorker version. An activation deploys the version to either the Akamai staging or production network.
 * 
 * Before activating on production, activate on staging first. This way you can detect any problems in staging before your changes progress to production.
 * 
 * ## Example Usage
 * 
 * Basic usage:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.akamai.EdgeWorkersActivation;
 * import com.pulumi.akamai.EdgeWorkersActivationArgs;
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
 *         var test = new EdgeWorkersActivation(&#34;test&#34;, EdgeWorkersActivationArgs.builder()        
 *             .edgeworkerId(1234)
 *             .network(&#34;STAGING&#34;)
 *             .version(&#34;test1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="akamai:index/edgeWorkersActivation:EdgeWorkersActivation")
public class EdgeWorkersActivation extends com.pulumi.resources.CustomResource {
    /**
     * (Required) Unique identifier of the activation.
     * 
     */
    @Export(name="activationId", type=Integer.class, parameters={})
    private Output<Integer> activationId;

    /**
     * @return (Required) Unique identifier of the activation.
     * 
     */
    public Output<Integer> activationId() {
        return this.activationId;
    }
    /**
     * A unique identifier for the EdgeWorker ID you want to activate.
     * 
     */
    @Export(name="edgeworkerId", type=Integer.class, parameters={})
    private Output<Integer> edgeworkerId;

    /**
     * @return A unique identifier for the EdgeWorker ID you want to activate.
     * 
     */
    public Output<Integer> edgeworkerId() {
        return this.edgeworkerId;
    }
    /**
     * The network you want to activate the policy version on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
     * 
     */
    @Export(name="network", type=String.class, parameters={})
    private Output<String> network;

    /**
     * @return The network you want to activate the policy version on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
     * 
     */
    public Output<String> network() {
        return this.network;
    }
    /**
     * The EdgeWorker version you want to activate.
     * 
     */
    @Export(name="version", type=String.class, parameters={})
    private Output<String> version;

    /**
     * @return The EdgeWorker version you want to activate.
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EdgeWorkersActivation(String name) {
        this(name, EdgeWorkersActivationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EdgeWorkersActivation(String name, EdgeWorkersActivationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EdgeWorkersActivation(String name, EdgeWorkersActivationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/edgeWorkersActivation:EdgeWorkersActivation", name, args == null ? EdgeWorkersActivationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EdgeWorkersActivation(String name, Output<String> id, @Nullable EdgeWorkersActivationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/edgeWorkersActivation:EdgeWorkersActivation", name, state, makeResourceOptions(options, id));
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
    public static EdgeWorkersActivation get(String name, Output<String> id, @Nullable EdgeWorkersActivationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EdgeWorkersActivation(name, id, state, options);
    }
}
