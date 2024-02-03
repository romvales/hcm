// @tags-mvp

import type {
  HCMPendingJoinRequestService,
  Organization, PendingJoinRequest
} from './index.d'

import type { 
  StandardShift,
} from './time'

export abstract class HCMRoleService implements BaseOrganizationEntityStatusChecker<Role> {
  createRole(name: string): Role

  async getRoleById(id: number): Promise<Role>
  async deleteRoleById<T>(id: number): Promise<T>
  async saveRole<T>(role: Role): Promise<T>

  changeRoleName<T>(team: Team): T
  changeRoleStatus<T>(role: Role, status: RoleStatus): T
}

export abstract class HCMTeamService implements BaseOrganizationEntityStatusChecker<role> {
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

export abstract class HCMWorkerService implements HCMPendingJoinRequestService<Worker> {
  createWorker(
    name: { firstName: string, middleName: string, lastName: string },
    email: string,
    username: string,
    mobileNumber?: string,
    birthdate?: number
  ): Worker

  async getWorkerById(id: number): Promise<Worker>
  async deleteWorkerById<T>(id: number): Promise<T>
  async saveWorker<T>(worker: Worker): Promise<T>

  changeWorkerStatus<T>(worker: Worker, status: WorkerStatus): T
  changeWorkerType<T>(worker: Worker, type: WorkerType): T
  changeWorkerRole<T>(worker: Worker, newRole: Role): T
  changeWorkerTeam<T>(worker: Worker, newTeam: Team): T
  changeWorkerPayCycle<T>(worker: Worker, newPayCycle: WorkerPayCycle): T

  suspend<T>(worker: Worker): T
  terminate<T>(worker: Worker): T
  resign<T>(worker: Worker): T

  addWorkerAddress<T>(worker: Worker, address: WorkerAddress): T
  addIdentityCards<T>(worker: Worker, identityCards: WorkerIdentityCard[]): T

  getWorkerType(worker: Worker): WorkerType
  getWorkerStatus(worker: Worker): WorkerStatus
  getAddresses(worker: Worker): WorkerAddress[]
  getIdentityCards(worker: Worker): WorkerIdentityCard[]

  async getOrganization(worker: Worker): Promise<Organization | undefined>
  async getRole(worker: Worker): Promise<Role | undefined>
  async getTeam(worker: Worker): Promise<Team | undefined>

  hasOverridenStandardRoleShift(worker: Worker): boolean
  isWorkerHired(worker: Worker): boolean
  isWorkerOnLeave(worker: Worker): boolean
  isWorkerRemote(worker: Worker): boolean
  isWorkerOnline(worker: Worker): boolean
  isWorkerRemotelyOnline(worker: Worker): boolean
  isWorkerOffline(worker: Worker): boolean
  isWorkerAway(worker: Worker): boolean
  isWorkerSuspended(worker: Worker): boolean
  isWorkerOnCall(worker: Worker): boolean
}

export interface BaseOrganizationEntityStatusChecker<Entity = unknown> {
  isActive(entity: Entity): boolean
  isInactive(entity: Entity): boolean
  isOnReview(entity: Entity): boolean
  isTerminated(entity: Entity): boolean
}

// 
export type Role = {
  id: number
  createdById: number
  updatedById?: number
  organizationId?: number

  createdAt: number
  lastUpdatedAt?: number

  createdBy: Worker
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

  createdAt: number
  lastUpdatedAt?: number

  createdBy: Worker
  updatedBy?: Worker

  organization?: Organization
  status?: TeamStatus
  
  name: string
  members: Worker[]
}

// 
export type Worker = {
  id: number
  createdById?: number
  updatedById?: number

  // Represents the collection of identity cards uploaded by the worker
  identityCards?: WorkerIdentityCard[]

  createdBy?: Worker
  updatedBy?: Worker

  createdAt: number
  lastUpdatedAt?: number

  pictureUrl?: string

  firstName: string
  middleName?: string
  lastName: string
  nickname?: string

  gender: WorkerGender

  birthdate?: number

  username?: string
  email: string
  mobile? : string

  addresses?: WorkerAddress[]
}

//
export type WorkerOrganizationInfo = {
  id: number
  organizationId: number
  workerId: number
  hiredById?: number
  scheduledSuspensionAt?: number

  status?: WorkerStatus
  type?: WorkerType

  // When was the employee hired in the organization?
  hiredAt?: number
  
  // A timestamp that represents when was the worker suspended
  suspendedAt?: number

  // Indicates when was the last time the worker leave in the work.
  leaveAt?: number

  // Termination date
  terminatedAt?: number

  // When did the worker returned after leaving the work?
  returnedAt?: number

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
  payCycle?: WorkerPayCycle
}

//
export type WorkerIdentityCard = {
  id: number
  workerId: number
  createdById: number
  updatedById: number

  worker: Worker
  createdBy: Worker
  updatedBy?: Worker
  
  createdAt: number
  lastUpdatedAt?: number

  frontImageUrl: string
  backImageUrl: string

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
export enum WorkerStatus {
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
export enum WorkerType {
  PART,
  FULL,
  SEASONAL,
  TEMPORARY,
  LEASED,
}

//
export enum WorkerPayCycle {
  WEEKLY,
  BIWEEKLY,
  SEMIMONTHLY,
  MONTHLY,
}

export enum TeamStatus {
  ACTIVE,
  INACTIVE,
  REVIEW,
  TERMINATED,
}

export enum RoleStatus {
  ACTIVE,
  INACTIVE,
  REVIEW,
  TERMINATED,
}