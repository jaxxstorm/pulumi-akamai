// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EdgeWorkersActivationState extends com.pulumi.resources.ResourceArgs {

    public static final EdgeWorkersActivationState Empty = new EdgeWorkersActivationState();

    /**
     * (Required) Unique identifier of the activation.
     * 
     */
    @Import(name="activationId")
    private @Nullable Output<Integer> activationId;

    /**
     * @return (Required) Unique identifier of the activation.
     * 
     */
    public Optional<Output<Integer>> activationId() {
        return Optional.ofNullable(this.activationId);
    }

    /**
     * A unique identifier for the EdgeWorker ID you want to activate.
     * 
     */
    @Import(name="edgeworkerId")
    private @Nullable Output<Integer> edgeworkerId;

    /**
     * @return A unique identifier for the EdgeWorker ID you want to activate.
     * 
     */
    public Optional<Output<Integer>> edgeworkerId() {
        return Optional.ofNullable(this.edgeworkerId);
    }

    /**
     * The network you want to activate the policy version on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
     * 
     */
    @Import(name="network")
    private @Nullable Output<String> network;

    /**
     * @return The network you want to activate the policy version on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
     * 
     */
    public Optional<Output<String>> network() {
        return Optional.ofNullable(this.network);
    }

    /**
     * The EdgeWorker version you want to activate.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return The EdgeWorker version you want to activate.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private EdgeWorkersActivationState() {}

    private EdgeWorkersActivationState(EdgeWorkersActivationState $) {
        this.activationId = $.activationId;
        this.edgeworkerId = $.edgeworkerId;
        this.network = $.network;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EdgeWorkersActivationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EdgeWorkersActivationState $;

        public Builder() {
            $ = new EdgeWorkersActivationState();
        }

        public Builder(EdgeWorkersActivationState defaults) {
            $ = new EdgeWorkersActivationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param activationId (Required) Unique identifier of the activation.
         * 
         * @return builder
         * 
         */
        public Builder activationId(@Nullable Output<Integer> activationId) {
            $.activationId = activationId;
            return this;
        }

        /**
         * @param activationId (Required) Unique identifier of the activation.
         * 
         * @return builder
         * 
         */
        public Builder activationId(Integer activationId) {
            return activationId(Output.of(activationId));
        }

        /**
         * @param edgeworkerId A unique identifier for the EdgeWorker ID you want to activate.
         * 
         * @return builder
         * 
         */
        public Builder edgeworkerId(@Nullable Output<Integer> edgeworkerId) {
            $.edgeworkerId = edgeworkerId;
            return this;
        }

        /**
         * @param edgeworkerId A unique identifier for the EdgeWorker ID you want to activate.
         * 
         * @return builder
         * 
         */
        public Builder edgeworkerId(Integer edgeworkerId) {
            return edgeworkerId(Output.of(edgeworkerId));
        }

        /**
         * @param network The network you want to activate the policy version on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
         * 
         * @return builder
         * 
         */
        public Builder network(@Nullable Output<String> network) {
            $.network = network;
            return this;
        }

        /**
         * @param network The network you want to activate the policy version on. For the Staging network, specify either `STAGING`, `STAG`, or `S`. For the Production network, specify either `PRODUCTION`, `PROD`, or `P`. All values are case insensitive.
         * 
         * @return builder
         * 
         */
        public Builder network(String network) {
            return network(Output.of(network));
        }

        /**
         * @param version The EdgeWorker version you want to activate.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The EdgeWorker version you want to activate.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public EdgeWorkersActivationState build() {
            return $;
        }
    }

}
