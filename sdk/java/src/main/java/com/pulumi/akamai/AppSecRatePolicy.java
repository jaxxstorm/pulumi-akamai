// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.akamai.AppSecRatePolicyArgs;
import com.pulumi.akamai.Utilities;
import com.pulumi.akamai.inputs.AppSecRatePolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * **Scopes**: Security configuration; rate policy
 * 
 * Creates, modifies, or deletes rate policies. Rate polices help you monitor and moderate the number and rate of all the requests you receive.
 * In turn, this helps you prevent your website from being overwhelmed by a dramatic and unexpected surge in traffic.
 * 
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/rate-policies](https://techdocs.akamai.com/application-security/reference/post-rate-policies)
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
 * import com.pulumi.akamai.AppSecRatePolicy;
 * import com.pulumi.akamai.AppSecRatePolicyArgs;
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
 *         var ratePolicy = new AppSecRatePolicy(&#34;ratePolicy&#34;, AppSecRatePolicyArgs.builder()        
 *             .configId(configuration.applyValue(getAppSecConfigurationResult -&gt; getAppSecConfigurationResult.configId()))
 *             .ratePolicy(Files.readString(Paths.get(String.format(&#34;%s/rate_policy.json&#34;, path.module()))))
 *             .build());
 * 
 *         ctx.export(&#34;ratePolicyId&#34;, ratePolicy.ratePolicyId());
 *     }
 * }
 * ```
 * ## Output Options
 * 
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 * 
 * - `rate_policy_id`. ID of the modified or newly-created rate policy.
 * 
 */
@ResourceType(type="akamai:index/appSecRatePolicy:AppSecRatePolicy")
public class AppSecRatePolicy extends com.pulumi.resources.CustomResource {
    /**
     * . Unique identifier of the security configuration associated with the rate policy being modified.
     * 
     */
    @Export(name="configId", type=Integer.class, parameters={})
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the rate policy being modified.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }
    /**
     * . Path to a JSON file containing a rate policy definition.
     * 
     */
    @Export(name="ratePolicy", type=String.class, parameters={})
    private Output<String> ratePolicy;

    /**
     * @return . Path to a JSON file containing a rate policy definition.
     * 
     */
    public Output<String> ratePolicy() {
        return this.ratePolicy;
    }
    /**
     * Unique identifier of the rate policy
     * 
     */
    @Export(name="ratePolicyId", type=Integer.class, parameters={})
    private Output<Integer> ratePolicyId;

    /**
     * @return Unique identifier of the rate policy
     * 
     */
    public Output<Integer> ratePolicyId() {
        return this.ratePolicyId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppSecRatePolicy(String name) {
        this(name, AppSecRatePolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppSecRatePolicy(String name, AppSecRatePolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppSecRatePolicy(String name, AppSecRatePolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecRatePolicy:AppSecRatePolicy", name, args == null ? AppSecRatePolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppSecRatePolicy(String name, Output<String> id, @Nullable AppSecRatePolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecRatePolicy:AppSecRatePolicy", name, state, makeResourceOptions(options, id));
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
    public static AppSecRatePolicy get(String name, Output<String> id, @Nullable AppSecRatePolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppSecRatePolicy(name, id, state, options);
    }
}
