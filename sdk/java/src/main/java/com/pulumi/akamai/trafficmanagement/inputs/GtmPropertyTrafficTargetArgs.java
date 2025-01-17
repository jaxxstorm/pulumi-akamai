// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.trafficmanagement.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GtmPropertyTrafficTargetArgs extends com.pulumi.resources.ResourceArgs {

    public static final GtmPropertyTrafficTargetArgs Empty = new GtmPropertyTrafficTargetArgs();

    /**
     * A unique identifier for an existing data center in the domain.
     * 
     */
    @Import(name="datacenterId")
    private @Nullable Output<Integer> datacenterId;

    /**
     * @return A unique identifier for an existing data center in the domain.
     * 
     */
    public Optional<Output<Integer>> datacenterId() {
        return Optional.ofNullable(this.datacenterId);
    }

    /**
     * A boolean indicating whether the traffic target is used. You can also omit the traffic target, which has the same result as the false value.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return A boolean indicating whether the traffic target is used. You can also omit the traffic target, which has the same result as the false value.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Specifies an optional data center for the property. Used when there are no servers configured for the property.
     * 
     */
    @Import(name="handoutCname")
    private @Nullable Output<String> handoutCname;

    /**
     * @return Specifies an optional data center for the property. Used when there are no servers configured for the property.
     * 
     */
    public Optional<Output<String>> handoutCname() {
        return Optional.ofNullable(this.handoutCname);
    }

    /**
     * Name of HTTP header.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of HTTP header.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * (List) Identifies the IP address or the hostnames of the servers.
     * 
     */
    @Import(name="servers")
    private @Nullable Output<List<String>> servers;

    /**
     * @return (List) Identifies the IP address or the hostnames of the servers.
     * 
     */
    public Optional<Output<List<String>>> servers() {
        return Optional.ofNullable(this.servers);
    }

    /**
     * Specifies the traffic weight for the target.
     * 
     */
    @Import(name="weight")
    private @Nullable Output<Double> weight;

    /**
     * @return Specifies the traffic weight for the target.
     * 
     */
    public Optional<Output<Double>> weight() {
        return Optional.ofNullable(this.weight);
    }

    private GtmPropertyTrafficTargetArgs() {}

    private GtmPropertyTrafficTargetArgs(GtmPropertyTrafficTargetArgs $) {
        this.datacenterId = $.datacenterId;
        this.enabled = $.enabled;
        this.handoutCname = $.handoutCname;
        this.name = $.name;
        this.servers = $.servers;
        this.weight = $.weight;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GtmPropertyTrafficTargetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GtmPropertyTrafficTargetArgs $;

        public Builder() {
            $ = new GtmPropertyTrafficTargetArgs();
        }

        public Builder(GtmPropertyTrafficTargetArgs defaults) {
            $ = new GtmPropertyTrafficTargetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param datacenterId A unique identifier for an existing data center in the domain.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(@Nullable Output<Integer> datacenterId) {
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
         * @param enabled A boolean indicating whether the traffic target is used. You can also omit the traffic target, which has the same result as the false value.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled A boolean indicating whether the traffic target is used. You can also omit the traffic target, which has the same result as the false value.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param handoutCname Specifies an optional data center for the property. Used when there are no servers configured for the property.
         * 
         * @return builder
         * 
         */
        public Builder handoutCname(@Nullable Output<String> handoutCname) {
            $.handoutCname = handoutCname;
            return this;
        }

        /**
         * @param handoutCname Specifies an optional data center for the property. Used when there are no servers configured for the property.
         * 
         * @return builder
         * 
         */
        public Builder handoutCname(String handoutCname) {
            return handoutCname(Output.of(handoutCname));
        }

        /**
         * @param name Name of HTTP header.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of HTTP header.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param servers (List) Identifies the IP address or the hostnames of the servers.
         * 
         * @return builder
         * 
         */
        public Builder servers(@Nullable Output<List<String>> servers) {
            $.servers = servers;
            return this;
        }

        /**
         * @param servers (List) Identifies the IP address or the hostnames of the servers.
         * 
         * @return builder
         * 
         */
        public Builder servers(List<String> servers) {
            return servers(Output.of(servers));
        }

        /**
         * @param servers (List) Identifies the IP address or the hostnames of the servers.
         * 
         * @return builder
         * 
         */
        public Builder servers(String... servers) {
            return servers(List.of(servers));
        }

        /**
         * @param weight Specifies the traffic weight for the target.
         * 
         * @return builder
         * 
         */
        public Builder weight(@Nullable Output<Double> weight) {
            $.weight = weight;
            return this;
        }

        /**
         * @param weight Specifies the traffic weight for the target.
         * 
         * @return builder
         * 
         */
        public Builder weight(Double weight) {
            return weight(Output.of(weight));
        }

        public GtmPropertyTrafficTargetArgs build() {
            return $;
        }
    }

}
