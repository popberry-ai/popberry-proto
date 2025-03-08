package ai.popberry.proto.user.v1alpha1

import ai.popberry.proto.user.v1alpha1.UserServiceGrpc.getServiceDescriptor
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
 * Holder for Kotlin coroutine-based client and server APIs for user.v1alpha1.UserService.
 */
public object UserServiceGrpcKt {
  public const val SERVICE_NAME: String = UserServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  public val createUserMethod: MethodDescriptor<CreateUserRequest, CreateUserResponse>
    @JvmStatic
    get() = UserServiceGrpc.getCreateUserMethod()

  public val getUserMethod: MethodDescriptor<GetUserRequest, GetUserResponse>
    @JvmStatic
    get() = UserServiceGrpc.getGetUserMethod()

  public val updateUserMethod: MethodDescriptor<UpdateUserRequest, UpdateUserResponse>
    @JvmStatic
    get() = UserServiceGrpc.getUpdateUserMethod()

  public val deleteUserMethod: MethodDescriptor<DeleteUserRequest, DeleteUserResponse>
    @JvmStatic
    get() = UserServiceGrpc.getDeleteUserMethod()

  public val listUsersMethod: MethodDescriptor<ListUsersRequest, ListUsersResponse>
    @JvmStatic
    get() = UserServiceGrpc.getListUsersMethod()

  /**
   * A stub for issuing RPCs to a(n) user.v1alpha1.UserService service as suspending coroutines.
   */
  @StubFor(UserServiceGrpc::class)
  public class UserServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<UserServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions): UserServiceCoroutineStub =
        UserServiceCoroutineStub(channel, callOptions)

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
    public suspend fun createUser(request: CreateUserRequest, headers: Metadata = Metadata()):
        CreateUserResponse = unaryRpc(
      channel,
      UserServiceGrpc.getCreateUserMethod(),
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
    public suspend fun getUser(request: GetUserRequest, headers: Metadata = Metadata()):
        GetUserResponse = unaryRpc(
      channel,
      UserServiceGrpc.getGetUserMethod(),
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
    public suspend fun updateUser(request: UpdateUserRequest, headers: Metadata = Metadata()):
        UpdateUserResponse = unaryRpc(
      channel,
      UserServiceGrpc.getUpdateUserMethod(),
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
    public suspend fun deleteUser(request: DeleteUserRequest, headers: Metadata = Metadata()):
        DeleteUserResponse = unaryRpc(
      channel,
      UserServiceGrpc.getDeleteUserMethod(),
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
    public suspend fun listUsers(request: ListUsersRequest, headers: Metadata = Metadata()):
        ListUsersResponse = unaryRpc(
      channel,
      UserServiceGrpc.getListUsersMethod(),
      request,
      callOptions,
      headers
    )
  }

  /**
   * Skeletal implementation of the user.v1alpha1.UserService service based on Kotlin coroutines.
   */
  public abstract class UserServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    /**
     * Returns the response to an RPC for user.v1alpha1.UserService.CreateUser.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun createUser(request: CreateUserRequest): CreateUserResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method user.v1alpha1.UserService.CreateUser is unimplemented"))

    /**
     * Returns the response to an RPC for user.v1alpha1.UserService.GetUser.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getUser(request: GetUserRequest): GetUserResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method user.v1alpha1.UserService.GetUser is unimplemented"))

    /**
     * Returns the response to an RPC for user.v1alpha1.UserService.UpdateUser.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun updateUser(request: UpdateUserRequest): UpdateUserResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method user.v1alpha1.UserService.UpdateUser is unimplemented"))

    /**
     * Returns the response to an RPC for user.v1alpha1.UserService.DeleteUser.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun deleteUser(request: DeleteUserRequest): DeleteUserResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method user.v1alpha1.UserService.DeleteUser is unimplemented"))

    /**
     * Returns the response to an RPC for user.v1alpha1.UserService.ListUsers.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun listUsers(request: ListUsersRequest): ListUsersResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method user.v1alpha1.UserService.ListUsers is unimplemented"))

    final override fun bindService(): ServerServiceDefinition = builder(getServiceDescriptor())
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = UserServiceGrpc.getCreateUserMethod(),
      implementation = ::createUser
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = UserServiceGrpc.getGetUserMethod(),
      implementation = ::getUser
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = UserServiceGrpc.getUpdateUserMethod(),
      implementation = ::updateUser
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = UserServiceGrpc.getDeleteUserMethod(),
      implementation = ::deleteUser
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = UserServiceGrpc.getListUsersMethod(),
      implementation = ::listUsers
    )).build()
  }
}
