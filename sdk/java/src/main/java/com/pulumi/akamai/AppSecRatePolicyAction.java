// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.akamai.AppSecRatePolicyActionArgs;
import com.pulumi.akamai.Utilities;
import com.pulumi.akamai.inputs.AppSecRatePolicyActionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * **Scopes**: Rate policy
 * 
 * Creates, modifies, or deletes the actions associated with a rate policy.
 * By default, rate policies take no action when triggered.
 * Note that you must set separate actions for requests originating from an IPv4 IP address and for requests originating from an IPv6 address.
 * 
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rate-policies/{ratePolicyId}](https://techdocs.akamai.com/application-security/reference/put-rate-policy-action)
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
 * import com.pulumi.akamai.AppSecRatePolicyAction;
 * import com.pulumi.akamai.AppSecRatePolicyActionArgs;
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
 *         var appsecRatePolicy = new AppSecRatePolicy(&#34;appsecRatePolicy&#34;, AppSecRatePolicyArgs.builder()        
 *             .configId(configuration.applyValue(getAppSecConfigurationResult -&gt; getAppSecConfigurationResult.configId()))
 *             .ratePolicy(Files.readString(Paths.get(String.format(&#34;%s/rate_policy.json&#34;, path.module()))))
 *             .build());
 * 
 *         var appsecRatePolicyAction = new AppSecRatePolicyAction(&#34;appsecRatePolicyAction&#34;, AppSecRatePolicyActionArgs.builder()        
 *             .configId(configuration.applyValue(getAppSecConfigurationResult -&gt; getAppSecConfigurationResult.configId()))
 *             .securityPolicyId(&#34;gms1_134637&#34;)
 *             .ratePolicyId(appsecRatePolicy.ratePolicyId())
 *             .ipv4Action(&#34;deny&#34;)
 *             .ipv6Action(&#34;deny&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="akamai:index/appSecRatePolicyAction:AppSecRatePolicyAction")
public class AppSecRatePolicyAction extends com.pulumi.resources.CustomResource {
    /**
     * . Unique identifier of the security configuration associated with the rate policy action being modified.
     * 
     */
    @Export(name="configId", type=Integer.class, parameters={})
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the rate policy action being modified.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }
    /**
     * . Rate policy action for requests coming from an IPv4 IP address. Allowed actions are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     * 
     */
    @Export(name="ipv4Action", type=String.class, parameters={})
    private Output<String> ipv4Action;

    /**
     * @return . Rate policy action for requests coming from an IPv4 IP address. Allowed actions are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     * 
     */
    public Output<String> ipv4Action() {
        return this.ipv4Action;
    }
    /**
     * . Rate policy action for requests coming from an IPv6 IP address. Allowed actions are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom{custom_deny_id}**. Take the action specified by the custom deny.
     * 
     */
    @Export(name="ipv6Action", type=String.class, parameters={})
    private Output<String> ipv6Action;

    /**
     * @return . Rate policy action for requests coming from an IPv6 IP address. Allowed actions are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom{custom_deny_id}**. Take the action specified by the custom deny.
     * 
     */
    public Output<String> ipv6Action() {
        return this.ipv6Action;
    }
    /**
     * . Unique identifier of the rate policy whose action is being modified.
     * 
     */
    @Export(name="ratePolicyId", type=Integer.class, parameters={})
    private Output<Integer> ratePolicyId;

    /**
     * @return . Unique identifier of the rate policy whose action is being modified.
     * 
     */
    public Output<Integer> ratePolicyId() {
        return this.ratePolicyId;
    }
    /**
     * . Unique identifier of the security policy associated with the rate policy whose action is being modified.
     * 
     */
    @Export(name="securityPolicyId", type=String.class, parameters={})
    private Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the rate policy whose action is being modified.
     * 
     */
    public Output<String> securityPolicyId() {
        return this.securityPolicyId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppSecRatePolicyAction(String name) {
        this(name, AppSecRatePolicyActionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppSecRatePolicyAction(String name, AppSecRatePolicyActionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppSecRatePolicyAction(String name, AppSecRatePolicyActionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecRatePolicyAction:AppSecRatePolicyAction", name, args == null ? AppSecRatePolicyActionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppSecRatePolicyAction(String name, Output<String> id, @Nullable AppSecRatePolicyActionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecRatePolicyAction:AppSecRatePolicyAction", name, state, makeResourceOptions(options, id));
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
    public static AppSecRatePolicyAction get(String name, Output<String> id, @Nullable AppSecRatePolicyActionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppSecRatePolicyAction(name, id, state, options);
    }
}
