// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: game/joker/v1alpha1/joker.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package ai.popberry.proto.game.joker.v1alpha1;

@kotlin.jvm.JvmName("-initializegetGameRoundStatusRequest")
public inline fun getGameRoundStatusRequest(block: ai.popberry.proto.game.joker.v1alpha1.GetGameRoundStatusRequestKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.game.joker.v1alpha1.GetGameRoundStatusRequest =
  ai.popberry.proto.game.joker.v1alpha1.GetGameRoundStatusRequestKt.Dsl._create(ai.popberry.proto.game.joker.v1alpha1.GetGameRoundStatusRequest.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `game.joker.v1alpha1.GetGameRoundStatusRequest`
 */
public object GetGameRoundStatusRequestKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: ai.popberry.proto.game.joker.v1alpha1.GetGameRoundStatusRequest.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
      internal fun _create(builder: ai.popberry.proto.game.joker.v1alpha1.GetGameRoundStatusRequest.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
  @kotlin.PublishedApi
    internal fun _build(): ai.popberry.proto.game.joker.v1alpha1.GetGameRoundStatusRequest = _builder.build()

    /**
     * `string round_id = 1 [json_name = "roundId"];`
     */
    public var roundId: kotlin.String
      @kotlin.jvm.JvmName("getRoundId")
        get() = _builder.roundId
      @kotlin.jvm.JvmName("setRoundId")
        set(value) {
        _builder.roundId = value
      }
    /**
     * <code>string round_id = 1 [json_name = "roundId"];</code>
     * @return This builder for chaining.
     */
    public fun clearRoundId() {
      _builder.clearRoundId()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun ai.popberry.proto.game.joker.v1alpha1.GetGameRoundStatusRequest.copy(block: `ai.popberry.proto.game.joker.v1alpha1`.GetGameRoundStatusRequestKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.game.joker.v1alpha1.GetGameRoundStatusRequest =
  `ai.popberry.proto.game.joker.v1alpha1`.GetGameRoundStatusRequestKt.Dsl._create(this.toBuilder()).apply { block() }._build()

