// @tags-mvp

import { Organization } from './index.d'
import { Worker } from './worker'

//
export abstract class PayrollService {

}

//
export type Payroll = {
  id: number
  createdById: number
  updatedById?: number
  verifiedById?: number
  organizationId: number

  createdAt: number
  lastUpdatedAt?: number

  createdBy: Worker
  updatedBy?: Worker
  verifiedBy?: Worker

  payCycleType?: PayrollPayCycleType
  status?: PayrollStatus
  organization: Organization

  compensations?: Compensation[]

  // A computed value that represents the total amount to be paid by the organization
  // to its workers.
  total: number
}

//
export enum PayrollStatus {
  PENDING,
  VERIFIED,
  PAID,
}

//
export enum PayrollPayCycleType {
  WEEKLY,
  BIWEEKLY,
  SEMIMONTHLY,
  MONTHLY,
  CUSTOM,
}