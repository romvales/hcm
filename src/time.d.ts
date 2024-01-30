// @tags-mvp

import type {
  Organization,
} from './index.d'

import type { 
  Worker,
} from './worker'

//
export abstract class HCMAttendanceService {

}

//
export abstract class HCMWorkerShiftService {
  
}

// 
export type Attendance = {
  id: number
  workerId: number
  createdById: number
  updatedById?: number

  createdAt: number
  lastUpdatedAt?: number

  createdBy?: Worker
  updatedBy?: Worker

  type?: AttendanceType
  status?: AttendanceStatus

  // Sets to true when the worker did not met the shift in the 7-day standard shift
  isLate?: boolean

  // If the standard shift is overriden by an HR/Manager, set this to true
  isOverride?: boolean

  // When a worker still works during a holiday, set this to true
  isHoliday?: boolean

  worker: Worker

  clockIn?: number
  clockOut?: number

  // Computed props stored in the database
  computed?: number
  underTime?: number
  overTime?: number
}

// 
export enum AttendanceType {
  ONCALL,
  DAYOFF,
  HALFDAY,
  NIGHT,
}

// 
export enum AttendanceStatus {
  PRESENT,
  LATE,
  OVERRIDE,
  HOLIDAY,
  OVERTIME,
}

// Every role should be assigned with a standard 7-day shift. Optionall
// a worker can have an overriden standard shift specify only to the worker.
export type StandardShift = {
  id: number
  organizationId?: number
  createdById: number
  updatedById: number

  createdAt: number
  lastUpdatedAt?: number

  createdBy: Worker
  updatedBy?: Worker

  organization?: Organization
  day?: StandardShiftDay

  name: string

  clockIn: number
  clockOut: number
}

// Override shift is a scheduled override shift assigned to a specific
// worker in the company.
export type OverrideShift = {
  id: number
  createdById: number
  updatedById: number
  organizationId?: number
  workerId: number

  createdBy: Worker
  updatedBy?: Worker

  createdAt: number
  lastUpdatedAt?: number
  startsOn?: number
  endsOn?: number
  
  organization?: Organization
  assignedTo: Worker
  
  overrideClockIn: number
  overrideClockOut: number
}

// 
enum StandardShiftDay {
  MONDAY,
  TUESDAY,
  WEDNESDAY,
  THURSDAY,
  FRIDAY,
  SATURDAY,
}