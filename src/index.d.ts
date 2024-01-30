// @tags-mvp

import type { 
  Worker,
} from './worker'

//
export abstract class HCMOrganizationService {
  
  static createOrg(
    name: string,
    industry: OrganizationIndustry,
    overrideIndustry?: string
  ): Organization

  static async getOrgById(organizationId: number): Promise<Organization>
  static async removeOrgById<T>(organizationId: number): Promise<T>
  static async saveOrg<T>(org: Organization): Promise<T>

  static changeOrgName<T>(org: Organization, name: string): T
  static changeOrgIndustry<T>(org: Organization, industry: OrganizationIndustry, overrideIndustry?: string): T
  static changeOrgStatus<T>(org: Organization, status: OrganizationStatus): T

  static async removeWorkerById<T>(workerId: number): Promise<T>
  static async addToOrgById<T>(org: Organization, workerId: number): Promise<T>

  static async getOrgCreator(org: Organization): Promise<Worker>
}

//
export interface HCMPendingJoinRequestService<Entity = unknown> {
  static async sendRequest<T>(entity: Entity, recepientId: number): Promise<T>
  static async cancelRequest<T>(entity: Entity, recepientId: number): Promise<T>

  static async getPendingRequests(entity: Entity): Promise<PendingJoinRequest[]>
  static async acceptPendingRequest<T>(entity: Entity, requestId: number): Promise<T>
  static async declinePendingRequest<T>(entity: Entity, requestId: number): Promise<T>
}

//
export type Organization = {
  id: number

  createdBy: Worker
  updatedBy?: Worker

  createdAt?: number
  lastUpdatedAt?: number

  status?: OrganizationStatus

  name: string
  industry?: OrganizationIndustry
  overrideIndustry?: string

  members?: Worker[]
  requests?: PendingOrganizationRequest[]
}

//
export type PendingJoinRequest = {
  id: number
  workerId: number
  organizationId: number

  createdAt: number
  expiredAt: number

  status?: PendingJoinRequestStatus
  type?: PendingJoinRequestInvitationType

  organization: Organization
  worker: Worker
}

export enum OrganizationStatus {
  ACTIVE,
  INACTIVE,
  SUSPENDED,
  DISSOLVED,
}

export enum OrganizationIndustry {
  AGRICULTURE,
  PRODUCTION,
  CHEMICAL,
  COMMERCE,
  CONSTRUCTION,
  EDUCATION,
  FINANCIAL,
  RETAIL,
  FORESTRY,
  HEALTH,
  HOSPITALITY,
  MINING,
  MECHANICAL,
  PUBLIC_SERVICE,
  TELECOM,
  SHIPPING,
  TEXTILE,
  TRANSPORT,
  EQUIPMENT,
  UTILITIES,
  OTHER
}

export enum PendingJoinRequestInvitationType {
  
}

export enum PendingJoinRequestStatus {
  PENDING,
  ACCEPTED,
  REJECTED,
  EXPIRED,
}