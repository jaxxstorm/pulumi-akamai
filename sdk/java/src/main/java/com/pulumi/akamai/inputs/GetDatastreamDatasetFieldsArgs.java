// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDatastreamDatasetFieldsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDatastreamDatasetFieldsArgs Empty = new GetDatastreamDatasetFieldsArgs();

    /**
     * - (Optional) The name of the data set template you use in your stream configuration. Currently, `EDGE_LOGS` is the only available data set template and the default value for this argument.
     * 
     */
    @Import(name="templateName")
    private @Nullable Output<String> templateName;

    /**
     * @return - (Optional) The name of the data set template you use in your stream configuration. Currently, `EDGE_LOGS` is the only available data set template and the default value for this argument.
     * 
     */
    public Optional<Output<String>> templateName() {
        return Optional.ofNullable(this.templateName);
    }

    private GetDatastreamDatasetFieldsArgs() {}

    private GetDatastreamDatasetFieldsArgs(GetDatastreamDatasetFieldsArgs $) {
        this.templateName = $.templateName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDatastreamDatasetFieldsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDatastreamDatasetFieldsArgs $;

        public Builder() {
            $ = new GetDatastreamDatasetFieldsArgs();
        }

        public Builder(GetDatastreamDatasetFieldsArgs defaults) {
            $ = new GetDatastreamDatasetFieldsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param templateName - (Optional) The name of the data set template you use in your stream configuration. Currently, `EDGE_LOGS` is the only available data set template and the default value for this argument.
         * 
         * @return builder
         * 
         */
        public Builder templateName(@Nullable Output<String> templateName) {
            $.templateName = templateName;
            return this;
        }

        /**
         * @param templateName - (Optional) The name of the data set template you use in your stream configuration. Currently, `EDGE_LOGS` is the only available data set template and the default value for this argument.
         * 
         * @return builder
         * 
         */
        public Builder templateName(String templateName) {
            return templateName(Output.of(templateName));
        }

        public GetDatastreamDatasetFieldsArgs build() {
            return $;
        }
    }

}