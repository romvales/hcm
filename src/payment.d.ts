// @tags-mvp

import { isTargetNotDefined } from './implementations/supabase'
import { Organization } from './index.d'
import type {
  Worker
} from './worker'

//
export abstract class HCMWorkerPayInfoService {

  setTarget(target: Worker) {
    this.target = target
    return this
  }

  createWorkerPayInfo(params: {
    
  }): WorkerPayInfo

  async getWorkerPayInfo(worker: Worker): Promise<WorkerPayInfo>
  async removeWorkerPayInfo<T>(worker: Worker): Promise<WorkerPayInfoOverride>
  async saveWorkerPayInfo<T>(payInfo: WorkerPayInfo): Promise<T>

}

//
export abstract class HCMCompensationService {

  setTarget(target: Compensation) {
    this.target = target
    return this
  }

  createCompensation(params: {
    worker: Worker
    periodStart: string
    periodEnd: string
  }): Compensation

  async getCompensationById(compensationId: number): Promise<Compensation>
  async removeCompensationById<T>(compensationId: number): Promise<T>
  async saveCompensation<T>(compensation: Compensation): Promise<T>

  getGrossValue(compensation: Compensation): number | null
  getAddedValue(compensation: Compensation): number | null
  getDeductedValue(compensation: Compensation): number | null
  getValue(compensation: Compensation): number | null

  async getAdditions(): Promise<Addition[]>
  async getDeductions(): Promise<Deduction[]>

  changeStatus(compensation: Compensation, status: CompensationStatus)
  
}

//
export abstract class HCMAdditionService {

  setTarget(target: Addition) {
    this.target = target
    return this
  }

  createAddition(name: string, value: number)

  async getAdditionById(additionId: number): Promise<Addition | undefined>
  async deleteAdditionById(additionId: number)
  async saveAddition(additionId: number)

  changeType(type: AdditionType) {
    isTargetNotDefined(this.target)

    const target = this.target as Addition
    target.type = type

    return this
  }

  changeScope(scope: AdditionScope) {
    isTargetNotDefined(this.target)

    const target = this.target as Addition
    target.scope = scope

    return this
  }

  changeStatus(status: AdditionStatus) {
    isTargetNotDefined(this.target)

    const target = this.target as Addition
    target.status = status

    return this
  }

  setEphemeral(state: boolean) {
    isTargetNotDefined(this.target)

    const target = this.target as Addition
    target.isEphemeral = state

    return this
  }

  changeName(newName: string) {
    isTargetNotDefined(this.target)

    const target = this.target as Addition
    target.name = newName

    return this
  }

  changeValue(newValue: number) {
    isTargetNotDefined(this.target)

    const target = this.target as Addition
    target.value = newValue

    return this
  }

  changeEffectiveDate(newDate: string) {
    isTargetNotDefined(this.target)

    const target = this.target as Addition
    target.effectiveAt = newDate

    return this
  }

  assignAdditionToWorker(worker: Worker) {
    isTargetNotDefined(this.target)

    const target = this.target as Addition
    target.workerId = worker.id

    return this
  }

}

//
export abstract class HCMDeductionService {

  setTarget(target: Deduction) {
    this.target = target
    return this
  }

  createDeduction(name: string, value: number)

  async getDeductionById(deductionId: number): Promise<Deduction | undefined>
  async deleteDeductionById(deductionId: number)
  async saveDeduction()

  setEphemeral(state: boolean) {
    isTargetNotDefined(this.target)

    const target = this.target as Deduction
    target.isEphemeral = state

    return this
  }

  changeType(type: DeductionType) {
    isTargetNotDefined(this.target)

    const target = this.target as Deduction
    target.type = type

    return this
  }

  changeScope(scope: DeductionScope) {
    isTargetNotDefined(this.target)

    const target = this.target as Deduction
    target.scope = scope

    return this
  }

  changeStatus(status: DeductionStatus) {
    isTargetNotDefined(this.target)

    const target = this.target as Deduction
    target.status = status

    return this
  }

  changeVoluntary(state: boolean) {
    isTargetNotDefined(this.target)

    const target = this.target as Deduction
    target.isVoluntary = state

    return this
  }

  changeName(newName: string) {
    isTargetNotDefined(this.target)

    const target = this.target as Deduction
    target.name = newName

    return this
  }

  changeValue(newValue: number) {
    isTargetNotDefined(this.target)

    const target = this.target as Deduction
    target.value = newValue

    return this
  }


  assignDeductionToWorker(worker: Worker) {
    isTargetNotDefined(this.target)

    const target = this.target as Deduction
    target.workerId = worker.id

    return this
  }

}

// Represents the default worker payment info that will be used in the computation
// it can be overriden by anyone by specifying a WorkerPayInfoOverride.
export type WorkerPayInfo = {
  id: number
  createdById: number
  updatedById: number
  workerId: number

  createdAt: string
  lastUpdatedAt?: string

  createdBy: Worker
  updatedBy?: Worker

  type?: WorkerPayInfoType

  worker: Worker

  // A collection of all benefits that the worker payment info has received.
  additions?: Addition[]

  // Same with the above collection, this contains the deductions the worker received.
  deductions?: Deduction[]

  hourly?: number
  salary?: number
  nonExempt?: number
}

// An HR can override a worker pay info either temporarily, for a given period,
// or permanently.
export type WorkerPayInfoOverride = {
  id: number
  payId: number
  createdById: number
  updatedById?: number

  createdAt: string
  lastUpdatedAt?: string
  startsOn?: string
  endsOn?: string
  
  pay: WorkerPayInfo
  type: WorkerPayInfoType
  status?: WorkerPayInfoOverrideStatus

  //
  isEnabled?: boolean

  //
  additions?: Addition[]

  //
  deductions?: Deduction[]

  //
  hourly?: number
  salary?: number
  nonExempt?: number
}

export enum WorkerPayInfoOverrideStatus {
  PENDING,
  ONGOING,
  OVERRIDE,
}

//
export type Compensation = {
  id: number
  createdById: number
  updatedById?: number
  organizationId: number
  workerId: number

  createdAt: string
  lastUpdatedAt?: string
  paidAt?: string
  approvedAt?: string
  rejectedAt?: string

  createdBy: Worker
  updatedBy?: Worker

  organization: Organization
  status?: CompensationStatus

  worker: Worker

  additions?: Addition[]
  deductions?: Deduction[]

  // 
  periodStart: string
  periodEnd: string

  // A value which holds the gross pay of the worker
  gvalue: number

  // Represents a computed value that is added to the final take-home pay of the worker
  avalue?: number

  // Computed value that denotes the total deductions in the take-home pay.
  dvalue?: number

  // 
  value: number
}

// Additions is a value added to the final pay for a worker on any specific pay period.
export type Addition = {
  id?: number
  createdById?: number
  updatedById?: number
  organizationId?: number
  workerId?: number

  createdAt?: string
  lastUpdatedAt?: string
  effectiveAt?: string

  createdBy?: Worker
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
  id?: number
  createdById?: number
  updatedById?: number
  workerId?: number

  createdAt?: string
  lastUpdatedAt?: string

  createdBy?: Worker
  updatedBy?: Worker

  effectiveAt?: string

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
