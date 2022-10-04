// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.akamai.EdgeWorkerArgs;
import com.pulumi.akamai.Utilities;
import com.pulumi.akamai.inputs.EdgeWorkerState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `akamai.EdgeWorker` resource lets you deploy custom code on thousands of edge servers and apply logic that creates powerful web experiences.
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
 * import com.pulumi.akamai.EdgeWorker;
 * import com.pulumi.akamai.EdgeWorkerArgs;
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
 *         var ew = new EdgeWorker(&#34;ew&#34;, EdgeWorkerArgs.builder()        
 *             .groupId(72297)
 *             .resourceTierId(100)
 *             .localBundle(var_.bundle_path())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Attributes reference
 * 
 * * `edgeworker_id` - Unique identifier for an EdgeWorker ID.
 * * `local_bundle_hash` - A SHA-256 hash digest of the EdgeWorkers code bundle.
 * * `version` - Unique identifier for a specific EdgeWorker version.
 * * `warnings` - List of validation warnings.
 * 
 */
@ResourceType(type="akamai:index/edgeWorker:EdgeWorker")
public class EdgeWorker extends com.pulumi.resources.CustomResource {
    /**
     * The unique identifier of the EdgeWorker
     * 
     */
    @Export(name="edgeworkerId", type=Integer.class, parameters={})
    private Output<Integer> edgeworkerId;

    /**
     * @return The unique identifier of the EdgeWorker
     * 
     */
    public Output<Integer> edgeworkerId() {
        return this.edgeworkerId;
    }
    /**
     * - (Required) Identifies a group to assign to the EdgeWorker ID.
     * 
     */
    @Export(name="groupId", type=Integer.class, parameters={})
    private Output<Integer> groupId;

    /**
     * @return - (Required) Identifies a group to assign to the EdgeWorker ID.
     * 
     */
    public Output<Integer> groupId() {
        return this.groupId;
    }
    /**
     * - (Optional) The path to the EdgeWorkers code bundle.
     * 
     */
    @Export(name="localBundle", type=String.class, parameters={})
    private Output</* @Nullable */ String> localBundle;

    /**
     * @return - (Optional) The path to the EdgeWorkers code bundle.
     * 
     */
    public Output<Optional<String>> localBundle() {
        return Codegen.optional(this.localBundle);
    }
    /**
     * The local bundle hash for the EdgeWorker
     * 
     */
    @Export(name="localBundleHash", type=String.class, parameters={})
    private Output<String> localBundleHash;

    /**
     * @return The local bundle hash for the EdgeWorker
     * 
     */
    public Output<String> localBundleHash() {
        return this.localBundleHash;
    }
    /**
     * - (Required) The name of the EdgeWorker ID.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return - (Required) The name of the EdgeWorker ID.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * - (Required) Unique identifier of the resource tier.
     * 
     */
    @Export(name="resourceTierId", type=Integer.class, parameters={})
    private Output<Integer> resourceTierId;

    /**
     * @return - (Required) Unique identifier of the resource tier.
     * 
     */
    public Output<Integer> resourceTierId() {
        return this.resourceTierId;
    }
    /**
     * The bundle version
     * 
     */
    @Export(name="version", type=String.class, parameters={})
    private Output<String> version;

    /**
     * @return The bundle version
     * 
     */
    public Output<String> version() {
        return this.version;
    }
    /**
     * The list of warnings returned by EdgeWorker validation
     * 
     */
    @Export(name="warnings", type=List.class, parameters={String.class})
    private Output<List<String>> warnings;

    /**
     * @return The list of warnings returned by EdgeWorker validation
     * 
     */
    public Output<List<String>> warnings() {
        return this.warnings;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EdgeWorker(String name) {
        this(name, EdgeWorkerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EdgeWorker(String name, EdgeWorkerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EdgeWorker(String name, EdgeWorkerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/edgeWorker:EdgeWorker", name, args == null ? EdgeWorkerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EdgeWorker(String name, Output<String> id, @Nullable EdgeWorkerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/edgeWorker:EdgeWorker", name, state, makeResourceOptions(options, id));
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
    public static EdgeWorker get(String name, Output<String> id, @Nullable EdgeWorkerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EdgeWorker(name, id, state, options);
    }
}
