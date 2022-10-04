// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GtmResourceResourceInstance {
    /**
     * @return A unique identifier for an existing data center in the domain.
     * 
     */
    private Integer datacenterId;
    /**
     * @return Identifies the load object file used to report real-time information about the current load, maximum allowable load, and target load on each resource.
     * 
     */
    private @Nullable String loadObject;
    /**
     * @return Specifies the TCP port of the `load_object`.
     * 
     */
    private @Nullable Integer loadObjectPort;
    /**
     * @return (List) Specifies a list of servers from which to request the load object.
     * 
     */
    private @Nullable List<String> loadServers;
    /**
     * @return A boolean that indicates whether a default `load_object` is used for the resources.
     * 
     */
    private @Nullable Boolean useDefaultLoadObject;

    private GtmResourceResourceInstance() {}
    /**
     * @return A unique identifier for an existing data center in the domain.
     * 
     */
    public Integer datacenterId() {
        return this.datacenterId;
    }
    /**
     * @return Identifies the load object file used to report real-time information about the current load, maximum allowable load, and target load on each resource.
     * 
     */
    public Optional<String> loadObject() {
        return Optional.ofNullable(this.loadObject);
    }
    /**
     * @return Specifies the TCP port of the `load_object`.
     * 
     */
    public Optional<Integer> loadObjectPort() {
        return Optional.ofNullable(this.loadObjectPort);
    }
    /**
     * @return (List) Specifies a list of servers from which to request the load object.
     * 
     */
    public List<String> loadServers() {
        return this.loadServers == null ? List.of() : this.loadServers;
    }
    /**
     * @return A boolean that indicates whether a default `load_object` is used for the resources.
     * 
     */
    public Optional<Boolean> useDefaultLoadObject() {
        return Optional.ofNullable(this.useDefaultLoadObject);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GtmResourceResourceInstance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer datacenterId;
        private @Nullable String loadObject;
        private @Nullable Integer loadObjectPort;
        private @Nullable List<String> loadServers;
        private @Nullable Boolean useDefaultLoadObject;
        public Builder() {}
        public Builder(GtmResourceResourceInstance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.datacenterId = defaults.datacenterId;
    	      this.loadObject = defaults.loadObject;
    	      this.loadObjectPort = defaults.loadObjectPort;
    	      this.loadServers = defaults.loadServers;
    	      this.useDefaultLoadObject = defaults.useDefaultLoadObject;
        }

        @CustomType.Setter
        public Builder datacenterId(Integer datacenterId) {
            this.datacenterId = Objects.requireNonNull(datacenterId);
            return this;
        }
        @CustomType.Setter
        public Builder loadObject(@Nullable String loadObject) {
            this.loadObject = loadObject;
            return this;
        }
        @CustomType.Setter
        public Builder loadObjectPort(@Nullable Integer loadObjectPort) {
            this.loadObjectPort = loadObjectPort;
            return this;
        }
        @CustomType.Setter
        public Builder loadServers(@Nullable List<String> loadServers) {
            this.loadServers = loadServers;
            return this;
        }
        public Builder loadServers(String... loadServers) {
            return loadServers(List.of(loadServers));
        }
        @CustomType.Setter
        public Builder useDefaultLoadObject(@Nullable Boolean useDefaultLoadObject) {
            this.useDefaultLoadObject = useDefaultLoadObject;
            return this;
        }
        public GtmResourceResourceInstance build() {
            final var o = new GtmResourceResourceInstance();
            o.datacenterId = datacenterId;
            o.loadObject = loadObject;
            o.loadObjectPort = loadObjectPort;
            o.loadServers = loadServers;
            o.useDefaultLoadObject = useDefaultLoadObject;
            return o;
        }
    }
}
