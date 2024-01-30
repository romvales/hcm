// @tags-mvp

import { Organization } from '.'
import type {
  Worker
} from './worker'

//
export abstract class HCMWorkerPayInfoService {

}

//
export abstract class HCMCompensationService {

}

//
export abstract class HCMAdditionService {

}

//
export abstract class HCMDeductionService {

}

// Represents the default worker payment info that will be used in the computation
// it can be overriden by anyone by specifying a WorkerPayInfoOverride.
export type WorkerPayInfo = {
  id: number
  createdById: number
  updatedById: number
  workerId: number

  createdAt: number
  lastUpdatedAt?: number

  createdBy: Worker
  updatedBy?: Worker

  type?: WorkerPayInfoType

  worker: Worker

  // A collection of all benefits that the worker payment info has received.
  additions?: Addition[]

  // Same with the above collection, this contains the deductions the worker received.
  deductions?: Deduction[]
}

//
export type WorkerPayInfoOverride = {
  id: number
  payId: number
  pay: WorkerPayInfo
}

//
export type Compensation = {
  id: number
  createdById: number
  updatedById: number
  organizationId: number
  workerId: number

  createdAt: number
  lastUpdatedAt?: number

  createdBy: Worker
  updatedBy?: Worker

  organization: Organization
  status?: CompensationStatus

  worker: Worker

  additions?: Addition[]
  deductions?: Deduction[]

  // A value which holds the gross pay of the worker
  gvalue: number

  // Represents a computed value that is added to the final take-home pay of the worker
  bvalue?: number

  // Computed value that denotes the total deductions in the take-home pay.
  dvalue?: number

  //
  value: number
  
}

// Additions is a value added to the final pay for a worker on any specific pay period.
export type Addition = {
  id: number
  createdById: number
  updatedById: number
  organizationId: number
  workerId?: number

  createdAt: number
  lastUpdatedAt?: number

  createdBy: Worker
  updatedBy?: Worker

  type?: AdditionType
  scope?: AdditionScope
  status?: AdditionStatus

  //
  isEphemeral?: boolean

  worker?: Worker

  name?: string
  value: number
}

//
export enum CompensationStatus {
  PENDING,
  APPROVED,
  REJECTED,
  PAID,
}

//
export enum AdditionScope {
  GLOBAL,
  ROLE,
  TEAM,
  WORKER,
}

//
export enum AdditionType {
  REIMBURSEMENT,
  BONUS,
  COMMISSION,
  OTHER,
}

//
export enum AdditionStatus {
  PENDING,
  VERIFIED,
  REJECT,
  WAITING,
  DISABLED,
}

// A type that represents a value that will either be deducted to the total pay
// of a worker on a specific pay period. Depending on the type of deduction, it can
// be cancel out and prevent any deduction to the final pay amount.
export type Deduction = {
  id: number
  createdById: number
  updatedById?: number
  organizationId: number
  workerId?: number

  createdAt: number
  lastUpdatedAt?: number

  createdBy: Worker
  updatedBy?: Worker

  effectiveAt?: number

  type?: DeductionType
  scope?: DeductionScope
  status?: DeductionStatus

  // Indicates whether this deduction is voluntary or not
  isVoluntary?: boolean

  //
  isEphemeral?: boolean

  worker?: Worker

  name?: string
  value: number
}

//
export enum DeductionScope {
  GLOBAL,
  ROLE,
  TEAM,
  WORKER,
}

//
export enum DeductionType {
  TAX,
  BENEFIT,
  REIMBURSEMENT,
  GARNISHMENT,
  OTHER,
}

//
export enum DeductionStatus {
  PENDING,
  VERIFIED,
  REJECT,
  WAITING,
  DISABLED,
}

//
export enum WorkerPayInfoType {
  HOURLY,
  SALARY,
  NONEXEMPT,
}
