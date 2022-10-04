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


public final class PropertyOriginArgs extends com.pulumi.resources.ResourceArgs {

    public static final PropertyOriginArgs Empty = new PropertyOriginArgs();

    @Import(name="cacheKeyHostname")
    private @Nullable Output<String> cacheKeyHostname;

    public Optional<Output<String>> cacheKeyHostname() {
        return Optional.ofNullable(this.cacheKeyHostname);
    }

    @Import(name="compress")
    private @Nullable Output<Boolean> compress;

    public Optional<Output<Boolean>> compress() {
        return Optional.ofNullable(this.compress);
    }

    @Import(name="enableTrueClientIp")
    private @Nullable Output<Boolean> enableTrueClientIp;

    public Optional<Output<Boolean>> enableTrueClientIp() {
        return Optional.ofNullable(this.enableTrueClientIp);
    }

    @Import(name="forwardHostname")
    private @Nullable Output<String> forwardHostname;

    public Optional<Output<String>> forwardHostname() {
        return Optional.ofNullable(this.forwardHostname);
    }

    @Import(name="hostname")
    private @Nullable Output<String> hostname;

    public Optional<Output<String>> hostname() {
        return Optional.ofNullable(this.hostname);
    }

    @Import(name="port")
    private @Nullable Output<Integer> port;

    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    private PropertyOriginArgs() {}

    private PropertyOriginArgs(PropertyOriginArgs $) {
        this.cacheKeyHostname = $.cacheKeyHostname;
        this.compress = $.compress;
        this.enableTrueClientIp = $.enableTrueClientIp;
        this.forwardHostname = $.forwardHostname;
        this.hostname = $.hostname;
        this.port = $.port;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PropertyOriginArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PropertyOriginArgs $;

        public Builder() {
            $ = new PropertyOriginArgs();
        }

        public Builder(PropertyOriginArgs defaults) {
            $ = new PropertyOriginArgs(Objects.requireNonNull(defaults));
        }

        public Builder cacheKeyHostname(@Nullable Output<String> cacheKeyHostname) {
            $.cacheKeyHostname = cacheKeyHostname;
            return this;
        }

        public Builder cacheKeyHostname(String cacheKeyHostname) {
            return cacheKeyHostname(Output.of(cacheKeyHostname));
        }

        public Builder compress(@Nullable Output<Boolean> compress) {
            $.compress = compress;
            return this;
        }

        public Builder compress(Boolean compress) {
            return compress(Output.of(compress));
        }

        public Builder enableTrueClientIp(@Nullable Output<Boolean> enableTrueClientIp) {
            $.enableTrueClientIp = enableTrueClientIp;
            return this;
        }

        public Builder enableTrueClientIp(Boolean enableTrueClientIp) {
            return enableTrueClientIp(Output.of(enableTrueClientIp));
        }

        public Builder forwardHostname(@Nullable Output<String> forwardHostname) {
            $.forwardHostname = forwardHostname;
            return this;
        }

        public Builder forwardHostname(String forwardHostname) {
            return forwardHostname(Output.of(forwardHostname));
        }

        public Builder hostname(@Nullable Output<String> hostname) {
            $.hostname = hostname;
            return this;
        }

        public Builder hostname(String hostname) {
            return hostname(Output.of(hostname));
        }

        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        public PropertyOriginArgs build() {
            return $;
        }
    }

}
