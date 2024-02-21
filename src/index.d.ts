// @tags-mvp
import { isTargetNotDefined } from './implementations/supabase'
import { Worker, HCMServiceUtility } from './index'

//
export abstract class HCMOrganizationService {

  setTarget(target: Organization) {
    this.target = target
    return this
  }

  createOrg(name: string): Organization

  async getOrgCreator(): Promise<Worker | undefined>
  async getOrgById(organizationId: number): Promise<Organization | undefined>
  async deleteOrgById(organizationId: number)
  async saveOrg()

  changeOrgName(name: string) {
    isTargetNotDefined(this.target)

    const target = this.target as Organization
    target.name = name

    return this
  }

  changeOrgIndustry(industry: OrganizationIndustry, overrideIndustry?: string) {
    isTargetNotDefined(this.target)

    const target = this.target as Organization
    if ((target.industry = industry) == OrganizationIndustry.OTHER) {
      target.overrideIndustry = overrideIndustry
    }

    return this
  }

  changeOrgStatus(status: OrganizationStatus) {
    isTargetNotDefined(this.target)

    const target = this.target as Organization
    target.status = status

    return this
  }

  async removeWorkerFromOrgById(workerId: number)
  async addWorkerToOrgById(workerId: number)

  private isOrganizationStatusCheckAgainst(status: OrganizationStatus) {
    isTargetNotDefined(this.target)

    const target = this.target as Organization
    
    return target.status == status
  }

  isActive = 
    () => this.isOrganizationStatusCheckAgainst(OrganizationStatus.ACTIVE)
  isInactive =
    () => this.isOrganizationStatusCheckAgainst(OrganizationStatus.INACTIVE)
  isSuspended =
    () => this.isOrganizationStatusCheckAgainst(OrganizationStatus.SUSPENDED)
  isDissolved =
    () => this.isOrganizationStatusCheckAgainst(OrganizationStatus.DISSOLVED)

}

//
export interface HCMPendingJoinRequestService<Entity = unknown> {
  async sendRequest(recepientId: number)
  async cancelRequest(recepientId: number)

  async getPendingRequests(): Promise<PendingJoinRequest[]>
  async acceptPendingRequest(requestId: number)
  async declinePendingRequest(requestId: number)
}

//
export type Organization = {
  id?: number

  createdById?: number
  updatedById?: number

  createdBy?: Worker
  updatedBy?: Worker

  createdAt?: string
  lastUpdatedAt?: string

  status?: OrganizationStatus

  // TODO: Add fields that are necessary for an organization

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

  createdAt: string
  expiredAt: string

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