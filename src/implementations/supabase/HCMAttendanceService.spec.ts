import { afterEach, test } from 'bun:test'
import { Supabase_HCMAttendanceService } from './HCMAttendanceService'
import { createClient } from '@supabase/supabase-js'
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
  await cleanHCMAttendanceServiceTest()
})

const setupHCMAttendanceServiceTest = async () => {
  const workerService = new Supabase_HCMWorkerService(localClient)
  const attendanceService = new Supabase_HCMAttendanceService(localClient, workerService)

  await localClient.auth.signUp({
    email: LOCALCLIENT_TEST_EMAIL,
    password: LOCALCLIENT_TEST_PASSWORD,
  })

  const { data, error } = await localClient.auth.signInWithPassword({
    email: LOCALCLIENT_TEST_EMAIL,
    password: LOCALCLIENT_TEST_PASSWORD,
  })

  const { user } = data

  if (user?.id) userId = user.id
  if (error) throw error
  

  return {
    attendanceService,
    workerService,

    user,
  }
}

const cleanHCMAttendanceServiceTest = async () => {
  await localClient.auth.signOut()
  await localClient.auth.admin.deleteUser(userId)
}


test('> HCMAttendanceService.createAttendance()', async () => {
  const {} = await setupHCMAttendanceServiceTest()
})

test('> HCMAttendanceService.getAttendanceById()', async () => {
  const {} = await setupHCMAttendanceServiceTest()
})

test('> HCMAttendanceService.deleteAttendanceById()', async () => {
  const {} = await setupHCMAttendanceServiceTest()
})

test('> HCMAttendanceService.saveAttendance()', async () => {
  const {} = await setupHCMAttendanceServiceTest()  
})

test('> HCMAttendanceService.changeStatus()', async () => {
  const {} = await setupHCMAttendanceServiceTest()
})

test('> HCMAttendanceService.changeType()', async () => {
  const {} = await setupHCMAttendanceServiceTest()
})

test('> HCMAttendanceService.changePerfLabel()', async () => {
  const {} = await setupHCMAttendanceServiceTest()
})

test('> HCMAttendanceService.changeClockInType()', async () => {
  const {} = await setupHCMAttendanceServiceTest()
})

test('> HCMAttendanceService.changeClockOutType()', async () => {
  const {} = await setupHCMAttendanceServiceTest()
})

test('> HCMAttendanceService.clockIn()', async () => {
  const {} = await setupHCMAttendanceServiceTest()
})

test('> HCMAttendanceService.clockOut()', async () => {
  const {} = await setupHCMAttendanceServiceTest()
})

test('> HCMAttendanceService.getShift()', async () => {
  const {} = await setupHCMAttendanceServiceTest()
})

test('> HCMAttendanceService.isLate()', async () => {
  const {} = await setupHCMAttendanceServiceTest()
})

test('> HCMAttendanceService.isOverride()', async () => {
  const {} = await setupHCMAttendanceServiceTest()
})

test('> HCMAttendanceService.isHoliday()', async () => {
  const {} = await setupHCMAttendanceServiceTest()
})

test('> HCMAttendanceService.isBreak()', async () => {
  const {} = await setupHCMAttendanceServiceTest()
})