// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.akamai.AppSecAdvancedSettingsLoggingArgs;
import com.pulumi.akamai.Utilities;
import com.pulumi.akamai.inputs.AppSecAdvancedSettingsLoggingState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * **Scopes**: Security configuration; security policy
 * 
 * Enables, disables, or updates HTTP header logging settings.
 * By default, this operation applies at the configuration level, which means that it applies to all the security policies within that configuration.
 * However, by using the `security_policy_id` parameter you can specify custom settings for an individual security policy.
 * 
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/advanced-settings/logging](https://techdocs.akamai.com/application-security/reference/put-policies-logging)
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
 * import com.pulumi.akamai.AppSecAdvancedSettingsLogging;
 * import com.pulumi.akamai.AppSecAdvancedSettingsLoggingArgs;
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
 *         var logging = new AppSecAdvancedSettingsLogging(&#34;logging&#34;, AppSecAdvancedSettingsLoggingArgs.builder()        
 *             .configId(configuration.applyValue(getAppSecConfigurationResult -&gt; getAppSecConfigurationResult.configId()))
 *             .logging(Files.readString(Paths.get(String.format(&#34;%s/logging.json&#34;, path.module()))))
 *             .build());
 * 
 *         var policyLogging = new AppSecAdvancedSettingsLogging(&#34;policyLogging&#34;, AppSecAdvancedSettingsLoggingArgs.builder()        
 *             .configId(configuration.applyValue(getAppSecConfigurationResult -&gt; getAppSecConfigurationResult.configId()))
 *             .securityPolicyId(&#34;gms1_134637&#34;)
 *             .logging(Files.readString(Paths.get(String.format(&#34;%s/logging.json&#34;, path.module()))))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="akamai:index/appSecAdvancedSettingsLogging:AppSecAdvancedSettingsLogging")
public class AppSecAdvancedSettingsLogging extends com.pulumi.resources.CustomResource {
    /**
     * . Unique identifier of the security configuration containing the logging settings being modified.
     * 
     */
    @Export(name="configId", type=Integer.class, parameters={})
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration containing the logging settings being modified.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }
    /**
     * . Path to a JSON file containing the logging settings to be configured.
     * 
     */
    @Export(name="logging", type=String.class, parameters={})
    private Output<String> logging;

    /**
     * @return . Path to a JSON file containing the logging settings to be configured.
     * 
     */
    public Output<String> logging() {
        return this.logging;
    }
    /**
     * . Unique identifier of the security policies whose settings are being modified. If not included, the logging settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
     * 
     */
    @Export(name="securityPolicyId", type=String.class, parameters={})
    private Output</* @Nullable */ String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policies whose settings are being modified. If not included, the logging settings are modified at the configuration scope and, as a result, apply to all the security policies associated with the configuration.
     * 
     */
    public Output<Optional<String>> securityPolicyId() {
        return Codegen.optional(this.securityPolicyId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppSecAdvancedSettingsLogging(String name) {
        this(name, AppSecAdvancedSettingsLoggingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppSecAdvancedSettingsLogging(String name, AppSecAdvancedSettingsLoggingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppSecAdvancedSettingsLogging(String name, AppSecAdvancedSettingsLoggingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecAdvancedSettingsLogging:AppSecAdvancedSettingsLogging", name, args == null ? AppSecAdvancedSettingsLoggingArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppSecAdvancedSettingsLogging(String name, Output<String> id, @Nullable AppSecAdvancedSettingsLoggingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecAdvancedSettingsLogging:AppSecAdvancedSettingsLogging", name, state, makeResourceOptions(options, id));
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
    public static AppSecAdvancedSettingsLogging get(String name, Output<String> id, @Nullable AppSecAdvancedSettingsLoggingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppSecAdvancedSettingsLogging(name, id, state, options);
    }
}
