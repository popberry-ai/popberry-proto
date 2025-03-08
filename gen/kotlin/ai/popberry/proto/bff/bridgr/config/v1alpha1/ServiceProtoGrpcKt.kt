package ai.popberry.proto.bff.bridgr.config.v1alpha1

import ai.popberry.proto.bff.bridgr.config.v1alpha1.ConfigServiceGrpc.getServiceDescriptor
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
import io.grpc.kotlin.ClientCalls.serverStreamingRpc
import io.grpc.kotlin.ClientCalls.unaryRpc
import io.grpc.kotlin.ServerCalls.serverStreamingServerMethodDefinition
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
 * bff.bridgr.config.v1alpha1.ConfigService.
 */
public object ConfigServiceGrpcKt {
  public const val SERVICE_NAME: String = ConfigServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  public val registerDeviceMethod: MethodDescriptor<RegisterDeviceRequest, RegisterDeviceResponse>
    @JvmStatic
    get() = ConfigServiceGrpc.getRegisterDeviceMethod()

  public val verifyDeviceMethod: MethodDescriptor<VerifyDeviceRequest, VerifyDeviceResponse>
    @JvmStatic
    get() = ConfigServiceGrpc.getVerifyDeviceMethod()

  public val assignDeviceGroupMethod:
      MethodDescriptor<AssignDeviceGroupRequest, AssignDeviceGroupResponse>
    @JvmStatic
    get() = ConfigServiceGrpc.getAssignDeviceGroupMethod()

  public val updateForwardingRulesMethod:
      MethodDescriptor<UpdateForwardingRulesRequest, UpdateForwardingRulesResponse>
    @JvmStatic
    get() = ConfigServiceGrpc.getUpdateForwardingRulesMethod()

  public val getForwardingRulesMethod:
      MethodDescriptor<GetForwardingRulesRequest, GetForwardingRulesResponse>
    @JvmStatic
    get() = ConfigServiceGrpc.getGetForwardingRulesMethod()

  public val getConfigMethod: MethodDescriptor<GetConfigRequest, GetConfigResponse>
    @JvmStatic
    get() = ConfigServiceGrpc.getGetConfigMethod()

  public val updateConfigMethod: MethodDescriptor<UpdateConfigRequest, UpdateConfigResponse>
    @JvmStatic
    get() = ConfigServiceGrpc.getUpdateConfigMethod()

  /**
   * A stub for issuing RPCs to a(n) bff.bridgr.config.v1alpha1.ConfigService service as suspending
   * coroutines.
   */
  @StubFor(ConfigServiceGrpc::class)
  public class ConfigServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<ConfigServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions): ConfigServiceCoroutineStub =
        ConfigServiceCoroutineStub(channel, callOptions)

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
      ConfigServiceGrpc.getRegisterDeviceMethod(),
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
    public suspend fun verifyDevice(request: VerifyDeviceRequest, headers: Metadata = Metadata()):
        VerifyDeviceResponse = unaryRpc(
      channel,
      ConfigServiceGrpc.getVerifyDeviceMethod(),
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
    public suspend fun assignDeviceGroup(request: AssignDeviceGroupRequest, headers: Metadata =
        Metadata()): AssignDeviceGroupResponse = unaryRpc(
      channel,
      ConfigServiceGrpc.getAssignDeviceGroupMethod(),
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
    public suspend fun updateForwardingRules(request: UpdateForwardingRulesRequest,
        headers: Metadata = Metadata()): UpdateForwardingRulesResponse = unaryRpc(
      channel,
      ConfigServiceGrpc.getUpdateForwardingRulesMethod(),
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
    public suspend fun getForwardingRules(request: GetForwardingRulesRequest, headers: Metadata =
        Metadata()): GetForwardingRulesResponse = unaryRpc(
      channel,
      ConfigServiceGrpc.getGetForwardingRulesMethod(),
      request,
      callOptions,
      headers
    )

    /**
     * Returns a [Flow] that, when collected, executes this RPC and emits responses from the
     * server as they arrive.  That flow finishes normally if the server closes its response with
     * [`Status.OK`][io.grpc.Status], and fails by throwing a [StatusException] otherwise.  If
     * collecting the flow downstream fails exceptionally (including via cancellation), the RPC
     * is cancelled with that exception as a cause.
     *
     * @param request The request message to send to the server.
     *
     * @param headers Metadata to attach to the request.  Most users will not need this.
     *
     * @return A flow that, when collected, emits the responses from the server.
     */
    public fun getConfig(request: GetConfigRequest, headers: Metadata = Metadata()):
        Flow<GetConfigResponse> = serverStreamingRpc(
      channel,
      ConfigServiceGrpc.getGetConfigMethod(),
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
      ConfigServiceGrpc.getUpdateConfigMethod(),
      request,
      callOptions,
      headers
    )
  }

  /**
   * Skeletal implementation of the bff.bridgr.config.v1alpha1.ConfigService service based on Kotlin
   * coroutines.
   */
  public abstract class ConfigServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    /**
     * Returns the response to an RPC for bff.bridgr.config.v1alpha1.ConfigService.RegisterDevice.
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
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.config.v1alpha1.ConfigService.RegisterDevice is unimplemented"))

    /**
     * Returns the response to an RPC for bff.bridgr.config.v1alpha1.ConfigService.VerifyDevice.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun verifyDevice(request: VerifyDeviceRequest): VerifyDeviceResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.config.v1alpha1.ConfigService.VerifyDevice is unimplemented"))

    /**
     * Returns the response to an RPC for
     * bff.bridgr.config.v1alpha1.ConfigService.AssignDeviceGroup.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun assignDeviceGroup(request: AssignDeviceGroupRequest):
        AssignDeviceGroupResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.config.v1alpha1.ConfigService.AssignDeviceGroup is unimplemented"))

    /**
     * Returns the response to an RPC for
     * bff.bridgr.config.v1alpha1.ConfigService.UpdateForwardingRules.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun updateForwardingRules(request: UpdateForwardingRulesRequest):
        UpdateForwardingRulesResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.config.v1alpha1.ConfigService.UpdateForwardingRules is unimplemented"))

    /**
     * Returns the response to an RPC for
     * bff.bridgr.config.v1alpha1.ConfigService.GetForwardingRules.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getForwardingRules(request: GetForwardingRulesRequest):
        GetForwardingRulesResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.config.v1alpha1.ConfigService.GetForwardingRules is unimplemented"))

    /**
     * Returns a [Flow] of responses to an RPC for
     * bff.bridgr.config.v1alpha1.ConfigService.GetConfig.
     *
     * If creating or collecting the returned flow fails with a [StatusException], the RPC
     * will fail with the corresponding [io.grpc.Status].  If it fails with a
     * [java.util.concurrent.CancellationException], the RPC will fail with status
     * `Status.CANCELLED`.  If creating
     * or collecting the returned flow fails for any other reason, the RPC will fail with
     * `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open fun getConfig(request: GetConfigRequest): Flow<GetConfigResponse> = throw
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.config.v1alpha1.ConfigService.GetConfig is unimplemented"))

    /**
     * Returns the response to an RPC for bff.bridgr.config.v1alpha1.ConfigService.UpdateConfig.
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
        StatusException(UNIMPLEMENTED.withDescription("Method bff.bridgr.config.v1alpha1.ConfigService.UpdateConfig is unimplemented"))

    final override fun bindService(): ServerServiceDefinition = builder(getServiceDescriptor())
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ConfigServiceGrpc.getRegisterDeviceMethod(),
      implementation = ::registerDevice
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ConfigServiceGrpc.getVerifyDeviceMethod(),
      implementation = ::verifyDevice
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ConfigServiceGrpc.getAssignDeviceGroupMethod(),
      implementation = ::assignDeviceGroup
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ConfigServiceGrpc.getUpdateForwardingRulesMethod(),
      implementation = ::updateForwardingRules
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ConfigServiceGrpc.getGetForwardingRulesMethod(),
      implementation = ::getForwardingRules
    ))
      .addMethod(serverStreamingServerMethodDefinition(
      context = this.context,
      descriptor = ConfigServiceGrpc.getGetConfigMethod(),
      implementation = ::getConfig
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ConfigServiceGrpc.getUpdateConfigMethod(),
      implementation = ::updateConfig
    )).build()
  }
}
