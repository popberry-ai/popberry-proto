// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: game/joker/v1alpha1/joker.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package ai.popberry.proto.game.joker.v1alpha1;

@kotlin.jvm.JvmName("-initializegameSignoutResponse")
public inline fun gameSignoutResponse(block: ai.popberry.proto.game.joker.v1alpha1.GameSignoutResponseKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.game.joker.v1alpha1.GameSignoutResponse =
  ai.popberry.proto.game.joker.v1alpha1.GameSignoutResponseKt.Dsl._create(ai.popberry.proto.game.joker.v1alpha1.GameSignoutResponse.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `game.joker.v1alpha1.GameSignoutResponse`
 */
public object GameSignoutResponseKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: ai.popberry.proto.game.joker.v1alpha1.GameSignoutResponse.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
      internal fun _create(builder: ai.popberry.proto.game.joker.v1alpha1.GameSignoutResponse.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
  @kotlin.PublishedApi
    internal fun _build(): ai.popberry.proto.game.joker.v1alpha1.GameSignoutResponse = _builder.build()

    /**
     * `bool is_success = 1 [json_name = "isSuccess"];`
     */
    public var isSuccess: kotlin.Boolean
      @kotlin.jvm.JvmName("getIsSuccess")
        get() = _builder.isSuccess
      @kotlin.jvm.JvmName("setIsSuccess")
        set(value) {
        _builder.isSuccess = value
      }
    /**
     * <code>bool is_success = 1 [json_name = "isSuccess"];</code>
     * @return This builder for chaining.
     */
    public fun clearIsSuccess() {
      _builder.clearIsSuccess()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun ai.popberry.proto.game.joker.v1alpha1.GameSignoutResponse.copy(block: `ai.popberry.proto.game.joker.v1alpha1`.GameSignoutResponseKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.game.joker.v1alpha1.GameSignoutResponse =
  `ai.popberry.proto.game.joker.v1alpha1`.GameSignoutResponseKt.Dsl._create(this.toBuilder()).apply { block() }._build()

