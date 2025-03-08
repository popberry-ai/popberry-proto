package ai.popberry.proto.bff.bridgr.v1alpha1

import ai.popberry.proto.bff.bridgr.v1alpha1.BridgrServiceGrpc.getServiceDescriptor
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
 * Holder for Kotlin coroutine-based client and server APIs for bff.bridgr.v1alpha1.BridgrService.
 */
public object BridgrServiceGrpcKt {
  public const val SERVICE_NAME: String = BridgrServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  public val registerDeviceMethod: MethodDescriptor<RegisterDeviceRequest, RegisterDeviceResponse>
    @JvmStatic
    get() = BridgrServiceGrpc.getRegisterDeviceMethod()

  public val unregisterDeviceMethod:
      MethodDescriptor<UnregisterDeviceRequest, UnregisterDeviceResponse>
    @JvmStatic
    get() = BridgrServiceGrpc.getUnregisterDeviceMethod()

  public val healthCheckMethod: MethodDescriptor<HealthCheckRequest, HealthCheckResponse>
    @JvmStatic
    get() = BridgrServiceGrpc.getHealthCheckMethod()

  public val getConfigMethod: MethodDescriptor<GetConfigRequest, GetConfigResponse>
    @JvmStatic
    get() = BridgrServiceGrpc.getGetConfigMethod()

  public val updateConfigMethod: MethodDescriptor<UpdateConfigRequest, UpdateConfigResponse>
    @JvmStatic
    get() = BridgrServiceGrpc.getUpdateConfigMethod()

  public val forwardSMSMethod: MethodDescriptor<ForwardSMSRequest, ForwardSMSResponse>
    @JvmStatic
    get() = BridgrServiceGrpc.getForwardSMSMethod()

  public val forwardNotificationMethod:
      MethodDescriptor<ForwardNotificationRequest, ForwardNotificationResponse>
    @JvmStatic
    get() = BridgrServiceGrpc.getForwardNotificationMethod()

  public val getForwardingHistoryMethod:
      MethodDescriptor<GetForwardingHistoryRequest, GetForwardingHistoryResponse>
    @JvmStatic
    get() = BridgrServiceGrpc.getGetForwardingHistoryMethod()

  /**
   * A stub for issuing RPCs to a(n) bff.bridgr.v1alpha1.BridgrService service as suspending
   * coroutines.
   */
  @StubFor(BridgrServiceGrpc::class)
  public class BridgrServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<BridgrServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions): BridgrServiceCoroutineStub =
        BridgrServiceCoroutineStub(channel, callOptions)

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
    public suspend fun registerDevice(request: RegisterDeviceRequest, headers: Metadata =
        Metadata()): RegisterDeviceResponse = unaryRpc(
      channel,
      BridgrServiceGrpc.getRegisterDeviceMethod(),
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
    public suspend fun unregisterDevice(request: UnregisterDeviceRequest, headers: Metadata =
        Metadata()): UnregisterDeviceResponse = unaryRpc(
      channel,
      BridgrServiceGrpc.getUnregisterDeviceMethod(),
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
    public suspend fun healthCheck(request: HealthCheckRequest, headers: Metadata = Metadata()):
        HealthCheckResponse = unaryRpc(
      channel,
      BridgrServiceGrpc.getHealthCheckMethod(),
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
    public suspend fun getConfig(request: GetConfigRequest, headers: Metadata = Metadata()):
        GetConfigResponse = unaryRpc(
      channel,
      BridgrServiceGrpc.getGetConfigMethod(),
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
    public suspend fun updateConfig(request: UpdateConfigRequest, headers: Metadata = Metadata()):
        UpdateConfigResponse = unaryRpc(
      channel,
      BridgrServiceGrpc.getUpdateConfigMethod(),
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
    public suspend fun forwardSMS(request: ForwardSMSRequest, headers: Metadata = Metadata()):
        ForwardSMSResponse = unaryRpc(
      channel,
      BridgrServiceGrpc.getForwardSMSMethod(),
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
    public suspend fun forwardNotification(request: ForwardNotificationRequest, headers: Metadata =
        Metadata()): ForwardNotificationResponse = unaryRpc(
      channel,
      BridgrServiceGrpc.getForwardNotificationMethod(),
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
    public suspend fun getForwardingHistory(request: GetForwardingHistoryRequest, headers: Metadata
        = Metadata()): GetForwardingHistoryResponse = unaryRpc(
      channel,
      BridgrServiceGrpc.getGetForwardingHistoryMethod(),
      request,
      callOptions,
      headers
    )
  }

  /**
   * Skeletal implementation of the bff.bridgr.v1alpha1.BridgrService service based on Kotlin
   * coroutines.
   */
  public abstract class BridgrServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    /**
     * Returns the response to an RPC for bff.bridgr.v1alpha1.BridgrService.RegisterDevice.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun registerDevice(request: RegisterDeviceRequest): RegisterDeviceResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.v1alpha1.BridgrService.RegisterDevice is unimplemented"))

    /**
     * Returns the response to an RPC for bff.bridgr.v1alpha1.BridgrService.UnregisterDevice.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun unregisterDevice(request: UnregisterDeviceRequest):
        UnregisterDeviceResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.v1alpha1.BridgrService.UnregisterDevice is unimplemented"))

    /**
     * Returns the response to an RPC for bff.bridgr.v1alpha1.BridgrService.HealthCheck.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun healthCheck(request: HealthCheckRequest): HealthCheckResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.v1alpha1.BridgrService.HealthCheck is unimplemented"))

    /**
     * Returns the response to an RPC for bff.bridgr.v1alpha1.BridgrService.GetConfig.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getConfig(request: GetConfigRequest): GetConfigResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.v1alpha1.BridgrService.GetConfig is unimplemented"))

    /**
     * Returns the response to an RPC for bff.bridgr.v1alpha1.BridgrService.UpdateConfig.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun updateConfig(request: UpdateConfigRequest): UpdateConfigResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.v1alpha1.BridgrService.UpdateConfig is unimplemented"))

    /**
     * Returns the response to an RPC for bff.bridgr.v1alpha1.BridgrService.ForwardSMS.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun forwardSMS(request: ForwardSMSRequest): ForwardSMSResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.v1alpha1.BridgrService.ForwardSMS is unimplemented"))

    /**
     * Returns the response to an RPC for bff.bridgr.v1alpha1.BridgrService.ForwardNotification.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun forwardNotification(request: ForwardNotificationRequest):
        ForwardNotificationResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.v1alpha1.BridgrService.ForwardNotification is unimplemented"))

    /**
     * Returns the response to an RPC for bff.bridgr.v1alpha1.BridgrService.GetForwardingHistory.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getForwardingHistory(request: GetForwardingHistoryRequest):
        GetForwardingHistoryResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.v1alpha1.BridgrService.GetForwardingHistory is unimplemented"))

    final override fun bindService(): ServerServiceDefinition = builder(getServiceDescriptor())
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = BridgrServiceGrpc.getRegisterDeviceMethod(),
      implementation = ::registerDevice
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = BridgrServiceGrpc.getUnregisterDeviceMethod(),
      implementation = ::unregisterDevice
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = BridgrServiceGrpc.getHealthCheckMethod(),
      implementation = ::healthCheck
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = BridgrServiceGrpc.getGetConfigMethod(),
      implementation = ::getConfig
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = BridgrServiceGrpc.getUpdateConfigMethod(),
      implementation = ::updateConfig
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = BridgrServiceGrpc.getForwardSMSMethod(),
      implementation = ::forwardSMS
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = BridgrServiceGrpc.getForwardNotificationMethod(),
      implementation = ::forwardNotification
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = BridgrServiceGrpc.getGetForwardingHistoryMethod(),
      implementation = ::getForwardingHistory
    )).build()
  }
}
