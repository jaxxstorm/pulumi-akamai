// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;


public final class GetAppSecVersionNotesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppSecVersionNotesPlainArgs Empty = new GetAppSecVersionNotesPlainArgs();

    /**
     * . Unique identifier of the security configuration you want to return information for.
     * 
     */
    @Import(name="configId", required=true)
    private Integer configId;

    /**
     * @return . Unique identifier of the security configuration you want to return information for.
     * 
     */
    public Integer configId() {
        return this.configId;
    }

    private GetAppSecVersionNotesPlainArgs() {}

    private GetAppSecVersionNotesPlainArgs(GetAppSecVersionNotesPlainArgs $) {
        this.configId = $.configId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppSecVersionNotesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppSecVersionNotesPlainArgs $;

        public Builder() {
            $ = new GetAppSecVersionNotesPlainArgs();
        }

        public Builder(GetAppSecVersionNotesPlainArgs defaults) {
            $ = new GetAppSecVersionNotesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configId . Unique identifier of the security configuration you want to return information for.
         * 
         * @return builder
         * 
         */
        public Builder configId(Integer configId) {
            $.configId = configId;
            return this;
        }

        public GetAppSecVersionNotesPlainArgs build() {
            $.configId = Objects.requireNonNull($.configId, "expected parameter 'configId' to be non-null");
            return $;
        }
    }

}
