// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: bff/bridgr/config/v1alpha1/service.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package ai.popberry.proto.bff.bridgr.config.v1alpha1;

@kotlin.jvm.JvmName("-initializeupdateForwardingRulesResponse")
public inline fun updateForwardingRulesResponse(block: ai.popberry.proto.bff.bridgr.config.v1alpha1.UpdateForwardingRulesResponseKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.bff.bridgr.config.v1alpha1.UpdateForwardingRulesResponse =
  ai.popberry.proto.bff.bridgr.config.v1alpha1.UpdateForwardingRulesResponseKt.Dsl._create(ai.popberry.proto.bff.bridgr.config.v1alpha1.UpdateForwardingRulesResponse.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `bff.bridgr.config.v1alpha1.UpdateForwardingRulesResponse`
 */
public object UpdateForwardingRulesResponseKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: ai.popberry.proto.bff.bridgr.config.v1alpha1.UpdateForwardingRulesResponse.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
      internal fun _create(builder: ai.popberry.proto.bff.bridgr.config.v1alpha1.UpdateForwardingRulesResponse.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
  @kotlin.PublishedApi
    internal fun _build(): ai.popberry.proto.bff.bridgr.config.v1alpha1.UpdateForwardingRulesResponse = _builder.build()

    /**
     * `bool success = 1 [json_name = "success"];`
     */
    public var success: kotlin.Boolean
      @kotlin.jvm.JvmName("getSuccess")
        get() = _builder.success
      @kotlin.jvm.JvmName("setSuccess")
        set(value) {
        _builder.success = value
      }
    /**
     * <code>bool success = 1 [json_name = "success"];</code>
     * @return This builder for chaining.
     */
    public fun clearSuccess() {
      _builder.clearSuccess()
    }

    /**
     * `string message = 2 [json_name = "message"];`
     */
    public var message: kotlin.String
      @kotlin.jvm.JvmName("getMessage")
        get() = _builder.message
      @kotlin.jvm.JvmName("setMessage")
        set(value) {
        _builder.message = value
      }
    /**
     * <code>string message = 2 [json_name = "message"];</code>
     * @return This builder for chaining.
     */
    public fun clearMessage() {
      _builder.clearMessage()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun ai.popberry.proto.bff.bridgr.config.v1alpha1.UpdateForwardingRulesResponse.copy(block: `ai.popberry.proto.bff.bridgr.config.v1alpha1`.UpdateForwardingRulesResponseKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.bff.bridgr.config.v1alpha1.UpdateForwardingRulesResponse =
  `ai.popberry.proto.bff.bridgr.config.v1alpha1`.UpdateForwardingRulesResponseKt.Dsl._create(this.toBuilder()).apply { block() }._build()

