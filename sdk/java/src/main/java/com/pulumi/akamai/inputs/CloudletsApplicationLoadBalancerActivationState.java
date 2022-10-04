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


public final class CloudletsApplicationLoadBalancerActivationState extends com.pulumi.resources.ResourceArgs {

    public static final CloudletsApplicationLoadBalancerActivationState Empty = new CloudletsApplicationLoadBalancerActivationState();

    /**
     * The network you want to activate the policy version on, either `staging`, `stag`,  and `s` for the Staging network, or `production`, `prod`, and `p` for the Production network. All values are case insensitive.
     * 
     */
    @Import(name="network")
    private @Nullable Output<String> network;

    /**
     * @return The network you want to activate the policy version on, either `staging`, `stag`,  and `s` for the Staging network, or `production`, `prod`, and `p` for the Production network. All values are case insensitive.
     * 
     */
    public Optional<Output<String>> network() {
        return Optional.ofNullable(this.network);
    }

    /**
     * The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
     * 
     */
    @Import(name="originId")
    private @Nullable Output<String> originId;

    /**
     * @return The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
     * 
     */
    public Optional<Output<String>> originId() {
        return Optional.ofNullable(this.originId);
    }

    /**
     * The activation status for this load balancing configuration.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The activation status for this load balancing configuration.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The Application Load Balancer Cloudlet configuration version you want to activate.
     * 
     */
    @Import(name="version")
    private @Nullable Output<Integer> version;

    /**
     * @return The Application Load Balancer Cloudlet configuration version you want to activate.
     * 
     */
    public Optional<Output<Integer>> version() {
        return Optional.ofNullable(this.version);
    }

    private CloudletsApplicationLoadBalancerActivationState() {}

    private CloudletsApplicationLoadBalancerActivationState(CloudletsApplicationLoadBalancerActivationState $) {
        this.network = $.network;
        this.originId = $.originId;
        this.status = $.status;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudletsApplicationLoadBalancerActivationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudletsApplicationLoadBalancerActivationState $;

        public Builder() {
            $ = new CloudletsApplicationLoadBalancerActivationState();
        }

        public Builder(CloudletsApplicationLoadBalancerActivationState defaults) {
            $ = new CloudletsApplicationLoadBalancerActivationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param network The network you want to activate the policy version on, either `staging`, `stag`,  and `s` for the Staging network, or `production`, `prod`, and `p` for the Production network. All values are case insensitive.
         * 
         * @return builder
         * 
         */
        public Builder network(@Nullable Output<String> network) {
            $.network = network;
            return this;
        }

        /**
         * @param network The network you want to activate the policy version on, either `staging`, `stag`,  and `s` for the Staging network, or `production`, `prod`, and `p` for the Production network. All values are case insensitive.
         * 
         * @return builder
         * 
         */
        public Builder network(String network) {
            return network(Output.of(network));
        }

        /**
         * @param originId The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
         * 
         * @return builder
         * 
         */
        public Builder originId(@Nullable Output<String> originId) {
            $.originId = originId;
            return this;
        }

        /**
         * @param originId The identifier of an origin that represents the data center. The Conditional Origin, which is defined in Property Manager, must have an origin type of either `CUSTOMER` or `NET_STORAGE` set in the `origin` behavior. See property rules for more information.
         * 
         * @return builder
         * 
         */
        public Builder originId(String originId) {
            return originId(Output.of(originId));
        }

        /**
         * @param status The activation status for this load balancing configuration.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The activation status for this load balancing configuration.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param version The Application Load Balancer Cloudlet configuration version you want to activate.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<Integer> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The Application Load Balancer Cloudlet configuration version you want to activate.
         * 
         * @return builder
         * 
         */
        public Builder version(Integer version) {
            return version(Output.of(version));
        }

        public CloudletsApplicationLoadBalancerActivationState build() {
            return $;
        }
    }

}
