// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.akamai.IamBlockedUserPropertiesArgs;
import com.pulumi.akamai.Utilities;
import com.pulumi.akamai.inputs.IamBlockedUserPropertiesState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Use the `akamai.IamBlockedUserProperties` resource to remove or grant access to properties. Administrators can block a user&#39;s access to any property, overriding any available role already assigned to that user.
 * 
 * ## Basic usage
 * 
 * This example returns the policy details based on the policy ID and optionally, a version:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.akamai.IamBlockedUserProperties;
 * import com.pulumi.akamai.IamBlockedUserPropertiesArgs;
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
 *         var example = new IamBlockedUserProperties(&#34;example&#34;, IamBlockedUserPropertiesArgs.builder()        
 *             .blockedProperties(            
 *                 1,
 *                 2,
 *                 3,
 *                 4,
 *                 5)
 *             .groupId(12345)
 *             .identityId(&#34;A-B-123456&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Attributes reference
 * 
 * This resource doesn&#39;t return any attributes.
 * 
 */
@ResourceType(type="akamai:index/iamBlockedUserProperties:IamBlockedUserProperties")
public class IamBlockedUserProperties extends com.pulumi.resources.CustomResource {
    /**
     * List of properties to block for a user. The property IDs must be an integer.
     * 
     */
    @Export(name="blockedProperties", type=List.class, parameters={Integer.class})
    private Output<List<Integer>> blockedProperties;

    /**
     * @return List of properties to block for a user. The property IDs must be an integer.
     * 
     */
    public Output<List<Integer>> blockedProperties() {
        return this.blockedProperties;
    }
    /**
     * A unique identifier for a group. Each identifier must be an integer.
     * 
     */
    @Export(name="groupId", type=Integer.class, parameters={})
    private Output<Integer> groupId;

    /**
     * @return A unique identifier for a group. Each identifier must be an integer.
     * 
     */
    public Output<Integer> groupId() {
        return this.groupId;
    }
    /**
     * A unique identifier that corresponds to a user&#39;s actual profile or client ID. Each identifier must be a string.
     * 
     */
    @Export(name="identityId", type=String.class, parameters={})
    private Output<String> identityId;

    /**
     * @return A unique identifier that corresponds to a user&#39;s actual profile or client ID. Each identifier must be a string.
     * 
     */
    public Output<String> identityId() {
        return this.identityId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IamBlockedUserProperties(String name) {
        this(name, IamBlockedUserPropertiesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IamBlockedUserProperties(String name, IamBlockedUserPropertiesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IamBlockedUserProperties(String name, IamBlockedUserPropertiesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/iamBlockedUserProperties:IamBlockedUserProperties", name, args == null ? IamBlockedUserPropertiesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IamBlockedUserProperties(String name, Output<String> id, @Nullable IamBlockedUserPropertiesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/iamBlockedUserProperties:IamBlockedUserProperties", name, state, makeResourceOptions(options, id));
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
    public static IamBlockedUserProperties get(String name, Output<String> id, @Nullable IamBlockedUserPropertiesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IamBlockedUserProperties(name, id, state, options);
    }
}