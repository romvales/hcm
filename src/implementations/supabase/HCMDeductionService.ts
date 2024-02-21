import { SupabaseClientDatabase, Supabase_HCMWorkerService, isClientNotUndefined, isServiceDefined, isTargetNotDefined } from '.'
import { Deduction, HCMDeductionService } from '../../../src'

export class Supabase_HCMDeductionService extends HCMDeductionService {

  constructor(
    private client: SupabaseClientDatabase,
    private workerService?: Supabase_HCMWorkerService,
    private target?: Deduction
  ) {
    super()
  }

  private dependencies() {
    const client = this.client
    const workerService = this.workerService as Supabase_HCMWorkerService
    const target = this.target as Deduction

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


  createDeduction(name: string, value: number) {
    const deduct: Deduction = { name, value }
    this.setTarget(deduct)
    return deduct
  }

  async getDeductionById(deductionId: number): Promise<Deduction | undefined> {
    this.ensureClientToBeDefined()

    const { client } = this.dependencies()

    const query = client.from('deductions')
      .select()
      .match({ id: deductionId })
      .throwOnError()

    return query
      .then(res => {
        const deduction = (res.data?.at(0) ?? {}) as Deduction

        if (deduction?.id) return deduction

        return
      })
  }

  async deleteDeductionById(deductionId: number) {
    this.ensureClientToBeDefined()

    const { client } = this.dependencies()

    const query = client.from('deductions')
      .delete()
      .match({ id: deductionId })
      .throwOnError()

    return query
  }

  async saveDeduction() {
    this.ensureClientWorkerServiceTargetToBeDefined()

    const { client, workerService, target } = this.dependencies()
    const updatorWorker = await workerService.getWorkerBySessionUser()

    // New deduction
    if (!target.createdById && updatorWorker) {
      target.createdById = updatorWorker.id
    }

    const query = client.from('deductions')
      .upsert(Object.assign(target))
      .select()
      .limit(1)
      .throwOnError()

    return query
      .then(res => {
        const updatedDeduction = (res.data?.at(0) ?? {}) as Deduction

        if (!updatedDeduction?.id) return {} as Deduction

        this.setTarget(updatedDeduction)

        return updatedDeduction
      })
  }

}

