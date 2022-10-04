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


public final class GtmDomainState extends com.pulumi.resources.ResourceArgs {

    public static final GtmDomainState Empty = new GtmDomainState();

    /**
     * A boolean that if set to `true`, GTM collapses CNAME redirections in DNS answers when it knows the target of the CNAME.
     * 
     */
    @Import(name="cnameCoalescingEnabled")
    private @Nullable Output<Boolean> cnameCoalescingEnabled;

    /**
     * @return A boolean that if set to `true`, GTM collapses CNAME redirections in DNS answers when it knows the target of the CNAME.
     * 
     */
    public Optional<Output<Boolean>> cnameCoalescingEnabled() {
        return Optional.ofNullable(this.cnameCoalescingEnabled);
    }

    /**
     * A descriptive note about changes to the domain. The maximum is 4000 characters.
     * 
     */
    @Import(name="comment")
    private @Nullable Output<String> comment;

    /**
     * @return A descriptive note about changes to the domain. The maximum is 4000 characters.
     * 
     */
    public Optional<Output<String>> comment() {
        return Optional.ofNullable(this.comment);
    }

    /**
     * If creating a domain, the contract ID.
     * 
     */
    @Import(name="contract")
    private @Nullable Output<String> contract;

    /**
     * @return If creating a domain, the contract ID.
     * 
     */
    public Optional<Output<String>> contract() {
        return Optional.ofNullable(this.contract);
    }

    /**
     * Specifies the download penalty score. The default is `75`. If the download encounters an error, the web agent computes a score that is either the download time in seconds or a penalty score.
     * 
     */
    @Import(name="defaultErrorPenalty")
    private @Nullable Output<Integer> defaultErrorPenalty;

    /**
     * @return Specifies the download penalty score. The default is `75`. If the download encounters an error, the web agent computes a score that is either the download time in seconds or a penalty score.
     * 
     */
    public Optional<Output<Integer>> defaultErrorPenalty() {
        return Optional.ofNullable(this.defaultErrorPenalty);
    }

    @Import(name="defaultHealthMax")
    private @Nullable Output<Double> defaultHealthMax;

    public Optional<Output<Double>> defaultHealthMax() {
        return Optional.ofNullable(this.defaultHealthMax);
    }

    @Import(name="defaultHealthMultiplier")
    private @Nullable Output<Double> defaultHealthMultiplier;

    public Optional<Output<Double>> defaultHealthMultiplier() {
        return Optional.ofNullable(this.defaultHealthMultiplier);
    }

    @Import(name="defaultHealthThreshold")
    private @Nullable Output<Double> defaultHealthThreshold;

    public Optional<Output<Double>> defaultHealthThreshold() {
        return Optional.ofNullable(this.defaultHealthThreshold);
    }

    @Import(name="defaultMaxUnreachablePenalty")
    private @Nullable Output<Integer> defaultMaxUnreachablePenalty;

    public Optional<Output<Integer>> defaultMaxUnreachablePenalty() {
        return Optional.ofNullable(this.defaultMaxUnreachablePenalty);
    }

    /**
     * Specifies an optional Base64-encoded certificate that corresponds with the private key for TLS-based liveness tests (HTTPS, SMTPS, POPS, and TCPS).
     * 
     */
    @Import(name="defaultSslClientCertificate")
    private @Nullable Output<String> defaultSslClientCertificate;

    /**
     * @return Specifies an optional Base64-encoded certificate that corresponds with the private key for TLS-based liveness tests (HTTPS, SMTPS, POPS, and TCPS).
     * 
     */
    public Optional<Output<String>> defaultSslClientCertificate() {
        return Optional.ofNullable(this.defaultSslClientCertificate);
    }

    /**
     * Specifies a Base64-encoded private key that corresponds with the TLS certificate for HTTPS, SMTPS, POPS, and TCPS liveness tests.
     * 
     */
    @Import(name="defaultSslClientPrivateKey")
    private @Nullable Output<String> defaultSslClientPrivateKey;

    /**
     * @return Specifies a Base64-encoded private key that corresponds with the TLS certificate for HTTPS, SMTPS, POPS, and TCPS liveness tests.
     * 
     */
    public Optional<Output<String>> defaultSslClientPrivateKey() {
        return Optional.ofNullable(this.defaultSslClientPrivateKey);
    }

    /**
     * Specifies the timeout penalty score. Default is `25`.
     * 
     */
    @Import(name="defaultTimeoutPenalty")
    private @Nullable Output<Integer> defaultTimeoutPenalty;

    /**
     * @return Specifies the timeout penalty score. Default is `25`.
     * 
     */
    public Optional<Output<Integer>> defaultTimeoutPenalty() {
        return Optional.ofNullable(this.defaultTimeoutPenalty);
    }

    @Import(name="defaultUnreachableThreshold")
    private @Nullable Output<Double> defaultUnreachableThreshold;

    public Optional<Output<Double>> defaultUnreachableThreshold() {
        return Optional.ofNullable(this.defaultUnreachableThreshold);
    }

    /**
     * A list of email addresses to notify when a change is made to the domain.
     * 
     */
    @Import(name="emailNotificationLists")
    private @Nullable Output<List<String>> emailNotificationLists;

    /**
     * @return A list of email addresses to notify when a change is made to the domain.
     * 
     */
    public Optional<Output<List<String>>> emailNotificationLists() {
        return Optional.ofNullable(this.emailNotificationLists);
    }

    /**
     * A boolean indicating whether whether the GTM Domain is using end user client subnet mapping.
     * 
     */
    @Import(name="endUserMappingEnabled")
    private @Nullable Output<Boolean> endUserMappingEnabled;

    /**
     * @return A boolean indicating whether whether the GTM Domain is using end user client subnet mapping.
     * 
     */
    public Optional<Output<Boolean>> endUserMappingEnabled() {
        return Optional.ofNullable(this.endUserMappingEnabled);
    }

    /**
     * If creating a domain, the currently selected group ID.
     * 
     */
    @Import(name="group")
    private @Nullable Output<String> group;

    /**
     * @return If creating a domain, the currently selected group ID.
     * 
     */
    public Optional<Output<String>> group() {
        return Optional.ofNullable(this.group);
    }

    /**
     * A boolean indicating whether one or more measurements of load (resources) are defined by you and supplied by each data center in real time to balance load.
     * 
     */
    @Import(name="loadFeedback")
    private @Nullable Output<Boolean> loadFeedback;

    /**
     * @return A boolean indicating whether one or more measurements of load (resources) are defined by you and supplied by each data center in real time to balance load.
     * 
     */
    public Optional<Output<Boolean>> loadFeedback() {
        return Optional.ofNullable(this.loadFeedback);
    }

    /**
     * Indicates the percentage of load imbalance factor (LIF) for the domain.
     * 
     */
    @Import(name="loadImbalancePercentage")
    private @Nullable Output<Double> loadImbalancePercentage;

    /**
     * @return Indicates the percentage of load imbalance factor (LIF) for the domain.
     * 
     */
    public Optional<Output<Double>> loadImbalancePercentage() {
        return Optional.ofNullable(this.loadImbalancePercentage);
    }

    @Import(name="mapUpdateInterval")
    private @Nullable Output<Integer> mapUpdateInterval;

    public Optional<Output<Integer>> mapUpdateInterval() {
        return Optional.ofNullable(this.mapUpdateInterval);
    }

    @Import(name="maxProperties")
    private @Nullable Output<Integer> maxProperties;

    public Optional<Output<Integer>> maxProperties() {
        return Optional.ofNullable(this.maxProperties);
    }

    @Import(name="maxResources")
    private @Nullable Output<Integer> maxResources;

    public Optional<Output<Integer>> maxResources() {
        return Optional.ofNullable(this.maxResources);
    }

    @Import(name="maxTestTimeout")
    private @Nullable Output<Double> maxTestTimeout;

    public Optional<Output<Double>> maxTestTimeout() {
        return Optional.ofNullable(this.maxTestTimeout);
    }

    @Import(name="maxTtl")
    private @Nullable Output<Integer> maxTtl;

    public Optional<Output<Integer>> maxTtl() {
        return Optional.ofNullable(this.maxTtl);
    }

    @Import(name="minPingableRegionFraction")
    private @Nullable Output<Double> minPingableRegionFraction;

    public Optional<Output<Double>> minPingableRegionFraction() {
        return Optional.ofNullable(this.minPingableRegionFraction);
    }

    @Import(name="minTestInterval")
    private @Nullable Output<Integer> minTestInterval;

    public Optional<Output<Integer>> minTestInterval() {
        return Optional.ofNullable(this.minTestInterval);
    }

    @Import(name="minTtl")
    private @Nullable Output<Integer> minTtl;

    public Optional<Output<Integer>> minTtl() {
        return Optional.ofNullable(this.minTtl);
    }

    /**
     * The DNS name for a collection of GTM Properties.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The DNS name for a collection of GTM Properties.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="pingInterval")
    private @Nullable Output<Integer> pingInterval;

    public Optional<Output<Integer>> pingInterval() {
        return Optional.ofNullable(this.pingInterval);
    }

    @Import(name="pingPacketSize")
    private @Nullable Output<Integer> pingPacketSize;

    public Optional<Output<Integer>> pingPacketSize() {
        return Optional.ofNullable(this.pingPacketSize);
    }

    @Import(name="roundRobinPrefix")
    private @Nullable Output<String> roundRobinPrefix;

    public Optional<Output<String>> roundRobinPrefix() {
        return Optional.ofNullable(this.roundRobinPrefix);
    }

    @Import(name="servermonitorLivenessCount")
    private @Nullable Output<Integer> servermonitorLivenessCount;

    public Optional<Output<Integer>> servermonitorLivenessCount() {
        return Optional.ofNullable(this.servermonitorLivenessCount);
    }

    @Import(name="servermonitorLoadCount")
    private @Nullable Output<Integer> servermonitorLoadCount;

    public Optional<Output<Integer>> servermonitorLoadCount() {
        return Optional.ofNullable(this.servermonitorLoadCount);
    }

    @Import(name="servermonitorPool")
    private @Nullable Output<String> servermonitorPool;

    public Optional<Output<String>> servermonitorPool() {
        return Optional.ofNullable(this.servermonitorPool);
    }

    /**
     * Th type of GTM domain. Options include `failover-only`, `static`, `weighted`, `basic`, or `full`.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Th type of GTM domain. Options include `failover-only`, `static`, `weighted`, `basic`, or `full`.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * A boolean that, if set to `true`, waits for transaction to complete.
     * 
     */
    @Import(name="waitOnComplete")
    private @Nullable Output<Boolean> waitOnComplete;

    /**
     * @return A boolean that, if set to `true`, waits for transaction to complete.
     * 
     */
    public Optional<Output<Boolean>> waitOnComplete() {
        return Optional.ofNullable(this.waitOnComplete);
    }

    private GtmDomainState() {}

    private GtmDomainState(GtmDomainState $) {
        this.cnameCoalescingEnabled = $.cnameCoalescingEnabled;
        this.comment = $.comment;
        this.contract = $.contract;
        this.defaultErrorPenalty = $.defaultErrorPenalty;
        this.defaultHealthMax = $.defaultHealthMax;
        this.defaultHealthMultiplier = $.defaultHealthMultiplier;
        this.defaultHealthThreshold = $.defaultHealthThreshold;
        this.defaultMaxUnreachablePenalty = $.defaultMaxUnreachablePenalty;
        this.defaultSslClientCertificate = $.defaultSslClientCertificate;
        this.defaultSslClientPrivateKey = $.defaultSslClientPrivateKey;
        this.defaultTimeoutPenalty = $.defaultTimeoutPenalty;
        this.defaultUnreachableThreshold = $.defaultUnreachableThreshold;
        this.emailNotificationLists = $.emailNotificationLists;
        this.endUserMappingEnabled = $.endUserMappingEnabled;
        this.group = $.group;
        this.loadFeedback = $.loadFeedback;
        this.loadImbalancePercentage = $.loadImbalancePercentage;
        this.mapUpdateInterval = $.mapUpdateInterval;
        this.maxProperties = $.maxProperties;
        this.maxResources = $.maxResources;
        this.maxTestTimeout = $.maxTestTimeout;
        this.maxTtl = $.maxTtl;
        this.minPingableRegionFraction = $.minPingableRegionFraction;
        this.minTestInterval = $.minTestInterval;
        this.minTtl = $.minTtl;
        this.name = $.name;
        this.pingInterval = $.pingInterval;
        this.pingPacketSize = $.pingPacketSize;
        this.roundRobinPrefix = $.roundRobinPrefix;
        this.servermonitorLivenessCount = $.servermonitorLivenessCount;
        this.servermonitorLoadCount = $.servermonitorLoadCount;
        this.servermonitorPool = $.servermonitorPool;
        this.type = $.type;
        this.waitOnComplete = $.waitOnComplete;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GtmDomainState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GtmDomainState $;

        public Builder() {
            $ = new GtmDomainState();
        }

        public Builder(GtmDomainState defaults) {
            $ = new GtmDomainState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cnameCoalescingEnabled A boolean that if set to `true`, GTM collapses CNAME redirections in DNS answers when it knows the target of the CNAME.
         * 
         * @return builder
         * 
         */
        public Builder cnameCoalescingEnabled(@Nullable Output<Boolean> cnameCoalescingEnabled) {
            $.cnameCoalescingEnabled = cnameCoalescingEnabled;
            return this;
        }

        /**
         * @param cnameCoalescingEnabled A boolean that if set to `true`, GTM collapses CNAME redirections in DNS answers when it knows the target of the CNAME.
         * 
         * @return builder
         * 
         */
        public Builder cnameCoalescingEnabled(Boolean cnameCoalescingEnabled) {
            return cnameCoalescingEnabled(Output.of(cnameCoalescingEnabled));
        }

        /**
         * @param comment A descriptive note about changes to the domain. The maximum is 4000 characters.
         * 
         * @return builder
         * 
         */
        public Builder comment(@Nullable Output<String> comment) {
            $.comment = comment;
            return this;
        }

        /**
         * @param comment A descriptive note about changes to the domain. The maximum is 4000 characters.
         * 
         * @return builder
         * 
         */
        public Builder comment(String comment) {
            return comment(Output.of(comment));
        }

        /**
         * @param contract If creating a domain, the contract ID.
         * 
         * @return builder
         * 
         */
        public Builder contract(@Nullable Output<String> contract) {
            $.contract = contract;
            return this;
        }

        /**
         * @param contract If creating a domain, the contract ID.
         * 
         * @return builder
         * 
         */
        public Builder contract(String contract) {
            return contract(Output.of(contract));
        }

        /**
         * @param defaultErrorPenalty Specifies the download penalty score. The default is `75`. If the download encounters an error, the web agent computes a score that is either the download time in seconds or a penalty score.
         * 
         * @return builder
         * 
         */
        public Builder defaultErrorPenalty(@Nullable Output<Integer> defaultErrorPenalty) {
            $.defaultErrorPenalty = defaultErrorPenalty;
            return this;
        }

        /**
         * @param defaultErrorPenalty Specifies the download penalty score. The default is `75`. If the download encounters an error, the web agent computes a score that is either the download time in seconds or a penalty score.
         * 
         * @return builder
         * 
         */
        public Builder defaultErrorPenalty(Integer defaultErrorPenalty) {
            return defaultErrorPenalty(Output.of(defaultErrorPenalty));
        }

        public Builder defaultHealthMax(@Nullable Output<Double> defaultHealthMax) {
            $.defaultHealthMax = defaultHealthMax;
            return this;
        }

        public Builder defaultHealthMax(Double defaultHealthMax) {
            return defaultHealthMax(Output.of(defaultHealthMax));
        }

        public Builder defaultHealthMultiplier(@Nullable Output<Double> defaultHealthMultiplier) {
            $.defaultHealthMultiplier = defaultHealthMultiplier;
            return this;
        }

        public Builder defaultHealthMultiplier(Double defaultHealthMultiplier) {
            return defaultHealthMultiplier(Output.of(defaultHealthMultiplier));
        }

        public Builder defaultHealthThreshold(@Nullable Output<Double> defaultHealthThreshold) {
            $.defaultHealthThreshold = defaultHealthThreshold;
            return this;
        }

        public Builder defaultHealthThreshold(Double defaultHealthThreshold) {
            return defaultHealthThreshold(Output.of(defaultHealthThreshold));
        }

        public Builder defaultMaxUnreachablePenalty(@Nullable Output<Integer> defaultMaxUnreachablePenalty) {
            $.defaultMaxUnreachablePenalty = defaultMaxUnreachablePenalty;
            return this;
        }

        public Builder defaultMaxUnreachablePenalty(Integer defaultMaxUnreachablePenalty) {
            return defaultMaxUnreachablePenalty(Output.of(defaultMaxUnreachablePenalty));
        }

        /**
         * @param defaultSslClientCertificate Specifies an optional Base64-encoded certificate that corresponds with the private key for TLS-based liveness tests (HTTPS, SMTPS, POPS, and TCPS).
         * 
         * @return builder
         * 
         */
        public Builder defaultSslClientCertificate(@Nullable Output<String> defaultSslClientCertificate) {
            $.defaultSslClientCertificate = defaultSslClientCertificate;
            return this;
        }

        /**
         * @param defaultSslClientCertificate Specifies an optional Base64-encoded certificate that corresponds with the private key for TLS-based liveness tests (HTTPS, SMTPS, POPS, and TCPS).
         * 
         * @return builder
         * 
         */
        public Builder defaultSslClientCertificate(String defaultSslClientCertificate) {
            return defaultSslClientCertificate(Output.of(defaultSslClientCertificate));
        }

        /**
         * @param defaultSslClientPrivateKey Specifies a Base64-encoded private key that corresponds with the TLS certificate for HTTPS, SMTPS, POPS, and TCPS liveness tests.
         * 
         * @return builder
         * 
         */
        public Builder defaultSslClientPrivateKey(@Nullable Output<String> defaultSslClientPrivateKey) {
            $.defaultSslClientPrivateKey = defaultSslClientPrivateKey;
            return this;
        }

        /**
         * @param defaultSslClientPrivateKey Specifies a Base64-encoded private key that corresponds with the TLS certificate for HTTPS, SMTPS, POPS, and TCPS liveness tests.
         * 
         * @return builder
         * 
         */
        public Builder defaultSslClientPrivateKey(String defaultSslClientPrivateKey) {
            return defaultSslClientPrivateKey(Output.of(defaultSslClientPrivateKey));
        }

        /**
         * @param defaultTimeoutPenalty Specifies the timeout penalty score. Default is `25`.
         * 
         * @return builder
         * 
         */
        public Builder defaultTimeoutPenalty(@Nullable Output<Integer> defaultTimeoutPenalty) {
            $.defaultTimeoutPenalty = defaultTimeoutPenalty;
            return this;
        }

        /**
         * @param defaultTimeoutPenalty Specifies the timeout penalty score. Default is `25`.
         * 
         * @return builder
         * 
         */
        public Builder defaultTimeoutPenalty(Integer defaultTimeoutPenalty) {
            return defaultTimeoutPenalty(Output.of(defaultTimeoutPenalty));
        }

        public Builder defaultUnreachableThreshold(@Nullable Output<Double> defaultUnreachableThreshold) {
            $.defaultUnreachableThreshold = defaultUnreachableThreshold;
            return this;
        }

        public Builder defaultUnreachableThreshold(Double defaultUnreachableThreshold) {
            return defaultUnreachableThreshold(Output.of(defaultUnreachableThreshold));
        }

        /**
         * @param emailNotificationLists A list of email addresses to notify when a change is made to the domain.
         * 
         * @return builder
         * 
         */
        public Builder emailNotificationLists(@Nullable Output<List<String>> emailNotificationLists) {
            $.emailNotificationLists = emailNotificationLists;
            return this;
        }

        /**
         * @param emailNotificationLists A list of email addresses to notify when a change is made to the domain.
         * 
         * @return builder
         * 
         */
        public Builder emailNotificationLists(List<String> emailNotificationLists) {
            return emailNotificationLists(Output.of(emailNotificationLists));
        }

        /**
         * @param emailNotificationLists A list of email addresses to notify when a change is made to the domain.
         * 
         * @return builder
         * 
         */
        public Builder emailNotificationLists(String... emailNotificationLists) {
            return emailNotificationLists(List.of(emailNotificationLists));
        }

        /**
         * @param endUserMappingEnabled A boolean indicating whether whether the GTM Domain is using end user client subnet mapping.
         * 
         * @return builder
         * 
         */
        public Builder endUserMappingEnabled(@Nullable Output<Boolean> endUserMappingEnabled) {
            $.endUserMappingEnabled = endUserMappingEnabled;
            return this;
        }

        /**
         * @param endUserMappingEnabled A boolean indicating whether whether the GTM Domain is using end user client subnet mapping.
         * 
         * @return builder
         * 
         */
        public Builder endUserMappingEnabled(Boolean endUserMappingEnabled) {
            return endUserMappingEnabled(Output.of(endUserMappingEnabled));
        }

        /**
         * @param group If creating a domain, the currently selected group ID.
         * 
         * @return builder
         * 
         */
        public Builder group(@Nullable Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group If creating a domain, the currently selected group ID.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param loadFeedback A boolean indicating whether one or more measurements of load (resources) are defined by you and supplied by each data center in real time to balance load.
         * 
         * @return builder
         * 
         */
        public Builder loadFeedback(@Nullable Output<Boolean> loadFeedback) {
            $.loadFeedback = loadFeedback;
            return this;
        }

        /**
         * @param loadFeedback A boolean indicating whether one or more measurements of load (resources) are defined by you and supplied by each data center in real time to balance load.
         * 
         * @return builder
         * 
         */
        public Builder loadFeedback(Boolean loadFeedback) {
            return loadFeedback(Output.of(loadFeedback));
        }

        /**
         * @param loadImbalancePercentage Indicates the percentage of load imbalance factor (LIF) for the domain.
         * 
         * @return builder
         * 
         */
        public Builder loadImbalancePercentage(@Nullable Output<Double> loadImbalancePercentage) {
            $.loadImbalancePercentage = loadImbalancePercentage;
            return this;
        }

        /**
         * @param loadImbalancePercentage Indicates the percentage of load imbalance factor (LIF) for the domain.
         * 
         * @return builder
         * 
         */
        public Builder loadImbalancePercentage(Double loadImbalancePercentage) {
            return loadImbalancePercentage(Output.of(loadImbalancePercentage));
        }

        public Builder mapUpdateInterval(@Nullable Output<Integer> mapUpdateInterval) {
            $.mapUpdateInterval = mapUpdateInterval;
            return this;
        }

        public Builder mapUpdateInterval(Integer mapUpdateInterval) {
            return mapUpdateInterval(Output.of(mapUpdateInterval));
        }

        public Builder maxProperties(@Nullable Output<Integer> maxProperties) {
            $.maxProperties = maxProperties;
            return this;
        }

        public Builder maxProperties(Integer maxProperties) {
            return maxProperties(Output.of(maxProperties));
        }

        public Builder maxResources(@Nullable Output<Integer> maxResources) {
            $.maxResources = maxResources;
            return this;
        }

        public Builder maxResources(Integer maxResources) {
            return maxResources(Output.of(maxResources));
        }

        public Builder maxTestTimeout(@Nullable Output<Double> maxTestTimeout) {
            $.maxTestTimeout = maxTestTimeout;
            return this;
        }

        public Builder maxTestTimeout(Double maxTestTimeout) {
            return maxTestTimeout(Output.of(maxTestTimeout));
        }

        public Builder maxTtl(@Nullable Output<Integer> maxTtl) {
            $.maxTtl = maxTtl;
            return this;
        }

        public Builder maxTtl(Integer maxTtl) {
            return maxTtl(Output.of(maxTtl));
        }

        public Builder minPingableRegionFraction(@Nullable Output<Double> minPingableRegionFraction) {
            $.minPingableRegionFraction = minPingableRegionFraction;
            return this;
        }

        public Builder minPingableRegionFraction(Double minPingableRegionFraction) {
            return minPingableRegionFraction(Output.of(minPingableRegionFraction));
        }

        public Builder minTestInterval(@Nullable Output<Integer> minTestInterval) {
            $.minTestInterval = minTestInterval;
            return this;
        }

        public Builder minTestInterval(Integer minTestInterval) {
            return minTestInterval(Output.of(minTestInterval));
        }

        public Builder minTtl(@Nullable Output<Integer> minTtl) {
            $.minTtl = minTtl;
            return this;
        }

        public Builder minTtl(Integer minTtl) {
            return minTtl(Output.of(minTtl));
        }

        /**
         * @param name The DNS name for a collection of GTM Properties.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The DNS name for a collection of GTM Properties.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder pingInterval(@Nullable Output<Integer> pingInterval) {
            $.pingInterval = pingInterval;
            return this;
        }

        public Builder pingInterval(Integer pingInterval) {
            return pingInterval(Output.of(pingInterval));
        }

        public Builder pingPacketSize(@Nullable Output<Integer> pingPacketSize) {
            $.pingPacketSize = pingPacketSize;
            return this;
        }

        public Builder pingPacketSize(Integer pingPacketSize) {
            return pingPacketSize(Output.of(pingPacketSize));
        }

        public Builder roundRobinPrefix(@Nullable Output<String> roundRobinPrefix) {
            $.roundRobinPrefix = roundRobinPrefix;
            return this;
        }

        public Builder roundRobinPrefix(String roundRobinPrefix) {
            return roundRobinPrefix(Output.of(roundRobinPrefix));
        }

        public Builder servermonitorLivenessCount(@Nullable Output<Integer> servermonitorLivenessCount) {
            $.servermonitorLivenessCount = servermonitorLivenessCount;
            return this;
        }

        public Builder servermonitorLivenessCount(Integer servermonitorLivenessCount) {
            return servermonitorLivenessCount(Output.of(servermonitorLivenessCount));
        }

        public Builder servermonitorLoadCount(@Nullable Output<Integer> servermonitorLoadCount) {
            $.servermonitorLoadCount = servermonitorLoadCount;
            return this;
        }

        public Builder servermonitorLoadCount(Integer servermonitorLoadCount) {
            return servermonitorLoadCount(Output.of(servermonitorLoadCount));
        }

        public Builder servermonitorPool(@Nullable Output<String> servermonitorPool) {
            $.servermonitorPool = servermonitorPool;
            return this;
        }

        public Builder servermonitorPool(String servermonitorPool) {
            return servermonitorPool(Output.of(servermonitorPool));
        }

        /**
         * @param type Th type of GTM domain. Options include `failover-only`, `static`, `weighted`, `basic`, or `full`.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Th type of GTM domain. Options include `failover-only`, `static`, `weighted`, `basic`, or `full`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param waitOnComplete A boolean that, if set to `true`, waits for transaction to complete.
         * 
         * @return builder
         * 
         */
        public Builder waitOnComplete(@Nullable Output<Boolean> waitOnComplete) {
            $.waitOnComplete = waitOnComplete;
            return this;
        }

        /**
         * @param waitOnComplete A boolean that, if set to `true`, waits for transaction to complete.
         * 
         * @return builder
         * 
         */
        public Builder waitOnComplete(Boolean waitOnComplete) {
            return waitOnComplete(Output.of(waitOnComplete));
        }

        public GtmDomainState build() {
            return $;
        }
    }

}
