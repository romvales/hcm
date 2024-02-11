// @tags-mvp

import type {
  Organization,
} from './index.d'

import type { 
  Worker,
} from './worker'

//
export abstract class HCMAttendanceService {

  createAttendance(params: {
    worker: Worker,
    clockIn: number,
    clockOut?: number,
    clockInType: AttendanceClockInType,
    clockOutType?: AttendanceClockOutType,
    shift?: StandardShift | OverrideShift,
    isOverride?: boolean,
    perfLevel?: AttendancePerformanceLabel,
  }): Attendance

  async getAttendanceById(attendanceId: number): Promise<Attendance>
  async removeAttendanceById<T>(attendanceId: number): Promise<T>
  async saveAttendance<T>(attendance: Attendance): Promise<T>

  changeStatus(attendance: Attendance, status: AttendanceStatus)
  changeType(attendance: Attendance, type: AttendanceType)
  changePerfLabel(attendance: Attendance, label: AttendancePerformanceLabel)
  changeClockInType(attendance: Attendance, type: AttendanceClockInType)
  changeClockOutType(attendance: Attendance, type: AttendanceClockOutType)

  async clockIn<T>(worker: Worker, type: AttendanceClockInType): Promise<T>
  async clockOut<T>(worker: Worker, type: AttendanceClockOutType): Promise<T>

  async getShift(attendance: Attendance): Promise<StandardShift | OverrideShift>

  isLate(attendance: Attendance): boolean
  isOverride(attendance: Attendance): boolean
  isHoliday(attendance: Attendance): boolean
  isBreak(attendance: Attendance): boolean
}

//
export abstract class HCMWorkerShiftService {
  
}

// 
export type Attendance = {
  id: number
  workerId: number
  shiftId?: number
  oshiftId?: number
  createdById: number
  updatedById?: number

  createdAt: string
  lastUpdatedAt?: string

  createdBy?: Worker
  updatedBy?: Worker

  clockInType?: AttendanceClockInType
  clockOutType?: AttendanceClockOutType
  type?: AttendanceType
  status?: AttendanceStatus
  perfLabel?: AttendancePerformanceLabel

  // Sets to true when the worker did not met the shift in the 7-day standard shift
  isLate?: boolean

  // If the standard shift is overriden by an HR/Manager, set this to true
  isOverride?: boolean

  // When a worker still works during a holiday, set this to true
  isHoliday?: boolean

  // Toggled to true when the HR is responsible for creating this attendance
  // on behalf of the worker.
  isManual?: boolean

  // 
  isOnBreak?: boolean

  worker: Worker
  shift: StandardShift | OverrideShift

  clockIn?: number
  clockOut?: number

  // Computed props stored in the database
  computed?: number
  underTime?: number
  overTime?: number
  lateTime?: number
  breakTime?: number
}

//
export enum AttendancePerformanceLabel {
  BELOW,
  POOR,
  NORMAL,
  GOOD,
  PRODUCTIVE,
}

//
export enum AttendanceClockInType {
  ONCALLSHIFT,
  HOLIDAYSHIFT,
  NIGHTSHIFT,
  NORMALSHIFT,
}

//
export enum AttendanceClockOutType {
  LUNCHTIME,
  BREAKTIME,
  ENDSHIFT,
  MEETING,
  EMERGENCY,
  CUSTOM,
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

  createdAt: string
  lastUpdatedAt?: string

  createdBy: Worker
  updatedBy?: Worker

  organization?: Organization
  day?: StandardShiftDay

  name: string

  clockIn: number
  clockOut: number
}

// Override shift is a scheduled override shift assigned to a specific
// worker in an organization.
export type OverrideShift = {
  id: number
  createdById: number
  updatedById: number
  organizationId?: number
  workerId: number

  createdAt: string
  lastUpdatedAt?: string
  verifiedAt?: string
  completedAt?: string
  
  startsOn?: number
  endsOn?: number

  createdBy: Worker
  updatedBy?: Worker
  
  organization?: Organization
  assignedTo: Worker
  day: StandardShiftDay
  status: OverrideShiftStatus
  
  overrideClockIn: number
  overrideClockOut: number
}

//
export enum OverrideShiftStatus {
  PENDING,
  VERIFIED,
  ASSIGNED,
  DONE
}

// 
enum StandardShiftDay {
  MONDAY,
  TUESDAY,
  WEDNESDAY,
  THURSDAY,
  FRIDAY,
  SATURDAY,
  SUNDAY,
}