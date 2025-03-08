package ai.popberry.proto.bff.bridgr.health.v1alpha1

import ai.popberry.proto.bff.bridgr.health.v1alpha1.HealthServiceGrpc.getServiceDescriptor
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
import io.grpc.kotlin.ClientCalls.bidiStreamingRpc
import io.grpc.kotlin.ClientCalls.unaryRpc
import io.grpc.kotlin.ServerCalls.bidiStreamingServerMethodDefinition
import io.grpc.kotlin.ServerCalls.unaryServerMethodDefinition
import io.grpc.kotlin.StubFor
import kotlin.String
import kotlin.coroutines.CoroutineContext
import kotlin.coroutines.EmptyCoroutineContext
import kotlin.jvm.JvmOverloads
import kotlin.jvm.JvmStatic
import kotlinx.coroutines.flow.Flow

/**
 * Holder for Kotlin coroutine-based client and server APIs for
 * bff.bridgr.health.v1alpha1.HealthService.
 */
public object HealthServiceGrpcKt {
  public const val SERVICE_NAME: String = HealthServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  public val reportHealthMethod: MethodDescriptor<ReportHealthRequest, ReportHealthResponse>
    @JvmStatic
    get() = HealthServiceGrpc.getReportHealthMethod()

  public val deviceHealthMethod: MethodDescriptor<DeviceHealthRequest, DeviceHealthResponse>
    @JvmStatic
    get() = HealthServiceGrpc.getDeviceHealthMethod()

  /**
   * A stub for issuing RPCs to a(n) bff.bridgr.health.v1alpha1.HealthService service as suspending
   * coroutines.
   */
  @StubFor(HealthServiceGrpc::class)
  public class HealthServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<HealthServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions): HealthServiceCoroutineStub =
        HealthServiceCoroutineStub(channel, callOptions)

    /**
     * Returns a [Flow] that, when collected, executes this RPC and emits responses from the
     * server as they arrive.  That flow finishes normally if the server closes its response with
     * [`Status.OK`][io.grpc.Status], and fails by throwing a [StatusException] otherwise.  If
     * collecting the flow downstream fails exceptionally (including via cancellation), the RPC
     * is cancelled with that exception as a cause.
     *
     * The [Flow] of requests is collected once each time the [Flow] of responses is
     * collected. If collection of the [Flow] of responses completes normally or
     * exceptionally before collection of `requests` completes, the collection of
     * `requests` is cancelled.  If the collection of `requests` completes
     * exceptionally for any other reason, then the collection of the [Flow] of responses
     * completes exceptionally for the same reason and the RPC is cancelled with that reason.
     *
     * @param requests A [Flow] of request messages.
     *
     * @param headers Metadata to attach to the request.  Most users will not need this.
     *
     * @return A flow that, when collected, emits the responses from the server.
     */
    public fun reportHealth(requests: Flow<ReportHealthRequest>, headers: Metadata = Metadata()):
        Flow<ReportHealthResponse> = bidiStreamingRpc(
      channel,
      HealthServiceGrpc.getReportHealthMethod(),
      requests,
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
    public suspend fun deviceHealth(request: DeviceHealthRequest, headers: Metadata = Metadata()):
        DeviceHealthResponse = unaryRpc(
      channel,
      HealthServiceGrpc.getDeviceHealthMethod(),
      request,
      callOptions,
      headers
    )
  }

  /**
   * Skeletal implementation of the bff.bridgr.health.v1alpha1.HealthService service based on Kotlin
   * coroutines.
   */
  public abstract class HealthServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    /**
     * Returns a [Flow] of responses to an RPC for
     * bff.bridgr.health.v1alpha1.HealthService.ReportHealth.
     *
     * If creating or collecting the returned flow fails with a [StatusException], the RPC
     * will fail with the corresponding [io.grpc.Status].  If it fails with a
     * [java.util.concurrent.CancellationException], the RPC will fail with status
     * `Status.CANCELLED`.  If creating
     * or collecting the returned flow fails for any other reason, the RPC will fail with
     * `Status.UNKNOWN` with the exception as a cause.
     *
     * @param requests A [Flow] of requests from the client.  This flow can be
     *        collected only once and throws [java.lang.IllegalStateException] on attempts to
     * collect
     *        it more than once.
     */
    public open fun reportHealth(requests: Flow<ReportHealthRequest>): Flow<ReportHealthResponse> =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.health.v1alpha1.HealthService.ReportHealth is unimplemented"))

    /**
     * Returns the response to an RPC for bff.bridgr.health.v1alpha1.HealthService.DeviceHealth.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun deviceHealth(request: DeviceHealthRequest): DeviceHealthResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.health.v1alpha1.HealthService.DeviceHealth is unimplemented"))

    final override fun bindService(): ServerServiceDefinition = builder(getServiceDescriptor())
      .addMethod(bidiStreamingServerMethodDefinition(
      context = this.context,
      descriptor = HealthServiceGrpc.getReportHealthMethod(),
      implementation = ::reportHealth
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = HealthServiceGrpc.getDeviceHealthMethod(),
      implementation = ::deviceHealth
    )).build()
  }
}
