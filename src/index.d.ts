// @tags-mvp

import type { 
  Worker,
} from './worker'

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
}

export enum OrganizationStatus {
  OS_ACTIVE,
  OS_INACTIVE,
  OS_SUSPENDED,
  OS_DISSOLVED,
}

export enum OrganizationIndustry {
  OI_AGRICULTURE,
  OI_PRODUCTION,
  OI_CHEMICAL,
  OI_COMMERCE,
  OI_CONSTRUCTION,
  OI_EDUCATION,
  OI_FINANCIAL,
  OI_RETAIL,
  OI_FORESTRY,
  OI_HEALTH,
  OI_HOSPITALITY,
  OI_MINING,
  OI_MECHANICAL,
  OI_PUBLIC_SERVICE,
  OI_TELECOM,
  OI_SHIPPING,
  OI_TEXTILE,
  OI_TRANSPORT,
  OI_EQUIPMENT,
  OI_UTILITIES,
  OI_OTHER
}
