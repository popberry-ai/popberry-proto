// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: user/v1/user.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package ai.popberry.proto.user.v1;

@kotlin.jvm.JvmName("-initializeupdateUserRequest")
public inline fun updateUserRequest(block: ai.popberry.proto.user.v1.UpdateUserRequestKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.user.v1.UpdateUserRequest =
  ai.popberry.proto.user.v1.UpdateUserRequestKt.Dsl._create(ai.popberry.proto.user.v1.UpdateUserRequest.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `user.v1.UpdateUserRequest`
 */
public object UpdateUserRequestKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: ai.popberry.proto.user.v1.UpdateUserRequest.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
      internal fun _create(builder: ai.popberry.proto.user.v1.UpdateUserRequest.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
  @kotlin.PublishedApi
    internal fun _build(): ai.popberry.proto.user.v1.UpdateUserRequest = _builder.build()

    /**
     * `string id = 1 [json_name = "id"];`
     */
    public var id: kotlin.String
      @kotlin.jvm.JvmName("getId")
        get() = _builder.id
      @kotlin.jvm.JvmName("setId")
        set(value) {
        _builder.id = value
      }
    /**
     * <code>string id = 1 [json_name = "id"];</code>
     * @return This builder for chaining.
     */
    public fun clearId() {
      _builder.clearId()
    }

    /**
     * `.user.v1.User user = 2 [json_name = "user"];`
     */
    public var user: ai.popberry.proto.user.v1.User
      @kotlin.jvm.JvmName("getUser")
        get() = _builder.user
      @kotlin.jvm.JvmName("setUser")
        set(value) {
        _builder.user = value
      }
    /**
     * <code>.user.v1.User user = 2 [json_name = "user"];</code>
     * @return This builder for chaining.
     */
    public fun clearUser() {
      _builder.clearUser()
    }
    /**
     * <code>.user.v1.User user = 2 [json_name = "user"];</code>
     * @return Whether the user field is set.
     * @return This builder for chaining.
     */
    public fun hasUser(): kotlin.Boolean {
      return _builder.hasUser()
    }

    public val UpdateUserRequestKt.Dsl.userOrNull: ai.popberry.proto.user.v1.User?
      get() = _builder.userOrNull
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun ai.popberry.proto.user.v1.UpdateUserRequest.copy(block: `ai.popberry.proto.user.v1`.UpdateUserRequestKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.user.v1.UpdateUserRequest =
  `ai.popberry.proto.user.v1`.UpdateUserRequestKt.Dsl._create(this.toBuilder()).apply { block() }._build()

public val ai.popberry.proto.user.v1.UpdateUserRequestOrBuilder.userOrNull: ai.popberry.proto.user.v1.User?
  get() = if (hasUser()) getUser() else null

