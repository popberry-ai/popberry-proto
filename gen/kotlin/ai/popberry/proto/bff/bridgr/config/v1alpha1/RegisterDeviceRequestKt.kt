// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: bff/bridgr/config/v1alpha1/service.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package ai.popberry.proto.bff.bridgr.config.v1alpha1;

@kotlin.jvm.JvmName("-initializeregisterDeviceRequest")
public inline fun registerDeviceRequest(block: ai.popberry.proto.bff.bridgr.config.v1alpha1.RegisterDeviceRequestKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.bff.bridgr.config.v1alpha1.RegisterDeviceRequest =
  ai.popberry.proto.bff.bridgr.config.v1alpha1.RegisterDeviceRequestKt.Dsl._create(ai.popberry.proto.bff.bridgr.config.v1alpha1.RegisterDeviceRequest.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `bff.bridgr.config.v1alpha1.RegisterDeviceRequest`
 */
public object RegisterDeviceRequestKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: ai.popberry.proto.bff.bridgr.config.v1alpha1.RegisterDeviceRequest.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
      internal fun _create(builder: ai.popberry.proto.bff.bridgr.config.v1alpha1.RegisterDeviceRequest.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
  @kotlin.PublishedApi
    internal fun _build(): ai.popberry.proto.bff.bridgr.config.v1alpha1.RegisterDeviceRequest = _builder.build()

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
     * `string device_name = 2 [json_name = "deviceName"];`
     */
    public var deviceName: kotlin.String
      @kotlin.jvm.JvmName("getDeviceName")
        get() = _builder.deviceName
      @kotlin.jvm.JvmName("setDeviceName")
        set(value) {
        _builder.deviceName = value
      }
    /**
     * <code>string device_name = 2 [json_name = "deviceName"];</code>
     * @return This builder for chaining.
     */
    public fun clearDeviceName() {
      _builder.clearDeviceName()
    }

    /**
     * `string device_model = 3 [json_name = "deviceModel"];`
     */
    public var deviceModel: kotlin.String
      @kotlin.jvm.JvmName("getDeviceModel")
        get() = _builder.deviceModel
      @kotlin.jvm.JvmName("setDeviceModel")
        set(value) {
        _builder.deviceModel = value
      }
    /**
     * <code>string device_model = 3 [json_name = "deviceModel"];</code>
     * @return This builder for chaining.
     */
    public fun clearDeviceModel() {
      _builder.clearDeviceModel()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun ai.popberry.proto.bff.bridgr.config.v1alpha1.RegisterDeviceRequest.copy(block: `ai.popberry.proto.bff.bridgr.config.v1alpha1`.RegisterDeviceRequestKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.bff.bridgr.config.v1alpha1.RegisterDeviceRequest =
  `ai.popberry.proto.bff.bridgr.config.v1alpha1`.RegisterDeviceRequestKt.Dsl._create(this.toBuilder()).apply { block() }._build()

