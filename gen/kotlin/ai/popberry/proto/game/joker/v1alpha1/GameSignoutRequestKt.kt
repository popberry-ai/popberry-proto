// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: game/joker/v1alpha1/joker.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package ai.popberry.proto.game.joker.v1alpha1;

@kotlin.jvm.JvmName("-initializegameSignoutRequest")
public inline fun gameSignoutRequest(block: ai.popberry.proto.game.joker.v1alpha1.GameSignoutRequestKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.game.joker.v1alpha1.GameSignoutRequest =
  ai.popberry.proto.game.joker.v1alpha1.GameSignoutRequestKt.Dsl._create(ai.popberry.proto.game.joker.v1alpha1.GameSignoutRequest.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `game.joker.v1alpha1.GameSignoutRequest`
 */
public object GameSignoutRequestKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: ai.popberry.proto.game.joker.v1alpha1.GameSignoutRequest.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
      internal fun _create(builder: ai.popberry.proto.game.joker.v1alpha1.GameSignoutRequest.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
  @kotlin.PublishedApi
    internal fun _build(): ai.popberry.proto.game.joker.v1alpha1.GameSignoutRequest = _builder.build()

    /**
     * `string username = 1 [json_name = "username"];`
     */
    public var username: kotlin.String
      @kotlin.jvm.JvmName("getUsername")
        get() = _builder.username
      @kotlin.jvm.JvmName("setUsername")
        set(value) {
        _builder.username = value
      }
    /**
     * <code>string username = 1 [json_name = "username"];</code>
     * @return This builder for chaining.
     */
    public fun clearUsername() {
      _builder.clearUsername()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun ai.popberry.proto.game.joker.v1alpha1.GameSignoutRequest.copy(block: `ai.popberry.proto.game.joker.v1alpha1`.GameSignoutRequestKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.game.joker.v1alpha1.GameSignoutRequest =
  `ai.popberry.proto.game.joker.v1alpha1`.GameSignoutRequestKt.Dsl._create(this.toBuilder()).apply { block() }._build()

