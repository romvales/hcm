import { HCMOrganizationService, HCMPendingJoinRequestService, Organization, PendingJoinRequest } from '../../index.d'
import { SupabaseClientDatabase, Supabase_HCMWorkerService, isClientNotUndefined, isServiceDefined, isTargetNotDefined } from '.'


export class Supabase_HCMOrganizationService extends HCMOrganizationService
  implements HCMPendingJoinRequestService<Organization> {

  constructor(
    private client?: SupabaseClientDatabase,
    private workerService?: Supabase_HCMWorkerService,
    private target?: Organization,
  ) {
    super()

    if (!this.workerService) 
      this.workerService = new Supabase_HCMWorkerService(client)

  }

  private dependencies() {
    const client = this.client as SupabaseClientDatabase
    const workerService = this.workerService as Supabase_HCMWorkerService
    const target = this.target as Organization

    return { client, workerService, target }
  }

  private ensureClientAndTargetToBeDefined() {
    isClientNotUndefined(this.client)
    isTargetNotDefined(this.target)
  }

  private ensureClientServiceAndTargetToBeDefined() {
    isClientNotUndefined(this.client)
    isTargetNotDefined(this.target)
    isServiceDefined(this.workerService)
  }

  createOrg(name: string): Organization {
    const org: Organization = { name }
    this.setTarget(org)
    return org
  }

  async getOrgCreator() {
    this.ensureClientAndTargetToBeDefined()

    const { workerService, target } = this.dependencies()

    return workerService.getWorkerById(target.createdById)
  }

  async saveOrg() {
    this.ensureClientServiceAndTargetToBeDefined()

    const { client, target } = this.dependencies()
    const sessionUser = await this.client?.auth.getUser()
      .then(res => {
        const { data, error } = res
        if (error) {
          return
        }

        return data.user
      })

    const updatorWorker = await this.workerService?.getWorkerByUser(sessionUser)

    // New organization
    if (!target.createdById && updatorWorker) {
      target.createdById = updatorWorker.id
    }

    const query = client.from('organizations')
      .upsert(Object.assign(target))
      .select()
      .limit(1)
      .throwOnError()

    return query
      .then(res => {
        const updatedOrg = (res.data?.at(0) ?? {}) as Organization

        if (!updatedOrg.id) return {}

        this.setTarget(updatedOrg)

        return updatedOrg
      })
  }

  async getOrgById(organizationId: number): Promise<Organization | undefined> {
    isClientNotUndefined(this.client)
    
    const { client } = this.dependencies()

    const query = client.from('organizations')
      .select()
      .match({ id: organizationId })
      .throwOnError()

    return query
      .then(res => {
        const [ org ] = res.data as Organization[]
        return org
      })
  }

  async deleteOrgById(organizationId: number | undefined) {
    isClientNotUndefined(this.client)

    const { client } = this.dependencies()

    const query = client.from('organizations')
      .delete()
      .match({ id: organizationId })
      .throwOnError()

    return await query
  }

  async removeWorkerFromOrgById(workerId: number) {
    this.ensureClientAndTargetToBeDefined()
    
    const { client, target } = this.dependencies()

  }

  async addWorkerToOrgById(workerId: number) {
    this.ensureClientAndTargetToBeDefined()

    const { client, target } = this.dependencies()



  }

  async sendRequest(recepientId: number) {
    
  }

  async cancelRequest(recepientId: number) {
    
  }
  
  async getPendingRequests(): Promise<PendingJoinRequest[]> {
    
  }

  async acceptPendingRequest(requestId: number) {
    
  }

  async declinePendingRequest(requestId: number) {
    
  }

}