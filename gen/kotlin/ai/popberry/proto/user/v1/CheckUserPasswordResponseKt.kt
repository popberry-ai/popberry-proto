// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: user/v1/user.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package ai.popberry.proto.user.v1;

@kotlin.jvm.JvmName("-initializecheckUserPasswordResponse")
public inline fun checkUserPasswordResponse(block: ai.popberry.proto.user.v1.CheckUserPasswordResponseKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.user.v1.CheckUserPasswordResponse =
  ai.popberry.proto.user.v1.CheckUserPasswordResponseKt.Dsl._create(ai.popberry.proto.user.v1.CheckUserPasswordResponse.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `user.v1.CheckUserPasswordResponse`
 */
public object CheckUserPasswordResponseKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: ai.popberry.proto.user.v1.CheckUserPasswordResponse.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
      internal fun _create(builder: ai.popberry.proto.user.v1.CheckUserPasswordResponse.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
  @kotlin.PublishedApi
    internal fun _build(): ai.popberry.proto.user.v1.CheckUserPasswordResponse = _builder.build()

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

    public val CheckUserPasswordResponseKt.Dsl.userOrNull: ai.popberry.proto.user.v1.User?
      get() = _builder.userOrNull
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun ai.popberry.proto.user.v1.CheckUserPasswordResponse.copy(block: `ai.popberry.proto.user.v1`.CheckUserPasswordResponseKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.user.v1.CheckUserPasswordResponse =
  `ai.popberry.proto.user.v1`.CheckUserPasswordResponseKt.Dsl._create(this.toBuilder()).apply { block() }._build()

public val ai.popberry.proto.user.v1.CheckUserPasswordResponseOrBuilder.userOrNull: ai.popberry.proto.user.v1.User?
  get() = if (hasUser()) getUser() else null

