package ai.popberry.proto.wallet.v1alpha1

import ai.popberry.proto.wallet.v1alpha1.WalletServiceGrpc.getServiceDescriptor
import io.grpc.CallOptions
import io.grpc.CallOptions.DEFAULT
import io.grpc.Channel
import io.grpc.Metadata
import io.grpc.MethodDescriptor
import io.grpc.ServerServiceDefinition
import io.grpc.ServerServiceDefinition.builder
import io.grpc.ServiceDescriptor
import io.grpc.Status.UNIMPLEMENTED
import io.grpc.StatusException
import io.grpc.kotlin.AbstractCoroutineServerImpl
import io.grpc.kotlin.AbstractCoroutineStub
import io.grpc.kotlin.ClientCalls.unaryRpc
import io.grpc.kotlin.ServerCalls.unaryServerMethodDefinition
import io.grpc.kotlin.StubFor
import kotlin.String
import kotlin.coroutines.CoroutineContext
import kotlin.coroutines.EmptyCoroutineContext
import kotlin.jvm.JvmOverloads
import kotlin.jvm.JvmStatic

/**
 * Holder for Kotlin coroutine-based client and server APIs for wallet.v1alpha1.WalletService.
 */
public object WalletServiceGrpcKt {
  public const val SERVICE_NAME: String = WalletServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  public val updateBalanceMethod: MethodDescriptor<UpdateBalanceRequest, UpdateBalanceResponse>
    @JvmStatic
    get() = WalletServiceGrpc.getUpdateBalanceMethod()

  public val identifyUserMethod: MethodDescriptor<IdentifyUserRequest, IdentifyUserResponse>
    @JvmStatic
    get() = WalletServiceGrpc.getIdentifyUserMethod()

  public val updateTransactionResponseMethod:
      MethodDescriptor<UpdateTransactionResponseRequest, UpdateTransactionResponseResponse>
    @JvmStatic
    get() = WalletServiceGrpc.getUpdateTransactionResponseMethod()

  /**
   * A stub for issuing RPCs to a(n) wallet.v1alpha1.WalletService service as suspending coroutines.
   */
  @StubFor(WalletServiceGrpc::class)
  public class WalletServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<WalletServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions): WalletServiceCoroutineStub =
        WalletServiceCoroutineStub(channel, callOptions)

    /**
     * Executes this RPC and returns the response message, suspending until the RPC completes
     * with [`Status.OK`][io.grpc.Status].  If the RPC completes with another status, a
     * corresponding
     * [StatusException] is thrown.  If this coroutine is cancelled, the RPC is also cancelled
     * with the corresponding exception as a cause.
     *
     * @param request The request message to send to the server.
     *
     * @param headers Metadata to attach to the request.  Most users will not need this.
     *
     * @return The single response from the server.
     */
    public suspend fun updateBalance(request: UpdateBalanceRequest, headers: Metadata = Metadata()):
        UpdateBalanceResponse = unaryRpc(
      channel,
      WalletServiceGrpc.getUpdateBalanceMethod(),
      request,
      callOptions,
      headers
    )

    /**
     * Executes this RPC and returns the response message, suspending until the RPC completes
     * with [`Status.OK`][io.grpc.Status].  If the RPC completes with another status, a
     * corresponding
     * [StatusException] is thrown.  If this coroutine is cancelled, the RPC is also cancelled
     * with the corresponding exception as a cause.
     *
     * @param request The request message to send to the server.
     *
     * @param headers Metadata to attach to the request.  Most users will not need this.
     *
     * @return The single response from the server.
     */
    public suspend fun identifyUser(request: IdentifyUserRequest, headers: Metadata = Metadata()):
        IdentifyUserResponse = unaryRpc(
      channel,
      WalletServiceGrpc.getIdentifyUserMethod(),
      request,
      callOptions,
      headers
    )

    /**
     * Executes this RPC and returns the response message, suspending until the RPC completes
     * with [`Status.OK`][io.grpc.Status].  If the RPC completes with another status, a
     * corresponding
     * [StatusException] is thrown.  If this coroutine is cancelled, the RPC is also cancelled
     * with the corresponding exception as a cause.
     *
     * @param request The request message to send to the server.
     *
     * @param headers Metadata to attach to the request.  Most users will not need this.
     *
     * @return The single response from the server.
     */
    public suspend fun updateTransactionResponse(request: UpdateTransactionResponseRequest,
        headers: Metadata = Metadata()): UpdateTransactionResponseResponse = unaryRpc(
      channel,
      WalletServiceGrpc.getUpdateTransactionResponseMethod(),
      request,
      callOptions,
      headers
    )
  }

  /**
   * Skeletal implementation of the wallet.v1alpha1.WalletService service based on Kotlin
   * coroutines.
   */
  public abstract class WalletServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    /**
     * Returns the response to an RPC for wallet.v1alpha1.WalletService.UpdateBalance.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun updateBalance(request: UpdateBalanceRequest): UpdateBalanceResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method wallet.v1alpha1.WalletService.UpdateBalance is unimplemented"))

    /**
     * Returns the response to an RPC for wallet.v1alpha1.WalletService.IdentifyUser.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun identifyUser(request: IdentifyUserRequest): IdentifyUserResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method wallet.v1alpha1.WalletService.IdentifyUser is unimplemented"))

    /**
     * Returns the response to an RPC for wallet.v1alpha1.WalletService.UpdateTransactionResponse.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun updateTransactionResponse(request: UpdateTransactionResponseRequest):
        UpdateTransactionResponseResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method wallet.v1alpha1.WalletService.UpdateTransactionResponse is unimplemented"))

    final override fun bindService(): ServerServiceDefinition = builder(getServiceDescriptor())
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = WalletServiceGrpc.getUpdateBalanceMethod(),
      implementation = ::updateBalance
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = WalletServiceGrpc.getIdentifyUserMethod(),
      implementation = ::identifyUser
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = WalletServiceGrpc.getUpdateTransactionResponseMethod(),
      implementation = ::updateTransactionResponse
    )).build()
  }
}
