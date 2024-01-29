// @tags-mvp

import { Organization } from './index.d'
import { Worker } from './worker'


//
export type Payroll = {
  id: number
  createdById: number
  updatedById?: number
  organizationId: number

  createdAt: number
  lastUpdatedAt?: number

  createdBy: Worker
  updatedBy?: Worker

  organization: Organization

  compensations?: Compensation[]

  // A computed value that represents the total amount to be paid by the organization
  // to its workers.
  total: number
}