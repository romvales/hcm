import { SupabaseClientDatabase, Supabase_HCMWorkerService, isClientNotUndefined } from '.'
import { HCMPayrollService, Payroll } from '../../../src'


export class Supabase_HCMPayrollService extends HCMPayrollService {

  constructor(
    private client: SupabaseClientDatabase,
    private workerService?: Supabase_HCMWorkerService,
    private target?: Payroll,
  ) {
    super()
  }

  private dependencies() {
    const client = this.client
    const workerService = this.workerService as Supabase_HCMWorkerService
    const target = this.target as Payroll

    return { client, workerService, target }
  }

  private ensureClientToBeDefined() {
    isClientNotUndefined(this.client)
  }

}