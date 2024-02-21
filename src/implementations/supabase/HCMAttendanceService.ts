import { SupabaseClientDatabase, Supabase_HCMWorkerService, isClientNotUndefined, isServiceDefined, isTargetNotDefined } from '.'
import { 
  Worker, 
  
  Attendance, AttendanceClockInType, AttendanceClockOutType, AttendancePerformanceLabel, 
  HCMAttendanceService, OverrideShift, StandardShift } from '../../../src'

export class Supabase_HCMAttendanceService extends HCMAttendanceService {

  constructor(
    private client: SupabaseClientDatabase,
    private workerService?: Supabase_HCMWorkerService,
    private target?: Attendance
  ) {
    super()

  }

  private dependencies() {
    const client = this.client
    const workerService = this.workerService as Supabase_HCMWorkerService
    const target = this.target as Attendance

    return { client, workerService, target }
  }

  private ensureClientToBeDefined() {
    isClientNotUndefined(this.client)
  }

  private ensureClientWorkerToBeDefined() {
    isClientNotUndefined(this.client)
    isServiceDefined(this.workerService)
  }

  private ensureClientWorkerServiceTargetToBeDefined() {
    isClientNotUndefined(this.client)
    isServiceDefined(this.workerService)
    isTargetNotDefined(this.target)
  }

  createAttendance(params: {
    worker: Worker
    clockIn: number
    clockOut?: number
    clockInType: AttendanceClockInType
    clockOutType?: AttendanceClockOutType
    shift?: StandardShift | OverrideShift
    isOverride?: boolean
    perfLevel?: AttendancePerformanceLabel
  }): Attendance {
    const { 
      worker, 
      shift,
      
      clockIn, 
      clockOut,
      clockInType,
      clockOutType,
      isOverride,
      perfLevel } = params

    if (worker?.id) return {} as Worker

    const attendance: Attendance = {
      workerId: worker?.id,
      clockIn,
      clockOut,
      clockInType,
      clockOutType,
      isOverride,
      perfLabel: perfLevel,
    }

    if ((shift as StandardShift).clockIn) {
      attendance.shiftId = shift?.id
    } else {
      attendance.oshiftId = shift?.id
    }

    return attendance
  }

  async getAttendanceById(attendanceId: number): Promise<Attendance> {
    
  }

  async deleteAttendanceById(attendanceId: number) {
    
  }

  async saveAttendance() {

  }


}