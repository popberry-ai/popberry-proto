// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: bff/bridgr/config/v1alpha1/service.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package ai.popberry.proto.bff.bridgr.config.v1alpha1;

@kotlin.jvm.JvmName("-initializeassignDeviceGroupRequest")
public inline fun assignDeviceGroupRequest(block: ai.popberry.proto.bff.bridgr.config.v1alpha1.AssignDeviceGroupRequestKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.bff.bridgr.config.v1alpha1.AssignDeviceGroupRequest =
  ai.popberry.proto.bff.bridgr.config.v1alpha1.AssignDeviceGroupRequestKt.Dsl._create(ai.popberry.proto.bff.bridgr.config.v1alpha1.AssignDeviceGroupRequest.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `bff.bridgr.config.v1alpha1.AssignDeviceGroupRequest`
 */
public object AssignDeviceGroupRequestKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: ai.popberry.proto.bff.bridgr.config.v1alpha1.AssignDeviceGroupRequest.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
      internal fun _create(builder: ai.popberry.proto.bff.bridgr.config.v1alpha1.AssignDeviceGroupRequest.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
  @kotlin.PublishedApi
    internal fun _build(): ai.popberry.proto.bff.bridgr.config.v1alpha1.AssignDeviceGroupRequest = _builder.build()

    /**
     * `string device_id = 1 [json_name = "deviceId"];`
     */
    public var deviceId: kotlin.String
      @kotlin.jvm.JvmName("getDeviceId")
        get() = _builder.deviceId
      @kotlin.jvm.JvmName("setDeviceId")
        set(value) {
        _builder.deviceId = value
      }
    /**
     * <code>string device_id = 1 [json_name = "deviceId"];</code>
     * @return This builder for chaining.
     */
    public fun clearDeviceId() {
      _builder.clearDeviceId()
    }

    /**
     * `string group_name = 2 [json_name = "groupName"];`
     */
    public var groupName: kotlin.String
      @kotlin.jvm.JvmName("getGroupName")
        get() = _builder.groupName
      @kotlin.jvm.JvmName("setGroupName")
        set(value) {
        _builder.groupName = value
      }
    /**
     * <code>string group_name = 2 [json_name = "groupName"];</code>
     * @return This builder for chaining.
     */
    public fun clearGroupName() {
      _builder.clearGroupName()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun ai.popberry.proto.bff.bridgr.config.v1alpha1.AssignDeviceGroupRequest.copy(block: `ai.popberry.proto.bff.bridgr.config.v1alpha1`.AssignDeviceGroupRequestKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.bff.bridgr.config.v1alpha1.AssignDeviceGroupRequest =
  `ai.popberry.proto.bff.bridgr.config.v1alpha1`.AssignDeviceGroupRequestKt.Dsl._create(this.toBuilder()).apply { block() }._build()

