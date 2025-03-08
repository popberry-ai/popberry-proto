// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: wallet/v1alpha1/wallet.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package ai.popberry.proto.wallet.v1alpha1;

@kotlin.jvm.JvmName("-initializeupdateTransactionResponseRequest")
public inline fun updateTransactionResponseRequest(block: ai.popberry.proto.wallet.v1alpha1.UpdateTransactionResponseRequestKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.wallet.v1alpha1.UpdateTransactionResponseRequest =
  ai.popberry.proto.wallet.v1alpha1.UpdateTransactionResponseRequestKt.Dsl._create(ai.popberry.proto.wallet.v1alpha1.UpdateTransactionResponseRequest.newBuilder()).apply { block() }._build()
/**
 * ```
 * Request message for updating the transaction response after another service has processed the request
 * ```
 *
 * Protobuf type `wallet.v1alpha1.UpdateTransactionResponseRequest`
 */
public object UpdateTransactionResponseRequestKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: ai.popberry.proto.wallet.v1alpha1.UpdateTransactionResponseRequest.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
      internal fun _create(builder: ai.popberry.proto.wallet.v1alpha1.UpdateTransactionResponseRequest.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
  @kotlin.PublishedApi
    internal fun _build(): ai.popberry.proto.wallet.v1alpha1.UpdateTransactionResponseRequest = _builder.build()

    /**
     * ```
     * Transaction ID to identify which record to update
     * ```
     *
     * `string txn_id = 1 [json_name = "txnId"];`
     */
    public var txnId: kotlin.String
      @kotlin.jvm.JvmName("getTxnId")
        get() = _builder.txnId
      @kotlin.jvm.JvmName("setTxnId")
        set(value) {
        _builder.txnId = value
      }
    /**
     * <pre>
     * Transaction ID to identify which record to update
     * </pre>
     *
     * <code>string txn_id = 1 [json_name = "txnId"];</code>
     * @return This builder for chaining.
     */
    public fun clearTxnId() {
      _builder.clearTxnId()
    }

    /**
     * ```
     * Updated response data (JSON or any other format)
     * ```
     *
     * `string response = 2 [json_name = "response"];`
     */
    public var response: kotlin.String
      @kotlin.jvm.JvmName("getResponse")
        get() = _builder.response
      @kotlin.jvm.JvmName("setResponse")
        set(value) {
        _builder.response = value
      }
    /**
     * <pre>
     * Updated response data (JSON or any other format)
     * </pre>
     *
     * <code>string response = 2 [json_name = "response"];</code>
     * @return This builder for chaining.
     */
    public fun clearResponse() {
      _builder.clearResponse()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun ai.popberry.proto.wallet.v1alpha1.UpdateTransactionResponseRequest.copy(block: `ai.popberry.proto.wallet.v1alpha1`.UpdateTransactionResponseRequestKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.wallet.v1alpha1.UpdateTransactionResponseRequest =
  `ai.popberry.proto.wallet.v1alpha1`.UpdateTransactionResponseRequestKt.Dsl._create(this.toBuilder()).apply { block() }._build()

