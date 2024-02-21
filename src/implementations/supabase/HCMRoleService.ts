import { SupabaseClientDatabase, Supabase_HCMWorkerService, isClientNotUndefined, isServiceDefined, isTargetNotDefined, isWorkerDefined } from '.'
import { HCMRoleService, Role } from '../../../src'

export class Supabase_HCMRoleService extends HCMRoleService {

  constructor(
    private client?: SupabaseClientDatabase,
    private workerService?: Supabase_HCMWorkerService,
    private target?: Role,
  ) {
    super()
  }

  private dependencies() {
    const target = this.target as Role
    const client = this.client as SupabaseClientDatabase
    const workerService = this.workerService as Supabase_HCMWorkerService

    return { target, workerService, client }
  }

  private ensureClientTargetWorkerServiceToBeDefined() {
    isClientNotUndefined(this.client)
    isTargetNotDefined(this.target)
    isServiceDefined(this.workerService)
  }

  private ensureClientTargetToBeDefined() {
    isClientNotUndefined(this.client)
    isTargetNotDefined(this.target)
  }

  createRole(name: string): Role {
    this.ensureClientTargetToBeDefined()
    const role: Role = { name }
    return role
  }

  async getRoleById(id: number): Promise<Role> {
    this.ensureClientTargetToBeDefined()

    const { client } = this.dependencies()

    const query = client.from('roles')
      .select()
      .match({ id })
      .limit(1)
      .throwOnError()

    return query
      .then(res => {
        const role = (res.data?.at(0) ?? {}) as Role
        return role
      })
  }

  async deleteRoleById(id: number) {
    this.ensureClientTargetToBeDefined()

    const { client } = this.dependencies()

    const query = client.from('roles')
      .delete()
      .match({ id })
      .throwOnError()

    return query
  }

  async saveRole() {
    this.ensureClientTargetWorkerServiceToBeDefined()

    const { client, target, workerService } = this.dependencies()
    const updatorWorker = await workerService.getWorkerBySessionUser()

    // New role
    if (!target.createdById && updatorWorker) {
      target.createdById = updatorWorker.id
    }

    const query = client.from('roles')
      .upsert(Object.assign(target))
      .select()
      .limit(1)
      .throwOnError()

    return query
      .then(res => {
        const updatedRole = (res.data?.at(0) ?? {}) as Role
        
        if (!updatedRole?.id) return {}

        this.setTarget(updatedRole)

        return updatedRole
      })
  }

}