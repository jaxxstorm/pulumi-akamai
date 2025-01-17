// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.akamai.AppSecCustomDenyArgs;
import com.pulumi.akamai.Utilities;
import com.pulumi.akamai.inputs.AppSecCustomDenyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * **Scopes**: Custom deny
 * 
 * Modifies a custom deny action. Custom denies enable you to craft your own error message or redirect pages for use when HTTP requests are denied.
 * 
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/custom-deny](https://techdocs.akamai.com/application-security/reference/get-custom-deny-actions)
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
 * import com.pulumi.akamai.AppSecCustomDeny;
 * import com.pulumi.akamai.AppSecCustomDenyArgs;
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
 *         var customDeny = new AppSecCustomDeny(&#34;customDeny&#34;, AppSecCustomDenyArgs.builder()        
 *             .configId(configuration.applyValue(getAppSecConfigurationResult -&gt; getAppSecConfigurationResult.configId()))
 *             .customDeny(Files.readString(Paths.get(String.format(&#34;%s/custom_deny.json&#34;, path.module()))))
 *             .build());
 * 
 *         ctx.export(&#34;customDenyId&#34;, customDeny.customDenyId());
 *     }
 * }
 * ```
 * ## Output Options
 * 
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 * 
 * - `custom_deny_id`. ID of the new custom deny action.
 * 
 */
@ResourceType(type="akamai:index/appSecCustomDeny:AppSecCustomDeny")
public class AppSecCustomDeny extends com.pulumi.resources.CustomResource {
    /**
     * . Unique identifier of the security configuration associated with the custom deny.
     * 
     */
    @Export(name="configId", type=Integer.class, parameters={})
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the custom deny.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }
    /**
     * . Path to a JSON file containing properties and property values for the custom deny.
     * 
     */
    @Export(name="customDeny", type=String.class, parameters={})
    private Output<String> customDeny;

    /**
     * @return . Path to a JSON file containing properties and property values for the custom deny.
     * 
     */
    public Output<String> customDeny() {
        return this.customDeny;
    }
    /**
     * custom_deny_id
     * 
     */
    @Export(name="customDenyId", type=String.class, parameters={})
    private Output<String> customDenyId;

    /**
     * @return custom_deny_id
     * 
     */
    public Output<String> customDenyId() {
        return this.customDenyId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppSecCustomDeny(String name) {
        this(name, AppSecCustomDenyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppSecCustomDeny(String name, AppSecCustomDenyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppSecCustomDeny(String name, AppSecCustomDenyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecCustomDeny:AppSecCustomDeny", name, args == null ? AppSecCustomDenyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppSecCustomDeny(String name, Output<String> id, @Nullable AppSecCustomDenyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecCustomDeny:AppSecCustomDeny", name, state, makeResourceOptions(options, id));
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
    public static AppSecCustomDeny get(String name, Output<String> id, @Nullable AppSecCustomDenyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppSecCustomDeny(name, id, state, options);
    }
}
