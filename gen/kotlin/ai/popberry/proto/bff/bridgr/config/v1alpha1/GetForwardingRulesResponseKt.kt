// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: bff/bridgr/config/v1alpha1/service.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package ai.popberry.proto.bff.bridgr.config.v1alpha1;

@kotlin.jvm.JvmName("-initializegetForwardingRulesResponse")
public inline fun getForwardingRulesResponse(block: ai.popberry.proto.bff.bridgr.config.v1alpha1.GetForwardingRulesResponseKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.bff.bridgr.config.v1alpha1.GetForwardingRulesResponse =
  ai.popberry.proto.bff.bridgr.config.v1alpha1.GetForwardingRulesResponseKt.Dsl._create(ai.popberry.proto.bff.bridgr.config.v1alpha1.GetForwardingRulesResponse.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `bff.bridgr.config.v1alpha1.GetForwardingRulesResponse`
 */
public object GetForwardingRulesResponseKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: ai.popberry.proto.bff.bridgr.config.v1alpha1.GetForwardingRulesResponse.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
      internal fun _create(builder: ai.popberry.proto.bff.bridgr.config.v1alpha1.GetForwardingRulesResponse.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
  @kotlin.PublishedApi
    internal fun _build(): ai.popberry.proto.bff.bridgr.config.v1alpha1.GetForwardingRulesResponse = _builder.build()

    /**
     * An uninstantiable, behaviorless type to represent the field in
     * generics.
     */
    @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
    public class AllowedSendersProxy private constructor() : com.google.protobuf.kotlin.DslProxy()
    /**
     * <code>repeated string allowed_senders = 1 [json_name = "allowedSenders"];</code>
     * @return A list containing the allowedSenders.
     * @return This builder for chaining.
     */
    public val allowedSenders: com.google.protobuf.kotlin.DslList<kotlin.String, AllowedSendersProxy>
    @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
      get() = com.google.protobuf.kotlin.DslList(
        _builder.allowedSendersList
      )
    /**
     * <code>repeated string allowed_senders = 1 [json_name = "allowedSenders"];</code>
     * @param value The allowedSenders to add.
     * @return This builder for chaining.
     */
    @kotlin.jvm.JvmSynthetic
@kotlin.jvm.JvmName("addAllowedSenders")
    public fun com.google.protobuf.kotlin.DslList<kotlin.String, AllowedSendersProxy>.add(value: kotlin.String) {
      _builder.addAllowedSenders(value)
    }
    /**
     * <code>repeated string allowed_senders = 1 [json_name = "allowedSenders"];</code>
     * @param value The allowedSenders to add.
     * @return This builder for chaining.
     */
    @kotlin.jvm.JvmSynthetic
@kotlin.jvm.JvmName("plusAssignAllowedSenders")
    @Suppress("NOTHING_TO_INLINE")
    public inline operator fun com.google.protobuf.kotlin.DslList<kotlin.String, AllowedSendersProxy>.plusAssign(value: kotlin.String) {
      add(value)
    }
    /**
     * <code>repeated string allowed_senders = 1 [json_name = "allowedSenders"];</code>
     * @param values The allowedSenders to add.
     * @return This builder for chaining.
     */
    @kotlin.jvm.JvmSynthetic
@kotlin.jvm.JvmName("addAllAllowedSenders")
    public fun com.google.protobuf.kotlin.DslList<kotlin.String, AllowedSendersProxy>.addAll(values: kotlin.collections.Iterable<kotlin.String>) {
      _builder.addAllAllowedSenders(values)
    }
    /**
     * <code>repeated string allowed_senders = 1 [json_name = "allowedSenders"];</code>
     * @param values The allowedSenders to add.
     * @return This builder for chaining.
     */
    @kotlin.jvm.JvmSynthetic
@kotlin.jvm.JvmName("plusAssignAllAllowedSenders")
    @Suppress("NOTHING_TO_INLINE")
    public inline operator fun com.google.protobuf.kotlin.DslList<kotlin.String, AllowedSendersProxy>.plusAssign(values: kotlin.collections.Iterable<kotlin.String>) {
      addAll(values)
    }
    /**
     * <code>repeated string allowed_senders = 1 [json_name = "allowedSenders"];</code>
     * @param index The index to set the value at.
     * @param value The allowedSenders to set.
     * @return This builder for chaining.
     */
    @kotlin.jvm.JvmSynthetic
@kotlin.jvm.JvmName("setAllowedSenders")
    public operator fun com.google.protobuf.kotlin.DslList<kotlin.String, AllowedSendersProxy>.set(index: kotlin.Int, value: kotlin.String) {
      _builder.setAllowedSenders(index, value)
    }/**
     * <code>repeated string allowed_senders = 1 [json_name = "allowedSenders"];</code>
     * @return This builder for chaining.
     */
    @kotlin.jvm.JvmSynthetic
@kotlin.jvm.JvmName("setAllowedSenders")
    public fun com.google.protobuf.kotlin.DslList<kotlin.String, AllowedSendersProxy>.clear() {
      _builder.clearAllowedSenders()
    }
    /**
     * An uninstantiable, behaviorless type to represent the field in
     * generics.
     */
    @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
    public class AllowedAppsProxy private constructor() : com.google.protobuf.kotlin.DslProxy()
    /**
     * <code>repeated string allowed_apps = 2 [json_name = "allowedApps"];</code>
     * @return A list containing the allowedApps.
     * @return This builder for chaining.
     */
    public val allowedApps: com.google.protobuf.kotlin.DslList<kotlin.String, AllowedAppsProxy>
    @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
      get() = com.google.protobuf.kotlin.DslList(
        _builder.allowedAppsList
      )
    /**
     * <code>repeated string allowed_apps = 2 [json_name = "allowedApps"];</code>
     * @param value The allowedApps to add.
     * @return This builder for chaining.
     */
    @kotlin.jvm.JvmSynthetic
@kotlin.jvm.JvmName("addAllowedApps")
    public fun com.google.protobuf.kotlin.DslList<kotlin.String, AllowedAppsProxy>.add(value: kotlin.String) {
      _builder.addAllowedApps(value)
    }
    /**
     * <code>repeated string allowed_apps = 2 [json_name = "allowedApps"];</code>
     * @param value The allowedApps to add.
     * @return This builder for chaining.
     */
    @kotlin.jvm.JvmSynthetic
@kotlin.jvm.JvmName("plusAssignAllowedApps")
    @Suppress("NOTHING_TO_INLINE")
    public inline operator fun com.google.protobuf.kotlin.DslList<kotlin.String, AllowedAppsProxy>.plusAssign(value: kotlin.String) {
      add(value)
    }
    /**
     * <code>repeated string allowed_apps = 2 [json_name = "allowedApps"];</code>
     * @param values The allowedApps to add.
     * @return This builder for chaining.
     */
    @kotlin.jvm.JvmSynthetic
@kotlin.jvm.JvmName("addAllAllowedApps")
    public fun com.google.protobuf.kotlin.DslList<kotlin.String, AllowedAppsProxy>.addAll(values: kotlin.collections.Iterable<kotlin.String>) {
      _builder.addAllAllowedApps(values)
    }
    /**
     * <code>repeated string allowed_apps = 2 [json_name = "allowedApps"];</code>
     * @param values The allowedApps to add.
     * @return This builder for chaining.
     */
    @kotlin.jvm.JvmSynthetic
@kotlin.jvm.JvmName("plusAssignAllAllowedApps")
    @Suppress("NOTHING_TO_INLINE")
    public inline operator fun com.google.protobuf.kotlin.DslList<kotlin.String, AllowedAppsProxy>.plusAssign(values: kotlin.collections.Iterable<kotlin.String>) {
      addAll(values)
    }
    /**
     * <code>repeated string allowed_apps = 2 [json_name = "allowedApps"];</code>
     * @param index The index to set the value at.
     * @param value The allowedApps to set.
     * @return This builder for chaining.
     */
    @kotlin.jvm.JvmSynthetic
@kotlin.jvm.JvmName("setAllowedApps")
    public operator fun com.google.protobuf.kotlin.DslList<kotlin.String, AllowedAppsProxy>.set(index: kotlin.Int, value: kotlin.String) {
      _builder.setAllowedApps(index, value)
    }/**
     * <code>repeated string allowed_apps = 2 [json_name = "allowedApps"];</code>
     * @return This builder for chaining.
     */
    @kotlin.jvm.JvmSynthetic
@kotlin.jvm.JvmName("setAllowedApps")
    public fun com.google.protobuf.kotlin.DslList<kotlin.String, AllowedAppsProxy>.clear() {
      _builder.clearAllowedApps()
    }}
}
@kotlin.jvm.JvmSynthetic
public inline fun ai.popberry.proto.bff.bridgr.config.v1alpha1.GetForwardingRulesResponse.copy(block: `ai.popberry.proto.bff.bridgr.config.v1alpha1`.GetForwardingRulesResponseKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.bff.bridgr.config.v1alpha1.GetForwardingRulesResponse =
  `ai.popberry.proto.bff.bridgr.config.v1alpha1`.GetForwardingRulesResponseKt.Dsl._create(this.toBuilder()).apply { block() }._build()

