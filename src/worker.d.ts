// @tags-mvp

import type {
  Organization, PendingJoinRequest
} from './index.d'

import type { 
  StandardShift,
} from './time'

export abstract class HCMRoleService {
  static createRole(name: string): Role

  static async getRoleById(id: number): Promise<Role>
  static async deleteRoleById<T>(id: number): Promise<T>
  static async saveRole<T>(role: Role): Promise<T>

  static changeRoleName<T>(team: Team): T
  static changeRoleStatus<T>(role: Role, status: RoleStatus): T
}

export abstract class HCMTeamService {
  static createTeam(name: string): Team

  static async getTeamById(id: number): Promise<Team>
  static async deleteTeamById<T>(id: number): Promise<T>
  static async saveTeam<T>(team: Team): Promise<T>

  static changeTeamName<T>(team: Team): T
  static changeTeamStatus<T>(team: Team, status: TeamStatus): T

  static async getWorkerMembers(team: Team): Promise<Worker[]>

  static async addWorkerToTeam<T>(team: Team, worker: Worker): Promise<T>
  static async removeWorkerFromTeam<T>(team: Team, worker: Worker): Promise<T>
}

export abstract class HCMWorkerService {
  static createWorker(): Worker

  static async getWorkerById(id: number): Promise<Worker>
  static async deleteWorkerById<T>(id: number): Promise<T>
  static async saveWorker<T>(worker: Worker): Promise<T>

  static changeWorkerStatus<T>(worker: Worker, status: WorkerStatus): T
  static changeWorkerType<T>(worker: Worker, type: WorkerType): T
  static changeWorkerRole<T>(worker: Worker, newRole: Role): T
  static changeWorkerTeam<T>(worker: Worker, newTeam: Team): T
  static changeWorkerPayCycle<T>(worker: Worker, newPayCycle: WorkerPayCycle): T

  static suspend<T>(worker: Worker): T
  static terminate<T>(worker: Worker): T
  static resign<T>(worker: Worker): T

  static addWorkerAddress<T>(worker: Worker, address: WorkerAddress): T
  static addIdentityCards<T>(worker: Worker, identityCards: WorkerIdentityCard[]): T

  static getWorkerType(worker: Worker): WorkerType
  static getWorkerStatus(worker: Worker): WorkerStatus
  static getAddresses(worker: Worker): WorkerAddress[]
  static getIdentityCards(worker: Worker): WorkerIdentityCard[]

  static async getOrganization(worker: Worker): Promise<Organization | undefined>
  static async getRole(worker: Worker): Promise<Role | undefined>
  static async getTeam(worker: Worker): Promise<Team | undefined>

  static hasOverridenStandardShifts(worker: Worker): boolean
  static isWorkerHired(worker: Worker): boolean
  static isWorkerOnLeave(worker: Worker): boolean
  static isWorkerRemote(worker: Worker): boolean
  static isWorkerOnline(worker: Worker): boolean
  static isWorkerRemotelyOnline(worker: Worker): boolean
  static isWorkerOffline(worker: Worker): boolean
  static isWorkerAway(worker: Worker): boolean
  static isWorkerSuspended(worker: Worker): boolean
  static isWorkerOnCall(worker: Worker): boolean
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
  createdById: number
  updatedById?: number
  hiredBy?: number
  organizationId?: number
  roleId?: number
  teamId?: number

  // Represents the collection of identity cards uploaded by the worker
  identityCards?: WorkerIdentityCard[]

  createdBy: Worker
  updatedBy?: Worker
  hiredBy?: Worker

  createdAt: number
  lastUpdatedAt?: number

  scheduledSuspensionAt?: number

  // When was the employee hired in the organization?
  hiredAt: number
  
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

  //
  organization?: Organization

  // Used to indicate what pay cycle is a worker in the payroll
  payCycle?: WorkerPayCycle

  status?: WorkerStatus
  type?: WorkerType

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

  pictureUrl?: string

  firstName: string
  middleName?: string
  lastName: string
  nickname?: string
  suffix?: string

  gender: WorkerGender

  birthdate?: number

  username?: string
  email: string
  mobile? : string

  addresses?: WorkerAddress[]
}

//
export type WorkerIdentityCard = {
  id: number

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