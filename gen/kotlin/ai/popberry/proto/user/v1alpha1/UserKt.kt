// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: user/v1alpha1/user.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package ai.popberry.proto.user.v1alpha1;

@kotlin.jvm.JvmName("-initializeuser")
public inline fun user(block: ai.popberry.proto.user.v1alpha1.UserKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.user.v1alpha1.User =
  ai.popberry.proto.user.v1alpha1.UserKt.Dsl._create(ai.popberry.proto.user.v1alpha1.User.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `user.v1alpha1.User`
 */
public object UserKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: ai.popberry.proto.user.v1alpha1.User.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
      internal fun _create(builder: ai.popberry.proto.user.v1alpha1.User.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
  @kotlin.PublishedApi
    internal fun _build(): ai.popberry.proto.user.v1alpha1.User = _builder.build()

    /**
     * `int32 id = 1 [json_name = "id"];`
     */
    public var id: kotlin.Int
      @kotlin.jvm.JvmName("getId")
        get() = _builder.id
      @kotlin.jvm.JvmName("setId")
        set(value) {
        _builder.id = value
      }
    /**
     * <code>int32 id = 1 [json_name = "id"];</code>
     * @return This builder for chaining.
     */
    public fun clearId() {
      _builder.clearId()
    }

    /**
     * `string username = 2 [json_name = "username"];`
     */
    public var username: kotlin.String
      @kotlin.jvm.JvmName("getUsername")
        get() = _builder.username
      @kotlin.jvm.JvmName("setUsername")
        set(value) {
        _builder.username = value
      }
    /**
     * <code>string username = 2 [json_name = "username"];</code>
     * @return This builder for chaining.
     */
    public fun clearUsername() {
      _builder.clearUsername()
    }

    /**
     * `string firstname = 3 [json_name = "firstname"];`
     */
    public var firstname: kotlin.String
      @kotlin.jvm.JvmName("getFirstname")
        get() = _builder.firstname
      @kotlin.jvm.JvmName("setFirstname")
        set(value) {
        _builder.firstname = value
      }
    /**
     * <code>string firstname = 3 [json_name = "firstname"];</code>
     * @return This builder for chaining.
     */
    public fun clearFirstname() {
      _builder.clearFirstname()
    }

    /**
     * `string lastname = 4 [json_name = "lastname"];`
     */
    public var lastname: kotlin.String
      @kotlin.jvm.JvmName("getLastname")
        get() = _builder.lastname
      @kotlin.jvm.JvmName("setLastname")
        set(value) {
        _builder.lastname = value
      }
    /**
     * <code>string lastname = 4 [json_name = "lastname"];</code>
     * @return This builder for chaining.
     */
    public fun clearLastname() {
      _builder.clearLastname()
    }

    /**
     * `.user.v1alpha1.Gender gender = 5 [json_name = "gender"];`
     */
    public var gender: ai.popberry.proto.user.v1alpha1.Gender
      @kotlin.jvm.JvmName("getGender")
        get() = _builder.gender
      @kotlin.jvm.JvmName("setGender")
        set(value) {
        _builder.gender = value
      }
    public var genderValue: kotlin.Int
      @kotlin.jvm.JvmName("getGenderValue")
        get() = _builder.genderValue
      @kotlin.jvm.JvmName("setGenderValue")
        set(value) {
        _builder.genderValue = value
      }
    /**
     * <code>.user.v1alpha1.Gender gender = 5 [json_name = "gender"];</code>
     * @return This builder for chaining.
     */
    public fun clearGender() {
      _builder.clearGender()
    }

    /**
     * `string lineid = 6 [json_name = "lineid"];`
     */
    public var lineid: kotlin.String
      @kotlin.jvm.JvmName("getLineid")
        get() = _builder.lineid
      @kotlin.jvm.JvmName("setLineid")
        set(value) {
        _builder.lineid = value
      }
    /**
     * <code>string lineid = 6 [json_name = "lineid"];</code>
     * @return This builder for chaining.
     */
    public fun clearLineid() {
      _builder.clearLineid()
    }

    /**
     * `string banknumber = 7 [json_name = "banknumber"];`
     */
    public var banknumber: kotlin.String
      @kotlin.jvm.JvmName("getBanknumber")
        get() = _builder.banknumber
      @kotlin.jvm.JvmName("setBanknumber")
        set(value) {
        _builder.banknumber = value
      }
    /**
     * <code>string banknumber = 7 [json_name = "banknumber"];</code>
     * @return This builder for chaining.
     */
    public fun clearBanknumber() {
      _builder.clearBanknumber()
    }

    /**
     * `.google.protobuf.Timestamp created_at = 8 [json_name = "createdAt"];`
     */
    public var createdAt: com.google.protobuf.Timestamp
      @kotlin.jvm.JvmName("getCreatedAt")
        get() = _builder.createdAt
      @kotlin.jvm.JvmName("setCreatedAt")
        set(value) {
        _builder.createdAt = value
      }
    /**
     * <code>.google.protobuf.Timestamp created_at = 8 [json_name = "createdAt"];</code>
     * @return This builder for chaining.
     */
    public fun clearCreatedAt() {
      _builder.clearCreatedAt()
    }
    /**
     * <code>.google.protobuf.Timestamp created_at = 8 [json_name = "createdAt"];</code>
     * @return Whether the createdAt field is set.
     * @return This builder for chaining.
     */
    public fun hasCreatedAt(): kotlin.Boolean {
      return _builder.hasCreatedAt()
    }

    public val UserKt.Dsl.createdAtOrNull: com.google.protobuf.Timestamp?
      get() = _builder.createdAtOrNull

    /**
     * `.google.protobuf.Timestamp updated_at = 9 [json_name = "updatedAt"];`
     */
    public var updatedAt: com.google.protobuf.Timestamp
      @kotlin.jvm.JvmName("getUpdatedAt")
        get() = _builder.updatedAt
      @kotlin.jvm.JvmName("setUpdatedAt")
        set(value) {
        _builder.updatedAt = value
      }
    /**
     * <code>.google.protobuf.Timestamp updated_at = 9 [json_name = "updatedAt"];</code>
     * @return This builder for chaining.
     */
    public fun clearUpdatedAt() {
      _builder.clearUpdatedAt()
    }
    /**
     * <code>.google.protobuf.Timestamp updated_at = 9 [json_name = "updatedAt"];</code>
     * @return Whether the updatedAt field is set.
     * @return This builder for chaining.
     */
    public fun hasUpdatedAt(): kotlin.Boolean {
      return _builder.hasUpdatedAt()
    }

    public val UserKt.Dsl.updatedAtOrNull: com.google.protobuf.Timestamp?
      get() = _builder.updatedAtOrNull

    /**
     * `.user.v1alpha1.UserStatus status = 10 [json_name = "status"];`
     */
    public var status: ai.popberry.proto.user.v1alpha1.UserStatus
      @kotlin.jvm.JvmName("getStatus")
        get() = _builder.status
      @kotlin.jvm.JvmName("setStatus")
        set(value) {
        _builder.status = value
      }
    public var statusValue: kotlin.Int
      @kotlin.jvm.JvmName("getStatusValue")
        get() = _builder.statusValue
      @kotlin.jvm.JvmName("setStatusValue")
        set(value) {
        _builder.statusValue = value
      }
    /**
     * <code>.user.v1alpha1.UserStatus status = 10 [json_name = "status"];</code>
     * @return This builder for chaining.
     */
    public fun clearStatus() {
      _builder.clearStatus()
    }

    /**
     * `int32 bank_categories_id = 11 [json_name = "bankCategoriesId"];`
     */
    public var bankCategoriesId: kotlin.Int
      @kotlin.jvm.JvmName("getBankCategoriesId")
        get() = _builder.bankCategoriesId
      @kotlin.jvm.JvmName("setBankCategoriesId")
        set(value) {
        _builder.bankCategoriesId = value
      }
    /**
     * <code>int32 bank_categories_id = 11 [json_name = "bankCategoriesId"];</code>
     * @return This builder for chaining.
     */
    public fun clearBankCategoriesId() {
      _builder.clearBankCategoriesId()
    }

    /**
     * `int32 ways_id = 12 [json_name = "waysId"];`
     */
    public var waysId: kotlin.Int
      @kotlin.jvm.JvmName("getWaysId")
        get() = _builder.waysId
      @kotlin.jvm.JvmName("setWaysId")
        set(value) {
        _builder.waysId = value
      }
    /**
     * <code>int32 ways_id = 12 [json_name = "waysId"];</code>
     * @return This builder for chaining.
     */
    public fun clearWaysId() {
      _builder.clearWaysId()
    }

    /**
     * `int32 wallets_id = 13 [json_name = "walletsId"];`
     */
    public var walletsId: kotlin.Int
      @kotlin.jvm.JvmName("getWalletsId")
        get() = _builder.walletsId
      @kotlin.jvm.JvmName("setWalletsId")
        set(value) {
        _builder.walletsId = value
      }
    /**
     * <code>int32 wallets_id = 13 [json_name = "walletsId"];</code>
     * @return This builder for chaining.
     */
    public fun clearWalletsId() {
      _builder.clearWalletsId()
    }

    /**
     * `.user.v1alpha1.UserRole role = 14 [json_name = "role"];`
     */
    public var role: ai.popberry.proto.user.v1alpha1.UserRole
      @kotlin.jvm.JvmName("getRole")
        get() = _builder.role
      @kotlin.jvm.JvmName("setRole")
        set(value) {
        _builder.role = value
      }
    public var roleValue: kotlin.Int
      @kotlin.jvm.JvmName("getRoleValue")
        get() = _builder.roleValue
      @kotlin.jvm.JvmName("setRoleValue")
        set(value) {
        _builder.roleValue = value
      }
    /**
     * <code>.user.v1alpha1.UserRole role = 14 [json_name = "role"];</code>
     * @return This builder for chaining.
     */
    public fun clearRole() {
      _builder.clearRole()
    }

    /**
     * `string line_user_id = 15 [json_name = "lineUserId"];`
     */
    public var lineUserId: kotlin.String
      @kotlin.jvm.JvmName("getLineUserId")
        get() = _builder.lineUserId
      @kotlin.jvm.JvmName("setLineUserId")
        set(value) {
        _builder.lineUserId = value
      }
    /**
     * <code>string line_user_id = 15 [json_name = "lineUserId"];</code>
     * @return This builder for chaining.
     */
    public fun clearLineUserId() {
      _builder.clearLineUserId()
    }

    /**
     * `string affiliate_id = 16 [json_name = "affiliateId"];`
     */
    public var affiliateId: kotlin.String
      @kotlin.jvm.JvmName("getAffiliateId")
        get() = _builder.affiliateId
      @kotlin.jvm.JvmName("setAffiliateId")
        set(value) {
        _builder.affiliateId = value
      }
    /**
     * <code>string affiliate_id = 16 [json_name = "affiliateId"];</code>
     * @return This builder for chaining.
     */
    public fun clearAffiliateId() {
      _builder.clearAffiliateId()
    }

    /**
     * `string rank = 17 [json_name = "rank"];`
     */
    public var rank: kotlin.String
      @kotlin.jvm.JvmName("getRank")
        get() = _builder.rank
      @kotlin.jvm.JvmName("setRank")
        set(value) {
        _builder.rank = value
      }
    /**
     * <code>string rank = 17 [json_name = "rank"];</code>
     * @return This builder for chaining.
     */
    public fun clearRank() {
      _builder.clearRank()
    }

    /**
     * `double point = 18 [json_name = "point"];`
     */
    public var point: kotlin.Double
      @kotlin.jvm.JvmName("getPoint")
        get() = _builder.point
      @kotlin.jvm.JvmName("setPoint")
        set(value) {
        _builder.point = value
      }
    /**
     * <code>double point = 18 [json_name = "point"];</code>
     * @return This builder for chaining.
     */
    public fun clearPoint() {
      _builder.clearPoint()
    }

    /**
     * `double exp = 19 [json_name = "exp"];`
     */
    public var exp: kotlin.Double
      @kotlin.jvm.JvmName("getExp")
        get() = _builder.exp
      @kotlin.jvm.JvmName("setExp")
        set(value) {
        _builder.exp = value
      }
    /**
     * <code>double exp = 19 [json_name = "exp"];</code>
     * @return This builder for chaining.
     */
    public fun clearExp() {
      _builder.clearExp()
    }

    /**
     * `string dob = 20 [json_name = "dob"];`
     */
    public var dob: kotlin.String
      @kotlin.jvm.JvmName("getDob")
        get() = _builder.dob
      @kotlin.jvm.JvmName("setDob")
        set(value) {
        _builder.dob = value
      }
    /**
     * <code>string dob = 20 [json_name = "dob"];</code>
     * @return This builder for chaining.
     */
    public fun clearDob() {
      _builder.clearDob()
    }

    /**
     * ```
     * Added Wallet reference
     * ```
     *
     * `.user.v1alpha1.Wallet wallet = 21 [json_name = "wallet"];`
     */
    public var wallet: ai.popberry.proto.user.v1alpha1.Wallet
      @kotlin.jvm.JvmName("getWallet")
        get() = _builder.wallet
      @kotlin.jvm.JvmName("setWallet")
        set(value) {
        _builder.wallet = value
      }
    /**
     * <pre>
     * Added Wallet reference
     * </pre>
     *
     * <code>.user.v1alpha1.Wallet wallet = 21 [json_name = "wallet"];</code>
     * @return This builder for chaining.
     */
    public fun clearWallet() {
      _builder.clearWallet()
    }
    /**
     * <pre>
     * Added Wallet reference
     * </pre>
     *
     * <code>.user.v1alpha1.Wallet wallet = 21 [json_name = "wallet"];</code>
     * @return Whether the wallet field is set.
     * @return This builder for chaining.
     */
    public fun hasWallet(): kotlin.Boolean {
      return _builder.hasWallet()
    }

    public val UserKt.Dsl.walletOrNull: ai.popberry.proto.user.v1alpha1.Wallet?
      get() = _builder.walletOrNull
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun ai.popberry.proto.user.v1alpha1.User.copy(block: `ai.popberry.proto.user.v1alpha1`.UserKt.Dsl.() -> kotlin.Unit): ai.popberry.proto.user.v1alpha1.User =
  `ai.popberry.proto.user.v1alpha1`.UserKt.Dsl._create(this.toBuilder()).apply { block() }._build()

public val ai.popberry.proto.user.v1alpha1.UserOrBuilder.createdAtOrNull: com.google.protobuf.Timestamp?
  get() = if (hasCreatedAt()) getCreatedAt() else null

public val ai.popberry.proto.user.v1alpha1.UserOrBuilder.updatedAtOrNull: com.google.protobuf.Timestamp?
  get() = if (hasUpdatedAt()) getUpdatedAt() else null

public val ai.popberry.proto.user.v1alpha1.UserOrBuilder.walletOrNull: ai.popberry.proto.user.v1alpha1.Wallet?
  get() = if (hasWallet()) getWallet() else null

