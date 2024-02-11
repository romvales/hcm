// @tags-mvp

import { PostgrestResponse, PostgrestSingleResponse } from '@supabase/supabase-js'
import type {
  HCMPendingJoinRequestService,
  Organization, PendingJoinRequest
} from './index.d'

import type { 
  StandardShift,
} from './time'
import { TargetUndefinedError, isTargetNotDefined } from './implementations/supabase'

export abstract class HCMRoleService implements BaseOrganizationEntityStatusChecker {
  createRole(name: string): Role
  
  setTarget(role: Role) {
    this.role = role
    return this
  }

  async getRoleById(id: number)
  async deleteRoleById(id: number)
  async saveRole()

  changeRoleName(newName: string) {
    isTargetNotDefined(this.target)

    const target = this.target as Role
    target.name = newName

    return this
  }

  changeRoleStatus(status: RoleStatus) {
    isTargetNotDefined(this.target)

    const target = this.target as Role
    target.status = status

    return this
  }

  private isRoleStatusCheckAgainst(status: RoleStatus) {
    isTargetNotDefined(this.target)

    const target = this.target as Role

    return target.status == status
  }

  isActive = () => this.isRoleStatusCheckAgainst(RoleStatus.ACTIVE)
  isInactive = () => this.isRoleStatusCheckAgainst(RoleStatus.INACTIVE)
  isOnReview = () => this.isRoleStatusCheckAgainst(RoleStatus.REVIEW)
  isTerminated = () => this.isRoleStatusCheckAgainst(RoleStatus.TERMINATED)
}

export abstract class HCMTeamService {
  createTeam(name: string): Team

  async getTeamById(id: number): Promise<Team>
  async deleteTeamById<T>(id: number): Promise<T>
  async saveTeam<T>(team: Team): Promise<T>

  changeTeamName<T>(team: Team): T
  changeTeamStatus<T>(team: Team, status: TeamStatus): T

  async getWorkerMembers(team: Team): Promise<Worker[]>

  async addWorkerToTeam<T>(team: Team, worker: Worker): Promise<T>
  async removeWorkerFromTeam<T>(team: Team, worker: Worker): Promise<T>
}

export abstract class HCMWorkerService {
  createWorker(email: string, username: string): Worker

  setTarget(worker: Worker): this {
    this.target = worker
    return this
  }

  async getWorkerById(id?: number)  
  async deleteWorkerById(id: number)

  getAddresses(): WorkerAddress[]
  async getIdentityCards(): Promise<WorkerIdentityCard[]>
  async deleteIdentityCardById(id: number)

  //
  changeName(name: {
    firstName?: string,
    lastName?: string,
    middleName?: string,
    suffix?: string,
  }): this

  changePictureUrl(pictureUrl: string): this {
    isTargetNotDefined(this.target)
    
    const target = this.target as Worker
    target.pictureUrl = pictureUrl

    return this
  }

  changeBirthdate(birthdate: number): this {
    isTargetNotDefined(this.target)
    const target = this.target as Worker
    target.birthdate = birthdate

    return this
  }

  changeEmailAddress(newEmail: string): this {
    isTargetNotDefined(this.target)
    const target = this.target as Worker
    target.email = newEmail
    return this
  }

  changeMobileNumber(newMobile: string): this {
    isTargetNotDefined(this.target)
    
    const target = this.target as Worker
    target.mobile = newMobile

    return this
  }

  changeUsername(newUsername: string): this {
    isTargetNotDefined(this.target)

    const target = this.target as Worker
    target.username = newUsername

    return this
  }

  changeGender(gender: WorkerGender): this {
    isTargetNotDefined(this.target)

    const target = this.target as Worker
    target.gender = gender

    return this
  }

  changeName(name: { 
    firstName?: string, 
    lastName?: string,
    middleName?: string,
    suffix?: string 
  }): this {
    isTargetNotDefined(this.target)
    
    const target = this.target as Worker
    const { firstName, lastName, middleName, suffix } = name
    
    if (firstName) target.firstName = firstName
    if (lastName) target.lastName = lastName
    if (middleName) target.middleName = middleName
    if (suffix) target.suffix = suffix

    return this
  }

  changeWorkerIndicator(indicator: WorkerIndicator): this {
    isTargetNotDefined(this.target)

    const target = this.target as Worker
    target.indicator = indicator

    return this
  }
  
  async saveWorker()
  async saveWorkerIdentityCard(id: WorkerIdentityCard)

  addWorkerAddress(address: WorkerAddress): this
  addIdentityCards(identityCards: WorkerIdentityCard[]): this
}

export abstract class HCMWorkerOrganizationService {

  setTarget(worker: Worker): this {
    this.target = worker
    return this
  }

  async suspend(org: Organization)
  async terminate(org: Organization)
  async resign(org: Organization)

  async getOrganizations(): Promise<Organization[] | undefined>
  async getRoles(): Promise<Role[] | undefined>
  async getTeams(): Promise<Team[] | undefined>

  changeWorkerStatus(status: WorkerStatus) {

    return this
  }

  changeWorkerType(type: WorkerType) {

    return this
  }

  changeWorkerRole(newRole: Role) {

    return this
  }

  changeWorkerTeam(newTeam: Team) {

    return this
  }

  changeWorkerPayCycle(newPayCycle: WorkerPayCycle) {

    return this
  }

  getWorkerType(): WorkerType
  getWorkerStatus(): WorkerStatus

  hasOverridenStandardRoleShift(): boolean
  isWorkerHired(): boolean
  isWorkerOnLeave(): boolean
  isWorkerRemote(): boolean
  isWorkerOnline(): boolean
  isWorkerRemotelyOnline(): boolean
  isWorkerOffline(): boolean
  isWorkerAway(): boolean
  isWorkerSuspended(): boolean
  isWorkerOnCall(): boolean

}

export interface BaseOrganizationEntityStatusChecker {
  isActive(): boolean
  isInactive(): boolean
  isOnReview(): boolean
  isTerminated(): boolean
}

// 
export type Role = {
  id?: number
  createdById?: number
  updatedById?: number
  organizationId?: number

  createdAt?: string
  lastUpdatedAt?: string

  createdBy?: Worker
  updatedBy?: Worker

  organization?: Organization
  status?: RoleStatus

  // Every role should have 7 standard shifts assigned
  shifts?: StandardShift[]

  name: string
}

// 
export type Team = {
  id: number
  createdById: number
  updatedById?: number
  organizationId?: number

  createdAt: string
  lastUpdatedAt?: string

  createdBy: Worker
  updatedBy?: Worker

  organization?: Organization
  status?: TeamStatus
  
  name: string
  members: Worker[]
}

// 
export type Worker = {
  id?: number
  userId?: string
  createdById?: number
  updatedById?: number

  // Indicates whether the worker is online, offline, or away. This is visible to all organizations
  // and workers in the platform.
  indicator?: WorkerIndicator

  // Represents the collection of identity cards uploaded by the worker
  identityCards?: WorkerIdentityCard[]

  createdBy?: Worker
  updatedBy?: Worker

  createdAt?: string
  lastUpdatedAt?: string

  pictureUrl?: string

  firstName?: string
  middleName?: string
  lastName?: string
  nickname?: string
  suffix?: string

  gender?: WorkerGender

  birthdate?: string

  username?: string
  email: string
  mobile? : string

  addresses?: WorkerAddress[]
}

// Indicates what information a worker has within an organization.
export type WorkerOrganizationInfo = {
  id: number
  workerId: number
  organizationId: number
  hiredById?: number

  //
  status?: WorkerOrganizationStatus
  type?: WorkerOrganizationType

  //
  scheduledSuspensionAt?: string

  // When was the employee hired in the organization?
  hiredAt?: string
  
  // A timestamp that represents when was the worker suspended
  suspendedAt?: string

  // Indicates when was the last time the worker leave in the work.
  leaveAt?: string

  // Termination date
  terminatedAt?: string

  // When did the worker returned after leaving the work?
  returnedAt?: string

  // Which team does this worker belongs? Does it belong to a team?
  team?: Team
  
  // What assigned role does this worker have? @IMPORTANT: if empty, do not include to the payroll.
  role?: Role

  // When an worker is onboarded to the team as remote, make sure that
  // isRemote is toggled to true
  isRemote?: boolean

  // When a worker is now hired by an organization, make sure that you toggle
  // this state to true
  isHired?: boolean

  //
  isDayOff?: boolean

  //
  isOnCall?: boolean

  //
  isOnLeave?: boolean

  //
  isTerminated?: boolean

  //
  isSuspended?: boolean

  // Overrides the standard shift assigned to the role for a specific worker.
  overridesShift?: StandardShift[]

  //
  organization?: Organization

  // Used to indicate what pay cycle is a worker in the payroll
  payCycle?: WorkerOrganizationPayCycle
}

//
export type WorkerIdentityCard = {
  id?: number
  workerId?: number
  createdById?: number
  updatedById?: number

  worker?: Worker
  createdBy?: Worker
  updatedBy?: Worker
  
  createdAt?: string
  lastUpdatedAt?: string

  frontImageUrl?: string
  backImageUrl?: string

  name: string
  extractedInfo?: unknown
}

//
export type WorkerAddress = {
  addrType: WorkerAddressType
  streetLines?: string[]
  city: string
  state: string
  postalCode?: string
  country?: string
}

//
export enum WorkerAddressType {
  HOME,
  BUSINESS,
  BILLING,
  SHIPPING
}

//
export enum WorkerGender {
  MALE,
  FEMALE,
  OTHER
}

//
export enum WorkerIndicator {
  OFFLINE,
  ONLINE,
  AWAY,
}

//
export enum WorkerOrganizationStatus {
  OFFLINE,
  ONLINE,
  RONLINE,
  AWAY,

  SUSPENDED,
  RESIGNED,
  TERMINATED,
  LEAVE,

  ONCALL,
}

//
export enum WorkerOrganizationType {
  PART,
  FULL,
  SEASONAL,
  TEMPORARY,
  LEASED,
}

//
export enum WorkerOrganizationPayCycle {
  WEEKLY,
  BIWEEKLY,
  SEMIMONTHLY,
  MONTHLY,
}

//
export enum TeamStatus {
  ACTIVE,
  INACTIVE,
  REVIEW,
  TERMINATED,
}

//
export enum RoleStatus {
  ACTIVE,
  INACTIVE,
  REVIEW,
  TERMINATED,
}