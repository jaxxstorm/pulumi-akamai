// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.akamai.inputs.GtmResourceResourceInstanceArgs;
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


public final class GtmResourceState extends com.pulumi.resources.ResourceArgs {

    public static final GtmResourceState Empty = new GtmResourceState();

    /**
     * Specifies how GTM handles different load numbers when multiple load servers are used for a data center or property.
     * 
     */
    @Import(name="aggregationType")
    private @Nullable Output<String> aggregationType;

    /**
     * @return Specifies how GTM handles different load numbers when multiple load servers are used for a data center or property.
     * 
     */
    public Optional<Output<String>> aggregationType() {
        return Optional.ofNullable(this.aggregationType);
    }

    /**
     * Specifies the name of the property that this resource constrains, enter `**` to constrain all properties.
     * 
     */
    @Import(name="constrainedProperty")
    private @Nullable Output<String> constrainedProperty;

    /**
     * @return Specifies the name of the property that this resource constrains, enter `**` to constrain all properties.
     * 
     */
    public Optional<Output<String>> constrainedProperty() {
        return Optional.ofNullable(this.constrainedProperty);
    }

    /**
     * For Akamai internal use only. You can omit the value or set it to `null`.
     * 
     */
    @Import(name="decayRate")
    private @Nullable Output<Double> decayRate;

    /**
     * @return For Akamai internal use only. You can omit the value or set it to `null`.
     * 
     */
    public Optional<Output<Double>> decayRate() {
        return Optional.ofNullable(this.decayRate);
    }

    /**
     * A descriptive note to help you track what the resource constrains.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A descriptive note to help you track what the resource constrains.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * DNS name for the GTM Domain set that includes this property.
     * 
     */
    @Import(name="domain")
    private @Nullable Output<String> domain;

    /**
     * @return DNS name for the GTM Domain set that includes this property.
     * 
     */
    public Optional<Output<String>> domain() {
        return Optional.ofNullable(this.domain);
    }

    /**
     * Optionally specifies the host header used when fetching the load object.
     * 
     */
    @Import(name="hostHeader")
    private @Nullable Output<String> hostHeader;

    /**
     * @return Optionally specifies the host header used when fetching the load object.
     * 
     */
    public Optional<Output<String>> hostHeader() {
        return Optional.ofNullable(this.hostHeader);
    }

    /**
     * Specifies the text that comes before the `load_object`.
     * 
     */
    @Import(name="leaderString")
    private @Nullable Output<String> leaderString;

    /**
     * @return Specifies the text that comes before the `load_object`.
     * 
     */
    public Optional<Output<String>> leaderString() {
        return Optional.ofNullable(this.leaderString);
    }

    /**
     * For internal use only. Unless Akamai indicates otherwise, omit the value or set it to null.
     * 
     */
    @Import(name="leastSquaresDecay")
    private @Nullable Output<Double> leastSquaresDecay;

    /**
     * @return For internal use only. Unless Akamai indicates otherwise, omit the value or set it to null.
     * 
     */
    public Optional<Output<Double>> leastSquaresDecay() {
        return Optional.ofNullable(this.leastSquaresDecay);
    }

    @Import(name="loadImbalancePercentage")
    private @Nullable Output<Double> loadImbalancePercentage;

    public Optional<Output<Double>> loadImbalancePercentage() {
        return Optional.ofNullable(this.loadImbalancePercentage);
    }

    /**
     * For Akamai internal use only. You can omit the value or set it to `null`.
     * 
     */
    @Import(name="maxUMultiplicativeIncrement")
    private @Nullable Output<Double> maxUMultiplicativeIncrement;

    /**
     * @return For Akamai internal use only. You can omit the value or set it to `null`.
     * 
     */
    public Optional<Output<Double>> maxUMultiplicativeIncrement() {
        return Optional.ofNullable(this.maxUMultiplicativeIncrement);
    }

    /**
     * A descriptive label for the GTM resource.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A descriptive label for the GTM resource.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * (multiple allowed) Contains information about the resources that constrain the properties within the data center. You can have multiple `resource_instance` entries. Requires these arguments:
     * 
     */
    @Import(name="resourceInstances")
    private @Nullable Output<List<GtmResourceResourceInstanceArgs>> resourceInstances;

    /**
     * @return (multiple allowed) Contains information about the resources that constrain the properties within the data center. You can have multiple `resource_instance` entries. Requires these arguments:
     * 
     */
    public Optional<Output<List<GtmResourceResourceInstanceArgs>>> resourceInstances() {
        return Optional.ofNullable(this.resourceInstances);
    }

    /**
     * Indicates the kind of `load_object` format used to determine the load on the resource.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Indicates the kind of `load_object` format used to determine the load on the resource.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * An optional sanity check that specifies the maximum allowed value for any component of the load object.
     * 
     */
    @Import(name="upperBound")
    private @Nullable Output<Integer> upperBound;

    /**
     * @return An optional sanity check that specifies the maximum allowed value for any component of the load object.
     * 
     */
    public Optional<Output<Integer>> upperBound() {
        return Optional.ofNullable(this.upperBound);
    }

    /**
     * A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
     * 
     */
    @Import(name="waitOnComplete")
    private @Nullable Output<Boolean> waitOnComplete;

    /**
     * @return A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
     * 
     */
    public Optional<Output<Boolean>> waitOnComplete() {
        return Optional.ofNullable(this.waitOnComplete);
    }

    private GtmResourceState() {}

    private GtmResourceState(GtmResourceState $) {
        this.aggregationType = $.aggregationType;
        this.constrainedProperty = $.constrainedProperty;
        this.decayRate = $.decayRate;
        this.description = $.description;
        this.domain = $.domain;
        this.hostHeader = $.hostHeader;
        this.leaderString = $.leaderString;
        this.leastSquaresDecay = $.leastSquaresDecay;
        this.loadImbalancePercentage = $.loadImbalancePercentage;
        this.maxUMultiplicativeIncrement = $.maxUMultiplicativeIncrement;
        this.name = $.name;
        this.resourceInstances = $.resourceInstances;
        this.type = $.type;
        this.upperBound = $.upperBound;
        this.waitOnComplete = $.waitOnComplete;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GtmResourceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GtmResourceState $;

        public Builder() {
            $ = new GtmResourceState();
        }

        public Builder(GtmResourceState defaults) {
            $ = new GtmResourceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param aggregationType Specifies how GTM handles different load numbers when multiple load servers are used for a data center or property.
         * 
         * @return builder
         * 
         */
        public Builder aggregationType(@Nullable Output<String> aggregationType) {
            $.aggregationType = aggregationType;
            return this;
        }

        /**
         * @param aggregationType Specifies how GTM handles different load numbers when multiple load servers are used for a data center or property.
         * 
         * @return builder
         * 
         */
        public Builder aggregationType(String aggregationType) {
            return aggregationType(Output.of(aggregationType));
        }

        /**
         * @param constrainedProperty Specifies the name of the property that this resource constrains, enter `**` to constrain all properties.
         * 
         * @return builder
         * 
         */
        public Builder constrainedProperty(@Nullable Output<String> constrainedProperty) {
            $.constrainedProperty = constrainedProperty;
            return this;
        }

        /**
         * @param constrainedProperty Specifies the name of the property that this resource constrains, enter `**` to constrain all properties.
         * 
         * @return builder
         * 
         */
        public Builder constrainedProperty(String constrainedProperty) {
            return constrainedProperty(Output.of(constrainedProperty));
        }

        /**
         * @param decayRate For Akamai internal use only. You can omit the value or set it to `null`.
         * 
         * @return builder
         * 
         */
        public Builder decayRate(@Nullable Output<Double> decayRate) {
            $.decayRate = decayRate;
            return this;
        }

        /**
         * @param decayRate For Akamai internal use only. You can omit the value or set it to `null`.
         * 
         * @return builder
         * 
         */
        public Builder decayRate(Double decayRate) {
            return decayRate(Output.of(decayRate));
        }

        /**
         * @param description A descriptive note to help you track what the resource constrains.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A descriptive note to help you track what the resource constrains.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param domain DNS name for the GTM Domain set that includes this property.
         * 
         * @return builder
         * 
         */
        public Builder domain(@Nullable Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain DNS name for the GTM Domain set that includes this property.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param hostHeader Optionally specifies the host header used when fetching the load object.
         * 
         * @return builder
         * 
         */
        public Builder hostHeader(@Nullable Output<String> hostHeader) {
            $.hostHeader = hostHeader;
            return this;
        }

        /**
         * @param hostHeader Optionally specifies the host header used when fetching the load object.
         * 
         * @return builder
         * 
         */
        public Builder hostHeader(String hostHeader) {
            return hostHeader(Output.of(hostHeader));
        }

        /**
         * @param leaderString Specifies the text that comes before the `load_object`.
         * 
         * @return builder
         * 
         */
        public Builder leaderString(@Nullable Output<String> leaderString) {
            $.leaderString = leaderString;
            return this;
        }

        /**
         * @param leaderString Specifies the text that comes before the `load_object`.
         * 
         * @return builder
         * 
         */
        public Builder leaderString(String leaderString) {
            return leaderString(Output.of(leaderString));
        }

        /**
         * @param leastSquaresDecay For internal use only. Unless Akamai indicates otherwise, omit the value or set it to null.
         * 
         * @return builder
         * 
         */
        public Builder leastSquaresDecay(@Nullable Output<Double> leastSquaresDecay) {
            $.leastSquaresDecay = leastSquaresDecay;
            return this;
        }

        /**
         * @param leastSquaresDecay For internal use only. Unless Akamai indicates otherwise, omit the value or set it to null.
         * 
         * @return builder
         * 
         */
        public Builder leastSquaresDecay(Double leastSquaresDecay) {
            return leastSquaresDecay(Output.of(leastSquaresDecay));
        }

        public Builder loadImbalancePercentage(@Nullable Output<Double> loadImbalancePercentage) {
            $.loadImbalancePercentage = loadImbalancePercentage;
            return this;
        }

        public Builder loadImbalancePercentage(Double loadImbalancePercentage) {
            return loadImbalancePercentage(Output.of(loadImbalancePercentage));
        }

        /**
         * @param maxUMultiplicativeIncrement For Akamai internal use only. You can omit the value or set it to `null`.
         * 
         * @return builder
         * 
         */
        public Builder maxUMultiplicativeIncrement(@Nullable Output<Double> maxUMultiplicativeIncrement) {
            $.maxUMultiplicativeIncrement = maxUMultiplicativeIncrement;
            return this;
        }

        /**
         * @param maxUMultiplicativeIncrement For Akamai internal use only. You can omit the value or set it to `null`.
         * 
         * @return builder
         * 
         */
        public Builder maxUMultiplicativeIncrement(Double maxUMultiplicativeIncrement) {
            return maxUMultiplicativeIncrement(Output.of(maxUMultiplicativeIncrement));
        }

        /**
         * @param name A descriptive label for the GTM resource.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A descriptive label for the GTM resource.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param resourceInstances (multiple allowed) Contains information about the resources that constrain the properties within the data center. You can have multiple `resource_instance` entries. Requires these arguments:
         * 
         * @return builder
         * 
         */
        public Builder resourceInstances(@Nullable Output<List<GtmResourceResourceInstanceArgs>> resourceInstances) {
            $.resourceInstances = resourceInstances;
            return this;
        }

        /**
         * @param resourceInstances (multiple allowed) Contains information about the resources that constrain the properties within the data center. You can have multiple `resource_instance` entries. Requires these arguments:
         * 
         * @return builder
         * 
         */
        public Builder resourceInstances(List<GtmResourceResourceInstanceArgs> resourceInstances) {
            return resourceInstances(Output.of(resourceInstances));
        }

        /**
         * @param resourceInstances (multiple allowed) Contains information about the resources that constrain the properties within the data center. You can have multiple `resource_instance` entries. Requires these arguments:
         * 
         * @return builder
         * 
         */
        public Builder resourceInstances(GtmResourceResourceInstanceArgs... resourceInstances) {
            return resourceInstances(List.of(resourceInstances));
        }

        /**
         * @param type Indicates the kind of `load_object` format used to determine the load on the resource.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Indicates the kind of `load_object` format used to determine the load on the resource.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param upperBound An optional sanity check that specifies the maximum allowed value for any component of the load object.
         * 
         * @return builder
         * 
         */
        public Builder upperBound(@Nullable Output<Integer> upperBound) {
            $.upperBound = upperBound;
            return this;
        }

        /**
         * @param upperBound An optional sanity check that specifies the maximum allowed value for any component of the load object.
         * 
         * @return builder
         * 
         */
        public Builder upperBound(Integer upperBound) {
            return upperBound(Output.of(upperBound));
        }

        /**
         * @param waitOnComplete A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
         * 
         * @return builder
         * 
         */
        public Builder waitOnComplete(@Nullable Output<Boolean> waitOnComplete) {
            $.waitOnComplete = waitOnComplete;
            return this;
        }

        /**
         * @param waitOnComplete A boolean indicating whether to wait for transaction to complete. Set to `true` by default.
         * 
         * @return builder
         * 
         */
        public Builder waitOnComplete(Boolean waitOnComplete) {
            return waitOnComplete(Output.of(waitOnComplete));
        }

        public GtmResourceState build() {
            return $;
        }
    }

}
