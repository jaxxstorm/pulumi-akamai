// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGroupResult {
    /**
     * @deprecated
     * The setting &#34;contract&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""contract"" has been deprecated. */
    private String contract;
    private String contractId;
    private String groupName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @deprecated
     * The setting &#34;name&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""name"" has been deprecated. */
    private String name;

    private GetGroupResult() {}
    /**
     * @deprecated
     * The setting &#34;contract&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""contract"" has been deprecated. */
    public String contract() {
        return this.contract;
    }
    public String contractId() {
        return this.contractId;
    }
    public String groupName() {
        return this.groupName;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @deprecated
     * The setting &#34;name&#34; has been deprecated.
     * 
     */
    @Deprecated /* The setting ""name"" has been deprecated. */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGroupResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String contract;
        private String contractId;
        private String groupName;
        private String id;
        private String name;
        public Builder() {}
        public Builder(GetGroupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.contract = defaults.contract;
    	      this.contractId = defaults.contractId;
    	      this.groupName = defaults.groupName;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder contract(String contract) {
            this.contract = Objects.requireNonNull(contract);
            return this;
        }
        @CustomType.Setter
        public Builder contractId(String contractId) {
            this.contractId = Objects.requireNonNull(contractId);
            return this;
        }
        @CustomType.Setter
        public Builder groupName(String groupName) {
            this.groupName = Objects.requireNonNull(groupName);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public GetGroupResult build() {
            final var o = new GetGroupResult();
            o.contract = contract;
            o.contractId = contractId;
            o.groupName = groupName;
            o.id = id;
            o.name = name;
            return o;
        }
    }
}
