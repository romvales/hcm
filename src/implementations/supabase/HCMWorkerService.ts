  import { SupabaseClientDatabase, isClientNotUndefined, isTargetNotDefined } from '.'
  import { HCMPendingJoinRequestService, Organization } from '../../index.d'
  import { 
    HCMWorkerOrganizationService,
    HCMWorkerService,
    Role,
    Team,
    Worker, 
    WorkerAddress,
    WorkerIdentityCard, 
    WorkerOrganizationInfo } from '../../worker.d'

  import { User, UserResponse } from '@supabase/supabase-js'

  export class Supabase_HCMWorkerService extends HCMWorkerService {

    constructor(
      private client?: SupabaseClientDatabase,
      private target?: Worker,
    ) {
      super()
    }

    private dependencies() {
      const target = this.target as Worker
      const client = this.client as SupabaseClientDatabase

      return { client, target }
    }

    private ensureClientAndTargetToBeDefined() {
      isClientNotUndefined(this.client)
      isTargetNotDefined(this.target)
    }

    assignUserToWorker(user: User | null) {
      this.ensureClientAndTargetToBeDefined()

      const target = this.target as Worker
      target.userId = user?.id

      return this
    }

    // createWorker() -> Creates a new worker record without saving it to the database
    createWorker(email: string, username: string): Worker {
      isClientNotUndefined(this.client)

      const worker = { 
        email, 
        username,
        createdById: undefined,
        updatedById: undefined,
        identityCards: [],
        addresses: [],
      }

      this.setTarget(worker)

      return worker
    }

    async getWorkerByUser(user?: User): Promise<Worker | undefined> {
      isClientNotUndefined(this.client)

      if (!user) return

      const { client } = this.dependencies()

      const query = client.from('workers')
        .select()
        .match({ userId: user?.id })
        .limit(1)
        .throwOnError()

      return query
        .then<Worker>(res => {
          const worker = (res.data?.at(0) ?? {}) as Worker
          
          return worker
        })
    }

    async getWorkerById(id?: number): Promise<Worker | undefined> {
      isClientNotUndefined(this.client)

      if (!id) return
      
      const { client } = this.dependencies()
      const query = client.from('workers')
        .select()
        .match({ id })
        .throwOnError()

      return query
        .then(res => {
          const worker = (res.data?.at(0) ?? {}) as Worker

          if (worker?.id) {


            return worker
          }

          return
        })
    }

    async getIdentityCards(): Promise<WorkerIdentityCard[]> {
      this.ensureClientAndTargetToBeDefined()

      const { client, target } = this.dependencies()

      const ids = await client.from('workersIdentityCards')
        .select(`
          cards!workersIdentityCards_cardId_fkey {
            *
          }
        `)
        .match({ workerId: target.id })
        .throwOnError()
        .then(res => res.data) ?? []

      return ids
    }

    getAddresses(): WorkerAddress[] {
      isTargetNotDefined(this.target)
      
      const { target } = this.dependencies()

      return target.addresses ?? []
    }

    async deleteWorkerById(id?: number) {
      isClientNotUndefined(this.client)

      if (typeof id == 'undefined') {
        return
      }

      const { client } = this.dependencies()
      const query = client.from('workers')
        .delete()
        .eq('id', id)
        .throwOnError()

      return await query
    }

    async deleteIdentityCardById(id: number) {
      isClientNotUndefined(this.client)

      const { client } = this.dependencies()
      const query = client.from('workersIdentityCards')
        .delete()
        .eq('id', id)
        .throwOnError()

      return await query
    }

    // Saves a worker related data to the postgres db
    async saveWorker(): Promise<Worker> {
      this.ensureClientAndTargetToBeDefined()

      const { client, target } = this.dependencies()

      const sessionUser = await client?.auth.getUser()
        .then(res => {
          const { data, error } = res
          if (error) {
            return undefined
          }

          return data.user
        })

      const updatorWorker = await this.getWorkerByUser(sessionUser)

      // New worker
      if (!target.createdById && updatorWorker) {
        target.createdById = updatorWorker.id
      }

      // Save the identityCards to the table for identity cards.
      const identityCards = target.identityCards ?? []
      target.identityCards = undefined

      const actions: Promise<WorkerIdentityCard>[] = []

      for (const id of identityCards) actions.push(this.saveWorkerIdentityCard(id))
      await Promise.all(actions)

      // When the updator is the same as the worker, don't update
      // who updated the record.
      if (updatorWorker && updatorWorker.id !== target.id) {    
        target.updatedById = updatorWorker.id
      }

      const query = client.from('workers')
        .upsert(Object.assign(target))
        .select()
        .limit(1)
        .throwOnError()

      return query
        .then(res => {
          const updatedWorker = (res.data?.at(0) ?? {}) as Worker

          if (!updatedWorker.id) return {} as Worker

          // TODO: Include relational data to the final result
          this.setTarget(updatedWorker)

          return updatedWorker
        })
    }

    async saveWorkerIdentityCard(id: WorkerIdentityCard) {
      isClientNotUndefined(this.client)

      const { client } = this.dependencies()

      const query = client.from('workerIdentityCards')
        .upsert(Object.assign(id))
        .select()
        .limit(1)
        .throwOnError()

      return query
        .then(res => {
          const id = (res.data?.at(0) ?? {}) as WorkerIdentityCard
          return id
        })
    }

    // Adds new identity cards for a worker
    addIdentityCards(identityCards: WorkerIdentityCard[]) {
      isTargetNotDefined(this.target)

      const { target } = this.dependencies()

      if (!target.identityCards?.length) target.identityCards = []
      target.identityCards.push(...identityCards)
      return this
    }

    // Pushes a new address for a worker
    addWorkerAddress(address: WorkerAddress) {
      isTargetNotDefined(this.target)

      const target = this.target as Worker

      if (!target.addresses?.length) target.addresses = []
      target.addresses?.unshift(address)
      return this
    }

  }

  export class Supabase_HCMWorkerOrganizationService extends HCMWorkerOrganizationService
    implements HCMPendingJoinRequestService<Worker> {

    constructor(
      private client?: SupabaseClientDatabase,
      private target?: Worker,
    ) {
      super()
    }

    private dependencies() {
      const target = this.target as Worker
      const client = this.client as SupabaseClientDatabase

      return { client, target }
    }

    private ensureClientAndTargetToBeDefined() {
      isClientNotUndefined(this.client)
      isTargetNotDefined(this.target)
    }
    
    // Retrieves all organizations where the worker is currently hired
    async getOrganizations(): Promise<Organization[] | undefined> {
      this.ensureClientAndTargetToBeDefined()

      const { client, target } = this.dependencies()

      const query = client.from('workerOrganizations')
        .select(`
          organizations!workerOrganizations_organizationId_fkey (
            *
          )
        `)
        .match({ workerId: target.id })
    }
    
    async getRoles(): Promise<Role[] | undefined> {
      this.ensureClientAndTargetToBeDefined()

    }
    
    async getTeams(): Promise<Team[] | undefined> {
      this.ensureClientAndTargetToBeDefined()

    }
    
    async suspend() {
      this.ensureClientAndTargetToBeDefined()

    }

    async terminate() {
      this.ensureClientAndTargetToBeDefined()

    }

    async resign() {
      this.ensureClientAndTargetToBeDefined()

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