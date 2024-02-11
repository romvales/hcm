import { SupabaseClientDatabase, Supabase_HCMWorkerService, isClientNotUndefined, isServiceDefined, isTargetNotDefined, isWorkerDefined } from '.'
import { HCMRoleService, Role } from '../../worker.d'

export class Supabase_HCMRoleService extends HCMRoleService {

  constructor(
    private client?: SupabaseClientDatabase,
    private workerService?: Supabase_HCMWorkerService,
    private target?: Role,
  ) {
    super()
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

  private dependencies() {
    const target = this.target as Role
    const client = this.client as SupabaseClientDatabase
    const workerService = this.workerService as Supabase_HCMWorkerService

    return { target, workerService, client }
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

    return await query
  }

  async saveRole() {
    this.ensureClientTargetWorkerServiceToBeDefined()

    const { client, target, workerService } = this.dependencies()
    const sessionUser = await client.auth.getUser()
      .then(res => {
        if (res.error) {
          return undefined
        }

        return res.data.user
      })

    const updatorWorker = await workerService.getWorkerByUser(sessionUser)

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