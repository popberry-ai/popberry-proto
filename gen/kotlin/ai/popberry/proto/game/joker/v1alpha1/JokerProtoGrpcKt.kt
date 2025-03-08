package ai.popberry.proto.game.joker.v1alpha1

import ai.popberry.proto.game.joker.v1alpha1.JokerServiceGrpc.getServiceDescriptor
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
 * Holder for Kotlin coroutine-based client and server APIs for game.joker.v1alpha1.JokerService.
 */
public object JokerServiceGrpcKt {
  public const val SERVICE_NAME: String = JokerServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  public val gameLoginMethod: MethodDescriptor<GameLoginRequest, GameLoginResponse>
    @JvmStatic
    get() = JokerServiceGrpc.getGameLoginMethod()

  public val getGameListMethod: MethodDescriptor<GetGameListRequest, GetGameListResponse>
    @JvmStatic
    get() = JokerServiceGrpc.getGetGameListMethod()

  public val getGameRoundStatusMethod:
      MethodDescriptor<GetGameRoundStatusRequest, GetGameRoundStatusResponse>
    @JvmStatic
    get() = JokerServiceGrpc.getGetGameRoundStatusMethod()

  public val getGameDetailMethod: MethodDescriptor<GetGameDetailRequest, GetGameDetailResponse>
    @JvmStatic
    get() = JokerServiceGrpc.getGetGameDetailMethod()

  public val gameSignoutMethod: MethodDescriptor<GameSignoutRequest, GameSignoutResponse>
    @JvmStatic
    get() = JokerServiceGrpc.getGameSignoutMethod()

  /**
   * A stub for issuing RPCs to a(n) game.joker.v1alpha1.JokerService service as suspending
   * coroutines.
   */
  @StubFor(JokerServiceGrpc::class)
  public class JokerServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<JokerServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions): JokerServiceCoroutineStub =
        JokerServiceCoroutineStub(channel, callOptions)

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
    public suspend fun gameLogin(request: GameLoginRequest, headers: Metadata = Metadata()):
        GameLoginResponse = unaryRpc(
      channel,
      JokerServiceGrpc.getGameLoginMethod(),
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
    public suspend fun getGameList(request: GetGameListRequest, headers: Metadata = Metadata()):
        GetGameListResponse = unaryRpc(
      channel,
      JokerServiceGrpc.getGetGameListMethod(),
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
    public suspend fun getGameRoundStatus(request: GetGameRoundStatusRequest, headers: Metadata =
        Metadata()): GetGameRoundStatusResponse = unaryRpc(
      channel,
      JokerServiceGrpc.getGetGameRoundStatusMethod(),
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
    public suspend fun getGameDetail(request: GetGameDetailRequest, headers: Metadata = Metadata()):
        GetGameDetailResponse = unaryRpc(
      channel,
      JokerServiceGrpc.getGetGameDetailMethod(),
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
    public suspend fun gameSignout(request: GameSignoutRequest, headers: Metadata = Metadata()):
        GameSignoutResponse = unaryRpc(
      channel,
      JokerServiceGrpc.getGameSignoutMethod(),
      request,
      callOptions,
      headers
    )
  }

  /**
   * Skeletal implementation of the game.joker.v1alpha1.JokerService service based on Kotlin
   * coroutines.
   */
  public abstract class JokerServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    /**
     * Returns the response to an RPC for game.joker.v1alpha1.JokerService.GameLogin.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun gameLogin(request: GameLoginRequest): GameLoginResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method game.joker.v1alpha1.JokerService.GameLogin is unimplemented"))

    /**
     * Returns the response to an RPC for game.joker.v1alpha1.JokerService.GetGameList.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getGameList(request: GetGameListRequest): GetGameListResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method game.joker.v1alpha1.JokerService.GetGameList is unimplemented"))

    /**
     * Returns the response to an RPC for game.joker.v1alpha1.JokerService.GetGameRoundStatus.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getGameRoundStatus(request: GetGameRoundStatusRequest):
        GetGameRoundStatusResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method game.joker.v1alpha1.JokerService.GetGameRoundStatus is unimplemented"))

    /**
     * Returns the response to an RPC for game.joker.v1alpha1.JokerService.GetGameDetail.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getGameDetail(request: GetGameDetailRequest): GetGameDetailResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method game.joker.v1alpha1.JokerService.GetGameDetail is unimplemented"))

    /**
     * Returns the response to an RPC for game.joker.v1alpha1.JokerService.GameSignout.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun gameSignout(request: GameSignoutRequest): GameSignoutResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method game.joker.v1alpha1.JokerService.GameSignout is unimplemented"))

    final override fun bindService(): ServerServiceDefinition = builder(getServiceDescriptor())
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = JokerServiceGrpc.getGameLoginMethod(),
      implementation = ::gameLogin
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = JokerServiceGrpc.getGetGameListMethod(),
      implementation = ::getGameList
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = JokerServiceGrpc.getGetGameRoundStatusMethod(),
      implementation = ::getGameRoundStatus
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = JokerServiceGrpc.getGetGameDetailMethod(),
      implementation = ::getGameDetail
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = JokerServiceGrpc.getGameSignoutMethod(),
      implementation = ::gameSignout
    )).build()
  }
}
