// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CloudletsApplicationLoadBalancerLivenessSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final CloudletsApplicationLoadBalancerLivenessSettingsArgs Empty = new CloudletsApplicationLoadBalancerLivenessSettingsArgs();

    /**
     * Maps additional case-insensitive HTTP header names included to the liveness testing requests.
     * 
     */
    @Import(name="additionalHeaders")
    private @Nullable Output<Map<String,String>> additionalHeaders;

    /**
     * @return Maps additional case-insensitive HTTP header names included to the liveness testing requests.
     * 
     */
    public Optional<Output<Map<String,String>>> additionalHeaders() {
        return Optional.ofNullable(this.additionalHeaders);
    }

    /**
     * The Host header for the liveness HTTP request.
     * 
     */
    @Import(name="hostHeader")
    private @Nullable Output<String> hostHeader;

    /**
     * @return The Host header for the liveness HTTP request.
     * 
     */
    public Optional<Output<String>> hostHeader() {
        return Optional.ofNullable(this.hostHeader);
    }

    /**
     * The frequency of liveness tests. Defaults to 60 seconds, minimum is 10 seconds.
     * 
     */
    @Import(name="interval")
    private @Nullable Output<Integer> interval;

    /**
     * @return The frequency of liveness tests. Defaults to 60 seconds, minimum is 10 seconds.
     * 
     */
    public Optional<Output<Integer>> interval() {
        return Optional.ofNullable(this.interval);
    }

    /**
     * The path to the test object used for liveness testing. The function of the test object is to help determine whether the data center is functioning.
     * 
     */
    @Import(name="path", required=true)
    private Output<String> path;

    /**
     * @return The path to the test object used for liveness testing. The function of the test object is to help determine whether the data center is functioning.
     * 
     */
    public Output<String> path() {
        return this.path;
    }

    /**
     * Whether to validate the origin certificate for an HTTPS request.
     * 
     */
    @Import(name="peerCertificateVerification")
    private @Nullable Output<Boolean> peerCertificateVerification;

    /**
     * @return Whether to validate the origin certificate for an HTTPS request.
     * 
     */
    public Optional<Output<Boolean>> peerCertificateVerification() {
        return Optional.ofNullable(this.peerCertificateVerification);
    }

    /**
     * The port for the test object. The default port is 80, which is standard for HTTP. Enter 443 if you are using HTTPS.
     * 
     */
    @Import(name="port", required=true)
    private Output<Integer> port;

    /**
     * @return The port for the test object. The default port is 80, which is standard for HTTP. Enter 443 if you are using HTTPS.
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }

    /**
     * The protocol or scheme for the database, either `HTTP` or `HTTPS`.
     * 
     */
    @Import(name="protocol", required=true)
    private Output<String> protocol;

    /**
     * @return The protocol or scheme for the database, either `HTTP` or `HTTPS`.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }

    /**
     * The request used for TCP and TCPS tests.
     * 
     */
    @Import(name="requestString")
    private @Nullable Output<String> requestString;

    /**
     * @return The request used for TCP and TCPS tests.
     * 
     */
    public Optional<Output<String>> requestString() {
        return Optional.ofNullable(this.requestString);
    }

    /**
     * The response used for TCP and TCPS tests.
     * 
     */
    @Import(name="responseString")
    private @Nullable Output<String> responseString;

    /**
     * @return The response used for TCP and TCPS tests.
     * 
     */
    public Optional<Output<String>> responseString() {
        return Optional.ofNullable(this.responseString);
    }

    /**
     * If set to `true`, marks the liveness test as failed when the request returns a 3xx (redirection) status code.
     * 
     */
    @Import(name="status3xxFailure")
    private @Nullable Output<Boolean> status3xxFailure;

    /**
     * @return If set to `true`, marks the liveness test as failed when the request returns a 3xx (redirection) status code.
     * 
     */
    public Optional<Output<Boolean>> status3xxFailure() {
        return Optional.ofNullable(this.status3xxFailure);
    }

    /**
     * If set to `true`, marks the liveness test as failed when the request returns a 4xx (client error) status code.
     * 
     */
    @Import(name="status4xxFailure")
    private @Nullable Output<Boolean> status4xxFailure;

    /**
     * @return If set to `true`, marks the liveness test as failed when the request returns a 4xx (client error) status code.
     * 
     */
    public Optional<Output<Boolean>> status4xxFailure() {
        return Optional.ofNullable(this.status4xxFailure);
    }

    /**
     * If set to `true`, marks the liveness test as failed when the request returns a 5xx (server error) status code.
     * 
     */
    @Import(name="status5xxFailure")
    private @Nullable Output<Boolean> status5xxFailure;

    /**
     * @return If set to `true`, marks the liveness test as failed when the request returns a 5xx (server error) status code.
     * 
     */
    public Optional<Output<Boolean>> status5xxFailure() {
        return Optional.ofNullable(this.status5xxFailure);
    }

    /**
     * The number of seconds the system waits before failing the liveness test.
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Double> timeout;

    /**
     * @return The number of seconds the system waits before failing the liveness test.
     * 
     */
    public Optional<Output<Double>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private CloudletsApplicationLoadBalancerLivenessSettingsArgs() {}

    private CloudletsApplicationLoadBalancerLivenessSettingsArgs(CloudletsApplicationLoadBalancerLivenessSettingsArgs $) {
        this.additionalHeaders = $.additionalHeaders;
        this.hostHeader = $.hostHeader;
        this.interval = $.interval;
        this.path = $.path;
        this.peerCertificateVerification = $.peerCertificateVerification;
        this.port = $.port;
        this.protocol = $.protocol;
        this.requestString = $.requestString;
        this.responseString = $.responseString;
        this.status3xxFailure = $.status3xxFailure;
        this.status4xxFailure = $.status4xxFailure;
        this.status5xxFailure = $.status5xxFailure;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudletsApplicationLoadBalancerLivenessSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudletsApplicationLoadBalancerLivenessSettingsArgs $;

        public Builder() {
            $ = new CloudletsApplicationLoadBalancerLivenessSettingsArgs();
        }

        public Builder(CloudletsApplicationLoadBalancerLivenessSettingsArgs defaults) {
            $ = new CloudletsApplicationLoadBalancerLivenessSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param additionalHeaders Maps additional case-insensitive HTTP header names included to the liveness testing requests.
         * 
         * @return builder
         * 
         */
        public Builder additionalHeaders(@Nullable Output<Map<String,String>> additionalHeaders) {
            $.additionalHeaders = additionalHeaders;
            return this;
        }

        /**
         * @param additionalHeaders Maps additional case-insensitive HTTP header names included to the liveness testing requests.
         * 
         * @return builder
         * 
         */
        public Builder additionalHeaders(Map<String,String> additionalHeaders) {
            return additionalHeaders(Output.of(additionalHeaders));
        }

        /**
         * @param hostHeader The Host header for the liveness HTTP request.
         * 
         * @return builder
         * 
         */
        public Builder hostHeader(@Nullable Output<String> hostHeader) {
            $.hostHeader = hostHeader;
            return this;
        }

        /**
         * @param hostHeader The Host header for the liveness HTTP request.
         * 
         * @return builder
         * 
         */
        public Builder hostHeader(String hostHeader) {
            return hostHeader(Output.of(hostHeader));
        }

        /**
         * @param interval The frequency of liveness tests. Defaults to 60 seconds, minimum is 10 seconds.
         * 
         * @return builder
         * 
         */
        public Builder interval(@Nullable Output<Integer> interval) {
            $.interval = interval;
            return this;
        }

        /**
         * @param interval The frequency of liveness tests. Defaults to 60 seconds, minimum is 10 seconds.
         * 
         * @return builder
         * 
         */
        public Builder interval(Integer interval) {
            return interval(Output.of(interval));
        }

        /**
         * @param path The path to the test object used for liveness testing. The function of the test object is to help determine whether the data center is functioning.
         * 
         * @return builder
         * 
         */
        public Builder path(Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path to the test object used for liveness testing. The function of the test object is to help determine whether the data center is functioning.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param peerCertificateVerification Whether to validate the origin certificate for an HTTPS request.
         * 
         * @return builder
         * 
         */
        public Builder peerCertificateVerification(@Nullable Output<Boolean> peerCertificateVerification) {
            $.peerCertificateVerification = peerCertificateVerification;
            return this;
        }

        /**
         * @param peerCertificateVerification Whether to validate the origin certificate for an HTTPS request.
         * 
         * @return builder
         * 
         */
        public Builder peerCertificateVerification(Boolean peerCertificateVerification) {
            return peerCertificateVerification(Output.of(peerCertificateVerification));
        }

        /**
         * @param port The port for the test object. The default port is 80, which is standard for HTTP. Enter 443 if you are using HTTPS.
         * 
         * @return builder
         * 
         */
        public Builder port(Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port for the test object. The default port is 80, which is standard for HTTP. Enter 443 if you are using HTTPS.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param protocol The protocol or scheme for the database, either `HTTP` or `HTTPS`.
         * 
         * @return builder
         * 
         */
        public Builder protocol(Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol The protocol or scheme for the database, either `HTTP` or `HTTPS`.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param requestString The request used for TCP and TCPS tests.
         * 
         * @return builder
         * 
         */
        public Builder requestString(@Nullable Output<String> requestString) {
            $.requestString = requestString;
            return this;
        }

        /**
         * @param requestString The request used for TCP and TCPS tests.
         * 
         * @return builder
         * 
         */
        public Builder requestString(String requestString) {
            return requestString(Output.of(requestString));
        }

        /**
         * @param responseString The response used for TCP and TCPS tests.
         * 
         * @return builder
         * 
         */
        public Builder responseString(@Nullable Output<String> responseString) {
            $.responseString = responseString;
            return this;
        }

        /**
         * @param responseString The response used for TCP and TCPS tests.
         * 
         * @return builder
         * 
         */
        public Builder responseString(String responseString) {
            return responseString(Output.of(responseString));
        }

        /**
         * @param status3xxFailure If set to `true`, marks the liveness test as failed when the request returns a 3xx (redirection) status code.
         * 
         * @return builder
         * 
         */
        public Builder status3xxFailure(@Nullable Output<Boolean> status3xxFailure) {
            $.status3xxFailure = status3xxFailure;
            return this;
        }

        /**
         * @param status3xxFailure If set to `true`, marks the liveness test as failed when the request returns a 3xx (redirection) status code.
         * 
         * @return builder
         * 
         */
        public Builder status3xxFailure(Boolean status3xxFailure) {
            return status3xxFailure(Output.of(status3xxFailure));
        }

        /**
         * @param status4xxFailure If set to `true`, marks the liveness test as failed when the request returns a 4xx (client error) status code.
         * 
         * @return builder
         * 
         */
        public Builder status4xxFailure(@Nullable Output<Boolean> status4xxFailure) {
            $.status4xxFailure = status4xxFailure;
            return this;
        }

        /**
         * @param status4xxFailure If set to `true`, marks the liveness test as failed when the request returns a 4xx (client error) status code.
         * 
         * @return builder
         * 
         */
        public Builder status4xxFailure(Boolean status4xxFailure) {
            return status4xxFailure(Output.of(status4xxFailure));
        }

        /**
         * @param status5xxFailure If set to `true`, marks the liveness test as failed when the request returns a 5xx (server error) status code.
         * 
         * @return builder
         * 
         */
        public Builder status5xxFailure(@Nullable Output<Boolean> status5xxFailure) {
            $.status5xxFailure = status5xxFailure;
            return this;
        }

        /**
         * @param status5xxFailure If set to `true`, marks the liveness test as failed when the request returns a 5xx (server error) status code.
         * 
         * @return builder
         * 
         */
        public Builder status5xxFailure(Boolean status5xxFailure) {
            return status5xxFailure(Output.of(status5xxFailure));
        }

        /**
         * @param timeout The number of seconds the system waits before failing the liveness test.
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Double> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout The number of seconds the system waits before failing the liveness test.
         * 
         * @return builder
         * 
         */
        public Builder timeout(Double timeout) {
            return timeout(Output.of(timeout));
        }

        public CloudletsApplicationLoadBalancerLivenessSettingsArgs build() {
            $.path = Objects.requireNonNull($.path, "expected parameter 'path' to be non-null");
            $.port = Objects.requireNonNull($.port, "expected parameter 'port' to be non-null");
            $.protocol = Objects.requireNonNull($.protocol, "expected parameter 'protocol' to be non-null");
            return $;
        }
    }

}
