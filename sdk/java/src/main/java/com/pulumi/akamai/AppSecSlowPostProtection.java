// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.akamai.AppSecSlowPostProtectionArgs;
import com.pulumi.akamai.Utilities;
import com.pulumi.akamai.inputs.AppSecSlowPostProtectionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * **Scopes**: Security policy
 * 
 * Enables or disables slow POST protection for a security configuration and security policy. Slow POST protections help defend a site against attacks that try to tie up the site by using extremely slow requests and responses.
 * 
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/protections](https://techdocs.akamai.com/application-security/reference/put-policy-protections)
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
 * import com.pulumi.akamai.AppSecSlowPostProtection;
 * import com.pulumi.akamai.AppSecSlowPostProtectionArgs;
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
 *         var protection = new AppSecSlowPostProtection(&#34;protection&#34;, AppSecSlowPostProtectionArgs.builder()        
 *             .configId(configuration.applyValue(getAppSecConfigurationResult -&gt; getAppSecConfigurationResult.configId()))
 *             .securityPolicyId(&#34;gms1_134637&#34;)
 *             .enabled(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Output Options
 * 
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 * 
 * - `output_text`. Tabular report showing the current protection settings.
 * 
 */
@ResourceType(type="akamai:index/appSecSlowPostProtection:AppSecSlowPostProtection")
public class AppSecSlowPostProtection extends com.pulumi.resources.CustomResource {
    /**
     * . Unique identifier of the security configuration associated with the slow POST protection settings being modified.
     * 
     */
    @Export(name="configId", type=Integer.class, parameters={})
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the slow POST protection settings being modified.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }
    /**
     * . Set to **true** to enable slow POST protection; set to **false** to disable slow POST protection.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output<Boolean> enabled;

    /**
     * @return . Set to **true** to enable slow POST protection; set to **false** to disable slow POST protection.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * Text representation
     * 
     */
    @Export(name="outputText", type=String.class, parameters={})
    private Output<String> outputText;

    /**
     * @return Text representation
     * 
     */
    public Output<String> outputText() {
        return this.outputText;
    }
    /**
     * . Unique identifier of the security policy associated with the slow POST protection settings being modified.
     * 
     */
    @Export(name="securityPolicyId", type=String.class, parameters={})
    private Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the slow POST protection settings being modified.
     * 
     */
    public Output<String> securityPolicyId() {
        return this.securityPolicyId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppSecSlowPostProtection(String name) {
        this(name, AppSecSlowPostProtectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppSecSlowPostProtection(String name, AppSecSlowPostProtectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppSecSlowPostProtection(String name, AppSecSlowPostProtectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecSlowPostProtection:AppSecSlowPostProtection", name, args == null ? AppSecSlowPostProtectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppSecSlowPostProtection(String name, Output<String> id, @Nullable AppSecSlowPostProtectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecSlowPostProtection:AppSecSlowPostProtection", name, state, makeResourceOptions(options, id));
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
    public static AppSecSlowPostProtection get(String name, Output<String> id, @Nullable AppSecSlowPostProtectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppSecSlowPostProtection(name, id, state, options);
    }
}
