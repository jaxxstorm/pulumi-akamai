// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.akamai.NetworkListActivationsArgs;
import com.pulumi.akamai.Utilities;
import com.pulumi.akamai.inputs.NetworkListActivationsState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Use the `akamai.NetworkListActivations` resource to activate a network list in either the STAGING or PRODUCTION
 * environment.
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
 * import com.pulumi.akamai.AkamaiFunctions;
 * import com.pulumi.akamai.inputs.GetNetworkListsArgs;
 * import com.pulumi.akamai.NetworkListActivations;
 * import com.pulumi.akamai.NetworkListActivationsArgs;
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
 *         final var networkListsFilter = AkamaiFunctions.getNetworkLists(GetNetworkListsArgs.builder()
 *             .name(var_.network_list())
 *             .build());
 * 
 *         var activation = new NetworkListActivations(&#34;activation&#34;, NetworkListActivationsArgs.builder()        
 *             .networkListId(networkListsFilter.applyValue(getNetworkListsResult -&gt; getNetworkListsResult.lists()[0]))
 *             .network(&#34;STAGING&#34;)
 *             .notes(&#34;TEST Notes&#34;)
 *             .notificationEmails(&#34;user@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="akamai:index/networkListActivations:NetworkListActivations")
public class NetworkListActivations extends com.pulumi.resources.CustomResource {
    /**
     * @deprecated
     * The setting &#34;activate&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""activate"" has been deprecated. */
    @Export(name="activate", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> activate;

    public Output<Optional<Boolean>> activate() {
        return Codegen.optional(this.activate);
    }
    /**
     * The network to be used, either `STAGING` or `PRODUCTION`. If not supplied, defaults to
     * `STAGING`.
     * 
     */
    @Export(name="network", type=String.class, parameters={})
    private Output</* @Nullable */ String> network;

    /**
     * @return The network to be used, either `STAGING` or `PRODUCTION`. If not supplied, defaults to
     * `STAGING`.
     * 
     */
    public Output<Optional<String>> network() {
        return Codegen.optional(this.network);
    }
    /**
     * The ID of the network list to be activated
     * 
     */
    @Export(name="networkListId", type=String.class, parameters={})
    private Output<String> networkListId;

    /**
     * @return The ID of the network list to be activated
     * 
     */
    public Output<String> networkListId() {
        return this.networkListId;
    }
    /**
     * A comment describing the activation.
     * 
     */
    @Export(name="notes", type=String.class, parameters={})
    private Output</* @Nullable */ String> notes;

    /**
     * @return A comment describing the activation.
     * 
     */
    public Output<Optional<String>> notes() {
        return Codegen.optional(this.notes);
    }
    /**
     * A bracketed, comma-separated list of email addresses that will be notified when the
     * operation is complete.
     * 
     */
    @Export(name="notificationEmails", type=List.class, parameters={String.class})
    private Output<List<String>> notificationEmails;

    /**
     * @return A bracketed, comma-separated list of email addresses that will be notified when the
     * operation is complete.
     * 
     */
    public Output<List<String>> notificationEmails() {
        return this.notificationEmails;
    }
    /**
     * The string `ACTIVATED` if the activation was successful, or a string identifying the reason why the network
     * list was not activated.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The string `ACTIVATED` if the activation was successful, or a string identifying the reason why the network
     * list was not activated.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NetworkListActivations(String name) {
        this(name, NetworkListActivationsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NetworkListActivations(String name, NetworkListActivationsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NetworkListActivations(String name, NetworkListActivationsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/networkListActivations:NetworkListActivations", name, args == null ? NetworkListActivationsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NetworkListActivations(String name, Output<String> id, @Nullable NetworkListActivationsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/networkListActivations:NetworkListActivations", name, state, makeResourceOptions(options, id));
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
    public static NetworkListActivations get(String name, Output<String> id, @Nullable NetworkListActivationsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NetworkListActivations(name, id, state, options);
    }
}
