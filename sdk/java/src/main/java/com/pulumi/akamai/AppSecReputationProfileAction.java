// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.akamai.AppSecReputationProfileActionArgs;
import com.pulumi.akamai.Utilities;
import com.pulumi.akamai.inputs.AppSecReputationProfileActionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * **Scopes**: Reputation profile
 * 
 * Modifies the action taken when a reputation profile is triggered.
 * 
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/reputation-profiles/{reputationProfileId}](https://techdocs.akamai.com/application-security/reference/put-reputation-profile-action)
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
 * import com.pulumi.akamai.AppSecReputationProfileAction;
 * import com.pulumi.akamai.AppSecReputationProfileActionArgs;
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
 *         var appsecReputationProfileAction = new AppSecReputationProfileAction(&#34;appsecReputationProfileAction&#34;, AppSecReputationProfileActionArgs.builder()        
 *             .configId(configuration.applyValue(getAppSecConfigurationResult -&gt; getAppSecConfigurationResult.configId()))
 *             .securityPolicyId(&#34;gms1_134637&#34;)
 *             .reputationProfileId(130713)
 *             .action(&#34;alert&#34;)
 *             .build());
 * 
 *         ctx.export(&#34;reputationProfileId&#34;, appsecReputationProfileAction.reputationProfileId());
 *         ctx.export(&#34;reputationProfileAction&#34;, appsecReputationProfileAction.action());
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="akamai:index/appSecReputationProfileAction:AppSecReputationProfileAction")
public class AppSecReputationProfileAction extends com.pulumi.resources.CustomResource {
    /**
     * . Action taken any time the reputation profile is triggered. Allows values are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     * 
     */
    @Export(name="action", type=String.class, parameters={})
    private Output<String> action;

    /**
     * @return . Action taken any time the reputation profile is triggered. Allows values are:
     * - **alert**. Record the event.
     * - **deny**. Block the request.
     * - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
     * - **none**. Take no action.
     * 
     */
    public Output<String> action() {
        return this.action;
    }
    /**
     * . Unique identifier of the security configuration associated with the reputation profile action being modified.
     * 
     */
    @Export(name="configId", type=Integer.class, parameters={})
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the reputation profile action being modified.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }
    /**
     * . Unique identifier of the reputation profile whose action is being modified.
     * 
     */
    @Export(name="reputationProfileId", type=Integer.class, parameters={})
    private Output<Integer> reputationProfileId;

    /**
     * @return . Unique identifier of the reputation profile whose action is being modified.
     * 
     */
    public Output<Integer> reputationProfileId() {
        return this.reputationProfileId;
    }
    /**
     * . Unique identifier of the security policy associated with the reputation profile action being modified.
     * 
     */
    @Export(name="securityPolicyId", type=String.class, parameters={})
    private Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the reputation profile action being modified.
     * 
     */
    public Output<String> securityPolicyId() {
        return this.securityPolicyId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppSecReputationProfileAction(String name) {
        this(name, AppSecReputationProfileActionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppSecReputationProfileAction(String name, AppSecReputationProfileActionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppSecReputationProfileAction(String name, AppSecReputationProfileActionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecReputationProfileAction:AppSecReputationProfileAction", name, args == null ? AppSecReputationProfileActionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppSecReputationProfileAction(String name, Output<String> id, @Nullable AppSecReputationProfileActionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecReputationProfileAction:AppSecReputationProfileAction", name, state, makeResourceOptions(options, id));
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
    public static AppSecReputationProfileAction get(String name, Output<String> id, @Nullable AppSecReputationProfileActionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppSecReputationProfileAction(name, id, state, options);
    }
}
