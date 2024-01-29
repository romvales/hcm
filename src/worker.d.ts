// @tags-mvp

import type {
  Organization
} from './index.d'

import type { 
  StandardShift,
} from './time'

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
  organizationId?: number
  roleId?: number
  teamId?: number

  // Represents the collection of identity cards uploaded by the worker
  identityCards?: WorkerIdentityCard[]

  createdBy: Worker
  updatedBy?: Worker

  createdAt: number
  lastUpdatedAt?: number

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

  // 
  isRemote?: boolean

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

  extractedInfo?: any
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
  WAT_HOME,
  WAT_BUSINESS,
  WAT_BILLING,
  WAT_SHIPPING
}

//
export enum WorkerGender {
  WG_MALE,
  WG_FEMALE,
  WG_OTHER
}

//
export enum WorkerStatus {
  WS_ONLINE,
  WS_OFFLINE,
  WS_AWAY,

  WS_SUSPENDED,
  WS_TERMINATED,
  WS_LEAVE,

  WT_ONCALL,
}

//
export enum WorkerType {
  WT_PART,
  WT_FULL,
  WT_SEASONAL,
  WT_TEMPORARY,
  WT_LEASED,
}

//
export enum WorkerPayCycle {
  WPC_WEEKLY,
  WPC_BIWEEKLY,
  WPC_SEMIMONTHLY,
  WPC_MONTHLY,
}

export enum TeamStatus {
  TS_ACTIVE,
  TS_INACTIVE,
  TS_REVIEW,
  TS_TERMINATED,
}

export enum RoleStatus {
  TS_ACTIVE,
  TS_INACTIVE,
  TS_REVIEW,
  TS_TERMINATED,
}