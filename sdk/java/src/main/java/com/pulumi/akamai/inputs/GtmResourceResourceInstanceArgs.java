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


public final class GtmResourceResourceInstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final GtmResourceResourceInstanceArgs Empty = new GtmResourceResourceInstanceArgs();

    /**
     * A unique identifier for an existing data center in the domain.
     * 
     */
    @Import(name="datacenterId", required=true)
    private Output<Integer> datacenterId;

    /**
     * @return A unique identifier for an existing data center in the domain.
     * 
     */
    public Output<Integer> datacenterId() {
        return this.datacenterId;
    }

    /**
     * Identifies the load object file used to report real-time information about the current load, maximum allowable load, and target load on each resource.
     * 
     */
    @Import(name="loadObject")
    private @Nullable Output<String> loadObject;

    /**
     * @return Identifies the load object file used to report real-time information about the current load, maximum allowable load, and target load on each resource.
     * 
     */
    public Optional<Output<String>> loadObject() {
        return Optional.ofNullable(this.loadObject);
    }

    /**
     * Specifies the TCP port of the `load_object`.
     * 
     */
    @Import(name="loadObjectPort")
    private @Nullable Output<Integer> loadObjectPort;

    /**
     * @return Specifies the TCP port of the `load_object`.
     * 
     */
    public Optional<Output<Integer>> loadObjectPort() {
        return Optional.ofNullable(this.loadObjectPort);
    }

    /**
     * (List) Specifies a list of servers from which to request the load object.
     * 
     */
    @Import(name="loadServers")
    private @Nullable Output<List<String>> loadServers;

    /**
     * @return (List) Specifies a list of servers from which to request the load object.
     * 
     */
    public Optional<Output<List<String>>> loadServers() {
        return Optional.ofNullable(this.loadServers);
    }

    /**
     * A boolean that indicates whether a default `load_object` is used for the resources.
     * 
     */
    @Import(name="useDefaultLoadObject")
    private @Nullable Output<Boolean> useDefaultLoadObject;

    /**
     * @return A boolean that indicates whether a default `load_object` is used for the resources.
     * 
     */
    public Optional<Output<Boolean>> useDefaultLoadObject() {
        return Optional.ofNullable(this.useDefaultLoadObject);
    }

    private GtmResourceResourceInstanceArgs() {}

    private GtmResourceResourceInstanceArgs(GtmResourceResourceInstanceArgs $) {
        this.datacenterId = $.datacenterId;
        this.loadObject = $.loadObject;
        this.loadObjectPort = $.loadObjectPort;
        this.loadServers = $.loadServers;
        this.useDefaultLoadObject = $.useDefaultLoadObject;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GtmResourceResourceInstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GtmResourceResourceInstanceArgs $;

        public Builder() {
            $ = new GtmResourceResourceInstanceArgs();
        }

        public Builder(GtmResourceResourceInstanceArgs defaults) {
            $ = new GtmResourceResourceInstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param datacenterId A unique identifier for an existing data center in the domain.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(Output<Integer> datacenterId) {
            $.datacenterId = datacenterId;
            return this;
        }

        /**
         * @param datacenterId A unique identifier for an existing data center in the domain.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(Integer datacenterId) {
            return datacenterId(Output.of(datacenterId));
        }

        /**
         * @param loadObject Identifies the load object file used to report real-time information about the current load, maximum allowable load, and target load on each resource.
         * 
         * @return builder
         * 
         */
        public Builder loadObject(@Nullable Output<String> loadObject) {
            $.loadObject = loadObject;
            return this;
        }

        /**
         * @param loadObject Identifies the load object file used to report real-time information about the current load, maximum allowable load, and target load on each resource.
         * 
         * @return builder
         * 
         */
        public Builder loadObject(String loadObject) {
            return loadObject(Output.of(loadObject));
        }

        /**
         * @param loadObjectPort Specifies the TCP port of the `load_object`.
         * 
         * @return builder
         * 
         */
        public Builder loadObjectPort(@Nullable Output<Integer> loadObjectPort) {
            $.loadObjectPort = loadObjectPort;
            return this;
        }

        /**
         * @param loadObjectPort Specifies the TCP port of the `load_object`.
         * 
         * @return builder
         * 
         */
        public Builder loadObjectPort(Integer loadObjectPort) {
            return loadObjectPort(Output.of(loadObjectPort));
        }

        /**
         * @param loadServers (List) Specifies a list of servers from which to request the load object.
         * 
         * @return builder
         * 
         */
        public Builder loadServers(@Nullable Output<List<String>> loadServers) {
            $.loadServers = loadServers;
            return this;
        }

        /**
         * @param loadServers (List) Specifies a list of servers from which to request the load object.
         * 
         * @return builder
         * 
         */
        public Builder loadServers(List<String> loadServers) {
            return loadServers(Output.of(loadServers));
        }

        /**
         * @param loadServers (List) Specifies a list of servers from which to request the load object.
         * 
         * @return builder
         * 
         */
        public Builder loadServers(String... loadServers) {
            return loadServers(List.of(loadServers));
        }

        /**
         * @param useDefaultLoadObject A boolean that indicates whether a default `load_object` is used for the resources.
         * 
         * @return builder
         * 
         */
        public Builder useDefaultLoadObject(@Nullable Output<Boolean> useDefaultLoadObject) {
            $.useDefaultLoadObject = useDefaultLoadObject;
            return this;
        }

        /**
         * @param useDefaultLoadObject A boolean that indicates whether a default `load_object` is used for the resources.
         * 
         * @return builder
         * 
         */
        public Builder useDefaultLoadObject(Boolean useDefaultLoadObject) {
            return useDefaultLoadObject(Output.of(useDefaultLoadObject));
        }

        public GtmResourceResourceInstanceArgs build() {
            $.datacenterId = Objects.requireNonNull($.datacenterId, "expected parameter 'datacenterId' to be non-null");
            return $;
        }
    }

}
