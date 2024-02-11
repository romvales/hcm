export type Json =
  | string
  | number
  | boolean
  | null
  | { [key: string]: Json | undefined }
  | Json[]

export interface Database {
  graphql_public: {
    Tables: {
      [_ in never]: never
    }
    Views: {
      [_ in never]: never
    }
    Functions: {
      graphql: {
        Args: {
          operationName?: string
          query?: string
          variables?: Json
          extensions?: Json
        }
        Returns: Json
      }
    }
    Enums: {
      [_ in never]: never
    }
    CompositeTypes: {
      [_ in never]: never
    }
  }
  public: {
    Tables: {
      additions: {
        Row: {
          createdAt: string
          createdById: number
          effectiveAt: string | null
          id: number
          isEphemeral: boolean
          lastUpdatedAt: string | null
          name: string | null
          scope: number | null
          status: number | null
          type: number | null
          updatedById: number | null
          value: number
          workerId: number | null
        }
        Insert: {
          createdAt?: string
          createdById: number
          effectiveAt?: string | null
          id?: number
          isEphemeral?: boolean
          lastUpdatedAt?: string | null
          name?: string | null
          scope?: number | null
          status?: number | null
          type?: number | null
          updatedById?: number | null
          value?: number
          workerId?: number | null
        }
        Update: {
          createdAt?: string
          createdById?: number
          effectiveAt?: string | null
          id?: number
          isEphemeral?: boolean
          lastUpdatedAt?: string | null
          name?: string | null
          scope?: number | null
          status?: number | null
          type?: number | null
          updatedById?: number | null
          value?: number
          workerId?: number | null
        }
        Relationships: [
          {
            foreignKeyName: "additions_createdById_fkey"
            columns: ["createdById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "additions_updatedById_fkey"
            columns: ["updatedById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "additions_workerId_fkey"
            columns: ["workerId"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          }
        ]
      }
      attendances: {
        Row: {
          breakTime: number
          clockIn: string
          clockInType: number | null
          clockOut: string | null
          clockOutType: number | null
          computed: number
          createdAt: string
          createdById: number
          id: number
          isHoliday: boolean
          isLate: boolean
          isManual: boolean
          isOnBreak: boolean
          isOverride: boolean
          lastUpdatedAt: string | null
          lateTime: number
          oShiftId: number | null
          overTime: number
          perfLabel: number | null
          shiftId: number | null
          status: number | null
          type: number | null
          underTime: number
          updatedById: number | null
          workerId: number
        }
        Insert: {
          breakTime?: number
          clockIn?: string
          clockInType?: number | null
          clockOut?: string | null
          clockOutType?: number | null
          computed?: number
          createdAt?: string
          createdById: number
          id?: number
          isHoliday?: boolean
          isLate?: boolean
          isManual?: boolean
          isOnBreak?: boolean
          isOverride?: boolean
          lastUpdatedAt?: string | null
          lateTime?: number
          oShiftId?: number | null
          overTime?: number
          perfLabel?: number | null
          shiftId?: number | null
          status?: number | null
          type?: number | null
          underTime?: number
          updatedById?: number | null
          workerId: number
        }
        Update: {
          breakTime?: number
          clockIn?: string
          clockInType?: number | null
          clockOut?: string | null
          clockOutType?: number | null
          computed?: number
          createdAt?: string
          createdById?: number
          id?: number
          isHoliday?: boolean
          isLate?: boolean
          isManual?: boolean
          isOnBreak?: boolean
          isOverride?: boolean
          lastUpdatedAt?: string | null
          lateTime?: number
          oShiftId?: number | null
          overTime?: number
          perfLabel?: number | null
          shiftId?: number | null
          status?: number | null
          type?: number | null
          underTime?: number
          updatedById?: number | null
          workerId?: number
        }
        Relationships: [
          {
            foreignKeyName: "attendances_createdById_fkey"
            columns: ["createdById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "attendances_oShiftId_fkey"
            columns: ["oShiftId"]
            isOneToOne: false
            referencedRelation: "overrideShifts"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "attendances_shiftId_fkey"
            columns: ["shiftId"]
            isOneToOne: false
            referencedRelation: "standardShifts"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "attendances_updatedById_fkey"
            columns: ["updatedById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "attendances_workerId_fkey"
            columns: ["workerId"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          }
        ]
      }
      compensations: {
        Row: {
          approvedAt: string | null
          avalue: number
          createdAt: string
          createdById: number
          dvalue: number
          gvalue: number
          id: number
          lastUpdatedAt: string | null
          organizationId: number
          paidAt: string | null
          periodEnd: string
          periodStart: string
          rejectedAt: string | null
          status: number | null
          updatedById: number | null
          value: number
          workerId: number
        }
        Insert: {
          approvedAt?: string | null
          avalue?: number
          createdAt?: string
          createdById: number
          dvalue?: number
          gvalue: number
          id?: number
          lastUpdatedAt?: string | null
          organizationId: number
          paidAt?: string | null
          periodEnd: string
          periodStart: string
          rejectedAt?: string | null
          status?: number | null
          updatedById?: number | null
          value?: number
          workerId: number
        }
        Update: {
          approvedAt?: string | null
          avalue?: number
          createdAt?: string
          createdById?: number
          dvalue?: number
          gvalue?: number
          id?: number
          lastUpdatedAt?: string | null
          organizationId?: number
          paidAt?: string | null
          periodEnd?: string
          periodStart?: string
          rejectedAt?: string | null
          status?: number | null
          updatedById?: number | null
          value?: number
          workerId?: number
        }
        Relationships: [
          {
            foreignKeyName: "compensations_createdById_fkey"
            columns: ["createdById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "compensations_organizationId_fkey"
            columns: ["organizationId"]
            isOneToOne: false
            referencedRelation: "organizations"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "compensations_updatedById_fkey"
            columns: ["updatedById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "compensations_workerId_fkey"
            columns: ["workerId"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          }
        ]
      }
      compensationsAdditions: {
        Row: {
          additionId: number
          compensationId: number
          id: number
        }
        Insert: {
          additionId: number
          compensationId: number
          id?: number
        }
        Update: {
          additionId?: number
          compensationId?: number
          id?: number
        }
        Relationships: [
          {
            foreignKeyName: "compensationsAdditions_additionId_fkey"
            columns: ["additionId"]
            isOneToOne: false
            referencedRelation: "additions"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "compensationsAdditions_compensationId_fkey"
            columns: ["compensationId"]
            isOneToOne: false
            referencedRelation: "compensations"
            referencedColumns: ["id"]
          }
        ]
      }
      compensationsDeductions: {
        Row: {
          compensationId: number
          deductionId: number
          id: number
        }
        Insert: {
          compensationId: number
          deductionId: number
          id?: number
        }
        Update: {
          compensationId?: number
          deductionId?: number
          id?: number
        }
        Relationships: [
          {
            foreignKeyName: "compensationsDeductions_compensationId_fkey"
            columns: ["compensationId"]
            isOneToOne: false
            referencedRelation: "compensations"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "compensationsDeductions_deductionId_fkey"
            columns: ["deductionId"]
            isOneToOne: false
            referencedRelation: "deductions"
            referencedColumns: ["id"]
          }
        ]
      }
      deductions: {
        Row: {
          createdAt: string
          createdById: number
          effectiveAt: string | null
          id: number
          isEphemeral: boolean | null
          isVoluntary: boolean | null
          lastUpdatedAt: string | null
          name: string | null
          scope: number | null
          status: number | null
          type: number | null
          updatedById: number | null
          value: number
          workerId: number | null
        }
        Insert: {
          createdAt?: string
          createdById: number
          effectiveAt?: string | null
          id?: number
          isEphemeral?: boolean | null
          isVoluntary?: boolean | null
          lastUpdatedAt?: string | null
          name?: string | null
          scope?: number | null
          status?: number | null
          type?: number | null
          updatedById?: number | null
          value?: number
          workerId?: number | null
        }
        Update: {
          createdAt?: string
          createdById?: number
          effectiveAt?: string | null
          id?: number
          isEphemeral?: boolean | null
          isVoluntary?: boolean | null
          lastUpdatedAt?: string | null
          name?: string | null
          scope?: number | null
          status?: number | null
          type?: number | null
          updatedById?: number | null
          value?: number
          workerId?: number | null
        }
        Relationships: [
          {
            foreignKeyName: "deductions_createdById_fkey"
            columns: ["createdById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "deductions_updatedById_fkey"
            columns: ["updatedById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "deductions_workerId_fkey"
            columns: ["workerId"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          }
        ]
      }
      organizations: {
        Row: {
          createdAt: string
          createdById: number
          id: number
          industry: number | null
          lastUpdatedAt: string | null
          name: string
          overrideIndustry: string | null
          status: number | null
          updatedById: number | null
        }
        Insert: {
          createdAt?: string
          createdById: number
          id?: number
          industry?: number | null
          lastUpdatedAt?: string | null
          name: string
          overrideIndustry?: string | null
          status?: number | null
          updatedById?: number | null
        }
        Update: {
          createdAt?: string
          createdById?: number
          id?: number
          industry?: number | null
          lastUpdatedAt?: string | null
          name?: string
          overrideIndustry?: string | null
          status?: number | null
          updatedById?: number | null
        }
        Relationships: [
          {
            foreignKeyName: "organizations_createdById_fkey"
            columns: ["createdById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "organizations_updatedById_fkey"
            columns: ["updatedById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          }
        ]
      }
      organizationsMembers: {
        Row: {
          createdAt: string
          hiredAt: string | null
          hiredById: number | null
          id: number
          isDayOff: boolean
          isHired: boolean
          isOnCall: boolean
          isOnLeave: boolean
          isRemote: boolean
          isSuspended: boolean
          isTerminated: boolean
          lastUpdatedAt: string | null
          leaveAt: string | null
          organizationId: number
          returnedAt: string | null
          scheduledSuspensionAt: string | null
          status: number | null
          suspendedAt: string | null
          terminatedAt: string | null
          type: number | null
          workerId: number
        }
        Insert: {
          createdAt?: string
          hiredAt?: string | null
          hiredById?: number | null
          id?: number
          isDayOff?: boolean
          isHired?: boolean
          isOnCall?: boolean
          isOnLeave?: boolean
          isRemote?: boolean
          isSuspended?: boolean
          isTerminated?: boolean
          lastUpdatedAt?: string | null
          leaveAt?: string | null
          organizationId: number
          returnedAt?: string | null
          scheduledSuspensionAt?: string | null
          status?: number | null
          suspendedAt?: string | null
          terminatedAt?: string | null
          type?: number | null
          workerId: number
        }
        Update: {
          createdAt?: string
          hiredAt?: string | null
          hiredById?: number | null
          id?: number
          isDayOff?: boolean
          isHired?: boolean
          isOnCall?: boolean
          isOnLeave?: boolean
          isRemote?: boolean
          isSuspended?: boolean
          isTerminated?: boolean
          lastUpdatedAt?: string | null
          leaveAt?: string | null
          organizationId?: number
          returnedAt?: string | null
          scheduledSuspensionAt?: string | null
          status?: number | null
          suspendedAt?: string | null
          terminatedAt?: string | null
          type?: number | null
          workerId?: number
        }
        Relationships: [
          {
            foreignKeyName: "organizationsMembers_hiredById_fkey"
            columns: ["hiredById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "organizationsMembers_organizationId_fkey"
            columns: ["organizationId"]
            isOneToOne: false
            referencedRelation: "organizations"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "organizationsMembers_workerId_fkey"
            columns: ["workerId"]
            isOneToOne: false
            referencedRelation: "organizations"
            referencedColumns: ["id"]
          }
        ]
      }
      organizationsPendingRequests: {
        Row: {
          id: number
          organizationId: number
          requestId: number
        }
        Insert: {
          id?: number
          organizationId: number
          requestId: number
        }
        Update: {
          id?: number
          organizationId?: number
          requestId?: number
        }
        Relationships: [
          {
            foreignKeyName: "organizationsPendingRequests_organizationId_fkey"
            columns: ["organizationId"]
            isOneToOne: false
            referencedRelation: "organizations"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "organizationsPendingRequests_requestId_fkey"
            columns: ["requestId"]
            isOneToOne: false
            referencedRelation: "pendingJoinRequests"
            referencedColumns: ["id"]
          }
        ]
      }
      overrideShifts: {
        Row: {
          completedAt: string | null
          createdAt: string
          createdById: number
          day: number
          endsOn: string | null
          id: number
          lastUpdatedAt: string | null
          name: string
          organizationId: number
          overrideClockIn: string
          overrideClockOut: string
          startsOn: string
          status: number
          updatedById: number | null
          verifiedAt: string | null
          workerId: number | null
        }
        Insert: {
          completedAt?: string | null
          createdAt?: string
          createdById: number
          day: number
          endsOn?: string | null
          id?: number
          lastUpdatedAt?: string | null
          name: string
          organizationId: number
          overrideClockIn: string
          overrideClockOut: string
          startsOn: string
          status: number
          updatedById?: number | null
          verifiedAt?: string | null
          workerId?: number | null
        }
        Update: {
          completedAt?: string | null
          createdAt?: string
          createdById?: number
          day?: number
          endsOn?: string | null
          id?: number
          lastUpdatedAt?: string | null
          name?: string
          organizationId?: number
          overrideClockIn?: string
          overrideClockOut?: string
          startsOn?: string
          status?: number
          updatedById?: number | null
          verifiedAt?: string | null
          workerId?: number | null
        }
        Relationships: [
          {
            foreignKeyName: "overrideShifts_createdById_fkey"
            columns: ["createdById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "overrideShifts_organizationId_fkey"
            columns: ["organizationId"]
            isOneToOne: false
            referencedRelation: "organizations"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "overrideShifts_updatedById_fkey"
            columns: ["updatedById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "overrideShifts_workerId_fkey"
            columns: ["workerId"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          }
        ]
      }
      payrollComputedCompensations: {
        Row: {
          compensationId: number
          id: number
          payrollId: number
        }
        Insert: {
          compensationId: number
          id?: number
          payrollId: number
        }
        Update: {
          compensationId?: number
          id?: number
          payrollId?: number
        }
        Relationships: [
          {
            foreignKeyName: "payrollComputedCompensations_compensationId_fkey"
            columns: ["compensationId"]
            isOneToOne: false
            referencedRelation: "compensations"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "payrollComputedCompensations_payrollId_fkey"
            columns: ["payrollId"]
            isOneToOne: false
            referencedRelation: "payrolls"
            referencedColumns: ["id"]
          }
        ]
      }
      payrolls: {
        Row: {
          createdAt: string
          createdById: number
          id: number
          lastUpdatedAt: string | null
          organizationId: number | null
          payCycleType: number | null
          status: number | null
          total: number
          updatedById: number | null
          verifiedById: number | null
        }
        Insert: {
          createdAt?: string
          createdById: number
          id?: number
          lastUpdatedAt?: string | null
          organizationId?: number | null
          payCycleType?: number | null
          status?: number | null
          total?: number
          updatedById?: number | null
          verifiedById?: number | null
        }
        Update: {
          createdAt?: string
          createdById?: number
          id?: number
          lastUpdatedAt?: string | null
          organizationId?: number | null
          payCycleType?: number | null
          status?: number | null
          total?: number
          updatedById?: number | null
          verifiedById?: number | null
        }
        Relationships: [
          {
            foreignKeyName: "payrolls_createdById_fkey"
            columns: ["createdById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "payrolls_organizationId_fkey"
            columns: ["organizationId"]
            isOneToOne: false
            referencedRelation: "organizations"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "payrolls_updatedById_fkey"
            columns: ["updatedById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "payrolls_verifiedById_fkey"
            columns: ["verifiedById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          }
        ]
      }
      pendingJoinRequests: {
        Row: {
          createdAt: string
          expiredAt: string
          id: number
          organizationId: number
          status: number | null
          type: number | null
          workerId: number
        }
        Insert: {
          createdAt?: string
          expiredAt: string
          id?: number
          organizationId: number
          status?: number | null
          type?: number | null
          workerId: number
        }
        Update: {
          createdAt?: string
          expiredAt?: string
          id?: number
          organizationId?: number
          status?: number | null
          type?: number | null
          workerId?: number
        }
        Relationships: [
          {
            foreignKeyName: "pendingJoinRequests_organizationId_fkey"
            columns: ["organizationId"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "pendingJoinRequests_workerId_fkey"
            columns: ["workerId"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          }
        ]
      }
      roles: {
        Row: {
          createdAt: string
          createdById: number
          id: number
          lastUpdatedAt: string | null
          name: string
          organizationId: number
          status: number | null
          updatedById: number | null
        }
        Insert: {
          createdAt?: string
          createdById: number
          id?: number
          lastUpdatedAt?: string | null
          name: string
          organizationId: number
          status?: number | null
          updatedById?: number | null
        }
        Update: {
          createdAt?: string
          createdById?: number
          id?: number
          lastUpdatedAt?: string | null
          name?: string
          organizationId?: number
          status?: number | null
          updatedById?: number | null
        }
        Relationships: [
          {
            foreignKeyName: "roles_createdById_fkey"
            columns: ["createdById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "roles_organizationId_fkey"
            columns: ["organizationId"]
            isOneToOne: false
            referencedRelation: "organizations"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "roles_updatedById_fkey"
            columns: ["updatedById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          }
        ]
      }
      rolesStandardShifts: {
        Row: {
          id: number
          roleId: number
          standardShiftId: number
        }
        Insert: {
          id?: number
          roleId: number
          standardShiftId: number
        }
        Update: {
          id?: number
          roleId?: number
          standardShiftId?: number
        }
        Relationships: [
          {
            foreignKeyName: "rolesStandardShifts_roleId_fkey"
            columns: ["roleId"]
            isOneToOne: false
            referencedRelation: "roles"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "rolesStandardShifts_standardShiftId_fkey"
            columns: ["standardShiftId"]
            isOneToOne: false
            referencedRelation: "standardShifts"
            referencedColumns: ["id"]
          }
        ]
      }
      standardShifts: {
        Row: {
          clockIn: string
          clockOut: string
          createdAt: string
          createdById: number
          day: number
          id: number
          lastUpdatedAt: string | null
          name: string
          organizationId: number
          updatedById: number
        }
        Insert: {
          clockIn: string
          clockOut: string
          createdAt?: string
          createdById: number
          day: number
          id?: number
          lastUpdatedAt?: string | null
          name: string
          organizationId: number
          updatedById: number
        }
        Update: {
          clockIn?: string
          clockOut?: string
          createdAt?: string
          createdById?: number
          day?: number
          id?: number
          lastUpdatedAt?: string | null
          name?: string
          organizationId?: number
          updatedById?: number
        }
        Relationships: [
          {
            foreignKeyName: "standardShifts_createdById_fkey"
            columns: ["createdById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "standardShifts_organizationId_fkey"
            columns: ["organizationId"]
            isOneToOne: false
            referencedRelation: "organizations"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "standardShifts_updatedById_fkey"
            columns: ["updatedById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          }
        ]
      }
      teams: {
        Row: {
          createdAt: string
          createdById: number
          id: number
          lastUpdatedAt: string | null
          name: string
          organizationId: number
          status: number | null
          updatedById: number | null
        }
        Insert: {
          createdAt?: string
          createdById: number
          id?: number
          lastUpdatedAt?: string | null
          name: string
          organizationId: number
          status?: number | null
          updatedById?: number | null
        }
        Update: {
          createdAt?: string
          createdById?: number
          id?: number
          lastUpdatedAt?: string | null
          name?: string
          organizationId?: number
          status?: number | null
          updatedById?: number | null
        }
        Relationships: [
          {
            foreignKeyName: "teams_createdById_fkey"
            columns: ["createdById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "teams_organizationId_fkey"
            columns: ["organizationId"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "teams_updatedById_fkey"
            columns: ["updatedById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          }
        ]
      }
      teamsMembers: {
        Row: {
          id: number
          teamId: number
          workerId: number
        }
        Insert: {
          id?: number
          teamId: number
          workerId: number
        }
        Update: {
          id?: number
          teamId?: number
          workerId?: number
        }
        Relationships: [
          {
            foreignKeyName: "teamsMembers_teamId_fkey"
            columns: ["teamId"]
            isOneToOne: false
            referencedRelation: "teams"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "teamsMembers_workerId_fkey"
            columns: ["workerId"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          }
        ]
      }
      workerIdentityCards: {
        Row: {
          backImageUrl: string
          createdAt: string
          createdById: number | null
          extractedInfo: Json | null
          frontImageUrl: string
          id: number
          lastUpdatedAt: string | null
          name: string
          updatedById: number | null
          workerId: number
        }
        Insert: {
          backImageUrl?: string
          createdAt?: string
          createdById?: number | null
          extractedInfo?: Json | null
          frontImageUrl?: string
          id?: number
          lastUpdatedAt?: string | null
          name: string
          updatedById?: number | null
          workerId: number
        }
        Update: {
          backImageUrl?: string
          createdAt?: string
          createdById?: number | null
          extractedInfo?: Json | null
          frontImageUrl?: string
          id?: number
          lastUpdatedAt?: string | null
          name?: string
          updatedById?: number | null
          workerId?: number
        }
        Relationships: [
          {
            foreignKeyName: "workerIdentityCards_createdById_fkey"
            columns: ["createdById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "workerIdentityCards_updatedById_fkey"
            columns: ["updatedById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "workerIdentityCards_workerId_fkey"
            columns: ["workerId"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          }
        ]
      }
      workerOrganizations: {
        Row: {
          id: number
          organizationId: number
          workerId: number
        }
        Insert: {
          id?: number
          organizationId: number
          workerId: number
        }
        Update: {
          id?: number
          organizationId?: number
          workerId?: number
        }
        Relationships: [
          {
            foreignKeyName: "workerOrganizations_organizationId_fkey"
            columns: ["organizationId"]
            isOneToOne: false
            referencedRelation: "organizations"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "workerOrganizations_workerId_fkey"
            columns: ["workerId"]
            isOneToOne: false
            referencedRelation: "organizations"
            referencedColumns: ["id"]
          }
        ]
      }
      workerPendingRequests: {
        Row: {
          id: number
          requestId: number
          workerId: number
        }
        Insert: {
          id?: number
          requestId: number
          workerId: number
        }
        Update: {
          id?: number
          requestId?: number
          workerId?: number
        }
        Relationships: [
          {
            foreignKeyName: "workerPendingRequests_requestId_fkey"
            columns: ["requestId"]
            isOneToOne: false
            referencedRelation: "pendingJoinRequests"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "workerPendingRequests_workerId_fkey"
            columns: ["workerId"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          }
        ]
      }
      workerRoles: {
        Row: {
          id: number
          roleId: number
          workerId: number
        }
        Insert: {
          id?: number
          roleId: number
          workerId: number
        }
        Update: {
          id?: number
          roleId?: number
          workerId?: number
        }
        Relationships: [
          {
            foreignKeyName: "workerRoles_roleId_fkey"
            columns: ["roleId"]
            isOneToOne: false
            referencedRelation: "roles"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "workerRoles_workerId_fkey"
            columns: ["workerId"]
            isOneToOne: false
            referencedRelation: "roles"
            referencedColumns: ["id"]
          }
        ]
      }
      workers: {
        Row: {
          addresses: Json
          birthdate: string | null
          createdAt: string
          createdById: number | null
          email: string
          firstName: string
          gender: number
          id: number
          indicator: number
          lastName: string
          lastUpdatedAt: string | null
          middleName: string | null
          mobile: string | null
          pictureUrl: string | null
          suffix: string | null
          updatedById: number | null
          userId: string | null
          username: string
        }
        Insert: {
          addresses?: Json
          birthdate?: string | null
          createdAt?: string
          createdById?: number | null
          email: string
          firstName: string
          gender: number
          id?: number
          indicator?: number
          lastName: string
          lastUpdatedAt?: string | null
          middleName?: string | null
          mobile?: string | null
          pictureUrl?: string | null
          suffix?: string | null
          updatedById?: number | null
          userId?: string | null
          username: string
        }
        Update: {
          addresses?: Json
          birthdate?: string | null
          createdAt?: string
          createdById?: number | null
          email?: string
          firstName?: string
          gender?: number
          id?: number
          indicator?: number
          lastName?: string
          lastUpdatedAt?: string | null
          middleName?: string | null
          mobile?: string | null
          pictureUrl?: string | null
          suffix?: string | null
          updatedById?: number | null
          userId?: string | null
          username?: string
        }
        Relationships: [
          {
            foreignKeyName: "workers_createdById_fkey"
            columns: ["createdById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "workers_updatedById_fkey"
            columns: ["updatedById"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "workers_userId_fkey"
            columns: ["userId"]
            isOneToOne: false
            referencedRelation: "users"
            referencedColumns: ["id"]
          }
        ]
      }
      workersIdentityCards: {
        Row: {
          cardId: number
          id: number
          ownerId: number
        }
        Insert: {
          cardId: number
          id?: number
          ownerId: number
        }
        Update: {
          cardId?: number
          id?: number
          ownerId?: number
        }
        Relationships: [
          {
            foreignKeyName: "workersIdentityCards_cardId_fkey"
            columns: ["cardId"]
            isOneToOne: false
            referencedRelation: "workerIdentityCards"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "workersIdentityCards_ownerId_fkey"
            columns: ["ownerId"]
            isOneToOne: false
            referencedRelation: "workers"
            referencedColumns: ["id"]
          }
        ]
      }
    }
    Views: {
      [_ in never]: never
    }
    Functions: {
      [_ in never]: never
    }
    Enums: {
      [_ in never]: never
    }
    CompositeTypes: {
      [_ in never]: never
    }
  }
  storage: {
    Tables: {
      buckets: {
        Row: {
          allowed_mime_types: string[] | null
          avif_autodetection: boolean | null
          created_at: string | null
          file_size_limit: number | null
          id: string
          name: string
          owner: string | null
          owner_id: string | null
          public: boolean | null
          updated_at: string | null
        }
        Insert: {
          allowed_mime_types?: string[] | null
          avif_autodetection?: boolean | null
          created_at?: string | null
          file_size_limit?: number | null
          id: string
          name: string
          owner?: string | null
          owner_id?: string | null
          public?: boolean | null
          updated_at?: string | null
        }
        Update: {
          allowed_mime_types?: string[] | null
          avif_autodetection?: boolean | null
          created_at?: string | null
          file_size_limit?: number | null
          id?: string
          name?: string
          owner?: string | null
          owner_id?: string | null
          public?: boolean | null
          updated_at?: string | null
        }
        Relationships: []
      }
      migrations: {
        Row: {
          executed_at: string | null
          hash: string
          id: number
          name: string
        }
        Insert: {
          executed_at?: string | null
          hash: string
          id: number
          name: string
        }
        Update: {
          executed_at?: string | null
          hash?: string
          id?: number
          name?: string
        }
        Relationships: []
      }
      objects: {
        Row: {
          bucket_id: string | null
          created_at: string | null
          id: string
          last_accessed_at: string | null
          metadata: Json | null
          name: string | null
          owner: string | null
          owner_id: string | null
          path_tokens: string[] | null
          updated_at: string | null
          version: string | null
        }
        Insert: {
          bucket_id?: string | null
          created_at?: string | null
          id?: string
          last_accessed_at?: string | null
          metadata?: Json | null
          name?: string | null
          owner?: string | null
          owner_id?: string | null
          path_tokens?: string[] | null
          updated_at?: string | null
          version?: string | null
        }
        Update: {
          bucket_id?: string | null
          created_at?: string | null
          id?: string
          last_accessed_at?: string | null
          metadata?: Json | null
          name?: string | null
          owner?: string | null
          owner_id?: string | null
          path_tokens?: string[] | null
          updated_at?: string | null
          version?: string | null
        }
        Relationships: [
          {
            foreignKeyName: "objects_bucketId_fkey"
            columns: ["bucket_id"]
            isOneToOne: false
            referencedRelation: "buckets"
            referencedColumns: ["id"]
          }
        ]
      }
    }
    Views: {
      [_ in never]: never
    }
    Functions: {
      can_insert_object: {
        Args: {
          bucketid: string
          name: string
          owner: string
          metadata: Json
        }
        Returns: undefined
      }
      extension: {
        Args: {
          name: string
        }
        Returns: string
      }
      filename: {
        Args: {
          name: string
        }
        Returns: string
      }
      foldername: {
        Args: {
          name: string
        }
        Returns: unknown
      }
      get_size_by_bucket: {
        Args: Record<PropertyKey, never>
        Returns: {
          size: number
          bucket_id: string
        }[]
      }
      search: {
        Args: {
          prefix: string
          bucketname: string
          limits?: number
          levels?: number
          offsets?: number
          search?: string
          sortcolumn?: string
          sortorder?: string
        }
        Returns: {
          name: string
          id: string
          updated_at: string
          created_at: string
          last_accessed_at: string
          metadata: Json
        }[]
      }
    }
    Enums: {
      [_ in never]: never
    }
    CompositeTypes: {
      [_ in never]: never
    }
  }
}

export type Tables<
  PublicTableNameOrOptions extends
    | keyof (Database["public"]["Tables"] & Database["public"]["Views"])
    | { schema: keyof Database },
  TableName extends PublicTableNameOrOptions extends { schema: keyof Database }
    ? keyof (Database[PublicTableNameOrOptions["schema"]]["Tables"] &
        Database[PublicTableNameOrOptions["schema"]]["Views"])
    : never = never
> = PublicTableNameOrOptions extends { schema: keyof Database }
  ? (Database[PublicTableNameOrOptions["schema"]]["Tables"] &
      Database[PublicTableNameOrOptions["schema"]]["Views"])[TableName] extends {
      Row: infer R
    }
    ? R
    : never
  : PublicTableNameOrOptions extends keyof (Database["public"]["Tables"] &
      Database["public"]["Views"])
  ? (Database["public"]["Tables"] &
      Database["public"]["Views"])[PublicTableNameOrOptions] extends {
      Row: infer R
    }
    ? R
    : never
  : never

export type TablesInsert<
  PublicTableNameOrOptions extends
    | keyof Database["public"]["Tables"]
    | { schema: keyof Database },
  TableName extends PublicTableNameOrOptions extends { schema: keyof Database }
    ? keyof Database[PublicTableNameOrOptions["schema"]]["Tables"]
    : never = never
> = PublicTableNameOrOptions extends { schema: keyof Database }
  ? Database[PublicTableNameOrOptions["schema"]]["Tables"][TableName] extends {
      Insert: infer I
    }
    ? I
    : never
  : PublicTableNameOrOptions extends keyof Database["public"]["Tables"]
  ? Database["public"]["Tables"][PublicTableNameOrOptions] extends {
      Insert: infer I
    }
    ? I
    : never
  : never

export type TablesUpdate<
  PublicTableNameOrOptions extends
    | keyof Database["public"]["Tables"]
    | { schema: keyof Database },
  TableName extends PublicTableNameOrOptions extends { schema: keyof Database }
    ? keyof Database[PublicTableNameOrOptions["schema"]]["Tables"]
    : never = never
> = PublicTableNameOrOptions extends { schema: keyof Database }
  ? Database[PublicTableNameOrOptions["schema"]]["Tables"][TableName] extends {
      Update: infer U
    }
    ? U
    : never
  : PublicTableNameOrOptions extends keyof Database["public"]["Tables"]
  ? Database["public"]["Tables"][PublicTableNameOrOptions] extends {
      Update: infer U
    }
    ? U
    : never
  : never

export type Enums<
  PublicEnumNameOrOptions extends
    | keyof Database["public"]["Enums"]
    | { schema: keyof Database },
  EnumName extends PublicEnumNameOrOptions extends { schema: keyof Database }
    ? keyof Database[PublicEnumNameOrOptions["schema"]]["Enums"]
    : never = never
> = PublicEnumNameOrOptions extends { schema: keyof Database }
  ? Database[PublicEnumNameOrOptions["schema"]]["Enums"][EnumName]
  : PublicEnumNameOrOptions extends keyof Database["public"]["Enums"]
  ? Database["public"]["Enums"][PublicEnumNameOrOptions]
  : never

