// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: bff/bridgr/v1alpha1/bridgr.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package ai.popberry.proto.bff.bridgr.v1alpha1;

@kotlin.jvm.JvmName("-initializeforwardingRecord")
public inline fun forwardingRecord(block: ai.popberry.proto.bff.bridgr.v1alpha1.ForwardingRecordKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.bff.bridgr.v1alpha1.ForwardingRecord =
  ai.popberry.proto.bff.bridgr.v1alpha1.ForwardingRecordKt.Dsl._create(ai.popberry.proto.bff.bridgr.v1alpha1.ForwardingRecord.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `bff.bridgr.v1alpha1.ForwardingRecord`
 */
public object ForwardingRecordKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: ai.popberry.proto.bff.bridgr.v1alpha1.ForwardingRecord.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
      internal fun _create(builder: ai.popberry.proto.bff.bridgr.v1alpha1.ForwardingRecord.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
  @kotlin.PublishedApi
    internal fun _build(): ai.popberry.proto.bff.bridgr.v1alpha1.ForwardingRecord = _builder.build()

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
     * `.bff.bridgr.v1alpha1.Type type = 2 [json_name = "type"];`
     */
    public var type: ai.popberry.proto.bff.bridgr.v1alpha1.Type
      @kotlin.jvm.JvmName("getType")
        get() = _builder.type
      @kotlin.jvm.JvmName("setType")
        set(value) {
        _builder.type = value
      }
    public var typeValue: kotlin.Int
      @kotlin.jvm.JvmName("getTypeValue")
        get() = _builder.typeValue
      @kotlin.jvm.JvmName("setTypeValue")
        set(value) {
        _builder.typeValue = value
      }
    /**
     * <code>.bff.bridgr.v1alpha1.Type type = 2 [json_name = "type"];</code>
     * @return This builder for chaining.
     */
    public fun clearType() {
      _builder.clearType()
    }

    /**
     * `string recipient = 3 [json_name = "recipient"];`
     */
    public var recipient: kotlin.String
      @kotlin.jvm.JvmName("getRecipient")
        get() = _builder.recipient
      @kotlin.jvm.JvmName("setRecipient")
        set(value) {
        _builder.recipient = value
      }
    /**
     * <code>string recipient = 3 [json_name = "recipient"];</code>
     * @return This builder for chaining.
     */
    public fun clearRecipient() {
      _builder.clearRecipient()
    }

    /**
     * `string content = 4 [json_name = "content"];`
     */
    public var content: kotlin.String
      @kotlin.jvm.JvmName("getContent")
        get() = _builder.content
      @kotlin.jvm.JvmName("setContent")
        set(value) {
        _builder.content = value
      }
    /**
     * <code>string content = 4 [json_name = "content"];</code>
     * @return This builder for chaining.
     */
    public fun clearContent() {
      _builder.clearContent()
    }

    /**
     * `.google.protobuf.Timestamp timestamp = 5 [json_name = "timestamp"];`
     */
    public var timestamp: com.google.protobuf.Timestamp
      @kotlin.jvm.JvmName("getTimestamp")
        get() = _builder.timestamp
      @kotlin.jvm.JvmName("setTimestamp")
        set(value) {
        _builder.timestamp = value
      }
    /**
     * <code>.google.protobuf.Timestamp timestamp = 5 [json_name = "timestamp"];</code>
     * @return This builder for chaining.
     */
    public fun clearTimestamp() {
      _builder.clearTimestamp()
    }
    /**
     * <code>.google.protobuf.Timestamp timestamp = 5 [json_name = "timestamp"];</code>
     * @return Whether the timestamp field is set.
     * @return This builder for chaining.
     */
    public fun hasTimestamp(): kotlin.Boolean {
      return _builder.hasTimestamp()
    }

    public val ForwardingRecordKt.Dsl.timestampOrNull: com.google.protobuf.Timestamp?
      get() = _builder.timestampOrNull

    /**
     * `bool success = 6 [json_name = "success"];`
     */
    public var success: kotlin.Boolean
      @kotlin.jvm.JvmName("getSuccess")
        get() = _builder.success
      @kotlin.jvm.JvmName("setSuccess")
        set(value) {
        _builder.success = value
      }
    /**
     * <code>bool success = 6 [json_name = "success"];</code>
     * @return This builder for chaining.
     */
    public fun clearSuccess() {
      _builder.clearSuccess()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun ai.popberry.proto.bff.bridgr.v1alpha1.ForwardingRecord.copy(block: `ai.popberry.proto.bff.bridgr.v1alpha1`.ForwardingRecordKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.bff.bridgr.v1alpha1.ForwardingRecord =
  `ai.popberry.proto.bff.bridgr.v1alpha1`.ForwardingRecordKt.Dsl._create(this.toBuilder()).apply { block() }._build()

public val ai.popberry.proto.bff.bridgr.v1alpha1.ForwardingRecordOrBuilder.timestampOrNull: com.google.protobuf.Timestamp?
  get() = if (hasTimestamp()) getTimestamp() else null

