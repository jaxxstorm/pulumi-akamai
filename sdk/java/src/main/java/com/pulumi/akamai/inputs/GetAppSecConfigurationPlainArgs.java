// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAppSecConfigurationPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecConfigurationPlainArgs Empty = new GetAppSecConfigurationPlainArgs();

    /**
     * . Name of the security configuration you want to return information for. If not included, information is returned for all your security configurations.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return . Name of the security configuration you want to return information for. If not included, information is returned for all your security configurations.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    private GetAppSecConfigurationPlainArgs() {}

    private GetAppSecConfigurationPlainArgs(GetAppSecConfigurationPlainArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecConfigurationPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecConfigurationPlainArgs $;

        public Builder() {
            $ = new GetAppSecConfigurationPlainArgs();
        }

        public Builder(GetAppSecConfigurationPlainArgs defaults) {
            $ = new GetAppSecConfigurationPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name . Name of the security configuration you want to return information for. If not included, information is returned for all your security configurations.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public GetAppSecConfigurationPlainArgs build() {
            return $;
        }
    }

}
