import { SupabaseClientDatabase, Supabase_HCMWorkerService, isClientNotUndefined, isServiceDefined, isTargetNotDefined } from '.'
import { Addition, HCMAdditionService } from '../../../src'

export class Supabase_HCMAdditionService extends HCMAdditionService {

  constructor(
    private client: SupabaseClientDatabase,
    private workerService?: Supabase_HCMWorkerService,
    private target?: Addition
  ) {
    super()

  }

  private dependencies() {
    const client = this.client as SupabaseClientDatabase
    const workerService = this.workerService as Supabase_HCMWorkerService
    const target = this.target as Addition

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

  createAddition(name: string, value: number) {
    const add: Addition = { name, value }

    this.setTarget(add)

    return add
  }

  async getAdditionById(id: number): Promise<Addition | undefined> {
    this.ensureClientToBeDefined()

    const { client } = this.dependencies()

    const query = client.from('additions')
      .select()
      .match({ id })
      .throwOnError()

    return query
      .then(res => {
        const addition = (res.data?.at(0) ?? {}) as Addition

        if (addition?.id) {
          return addition
        }

        return
      })
  }

  async deleteAdditionById(id: number) {
    this.ensureClientToBeDefined()

    const { client } = this.dependencies()
    const query = client.from('additions')
      .delete()
      .match({ id })
      .throwOnError()

    return query
  }

  async saveAddition() {
    this.ensureClientWorkerServiceTargetToBeDefined()

    const { client, workerService, target } = this.dependencies()
    const updatorWorker = await workerService.getWorkerBySessionUser()

    // New addition
    if (!target.createdById && updatorWorker) {
      target.createdById = updatorWorker.id
    }

    const query = client.from('additions')
      .upsert(Object.assign(target))
      .select()
      .limit(1)
      .throwOnError()

    return query
      .then(res => {
        const updatedAddition = (res.data?.at(0) ?? {}) as Addition

        if (!updatedAddition?.id) return {} as Addition

        this.setTarget(updatedAddition)

        return updatedAddition
      })
  }

}