// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DatastreamDatadogConnectorArgs extends com.pulumi.resources.ResourceArgs {

    public static final DatastreamDatadogConnectorArgs Empty = new DatastreamDatadogConnectorArgs();

    /**
     * **Secret**. The API key associated with your Datadog account. See [View API keys in Datadog](https://docs.datadoghq.com/account_management/api-app-keys/#api-keys).
     * * `compress logs` - (Optional) Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `false`.
     * 
     */
    @Import(name="authToken", required=true)
    private Output<String> authToken;

    /**
     * @return **Secret**. The API key associated with your Datadog account. See [View API keys in Datadog](https://docs.datadoghq.com/account_management/api-app-keys/#api-keys).
     * * `compress logs` - (Optional) Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `false`.
     * 
     */
    public Output<String> authToken() {
        return this.authToken;
    }

    /**
     * Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `true`.
     * 
     */
    @Import(name="compressLogs")
    private @Nullable Output<Boolean> compressLogs;

    /**
     * @return Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> compressLogs() {
        return Optional.ofNullable(this.compressLogs);
    }

    @Import(name="connectorId")
    private @Nullable Output<Integer> connectorId;

    public Optional<Output<Integer>> connectorId() {
        return Optional.ofNullable(this.connectorId);
    }

    /**
     * The name of the connector.
     * 
     */
    @Import(name="connectorName", required=true)
    private Output<String> connectorName;

    /**
     * @return The name of the connector.
     * 
     */
    public Output<String> connectorName() {
        return this.connectorName;
    }

    /**
     * The service of the Datadog connector. A service groups together endpoints, queries, or jobs for the purposes of scaling instances. See [View Datadog reserved attribute list](https://docs.datadoghq.com/logs/log_configuration/attributes_naming_convention/#reserved-attributes).
     * 
     */
    @Import(name="service")
    private @Nullable Output<String> service;

    /**
     * @return The service of the Datadog connector. A service groups together endpoints, queries, or jobs for the purposes of scaling instances. See [View Datadog reserved attribute list](https://docs.datadoghq.com/logs/log_configuration/attributes_naming_convention/#reserved-attributes).
     * 
     */
    public Optional<Output<String>> service() {
        return Optional.ofNullable(this.service);
    }

    /**
     * The source of the Datadog connector. See [View Datadog reserved attribute list](https://docs.datadoghq.com/logs/log_collection/?tab=http#reserved-attributes).
     * 
     */
    @Import(name="source")
    private @Nullable Output<String> source;

    /**
     * @return The source of the Datadog connector. See [View Datadog reserved attribute list](https://docs.datadoghq.com/logs/log_collection/?tab=http#reserved-attributes).
     * 
     */
    public Optional<Output<String>> source() {
        return Optional.ofNullable(this.source);
    }

    /**
     * The tags of the Datadog connector. See [View Datadog tags](https://docs.datadoghq.com/getting_started/tagging/).
     * 
     */
    @Import(name="tags")
    private @Nullable Output<String> tags;

    /**
     * @return The tags of the Datadog connector. See [View Datadog tags](https://docs.datadoghq.com/getting_started/tagging/).
     * 
     */
    public Optional<Output<String>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Enter the secure URL where you want to send and store your logs.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return Enter the secure URL where you want to send and store your logs.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    private DatastreamDatadogConnectorArgs() {}

    private DatastreamDatadogConnectorArgs(DatastreamDatadogConnectorArgs $) {
        this.authToken = $.authToken;
        this.compressLogs = $.compressLogs;
        this.connectorId = $.connectorId;
        this.connectorName = $.connectorName;
        this.service = $.service;
        this.source = $.source;
        this.tags = $.tags;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatastreamDatadogConnectorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatastreamDatadogConnectorArgs $;

        public Builder() {
            $ = new DatastreamDatadogConnectorArgs();
        }

        public Builder(DatastreamDatadogConnectorArgs defaults) {
            $ = new DatastreamDatadogConnectorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authToken **Secret**. The API key associated with your Datadog account. See [View API keys in Datadog](https://docs.datadoghq.com/account_management/api-app-keys/#api-keys).
         * * `compress logs` - (Optional) Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder authToken(Output<String> authToken) {
            $.authToken = authToken;
            return this;
        }

        /**
         * @param authToken **Secret**. The API key associated with your Datadog account. See [View API keys in Datadog](https://docs.datadoghq.com/account_management/api-app-keys/#api-keys).
         * * `compress logs` - (Optional) Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder authToken(String authToken) {
            return authToken(Output.of(authToken));
        }

        /**
         * @param compressLogs Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder compressLogs(@Nullable Output<Boolean> compressLogs) {
            $.compressLogs = compressLogs;
            return this;
        }

        /**
         * @param compressLogs Enables GZIP compression for a log file sent to a destination. If unspecified, this defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder compressLogs(Boolean compressLogs) {
            return compressLogs(Output.of(compressLogs));
        }

        public Builder connectorId(@Nullable Output<Integer> connectorId) {
            $.connectorId = connectorId;
            return this;
        }

        public Builder connectorId(Integer connectorId) {
            return connectorId(Output.of(connectorId));
        }

        /**
         * @param connectorName The name of the connector.
         * 
         * @return builder
         * 
         */
        public Builder connectorName(Output<String> connectorName) {
            $.connectorName = connectorName;
            return this;
        }

        /**
         * @param connectorName The name of the connector.
         * 
         * @return builder
         * 
         */
        public Builder connectorName(String connectorName) {
            return connectorName(Output.of(connectorName));
        }

        /**
         * @param service The service of the Datadog connector. A service groups together endpoints, queries, or jobs for the purposes of scaling instances. See [View Datadog reserved attribute list](https://docs.datadoghq.com/logs/log_configuration/attributes_naming_convention/#reserved-attributes).
         * 
         * @return builder
         * 
         */
        public Builder service(@Nullable Output<String> service) {
            $.service = service;
            return this;
        }

        /**
         * @param service The service of the Datadog connector. A service groups together endpoints, queries, or jobs for the purposes of scaling instances. See [View Datadog reserved attribute list](https://docs.datadoghq.com/logs/log_configuration/attributes_naming_convention/#reserved-attributes).
         * 
         * @return builder
         * 
         */
        public Builder service(String service) {
            return service(Output.of(service));
        }

        /**
         * @param source The source of the Datadog connector. See [View Datadog reserved attribute list](https://docs.datadoghq.com/logs/log_collection/?tab=http#reserved-attributes).
         * 
         * @return builder
         * 
         */
        public Builder source(@Nullable Output<String> source) {
            $.source = source;
            return this;
        }

        /**
         * @param source The source of the Datadog connector. See [View Datadog reserved attribute list](https://docs.datadoghq.com/logs/log_collection/?tab=http#reserved-attributes).
         * 
         * @return builder
         * 
         */
        public Builder source(String source) {
            return source(Output.of(source));
        }

        /**
         * @param tags The tags of the Datadog connector. See [View Datadog tags](https://docs.datadoghq.com/getting_started/tagging/).
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<String> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags of the Datadog connector. See [View Datadog tags](https://docs.datadoghq.com/getting_started/tagging/).
         * 
         * @return builder
         * 
         */
        public Builder tags(String tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param url Enter the secure URL where you want to send and store your logs.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Enter the secure URL where you want to send and store your logs.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public DatastreamDatadogConnectorArgs build() {
            $.authToken = Objects.requireNonNull($.authToken, "expected parameter 'authToken' to be non-null");
            $.connectorName = Objects.requireNonNull($.connectorName, "expected parameter 'connectorName' to be non-null");
            $.url = Objects.requireNonNull($.url, "expected parameter 'url' to be non-null");
            return $;
        }
    }

}
