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


public final class ProviderAppsecArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderAppsecArgs Empty = new ProviderAppsecArgs();

    @Import(name="accessToken")
    private @Nullable Output<String> accessToken;

    public Optional<Output<String>> accessToken() {
        return Optional.ofNullable(this.accessToken);
    }

    @Import(name="accountKey")
    private @Nullable Output<String> accountKey;

    public Optional<Output<String>> accountKey() {
        return Optional.ofNullable(this.accountKey);
    }

    @Import(name="clientSecret")
    private @Nullable Output<String> clientSecret;

    public Optional<Output<String>> clientSecret() {
        return Optional.ofNullable(this.clientSecret);
    }

    @Import(name="clientToken")
    private @Nullable Output<String> clientToken;

    public Optional<Output<String>> clientToken() {
        return Optional.ofNullable(this.clientToken);
    }

    @Import(name="host")
    private @Nullable Output<String> host;

    public Optional<Output<String>> host() {
        return Optional.ofNullable(this.host);
    }

    @Import(name="maxBody")
    private @Nullable Output<Integer> maxBody;

    public Optional<Output<Integer>> maxBody() {
        return Optional.ofNullable(this.maxBody);
    }

    private ProviderAppsecArgs() {}

    private ProviderAppsecArgs(ProviderAppsecArgs $) {
        this.accessToken = $.accessToken;
        this.accountKey = $.accountKey;
        this.clientSecret = $.clientSecret;
        this.clientToken = $.clientToken;
        this.host = $.host;
        this.maxBody = $.maxBody;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderAppsecArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderAppsecArgs $;

        public Builder() {
            $ = new ProviderAppsecArgs();
        }

        public Builder(ProviderAppsecArgs defaults) {
            $ = new ProviderAppsecArgs(Objects.requireNonNull(defaults));
        }

        public Builder accessToken(@Nullable Output<String> accessToken) {
            $.accessToken = accessToken;
            return this;
        }

        public Builder accessToken(String accessToken) {
            return accessToken(Output.of(accessToken));
        }

        public Builder accountKey(@Nullable Output<String> accountKey) {
            $.accountKey = accountKey;
            return this;
        }

        public Builder accountKey(String accountKey) {
            return accountKey(Output.of(accountKey));
        }

        public Builder clientSecret(@Nullable Output<String> clientSecret) {
            $.clientSecret = clientSecret;
            return this;
        }

        public Builder clientSecret(String clientSecret) {
            return clientSecret(Output.of(clientSecret));
        }

        public Builder clientToken(@Nullable Output<String> clientToken) {
            $.clientToken = clientToken;
            return this;
        }

        public Builder clientToken(String clientToken) {
            return clientToken(Output.of(clientToken));
        }

        public Builder host(@Nullable Output<String> host) {
            $.host = host;
            return this;
        }

        public Builder host(String host) {
            return host(Output.of(host));
        }

        public Builder maxBody(@Nullable Output<Integer> maxBody) {
            $.maxBody = maxBody;
            return this;
        }

        public Builder maxBody(Integer maxBody) {
            return maxBody(Output.of(maxBody));
        }

        public ProviderAppsecArgs build() {
            return $;
        }
    }

}
