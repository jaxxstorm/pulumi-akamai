// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AppSecAdvancedSettingsPrefetchState extends com.pulumi.resources.ResourceArgs {

    public static final AppSecAdvancedSettingsPrefetchState Empty = new AppSecAdvancedSettingsPrefetchState();

    /**
     * . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
     * 
     */
    @Import(name="allExtensions")
    private @Nullable Output<Boolean> allExtensions;

    /**
     * @return . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
     * 
     */
    public Optional<Output<Boolean>> allExtensions() {
        return Optional.ofNullable(this.allExtensions);
    }

    /**
     * . Unique identifier of the security configuration associated with the prefetch settings being modified.
     * 
     */
    @Import(name="configId")
    private @Nullable Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the prefetch settings being modified.
     * 
     */
    public Optional<Output<Integer>> configId() {
        return Optional.ofNullable(this.configId);
    }

    /**
     * . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
     * 
     */
    @Import(name="enableAppLayer")
    private @Nullable Output<Boolean> enableAppLayer;

    /**
     * @return . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
     * 
     */
    public Optional<Output<Boolean>> enableAppLayer() {
        return Optional.ofNullable(this.enableAppLayer);
    }

    /**
     * . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
     * 
     */
    @Import(name="enableRateControls")
    private @Nullable Output<Boolean> enableRateControls;

    /**
     * @return . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
     * 
     */
    public Optional<Output<Boolean>> enableRateControls() {
        return Optional.ofNullable(this.enableRateControls);
    }

    /**
     * . If `all_extensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won&#39;t be used with any file extensions not included in the array. If `all_extensions` is **true**, then this argument must be set to an empty array: **[]**.
     * 
     */
    @Import(name="extensions")
    private @Nullable Output<List<String>> extensions;

    /**
     * @return . If `all_extensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won&#39;t be used with any file extensions not included in the array. If `all_extensions` is **true**, then this argument must be set to an empty array: **[]**.
     * 
     */
    public Optional<Output<List<String>>> extensions() {
        return Optional.ofNullable(this.extensions);
    }

    private AppSecAdvancedSettingsPrefetchState() {}

    private AppSecAdvancedSettingsPrefetchState(AppSecAdvancedSettingsPrefetchState $) {
        this.allExtensions = $.allExtensions;
        this.configId = $.configId;
        this.enableAppLayer = $.enableAppLayer;
        this.enableRateControls = $.enableRateControls;
        this.extensions = $.extensions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecAdvancedSettingsPrefetchState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecAdvancedSettingsPrefetchState $;

        public Builder() {
            $ = new AppSecAdvancedSettingsPrefetchState();
        }

        public Builder(AppSecAdvancedSettingsPrefetchState defaults) {
            $ = new AppSecAdvancedSettingsPrefetchState(Objects.requireNonNull(defaults));
        }

        /**
         * @param allExtensions . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
         * 
         * @return builder
         * 
         */
        public Builder allExtensions(@Nullable Output<Boolean> allExtensions) {
            $.allExtensions = allExtensions;
            return this;
        }

        /**
         * @param allExtensions . Set to **true** to enable prefetch requests for all file extensions; set to **false** to enable prefetch requests on only a specified set of file extensions. If set to false you must include the `extensions` argument.
         * 
         * @return builder
         * 
         */
        public Builder allExtensions(Boolean allExtensions) {
            return allExtensions(Output.of(allExtensions));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the prefetch settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(@Nullable Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the prefetch settings being modified.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param enableAppLayer . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
         * 
         * @return builder
         * 
         */
        public Builder enableAppLayer(@Nullable Output<Boolean> enableAppLayer) {
            $.enableAppLayer = enableAppLayer;
            return this;
        }

        /**
         * @param enableAppLayer . Set to **true** to enable prefetch requests; set to **false** to disable prefetch requests.
         * 
         * @return builder
         * 
         */
        public Builder enableAppLayer(Boolean enableAppLayer) {
            return enableAppLayer(Output.of(enableAppLayer));
        }

        /**
         * @param enableRateControls . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
         * 
         * @return builder
         * 
         */
        public Builder enableRateControls(@Nullable Output<Boolean> enableRateControls) {
            $.enableRateControls = enableRateControls;
            return this;
        }

        /**
         * @param enableRateControls . Set to **true** to enable prefetch requests for rate controls; set to **false** to disable prefetch requests for rate controls.
         * 
         * @return builder
         * 
         */
        public Builder enableRateControls(Boolean enableRateControls) {
            return enableRateControls(Output.of(enableRateControls));
        }

        /**
         * @param extensions . If `all_extensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won&#39;t be used with any file extensions not included in the array. If `all_extensions` is **true**, then this argument must be set to an empty array: **[]**.
         * 
         * @return builder
         * 
         */
        public Builder extensions(@Nullable Output<List<String>> extensions) {
            $.extensions = extensions;
            return this;
        }

        /**
         * @param extensions . If `all_extensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won&#39;t be used with any file extensions not included in the array. If `all_extensions` is **true**, then this argument must be set to an empty array: **[]**.
         * 
         * @return builder
         * 
         */
        public Builder extensions(List<String> extensions) {
            return extensions(Output.of(extensions));
        }

        /**
         * @param extensions . If `all_extensions` is **false**, this must be a JSON array of all the file extensions for which prefetch requests are enabled: prefetch requests won&#39;t be used with any file extensions not included in the array. If `all_extensions` is **true**, then this argument must be set to an empty array: **[]**.
         * 
         * @return builder
         * 
         */
        public Builder extensions(String... extensions) {
            return extensions(List.of(extensions));
        }

        public AppSecAdvancedSettingsPrefetchState build() {
            return $;
        }
    }

}
