import { SupabaseClient, User } from '@supabase/supabase-js'
import { Database } from '../../database'
import { Worker } from '../../worker'

export { Supabase_HCMAdditionService } from './HCMAdditionService'
export { Supabase_HCMAttendanceService } from './HCMAttendanceService'
export { Supabase_HCMCompensationService } from './HCMCompensationService'
export { Supabase_HCMDeductionService } from './HCMDeductionService'
export { Supabase_HCMOrganizationService } from './HCMOrganizationService'
export { Supabase_HCMPayrollService } from './HCMPayrollService'
export { Supabase_HCMRoleService } from './HCMRoleService'
export { Supabase_HCMTeamService } from './HCMTeamService'
export { Supabase_HCMWorkerPayInfoService } from './HCMWorkerPayInfoService'
export { Supabase_HCMWorkerService } from './HCMWorkerService'
export { Supabase_HCMWorkerShiftService } from './HCMWorkerShiftService'

export type SupabaseClientDatabase = SupabaseClient<Database>

export class ClientUndefinedError extends Error {

  constructor(message?: string | undefined, options?: ErrorOptions | undefined) {
    super(message, options)
  }

}

export class TargetUndefinedError extends Error {

  constructor(message?: string | undefined, options?: ErrorOptions | undefined) {
    super(message, options)
  }

}

export class UserUndefinedError extends Error {

  constructor(message?: string | undefined, options?: ErrorOptions | undefined) {
    super(message, options)
  }

}

export class ServiceUndefinedError extends Error {

  constructor(message?: string | undefined, options?: ErrorOptions | undefined) {
    super(message, options)
  }

}

export class WorkerUndefinedError extends Error {

  constructor(message?: string | undefined, options?: ErrorOptions | undefined) {
    super(message, options)
  }

}

export const isClientNotUndefined = (client: SupabaseClient | undefined) => {
  if (typeof client === 'undefined') throw new ClientUndefinedError()
}

export const isTargetNotDefined = (target: unknown | undefined) => {
  if (typeof target === 'undefined') throw new TargetUndefinedError()
}

export const isUserDefined = (user: User | null | undefined) => {
  if (!user) throw new UserUndefinedError()
}

export const isServiceDefined = (service: unknown | undefined) => {
  if (typeof service === 'undefined') throw new ServiceUndefinedError()
}

export const isWorkerDefined = (worker: Worker | undefined) => {
  if (typeof worker === 'undefined') throw new WorkerUndefinedError()
}