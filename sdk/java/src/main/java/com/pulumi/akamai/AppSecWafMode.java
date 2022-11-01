// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.akamai.AppSecWafModeArgs;
import com.pulumi.akamai.Utilities;
import com.pulumi.akamai.inputs.AppSecWafModeState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * **Scopes**: Security policy
 * 
 * Modifies the way your Kona Rule Set rules are updated.
 * Use **KRS** mode to update the rule sets manually or **AAG** to have those rule sets automatically updated.
 * 
 * **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/mode](https://techdocs.akamai.com/application-security/reference/put-policy-mode)
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
 * import com.pulumi.akamai.AppSecWafMode;
 * import com.pulumi.akamai.AppSecWafModeArgs;
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
 *         var wafMode = new AppSecWafMode(&#34;wafMode&#34;, AppSecWafModeArgs.builder()        
 *             .configId(configuration.applyValue(getAppSecConfigurationResult -&gt; getAppSecConfigurationResult.configId()))
 *             .securityPolicyId(&#34;gms1_134637&#34;)
 *             .mode(&#34;KRS&#34;)
 *             .build());
 * 
 *         ctx.export(&#34;wafModeMode&#34;, wafMode.mode());
 *         ctx.export(&#34;wafModeCurrentRuleset&#34;, wafMode.currentRuleset());
 *         ctx.export(&#34;wafModeEvalStatus&#34;, wafMode.evalStatus());
 *         ctx.export(&#34;wafModeEvalRuleset&#34;, wafMode.evalRuleset());
 *         ctx.export(&#34;wafModeEvalExpirationDate&#34;, wafMode.evalExpirationDate());
 *     }
 * }
 * ```
 * ## Output Options
 * 
 * The following options can be used to determine the information returned, and how that returned information is formatted:
 * 
 * - `current_ruleset` – Versioning information for the current Kona Rule Set.
 * - `eval_ruleset`. Versioning information for the Kona Rule Set being evaluated (if applicable).
 * - `eval_status`. Returns **enabled** if an evaluation is currently in progress; otherwise returns **disabled**.
 * - `eval_expiration_date`. Date on which the evaluation period ends (if applicable).
 * - `output_text`. Tabular report showing the current rule set, WAF mode and evaluation status.
 * 
 */
@ResourceType(type="akamai:index/appSecWafMode:AppSecWafMode")
public class AppSecWafMode extends com.pulumi.resources.CustomResource {
    /**
     * . Unique identifier of the security configuration associated with the WAF mode settings being modified.
     * 
     */
    @Export(name="configId", type=Integer.class, parameters={})
    private Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the WAF mode settings being modified.
     * 
     */
    public Output<Integer> configId() {
        return this.configId;
    }
    /**
     * Versioning information for the current Kona Rule Set
     * 
     */
    @Export(name="currentRuleset", type=String.class, parameters={})
    private Output<String> currentRuleset;

    /**
     * @return Versioning information for the current Kona Rule Set
     * 
     */
    public Output<String> currentRuleset() {
        return this.currentRuleset;
    }
    /**
     * Date on which the evaluation period ends, if applicable
     * 
     */
    @Export(name="evalExpirationDate", type=String.class, parameters={})
    private Output<String> evalExpirationDate;

    /**
     * @return Date on which the evaluation period ends, if applicable
     * 
     */
    public Output<String> evalExpirationDate() {
        return this.evalExpirationDate;
    }
    /**
     * Versioning information for the Kona Rule Set being evaluated, if applicable
     * 
     */
    @Export(name="evalRuleset", type=String.class, parameters={})
    private Output<String> evalRuleset;

    /**
     * @return Versioning information for the Kona Rule Set being evaluated, if applicable
     * 
     */
    public Output<String> evalRuleset() {
        return this.evalRuleset;
    }
    /**
     * Whether an evaluation is currently in progress
     * 
     */
    @Export(name="evalStatus", type=String.class, parameters={})
    private Output<String> evalStatus;

    /**
     * @return Whether an evaluation is currently in progress
     * 
     */
    public Output<String> evalStatus() {
        return this.evalStatus;
    }
    /**
     * . Specifies how Kona Rule Set rules are upgraded. Allowed values are:
     * 
     */
    @Export(name="mode", type=String.class, parameters={})
    private Output<String> mode;

    /**
     * @return . Specifies how Kona Rule Set rules are upgraded. Allowed values are:
     * 
     */
    public Output<String> mode() {
        return this.mode;
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
     * . Unique identifier of the security policy associated with the WAF mode settings being modified.
     * 
     */
    @Export(name="securityPolicyId", type=String.class, parameters={})
    private Output<String> securityPolicyId;

    /**
     * @return . Unique identifier of the security policy associated with the WAF mode settings being modified.
     * 
     */
    public Output<String> securityPolicyId() {
        return this.securityPolicyId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppSecWafMode(String name) {
        this(name, AppSecWafModeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppSecWafMode(String name, AppSecWafModeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppSecWafMode(String name, AppSecWafModeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecWafMode:AppSecWafMode", name, args == null ? AppSecWafModeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppSecWafMode(String name, Output<String> id, @Nullable AppSecWafModeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/appSecWafMode:AppSecWafMode", name, state, makeResourceOptions(options, id));
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
    public static AppSecWafMode get(String name, Output<String> id, @Nullable AppSecWafModeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppSecWafMode(name, id, state, options);
    }
}