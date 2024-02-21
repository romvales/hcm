import { createClient } from '@supabase/supabase-js'
import { afterEach, test } from 'bun:test'
import { Database } from '../../database'
import { Supabase_HCMWorkerService } from './HCMWorkerService'

const localClient = createClient<Database>(
  process.env.SUPABASE_URL as string, 
  process.env.SUPABASE_LOCAL_SERVICE_ROLE_KEY as string,
)

const LOCALCLIENT_TEST_EMAIL = 'test@romvales.com'
const LOCALCLIENT_TEST_PASSWORD = 'testpassword'

let userId: string

afterEach(async () => {
  await cleanHCMTeamServiceTest()
})

const setupHCMTeamServiceTest = async () => {
  const workerService = new Supabase_HCMWorkerService(localClient)
  const teamService = new Supabase_HCMWorkerService(localClient)

  return {
    teamService,
    workerService,
  }
}

const cleanHCMTeamServiceTest = async () => {

}


test('> HCMTeamService.createTeam()', async () => {
  const {} = setupHCMTeamServiceTest()
})

test('> HCMTeamService.getTeamById()', async () => {
  const {} = setupHCMTeamServiceTest()
})

test('> HCMTeamService.deleteTeamById()', async () => {
  const {} = setupHCMTeamServiceTest()
})

test('> HCMTeamService.saveTeam()', async () => {
  const {} = setupHCMTeamServiceTest()
})

