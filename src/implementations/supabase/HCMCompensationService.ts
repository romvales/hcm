import { SupabaseClientDatabase, Supabase_HCMWorkerService, isClientNotUndefined, isServiceDefined, isTargetNotDefined } from '.'
import { Compensation, HCMCompensationService } from '../../../src'

export class Supabase_HCMCompensationService extends HCMCompensationService {

  constructor(
    private client: SupabaseClientDatabase,
    private workerService: Supabase_HCMWorkerService,
    private target: Compensation
  ) {
    super()
  }

  private dependencies() {
    const client = this.client
    const workerService = this.workerService as Supabase_HCMWorkerService
    const target = this.target as Compensation

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

}