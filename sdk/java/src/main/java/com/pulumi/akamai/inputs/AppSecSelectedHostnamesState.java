// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AppSecSelectedHostnamesState extends com.pulumi.resources.ResourceArgs {

    public static final AppSecSelectedHostnamesState Empty = new AppSecSelectedHostnamesState();

    /**
     * . Unique identifier of the security configuration associated with the hostnames.
     * 
     */
    @Import(name="configId")
    private @Nullable Output<Integer> configId;

    /**
     * @return . Unique identifier of the security configuration associated with the hostnames.
     * 
     */
    public Optional<Output<Integer>> configId() {
        return Optional.ofNullable(this.configId);
    }

    /**
     * . JSON array of hostnames to be added or removed from the protected hosts list.
     * 
     */
    @Import(name="hostnames")
    private @Nullable Output<List<String>> hostnames;

    /**
     * @return . JSON array of hostnames to be added or removed from the protected hosts list.
     * 
     */
    public Optional<Output<List<String>>> hostnames() {
        return Optional.ofNullable(this.hostnames);
    }

    /**
     * . Indicates how the `hostnames` array is to be applied. Allowed values are:
     * - **APPEND**. Hosts listed in the `hostnames` array are added to the current list of selected hostnames.
     * - **REPLACE**. Hosts listed in the `hostnames`  array overwrite the current list of selected hostnames: the “old” hostnames are replaced by the specified set of hostnames.
     * - **REMOVE**, Hosts listed in the `hostnames` array are removed from the current list of select hostnames.
     * 
     */
    @Import(name="mode")
    private @Nullable Output<String> mode;

    /**
     * @return . Indicates how the `hostnames` array is to be applied. Allowed values are:
     * - **APPEND**. Hosts listed in the `hostnames` array are added to the current list of selected hostnames.
     * - **REPLACE**. Hosts listed in the `hostnames`  array overwrite the current list of selected hostnames: the “old” hostnames are replaced by the specified set of hostnames.
     * - **REMOVE**, Hosts listed in the `hostnames` array are removed from the current list of select hostnames.
     * 
     */
    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    private AppSecSelectedHostnamesState() {}

    private AppSecSelectedHostnamesState(AppSecSelectedHostnamesState $) {
        this.configId = $.configId;
        this.hostnames = $.hostnames;
        this.mode = $.mode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSecSelectedHostnamesState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSecSelectedHostnamesState $;

        public Builder() {
            $ = new AppSecSelectedHostnamesState();
        }

        public Builder(AppSecSelectedHostnamesState defaults) {
            $ = new AppSecSelectedHostnamesState(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the hostnames.
         * 
         * @return builder
         * 
         */
        public Builder configId(@Nullable Output<Integer> configId) {
            $.configId = configId;
            return this;
        }

        /**
         * @param configId . Unique identifier of the security configuration associated with the hostnames.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            return configId(Output.of(configId));
        }

        /**
         * @param hostnames . JSON array of hostnames to be added or removed from the protected hosts list.
         * 
         * @return builder
         * 
         */
        public Builder hostnames(@Nullable Output<List<String>> hostnames) {
            $.hostnames = hostnames;
            return this;
        }

        /**
         * @param hostnames . JSON array of hostnames to be added or removed from the protected hosts list.
         * 
         * @return builder
         * 
         */
        public Builder hostnames(List<String> hostnames) {
            return hostnames(Output.of(hostnames));
        }

        /**
         * @param hostnames . JSON array of hostnames to be added or removed from the protected hosts list.
         * 
         * @return builder
         * 
         */
        public Builder hostnames(String... hostnames) {
            return hostnames(List.of(hostnames));
        }

        /**
         * @param mode . Indicates how the `hostnames` array is to be applied. Allowed values are:
         * - **APPEND**. Hosts listed in the `hostnames` array are added to the current list of selected hostnames.
         * - **REPLACE**. Hosts listed in the `hostnames`  array overwrite the current list of selected hostnames: the “old” hostnames are replaced by the specified set of hostnames.
         * - **REMOVE**, Hosts listed in the `hostnames` array are removed from the current list of select hostnames.
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode . Indicates how the `hostnames` array is to be applied. Allowed values are:
         * - **APPEND**. Hosts listed in the `hostnames` array are added to the current list of selected hostnames.
         * - **REPLACE**. Hosts listed in the `hostnames`  array overwrite the current list of selected hostnames: the “old” hostnames are replaced by the specified set of hostnames.
         * - **REMOVE**, Hosts listed in the `hostnames` array are removed from the current list of select hostnames.
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        public AppSecSelectedHostnamesState build() {
            return $;
        }
    }

}
