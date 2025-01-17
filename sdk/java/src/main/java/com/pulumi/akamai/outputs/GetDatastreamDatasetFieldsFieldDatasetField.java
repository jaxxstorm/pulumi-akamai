// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDatastreamDatasetFieldsFieldDatasetField {
    private String datasetFieldDescription;
    private Integer datasetFieldId;
    private String datasetFieldJsonKey;
    private String datasetFieldName;

    private GetDatastreamDatasetFieldsFieldDatasetField() {}
    public String datasetFieldDescription() {
        return this.datasetFieldDescription;
    }
    public Integer datasetFieldId() {
        return this.datasetFieldId;
    }
    public String datasetFieldJsonKey() {
        return this.datasetFieldJsonKey;
    }
    public String datasetFieldName() {
        return this.datasetFieldName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDatastreamDatasetFieldsFieldDatasetField defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String datasetFieldDescription;
        private Integer datasetFieldId;
        private String datasetFieldJsonKey;
        private String datasetFieldName;
        public Builder() {}
        public Builder(GetDatastreamDatasetFieldsFieldDatasetField defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.datasetFieldDescription = defaults.datasetFieldDescription;
    	      this.datasetFieldId = defaults.datasetFieldId;
    	      this.datasetFieldJsonKey = defaults.datasetFieldJsonKey;
    	      this.datasetFieldName = defaults.datasetFieldName;
        }

        @CustomType.Setter
        public Builder datasetFieldDescription(String datasetFieldDescription) {
            this.datasetFieldDescription = Objects.requireNonNull(datasetFieldDescription);
            return this;
        }
        @CustomType.Setter
        public Builder datasetFieldId(Integer datasetFieldId) {
            this.datasetFieldId = Objects.requireNonNull(datasetFieldId);
            return this;
        }
        @CustomType.Setter
        public Builder datasetFieldJsonKey(String datasetFieldJsonKey) {
            this.datasetFieldJsonKey = Objects.requireNonNull(datasetFieldJsonKey);
            return this;
        }
        @CustomType.Setter
        public Builder datasetFieldName(String datasetFieldName) {
            this.datasetFieldName = Objects.requireNonNull(datasetFieldName);
            return this;
        }
        public GetDatastreamDatasetFieldsFieldDatasetField build() {
            final var o = new GetDatastreamDatasetFieldsFieldDatasetField();
            o.datasetFieldDescription = datasetFieldDescription;
            o.datasetFieldId = datasetFieldId;
            o.datasetFieldJsonKey = datasetFieldJsonKey;
            o.datasetFieldName = datasetFieldName;
            return o;
        }
    }
}
