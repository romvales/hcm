// @tags-mvp

import type { 
  Worker,
} from './worker'

//
export abstract class HCMOrganizationService implements HCMPendingJoinRequestService<Organization> {
  createOrg(
    name: string,
    industry: OrganizationIndustry,
    overrideIndustry?: string
  ): Organization

  async getOrgById(organizationId: number): Promise<Organization>
  async removeOrgById<T>(organizationId: number): Promise<T>
  async saveOrg<T>(org: Organization): Promise<T>

  changeOrgName<T>(org: Organization, name: string): T
  changeOrgIndustry<T>(org: Organization, industry: OrganizationIndustry, overrideIndustry?: string): T
  changeOrgStatus<T>(org: Organization, status: OrganizationStatus): T

  async removeWorkerById<T>(workerId: number): Promise<T>
  async addToOrgById<T>(org: Organization, workerId: number): Promise<T>

  async getOrgCreator(org: Organization): Promise<Worker>
}

//
export interface HCMPendingJoinRequestService<Entity = unknown> {
  async sendRequest<T>(entity: Entity, recepientId: number): Promise<T>
  async cancelRequest<T>(entity: Entity, recepientId: number): Promise<T>

  async getPendingRequests(entity: Entity): Promise<PendingJoinRequest[]>
  async acceptPendingRequest<T>(entity: Entity, requestId: number): Promise<T>
  async declinePendingRequest<T>(entity: Entity, requestId: number): Promise<T>
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
  ORGANIZATION,
  WORKER,
}

export enum PendingJoinRequestStatus {
  PENDING,
  ACCEPTED,
  REJECTED,
  EXPIRED,
}