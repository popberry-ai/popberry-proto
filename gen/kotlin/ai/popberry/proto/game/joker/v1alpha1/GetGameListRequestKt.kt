// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: game/joker/v1alpha1/joker.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package ai.popberry.proto.game.joker.v1alpha1;

@kotlin.jvm.JvmName("-initializegetGameListRequest")
public inline fun getGameListRequest(block: ai.popberry.proto.game.joker.v1alpha1.GetGameListRequestKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.game.joker.v1alpha1.GetGameListRequest =
  ai.popberry.proto.game.joker.v1alpha1.GetGameListRequestKt.Dsl._create(ai.popberry.proto.game.joker.v1alpha1.GetGameListRequest.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `game.joker.v1alpha1.GetGameListRequest`
 */
public object GetGameListRequestKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: ai.popberry.proto.game.joker.v1alpha1.GetGameListRequest.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
      internal fun _create(builder: ai.popberry.proto.game.joker.v1alpha1.GetGameListRequest.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
  @kotlin.PublishedApi
    internal fun _build(): ai.popberry.proto.game.joker.v1alpha1.GetGameListRequest = _builder.build()

    /**
     * `int32 page_size = 1 [json_name = "pageSize"];`
     */
    public var pageSize: kotlin.Int
      @kotlin.jvm.JvmName("getPageSize")
        get() = _builder.pageSize
      @kotlin.jvm.JvmName("setPageSize")
        set(value) {
        _builder.pageSize = value
      }
    /**
     * <code>int32 page_size = 1 [json_name = "pageSize"];</code>
     * @return This builder for chaining.
     */
    public fun clearPageSize() {
      _builder.clearPageSize()
    }

    /**
     * `string page_token = 2 [json_name = "pageToken"];`
     */
    public var pageToken: kotlin.String
      @kotlin.jvm.JvmName("getPageToken")
        get() = _builder.pageToken
      @kotlin.jvm.JvmName("setPageToken")
        set(value) {
        _builder.pageToken = value
      }
    /**
     * <code>string page_token = 2 [json_name = "pageToken"];</code>
     * @return This builder for chaining.
     */
    public fun clearPageToken() {
      _builder.clearPageToken()
    }

    /**
     * `optional string query = 3 [json_name = "query"];`
     */
    public var query: kotlin.String
      @kotlin.jvm.JvmName("getQuery")
        get() = _builder.query
      @kotlin.jvm.JvmName("setQuery")
        set(value) {
        _builder.query = value
      }
    /**
     * <code>optional string query = 3 [json_name = "query"];</code>
     * @return This builder for chaining.
     */
    public fun clearQuery() {
      _builder.clearQuery()
    }
    /**
     * <code>optional string query = 3 [json_name = "query"];</code>
     * @return Whether the query field is set.
     * @return This builder for chaining.
     */
    public fun hasQuery(): kotlin.Boolean {
      return _builder.hasQuery()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun ai.popberry.proto.game.joker.v1alpha1.GetGameListRequest.copy(block: `ai.popberry.proto.game.joker.v1alpha1`.GetGameListRequestKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.game.joker.v1alpha1.GetGameListRequest =
  `ai.popberry.proto.game.joker.v1alpha1`.GetGameListRequestKt.Dsl._create(this.toBuilder()).apply { block() }._build()

