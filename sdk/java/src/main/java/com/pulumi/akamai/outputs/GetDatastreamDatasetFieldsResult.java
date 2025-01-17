// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.akamai.outputs.GetDatastreamDatasetFieldsField;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDatastreamDatasetFieldsResult {
    private List<GetDatastreamDatasetFieldsField> fields;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String templateName;

    private GetDatastreamDatasetFieldsResult() {}
    public List<GetDatastreamDatasetFieldsField> fields() {
        return this.fields;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> templateName() {
        return Optional.ofNullable(this.templateName);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDatastreamDatasetFieldsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetDatastreamDatasetFieldsField> fields;
        private String id;
        private @Nullable String templateName;
        public Builder() {}
        public Builder(GetDatastreamDatasetFieldsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fields = defaults.fields;
    	      this.id = defaults.id;
    	      this.templateName = defaults.templateName;
        }

        @CustomType.Setter
        public Builder fields(List<GetDatastreamDatasetFieldsField> fields) {
            this.fields = Objects.requireNonNull(fields);
            return this;
        }
        public Builder fields(GetDatastreamDatasetFieldsField... fields) {
            return fields(List.of(fields));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder templateName(@Nullable String templateName) {
            this.templateName = templateName;
            return this;
        }
        public GetDatastreamDatasetFieldsResult build() {
            final var o = new GetDatastreamDatasetFieldsResult();
            o.fields = fields;
            o.id = id;
            o.templateName = templateName;
            return o;
        }
    }
}
