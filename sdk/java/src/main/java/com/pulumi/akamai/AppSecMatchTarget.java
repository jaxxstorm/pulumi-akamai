// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.akamai.AppSecMatchTargetArgs;
import com.pulumi.akamai.Utilities;
import com.pulumi.akamai.inputs.AppSecMatchTargetState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * **Scopes**: Security configuration
 * 
 * Creates a match target associated with a security configuration. Match targets determine which security policy should apply to an API, hostname or path.
 * 
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/match-targets](https://techdocs.akamai.com/application-security/reference/post-match-targets)
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
 * import com.pulumi.akamai.inputs.GetAppSecConfigurationArgs;
 * import com.pulumi.akamai.AppSecMatchTarget;
 * import com.pulumi.akamai.AppSecMatchTargetArgs;
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
 *         final var configuration = AkamaiFunctions.getAppSecConfiguration(GetAppSecConfigurationArgs.builder()
 *             .name(&#34;Documentation&#34;)
 *             .build());
 * 
 *         var matchTarget = new AppSecMatchTarget(&#34;matchTarget&#34;, AppSecMatchTargetArgs.builder()        
 *             .configId(configuration.applyValue(getAppSecConfigurationResult -&gt; getAppSecConfigurationResult.configId()))
 *             .matchTarget(Files.readString(Paths.get(String.format(&#34;%s/match_targets.json&#34;, path.module()))))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Output Options
 * 
 * In addition to the arguments above, the following attribute is exported:
 * 
 * - `match_target_id`. ID of the match target.
 * 
 */
@ResourceType(type="akamai:index/appSecMatchTarget:AppSecMatchTarget")
public class AppSecMatchTarget extends com.pulumi.resources.CustomResource {
    /**
     * . Unique identifier of the security configuration associated with the match target being modified.
     * 
     */
    @Export(name="configId", type=Integer.class, parameters={})
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the match target being modified.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }
    /**
     * . Path to a JSON file containing one or more match target definitions.
     * 
     */
    @Export(name="matchTarget", type=String.class, parameters={})
    private Output<String> matchTarget;

    /**
     * @return . Path to a JSON file containing one or more match target definitions.
     * 
     */
    public Output<String> matchTarget() {
        return this.matchTarget;
    }
    /**
     * Unique identifier of the match target
     * 
     */
    @Export(name="matchTargetId", type=Integer.class, parameters={})
    private Output<Integer> matchTargetId;

    /**
     * @return Unique identifier of the match target
     * 
     */
    public Output<Integer> matchTargetId() {
        return this.matchTargetId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppSecMatchTarget(String name) {
        this(name, AppSecMatchTargetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppSecMatchTarget(String name, AppSecMatchTargetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppSecMatchTarget(String name, AppSecMatchTargetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecMatchTarget:AppSecMatchTarget", name, args == null ? AppSecMatchTargetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppSecMatchTarget(String name, Output<String> id, @Nullable AppSecMatchTargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecMatchTarget:AppSecMatchTarget", name, state, makeResourceOptions(options, id));
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
    public static AppSecMatchTarget get(String name, Output<String> id, @Nullable AppSecMatchTargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppSecMatchTarget(name, id, state, options);
    }
}
