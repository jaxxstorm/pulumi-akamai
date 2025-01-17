// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIamRolesRole {
    private String createdBy;
    private String description;
    private String modifiedBy;
    private String name;
    private String roleId;
    private String timeCreated;
    private String timeModified;
    private String type;

    private GetIamRolesRole() {}
    public String createdBy() {
        return this.createdBy;
    }
    public String description() {
        return this.description;
    }
    public String modifiedBy() {
        return this.modifiedBy;
    }
    public String name() {
        return this.name;
    }
    public String roleId() {
        return this.roleId;
    }
    public String timeCreated() {
        return this.timeCreated;
    }
    public String timeModified() {
        return this.timeModified;
    }
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIamRolesRole defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdBy;
        private String description;
        private String modifiedBy;
        private String name;
        private String roleId;
        private String timeCreated;
        private String timeModified;
        private String type;
        public Builder() {}
        public Builder(GetIamRolesRole defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdBy = defaults.createdBy;
    	      this.description = defaults.description;
    	      this.modifiedBy = defaults.modifiedBy;
    	      this.name = defaults.name;
    	      this.roleId = defaults.roleId;
    	      this.timeCreated = defaults.timeCreated;
    	      this.timeModified = defaults.timeModified;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder createdBy(String createdBy) {
            this.createdBy = Objects.requireNonNull(createdBy);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder modifiedBy(String modifiedBy) {
            this.modifiedBy = Objects.requireNonNull(modifiedBy);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder roleId(String roleId) {
            this.roleId = Objects.requireNonNull(roleId);
            return this;
        }
        @CustomType.Setter
        public Builder timeCreated(String timeCreated) {
            this.timeCreated = Objects.requireNonNull(timeCreated);
            return this;
        }
        @CustomType.Setter
        public Builder timeModified(String timeModified) {
            this.timeModified = Objects.requireNonNull(timeModified);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetIamRolesRole build() {
            final var o = new GetIamRolesRole();
            o.createdBy = createdBy;
            o.description = description;
            o.modifiedBy = modifiedBy;
            o.name = name;
            o.roleId = roleId;
            o.timeCreated = timeCreated;
            o.timeModified = timeModified;
            o.type = type;
            return o;
        }
    }
}
